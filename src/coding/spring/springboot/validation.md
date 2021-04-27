# Validation
Implemenations:
1. hibernate validation API

Aspects:
1. Annotations
1. Groups

Spring Integration:
1. where
1. when
1. manually
1. [[custom validator][custom-validator]]
1. message interpolation
1. handle vailidation exception

## Manually Validate
```java
// get the Validator
Validator validator = Validation.buildDefaultValidatorFactory().getValidator();
// or @Autowired
@Autowired Validator validator;

// validate whole bean
Set<ConstraintViolation<TargetType>> violations = validator.validate(targetBean);

// validate single property instead of all
String message = validator.validateProperty(tagetBean, "propertyName")
    .iterator()
    .next()
    .getMessage();

```

## Message Interpolation
`message` field of all annotations could be interpolated with annotation-specific values like:
1. values referred from annotation instance: `{value} {max} {min}` from `@Size`
1. `${validatedValue}` refered to current value being validated
1. `${externalValue}` refered to values outside of annotation
1. `${formatter.format('%1$.2f', validatedValue)}`
1. [custom inject message context](./validate_message_interpolation.md)

---

## Handle Validation Result/Exception
Validation related SpringBoot exceptions:

Exception | Thrown From
---|---
`BindException` | @Validated on class
`MethodArgumentNotValidException` | @Valid on method parameter
`ConstraintViolationException` | validator annotations on @Valid class

Handle:
```java
/* 
BindException.class,
MethodArgumentNotValidException.class
directly relate to BindingResult
*/
BindingResult bindingResult = null;
if (raw instanceof BindException) {
    bindingResult = ((BindException)(raw)).getBindingResult();
} else if (raw instanceof MethodArgumentNotValidException) {
    bindingResult = ((MethodArgumentNotValidException)raw).getBindingResult();
}
// handle BindingResult then
List<String> messages = new ArrayList<>();
if (results != null) {
    // errors on the object (root)
    if (results.hasGlobalErrors()) {
        results.getGlobalErrors().stream().forEach(error -> {
            messages.add(error.getDefaultMessage());
        });
    }

    // errors on fields (path)
    if (results.hasFieldErrors()) {
        results.getFieldErrors().stream().forEach(error -> {
            messages.add(error.getDefaultMessage());
        });
    }
}

/*
Handle ConstraintViolationException.class
*/
Iterator<ConstraintViolation<?>> ite = raw.getConstraintViolations().iterator();
while (ite.hasNext()) {
    ConstraintViolation<?> error = ite.next();
    // only need the last part of property.Path
    String lastPathName = StreamSupport.stream(error.getPropertyPath().spliterator(), false)
            .reduce( (f, s) -> s )
            .map(node -> node.getName())
            .orElseGet(() -> error.getPropertyPath().toString());
    String errorMessage = String.format("%s: %s", lastPathName, error.getMessage());
    messages.add(errorMessage);
}
```


[validation-group]: https://nullbeans.com/how-to-use-java-bean-validation-in-spring-boot/
[custom-validator]: https://www.baeldung.com/spring-mvc-custom-validator
[message-interpolation]: https://www.baeldung.com/spring-validation-message-interpolation