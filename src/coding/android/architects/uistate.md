# Save State
Considerations:
1. speed of retrieval
1. memory usage

Strategy of saving UI states:
1.
1.

Tools:
1. ViewModel in memory
1. onSaveInstanceState() onto disk

When states should remain
>Uses initiated
1. press back button
1. swip it off on the screen
1. navigate up from it
1. kill the app from Settings screen
1. any other kinds of finishing
>System initiated
1. rotation
1. switch into multi-window mode

>Saved State module for ViewModel
[[guide][save-viewmodel]]

1. implement `SavedStateRegistry.SavedStateProvider`
```java
class MyComponent: SavedStateRegistry.SavedStateProvider {

    companion object {
        private const val QUERY = "query"
    }

    override fun saveState(): Bundle {
        return bundleOf(QUERY to dataToSave)
    }
}
```

2. register by `SavedStateRegistry.registerSavedStateProvider`

3. 

[save-viewmodel]: https://developer.android.com/topic/libraries/architecture/viewmodel-savedstate