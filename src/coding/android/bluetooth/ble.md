# Bluetooth Low Energy BLE Support
[android-ble](https://developer.android.com/guide/topics/connectivity/bluetooth-le)

    requirements:
    ^ Android 4.3 (API Level 18)

## Key concepts:
* GATT Generic Attribute Profile
    * GATT/ATT
* ATT ATT protocol
    * characteristics and services
    * by UUID 128-bits

_service_ := _characteristic_ [ ,_characteristic_ ]  
_characteristic_ := value, [_descriptor_ [,_descriptor_ ]]

---
## Roles and the Play:
* __one__ Central role and __one__ peripheral during scanning and advertising
* GATT __client__ and GATT __server__ during communicating

---
## BLE permissions involved
* BLUETOOTH - request/receive connection and transfer data
* BLUETOOTH_ADMIN - device discover and configure settings
* ACCESS_FINE_LOCATION when **NOT** using _Companion Device Manager_ API

```xml
<uses-permission android:name="android.permission.BLUETOOTH"/>
<uses-permission android:name="android.permission.BLUETOOTH_ADMIN"/>

<!-- Required only if your app isn't using the Device Companion Manager. -->
<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />
```

Support BLE or not
```xml
<uses-feature android:name="android.hardware.bluetooth_le" android:required="true"/>
```
_required_:
* true only support BLE android device
* false otherwise

Check BLE Support
```java
// Use this check to determine whether BLE is supported on the device. Then
// you can selectively disable BLE-related features.
if(!packageManager.hasSystemFeature(PackageManager.FEATURE_BLUETOOTH_LE)) {
    Toast.makeText(this, R.string.ble_not_supported, Toast.LENGTH_SHORT).show()
    finish()
}
```

## Use BLE
1. setup BLE
1. Find BLE devices
1. Connect to GATT server
1. Read BLE attributes
1. Receive GATT notifications

### Setup BLE
To check BLE support and enable it.

>Enable BLE with __`BluetoothAdapter`__

```java
val bluetoothManager = getSystemService(BluetoothManager::class.java)
val bluetoothManagerAdapter: BluetoothAdapter? = bluetoothManager?.adapter
// not null only when a Bluetooth radio is installed
if (bluetoothManagerAdapter != null && !bluetoothAdapter.isEnabled) {
    val enableIntent = Intent(BluetoothAdatper.ACTION_REQUEST_ENABLE)
    startActivityForResult(enableIntent, REQUEST_ENABLE_BT)
}

// ...
onActivityResult(int, int, android.content.Intnet) {
    // requestCode
}
```

