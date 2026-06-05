# PostOauthDeviceTokenRequestBody

The request schema for the device access token request.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `GrantType`                                                                         | `string`                                                                            | :heavy_check_mark:                                                                  | Value MUST be set to "urn:ietf:params:oauth:grant-type:device_code".                |
| `DeviceCode`                                                                        | `string`                                                                            | :heavy_check_mark:                                                                  | The device verification code, "device_code" from the device authorization response. |
| `ClientID`                                                                          | `string`                                                                            | :heavy_check_mark:                                                                  | The client identifier.                                                              |