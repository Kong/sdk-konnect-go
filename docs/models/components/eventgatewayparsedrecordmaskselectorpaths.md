# EventGatewayParsedRecordMaskSelectorPaths

Selects which fields of the parsed record to mask. A maximum of 50 path entries are allowed.


## Supported Types

### 

```go
eventGatewayParsedRecordMaskSelectorPaths := components.CreateEventGatewayParsedRecordMaskSelectorPathsArrayOfEventGatewayParsedRecordFieldPathsArray([]components.EventGatewayParsedRecordFieldPathsArray{/* values here */})
```

### 

```go
eventGatewayParsedRecordMaskSelectorPaths := components.CreateEventGatewayParsedRecordMaskSelectorPathsStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayParsedRecordMaskSelectorPaths.Type {
	case components.EventGatewayParsedRecordMaskSelectorPathsTypeArrayOfEventGatewayParsedRecordFieldPathsArray:
		// eventGatewayParsedRecordMaskSelectorPaths.ArrayOfEventGatewayParsedRecordFieldPathsArray is populated
	case components.EventGatewayParsedRecordMaskSelectorPathsTypeStr:
		// eventGatewayParsedRecordMaskSelectorPaths.Str is populated
}
```
