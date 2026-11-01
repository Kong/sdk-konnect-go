# InvalidParameters


## Supported Types

### One

```go
invalidParameters := sdkerrors.CreateInvalidParametersOne(sdkerrors.One{/* values here */})
```

### Two

```go
invalidParameters := sdkerrors.CreateInvalidParametersTwo(sdkerrors.Two{/* values here */})
```

### Three

```go
invalidParameters := sdkerrors.CreateInvalidParametersThree(sdkerrors.Three{/* values here */})
```

### Four

```go
invalidParameters := sdkerrors.CreateInvalidParametersFour(sdkerrors.Four{/* values here */})
```

### Five

```go
invalidParameters := sdkerrors.CreateInvalidParametersFive(sdkerrors.Five{/* values here */})
```

### Six

```go
invalidParameters := sdkerrors.CreateInvalidParametersSix(sdkerrors.Six{/* values here */})
```

### Seven

```go
invalidParameters := sdkerrors.CreateInvalidParametersSeven(sdkerrors.Seven{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch invalidParameters.Type {
	case sdkerrors.InvalidParametersTypeOne:
		// invalidParameters.One is populated
	case sdkerrors.InvalidParametersTypeTwo:
		// invalidParameters.Two is populated
	case sdkerrors.InvalidParametersTypeThree:
		// invalidParameters.Three is populated
	case sdkerrors.InvalidParametersTypeFour:
		// invalidParameters.Four is populated
	case sdkerrors.InvalidParametersTypeFive:
		// invalidParameters.Five is populated
	case sdkerrors.InvalidParametersTypeSix:
		// invalidParameters.Six is populated
	case sdkerrors.InvalidParametersTypeSeven:
		// invalidParameters.Seven is populated
}
```
