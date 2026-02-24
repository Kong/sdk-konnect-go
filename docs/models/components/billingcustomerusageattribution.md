# BillingCustomerUsageAttribution

Mapping to attribute metered usage to the customer by the event subject.


## Fields

| Field                                                                                                                         | Type                                                                                                                          | Required                                                                                                                      | Description                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `SubjectKeys`                                                                                                                 | []*string*                                                                                                                    | :heavy_check_mark:                                                                                                            | The subjects that are attributed to the customer.<br/>Can be empty when no usage event subjects are associated with the customer. |