# PartialRedisEePort

An integer representing a port number between 0 and 65535, inclusive.


## Supported Types

### 

```go
partialRedisEePort := components.CreatePartialRedisEePortInteger(int64{/* values here */})
```

### 

```go
partialRedisEePort := components.CreatePartialRedisEePortStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch partialRedisEePort.Type {
	case components.PartialRedisEePortTypeInteger:
		// partialRedisEePort.Integer is populated
	case components.PartialRedisEePortTypeStr:
		// partialRedisEePort.Str is populated
}
```
