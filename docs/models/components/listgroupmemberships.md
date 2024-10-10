# ListGroupMemberships

A paginated list response for a collection of control plane group memberships.


## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `Meta`                                                                                                            | [*components.CursorPaginatedMetaWithSizeAndTotal](../../models/components/cursorpaginatedmetawithsizeandtotal.md) | :heavy_minus_sign:                                                                                                | returns the pagination information                                                                                |
| `Data`                                                                                                            | [][components.ControlPlaneSummary](../../models/components/controlplanesummary.md)                                | :heavy_minus_sign:                                                                                                | Array of control planes summary who are a child to this control plane group.                                      |