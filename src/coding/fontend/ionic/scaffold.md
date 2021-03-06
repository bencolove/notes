# Scaffold with `Ionic`
[How-to build capacitor plugin][custom-capacitor]
[[Make sure typescript comfiled against es5 not es2015][ts-config-es5]]
[access localhost from Android emulator [`10.0.2.2`]]

1. scaffold
1. workflow
    * modify typescript if using angular `tsconfig.json[14]: "target": "es5"`

<br>

---

## 1. Scaffold Plugin Project
```
npm -g @capacitor/cli
npx @capacitor/cli plugin:generate
# folder created
cd pluginfolder
npm run build
```

## 2. Add Plugin Implementions (Web/Android/Ios)
> 2.1 Typescript Definitions

--_`plugin/src/definitions.ts`_--
```typescript
declare module "@capacitor/core" {
  interface PluginRegistry {
      // ContactsPlugin is the name that should be imported
    ContactsPlugin: ContactsPluginPlugin;
  }
}
 
export interface MyPluginInterface {
  echo(options: { value: string }): Promise<{value: string}>;
  
  // --- ADD your custom callback function interfaces here --- 
  getContacts(filter: string): Promise<{results: any[]}>;
}
```

>2.2 Web Implementions  

--_`plugin/src/web.ts`_--
```typescript
export class MyPlugin extends WebPlugin implements MyPluginInterface {
    // ...
    // --- implementation for web callback ---
    // other native impls like android/ios will be located elsewhere
    async getContacts(filter: string): Promise<{results: any[]}> {
        console.log('Your web implementation here')
        return {
            results: [{
                // ...
            }]
        }
    }
}
```

>2.3 Ios Implementations

--_`ios/Plugin/Plugin.m`_--

>2.4 Android Implementations

`Android Studio` open the plugin folder `android` to locate your class to implement `MyPlugin.java`
```java
@NativePlugin(requestCodes)
class MyPlugin extends Plugin {

// --- your custom native code goes here ---
    @PluginMethod()
    public void getContacts(PluginCall call) {

    }
    /* How to handle permission authentication with dialog
    1. saveCall() before authen dialog
    2. call pluginRequestPermission() to trigger authen process by showing dialog
    3. deal in callback handleRequestPermissionResult
        3.1 loop to check whether perm is authenticated
        3.2 proceed when permitted
    */
}

```

---

## Use MyPlugin in Ionic App
>1. Prepare Ionic App if not created
```sh
npm i -g @ionic/cli
ionic start myapp blank --type=angular --capacitor 
cd myapp
# use local relative path    
npm i path/to/plugin

# build app
ionic build
# add native platforms
npx cap add android
```
Once done, `"myplugin": "file:path/to/plugin",` should be in depedencies of `package.json`

>2. Add permissions   

--_`android/app/src/main/AndroidManifest.xml`_--

    <uses-permission android:name="android.permission.READ_CONTACTS" />

>3. Register plugin

--_`android/app/src/main/io/ionic/starter/MainActivity.java`_--
```java
import MyPlugin;

public class MainActivity extends BridgeActivity {
    onCreate
        this.init(bundle, new ArrayList<>{}{{
            // add your plugin here to let inoic be aware of 
            add(MyPlugin.class);
        }})
}
```

>4. Used in ionic Pages

--_`home/home.page.ts`_--
```typescript
import { Plugins } from '@capacitor/core';
// ADD
import 'myplugin';
const { MyPlugin } = Plugins;

...
await MyPlugin.getContacts('filter').results

```

>5. Build ionic app

```sh
// plugin code changed
plugin_dir> npm run build

// ionic app must be rebuilt after changes
ionic_dir> ionic build
// push updates to android studio
ionic_dir> npx cap sync
```

<br>

---

## 1.Create Plugin
```shell
$ npm i -g @capacitor/cli
$ npx @capacitor/cli plugin:generate
 ... Q&A
# folder created
$ cd pluginfolder
$ npm run build
```
Folder layout:
* android - native side
* dist - build output of web
* ios - native side
* src - web implementation

## 2.Create Ionic App if not Existed
```sh
$ npm i -g @ionic/cli
$ ionic start myapp blank --type=angular --capacitor 
$ ionic version
```

## 3.Add Plugin into Ionic App
```shell
$ cd myapp
# use local relative path    
$ npm i path/to/plugin
```

## 4.Add native platform
```shell
# build app
$ ionic build
# add native platforms
$ npx cap add android
# or
$ ionic capacitor add android
```

## 5.Workflow of capacitor
1. Open ionic app folder  
`npx cap open android`
1. Make changes to plugin code
1. Fresh changes into ionic app  
`npx cap sync`
1. See udpates from Android Studio

<br>

## Additional Dependencies in Plugin
By looking at `plugin/android/build.gradle`
```gradle
dependencies {
    // every local jar in plugin/android/libs will be included
    implementation fileTree(dir: 'libs', include: ['*.jar'])

    // additional dependencies
    // somehow when plugin needs Activiy, it is required
	implementation "androidx.appcompat:appcompat:1.1.0"
	
	// kotlin koroutine
	implementation "org.jetbrains.kotlinx:kotlinx-coroutines-core:1.4.2"
    implementation "org.jetbrains.kotlinx:kotlinx-coroutines-android:1.4.2"

    // add maven jar slf4j android
    implementation group: 'org.slf4j', name: 'slf4j-api', version: '1.7.30'
    implementation group: 'org.slf4j', name: 'slf4j-android', version: '1.7.30'
}
```
Note: remember to `npm run build` your plugin before proceed

[custom-capacitor]: https://devdactic.com/build-capacitor-plugin/
[ts-config-es5]: https://angular.io/guide/typescript-configuration