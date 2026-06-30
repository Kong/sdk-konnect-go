# ListGroupMemberships

A paginated list response for a collection of control plane group memberships.


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `Meta`                                                                                                           | [components.CursorPaginatedMetaWithSizeAndTotal](../../models/components/cursorpaginatedmetawithsizeandtotal.md) | :heavy_check_mark:                                                                                               | returns the pagination information                                                                               |
| `Data`                                                                                                           | [][components.ControlPlane](../../models/components/controlplane.md)                                             | :heavy_check_mark:                                                                                               | Array of control planes summary who are a child to this control plane group.                                     |