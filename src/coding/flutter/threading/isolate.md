# Isolate
[[example1][example1]]

Toolkit:
1. **`Isolate`**
1. _`compute`_
1. _`loadBalancer`_

## Isolate
```java

// top-level funciton
void topLevelHeavyTask(SendPort sendPort) {

}

void scheduleHeavyTask() async {
    ReceivePort receivePort = ReceivePort();
    Isolate isolate = await Isolate.spawn(topLevelHeavyTask, receivePort.sendPort);
    receivePort.listen((message) {
        print("message from worker: $message");
    })
    // suspend by waiting for user keystoke
    await stdin.first;
    // stop opened isolate
    isolate.kill(priority: Isolate.immediate);
}
```



[example1]: https://danielkao.medium.com/%E5%B9%BB%E6%BB%85-%E6%98%AF%E6%88%90%E9%95%B7%E7%9A%84%E9%96%8B%E5%A7%8B-flutter-%E7%9A%84-async-%E8%88%87-isolate-2f87321a7ba8
