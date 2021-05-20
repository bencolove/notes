# Text Input
1. TextField
1. Form

## TextField
Props:
1. `TextEditingController` controller
1. `TextInputType` keyboardType  
    * text
    * multiline |
    * number |
    * phone |
    * datetime |
    * emailAddress |
    * url  
1. `onChange`, `onEditingComplete`, `onSubmitted`:
    * onSubmitted - `ValueChanged<String>`
    * onEditingComplete - `Function()`

```java
TextField(
    decoration: InputDecoration(
        labelText: "密码",
        hintText: "您的登录密码",
        prefixIcon: Icon(Icons.lock)
    ),
    obscureText: true,
    ),
```

>Interactive
1. defined String
    * `build()` based it
    * `onChange` => update
1. `TextEditingController`
    * TextField.controller = controller
    * controller.text

>Focus


## Form
Set of TextField (FormField)

`FormState` (via `Form.of(BuildContext)` )
1. validate() -> Form's children's (FormField) validate()
1. save() -> FormField's save()
1. reset() -> FormField's reset()





