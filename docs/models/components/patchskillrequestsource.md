# PatchSkillRequestSource

Replaces the existing source in its entirety — partial updates to individual source attributes are not supported. The source `type` cannot be changed after creation.


## Supported Types

### RawSkillSourcePayload

```go
patchSkillRequestSource := components.CreatePatchSkillRequestSourceRaw(components.RawSkillSourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch patchSkillRequestSource.Type {
	case components.PatchSkillRequestSourceTypeRaw:
		// patchSkillRequestSource.RawSkillSourcePayload is populated
}
```
