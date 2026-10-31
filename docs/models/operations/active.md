# Active


## Supported Types

###

```go
active := operations.CreateActiveBoolean(bool{/* values here */})
```

### StringFieldOEQFilter

```go
active := operations.CreateActiveStringFieldOEQFilter(components.StringFieldOEQFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch active.Type {
	case operations.ActiveTypeBoolean:
		// active.Boolean is populated
	case operations.ActiveTypeStringFieldOEQFilter:
		// active.StringFieldOEQFilter is populated
}
```
