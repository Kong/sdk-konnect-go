# AIGatewayModelVectorDBConfigRedisCloudAuthenticationOutput

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthenticationOutput

```go
aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationOutputAws(components.AIGatewayRedisAWSAuthenticationOutput{/* values here */})
```

### AIGatewayRedisAzureAuthenticationOutput

```go
aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationOutputAzure(components.AIGatewayRedisAzureAuthenticationOutput{/* values here */})
```

### AIGatewayRedisGCPAuthenticationOutput

```go
aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationOutputGcp(components.AIGatewayRedisGCPAuthenticationOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput.Type {
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationOutputTypeAws:
		// aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput.AIGatewayRedisAWSAuthenticationOutput is populated
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationOutputTypeAzure:
		// aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput.AIGatewayRedisAzureAuthenticationOutput is populated
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationOutputTypeGcp:
		// aiGatewayModelVectorDBConfigRedisCloudAuthenticationOutput.AIGatewayRedisGCPAuthenticationOutput is populated
}
```
