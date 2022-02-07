1. instrumentation API
1. Java Attach API
1. Javassist

## Loading Agent
1. static (load at JVM boot time)
2. dynamic (attach to running JVM)

> Static Load  
`java -javaagent:agent.jar -jar application.jar`

> Dynamic Load  
```java
public class Launcher {
    public static void main(String[] args) throws Exception {
        ...
        VirtualMachine jvm = VirtualMachine.attach(jvmPID);
        jvm.loadAgent(agentFile.getAbsolutePath());
        jvm.detach();
    }
}
```

## Create Agent
By utilizing Java Instrument API methods:
* addTransformer
* getAllLoadedClasses
* retransformClasses
* removeTransformer
* redefineClasses

Define for a agent:
* premain 
* agentmain

```java
static void premain(String args, Instrumentation inst) { processClass(inst); }

static void agentmain(String args, Instrumentation inst) { processClass(inst); }

private static void processClass(Instrumentation inst) {
    String className = "full.qualified.name";

    Class<?> targetClass = null;
    ClassLoader targetClassLoader = null;

    try {
        // if worked
        targetClass = Class.forName(className);
        transform(targetClass, targetClass.getClassLoader, inst);
        return;
    } catch(Exception ex) {}

    // otherwise, look from loaded
    for (Class<?> clazz: inst.getAllLoadedClasses()) {
        if (clazz.getName().equals(className)) {
            targetClass = clazz;
            transform(clazz, clazz.getClassLoader, inst);
            return;
        }
    }
    // ooops
    throw new RuntimeException();
}

private static void transform() {
    Class<?> clazz,
    ClassLoader classLoader,
    Instrumentation inst) {
        MyTransformer ts = new MyTransformer(
            clazz, classLoader
        );

        // instrument API
        inst.addTransformer(ts, true);

        try {
            // instrument API
            inst.restransformClases(clazz);
        } catch(Exception ex) {}
    }
}

// --- instrument API ---
// custom transformer class
public class MyTransformer implements ClassFileTransformer {
    @Overriede
    public byte[] transform(
        ClassLoader loader,
        String className,
        Class<?> classBeingRedefined,
        ProtectionDomain protectionDomain,
        byte[] classfileBuffer
    ) {}
}

```

## Package into agent.jar
[[package agent jar][instrut_package]]
Package agent class files into a jar with MANIFEST.MF:
```txt
Agent-Class: my.agent.entryclass
Can-Redefine-classes: true
Can-Retransform-Classes: true
Premain-Class: my.agent.entryclass
```

[instrut_package]: https://docs.oracle.com/en/java/javase/11/docs/api/java.instrument/java/lang/instrument/package-summary.html