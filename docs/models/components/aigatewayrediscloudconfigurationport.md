# AIGatewayRedisCloudConfigurationPort

An integer representing a port number between 0 and 65535, inclusive.
This field is [referenceable](https://developer.konghq.com/gateway/entities/vault/#how-do-i-reference-secrets-stored-in-a-vault).



## Supported Types

### 

```go
aiGatewayRedisCloudConfigurationPort := components.CreateAIGatewayRedisCloudConfigurationPortInteger(int64{/* values here */})
```

### 

```go
aiGatewayRedisCloudConfigurationPort := components.CreateAIGatewayRedisCloudConfigurationPortStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayRedisCloudConfigurationPort.Type {
	case components.AIGatewayRedisCloudConfigurationPortTypeInteger:
		// aiGatewayRedisCloudConfigurationPort.Integer is populated
	case components.AIGatewayRedisCloudConfigurationPortTypeStr:
		// aiGatewayRedisCloudConfigurationPort.Str is populated
}
```
