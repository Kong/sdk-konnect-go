# Port

An integer representing a port number between 0 and 65535, inclusive.


## Supported Types

### 

```go
port := components.CreatePortInteger(int64{/* values here */})
```

### 

```go
port := components.CreatePortStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch port.Type {
	case components.PortTypeInteger:
		// port.Integer is populated
	case components.PortTypeStr:
		// port.Str is populated
}
```
