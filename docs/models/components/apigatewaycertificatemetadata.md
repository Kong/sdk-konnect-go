# APIGatewayCertificateMetadata

Metadata extracted from a certificate.


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Issuer`                                                | **string*                                               | :heavy_minus_sign:                                      | The issuer of the certificate.                          |
| `Subject`                                               | **string*                                               | :heavy_minus_sign:                                      | The subject of the certificate.                         |
| `KeyUsages`                                             | []*string*                                              | :heavy_minus_sign:                                      | The key usages of the certificate.                      |
| `Expiry`                                                | **int64*                                                | :heavy_minus_sign:                                      | The expiry date of the certificate as a unix timestamp. |
| `SanNames`                                              | []*string*                                              | :heavy_minus_sign:                                      | The Subject Alternative Names (SAN) of the certificate. |
| `DNSNames`                                              | []*string*                                              | :heavy_minus_sign:                                      | The DNS names in the certificate SAN.                   |
| `EmailAddresses`                                        | []*string*                                              | :heavy_minus_sign:                                      | The email addresses in the certificate SAN.             |
| `IPAddresses`                                           | []*string*                                              | :heavy_minus_sign:                                      | The IP addresses in the certificate SAN.                |
| `Uris`                                                  | []*string*                                              | :heavy_minus_sign:                                      | The URIs in the certificate SAN.                        |
| `Sha256Fingerprint`                                     | **string*                                               | :heavy_minus_sign:                                      | The SHA-256 fingerprint of the certificate.             |