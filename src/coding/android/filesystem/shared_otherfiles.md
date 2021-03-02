# System Picker 
Aka *`Storage Access Framework`*, involves *`document provider`* and **`USER`** involved.

* \>= Android version 4.4 (API level 19)
* `READ_EXTERNAL_STORAGE` permission needed
* workflow:
    1. `Intent` with resource location Uri and supported action
    1. User interact with *`system picker`* to browse and choose
    1. granted permission to act upon


## Usecase 1. create a new file
Like *`save as`* workflow:
```java
// Request code
const val CREATE_FILE = 1

val intent = Intent(Intent.ACTION_CREATE_DOCUMENT).apply {
    addCategory(Intent.CATEGORY_OPENABLE)
    type = "application/pdf"
    putExtra(Intent.EXTRA_TITLE, "invoice.pdf")

    // optionally specify directory Uri to open into
    putExtra(DocumentsContract.EXTRA_INITIAL_URI, folderUri)
}
startActivityForResult(intent, CREATE_FILE)
```

## Usecase 2. open a file
```java
// custom request code
const val PICK_PDF_FILE = 2

val intent = Intent(Intent.CTION_OPEN_DOCUMENT).apply {
    addCategory(Intent.CATEGORY_OPENABLE)
    type = "application/pdf"
    putExtra(DocumentsContract.EXTRA_INITIAL_URI, folderUri)
}
startActivityForResult(intent, PICK_PDF_FILE)
```

## Usecase 3. operates on the selected file
>1. get back Uri of user selection

```java
override fun onActivityResult(
    requestCode: Int, resultCode: Int, resultData: Intent?
) {
    if (requestCode == your-request-code &&
        resultCode == Activity.RESULT_OK
    ) {
        resultData?.data?.also { uri ->
            // the Uri of file or dir use selected
        }
    }
}
```

>2. permissions

When a file is picked by user, a Uri permission lasting until device restart is granted to app unless the app "take" them:
```java
val contentResolver = applicationContenxt.contentResolver

val takeFlags: Int = Intent.FLAG_GRANT_READ_URI_PERMISSION or Intent.FLAG_GRANT_WRITE_URI_PERMISSION

// check for them
contentResolver.takePersistableUriPermission(uri, takeFlags)
```

>3. examine document metadata

```java
val cursor: Cursor? = contentResolver.query(
    uri, null, null, null, null, null
)?.use {
    // true only it exists
    if (it.moveToFirst()) {
        val displayName: String = it.getString(it.getColumnIndex(OpenableColumns.DISPLAY_NAME))
        // check if size is unknow before assignment
        val sizeIndex: Int = it.getColumnIndex(OpenableColumns.SIZE)
        val size: String = if (!it.isNull(sizeIndex)) {
            it.getString(sizeIndex)
        } else {
            "unknown"
        }
    }
}
```

>4. open a document
```java
// by file descriptor
val parcelFileDescriptor = contentResolver.openFileDescriptor(uri, "r")
val fileDescriptor = pacelFileDescriptor.fileDescriptor
val bitmap: Bitmap = BitmapFactory.decodeFileDescriptor(fileDescriptor)
parcelFileDescriptor.close()

//. by file input stream
contentResolver.openInputStream(uri)?.use {
    inputStream ->
        BufferReader(InputStreamReader(inputStream)).use {
            reader ->
                var line: String? = reader.readLine()
                while (line != null) {
                    line = reader.readLine()
                }
        }
}
```

>5. write a document

check `Document.COLUMN_FLAGS` with `FLAG_SUPPORTES_WRITE`

```java
// check Document.COLUMN_FLAGS with FLAG_SUPPORTES_WRITE
DocumentsContract.isDocumentUri(context, uri)

val cursor: Cursor? = contentResovler.query(
    uri,
    arrayOf(DocumentsContract.Document.CONLUMN_FLAGS),
    null, null, null
)

val flags: Int = cursor?.use {
    if (cursor.moveToFirst()) {
        cursor.getInt(0)
    } else { 0 }
} ?: 0

flags and DocumentsContract.Document.FLAG_SUPPORTS_WRITE != 0

try {
    contentResolver.openFileDescriptor(uri, "w")?.use {
        FileOutputStream(it.fileDescriptor).use {
            it.write("balar".toByteArray())
        }
    }
} catch (e: FileNotFoundException) {}
catch (e: IOException) {}
```