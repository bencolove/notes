# Intent
[[doc][doc]]  
`Intent` as message between actors:
1. Activity, 
1. Service, 
1. Broadcast

Types:
1. Explicit with target/destination
1. Implicit system find and match(manifest.intent-filter)
![implicit intents](../res/intent/intent-filters_2x.png)

## Activity

`startActvity(intent)`  
`startActivityForResult(intent)` -> `onActivityResult()`

## Service

> 5.0 (API-21), `JobScheduler`

## Broadcast

`sendBroadcast()` or `sendOrderedBroadcast()`


## Properties
Fields | Purpose
---|---
Component name | supplier reacting to requested Intent
Action | generic names: <br> 1. `ACTION_VIEW` <br> 2. `ACTION_SEND` <br> 3. ... 
Data | URI and MIME type: <br> *`content://`* means data on the device controlled by `ContentProvider`
Category | `CATEGORY_BROWSABLE` <br> `CATEGORY_LAUNCHER`
Extras | key-value pair or `Bundle`
Flags | how to launch target component


## Sending Out
> 1. user may choose default app
```java
// Create the text message with a string.
val sendIntent = Intent().apply {
    action = Intent.ACTION_SEND
    putExtra(Intent.EXTRA_TEXT, textMessage)
    type = "text/plain"
}

// Try to invoke the intent.
try {
    startActivity(sendIntent)
} catch (e: ActivityNotFoundException) {
    // Define what your app should do if no activity can handle the intent.
}
```

> 2. Force use to choose each time
```java
val sendIntent = Intent(Intent.ACTION_SEND)
...

// Always use string resources for UI text.
// This says something like "Share this photo with"
val title: String = resources.getString(R.string.chooser_title)
// Create intent to show the chooser dialog
val chooser: Intent = Intent.createChooser(sendIntent, title)

// Verify the original intent will resolve to at least one activity
if (sendIntent.resolveActivity(packageManager) != null) {
    startActivity(chooser)
}
```

## Recieve In
Intent Filter declares what action and data the app will accepts and reacts to from other apps on device.

Multiple filters act like `or` combination.

Disable by:
1. no intent filters declared
1. `exported` to false 

```java
<activity android:name="ShareActivity">
    <intent-filter>
        <action android:name="android.intent.action.SEND"/>
        <category android:name="android.intent.category.DEFAULT"/>
        <data android:mimeType="text/plain"/>
    </intent-filter>
</activity>
```

## Pending Intent
For:
1. Notification (NotificationManager executes Intent)
1. App Widget (Home screen app executes Intent)
1. Alarm (AlarmManager executes Intent)

## Intent Resolution
Test target | Intent Multicity | Filter Multicity | Rules | Default
---|---|---|---|---
Action | 1 | n | one | fail 
Data (bot URI and data type) | 1 | n | fail 
Category | n | n | all | pass

>For data, `content://` and `file://` are always supported


>Examples
```xml
<intent-filter>
    <data android:mimeType="image/*" />
    ...
</intent-filter>
```
Supports local image file retrieval

```xml
<intent-filter>
    <data android:scheme="http" android:mimeType="video/*" />
    ...
</intent-filter>
```

Supports internet video file retrieval
<br>

>`PackageManager` find matched intents programatically
1. query...()
1. resolt...()

[doc]: https://developer.android.com/guide/components/intents-filters