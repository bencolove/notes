# JSON Serialization and Deserialization
* [[Gson][gson-guide]] `implementation 'com.google.code.gson:gson:2.8.6'`
* Jackson

Considerations | Gson | Jackson
---|---|---
field rename | @SerializedName | @JsonProperty @JsonSetter/Getter
output null |
input null |
field transform |
field custom |
Views | @Since @Until | @JsonView

## Gson
[[tutorial][gson-tut]]  
[[Gson default constructor](https://blog.csdn.net/lmj623565791/article/details/106631386)]

>NOTE  
Gson creates an object from String using default constructor( no-arg version ). If it is not provided, Gson use an *unsafe* way to directly alloc free space for the object which may bypass **NON-NULL** declaration of Kotlin upholds leading to `null` value causing runtime exception. In order to make a `data class` with *no-arg* constructor, assign all proproties with default values.

>Primitives

```java
Gson gson = new Gson();

int i = gson.fromJson("100", int.class);              //100
double d = gson.fromJson("\"99.99\"", double.class);  //99.99
boolean b = gson.fromJson("true", boolean.class);     // true
String str = gson.fromJson("String", String.class);   // String

String jsonNumber = gson.toJson(100);       // 100
String jsonBoolean = gson.toJson(false);    // false
String jsonString = gson.toJson("String"); //"String"
```

>POJO

```java
public class Person {
    @SerializedName("last_name")
    public String lastName;
}

```

>Generic types

```java
public class Result<T> {
    public int code;
    public String message;
    public T data;
}

Type userListType = new TypeToken<Result<List<User>>>(){}.getType();
Result<List<User>> userListResult = gson.fromJson(json,userListType);
List<User> users = userListResult.data;

// shortcut with inline function
inline fun <reified T> Gson.fromJson(json: String): T = fromJson(json, object: TypeToken<T>() {}.type)

val result: Result<InnerType> = Gson().fromJson<Result<InnerType>>(jsonString)
```

> Streaming

```java
String json = "{\"name\":\"怪盗kidou\",\"age\":\"24\"}";
User user = new User();
JsonReader reader = new JsonReader(new StringReader(json));
reader.beginObject(); // throws IOException
while (reader.hasNext()) {
    String s = reader.nextName();
    switch (s) {
        case "name":
            user.name = reader.nextString();
            break;
        case "age":
            user.age = reader.nextInt(); //自动转换
            break;
        case "email":
            user.email = reader.nextString();
            break;
    }
}
reader.endObject(); // throws IOException
System.out.println(user.name);  // 怪盗kidou
System.out.println(user.age);   // 24
System.out.println(user.email); // ikidou@example.com

// programatically
JsonWriter writer = new JsonWriter(new OutputStreamWriter(System.out));
writer.beginObject() // throws IOException
        .name("name").value("怪盗kidou")
        .name("age").value(24)
        .name("email").nullValue() //演示null
        .endObject(); // throws IOException
writer.flush(); // throws IOException
//{"name":"怪盗kidou","age":24,"email":null}
```

> Custom builder
```java
Gson gson = new GsonBuilder()
    .serialzeNulls() // output null field
    .setDateFormat("yyyy-MM-dd")
    .disableInnerClassSerialization() // no inner class deser
    .generateNonExecuteableJson() // append )]}' four chars
    .disableHtmlEscaping() // no &lt;
    .setPrettyPrinting()
    // only ser/deser @Expose fields
    .excludeFieldsWithoutExposeAnnotation()
    .create();
```

## `TypeToken`
Generic parameterized type for `Container` classes are eventually erased in runtime. In order to capture the contained class, `TypeToken` is used as:
```java
TypeToken<?> listType = new TypeToken<List<String>>(){}.getType();

TypeToken<?> containedTypeToken = listType.resolveType(ArrayList.class.getTypeParameters()[0]);
// java.lang.String
containedTypeToken.getType()
```

[gson-guide]: https://github.com/google/gson/blob/master/UserGuide.md
[gson-tut]: https://www.jianshu.com/p/3108f1e44155