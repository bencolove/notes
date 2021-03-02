# Share Simle Data by Intents
1. ShareSheet
1. Intent Resolver
1. [[Recieving data][receive]]

## ShareSheet
`Intent.ACTION_SEND` with `Intent.createChooser()`
1. url
1. email
1. binary(image)

```java
// url
val sendIntent: Intent = Intent().apply {
    action = Intent.ACTION_SEND
    putExtra(Intent.EXTRA_TEXT, "url")
    type = "text/plain"
}

val shareIntent = Intent.createChooser(sendIntent, null)
startActivity(shareIntent)

// e-mail
EXTRA_EMAIL, String[]
EXTRA_CC, String[]
EXTRA_BCC
EXTRA_SUBJECT

// image, Intent.EXTRA_STREAM with Uri(not data itself)
// permission required
val shareIntent = Intent().apply {
    action = Intent.ACTION_SEND
    putExtra(Intent.EXTRA_STREAM, uriToImage)
    type = "image/jpeg"
}
startActivity(Intent.createChooser(shareIntent, "title"))
```

## Handling Temporary Per-Uri Permission
1. `Uri` from
    1. [FileProvider](./file_provider.md) kind of `ContentProvider`
    defined in *Manifest.xml*(with `getUriForFile()`)
    1. `MediaStore` (with `MediaStore.Files`)
    `scanFile()` -> `onScanCompleted()`
1. grant by `Intent.setFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION)`

## Mime Types
Sending | Recieving for
---|---
*text/plain* <br> *text/rtf* <br> *text/html* <br> *text/json* | *text/**
*image/jpg* <br> *image/png* <br> *image/gif* | *image/**
*video/mp4* <br> *video/3gp* | *video/**
*application/pdf* | same

### Sending Mulpitle Pieces
 ```java
 val imageUris: ArrayList<Uri> = arrayListOf(
        // Add your image URIs here
        imageUri1,
        imageUri2
)

val shareIntent = Intent().apply {
    action = Intent.ACTION_SEND_MULTIPLE
    putParcelableArrayListExtra(Intent.EXTRA_STREAM, imageUris)
    type = "image/*"
    setFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION)
}
startActivity(Intent.createChooser(shareIntent, "Share images to.."))
```

### Rich Preview with Sharesheet (>= Andoird 10 Q API 29)
![rich contents](../res/files/sharesheet_richcontent.png)
```java
 val share = Intent.createChooser(Intent().apply {
      action = Intent.ACTION_SEND
      putExtra(Intent.EXTRA_TEXT, "https://developer.android.com/training/sharing/")

      // (Optional) Here we're setting the title of the content
      putExtra(Intent.EXTRA_TITLE, "Introducing content previews")

      // (Optional) Here we're passing a content URI to an image to be displayed, better off placed in <cache-path>
      data = contentUri
      // (to-do) Add a relevant thumbnail via ClipData.
      flags = Intent.FLAG_GRANT_READ_URI_PERMISSION
  }, null)
  startActivity(share)
  ```

  ### Includes Custom Targets and Excludes
  ```java
  val share = Intent().createChooser(shareIntent, null).apply {
      putExtra(Intent.EXTRA_CHOOSER_TARGETS, [])
      putExtra(Intent.EXTRA_INITIAL_INTENTS, [extra_action_intent])
      putExtra(Intent.EXTRA_EXCLUDE_COMPONENTS, [excluded_array])
  }
  ```

### Notified with Share Results
Hook used to be notified when *Sharesheet* is completed with target:
```java
var shareIntent = Intent(Intent.ACTION_SEND)

val pi = PendingIntent.getBroadcast(context, requestCode, 
    Intent(context, MyBroadcastReceiver::class.java),
    PendingIntent.FLAG_UPDATE_CURRENT)

val share = Intent.createChooser(share, null, pi.intentSender)

// callback in MyBroadcastReceiver
override fun onReceive(context: Context, intent: Intent) {
    val clickedComponent: ComponentName = intent.getParcelableExtra(EXTRA_CHOSEN_COMPONENT)
}
```

## Intent Resolver
Without `Intent.createChooser()` wrapping sharing `Intent`:
```java
val sendIntent: Intent = Intent().apply {
    action = Intent.ACTION_SEND
    putExtra(Intent.EXTRA_TEXT, "This is my text to send.")
    type = "text/plain"
}
startActivity(sendIntent)
```
[send]: https://developer.android.com/training/sharing/send#kotlin
[receive]: https://developer.android.com/training/sharing/receive
