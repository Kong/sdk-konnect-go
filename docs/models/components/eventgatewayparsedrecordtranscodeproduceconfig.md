# EventGatewayParsedRecordTranscodeProduceConfig

The configuration of the transcode parsed record policy applied to produced records.


## Supported Types

### EventGatewayParsedRecordTranscodeProduceConfigJSON

```go
eventGatewayParsedRecordTranscodeProduceConfig := components.CreateEventGatewayParsedRecordTranscodeProduceConfigJSON(components.EventGatewayParsedRecordTranscodeProduceConfigJSON{/* values here */})
```

### EventGatewayParsedRecordTranscodeProduceConfigAvro

```go
eventGatewayParsedRecordTranscodeProduceConfig := components.CreateEventGatewayParsedRecordTranscodeProduceConfigAvro(components.EventGatewayParsedRecordTranscodeProduceConfigAvro{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordTranscodeProduceConfig.Type {
	case components.EventGatewayParsedRecordTranscodeProduceConfigTypeJSON:
		// eventGatewayParsedRecordTranscodeProduceConfig.EventGatewayParsedRecordTranscodeProduceConfigJSON is populated
	case components.EventGatewayParsedRecordTranscodeProduceConfigTypeAvro:
		// eventGatewayParsedRecordTranscodeProduceConfig.EventGatewayParsedRecordTranscodeProduceConfigAvro is populated
}
```
