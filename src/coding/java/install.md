## Installation
[linux][instruction] | [windows][java-win]

Downloads:
Vendor\Version | 8 | 11
---|---|---
Sun | [8][java-se-8] | [11][java-se-11]
OpenJDK | [8][openjdk-archive] | [11][openjdk-archive]


```bash

$ sudo apt install openjdk-11-jre-headless

$ sudo apt install openjdk-8-jdk-headless

# manage multiple java envs and set default one

$ sudo update-alternatives --config java

#Output
There are 2 choices for the alternative java (providing /usr/bin/java).

  Selection    Path                                            Priority   Status
------------------------------------------------------------
  0            /usr/lib/jvm/java-11-openjdk-amd64/bin/java      1111      auto mode
  1            /usr/lib/jvm/java-11-openjdk-amd64/bin/java      1111      manual mode
* 2            /usr/lib/jvm/java-11-oracle/bin/java             1091      manual mode

```

## Set JAVA_HOME for All Users
_`/etc/environment`_
>Append:  
>JAVA_HOME="/usr/lib/jvm/java-11-openjdk-amd64"

`source /etc/environment`


[instruction]: https://www.digitalocean.com/community/tutorials/how-to-install-java-with-apt-on-ubuntu-18-04

[java-win]: https://www.happycoders.eu/java/how-to-switch-multiple-java-versions-windows/

[java-se-8]: https://www.oracle.com/java/technologies/javase/javase8-archive-downloads.html

[java-se-11]: https://www.oracle.com/java/technologies/javase/jdk11-archive-downloads.html

[openjdk-archive]: https://jdk.java.net/archive/