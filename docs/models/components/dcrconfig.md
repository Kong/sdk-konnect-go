# DcrConfig


## Supported Types

### UpdateDcrConfigAuth0InRequest

```go
dcrConfig := components.CreateDcrConfigUpdateDcrConfigAuth0InRequest(components.UpdateDcrConfigAuth0InRequest{/* values here */})
```

### UpdateDcrConfigAzureAdInRequest

```go
dcrConfig := components.CreateDcrConfigUpdateDcrConfigAzureAdInRequest(components.UpdateDcrConfigAzureAdInRequest{/* values here */})
```

### UpdateDcrConfigCurityInRequest

```go
dcrConfig := components.CreateDcrConfigUpdateDcrConfigCurityInRequest(components.UpdateDcrConfigCurityInRequest{/* values here */})
```

### UpdateDcrConfigOktaInRequest

```go
dcrConfig := components.CreateDcrConfigUpdateDcrConfigOktaInRequest(components.UpdateDcrConfigOktaInRequest{/* values here */})
```

### UpdateDcrConfigHTTPInRequest

```go
dcrConfig := components.CreateDcrConfigUpdateDcrConfigHTTPInRequest(components.UpdateDcrConfigHTTPInRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dcrConfig.Type {
	case components.DcrConfigTypeUpdateDcrConfigAuth0InRequest:
		// dcrConfig.UpdateDcrConfigAuth0InRequest is populated
	case components.DcrConfigTypeUpdateDcrConfigAzureAdInRequest:
		// dcrConfig.UpdateDcrConfigAzureAdInRequest is populated
	case components.DcrConfigTypeUpdateDcrConfigCurityInRequest:
		// dcrConfig.UpdateDcrConfigCurityInRequest is populated
	case components.DcrConfigTypeUpdateDcrConfigOktaInRequest:
		// dcrConfig.UpdateDcrConfigOktaInRequest is populated
	case components.DcrConfigTypeUpdateDcrConfigHTTPInRequest:
		// dcrConfig.UpdateDcrConfigHTTPInRequest is populated
}
```
