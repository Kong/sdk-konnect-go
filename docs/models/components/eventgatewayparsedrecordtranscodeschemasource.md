# EventGatewayParsedRecordTranscodeSchemaSource

Determines how to look up the schema to use for the transcoded output data.
Leave this unset if the output data schema isn't needed.



## Supported Types

### EventGatewayParsedRecordTranscodeSchemaSourceReference

```go
eventGatewayParsedRecordTranscodeSchemaSource := components.CreateEventGatewayParsedRecordTranscodeSchemaSourceReference(components.EventGatewayParsedRecordTranscodeSchemaSourceReference{/* values here */})
```

### EventGatewayParsedRecordTranscodeSchemaSourceInline

```go
eventGatewayParsedRecordTranscodeSchemaSource := components.CreateEventGatewayParsedRecordTranscodeSchemaSourceInline(components.EventGatewayParsedRecordTranscodeSchemaSourceInline{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordTranscodeSchemaSource.Type {
	case components.EventGatewayParsedRecordTranscodeSchemaSourceTypeReference:
		// eventGatewayParsedRecordTranscodeSchemaSource.EventGatewayParsedRecordTranscodeSchemaSourceReference is populated
	case components.EventGatewayParsedRecordTranscodeSchemaSourceTypeInline:
		// eventGatewayParsedRecordTranscodeSchemaSource.EventGatewayParsedRecordTranscodeSchemaSourceInline is populated
}
```
