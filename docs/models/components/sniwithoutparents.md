# SNIWithoutParents

An SNI object represents a many-to-one mapping of hostnames to a certificate. That is, a certificate object can have many hostnames associated with it; when Kong receives an SSL request, it uses the SNI field in the Client Hello to lookup the certificate object based on the SNI associated with the certificate.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ID`                                                                            | **string*                                                                       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Name`                                                                          | *string*                                                                        | :heavy_check_mark:                                                              | The SNI name to associate with the given certificate.                           |
| `Tags`                                                                          | []*string*                                                                      | :heavy_minus_sign:                                                              | An optional set of strings associated with the SNIs for grouping and filtering. |