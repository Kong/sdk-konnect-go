# CloudAuthentication

Auth related config for connecting to a Cloud Provider's Redis instance.


## Supported Types

### AIGatewayRedisAWSAuthentication

```go
cloudAuthentication := components.CreateCloudAuthenticationAws(components.AIGatewayRedisAWSAuthentication{/* values here */})
```

### AIGatewayRedisAzureAuthentication

```go
cloudAuthentication := components.CreateCloudAuthenticationAzure(components.AIGatewayRedisAzureAuthentication{/* values here */})
```

### AIGatewayRedisGCPAuthentication

```go
cloudAuthentication := components.CreateCloudAuthenticationGcp(components.AIGatewayRedisGCPAuthentication{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch cloudAuthentication.Type {
	case components.CloudAuthenticationTypeAws:
		// cloudAuthentication.AIGatewayRedisAWSAuthentication is populated
	case components.CloudAuthenticationTypeAzure:
		// cloudAuthentication.AIGatewayRedisAzureAuthentication is populated
	case components.CloudAuthenticationTypeGcp:
		// cloudAuthentication.AIGatewayRedisGCPAuthentication is populated
}
```
