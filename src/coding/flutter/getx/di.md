## Dependency Injection
1. `Get.pu` and `Get.lazyPut` in constructor or build in top widget 
1. `init` parameter of _`GetBuilder`_ in top widget
1. bindings

## Bindings

> Separated class  
```java
class HomeBinding implements Bindings {
    @override
    void dependencies() {
        Get.lazyPut<HomeController>(() => HomeController());
        Get.put<Service>(() => Api());
    }
}
```

### 1. Global
```java
    GetMaterialApp(
        initialBinding: GlobalBindings(),
        ...
    )
```
### 2. Routes
```java
// static
getPages: [
    GetPage(
        name: '/',
        page: () => HomeView(),
        // per-route binding
        binding: HomeBiding(),
    )
]

// static with BindingBuilder
GetPage(
    name: '/',
    page: () => HomeView(),
    binding: BindingsBuilder(() {
      Get.lazyPut<ControllerX>(() => ControllerX());
      Get.put<Service>(()=> Api());
    }),
  ),

// dynamic
Get.to(Home(), binding: HomeBinding());
```

### Per widget
Reuse same controller class for multiple widgets.
```java
// put and find with tag(like id)
Get.put(Controller(), tag: "inst_1");

Get.find<Controller>(tag: "inst_1");

// getbuilder
GetBuilder(
    init: Controller(),
    global: false
    ...
    )
```

