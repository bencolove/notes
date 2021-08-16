# Focus Management

[[example](https://segmentfault.com/a/1190000023101716)]

1. Widget on focus:
    keyboard popup
1. Widget loses focus:
    keyboard hides
1. Jump to next focus:
    by **`FocusScope`**

>Hide keyboard:  

```java
    focusNode.unfocus();
    // or
    FocusManager.instance.primaryFocus.unfocus();
```

>Jump to next widget:

```
FocusScopeNode scopeNode = FocusScopeNode();

FocusScope focusScope = FocusScope(
    node: scopeNode,
    child: [
        widget(),
        widget(),
    ],
);

// jump to next focus
scopeNode.nextFocus();
```