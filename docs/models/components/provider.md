# Provider

Description of cloud provider that runs cloud gateway instances.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Provider`                                                                              | [components.ProviderName](../../models/components/providername.md)                      | :heavy_check_mark:                                                                      | Name of cloud provider.                                                                 | aws                                                                                     |
| `Regions`                                                                               | [][components.ProviderRegion](../../models/components/providerregion.md)                | :heavy_check_mark:                                                                      | List of available regions to run cloud gateway instances on for a given cloud provider. |                                                                                         |