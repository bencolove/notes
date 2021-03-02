# Save Key-value Data by `SharedPreferences` API
Preference files:
1. one for entire app
1. one for single Activity
1. shared by name among all Activities

Preference File | Method | Purpose
---|---|---
app | `getDefaultSharedPreferences()` | one for entire app
activity | `getPreferences()` | one for Activity
shared by all actitities | `getSharedPreferences()` | multiple by name

## Read
```java
val sharedPref = actitivity?.getPreferences(Context.MODE_PRIVATE) ?: return
sharedPref.getInt(keyName, defaultValue)
```

## Write
```java
val sharedPref = actitivy?.getPreferences(Context.MODE_PRIVATE) ?: return
with (sharedPref.edit()) {
    putInt(getString(prefName), defaultValue)
    // write in-memory immediately and to disk async
    apply()
    // write to disk sync, avoid calling it on UI thread
    commit()
}
```

[shared-preferences]: https://developer.android.com/training/data-storage/shared-preferences