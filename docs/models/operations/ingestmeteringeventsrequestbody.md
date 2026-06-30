# IngestMeteringEventsRequestBody


## Supported Types

### MeteringEvent

```go
ingestMeteringEventsRequestBody := operations.CreateIngestMeteringEventsRequestBodyMeteringEvent(components.MeteringEvent{/* values here */})
```

### 

```go
ingestMeteringEventsRequestBody := operations.CreateIngestMeteringEventsRequestBodyArrayOfMeteringEvent([]components.MeteringEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ingestMeteringEventsRequestBody.Type {
	case operations.IngestMeteringEventsRequestBodyTypeMeteringEvent:
		// ingestMeteringEventsRequestBody.MeteringEvent is populated
	case operations.IngestMeteringEventsRequestBodyTypeArrayOfMeteringEvent:
		// ingestMeteringEventsRequestBody.ArrayOfMeteringEvent is populated
}
```
