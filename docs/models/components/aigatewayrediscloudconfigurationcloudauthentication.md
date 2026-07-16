# AIGatewayRedisCloudConfigurationCloudAuthentication

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthenticationOutput

```go
aiGatewayRedisCloudConfigurationCloudAuthentication := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationAws(components.AIGatewayRedisAWSAuthenticationOutput{/* values here */})
```

### AIGatewayRedisAzureAuthenticationOutput

```go
aiGatewayRedisCloudConfigurationCloudAuthentication := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationAzure(components.AIGatewayRedisAzureAuthenticationOutput{/* values here */})
```

### AIGatewayRedisGCPAuthenticationOutput

```go
aiGatewayRedisCloudConfigurationCloudAuthentication := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationGcp(components.AIGatewayRedisGCPAuthenticationOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayRedisCloudConfigurationCloudAuthentication.Type {
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationTypeAws:
		// aiGatewayRedisCloudConfigurationCloudAuthentication.AIGatewayRedisAWSAuthenticationOutput is populated
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationTypeAzure:
		// aiGatewayRedisCloudConfigurationCloudAuthentication.AIGatewayRedisAzureAuthenticationOutput is populated
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationTypeGcp:
		// aiGatewayRedisCloudConfigurationCloudAuthentication.AIGatewayRedisGCPAuthenticationOutput is populated
}
```
