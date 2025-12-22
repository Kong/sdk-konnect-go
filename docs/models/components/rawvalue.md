# RawValue

Raw value of the criteria evaluation.

For example, a service may be failing the `time_to_acknowledge` criteria evaluation when
the the mean time-to-ack is greater than 15 minutes. This value represents the _actual_
mean time-to-ack value for the service which if not passing could be 25 minutes.



## Supported Types

### 

```go
rawValue := components.CreateRawValueNumber(float64{/* values here */})
```

### 

```go
rawValue := components.CreateRawValueStr(string{/* values here */})
```

### 

```go
rawValue := components.CreateRawValueBoolean(bool{/* values here */})
```

### TimeValue

```go
rawValue := components.CreateRawValueTimeValue(components.TimeValue{/* values here */})
```

### CriteriaEvaluationRelationMap

```go
rawValue := components.CreateRawValueCriteriaEvaluationRelationMap(components.CriteriaEvaluationRelationMap{/* values here */})
```

### 

```go
rawValue := components.CreateRawValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
rawValue := components.CreateRawValueAny(any{/* values here */})
```

