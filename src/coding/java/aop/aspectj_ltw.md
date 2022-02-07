# AspectJ Load-Time Weaving
Setup:
1. aop definitions(java file with @Aspect)
1. aspectj javaagent `-javaagent:path/to/aspectjweaver.jar`
1. config file `META-INF/aop.xml`

After compile-weaving/post-compile-weaving:
1. singleton patter per aspect class by `class.aspectOf()`
1. execution join point
```java
// origin method
public String origin(String value) {
    return "hello " + value;
}

// woven method
public String originNameButWoven(String value) {
    JoinPoint jp = Factory.makeJP(singleton, this, this, value);
    return (String)origin_aroundBody1$advice(this, value, jp, AspectClass.aspectOf(), (ProceedingJoinPoint)jp);
}

private static final Object origin_aroundBody1$advice(CurrentInstance ajc$this, String value, JoinPoint thisJoinPoint, AspectClass ajc$aspectInstance, ProceedingJoinPoint proceedingJoinPoint) {
    // nearly copy of aspect around method
    // before
    String ret = origin_aroundBody0$(ajc$this, value, (JoinPoint)proceedingJoinPoint)
    // after
}

private static final String origin_aroundBody0(CurrentInstance ajc$this, String value, JoinPoint paramJoinPoint) {
    // nearly copy of origin method
    // but rewritten as a static one
}

```

## the `aop.xml` Config File
[[ltw config][ltw-config]]  
-- _`META-INF/aop.mxl`_ --
```xml
<aspectj>
    <aspects>
        <!-- ① 切面類 @Aspect的文件-->
        <aspect name="com.dsc.spos.utils.stats.StatsAop"/>
    </aspects>

    <weaver options="-verbose -showWeaveInfo"> <!-- -debug"> -->
        <!-- ② 指定需要進行織入操作的目標類範圍 -->
        <!-- 一定要包括上面那步定義aspcet的文件 -->
        <include within="com.dsc.spos.utils.stats.StatsAop"/>

        <include within="com.dsc.spos.restfulservice.SPosRestfulService"/>
        <include within="com.dsc.spos.service.utils.DispatchService"/>
        <include within="com.dsc.spos.service.imp.json.GoodsFeatureGetDCP*"/>
        <include within="com.dsc.spos.utils.PosPub"/>

        <!-- 把runtime 修改的class輸出到 ./_ajdump目錄 -->
        <dump within="com.dsc.spos.utils.stats.StatsAop" beforeandafter="true" />
        <dump within="com.dsc.spos.service.utils.DispatchService" beforeandafter="true" />
        <dump within="com.dsc.spos.service.imp.json.GoodsFeatureGetDCP*" beforeandafter="true"/>

        <!-- 不碰 CGLIB 修改過的class -->
        <exclude within="com.foo.bar..*CGLIB*"/>
    </weaver>

</aspectj>
```
[ltw-config]: https://www.eclipse.org/aspectj/doc/released/devguide/ltw-configuration.html