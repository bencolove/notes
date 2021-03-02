# Internal & External Storage
[[blog](http://huzhengyu.com/2019/08/10/storage/)]

1. internal storage (ROM.internal)
1. external storage (ROM.external + removable)
1. removable (SDCard + OTG/USB)

Mounted Points:
1. internal: `/data/user/0` -> `/data/data`
1. external: `/storage/emulated/0` -> `/sdcard/`

## Internal Storage
Always available app-specific (Linux user space) `rw-rw----`
`/data/user/0` will be mounted to `/data/data`

Internal Storage Folders | Method | Meaning | Mounted Point
---|---|---|---
/data/user/0/*pkg* | `Context.getDataDir()` | root | /data/data/*pkg*
/data/user/0/*pkg*/files | `Context.getFilesDir()` | data | /dta/data/*pkg*/files
/data/user/0/*pkg*/databases | 
/data/user/0/*pkg*/shared_prefs |
/data/user/0/*pkg*/cache | `Context.getCacheDir` | cache, deleted any time | /data/data/*pkg*/cache
/data/user/0/*pkg*/*app-custom* | `Context.getDir(custom)` | custom files | /data/data/*pkg*/app_custom

<br>

>Access method 1.
```java
// file in /files
Context.openFileOutput(filename, Context.MODE_PRIVATE)
// cache /cache
File.createTempFile(cacheFilename, null, context.getCacheDir())
```

## External Storage
Globally visible starting with `/storage`:
* `/storage/emulated/0`, mounted to `/sdcard/`
* /storage/sdcardname

Types:
1. app specific
1. global (`WRITE_EXTERNAL_STORAGE` needed)

### Readable and Writetable
Not reliable, check everytime before access:
```java
// writetable
val externalStorageWritetable = Environment.getExternalStorageState().equals(Environment.MEDIA_MOUNTED)
// readable
val externalStorageReadable = Environment.getExternalStorageState() in setOf(Enviornment.MEDIA_MOUNTED, Environment.MEDIA_MOUNTED_READ_ONLY)
```

### External Private Storage (App Specific)
Removed when uninstalled

External App Specific Storage | Method | Meaning | Mounted Point
---|---|---|---
/storage/emultated/0/Android/data/*pkg*/`files/Music` | `Context.getExternalFilesDir(Environment.DIRECTORY_MUSIC)` | external music, pitures, download folders | /sdcard/Android/data/*pkg*/files/Music
/storage/emulated/0/Android/data/*pkg*/`cache` | `Context.getExternalCacheDir` | external cache | /sdcard/Android/data/*pkg*/cache
/storage/emulated/0/Android/`media/*pkg*` <br> /storage/`sdcard_name`/Android/`media/*pkg*`| `Context.getExternalMediaDirs` | sdcard folders | /sdcard/Android/media

```java
// create external private folder
File(context.getExternalFilesDir(Environment.DIRECTORY_PICTURES), albumName).also {
    if (!it.mkdirs()) {
        // error create folder in /storage/emulated/0/Android/
    }
}
```

### External Public Storage

External Public | Method | Mounted Point
---|---|---
/storage/emulated/0 | `Environment.getExternalStorageDirectory()` | /sdcard/
/storage/emulated/0/Pictures | `Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_PICTURE)` | /sdcard/Pictures/
/system | `Environment.getRootDirectory()`
/cache | `Environment.getDownloadCacheDirectory()`
/data | `Environment.getDataDirectory()`

<br>

---
### Clean up **DATA** and **CACHE** from AppManager
1. clean **DATA**
    * /data/user/0/*pkg*/
    * /storage/emulated/0/Android/data/*pkg*/
1. clean **CACHE**
    * /data/user/0/*pkg*/cache
    * /storage/emulated/0/Android/data/*pkg*/cache