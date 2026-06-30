# APIImplementationListItem

An entity that implements an API


## Supported Types

### APIImplementationListItemGatewayServiceEntity

```go
apiImplementationListItem := components.CreateAPIImplementationListItemAPIImplementationListItemGatewayServiceEntity(components.APIImplementationListItemGatewayServiceEntity{/* values here */})
```

### APIImplementationListItemControlPlaneEntity

```go
apiImplementationListItem := components.CreateAPIImplementationListItemAPIImplementationListItemControlPlaneEntity(components.APIImplementationListItemControlPlaneEntity{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch apiImplementationListItem.Type {
	case components.APIImplementationListItemTypeAPIImplementationListItemGatewayServiceEntity:
		// apiImplementationListItem.APIImplementationListItemGatewayServiceEntity is populated
	case components.APIImplementationListItemTypeAPIImplementationListItemControlPlaneEntity:
		// apiImplementationListItem.APIImplementationListItemControlPlaneEntity is populated
}
```
