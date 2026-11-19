# AICostExportSchedule

An export schedule exports the organization's usage period by period, such as month by month. An `active` schedule creates one export of each period as the period ends. Each export is an ordinary export whose `schedule` names the schedule: it's listed, downloaded, and expires like a one-time export. Its files follow the dataset's rule: one for each day for `usage_events`, and one for each calendar month for `focus`. Each export of a `daily` `usage_events` schedule or a `monthly` `focus` schedule therefore has one file.

- The export of a period has the period's start as `starts_at` and its end as `ends_at`, and the schedule's other settings as they are when the export is created.
- When a schedule becomes `active`, by creation or by resuming, its first period is the one in progress. That period's export covers the whole period, including the time before the schedule became `active`.
- A schedule exports each period once and never refreshes it, even when events of the period arrive later. A scheduled export that fails stays `failed`, and the schedule doesn't create it again. To include later events or to fill a failed period, a client creates a one-time export of the same range.
- When runs are late, for example after an outage, the schedule catches up on at most the 31 most recent periods that ended while it was `active`, and skips older ones.


## Supported Types

### AICostExportSchedule1

```go
aiCostExportSchedule := components.CreateAICostExportScheduleAICostExportSchedule1(components.AICostExportSchedule1{/* values here */})
```

### Two

```go
aiCostExportSchedule := components.CreateAICostExportScheduleTwo(components.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostExportSchedule.Type {
	case components.AICostExportScheduleTypeAICostExportSchedule1:
		// aiCostExportSchedule.AICostExportSchedule1 is populated
	case components.AICostExportScheduleTypeTwo:
		// aiCostExportSchedule.Two is populated
}
```
