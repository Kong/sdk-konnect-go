# TaxIDCollection

Configuration for collecting tax IDs during checkout.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Enabled`                                                     | `*bool`                                                       | :heavy_minus_sign:                                            | Enable tax ID collection during checkout.<br/><br/>Defaults to false. |
| `Required`                                                    | [*components.Required](../../models/components/required.md)   | :heavy_minus_sign:                                            | Whether tax ID collection is required.<br/><br/>Defaults to "never". |