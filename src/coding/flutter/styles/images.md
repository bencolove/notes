# Image
Sources:
1. *asset*
1. filesystem
1. memory
1. remote(internet)

Loader:   
*`ImageLoader`*  - `ImageLoader.load()`
1. 

## Asset Images
>resource

    root-dir/images/avatar.png  

>config file

-- _`pubspec.ymal`_ --
```yaml
assets:
  - images/avatar.png
``` 

>usage

```java
Image(
    image: AssetImage("images/avatar.png"),
    width: 100.0
)
// Or
Image.asset("iamges/avatar.png",
    width: 100.0
)

// Internet
Image.network(
    "https://www.images.com/path",
    width: 100.0
)
```