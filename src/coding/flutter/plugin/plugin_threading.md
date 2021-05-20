# Plugin threading

## Android
>To UI thread

```java
Handler(Looper.getMainLooper()).post {
    // Runnable on UI thread
}
```
