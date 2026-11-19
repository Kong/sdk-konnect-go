# ReportedCost

The producer's estimate of the invocation's cost, such as `{"value": "0.0054", "currency": "USD"}`. Usage queries, forecasts, and budgets sum it as reported.

`value` is `0` to `999999999.999999999`, with at most 9 digits before the decimal point and at most 9 after it. It has no sign, no exponent, and no extra leading zeros (`0.5` is valid, `00.5` isn't). Trailing fractional zeros are accepted, and the event list returns the value without them.

An omitted cost is unknown, not zero. A producer that knows the cost is zero sends a `value` of `"0"`. The event list returns an omitted cost with a `value` of `"0"`.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Value`                                                                | `string`                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `Currency`                                                             | [components.AICostCurrency](../../models/components/aicostcurrency.md) | :heavy_check_mark:                                                     | The currency of `value`.                                               |