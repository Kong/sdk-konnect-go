# SchedulerState


## Supported Types

### ResourceSchedulerStateOk

```go
schedulerState := components.CreateSchedulerStateResourceSchedulerStateOk(components.ResourceSchedulerStateOk{/* values here */})
```

### ResourceSchedulerStateNotOk

```go
schedulerState := components.CreateSchedulerStateResourceSchedulerStateNotOk(components.ResourceSchedulerStateNotOk{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch schedulerState.Type {
	case components.SchedulerStateTypeResourceSchedulerStateOk:
		// schedulerState.ResourceSchedulerStateOk is populated
	case components.SchedulerStateTypeResourceSchedulerStateNotOk:
		// schedulerState.ResourceSchedulerStateNotOk is populated
}
```
