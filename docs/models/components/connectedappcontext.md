# ConnectedAppContext

The connected app authorization context.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `State`                                                                    | `string`                                                                   | :heavy_check_mark:                                                         | The state parameter from the connected app authorization request.          |                                                                            |
| `Source`                                                                   | `string`                                                                   | :heavy_check_mark:                                                         | The source identifier of the connected app.                                |                                                                            |
| `Subdomain`                                                                | `string`                                                                   | :heavy_check_mark:                                                         | The subdomain of the authserver.                                           |                                                                            |
| `ConsentedScopeNames`                                                      | `*string`                                                                  | :heavy_minus_sign:                                                         | The (space-separated) list of consented scope names for the connected app. | konnect:read konnect:write                                                 |