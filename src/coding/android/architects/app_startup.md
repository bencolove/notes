# App Startup Library
`implementation "android.startup:startup-runtime:1.0.0"`

> X ContentProvider to init deps on startup

> O Use App Startup

Entry point of init process:
1. `InitializationProvider` in manifest
1. `<meta-data>` for all deps
1. `Initializer.dependencies()` to return all deps

```xml
<provider
    android:name="androidx.startup.InitializationProvider"
    android:authorities="${applicationId}.androidx-startup"
    android:exported="false"
    tools:node="merge">
    <!-- This entry makes ExampleLoggerInitializer discoverable. -->
    <meta-data  android:name="com.example.ExampleLoggerInitializer"
          android:value="androidx.startup" />
</provider>
```

## Manually

> Disable individual component

`tools:node="remove"` instead of `android:value="androidx.startup"`

```xml
<provider
    android:name="androidx.startup.InitializationProvider"
    android:authorities="${applicationId}.androidx-startup"
    android:exported="false"
    tools:node="merge">
    <meta-data android:name="com.example.ExampleLoggerInitializer"
              tools:node="remove" />
</provider>
```
> Disable all
```xml
<provider
    android:name="androidx.startup.InitializationProvider"
    android:authorities="${applicationId}.androidx-startup"
    tools:node="remove" />
```

> Manually init

```java
AppInitializer.getInstance(context)
    .initializeComponent(ExampleLoggerInitializer::class.java)
```