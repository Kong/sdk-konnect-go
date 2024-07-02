# Sni

An SNI object represents a many-to-one mapping of hostnames to a certificate. That is, a certificate object can have many hostnames associated with it; when Kong receives an SSL request, it uses the SNI field in the Client Hello to lookup the certificate object based on the SNI associated with the certificate.


## Fields

| Field                                                                                                                                                                       | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                                                                      | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | The SNI name to associate with the given certificate.                                                                                                                       |
| `Tags`                                                                                                                                                                      | []*string*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                          | An optional set of strings associated with the SNIs for grouping and filtering.                                                                                             |
| `Certificate`                                                                                                                                                               | [components.SNICertificate](../../models/components/snicertificate.md)                                                                                                      | :heavy_check_mark:                                                                                                                                                          | The id (a UUID) of the certificate with which to associate the SNI hostname. The Certificate must have a valid private key associated with it to be used by the SNI object. |
| `CreatedAt`                                                                                                                                                                 | **int64*                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                          | Unix epoch when the resource was created.                                                                                                                                   |
| `ID`                                                                                                                                                                        | **string*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                          | N/A                                                                                                                                                                         |