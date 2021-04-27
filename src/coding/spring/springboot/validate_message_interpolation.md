# Custom Interpolate Validation Message

```java
// usage validate over two fields
// needed to be a TYPE annotation
@DatetimeRange(type = ValidateDatetimeType.DATE, 
begin = "beginDate", end = "endDate", 
message = "時段[#{get('begin')} - #{get('end')}] 錯誤")
class NeedToValidateItsBeginAndEnd { ... }

// declare
@Target({ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = DatetimeRangeValidator.class)
public @interface DatetimeRange {
    String message() default "invalid";
    Class<?>[] groups() default {};
    Class<? extends Payload>[] payload() default {};
    ValidateDatetimeType type() default ValidateDatetimeType.DATE;
    String begin() default "";
    String end() default "";
    // multiple instances
    @Target({ElementType.TYPE})
    @Retention(RetentionPolicy.RUNTIME)
    @interface List {
        DatetimeRange[] value();
    }
}

/*
Needed to interpolate "時段[#{get('begin')} - #{get('end')}] 錯誤"
from bean's properties
#{} - El expression
get('begin') - injected context as a Map, its Map.get(String)
*/

public class CustomValidator implements ConstraintValidator<TargetAnnoClazz, TagetValue> {
    @Override
    public boolean isValid(TargetValue o, ConstraintValidatorContext constraintValidatorContext) {
        // use Spring El expression to interpolate message
        ExpressionParser parser =  new SpelExpressionParser();
        // recognize #{} as El expression
        TemplateParserContext elDelim = new TemplateParserContext();
        // inject context as a Map
        StandardEvaluationContext context =new StandardEvaluationContext(ImmutableMap.of("begin", beginValue, "end", endValue));
        // or inject context as the bean itself
        // and get value by #{begin} and #{end}
        StandardEvaluationContext context =new StandardEvaluationContext(o);

        // "時段[#{get('begin')} - #{get('end')}] 錯誤" to interpolate
        String template = constraintValidatorContext.getDefaultConstraintMessageTemplate();
        // parse and interpolate from context
        String errorMessage = parser.parseExpression(template, elDelim).getValue(context, String.class);
        // replace default message
        constraintValidatorContext.disableDefaultConstraintViolation();
        constraintValidatorContext.buildConstraintViolationWithTemplate(errorMessage).addConstraintViolation();

    }
}
```



[out-of-box]: https://www.baeldung.com/spring-validation-message-interpolation
[ConstraintValidatorContext]: https://docs.oracle.com/javaee/7/api/javax/validation/ConstraintValidatorContext.html
