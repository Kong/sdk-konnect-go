# BackendClusterAuthenticationSaslAwsIam

AWS IAM-based OAUTHBEARER authentication scheme for the backend cluster, for example when connecting to
Amazon MSK with IAM authentication.

**Requires a minimum runtime version of `1.3`**.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Type`                                                         | `string`                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `SaslAwsIam`                                                   | [components.SaslAwsIam](../../models/components/saslawsiam.md) | :heavy_check_mark:                                             | N/A                                                            |