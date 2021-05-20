# Flutter's RWD Way
Methods to adopt:
1. `MediaQueryData` from `InheritedWidget` provides a `Size` property of current screen
1. `LayoutBuilder` provides a `BoxContraints` with `min/max` `width/height`

## MediaQuery
With `MediaQuery.of(context).size` you build the Widget based on it's `width` and `height`

## LayoutBuilder


[show-how]: https://www.smashingmagazine.com/2020/04/responsive-web-desktop-development-flutter/