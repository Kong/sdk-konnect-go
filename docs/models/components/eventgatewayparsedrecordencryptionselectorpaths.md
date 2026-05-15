# EventGatewayParsedRecordEncryptionSelectorPaths

Selects which fields of the parsed record to encrypt. A maximum of 50 path entries are allowed.


## Supported Types

### 

```go
eventGatewayParsedRecordEncryptionSelectorPaths := components.CreateEventGatewayParsedRecordEncryptionSelectorPathsArrayOfPaths1([]components.Paths1{/* values here */})
```

### 

```go
eventGatewayParsedRecordEncryptionSelectorPaths := components.CreateEventGatewayParsedRecordEncryptionSelectorPathsStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordEncryptionSelectorPaths.Type {
	case components.EventGatewayParsedRecordEncryptionSelectorPathsTypeArrayOfPaths1:
		// eventGatewayParsedRecordEncryptionSelectorPaths.ArrayOfPaths1 is populated
	case components.EventGatewayParsedRecordEncryptionSelectorPathsTypeStr:
		// eventGatewayParsedRecordEncryptionSelectorPaths.Str is populated
}
```
