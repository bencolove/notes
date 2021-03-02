# Okhttp
Stacks:
* retrofit + okhttp + rxjava
* retrofit + okhttp + koroutine

## Fuel
[[github][fuel-github]]

Import
```gradle
dependencies {
	//OkHttp3
	implementation "com.squareup.okhttp3:okhttp:4.9.0"

        // define a BOM and its version
       implementation(platform("com.squareup.okhttp3:okhttp-bom:4.9.0"))
       
       // define any required OkHttp artifacts without version
       implementation("com.squareup.okhttp3:okhttp")
       implementation("com.squareup.okhttp3:logging-interceptor")

    //Retrofit
    api "com.squareup.retrofit2:retrofit:2.9.0"
    api "com.squareup.retrofit2:converter-gson:2.9.0"
    implementation "com.jakewharton.retrofit:retrofit2-kotlin-coroutines-adapter:0.9.2"
    //Kotlin Coroutines
    api "org.jetbrains.kotlinx:kotlinx-coroutines-android:1.4.2"
}
```

Patterns
1. DSL
1. Callback
1. LiveData

--_`api.kt(retrofit)`_--
```kotlin
interface TestService {
    @GET("/banner/json")
    suspend fun getBanner(): WanResponse<List<Banner>>
}
```


Workflows with API:
1. common setting (timeout, host, ssl, ...)
1. create request (header, body, method, ...)
1. POJO -> Json (request)
1. Json -> POJO (response)
1. send request

OKhttp: `1` and `5`  
Gson(Jackson): `3` and `4`  
Retrofit: `2` integrates all

>1. Common Settings

```java
val baseUrl = "https://myhost.com.tw/"
val retrofit = Retrofit.Builder()
        .baseUrl(baseUrl)
        .addConverterFactory(GsonConverterFactory.create(gson)) //3 and 4
        .client(createHttpClient()) // low-level httpclient
        .build()
```

>2. Create Request

```java
interface LoginService {
    @POST("/login")
    suspend fun login(@Body body: LoginParams): LoginResult
}
```

>3. Send Request

```java
val service = retrofit.create(LoginService::class.java)
// MainScope
MainScope().launch {
    showLoading()
    val result = withContext(Dispatchers.IO) {
        // withContext is suspend
        service.login(request)
    }
    closeLoading()

    if (result.isSuccess) { ... }
}
```


[okhttp-github]: https://github.com/square/okhttp
[retrofit-github]: https://github.com/square/retrofit
[kotlin-coroutine]: https://github.com/Kotlin/kotlinx.coroutines