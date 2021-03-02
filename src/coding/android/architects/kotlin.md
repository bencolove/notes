# Architecture Components with Kotlin
[[tutor][tutor]]

>1. Add KTX dependencies
* ViewModelScope
    `androidx.lifecycle:lifecycle-viewmodel-ktx:2.2.0`
* LifecycleScope
    `androidx.lifecycle:lifecycle-runtime-ktx:2.2.0`
* liveData
    `androidx.lfecycle:lifecycle-livedata-ktx:2.2.0`


## lifecycle-aware Coroutine Scopes

CoroutineScope | Container | Boundaries | Get
---|---|---|---
*ViewModel.viewModelScope* | within `ViewModel()` | canceled when ViewModel cleared | *`viewModelScope`*
*View.viewLifecycleOnwer.lifecycleScope* <br> *View.lifecycleScope*| `Lifecycle` with View | canceled when `Lifecycle` destroyed | `lifecycle.coroutineScope` <br> `lifecycleOwner.lifecycleScope`

## Suspend( await ) fro lifecycle-aware events

```java
View
    init {
        lifecycleScope.launch {
            // suspend function, awaits till complete
            whenStarted {

            }
        }
    }

// or
View
    init {
        // just like .launch{ whenStarted { ... } }
        lifecycleScope.launchWhenStarted {
            try {
                // suspend function calls
            } finally {
                // CancellationException
                if (lifecycle.state >= STARTED) { ... }
            }
        }
    }

```


[tutor]: https://developer.android.com/topic/libraries/architecture/coroutines