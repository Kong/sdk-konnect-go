# BillingSubscriptionChangeResponse

Response for changing a subscription.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Current`                                                | [components.Current](../../models/components/current.md) | :heavy_check_mark:                                       | The current subscription before the change.              |
| `Next`                                                   | [components.Next](../../models/components/next.md)       | :heavy_check_mark:                                       | The new state of the subscription after the change.      |