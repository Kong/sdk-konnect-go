# BodySanitizationStrategy

Strategy for body content sanitization.


## Supported Types

### JSONPathStrategy

```go
bodySanitizationStrategy := components.CreateBodySanitizationStrategyJSONPathStrategy(components.JSONPathStrategy{/* values here */})
```

### RegexStrategy

```go
bodySanitizationStrategy := components.CreateBodySanitizationStrategyRegexStrategy(components.RegexStrategy{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bodySanitizationStrategy.Type {
	case components.BodySanitizationStrategyTypeJSONPathStrategy:
		// bodySanitizationStrategy.JSONPathStrategy is populated
	case components.BodySanitizationStrategyTypeRegexStrategy:
		// bodySanitizationStrategy.RegexStrategy is populated
}
```
