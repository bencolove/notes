# WrokManager
![](../res/arch/overview-criteria.png)

Schedules background tasks not belonging to any views.

Features:
1. scheduling to right place
1. `WorkConstraint` to declare conditions
1. *one-time* or *repeatedly*
1. retry policy
1. chain
```java
WorkManager.getInstance(...)
    .beginWith(listOf(workA,workB))
    .then(workC)
    .enqueue()
```
6. support `rxJava` and `coroutine`

<br>

---
## How `Worker` Works
A `Worker` defines the task like `Runnable`. 

`WorkerManage` schedules automatically your `Worker.work()` on a background thread from `Executor` configured in `WorkerManager`'s `Configuration`.

Manually custom Executor for `WorkManager`:
```java
WorkManager.initialize(
    context,
    Configuration.Builder()
        .setExecutor(Executors.newFixedThreadPool(8))
        .build())
)
```


## Coroutine Worker
`implementation 'com.android:work-runtime-ktx`

> koro by default schedule to Dispatchers.Default

Instead of extending `Worker`, use `CoroutineWorker`:
```java
class CoroutineDownloadWorker(
    context: Context,
    params: WorkerParameters
) : CoroutineWorker(context, params) {

    override suspend fun doWork(): Result = {
        val data = downloadSynchronously("https://www.google.com")
        saveData(data)

        Result.success()
    }
}
```