# How-to Share Data
1. *`media store`* collections
1. [[MediaStore example][mediastore-example]]
1. other types of document
1. by `content://`

---

## Media Store
>Permissions

`READ_EXTERNAL_STORAGE` AND `WRITE_EXTERNAL_STORAGE`


By `ContentResolver`
```java
val projection = arrayOf(media-database-columns-to-retrieve)
val selection = sql-where-clause-placeholder-variables
val selectionArgs = values-of-placeholder-variables
val sortOrder = sql-order-by-clause

applicationContext.contentResolver.query(
    MediaStore.media-type.Media.EXTERNAL_CONTENT_URI,
    projection,
    selection,
    selectionArgs,
    sortOrder
)?.use { cursor ->
    while(cursor.moveToNext()) {

    }
}

```

Media Collections | Locations | Table  
---|---|---
Images | `DCIM/` <br> `Pictures/` | `MediaStoer.Images`
Videos | `DCIM/` <br> `Movies/` <br> `Pictures/` | `MediaStore.Video`
Audio files | `Alarms/` <br> `Audiobooks/` <br> `Music/` <br> `Notifications/` <br> `Podcasts/` <br> `Ringtones/` <br> `Movies` playlist | `MediaStore.Audio`
Downloaded files | `Download/` | `MediaStore.Downloads` >= Android 10 (API level 29)
Scoped storage | depends | `MediaStore.Files` >= Android 10 (API level 29)

---

<br>

>QUERY  

```kotlin
// need READ_EXTERNAL_STORAGE permission for files not created by app
// Android Q (10) API level 29
val collection = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
    MediaStore.Video.Media.getContentUri(
        MediaStore.VOLUME_EXTERNAL
    )
} else {
    MediaStore.Video.Media.EXTERNAL_CONTENT_URI
}

val projection = arrayOf(
    MediaStore.Video.Media._ID,
    MediaStore.Video.Media.DISPLAY_NAME,
    MediaStore.Video.Media.DURATION,
    MediaStore.Video.Media.SIZE
)

val selection = "${MediaStore.Video.Media.CURATION} >= ?"
val selectionArgs = arrayOf(
    TimeUnit.MILLISECONDS.convert(5, TimeUnit.MINUTES).toString()
)

val sortOrder = "${MediaStore.Video.Media.DISPALY_NAME} ASC"

ContentResolver.query(
    collection,
    projection,
    selection,
    selectionArgs,
    sortOrder
)?.use { cursor ->
    // cache column indices
    val idColumn = cursor.getColumnIndexOrThrow(MediaStoer.Video.Media._ID)
    val nameColmn = ...

    while(cursor.moveToNext()) {
        // values of each column
        val id = cursor.getLong(idColumn)
        ...
        // contentUri used to access ?
        val contentUri: Uri = ContentUris.withAppendedId(
            MediaStore.Video.Media.EXTERNAL_CONTENT_URI,
            id
        )
    }
}
```

>Load

```kotlin
// load as thumbnail
val thumbnail: Bitmap = applicationContext.contentResolver.loadThumbnail(
    contentUri, Size(640, 480), null
)
// open as File
val resolver = applicationContext.contentResolver
// rwt -- truncating/overwriting existing files
val readMode = "r"
resolver.openFileDescripor(contentUri, readMode).use {
    // ParcelFileDescriptor
    pfd ->
}
// open as Stream
resolver.openInputStream(contentUri).use {
    stream ->
}
```

> Add Item

```kotlin
val resolver = applicationContext.contentResolver

// desinate the collection
val audioCollection = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.Q) {
    MediaStore.Audio.Media.getContentUri(
        MediaStore.VOLUME_EXTERNAL_PRIMARY
    )
} else {
    MediaStore.Audio.Media.EXTERNAL_CONTENT_URI
}

// insert a placeholder into the collection 
val newSongDetail = ContentValues().apply {
    put(MediaStore.Audio.Media.DISPLAY_NAME, "My Song.mp3")
}
val uriHandle = resolver.insert(audioCollection, newSongDetail)

resolver.openFileDescriptor(uriHandle, "w", null).use {
    pfd ->
    // write data
}

// finish up
uriHandle.clear()
uriHandle.put(MediaStore.Audio.Media.IS_PENDING, 0)
resolver.update(uriHandle, newSongDetail, null, null)
```

>Update Item  

Change path `mv` by using `MediaColumns.RELATIVE_PATH`.

* without `scoped storage` permission is needed to access non-owned files
* with `scoped storage` extra effort needed

```java
val mediaId = // MediaStore.Audio.Media._ID of the item
val resolver = applicationContext.contentResolver

val selection = "${MediaStore.Audio.Media._ID} = ?"

val selectionArgs = arrayOf(mediaId.toString())

val updatedSongDetails = ContentValues().apply {
    put(MediaStore.Audio.Media.DISPLAY_NAME, "My Song.mp3")
}

val numSongsUpdated = resolver.update(
    songUri,
    details,
    selection,
    selectionArgs
)
```

>Remove Item  
```java
val resolver = applicationResoler.contentResolver
// Uri
val imageUri = 
// selection
val selection = ""
val selectionArgs = ""
// delete
val numImagesRemoved = resolver.delete(
    imageUri,
    selection,
    selectionArgs
)
```


[shared-media]: https://developer.android.com/training/data-storage/shared/media
[content-provider]: https://developer.android.com/guide/topics/providers/content-provider-creating
[mediastore-example]: https://github.com/android/storage-samples/tree/main/MediaStore