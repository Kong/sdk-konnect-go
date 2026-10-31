# EventGatewayParsedRecordTranscodeProduceConfig

The configuration of the transcode parsed record policy applied to produced records.


## Supported Types

### EventGatewayParsedRecordTranscodeProduceConfigAvro

```go
eventGatewayParsedRecordTranscodeProduceConfig := components.CreateEventGatewayParsedRecordTranscodeProduceConfigAvro(components.EventGatewayParsedRecordTranscodeProduceConfigAvro{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordTranscodeProduceConfig.Type {
	case components.EventGatewayParsedRecordTranscodeProduceConfigTypeAvro:
		// eventGatewayParsedRecordTranscodeProduceConfig.EventGatewayParsedRecordTranscodeProduceConfigAvro is populated
}
```
