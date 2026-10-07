# BillingEntitlementHistoryWindow

Usage and balance of a single history window.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Period`                                               | [components.Period](../../models/components/period.md) | :heavy_check_mark:                                     | The period the window covers.                          |                                                        |
| `Usage`                                                | `string`                                               | :heavy_check_mark:                                     | The usage recorded in the window.                      | 100                                                    |
| `BalanceAtStart`                                       | `string`                                               | :heavy_check_mark:                                     | The entitlement balance at the start of the window.    | 100                                                    |