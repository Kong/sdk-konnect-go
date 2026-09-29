# SkillSourcePayload

Where the skill's content comes from, used in create and update requests. The shape of `config` is determined by `type`.


## Supported Types

### RawSkillSourcePayload

```go
skillSourcePayload := components.CreateSkillSourcePayloadRaw(components.RawSkillSourcePayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch skillSourcePayload.Type {
	case components.SkillSourcePayloadTypeRaw:
		// skillSourcePayload.RawSkillSourcePayload is populated
}
```
