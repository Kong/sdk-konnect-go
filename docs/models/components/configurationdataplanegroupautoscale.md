# ConfigurationDataPlaneGroupAutoscale


## Supported Types

### ConfigurationDataPlaneGroupAutoscaleStatic

```go
configurationDataPlaneGroupAutoscale := components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleStatic(components.ConfigurationDataPlaneGroupAutoscaleStatic{/* values here */})
```

### ConfigurationDataPlaneGroupAutoscaleAutopilot

```go
configurationDataPlaneGroupAutoscale := components.CreateConfigurationDataPlaneGroupAutoscaleConfigurationDataPlaneGroupAutoscaleAutopilot(components.ConfigurationDataPlaneGroupAutoscaleAutopilot{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch configurationDataPlaneGroupAutoscale.Type {
	case components.ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleStatic:
		// configurationDataPlaneGroupAutoscale.ConfigurationDataPlaneGroupAutoscaleStatic is populated
	case components.ConfigurationDataPlaneGroupAutoscaleTypeConfigurationDataPlaneGroupAutoscaleAutopilot:
		// configurationDataPlaneGroupAutoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot is populated
}
```
