# Java Keytool Commands
[Source][source]

## Keystore VS Truststore
When communicating over SSL/TLS, a keystore/truststore is needed. These are password-protected files storing encrpyted **private keys** and **certificates**.

Format of keystore/truststore:
* __JKS__, Before and including JAVA 8, a JAVA specific format
* __PKCS12__, Since JAVA 9, a language-netural format

Both keystore/truststore store same piece of info:
* Private key entries,
* Certificates with public keys,
* Secret Keys

So JAVA does not essentially treat them any differently.

```bash
$ keytool -list -keystore cacerts
Enter keystore password:
Keystore type: JKS
Keystore Provider: SUN

Your keystore contains 92 entries
 
verisignclass2g2ca [jdk], 2018-06-13, trustedCertEntry,
Certificate fingerprint (SHA1): B3:EA:C4:47:76:C9:C8:1C:EA:F2:9D:95:B6:CC:A0:08:1B:67:EC:9D
...
```

### JAVA Keystore
A keystore is used to identify **US**.  

No default keystore, javax.net.ssl.KeyStore and 
javax.net.ssl.KeyStorePassword could be used if needed.
[Programmatically deal with a Keystore][program-keystore]

### JAVA Truststore
As opposite to Keystore, a Truststore is used to identify **others**.
It stores as the Keystore.

JRE comes with a bundled truststore **cacerts** located in 
    `%JRE%/lib/security/cacerts`

### How are They Used

Client ---> HTTPS ---> Server (1.)
(2.)

1. Server presents Client with its credential from its **Keystore**  
    1.1 Server lookup an associated key client requested for  
    1.2 Server presents client with its public key and certificates

2. Client authenticates Server with its **Truststore**  
    2.1 Client lookup an associated certificate from its Truststore  
    2.2 __SSLHandshakeException__ will be thrown if not matched found 


## Checking
* Check a stand-alone certificate

`$ keytool -printcert -v -file mydomain.crt`

* Check which certificates are in a Java keystore

`$ keytool -list -v -keystore keystore.jks`

* Check a particular keystore entry using an alias

`$ keytool -list -v -keystore keystore.jks -alias mydomain`

## Creating and Importing
* Generate a Java keystore and key pair  

`$ keytool -genkey -alias mydomain -keyalg RSA -keystore keystore.jks  -keysize 2048`  

* Generate a certificate signing request (CSR) for an existing Java keystore  

`$ keytool -certreq -alias mydomain -keystore keystore.jks -file mydomain.csr`  

* Import a root or intermediate CA certificate to an existing Java keystore

`$ keytool -import -trustcacerts -alias root -file Thawte.crt -keystore keystore.jks`  

* Import a signed primary certificate to an existing Java keystore  

`$ keytool -import -trustcacerts -alias mydomain -file mydomain.crt -keystore keystore.jks`

* Generate a keystore and self-signed certificate (see How to Create a Self Signed Certificate using Java Keytoolfor more info)  

`$ keytool -genkey -keyalg RSA -alias selfsigned -keystore keystore.jks -storepass password -validity 360 -keysize 2048`


## Other Keytool Commands
* Delete a certificate from a Java Keytool keystore

`$ keytool -delete -alias mydomain -keystore keystore.jks`

* Change a Java keystore password

`$ keytool -storepasswd -new new_storepass -keystore keystore.jks`

* Export a certificate from a keystore

`$ keytool -export -alias mydomain -file mydomain.crt -keystore keystore.jks`

* List Trusted CA Certs

`$ keytool -list -v -keystore $JAVA_HOME/jre/lib/security/cacerts`

* Import New CA into Trusted Certs  

`$ keytool -import -trustcacerts -file /path/to/ca/ca.pem -alias
CA_ALIAS -keystore $JAVA_HOME/jre/lib/security/cacerts`




[source]: https://www.sslshopper.com/article-most-common-java-keytool-keystore-commands.html
[program-keystore]: https://www.baeldung.com/java-keystore