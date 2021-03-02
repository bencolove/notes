## Basics

An Android app runs in its own linux process of VM with unique linux ID unknown to app before hand. Differenct app would be able to share files and process (with same linux ID) signed with same certificate.


### Four Types of Building-block Components
1. Activity
1. Service
1. Broadcast Receiver
1. Content Provider

### Activity
An independent screen of entry points from an App.

Responsibility:
Keeping user states before transmitting to antother actitivity or after coming back from another.

### Service
Keeping an app running in the background without user interface. A service can be bound to an activity to interact between each other.
Types:
1. Started Service
    1. User Aware
    1. User Non-aware (GCed anytime)
1. Bound Service

>Bound Service  
Component-A (in its own process-A) depends on Service-B, Android will start the Service-B in another own process and try to keep it running.


## Broadcast Receiver
Android delivers events to app and may instantiate a Broadcast Receiver without interface but a status bar notification.

A receiver may react to the interesting events and schedule a `JosbService` with `JobScheduler`.

The message payload those events carry are `Intent`s.

## Content Provider
Abstract layers that provide
access to shared set of app data like files, SQLite databases, web resources.

A *content provider* like system-provided *Contracts* comes with may restricted CURD actions with named data assigned as URI like *`ContactsContract.data`*.

While Activity, Service, Broadcast Receiver are activated by **`Intent`**, 
content providers are activated by **`Content Resolver`**.


## Component Activation

* `startActivity()` or `startActivityForResult` with `Intent`
* Android >= 5.0 (API level 21) `JobScheduler`
* Android < 5.0 `startService` or `bindSercie` with `Intent`
* `sendBroadcast`, `sendOrderBroadcast` or `sendStickyBroadcast` with `Intent`
* `query` on a `ContentResolver`


>Andoird(>5.0 API level>21) use [[JobScheduler][JobScheduler]] to schedule *`actions`* for battery-constrainted devices with the [[Doze][Doze]] API

[JobScheduler]: https://developer.android.com/reference/android/app/job/JobScheduler
[Doze]: https://developer.android.com/training/monitoring-device-state/doze-standby
[content-provider]: https://developer.android.com/guide/topics/providers/content-providers