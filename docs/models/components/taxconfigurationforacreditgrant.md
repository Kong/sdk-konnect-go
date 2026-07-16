# TaxConfigurationForACreditGrant

Tax configuration for the grant.

For `invoice` and `external` funding methods, tax configuration should be
provided to ensure correct revenue recognition. When not provided, the default
credit grant tax code is applied, if that's not set the global default taxcode
is used.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Behavior`                                                  | [*components.Behavior](../../models/components/behavior.md) | :heavy_minus_sign:                                          | Tax behavior applied to the invoice line item.              |
| `TaxCode`                                                   | [*components.TaxCode](../../models/components/taxcode.md)   | :heavy_minus_sign:                                          | Tax code applied to the invoice line item.                  |