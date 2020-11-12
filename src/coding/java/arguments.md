# System properties VS Environment variables

Means | System Properties | Environment Variables
---|---|---
Programatically | System.getProperties() | System.getenv()
Command Line | -Dpropertyname=value | export/SET 

## SpringBoot Executable
_application.properties_ are loaded as `System.properties` by key `spring.config.location`.

* override whole config file  
`java -Dspring.config.location=xxx -jar springboot.jar`
* override single property in _application.properties_  
`java -jar springboot.jar --prop=value`