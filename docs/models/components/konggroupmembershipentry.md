# KongGroupMembershipEntry

A single group membership entry, referencing either a principal or another group.



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ID`                                                                     | `string`                                                                 | :heavy_check_mark:                                                       | Contains a unique identifier used for this resource.                     | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                     |
| `Member`                                                                 | [components.KongGroupMember](../../models/components/konggroupmember.md) | :heavy_check_mark:                                                       | Identifies a single member (principal or group) of a group.<br/>         |                                                                          |
| `CreatedAt`                                                              | [time.Time](https://pkg.go.dev/time#Time)                                | :heavy_check_mark:                                                       | An ISO-8601 timestamp representation of entity creation date.            | 2022-11-04T20:10:06.927Z                                                 |