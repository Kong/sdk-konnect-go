# EventGatewayParsedRecordTranscodeConsumeConfig

The configuration of the transcode parsed record policy applied to consumed records.


## Supported Types

### EventGatewayParsedRecordTranscodeConsumeConfigJSON

```go
eventGatewayParsedRecordTranscodeConsumeConfig := components.CreateEventGatewayParsedRecordTranscodeConsumeConfigJSON(components.EventGatewayParsedRecordTranscodeConsumeConfigJSON{/* values here */})
```

### EventGatewayParsedRecordTranscodeConsumeConfigAvro

```go
eventGatewayParsedRecordTranscodeConsumeConfig := components.CreateEventGatewayParsedRecordTranscodeConsumeConfigAvro(components.EventGatewayParsedRecordTranscodeConsumeConfigAvro{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordTranscodeConsumeConfig.Type {
	case components.EventGatewayParsedRecordTranscodeConsumeConfigTypeJSON:
		// eventGatewayParsedRecordTranscodeConsumeConfig.EventGatewayParsedRecordTranscodeConsumeConfigJSON is populated
	case components.EventGatewayParsedRecordTranscodeConsumeConfigTypeAvro:
		// eventGatewayParsedRecordTranscodeConsumeConfig.EventGatewayParsedRecordTranscodeConsumeConfigAvro is populated
}
```
