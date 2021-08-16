# Built-in Route Management

Components:
1. __`Route`__
1. __`Navigator`__

Advanced:  
1. Basic Navigator(routes stack) methods
1. Pass values through reoutes 
1. Observe routes changes
1. Intercept before pop

---

>1. Route  
A page (Widget), `MaterialPageRoute(builder: (context) {
            return BPage();
          })`

>2. Navigator  
Navigator is a stack-like Widget manages the routes. Widget that exposes inner state via _`of`_ `Navigator.of(context)` 
---
## Navigator Methods:
1. push
1. pushNamed(path)  
    `MaterialApp.routes<String, WidgetBuilder>{ path: (context){} }`
1. pushReplacementNamed(path)  
    `routes.pop() && routes.pushNamed(path)`
1. popAndPushNamed(path)  
1. pop
1. maybePop()  
    `if (routes.isNotEmpty) routes.pop();`
1. canPop()  
    `routes.isNoEmpty`

1. pushNamedAndRemoveUntil(path, ModalRoute.withName(path))
1. pushNamedAndRemoveUntil(path, predicate: boolean (Route))
1. popUntil(predicateOrNamedRoute)

---
## Pass values to a Route:
1. Widget constructed from WidgetBuilder
```java
Navigator.of(context).push(MaterialPageRoute(builder: (context){
    return YourWidget(passedArguments);
}));
```
2. Named route
```java
Navigator.of(context).pushNamed('/pageB', arguments: 'fromPageA');

ModaRoute.of(context).settings.arguments
```

3. await pop
```java
onPressed: () async {
    var result = await Navigator.of(context).pushNamed('/pageB', arguments: 'fromPageA');
}

...
// page B
Navigator.of(context).pop('backFromPageB');
```
---

## Observe Routes Change  
`RouteOberserver`
```java
RouteObserver<PageRoute> routeObserver = RouteObserver<PageRoute>();

// observable: MaterialApp.routes
// MaterialApp.routes -> RouteObserver<Route> -> it's observers
MaterialApp.navigatorObservers: [routeObserver];

// in a Widget
@override
void didChangeDependencies() {
    super.didChangeDependencies();
    routeObserver.subscribe(this, ModalRoute.of(context));
}

@override
void dispose() {
    super.dispose();
    routeObserver.unsubscribe(this);
}
/* RouteObserver interface */
@override
  void didPush() {
    final route = ModalRoute.of(context).settings.name;
    print('A-didPush route: $route');
  }

  @override
  void didPopNext() {
    final route = ModalRoute.of(context).settings.name;
    print('A-didPopNext route: $route');
  }

  @override
  void didPushNext() {
    final route = ModalRoute.of(context).settings.name;
    print('A-didPushNext route: $route');
  }

  @override
  void didPop() {
    final route = ModalRoute.of(context).settings.name;
    print('A-didPop route: $route');
  }
```

Route Change Logger
```java
class MyRouteObserver<R extends Route<dynamic>> extends RouteObserver<R> {
  @override
  void didPush(Route route, Route previousRoute) {
    super.didPush(route, previousRoute);
    print('didPush route: $route,previousRoute:$previousRoute');
  }

  @override
  void didPop(Route route, Route previousRoute) {
    super.didPop(route, previousRoute);
    print('didPop route: $route,previousRoute:$previousRoute');
  }

  @override
  void didReplace({Route newRoute, Route oldRoute}) {
    super.didReplace(newRoute: newRoute, oldRoute: oldRoute);
    print('didReplace newRoute: $newRoute,oldRoute:$oldRoute');
  }

  @override
  void didRemove(Route route, Route previousRoute) {
    super.didRemove(route, previousRoute);
    print('didRemove route: $route,previousRoute:$previousRoute');
  }

  @override
  void didStartUserGesture(Route route, Route previousRoute) {
    super.didStartUserGesture(route, previousRoute);
    print('didStartUserGesture route: $route,previousRoute:$previousRoute');
  }

  @override
  void didStopUserGesture() {
    super.didStopUserGesture();
    print('didStopUserGesture');
  }
}
```

---

## Intercept before Pop
When:
1. Ask end-user before pop(exit)
1. Extra business logics

Tool: `WillPopScope.onWillPop( Future<boolean> Function() async)`
```java
WillPopScope(
    onWillPop: () async => showDialog(
        context: context,
        builder: (context) =>
            AlertDialog(title: Text('你确定要退出吗？'), actions: <Widget>[
              RaisedButton(
                  child: Text('退出'),
                  onPressed: () => Navigator.of(context).pop(true)),
              RaisedButton(
                  child: Text('取消'),
                  onPressed: () => Navigator.of(context).pop(false)),
            ])),
    child: Container(
      alignment: Alignment.center,
      child: Text('点击后退按钮，询问是否退出。'),
    ))
```

Double-click _back_ button to quit:
```java
DateTime _lastQuitTime;
WillPopScope(
    onWillPop: () async {
      if (_lastQuitTime == null ||
          DateTime.now().difference(_lastQuitTime).inSeconds > 1) {
        print('再按一次 Back 按钮退出');
        Scaffold.of(context)
            .showSnackBar(SnackBar(content: Text('再按一次 Back 按钮退出')));
        _lastQuitTime = DateTime.now();
        return false;
      } else {
        print('退出');
        Navigator.of(context).pop(true);
        return true;
      }
    },
    child: Container(
      alignment: Alignment.center,
      child: Text('点击后退按钮，询问是否退出。'),
    ))
    ```