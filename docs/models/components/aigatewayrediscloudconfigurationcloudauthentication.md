# AIGatewayRedisCloudConfigurationCloudAuthentication

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthentication

```go
aiGatewayRedisCloudConfigurationCloudAuthentication := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationAws(components.AIGatewayRedisAWSAuthentication{/* values here */})
```

### AIGatewayRedisAzureAuthentication

```go
aiGatewayRedisCloudConfigurationCloudAuthentication := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationAzure(components.AIGatewayRedisAzureAuthentication{/* values here */})
```

### AIGatewayRedisGCPAuthentication

```go
aiGatewayRedisCloudConfigurationCloudAuthentication := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationGcp(components.AIGatewayRedisGCPAuthentication{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayRedisCloudConfigurationCloudAuthentication.Type {
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationTypeAws:
		// aiGatewayRedisCloudConfigurationCloudAuthentication.AIGatewayRedisAWSAuthentication is populated
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationTypeAzure:
		// aiGatewayRedisCloudConfigurationCloudAuthentication.AIGatewayRedisAzureAuthentication is populated
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationTypeGcp:
		// aiGatewayRedisCloudConfigurationCloudAuthentication.AIGatewayRedisGCPAuthentication is populated
}
```
