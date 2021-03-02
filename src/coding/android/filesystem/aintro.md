# Intro to Android File System
1. app-specific storage
1. shared storage
1. preferences
1. databases

## Sotrage Locations
* internal (always available at `/data/data`)
* external (removable like SD card at `/sdcard`)
* they may vary from device to device

*APK* itself by default is stored internally, but can be changed by:
```xml
<manifext ...
    android:installLocation="preferExternal"
>
```

## The Scoped Storage
1. app-specific dirs
1. `MediaStore`

## the Bundled Resources
Access file from the bundled APK
```java
// from /res/raw folder
openRawResource(R.raw.resid)
```

## the App-specific Directories
[[Didicated folders][app-specific-volumes]] for files and caches obtained by:
1. `filesDir` and `cachesDir` on internal (`Security library` or >= Android 10 to encrypt)
1. `` on external
1. predefined external directory [[names][predefined-subdirs]]

### the App-specific Storage on Internal
Files within these those two locations on internal can only be shared by using `FileProvider` with `FLAG_GRANT_READ_URI_PERMISSION` attribute.

-- *persist files* --
```java
// File
val file = File(context.filesDir, filename)

// FileOutputStream
context.openFileOutput(filename, Context.MODE_PRIVATE).use {
    it.write(...contents...)
}

context.openFileInput(filename).bufferedReader().useLines {
    lines -> lines.fold("") {
        some, text -> "$some\n$text"
    }
}

// list files
var files: Array<String> = context.fileList()
// create folder automatically
context.getDir(dirName, Context.MODE_PRIVATE)
```

-- *cache files* --
```java
// query free cache space before writing
context.getCacheQuotaBytes()
// create cache file
File.createTempFile(filename, null, context.cacheDir)
// access
var cacheFile = File(context.cacheDir, filename)
// remove
cacheFile.delete()
context.deleteFile(cacheFileName)
```

### the App-specfic Storage on External
on **`<= Android 9(API level 28)`**, files can be accessed with appropriate permissions.
on **`>= Android 10(API level 29)`**, *`scoped storage`* should be used.

`$ adb shell sm set-virtual-disk true` to enable virtual removable volume for testing.

Query before access:
```java
// Readable and writable
val externalWritable = Environment.getExternalStorageState() == Environment.MEDIA_MOUNTED
val externalReadable = Environment.getExternalStorageState() in setOf(Environment.MEDIA_MOUNTED, Environment.MEDIA_MOUNTED_READ_ONLY)
// primary external storage
val externalVolums: Array<out File> =
ContenxtCompat.getExternalFilesDirs(context, null)
val primaryExternalVolume = externalVolumes[0]

// access
val appSpecificExternalDir = File(context.getExternalFilesDir(), filenname)
val appSpecificCacheFile = File(context.externalCacheDir, filename)
appSecificCacheFile.delete()

// large files like media on external within specific folder
val file = File(context.getExternalFilesDir(Environment.DIRECTORY_PICTURES), albumName)
if (!file?.mkdirs()) { ...error create dir... }
```

Query and remove cache files:
```java
const val NUM_BYTES_NEEDED = 1024*1024*10L;

val storageManager = applicationContext.getSystemService<StorageManager>()!!
val appSpecificInternalDirUuid: UUID = storageManager.getUuidForPath(filesDir)
val availableBytes: Long = storageManager.getAllocatableBytes(appSpecificInternalDirUuid)
if (avaialbleBytes >= NUM_BYTES_NEEDED) {
    // do the alloc
    storageManager.allocateBytes(appSpecificInternalDirUuid, NUM_BYTES_NEEDED)
} else {
    // do the clean
    // StorageManager.getFreeBytes() / StorageManager.getTotalBytes()
    val storageIntent = Intent().apply {
        action = ACTION_MAMAGE_STORAGE
    }
    // clean cache files
    Intent().apply {
        action = ACTION_CLEAR_APP_CACHE
    }
}
```


## the Shared Storage
1. *media colleciton*

## Common Directories

Method | Folder | Deprecated
---|---|---
context.filesDir | `/data/user/0/io.ionic.starter/files` |
context.cacheDir | `/data/user/0/io.ionic.starter/cache` |
context.externalCacheDir | `/storage/emulated/0/Android/data/io.ionic.starter/cache` |
context.getExternalFilesDir(null) | `/storage/emulated/0/Android/data/io.ionic.starter/files` |
Environment.getRootDirectory() | `/system` | 
Environment.getDataDirectory() | `/data` |
Environment.getDownloadCacheDirectory() | `/cache`
Environment.getExternalStorageDirectory() | `/storage/emulated/0` | X
Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS) | `/storage/emulated/0/Download` | X


## App specific folders
App's package name is like Linux uid, uniquely isolated file access permissions.

`Context.openFileOutput(fileName, Context.MODE_PRIVATE)` will automatically open (create if not existed) in `/data/data/<package>/files` folder.
Context.Mode | Value | Meaning
---|---|---
MODE_PRIVE | 0 | Default, app specific, overwrite
MODE_APPEND | 32768 | create or append
MODE_WORLD_READABLE | 1 | accessible for other apps
MODE_WORLD_WRITEABLE | 2 |


[app-specific-volumes]: https://developer.android.com/training/data-storage/app-specific#kotlin
[predefined-subdirs]: https://developer.android.com/reference/android/os/Environment#fields