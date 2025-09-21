# Data Generation

You will run into an issue when you try to use Fabric Datagen. Something along the lines of

```
[21:08:08] [main/ERROR] (FabricDataGenHelper) Failed to run data generation
 java.lang.RuntimeException: Failed to run data generator from mod (porting_lib_tool_actions)
	at net.fabricmc.fabric.impl.datagen.FabricDataGenHelper.runInternal(FabricDataGenHelper.java:150) ~[fabric-data-generation-api-v1-ceaf4ab3-12.3.7+1802ada577.jar:?]
    ...
```

This is caused by you not specifying which mod the datagen is supposed to generate. In this case, it generates
everything, which you don't want. Simply add `-Dfabric-api.datagen.modid=<yourmodidhere>` to your Datagen VM
options.

![datagen highlighted in IntelliJ VM options](/static/img/cfa/datagen-info.png)