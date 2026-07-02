# AIGatewayEmbeddingsModelConfig

Configuration for an embeddings model.


## Supported Types

### AIGatewayAzureEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigAzure(components.AIGatewayAzureEmbeddingsModelConfig{/* values here */})
```

### AIGatewayBedrockEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigBedrock(components.AIGatewayBedrockEmbeddingsModelConfig{/* values here */})
```

### AIGatewayDatabricksEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigDatabricks(components.AIGatewayDatabricksEmbeddingsModelConfig{/* values here */})
```

### AIGatewayGeminiEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigGemini(components.AIGatewayGeminiEmbeddingsModelConfig{/* values here */})
```

### AIGatewayHuggingfaceEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigHuggingface(components.AIGatewayHuggingfaceEmbeddingsModelConfig{/* values here */})
```

### AIGatewayVercelEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigVercel(components.AIGatewayVercelEmbeddingsModelConfig{/* values here */})
```

### AIGatewayVertexEmbeddingsModelConfig

```go
aiGatewayEmbeddingsModelConfig := components.CreateAIGatewayEmbeddingsModelConfigVertex(components.AIGatewayVertexEmbeddingsModelConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayEmbeddingsModelConfig.Type {
	case components.AIGatewayEmbeddingsModelConfigTypeAzure:
		// aiGatewayEmbeddingsModelConfig.AIGatewayAzureEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeBedrock:
		// aiGatewayEmbeddingsModelConfig.AIGatewayBedrockEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeDatabricks:
		// aiGatewayEmbeddingsModelConfig.AIGatewayDatabricksEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeGemini:
		// aiGatewayEmbeddingsModelConfig.AIGatewayGeminiEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeHuggingface:
		// aiGatewayEmbeddingsModelConfig.AIGatewayHuggingfaceEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeVercel:
		// aiGatewayEmbeddingsModelConfig.AIGatewayVercelEmbeddingsModelConfig is populated
	case components.AIGatewayEmbeddingsModelConfigTypeVertex:
		// aiGatewayEmbeddingsModelConfig.AIGatewayVertexEmbeddingsModelConfig is populated
}
```
