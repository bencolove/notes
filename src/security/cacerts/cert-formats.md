# Formats of Certificate
[Source][source]  
X.509 v3 certificates conform with RFC 5280.
Commonly referred to as PKIX for Public Key Infrastructure X.509

Formats of certificate:
* DER
* .CRT
* .CER
* PEM

They should be distuiguished essentially by:
* DER encoded: binary encoded, may bear .CER or .CRT extension
* PEM encoded: ASCII (Base64) encoded, armored data prefixed with `--- BEGIN...` line

## View PEM encoded certificate
`$ openssl x509 -in cert.cer|crt|pem -text -noout`  
## View DER encoded certificate  
`$ openssl x509 -in cert.der -inform der -text -noout`  
NOTE the key opt `-inform der`    

## Convert
* PEM to DER  

`$ openssl x509 -in cert.crt -outform der -out cert.der`  

* DER to PEM  

`$ openssl x509 -in cert.crt -inform der -outform pem -out cert.pem`  

Extension:
* CET mainly on Linux, may be DER or PEM encoded
* CER mainly on Windows, may be any encoding

## Example
Exported by Chrome a Certificate as PEM (Base64 encoded), outouted file as:  
_hncb.cer_
```bash
-----BEGIN CERTIFICATE-----
MIIICjCCBvKgAwIBAgIQR+MAAAAABMBAav0g2yKflDANBgkqhkiG9w0BAQsFADBz
MQswCQYDVQQGEwJUVzESMBAGA1UEChMJVEFJV0FOLUNBMRwwGgYDVQQLExNHbG9i
YWwgRVZTU0wgU3ViLUNBMTIwMAYDVQQDEylUV0NBIEdsb2JhbCBFVlNTTCBDZXJ0
aWZpY2F0aW9uIEF1dGhvcml0eTAeFw0xOTAxMTgxMDQyMzNaFw0yMTAzMTYxNTU5
NTlaMIIBWDELMAkGA1UEBhMCVFcxDzANBgNVBAgTBlRhaXdhbjEPMA0GA1UEBxMG
VGFpcGVpMSQwIgYDVQQKExtIVUEgTkFOIENPTU1FUkNJQUwgQkFOSyBMVEQxDDAK
BgNVBAsTA01JUzEbMBkGA1UEAxMSZXBvc250LmhuY2IuY29tLnR3MR0wGwYDVQQP
ExRQcml2YXRlIE9yZ2FuaXphdGlvbjEXMBUGCysGAQQBgjc8AgEBEwZUYWlwZWkx
FzAVBgsrBgEEAYI3PAIBAhMGVGFpd2FuMRMwEQYLKwYBBAGCNzwCAQMTAlRXMREw
DwYDVQQFEwgwMzc0MjMwMTFNMEsGA1UECRNETm8uMTIzLCBTb25ncmVuIFJkLiwg
WGlueWkgRGlzdC4sIFRhaXBlaSBDaXR5IDExMDEwLCBUYWl3YW4gKFIuTy5DLikx
DjAMBgNVBBETBTExMDEwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
qKfDpXkYqBOPPdiJivPHJ/uBLj63BBhrabiX3zDYrpY1R5iz/978nqyBgLj7jSIH
yLOLzMiM8f42YrQPgt00E+y9diIzVs0zL1WwFwUieiYBEQTxkrcHuNsEEegtwTq2
YHU2tX6pfEdb5O6TgXADeNkxN5fdG+s1ZykfWb1pBo4O0Hq3J7Aqyu0f7kBX2uUT
PvHNHyE11H1qQSnoRfh031W4yRZrvuFzGn1LKrrFhUTlHXKOdhCpfZRUPo8yO3Pg
wI1wjBwjE2E89EFdwJcNbFc88Ez51wtX7gy8oqvjfH9MSxr085pGgTlnziWiQmYQ
yly7+kTYzbo8bre1CGbWJwIDAQABo4IDsTCCA60wHwYDVR0jBBgwFoAUbr2hK87k
wtUodFy92YxvBHIqBt4wHQYDVR0OBBYEFIUYoU3zYE1xUPyx96tZ869Yw2Q7MFMG
A1UdHwRMMEowSKBGoESGQmh0dHA6Ly9zc2xzZXJ2ZXIudHdjYS5jb20udHcvc3Ns
c2VydmVyL0dsb2JhbEVWU1NMX1Jldm9rZV8yMDEyLmNybDAdBgNVHREEFjAUghJl
cG9zbnQuaG5jYi5jb20udHcwfwYIKwYBBQUHAQEEczBxMEQGCCsGAQUFBzAChjho
dHRwOi8vc3Nsc2VydmVyLnR3Y2EuY29tLnR3L2NhY2VydC9HbG9iYWxFdnNzbF8y
MDEyLnA3YjApBggrBgEFBQcwAYYdaHR0cDovL2V2c3Nsb2NzcC50d2NhLmNvbS50
dy8wPwYDVR0gBDgwNjA0BgwrBgEEAYK/JQEBFgMwJDAiBggrBgEFBQcCARYWaHR0
cDovL3d3dy50d2NhLmNvbS50dzAJBgNVHRMEAjAAMA4GA1UdDwEB/wQEAwIFoDAd
BgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwggH5BgorBgEEAdZ5AgQCBIIB
6QSCAeUB4wB2AId1v+dZfPiMQ5lfvfNu/1aNR1Y2/0q1YMG06v9eoIMPAAABaGCO
IXcAAAQDAEcwRQIhAIycFTPzGvp4Iulq88TbIkjEyMMIFn8ZcHlRz3AW4prfAiBA
8ljf0k6D11IG4A+pWw/zoMmviun3VJQ/mMUro3f5LAB3AFWB1MIWkDYBSuoLm1c8
U/DA5Dh4cCUIFy+jqh0HE9MMAAABaGCOIq0AAAQDAEgwRgIhAKmUUpybpmQymzdc
VAg88GijCwuQiNjR9l6/MuKXDhOYAiEAma9FCp404JTbEKr/JeEIveDLACfgIwPu
4+QCc1vC7EwAdwBvU3asMfAxGdiZAKRRFf93FRwR2QLBACkGjbIImjfZEwAAAWhg
jiKYAAAEAwBIMEYCIQDHkuog6v0iqHmMD0ycKkG2q4aBziefrehHS606oSSh9AIh
AKDGxJb7fRDDR0O+f2w4h7rkSSxQXwOCCOge46rXI6f4AHcApLkJkLQYWBSHuxOi
zGdwCjw1mAT5G9+443fNDsgN3BAAAAFoYI4eUQAABAMASDBGAiEA4xh+L6Li7jLi
PSK1CRAZCNUW4aWOviiTrE9qk/nD5iACIQDAu+C3d/rkadooIZB2gK9BBNyM8o61
9eCeUKTBA/NNETANBgkqhkiG9w0BAQsFAAOCAQEADxYSZcEXytIhyKiMIMr3vB+4
iOMWGerxxvrngnRf4FLu9dP0PxxYv876akoNAYlocVXGdeYAq0M/SvD0lDc2d8jc
qvx4a5uCsXsL839L4A0ri9yaX3vdNhVum6SOZqUzwfsb3vGXM1GzIK3tcQgqsjct
peVqS/kJNJqCfc86+6KqoULyRq9oRvdi+5vpNOHqSr3RrFxEeuDsjwk+iANZKJBn
likdVuuXfzOJ1smaTVjzZl2siJ++V1NKmj0PCcxrJLEJCkIl2jaW7/XkgSFmXqrc
S1iaLItmf+EqdXgElr2MjjVesPNRiY5h3xI9sYPKkMbeUPgHamfFHkItgHxySQ==
-----END CERTIFICATE-----

```

