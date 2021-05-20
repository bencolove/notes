# Basic MethodChannel
![](https://flutter.dev/images/PlatformChannels.png)

Roles:
    * Client - Flutter app
    * Host - native (Andoird/iOS)


## Client calling site
```java
class _State extends State<Widget> {
    static const platform = const MethodChannel('unique_name');

    Future<void> _getBatteryLevel() async {
        String level;
        try {
            final int result = await platform.invokeMethod('getBatteryLevel');
            batteryLevel = 'Battery $result %';
        } on PlatformException catch (e) {
            level = "Failed by: '${e.message}'";
        }

        setState((){
            _batteryLevel = level;
        })
    }

    @override
    Widget build(BuildContext context) {
        return Material(
            child: Center(
                child: Column(
                    children: [
                        ElevatedButton(
                            child: Text('Get battery level'),
                            onPressed: _getBatteryLevel,
                        ),
                        Text(_batteryLevel),
                    ]
                )
            )
        );
    }
}

```

## Host native site Android
    _-- root/android/app/src/main/kotlin/{pkg}/MainActivity.kt--_
```java
  private fun getBatteryLevel(): Int {
    val batteryLevel: Int
    if (VERSION.SDK_INT >= VERSION_CODES.LOLLIPOP) {
      val batteryManager = getSystemService(Context.BATTERY_SERVICE) as BatteryManager
      batteryLevel = batteryManager.getIntProperty(BatteryManager.BATTERY_PROPERTY_CAPACITY)
    } else {
      val intent = ContextWrapper(applicationContext).registerReceiver(null, IntentFilter(Intent.ACTION_BATTERY_CHANGED))
      batteryLevel = intent!!.getIntExtra(BatteryManager.EXTRA_LEVEL, -1) * 100 / intent.getIntExtra(BatteryManager.EXTRA_SCALE, -1)
    }

    return batteryLevel
  }

  override fun configureFlutterEngine(@NonNull flutterEngine: FlutterEngine) {
        super.configureFlutterEngine(flutterEngine)
        MethodChannel(flutterEngine.dartExecutor.binaryMessenger, CHANNEL).setMethodCallHandler {
                call, result ->
  // Note: this method is invoked on the main thread
    if (call.method == "getBatteryLevel") {
        val level = getBatteryLevel()
        if (level != -1) {
            result.success(level)
        } else {
            result.error("unavailable", "blarblar", null)
        }
    } else {
        result.notImplemented()
    }

```