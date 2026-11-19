# ReportedCacheCost

The producer's estimate of the part of `reported_cost` that comes from cache reads and cache writes. Its `value` has the same format and bounds as `reported_cost`'s.

An omitted value is unknown, not zero. The event list returns it with a `value` of `"0"`.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Value`                                                                | `string`                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `Currency`                                                             | [components.AICostCurrency](../../models/components/aicostcurrency.md) | :heavy_check_mark:                                                     | The currency of `value`.                                               |