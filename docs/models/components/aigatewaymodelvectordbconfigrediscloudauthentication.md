# AIGatewayModelVectorDBConfigRedisCloudAuthentication

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthenticationOutput

```go
aiGatewayModelVectorDBConfigRedisCloudAuthentication := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationAws(components.AIGatewayRedisAWSAuthenticationOutput{/* values here */})
```

### AIGatewayRedisAzureAuthenticationOutput

```go
aiGatewayModelVectorDBConfigRedisCloudAuthentication := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationAzure(components.AIGatewayRedisAzureAuthenticationOutput{/* values here */})
```

### AIGatewayRedisGCPAuthenticationOutput

```go
aiGatewayModelVectorDBConfigRedisCloudAuthentication := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationGcp(components.AIGatewayRedisGCPAuthenticationOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelVectorDBConfigRedisCloudAuthentication.Type {
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationTypeAws:
		// aiGatewayModelVectorDBConfigRedisCloudAuthentication.AIGatewayRedisAWSAuthenticationOutput is populated
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationTypeAzure:
		// aiGatewayModelVectorDBConfigRedisCloudAuthentication.AIGatewayRedisAzureAuthenticationOutput is populated
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationTypeGcp:
		// aiGatewayModelVectorDBConfigRedisCloudAuthentication.AIGatewayRedisGCPAuthenticationOutput is populated
}
```
