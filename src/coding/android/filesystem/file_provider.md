# File Provider

>Background:

From >= Android 7.0 (API level 24), files denoted `file:///` is not permitted instead use `content://` with ContentResolver to load resources.

In order to provide app specific file access to other apps, use `FileProvider` defined in *`manifest`*.

>NOTE:  

`FileProvider.getUriForFile` works for app-specific data and external public as follows:
1. internal: /data/data/*pkg*
1. external: /storage/emulated/0/Android/data/*pkg*
1. external public: /storage/emulated/0 

## How it works
`FileProvider` builds mapping from *< name >* to
1. *< files-path >*/*path* -> `context.filesDir`/*path*
1. *< cache-path >*/*path* -> `context.cacheDir`/*path*
1. *< external-files-path >*/*path* -> `context.getExternalFilesDirs()[0]`/*path*
1. *< external-cache-path >*/*path* -> `context.getExternalCacheDirs()[0]`/*path*
1. *< external-path >*/*path* -> `Environment.getExternalStorageDirectory()`/*path*

when calling `FileProvider.getUriForFile(context, <pkg>.fileprovider, file)`, it loops the mappings searching for most matched against path and generate content Uri as `content://<pkg>.fileprovider/<name>/<path>`.

```java
val file = File(Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS),"report.xls")
val uri = FileProvider.getUriForFile(activity, "${context.packageName}.fileprovider", file)

// result
// file: /storage/emulated/0/Download/report.xls
// uri: content://io.ionic.starter.fileprovider/downloads/report.xls
```
As seen, `Environment.getExternalStoragePublicDirectory(Environment.DIRECTORY_DOWNLOADS), "file.txt")` will be matched `content://<authorities>/downloads` from `/storage/emulated/0` 

## Configure Files
1. Manifest.xml
1. res/xml/file_paths.xml

```xml
<manifext xmlns...>
    <appliction ...>
        <provider
            android:name="androidx.core.content.FileProvider" 
            // or android.support.v4.content.FileProvider
            android:authorities="${applicationId}.fileprovider"
            // no export fileprovider
            android:exported="false"
            // grant perm on uri when request
            android:grantUriPermissions="true"
        >
            <meta-data
                android:name="android.support.FILE_PROVIDER_PATHS"
                <!-- res/xml/file_paths.xml -->
                android:resource="@xml/file_paths"
            />
```

>Configure Shared Folders 

-- *`res/xmnl/file_paths.xml`* --
```xml
<paths xmlns:android="http://schemas.android.com/apk/res/android"> 
    <files-path name="my_images" path="images/"/> 
    ...
</paths>
```

Means `content://< appid>.fileprovider/my_images` will be searched for within `Context.getFilesDir() + "/images"`.

Xml Elem | Device Folder | Android Method
---|---|---
*files-path* | `/data/data/<pkg>/files` | `context.filesDir`
*cache-path* | `/data/data/<pkg>/cache` | `context.cacheDir`
*external-path* | `/storage/emulated/0` | `Environment.getExternalStorageDirectory()`
*external-files-path* | `/storage/emulated/0/Android/data/<pkg>/cache` | `context.externalFilesDir`
*external-cache-path* | `/storage/emulated/0/Android/data/<pkg>/cache` | `context.externalCacheDir`

## Provide External Access
```java
val file = File(context.filesDir, "images").let {
    File(it, "img.jpg")
}
val contentUri = FileProvider.getUriForFile(activityContext, "<appid>.fileprovider", file)

Intent().apply {
    setDataAndType(contentUri, contentResolver.getType(contentUri))
    addFlags(Intent.GLAG_GRANT_READ_URI_PERMISSION)
}
```

[declare]: https://developer.android.com/training/secure-file-sharing/setup-sharing