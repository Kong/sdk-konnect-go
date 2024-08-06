# SNIInput


## Fields

| Field                                                                                                                                                                       | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                                                                      | **string*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                          | The SNI name to associate with the given certificate.                                                                                                                       |
| `Tags`                                                                                                                                                                      | []*string*                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                          | An optional set of strings associated with the SNIs for grouping and filtering.                                                                                             |
| `Certificate`                                                                                                                                                               | [*components.SNICertificate](../../models/components/snicertificate.md)                                                                                                     | :heavy_minus_sign:                                                                                                                                                          | The id (a UUID) of the certificate with which to associate the SNI hostname. The Certificate must have a valid private key associated with it to be used by the SNI object. |