# CloudAuthentication

Metadata describing the cloud authentication details for managed cache add-on.



## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AuthProvider`                                        | `*string`                                             | :heavy_minus_sign:                                    | Env vault path to cache auth provider.                | {vault://env/ADDON_MANAGED_CACHE_AUTH_PROVIDER}       |
| `AwsCacheName`                                        | `*string`                                             | :heavy_minus_sign:                                    | Env vault path to aws cache name.                     | {vault://env/ADDON_MANAGED_CACHE_AWS_CACHE_NAME}      |
| `AwsRegion`                                           | `*string`                                             | :heavy_minus_sign:                                    | Env vault path to aws region.                         | {vault://env/ADDON_MANAGED_CACHE_AWS_REGION}          |
| `AwsAssumeRoleArn`                                    | `*string`                                             | :heavy_minus_sign:                                    | Env vault path to aws assume role arn.                | {vault://env/ADDON_MANAGED_CACHE_AWS_ASSUME_ROLE_ARN} |
| `AzureTenantID`                                       | `*string`                                             | :heavy_minus_sign:                                    | Env vault path to azure tenant id.                    | {vault://env/ADDON_MANAGED_CACHE_AZURE_TENANT_ID}     |