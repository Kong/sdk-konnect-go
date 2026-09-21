# AuthorizeConnectedAppRequest

Request to authorize a connected app.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Consent`                                                                        | `bool`                                                                           | :heavy_check_mark:                                                               | Whether to authorize the connected app.                                          | true                                                                             |
| `Capps`                                                                          | [components.ConnectedAppContext](../../models/components/connectedappcontext.md) | :heavy_check_mark:                                                               | The connected app authorization context.                                         |                                                                                  |