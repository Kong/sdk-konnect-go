# OrganizationFeature

A response for getting or updating a feature configuration for an organization.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | `*string`                                                        | :heavy_minus_sign:                                               | The unique identifier for the feature configuration.             |                                                                  |
| `Name`                                                           | [components.FeatureName](../../models/components/featurename.md) | :heavy_check_mark:                                               | N/A                                                              |                                                                  |
| `Enabled`                                                        | `bool`                                                           | :heavy_check_mark:                                               | Whether this feature is enabled for the organization.            |                                                                  |
| `CreatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | An ISO-8601 timestamp representation of entity creation date.    | 2022-11-04T20:10:06.927Z                                         |
| `UpdatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | An ISO-8601 timestamp representation of entity update date.      | 2022-11-04T20:10:06.927Z                                         |