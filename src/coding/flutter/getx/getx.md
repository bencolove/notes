# GetX

1. Install
1. StateMangement
1. RouteManagement
1. Dependency Injection
1. Internationalization

## Install
--_`pubspec.yaml`_--
```yaml
depdencies:
  get:
```

`import 'package:get/get.dart'`;

## StateManagement
Bindings API
1. GetBuilder
1. GetX
1. Obx
1. MixinBuilder

> Overall  

Method | RAM Consumption | Reactive
---|---|---
`GetBuilder` | Lowest | NO `update()` needed
`Obx` | Medium | YES
`GetX` | High | YES

---

### Reactive Types
>NOTE: These are essentially Stream types under the hood:

Wrapping Type | _.obs_ | Rx Form | Rx Generics 
---|---|---|---
String | ''.obs | RxString('') | Rx<String>('')
boolean | false.obs | RxBool(false) | Rx<Bool>(fase)
int | 0.obs | RxInt(0) | Rx<Int>(0)
double | 0.0.obs | RxDouble(0.0) | Rx<Double>(0.0)
number | 0.obs | x | Rx<Num>(0)
list | <String>[].obs | RxList<String>([]) | Rx<List<String>>([])
map | <String, int>{}.obs | RxMap<String, int>({}) | Rx<Map<String, int>>({})
custom class | x | Rx<User>() 

### Controller
Single source of truth, and the source of changes.

```java
class Controller extends GetxController {
    /* maintain state data here */

    // util method to access this controller
    // need to Get.put(Controller()) first ???
    static Controller get instance => Get.find();

    /* do sth like subscribe to Streams */
    @override
    void onInit();

    /* do sth like unscribe from Streams*/
    @override
    void onClose();
}
``` 

Ways to get hold of **`THE`** instance of Controller:
1. `Get.find<Controller>()`
1. `var controller = Get.find()`
1. `Controller.instance`

### GetBuilder
Mechanics of State Management instead of __`ChangeNotifier`__ (needing `udpate()` to notify listeners). 

```java
GetBuilder<Controller>(
    // optional
    initState: // analogue to Widget's initState
    // optional
    dispose: // analogue to Widget's dispose

    // Unique id for Get.update(['id']) to notify
    id: 'id',
    // initial value, only the first one will be created
    init: Controller(),
    builder: void Function(Controller)
)

// update GetBulder programmatically:
update(['id']);
// conditionally
update(['id'], propPredicate);
```

> Observable:  
```java
class Controller extends GetxController{
    var count = 0.obs;
    increment() => count++;
}
```

> Observer:  
```java
// DI for all child routes ??? not sub Widgets
final Controller c = Get.put(Controller());

Obx(() => Text('${c.count}'))
```

> Observer:  
```java
class Other extends StatelessWidget {
    final Controller c = Get.find();

    @override
    Widget build(context) {
        return Scaffold(body: Center(
            child: Text('${c.count}')
        ));
    }
}
```

## Routes Management

Comparisons:

Intention | GetX | Navigator Version
---|---|---
jumpTo | Get.to(Widget) | ```Navigator.push(context, MaterialPageRoute<void>(builder: (context) => Widget))```
back | Get.back() | ```Navigator.pop(context)```
replaceCurrent | Get.off(Widget) | `Navigator.pushReplacement(context, MaterialPageRoute)`
jumpTo and clear all | Get.offAll(Widget) | ```Navigator.pushAndRemoveUntil(context, MaterialPageRoute, (Route<dynamic> route)=>false)```
await for jumpTo | await Get.to(Widget) <br> Get.back(result: 'success') | `await Navigator.push(context, MaterialPageRoute)` <br> `Navigator.pop(context, 'success')`
jumpToNamed | Get.toNamed(path) | `Navigator.pushNamed()`
replace current | Get.offNamed(path) | `Navigator.pushNamedReplacement`
clear routes to one | Get.offAllNamed(path) |
arguments | Get.toNamed(path, arguments: dynamic) <br> Get.arguments | `Navigator.pushNamed(path, arguments: dynamic)` <br> `ModaRoute.of(context).settings.arguments`

>Pass arguments like WEB:
```java
// passing predefined(place holder)
GetPage(
    name: '/profile/:user',
    page: () => UserProfile(),
)

// passing
Get.toNamed('/profile/1838');
Get.offAllNamed('/nextpage?device=iphone&id=xxy');

// retrieving
String userId = Get.parameters['user'];
String device = Get.parameters['device'];
int id = Get.parameters['id'];
```



### Declare Routes
```java
abstract class Routes {
    static const Initial = '/';
    static const NextScreen = '/nextscreen';
}

abstract class AppPages {
    static final pages = [
        GetPage(
            name: Routes.Initial,
            page: () => HomePage(),
        ),
        GetPage(
            name: Routes.NextScreen,
            page: () => NextScreen(),
            transition: Transition.zoom,
            // applied middlewares
            middlewares: [
                GetMiddleware()
            ]
        )
    ];
}

void main() {
    runApp(
        GetMaterialAPP(
            debugShowCheckedModeBanner: false,
            theme: appThemeData,
            defaultTransition: Transition.fade,

            // predefined routes
            initialRoute: '/',
            getPages: AppPages.pages,

            home: HomePage(),
            // unmatched route
            unknownRoute: GetPage(name: '/notfound', page: () => UnknownRoutePage()),
        )
    );
}

```

### Middleware around Route Change
Route guard:
```java
GetMaterialApp(
    routingCallback: (routing) {
        if (routing.current == '/login') {

        }
    }
)
```

--_`lib/common/middleware/route_auth.dart`_--
```java
class RouteAuthMiddleware extends GetMiddleware {
    @override
    int priority = 0;

    RouteAuthMiddleware({required this.priority});

    @override
    RouteSettings? redirect(String route) {
        Future.delayed(Duration(seconds: 1), () => Get.snackbar('提示','請登錄'));
        return RouteSettings(name: AppRoutes.Login);
    }
}

// whitelistt
GetPage(
    name: AppRoutes.Login,
    page: () => LoginView(),
),
GetPage(
    name: AppRoutes.NeedLogin,
    page: () => UserProfile(),
    middlewares: [
        RouteAuthMiddleware(priority: 1),
    ]
)
```

### SnackBars
`Get.snackbar('Hi', 'i am a modern snackbar')`  vs  
```java
final snackbar = SnackBar(
    content: Text('Hi'),
    action: SnackAction(
        label: 'I am an old snackbar',
        onPressed(){}
    )
);

Scaffold.of(context).showSnackBar(snackbar);
```

### Dialogs
```java
Get.defulatDialog(
    onConfirm: () => print('OK'),
    middleText: 'Dialog made of 3 lines of code'
);

Get.dialog(Widget);
```

### BottomSheets
```java
Get.bottomSheet(
    Container(
        child: Wrap(
            children: [
                ListTile(
                    leading: Icon(Icons.music_note),
                    title: Text('music'),
                    onTap: (){}
                ),
                ListTile(
                    leading: Icon(Icons.videocam),
                    title: Text('video'),
                    onTap: (){}
                )
            ]
        )
    )
)
```

### Transition
```java
GetPage(
    name: AppRoutes.Detail_ID,
    page: () => DetailView(),
    transition: Transition.downToUp,
)
```

## i18n
```java
import 'package:get/get.dart';

class I18nMessages extends Translations {
    @override
    Map<String, Map<String, String>> get keys => {
        'en_US': {
            'hello': 'Hello world',
        },
        'zh_TW': {
            'hello': '哈囖',
        }
    };
}


// usage reference
Text('hello'.tr);

```