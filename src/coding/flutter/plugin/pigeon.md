# Plugin by Pigeon
[flutter](https://flutter.dev/docs/development/platform-integration/platform-channels)

[pigeon](https://pub.dev/packages/pigeon)

[example](https://www.gushiciku.cn/pl/ptM3/zh-tw)

3 kinds of files needed:
1. pigeon definitions
    generated:
    * dart side code
    * java code
1. dart calling code
1. java handling code

## Generat from dart definition
>pigeon definitions:

-- root/tools/pigeon/definitions.dart
```java
class BatteryLevelRequest {
  int? default;
}

class BatteryLevelReply {
  int? data;
}

@HostApi()
abstract class BatteryService {
  @async
  BatteryLevelReply getBatteryLevel(BatteryLevelRequest);
}
```
```sh
$ flutter pub run pigeon \
--input tools/pigeon/definitions.dart
--dart_out tools/pigeon/platform_call.dart
--java_out tools/pigeon/PlatformService.java
```

## Calling from dart
```java
BatterLevelResult result = await BatteryService().getBatteryLevel();
```

## Respond on Android
_--MainActivity.kt--_
```java
inner class BatteryServiceCallbackHandler: PlatformService.BatteryService {
        override fun getBatteryLevel(result: PlatformService.Result<PlatformService.BatteryLevelResult>?) {
            val level = getBatteryLevel()
            val levelResult = PlatformService.BatteryLevelResult().apply {
                data = level.toLong()
            }.also {
                Log.i("Flutter", "getBatteryLevel callback, result=${it.data}")
            }
            result?.success(levelResult)
        }
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        // call pigeon framework
        PlatformService.BatteryService.setup(
            flutterEngine?.dartExecutor?.binaryMessenger, BatteryServiceCallbackHandler())
    }
```