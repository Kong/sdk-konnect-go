# GetMFASettingsRequest

Request body for getting Multi-Factor Authentication (MFA) settings.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Email`                                                | `string`                                               | :heavy_check_mark:                                     | The email address of the user to get MFA settings for. | user@example.com                                       |
| `Auth0id`                                              | `string`                                               | :heavy_check_mark:                                     | The Auth0 ID of the user to get MFA settings for.      | auth0\|1234567890                                      |