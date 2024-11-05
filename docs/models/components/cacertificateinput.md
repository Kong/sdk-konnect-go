# CACertificateInput

A CA certificate object represents a trusted CA. These objects are used by Kong to verify the validity of a client or server certificate.


## Fields

| Field                                                                                                                                         | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `Cert`                                                                                                                                        | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | PEM-encoded public certificate of the CA.                                                                                                     |
| `CertDigest`                                                                                                                                  | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | SHA256 hex digest of the public certificate. This field is read-only and it cannot be set by the caller, the value is automatically computed. |
| `ID`                                                                                                                                          | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | N/A                                                                                                                                           |
| `Tags`                                                                                                                                        | []*string*                                                                                                                                    | :heavy_minus_sign:                                                                                                                            | An optional set of strings associated with the Certificate for grouping and filtering.                                                        |