# BillingSubscriptionCancel

Request for canceling a subscription.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Timing`                                                  | [*metering.Timing](../../models/metering/timing.md)       | :heavy_minus_sign:                                        | If not provided the subscription is canceled immediately. | immediate                                                 |