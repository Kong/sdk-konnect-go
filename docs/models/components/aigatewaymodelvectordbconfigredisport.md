# AIGatewayModelVectorDBConfigRedisPort

An integer representing a port number between 0 and 65535, inclusive.
This field is [referenceable](https://developer.konghq.com/gateway/entities/vault/#how-do-i-reference-secrets-stored-in-a-vault).



## Supported Types

### 

```go
aiGatewayModelVectorDBConfigRedisPort := components.CreateAIGatewayModelVectorDBConfigRedisPortInteger(int64{/* values here */})
```

### 

```go
aiGatewayModelVectorDBConfigRedisPort := components.CreateAIGatewayModelVectorDBConfigRedisPortStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelVectorDBConfigRedisPort.Type {
	case components.AIGatewayModelVectorDBConfigRedisPortTypeInteger:
		// aiGatewayModelVectorDBConfigRedisPort.Integer is populated
	case components.AIGatewayModelVectorDBConfigRedisPortTypeStr:
		// aiGatewayModelVectorDBConfigRedisPort.Str is populated
}
```
