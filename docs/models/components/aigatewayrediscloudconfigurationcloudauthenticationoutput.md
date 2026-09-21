# AIGatewayRedisCloudConfigurationCloudAuthenticationOutput

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthenticationOutput

```go
aiGatewayRedisCloudConfigurationCloudAuthenticationOutput := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationOutputAws(components.AIGatewayRedisAWSAuthenticationOutput{/* values here */})
```

### AIGatewayRedisAzureAuthenticationOutput

```go
aiGatewayRedisCloudConfigurationCloudAuthenticationOutput := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationOutputAzure(components.AIGatewayRedisAzureAuthenticationOutput{/* values here */})
```

### AIGatewayRedisGCPAuthenticationOutput

```go
aiGatewayRedisCloudConfigurationCloudAuthenticationOutput := components.CreateAIGatewayRedisCloudConfigurationCloudAuthenticationOutputGcp(components.AIGatewayRedisGCPAuthenticationOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayRedisCloudConfigurationCloudAuthenticationOutput.Type {
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationOutputTypeAws:
		// aiGatewayRedisCloudConfigurationCloudAuthenticationOutput.AIGatewayRedisAWSAuthenticationOutput is populated
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationOutputTypeAzure:
		// aiGatewayRedisCloudConfigurationCloudAuthenticationOutput.AIGatewayRedisAzureAuthenticationOutput is populated
	case components.AIGatewayRedisCloudConfigurationCloudAuthenticationOutputTypeGcp:
		// aiGatewayRedisCloudConfigurationCloudAuthenticationOutput.AIGatewayRedisGCPAuthenticationOutput is populated
}
```
