# BackendClusterAuthenticationSaslAwsIamAssumeRole

Configures whether to authenticate using credentials obtained by first assuming a role, using the AWS default credentials provider chain

**Requires a minimum runtime version of `1.3`**.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Type`                                                                         | `string`                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `AssumeRole`                                                                   | [components.AssumeRole](../../models/components/assumerole.md)                 | :heavy_check_mark:                                                             | Configuration for assuming an IAM role. Required when `type` is `assume_role`. |