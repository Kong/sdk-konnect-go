# OauthConfig


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Type`                                                  | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |                                                         |
| `ClientID`                                              | *string*                                                | :heavy_check_mark:                                      | The OAuth client identifier.                            | d745213a-b7e8-4998-abe3-41f164001970                    |
| `ClientSecret`                                          | *string*                                                | :heavy_check_mark:                                      | The OAuth client secret.                                | s3cr3t4p1cl13ntt0k3n1234567890abcdef                    |
| `AuthorizationEndpoint`                                 | *string*                                                | :heavy_check_mark:                                      | The URL where users are redirected to authorize access. | https://identity.service.com/oauth/authorize            |
| `TokenEndpoint`                                         | *string*                                                | :heavy_check_mark:                                      | The URL used to retrieve access tokens.                 | https://identity.service.com/oauth/token                |