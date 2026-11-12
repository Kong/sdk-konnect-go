# RedisEEDatastoreConfigPort

An integer representing a port number between 0 and 65535, inclusive.
This field is [referenceable](https://developer.konghq.com/gateway/entities/vault/#how-do-i-reference-secrets-stored-in-a-vault).



## Supported Types

### 

```go
redisEEDatastoreConfigPort := components.CreateRedisEEDatastoreConfigPortInteger(int64{/* values here */})
```

### 

```go
redisEEDatastoreConfigPort := components.CreateRedisEEDatastoreConfigPortStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch redisEEDatastoreConfigPort.Type {
	case components.RedisEEDatastoreConfigPortTypeInteger:
		// redisEEDatastoreConfigPort.Integer is populated
	case components.RedisEEDatastoreConfigPortTypeStr:
		// redisEEDatastoreConfigPort.Str is populated
}
```
