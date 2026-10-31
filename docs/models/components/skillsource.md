# SkillSource

Where the skill's content comes from, as returned by the API. Never includes the content itself — read it through the skill contents endpoint.


## Supported Types

### RawSkillSource

```go
skillSource := components.CreateSkillSourceRaw(components.RawSkillSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch skillSource.Type {
	case components.SkillSourceTypeRaw:
		// skillSource.RawSkillSource is populated
}
```
