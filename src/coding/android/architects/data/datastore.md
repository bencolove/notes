# DataStore Solution
`androidx.datastore`

```java
// Preferences DataStore (SharedPreferences like APIs)
dependencies {
  implementation "androidx.datastore:datastore-preferences:1.0.0-alpha06"

  // optional - RxJava2 support
  implementation "androidx.datastore:datastore-preferences-rxjava2:1.0.0-alpha06"

  // optional - RxJava3 support
  implementation "androidx.datastore:datastore-preferences-rxjava3:1.0.0-alpha06"
}
// Alternatively - use the following artifact without an Android dependency.
dependencies {
  implementation "androidx.datastore:datastore-preferences-core:1.0.0-alpha06"
}
```

```java
// Typed DataStore (Typed API surface, such as Proto)
dependencies {
  implementation "androidx.datastore:datastore:1.0.0-alpha06"

  // optional - RxJava2 support
  implementation "androidx.datastore:datastore-rxjava2:1.0.0-alpha06"

  // optional - RxJava3 support
  implementation "androidx.datastore:datastore-rxjava3:1.0.0-alpha06"
}
// Alternatively - use the following artifact without an Android dependency.
dependencies {
  implementation "androidx.datastore:datastore-core:1.0.0-alpha06"
}
```


Koroutine based key-value pairs or typed objects storage solution with `protocol buffers`.

>Should migrating `SharedPreferences` to `DataStore`

>`Room` suits more for large and complex datasets

Implementations:
1. key-value pairs: Preferences DataStore
1. Typed Objects: Proto DataStore on `protocol buffers`

## Preference DataStore
```java
val prefStore: DataStore<Preferences> = context.createDataStore(name = "settings")

// read as Flow
val COUNTER_KEY = intPreferencesKey("key_name")
val countFlow: Flow<Int> = prefStore.data.map {
    it[COUNTER_KEY] ?: 0
}

// write
suspend fun updateCounter() {
    prefStore.edit { settings ->
        // single transaction
        val curValue = settings[COUNTER_KEY] ?: 0
        settings[COUNTER_KEY] = curValue + 1
    }
}
```

## Proto DataStore
* protocol buffers
* schema
* rebuild needed

--_`app/src/main/proto/`_--
```text
syntax = "proto3";

option java_package = "com.example.application";
option java_multiple_files = true;

message Settings {
    int32 example_counter = 1;
}
```

```java
// 1. ser/deser for out typed object
object SettingsSerializer : Serializer<Settings> {
  override val defaultValue: Settings = Settings.getDefaultInstance()

  override fun readFrom(input: InputStream): Settings {
    try {
      return Settings.parseFrom(input)
    } catch (exception: InvalidProtocolBufferException) {
      throw CorruptionException("Cannot read proto.", exception)
    }
  }

  override fun writeTo(
    t: Settings,
    output: OutputStream) = t.writeTo(output)
}

// 2. create
val settingsDataStore: DataStore<Settings> = context.createDataStore(
    fileName = "settings.pb",
    serializer = SettingsSerializer
)

// read
val exampleCounterFlow: Flow<Int> = settingsDataStore.data
  .map { settings ->
    // The exampleCounter property is generated from the proto schema.
    settings.exampleCounter
  }

// write atomic read-write-modify
suspend fun incrementCounter() {
  settingsDataStore.updateData { currentSettings ->
    currentSettings.toBuilder()
      .setExampleCounter(currentSettings.exampleCounter + 1)
      .build()
    }
}
```


