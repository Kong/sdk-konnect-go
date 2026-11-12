# CreateDcrSettings

Dynamic Client Registration settings for the auth server. Both fields are optional; an omitted field falls back to its default.


## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `DefaultAccessTokenDuration`                                                                              | `*int64`                                                                                                  | :heavy_minus_sign:                                                                                        | The default access token duration, in seconds, applied to DCR clients registered against this auth server |