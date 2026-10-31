# AIGatewayModelVectorDBConfigRedisCloudAuthentication

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthentication

```go
aiGatewayModelVectorDBConfigRedisCloudAuthentication := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationAws(components.AIGatewayRedisAWSAuthentication{/* values here */})
```

### AIGatewayRedisAzureAuthentication

```go
aiGatewayModelVectorDBConfigRedisCloudAuthentication := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationAzure(components.AIGatewayRedisAzureAuthentication{/* values here */})
```

### AIGatewayRedisGCPAuthentication

```go
aiGatewayModelVectorDBConfigRedisCloudAuthentication := components.CreateAIGatewayModelVectorDBConfigRedisCloudAuthenticationGcp(components.AIGatewayRedisGCPAuthentication{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayModelVectorDBConfigRedisCloudAuthentication.Type {
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationTypeAws:
		// aiGatewayModelVectorDBConfigRedisCloudAuthentication.AIGatewayRedisAWSAuthentication is populated
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationTypeAzure:
		// aiGatewayModelVectorDBConfigRedisCloudAuthentication.AIGatewayRedisAzureAuthentication is populated
	case components.AIGatewayModelVectorDBConfigRedisCloudAuthenticationTypeGcp:
		// aiGatewayModelVectorDBConfigRedisCloudAuthentication.AIGatewayRedisGCPAuthentication is populated
}
```
