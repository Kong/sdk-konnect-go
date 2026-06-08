# EventGatewayParsedRecordEncryptionSelectorPaths

Selects which fields of the parsed record to encrypt. A maximum of 50 path entries are allowed.


## Supported Types

### 

```go
eventGatewayParsedRecordEncryptionSelectorPaths := components.CreateEventGatewayParsedRecordEncryptionSelectorPathsArrayOfEventGatewayParsedRecordFieldPathsArray([]components.EventGatewayParsedRecordFieldPathsArray{/* values here */})
```

### 

```go
eventGatewayParsedRecordEncryptionSelectorPaths := components.CreateEventGatewayParsedRecordEncryptionSelectorPathsStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordEncryptionSelectorPaths.Type {
	case components.EventGatewayParsedRecordEncryptionSelectorPathsTypeArrayOfEventGatewayParsedRecordFieldPathsArray:
		// eventGatewayParsedRecordEncryptionSelectorPaths.ArrayOfEventGatewayParsedRecordFieldPathsArray is populated
	case components.EventGatewayParsedRecordEncryptionSelectorPathsTypeStr:
		// eventGatewayParsedRecordEncryptionSelectorPaths.Str is populated
}
```
