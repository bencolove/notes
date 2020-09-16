# Install Tomcat as Service

1. OpenJDK
2. Create Tomcat user
3. Install Tomcat
4. Manager GUI

---

## Create Tomcat User (Service Account)
`sudo useradd -r -m -U -d /opt/tomcat -s /bin/false tomcat`  
* -r --system
* -m --create-home          : create /home/user
* -d --home-dir HOME_DIR    : set as home dir
* -U --user-group           : create a group with same name
* -s --shell                : set login shell

## Download Tomcat9.0
```bash
wget http://www-eu.apache.org/dist/tomcat/tomcat-9/v9.0.37/bin/apache-tomcat-9.0.37.tar.gz -P /tmp
```
## Install Tomcat
`sudo tar xf /tmp/apache-tomcat-9*.tar.gz -C /opt/tomcat`  

## Version Management
`sudo ln -s /opt/tomcat/apache-tomcat-9.0.37 /opt/tomcat/latest`  

## File Persmissions for `tomcat`
`sudo chown -RH tomcat: /opt/tomcat/latest`
* -R --recursive
* -H:   traverse the argument if symbolic link

`sudo sh -c 'chmod +x /opt/tomcat/latest/bin/*.sh'`  
>`sudo chmod +x /opt/tomcat/latest/bin/*.sh`  
This way *.sh will not be globbed !??? why

## Make `tomcat` as Service
1. Service File  
-- _`/etc/systemd/system/tomcat.service`_ --  
```txt
[Unit]
Description=Tomcat 9 servlet container
After=network.target

[Service]
Type=forking

User=tomcat
Group=tomcat

Environment="JAVA_HOME=/usr/lib/jvm/default-java"
Environment="JAVA_OPTS=-Djava.security.egd=file:///dev/urandom -Djava.awt.headless=true"

Environment="CATALINA_BASE=/opt/tomcat/latest"
Environment="CATALINA_HOME=/opt/tomcat/latest"
Environment="CATALINA_PID=/opt/tomcat/latest/temp/tomcat.pid"
Environment="CATALINA_OPTS=-Xms512M -Xmx1024M -server -XX:+UseParallelGC"

ExecStart=/opt/tomcat/latest/bin/startup.sh
ExecStop=/opt/tomcat/latest/bin/shutdown.sh

[Install]
WantedBy=multi-user.target
```

2. Reload `systemd`
```bash
$ sudo systemctl daemon-reload
$ sudo systemctl start tomcat
$ sudo systemctl status tomcat
$ sudo systemctl enable tomcat
```

3. Firewall  
`sudo ufw allow 8080/tcp`  

## Management GUI
1. Manager User  
-- _`/opt/tomcat/latest/conf/tomcat-users.xml`_ --
```txt
<tomcat-users>
<!--
    Comments
-->
   <role rolename="admin-gui"/>
   <role rolename="manager-gui"/>
   <user username="admin" password="admin_password" roles="admin-gui,manager-gui"/>
</tomcat-users>
```

2. Allow for Non-localhost Access  
* manager app: _`$TOMCAT_HOME/webapps/manager/META-INF/context.xml`_ 
* host-manager app: _`$TOMCAT_HOME/webapps/host-manager/META-INF/context.xml`_ 

>Allow ALL
```
<Context antiResourceLocking="false" privileged="true" >
<!--
  <Valve className="org.apache.catalina.valves.RemoteAddrValve"
         allow="127\.\d+\.\d+\.\d+|::1|0:0:0:0:0:0:0:1" />
-->
</Context>
```

>White-list:  
```
<Context antiResourceLocking="false" privileged="true" >
  <Valve className="org.apache.catalina.valves.RemoteAddrValve"
         allow="127\.\d+\.\d+\.\d+|::1|0:0:0:0:0:0:0:1|45.45.45.45" />
</Context>
```

3. Restart  
`$ sudo systemctl restart tomcat`  


