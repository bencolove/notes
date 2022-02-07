# Modify Java Byte-Code
Utilize Javaassist
```java
byte[] transform(
    ClassLoader loader,
    String className,
    Class<?> classBeingRedefined,
    ProtectionDomain protectionDomain,
    byte[] classfileBuffer
) {
    // what we do here is to based on `classfileBuffer` return new byte[] of the class

    // check target class and the loader are identical

    ClassPool cp = ClassPool.getDefault();
    CtClass cc = cp.get(targetClassName);
    CtMethod m = cc.getDeclaredMethod("methodName");

    m.addLocalVariable("startTime", CtClass.longType);
    m.insertBefore(
        "startTime = System.currentTimeMillis();"
    )

    m.addLocalVariable("endTime", CtClass.longType);
    m.addLocalVariable("opTime", CtClass.longType);
    m.insertAfter(
        "endTime = System.currentTimeMillis();" +
        "opTime = (endTime - startTime) / 1000;" +
        "//log sth"
    );

    // modified
    byte[] modifiedBytecode = cc.toBytecode();
    cc.detach();
    return modifiedBytecode;
}


```