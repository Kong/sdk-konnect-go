# UpsertNode

The request schema to upsert a Node.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Hostname`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            | node1.my.domain.com                                            |
| `Type`                                                         | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            | knep                                                           |
| `Version`                                                      | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            | 1.0                                                            |
| `ConfigVersion`                                                | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            |                                                                |
| `Errors`                                                       | [][components.NodeError](../../models/components/nodeerror.md) | :heavy_minus_sign:                                             | N/A                                                            |                                                                |