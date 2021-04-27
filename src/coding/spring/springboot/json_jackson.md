# Json De/Serialization with Jackson
Out of box:
1. @JsonProperty on field, read/write directly
1. @JsonSetter/@JsonGetter on field/method (via accessors)

Conversion:  
[[@JsonSerialize/@JsonDeserialize][json-deser]]
1. `using` -- convert from/to directly from JSON(string)
1. `converter` -- convert from/to after de/serialization

>Example:  
"2021-04-21" deserialize to "20210421"
```java
// usage
@JsonSetter("JsonPropertyName")
@JsonDeserialze(using = DatetimeStripper.class)
// pass optional parameter to @JsonDeserialize
@DatetimeStripperAttributes(trim="-")
String datetime;

// implement
public class DatetimeStripper 
extends StdDeserializer<String> 
implements ContextualDeserializer {

    private String trim;

    protected DatetimeStripper(String trim) {
        super(String.class);
        this.trim = trim;
    }

    /*
    A way to get @DatetimeStripperAttributes(trim="-")
    */
    @Override
    public JsonDeserializer<?> createContextual(DeserializationContext deserializationContext, BeanProperty beanProperty) 
    throws JsonMappingException {
        String trim = "";
        if (beanProperty != null) {
            // get parameter annotation
           DatetimeStripperAttributes anno = beanProperty.getAnnotation(DatetimeStripperAttributes.class);
           trim = anno.trim()
        }

        // pass to deserializer
        return new DatetimeStripper(trim);
    }

    @Override
    public String deserialize(JsonParser jsonParser, DeserializationContext deserializationContext) 
    throws IOException, JsonProcessingException {
        // string from JSON
        String value = jsonParser.getValueAsString();
        // convert
        return value == null ? value : value.replaceAll(trim,"");
    }
```

[json-deser]: https://juconcurrent.com/2016/11/03/jackson-serializer/