# CreateAddOnRequest

Request schema for creating an add-on.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Name`                                                                       | `string`                                                                     | :heavy_check_mark:                                                           | Unique human-readable name of the add-on.                                    | my-add-on                                                                    |
| `Owner`                                                                      | [components.AddOnOwner](../../models/components/addonowner.md)               | :heavy_check_mark:                                                           | Owner for the add-on.                                                        |                                                                              |
| `Config`                                                                     | [components.CreateAddOnConfig](../../models/components/createaddonconfig.md) | :heavy_check_mark:                                                           | Configuration for creating different types of add-ons.                       |                                                                              |