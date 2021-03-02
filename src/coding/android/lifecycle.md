# Lifecycle and Live Data
`Lifecycle` = `Observable` on STATE changes  
`LiveData` = `Lifecycle` + `Observable`

LiveData only pushes updates to components that are *`ACTIVE`*.

`LiveData.register(Observer, lifecycle)` will remove the observer when associated `lifecycle` *`DESTROYED`*.

>1. Create LiveData Objects

```java
class NameViewModel : ViewModel() {

    // Create a LiveData with a String
    val currentName: MutableLiveData<String> by lazy {
        MutableLiveData<String>()
    }

    // Rest of the ViewModel...
}
```

>2. Observe LiveData Objects

```java
// put init in app component's onCreate() method


```

[livedata]: https://developer.android.com/topic/libraries/architecture/livedata