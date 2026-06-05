# BillingCreditBalances

The balances of the credits of a customer.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `RetrievedAt`                                                          | [time.Time](https://pkg.go.dev/time#Time)                              | :heavy_check_mark:                                                     | The timestamp of the balance retrieval.                                | 2023-01-01T01:01:01.001Z                                               |
| `Balances`                                                             | [][components.CreditBalance](../../models/components/creditbalance.md) | :heavy_check_mark:                                                     | The balances by currencies.                                            |                                                                        |