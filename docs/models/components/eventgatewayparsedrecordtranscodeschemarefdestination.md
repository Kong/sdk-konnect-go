# EventGatewayParsedRecordTranscodeSchemaRefDestination

Defines how to record the schema id for the transcoded output data. See the
[Confluent docs](https://docs.confluent.io/platform/current/schema-registry/fundamentals/serdes-develop/index.html#wire-format)
for more about the wire format.



## Supported Types

### EventGatewayParsedRecordTranscodeSchemaRefDestinationConfluentFormat

```go
eventGatewayParsedRecordTranscodeSchemaRefDestination := components.CreateEventGatewayParsedRecordTranscodeSchemaRefDestinationConfluentFormat(components.EventGatewayParsedRecordTranscodeSchemaRefDestinationConfluentFormat{/* values here */})
```

### EventGatewayParsedRecordTranscodeSchemaRefDestinationRecordHeader

```go
eventGatewayParsedRecordTranscodeSchemaRefDestination := components.CreateEventGatewayParsedRecordTranscodeSchemaRefDestinationRecordHeader(components.EventGatewayParsedRecordTranscodeSchemaRefDestinationRecordHeader{/* values here */})
```

### EventGatewayParsedRecordTranscodeSchemaRefDestinationNone

```go
eventGatewayParsedRecordTranscodeSchemaRefDestination := components.CreateEventGatewayParsedRecordTranscodeSchemaRefDestinationNone(components.EventGatewayParsedRecordTranscodeSchemaRefDestinationNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordTranscodeSchemaRefDestination.Type {
	case components.EventGatewayParsedRecordTranscodeSchemaRefDestinationTypeConfluentFormat:
		// eventGatewayParsedRecordTranscodeSchemaRefDestination.EventGatewayParsedRecordTranscodeSchemaRefDestinationConfluentFormat is populated
	case components.EventGatewayParsedRecordTranscodeSchemaRefDestinationTypeRecordHeader:
		// eventGatewayParsedRecordTranscodeSchemaRefDestination.EventGatewayParsedRecordTranscodeSchemaRefDestinationRecordHeader is populated
	case components.EventGatewayParsedRecordTranscodeSchemaRefDestinationTypeNone:
		// eventGatewayParsedRecordTranscodeSchemaRefDestination.EventGatewayParsedRecordTranscodeSchemaRefDestinationNone is populated
}
```
