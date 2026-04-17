# EventGatewayKeySource

A key source that describes how to find a symmetric key for encryption or decryption.
It can be an AWS KMS key source that uses a KMS to find a symmetric key,
or a static key source that uses a static symmetric key provided as secrets.



## Supported Types

### EventGatewayAWSKeySource

```go
eventGatewayKeySource := components.CreateEventGatewayKeySourceAws(components.EventGatewayAWSKeySource{/* values here */})
```

### EventGatewayStaticKeySource

```go
eventGatewayKeySource := components.CreateEventGatewayKeySourceStatic(components.EventGatewayStaticKeySource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayKeySource.Type {
	case components.EventGatewayKeySourceTypeAws:
		// eventGatewayKeySource.EventGatewayAWSKeySource is populated
	case components.EventGatewayKeySourceTypeStatic:
		// eventGatewayKeySource.EventGatewayStaticKeySource is populated
}
```