View its decrypted content by:  
`$ openssl x509 -in hncb.cert -inform pem -noout -text `  
```bash
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            47:e3:00:00:00:00:04:c0:40:6a:fd:20:db:22:9f:94
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = TW, O = TAIWAN-CA, OU = Global EVSSL Sub-CA, CN = TWCA Global EVSSL Certification Authority
        Validity
            Not Before: Jan 18 10:42:33 2019 GMT
            Not After : Mar 16 15:59:59 2021 GMT
        Subject: C = TW, ST = Taiwan, L = Taipei, O = HUA NAN COMMERCIAL BANK LTD, OU = MIS, CN = eposnt.hncb.com.tw, businessCategory = Private Organization, jurisdictionL = Taipei, jurisdictionST = Taiwan, jurisdictionC = TW, serialNumber = 03742301, street = "No.123, Songren Rd., Xinyi Dist., Taipei City 11010, Taiwan (R.O.C.)", postalCode = 11010
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:a8:a7:c3:a5:79:18:a8:13:8f:3d:d8:89:8a:f3:
                    c7:27:fb:81:2e:3e:b7:04:18:6b:69:b8:97:df:30:
                    d8:ae:96:35:47:98:b3:ff:de:fc:9e:ac:81:80:b8:
                    fb:8d:22:07:c8:b3:8b:cc:c8:8c:f1:fe:36:62:b4:
                    0f:82:dd:34:13:ec:bd:76:22:33:56:cd:33:2f:55:
                    b0:17:05:22:7a:26:01:11:04:f1:92:b7:07:b8:db:
                    04:11:e8:2d:c1:3a:b6:60:75:36:b5:7e:a9:7c:47:
                    5b:e4:ee:93:81:70:03:78:d9:31:37:97:dd:1b:eb:
                    35:67:29:1f:59:bd:69:06:8e:0e:d0:7a:b7:27:b0:
                    2a:ca:ed:1f:ee:40:57:da:e5:13:3e:f1:cd:1f:21:
                    35:d4:7d:6a:41:29:e8:45:f8:74:df:55:b8:c9:16:
                    6b:be:e1:73:1a:7d:4b:2a:ba:c5:85:44:e5:1d:72:
                    8e:76:10:a9:7d:94:54:3e:8f:32:3b:73:e0:c0:8d:
                    70:8c:1c:23:13:61:3c:f4:41:5d:c0:97:0d:6c:57:
                    3c:f0:4c:f9:d7:0b:57:ee:0c:bc:a2:ab:e3:7c:7f:
                    4c:4b:1a:f4:f3:9a:46:81:39:67:ce:25:a2:42:66:
                    10:ca:5c:bb:fa:44:d8:cd:ba:3c:6e:b7:b5:08:66:
                    d6:27
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Authority Key Identifier:
                keyid:6E:BD:A1:2B:CE:E4:C2:D5:28:74:5C:BD:D9:8C:6F:04:72:2A:06:DE

            X509v3 Subject Key Identifier:
                85:18:A1:4D:F3:60:4D:71:50:FC:B1:F7:AB:59:F3:AF:58:C3:64:3B
            X509v3 CRL Distribution Points:

                Full Name:
                  URI:http://sslserver.twca.com.tw/sslserver/GlobalEVSSL_Revoke_2012.crl

            X509v3 Subject Alternative Name:
                DNS:eposnt.hncb.com.tw
            Authority Information Access:
                CA Issuers - URI:http://sslserver.twca.com.tw/cacert/GlobalEvssl_2012.p7b
                OCSP - URI:http://evsslocsp.twca.com.tw/

            X509v3 Certificate Policies:
                Policy: 1.3.6.1.4.1.40869.1.1.22.3
                  CPS: http://www.twca.com.tw

            X509v3 Basic Constraints:
                CA:FALSE
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Server Authentication, TLS Web Client Authentication
            CT Precertificate SCTs:
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 87:75:BF:E7:59:7C:F8:8C:43:99:5F:BD:F3:6E:FF:56:
                                8D:47:56:36:FF:4A:B5:60:C1:B4:EA:FF:5E:A0:83:0F
                    Timestamp : Jan 18 10:42:33.975 2019 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:21:00:8C:9C:15:33:F3:1A:FA:78:22:E9:6A:
                                F3:C4:DB:22:48:C4:C8:C3:08:16:7F:19:70:79:51:CF:
                                70:16:E2:9A:DF:02:20:40:F2:58:DF:D2:4E:83:D7:52:
                                06:E0:0F:A9:5B:0F:F3:A0:C9:AF:8A:E9:F7:54:94:3F:
                                98:C5:2B:A3:77:F9:2C
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 55:81:D4:C2:16:90:36:01:4A:EA:0B:9B:57:3C:53:F0:
                                C0:E4:38:78:70:25:08:17:2F:A3:AA:1D:07:13:D3:0C
                    Timestamp : Jan 18 10:42:34.285 2019 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:46:02:21:00:A9:94:52:9C:9B:A6:64:32:9B:37:5C:
                                54:08:3C:F0:68:A3:0B:0B:90:88:D8:D1:F6:5E:BF:32:
                                E2:97:0E:13:98:02:21:00:99:AF:45:0A:9E:34:E0:94:
                                DB:10:AA:FF:25:E1:08:BD:E0:CB:00:27:E0:23:03:EE:
                                E3:E4:02:73:5B:C2:EC:4C
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 6F:53:76:AC:31:F0:31:19:D8:99:00:A4:51:15:FF:77:
                                15:1C:11:D9:02:C1:00:29:06:8D:B2:08:9A:37:D9:13
                    Timestamp : Jan 18 10:42:34.264 2019 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:46:02:21:00:C7:92:EA:20:EA:FD:22:A8:79:8C:0F:
                                4C:9C:2A:41:B6:AB:86:81:CE:27:9F:AD:E8:47:4B:AD:
                                3A:A1:24:A1:F4:02:21:00:A0:C6:C4:96:FB:7D:10:C3:
                                47:43:BE:7F:6C:38:87:BA:E4:49:2C:50:5F:03:82:08:
                                E8:1E:E3:AA:D7:23:A7:F8
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : A4:B9:09:90:B4:18:58:14:87:BB:13:A2:CC:67:70:0A:
                                3C:35:98:04:F9:1B:DF:B8:E3:77:CD:0E:C8:0D:DC:10
                    Timestamp : Jan 18 10:42:33.169 2019 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:46:02:21:00:E3:18:7E:2F:A2:E2:EE:32:E2:3D:22:
                                B5:09:10:19:08:D5:16:E1:A5:8E:BE:28:93:AC:4F:6A:
                                93:F9:C3:E6:20:02:21:00:C0:BB:E0:B7:77:FA:E4:69:
                                DA:28:21:90:76:80:AF:41:04:DC:8C:F2:8E:B5:F5:E0:
                                9E:50:A4:C1:03:F3:4D:11
    Signature Algorithm: sha256WithRSAEncryption
         0f:16:12:65:c1:17:ca:d2:21:c8:a8:8c:20:ca:f7:bc:1f:b8:
         88:e3:16:19:ea:f1:c6:fa:e7:82:74:5f:e0:52:ee:f5:d3:f4:
         3f:1c:58:bf:ce:fa:6a:4a:0d:01:89:68:71:55:c6:75:e6:00:
         ab:43:3f:4a:f0:f4:94:37:36:77:c8:dc:aa:fc:78:6b:9b:82:
         b1:7b:0b:f3:7f:4b:e0:0d:2b:8b:dc:9a:5f:7b:dd:36:15:6e:
         9b:a4:8e:66:a5:33:c1:fb:1b:de:f1:97:33:51:b3:20:ad:ed:
         71:08:2a:b2:37:2d:a5:e5:6a:4b:f9:09:34:9a:82:7d:cf:3a:
         fb:a2:aa:a1:42:f2:46:af:68:46:f7:62:fb:9b:e9:34:e1:ea:
         4a:bd:d1:ac:5c:44:7a:e0:ec:8f:09:3e:88:03:59:28:90:67:
         96:29:1d:56:eb:97:7f:33:89:d6:c9:9a:4d:58:f3:66:5d:ac:
         88:9f:be:57:53:4a:9a:3d:0f:09:cc:6b:24:b1:09:0a:42:25:
         da:36:96:ef:f5:e4:81:21:66:5e:aa:dc:4b:58:9a:2c:8b:66:
         7f:e1:2a:75:78:04:96:bd:8c:8e:35:5e:b0:f3:51:89:8e:61:
         df:12:3d:b1:83:ca:90:c6:de:50:f8:07:6a:67:c5:1e:42:2d:
         80:7c:72:49
```

[source]: https://support.ssl.com/Knowledgebase/Article/View/19/0/der-vs-crt-vs-cer-vs-pem-certificates-and-how-to-convert-them