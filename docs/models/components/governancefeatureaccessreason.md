# GovernanceFeatureAccessReason

Optional reason when the customer does not have access to the feature. Populated
when `has_access` is `false`.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Code`                                                                                           | [components.GovernanceFeatureAccessCode](../../models/components/governancefeatureaccesscode.md) | :heavy_check_mark:                                                                               | Machine-readable error code.                                                                     |
| `Message`                                                                                        | `string`                                                                                         | :heavy_check_mark:                                                                               | Human-readable description of the error.                                                         |
| `Attributes`                                                                                     | map[string]`any`                                                                                 | :heavy_minus_sign:                                                                               | Additional structured context.                                                                   |