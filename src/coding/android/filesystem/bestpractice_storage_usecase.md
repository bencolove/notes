# Best Practice for Storage Access
1. media and non-media
1. all version or version specific


## Media Files
Image, video and audio files:

Usecases | Method | All version
---|---|---
list all | `MediaStore.query` collection | O
list dir | `MediaStore.query` absolute fs path | O
import exsiting image | `ACTION_GET_CONTENT` sysetm picker with user interaction | O
capture image | `ACTION_IMAGE_CAPTURE` storing into `MediaStore.Images` | O
share with other apps | `ContentResolver.insert()` into `MediaStore` | O
share with specific app | `FileProvder` and start Intent with it | O

## Non-media Files

Usecase | Method | All Version
---|---|---
open document | `ACTION_OPEN_DOCUMENT` with system picker <br> category: Intent.CATEGORY_OPENABLE <br> type: */* <br> extra: Intent.EXTRA_MIME_TYPES | O
Share content with other apps | `FileProvider` | O
cache non-media files | `Contentx.cacheDir` <br> `Context.externalCacheDir` | O


[use-cases]: https://developer.android.com/training/data-storage/use-cases