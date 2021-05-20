# Fine Tune of Local State
Tools:
1. _`ValueNotifier`_ and _`ValueListenableBuilder`_

## ValuelNotifier and ValueListenableBuilder
The key of reuse already built widgets is: `ValueListenableBuilder.child`.
It will pass the built **`child`** widgets of last run (before value changed) and make the most out of it.

How to:
```java
ValueNotifier<Type>

ValueListenableBuilder(
    valueListenable: valueNotifier,
    child: initChildMayBeNull,
    builder: (context, valueInterested, old) {
        /*
        The old (Widget?) passed-in is the Widget built from last round (before changed) 
        */
    }
)

```