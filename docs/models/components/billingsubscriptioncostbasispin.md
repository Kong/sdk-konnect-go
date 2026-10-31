# BillingSubscriptionCostBasisPin

A cost basis pinned to a custom-currency pair for the subscription.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `CustomCurrencyID`                                       | `string`                                                 | :heavy_check_mark:                                       | The managed custom currency ID.                          | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `InvoiceCurrency`                                        | `string`                                                 | :heavy_check_mark:                                       | The fiat currency in which the subscription is invoiced. | USD                                                      |
| `CostBasisID`                                            | `string`                                                 | :heavy_check_mark:                                       | The pinned cost basis resource ID.                       | 01G65Z755AFWAKHE12NY0CQ9FH                               |