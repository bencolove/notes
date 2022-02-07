# AspectJ

Lib jar | Purpose
---|---
*aspectjrt.jar* | core and needed
*aspectjtools.jar* | _`ajc`_ compiler
*aspectjweaver.jar* | java agent

Weaving Methods | What
---|--
1. compile-time (ctw) | use _`ajc`_ with java and aj source 
1. post-compile (pcw) | use _`ajc`_ with compiled class files
1. load-time weave (ltw) | java agent from _aspecjweaver.jar_

## Setup
-- _`pom.xml`_ --
```xml
<dependency>
  <groupId>org.aspectj
  <artifactId>aspectjrt
  <version>

<build>
  <plugins>
    <plugin>
      <groupId>org.apache.maven.plugins</groupId>
      <artifactId>maven-compiler-plugin
      <version>3.8.7
      <configuration>
        <source>1.8
        <target>1.8
```

-- _AnnoAspectj.java_ --
```java
@Aspect
public class AnnoAspect {
    @Pointcut(
        "execution(* com.pkg.MyClass.say(..))"
    )
    public void jointPoint(){}

    @Before("jointPoint()") 
    public void before(){}

    @After("jointPoint()")
    public void after() {}
}
```

## Compile Time Weaving
You have to use _`ajc`_ instead of ordinary java compiler:  
`java -jar $Aspect_Tools -cp $Aspectj_Rt -source 1.5 -sourceroots . -d target/classes` 

And the _`aspectjrt.jar`_ has to be included:  
`java -cp $Aspectj_rt;. myapp`

## Post-compile Weaving
>1. compile all java files  
`javac -cp $AspectjRt -d target/classes path/to/*.java`

>2. ajc to weave in  
`java -jar $Aspectj_Tools -cp $AspectJ_Rt -source 1.5 -inpath target/classes -d target/classes`


## Load-time Weave
* java agent
* src/main/resources/META-INF/aop.xml
```xml
<aspectj>
  <aspects>
    <aspect name="com.pkg.AnnoAspect"/>
  </aspects>
<aspectj>
```
`java -javaagent:$Aspectj_Weaver -cp $Aspecj_Rt;taget/classes/ com.pkg.App`

