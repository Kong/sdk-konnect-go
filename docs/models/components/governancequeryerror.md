# GovernanceQueryError

Query error within a partially successful governance query response.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Code`                                                             | [components.Code](../../models/components/code.md)                 | :heavy_check_mark:                                                 | Machine-readable error code.                                       |
| `Message`                                                          | `string`                                                           | :heavy_check_mark:                                                 | Human-readable description of the error.                           |
| `Attributes`                                                       | map[string]`any`                                                   | :heavy_minus_sign:                                                 | Additional structured context.                                     |
| `Customer`                                                         | `*string`                                                          | :heavy_minus_sign:                                                 | The customer identifier from the request that produced this error. |