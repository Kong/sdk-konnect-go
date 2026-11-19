# AICostExport

An export copies the organization's stored usage events in a time range into files to download. It runs asynchronously: `status` follows its progress, and once it succeeds, `files` lists the data files, which can be downloaded until `expires_at`.

After creation, only an export's status and results change. A client creates a one-time export, and an export schedule creates one export of each of its periods.


## Supported Types

### AICostExport1

```go
aiCostExport := components.CreateAICostExportAICostExport1(components.AICostExport1{/* values here */})
```

### AICostExport2

```go
aiCostExport := components.CreateAICostExportAICostExport2(components.AICostExport2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiCostExport.Type {
	case components.AICostExportTypeAICostExport1:
		// aiCostExport.AICostExport1 is populated
	case components.AICostExportTypeAICostExport2:
		// aiCostExport.AICostExport2 is populated
}
```
