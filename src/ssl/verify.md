# Verify Certificate
* online
* local cert file

## Certificate Chain
![cert chain](http://sslhowto.files.wordpress.com/2021/08/fc598-59b8e0_601624b8513544189ee32dd8d022659bmv2.png)

## Online Certificate

> View 

`echo -n | openssl s_client -connect google.com:443 -showcerts`  
* -text print decrypted cert content
* -notout/-nocert do not print cert content

> Download
1. download only site cert
`echo -n | openssl s_client -connect google.com:443 2>/dev/null | openssl x509 > site.cert`
1. download full chain
`$echo -n | openssl s_client -connect google.com:443 -showcerts | awk '/BEGIN CERTIFICATE/,/END CERTIFICATE/{if(/BEGIN/){a++}; out="cert"a".pem"; print > out }'`  
`openssl x509 -text -noout -in cert1.pem` to view each cert  
`openssl x509 -in cert3.pem -subject -subject_hash -issuer -issuer_hash -noout`
    * -subject/-issuer print data
    * -subject_hash/-issuer_hash print hash
    * cert1.pem => site cert
    * cert2.pem => intermedia cert in order
    * no root cert

## Verify _cert_ file

1. view and compare subject/issuer in order from downded  
`ls cert*.pem | sort | while read $f; do openssl x509 -in $f -subject -subject_hash -issuer -issuer_hash; done`  
1. verify site.cert aginst chain  
`cat cert2.pem cert3.pem > chain.pem && openssl verify -show_chain -CAfile chain.pem site.pem` 
