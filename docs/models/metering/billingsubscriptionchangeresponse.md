# BillingSubscriptionChangeResponse

Response for changing a subscription.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Current`                                            | [metering.Current](../../models/metering/current.md) | :heavy_check_mark:                                   | The current subscription before the change.          |
| `Next`                                               | [metering.Next](../../models/metering/next.md)       | :heavy_check_mark:                                   | The new state of the subscription after the change.  |