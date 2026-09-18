# BillingSubscriptionEditAddPhase

Add a new phase to the subscription. The phase is created without items; use
add-item operations to populate it.


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                           | [components.BillingSubscriptionEditAddPhaseType](../../models/components/billingsubscriptioneditaddphasetype.md) | :heavy_check_mark:                                                                                               | Discriminator for the add-phase operation.                                                                       |
| `Phase`                                                                                                          | [components.Phase](../../models/components/phase.md)                                                             | :heavy_check_mark:                                                                                               | The phase to add.                                                                                                |