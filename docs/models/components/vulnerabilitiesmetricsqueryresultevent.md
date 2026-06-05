# VulnerabilitiesMetricsQueryResultEvent

Payload containing the aggregated metric data and associated dimensions included in the query.


## Supported Types

### VulnerabilityCountMetricEvent

```go
vulnerabilitiesMetricsQueryResultEvent := components.CreateVulnerabilitiesMetricsQueryResultEventVulnerabilityCountMetricEvent(components.VulnerabilityCountMetricEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch vulnerabilitiesMetricsQueryResultEvent.Type {
	case components.VulnerabilitiesMetricsQueryResultEventTypeVulnerabilityCountMetricEvent:
		// vulnerabilitiesMetricsQueryResultEvent.VulnerabilityCountMetricEvent is populated
}
```
