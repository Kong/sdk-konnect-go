# APIPackageJSONPatchItem

A JSON Patch operation to add or remove an API Operation from a Package.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Op`                                                                             | [components.Op](../../models/components/op.md)                                   | :heavy_check_mark:                                                               | The operation to perform.                                                        |
| `Path`                                                                           | *string*                                                                         | :heavy_check_mark:                                                               | The path to the API Operation in the Package, in the format "operations/<uuid>". |
| `Value`                                                                          | [*components.Value](../../models/components/value.md)                            | :heavy_minus_sign:                                                               | N/A                                                                              |