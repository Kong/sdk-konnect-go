# EventsAttributesFieldFilter

Filters the events by values contained within the `attributes` field. Filters must use dot-notation to identify
the details property that will be used to filter the results.



## Supported Types

### StringFieldFilter

```go
eventsAttributesFieldFilter := components.CreateEventsAttributesFieldFilterStringFieldFilter(components.StringFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsAttributesFieldFilter.Type {
	case components.EventsAttributesFieldFilterTypeStringFieldFilter:
		// eventsAttributesFieldFilter.StringFieldFilter is populated
}
```
