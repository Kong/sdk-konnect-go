# MultiKeyAuth1

Defines an authentication strategy based on one or more API keys passed via HTTP headers.
This strategy supports integrations that require custom headers for credential-based access,
allowing flexibility across providers with different authentication header requirements.



## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Type`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `Config`                                                                       | [components.MultiKeyAuthConfig](../../models/components/multikeyauthconfig.md) | :heavy_check_mark:                                                             | N/A                                                                            |