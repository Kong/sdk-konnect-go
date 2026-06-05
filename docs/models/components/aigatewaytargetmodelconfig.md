# AIGatewayTargetModelConfig

Configuration for a target model.


## Supported Types

### AIGatewayTargetModelAnthropicConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigAnthropic(components.AIGatewayTargetModelAnthropicConfig{/* values here */})
```

### AIGatewayTargetModelAzureConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigAzure(components.AIGatewayTargetModelAzureConfig{/* values here */})
```

### AIGatewayTargetModelBedrockConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigBedrock(components.AIGatewayTargetModelBedrockConfig{/* values here */})
```

### AIGatewayTargetModelCerebrasConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigCerebras(components.AIGatewayTargetModelCerebrasConfig{/* values here */})
```

### AIGatewayTargetModelCohereConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigCohere(components.AIGatewayTargetModelCohereConfig{/* values here */})
```

### AIGatewayTargetModelDashscopeConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigDashscope(components.AIGatewayTargetModelDashscopeConfig{/* values here */})
```

### AIGatewayTargetModelDatabricksConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigDatabricks(components.AIGatewayTargetModelDatabricksConfig{/* values here */})
```

### AIGatewayTargetModelDeepseekConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigDeepseek(components.AIGatewayTargetModelDeepseekConfig{/* values here */})
```

### AIGatewayTargetModelGeminiConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigGemini(components.AIGatewayTargetModelGeminiConfig{/* values here */})
```

### AIGatewayTargetModelHuggingfaceConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigHuggingface(components.AIGatewayTargetModelHuggingfaceConfig{/* values here */})
```

### AIGatewayTargetModelKimiConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigKimi(components.AIGatewayTargetModelKimiConfig{/* values here */})
```

### AIGatewayTargetModelLlama2Config

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigLlama2(components.AIGatewayTargetModelLlama2Config{/* values here */})
```

### AIGatewayTargetModelMistralConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigMistral(components.AIGatewayTargetModelMistralConfig{/* values here */})
```

### AIGatewayTargetModelOllamaConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigOllama(components.AIGatewayTargetModelOllamaConfig{/* values here */})
```

### AIGatewayTargetModelOpenaiConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigOpenai(components.AIGatewayTargetModelOpenaiConfig{/* values here */})
```

### AIGatewayTargetModelVercelConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigVercel(components.AIGatewayTargetModelVercelConfig{/* values here */})
```

### AIGatewayTargetModelVllmConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigVllm(components.AIGatewayTargetModelVllmConfig{/* values here */})
```

### AIGatewayTargetModelXaiConfig

```go
aiGatewayTargetModelConfig := components.CreateAIGatewayTargetModelConfigXai(components.AIGatewayTargetModelXaiConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayTargetModelConfig.Type {
	case components.AIGatewayTargetModelConfigTypeAnthropic:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelAnthropicConfig is populated
	case components.AIGatewayTargetModelConfigTypeAzure:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelAzureConfig is populated
	case components.AIGatewayTargetModelConfigTypeBedrock:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelBedrockConfig is populated
	case components.AIGatewayTargetModelConfigTypeCerebras:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelCerebrasConfig is populated
	case components.AIGatewayTargetModelConfigTypeCohere:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelCohereConfig is populated
	case components.AIGatewayTargetModelConfigTypeDashscope:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelDashscopeConfig is populated
	case components.AIGatewayTargetModelConfigTypeDatabricks:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelDatabricksConfig is populated
	case components.AIGatewayTargetModelConfigTypeDeepseek:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelDeepseekConfig is populated
	case components.AIGatewayTargetModelConfigTypeGemini:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelGeminiConfig is populated
	case components.AIGatewayTargetModelConfigTypeHuggingface:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelHuggingfaceConfig is populated
	case components.AIGatewayTargetModelConfigTypeKimi:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelKimiConfig is populated
	case components.AIGatewayTargetModelConfigTypeLlama2:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelLlama2Config is populated
	case components.AIGatewayTargetModelConfigTypeMistral:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelMistralConfig is populated
	case components.AIGatewayTargetModelConfigTypeOllama:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelOllamaConfig is populated
	case components.AIGatewayTargetModelConfigTypeOpenai:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelOpenaiConfig is populated
	case components.AIGatewayTargetModelConfigTypeVercel:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelVercelConfig is populated
	case components.AIGatewayTargetModelConfigTypeVllm:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelVllmConfig is populated
	case components.AIGatewayTargetModelConfigTypeXai:
		// aiGatewayTargetModelConfig.AIGatewayTargetModelXaiConfig is populated
}
```
