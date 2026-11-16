# OrgSubscriptionListResponse

A paginated list of the org subscriptions of one event type, or of the org subscriptions that hold one destination.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Data`                                                                     | [][components.OrgSubscription](../../models/components/orgsubscription.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `Meta`                                                                     | [components.CursorMeta](../../models/components/cursormeta.md)             | :heavy_check_mark:                                                         | Pagination metadata.                                                       |