# _fileio_

```
import `dart:async`;
import `dart:io`;
import `package:path_provider/path_provider.dart`;
```

## 1. Location
Platform agnostic 
1. temporary directory
    * iOS - `NSCachesDirectory`
    * android - `getCacheDir()`
1. document directory
    * iOS - `NSDocumentDirectory`
    * android - `AppData`

```dart
Future<String> get _localPath async {
    final directory = await getApplicationDocumentDirectory();
    return directory.path;
}
```

## 2. Ref
`File` class from `dart:io`
```
Future<File> get _localFile async {
    final path = await _localPath;
    return File(`$path/counter.txt`)
}
``` 

## 3. Write
```
Future<File> writeInt(int counter) async {
    final file = await _localFile;
    return file.writeAsString('1'.toString());
}
```

## 4. Read
```
Future<int> readInt() async {
    try {
        final file = await _localFile;

        String contents = file.readAsString();

        return int.parse(contents);
    } catch (e) {
        return 0;
    }
}
```