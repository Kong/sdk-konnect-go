# SyncStatus

The outcome of the most recent attempt to fetch the skill's content from its source. `null` for sources whose content is supplied inline.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Status`                                                                             | [components.SkillSyncStatusStatus](../../models/components/skillsyncstatusstatus.md) | :heavy_check_mark:                                                                   | The state of the most recent synchronization attempt.                                |
| `Timestamp`                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                            | :heavy_check_mark:                                                                   | The time the synchronization reached this state.                                     |
| `ErrorMessage`                                                                       | `*string`                                                                            | :heavy_minus_sign:                                                                   | The reason the synchronization failed. `null` unless `status` is `failed`.           |