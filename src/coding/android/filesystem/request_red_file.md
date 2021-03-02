# Request Files and Read Files
1. app-specific files, `File`
1. external public files, `File`
1. Uri returned by server app, `Uri`

## Getting `File` by Absolute Path

## Request Files for Uri
For access to files owned by other app, request them by using the app's `ContentProvider` if any (like `FileProvider`). Steps are:
1. client app request a file by sending `Intent`
1. android ask a server app (with `ContentProvider`) to respond to it
1. the server app (should interact with user on UI????) returns a `Uri` with temprary permission
1. the client app access the file

## Access File by Uri
Once use selected a file, then a `Uri` is returned and permission is grantted:
```java
// server app grant permission by
// 1. set result
// 2. finish()
overriede fun onActivityResult(requestCode: Int, resultCode: Int, returnIntent: Intent) {
    // check if granted
    if (resultCode != Activity.RESULT_OK) {
        // user did not select
        return
    } else {
        // get Uri
        returnIntent.data?.also { returnUri ->
            // ParcelFileDescriptor
            val inputPFD = try {
                contentResolver.openFileDescriptor(returnUri, "r")
            } catch (e: FileNotFoundException) {
                return
            }
            // file descriptor as for FileInput/OutputStream
            val fd = inputPFD.fileDescriptor

            // Mime type by ContentResolver
            val mimeType: String? = contentResolver.getType(returnUri)

            // meta data like name and size
            contentResolver.query(returnUri, null, null, null, null) ?.use {
                cursor ->
                    val nameIndex = cursor.getColumnIndex(OpenableColums.DISPLAY_NAME)
                    val sizeIndex = cursor.getColumnIndex(OpenableColumns.SIZE)
                    cursor.moveToFirst()
                    val name = cursor.getString(nameIndex)
                    val size = if (cursor.isNull(sizeIndex)) "Unknown" else cursor.getLong(sizeIndex).toString()
            }
        }
    }
}
```
