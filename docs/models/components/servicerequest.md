# ServiceRequest

Service entities, as the name implies, are abstractions of each of your own upstream services. Examples of Services would be a data transformation microservice, a billing API, etc. The main attribute of a Service is its URL (where Kong should proxy traffic to), which can be set as a single string or by specifying its `protocol`, `host`, `port` and `path` individually. Services are associated to Routes (a Service can have many Routes associated with it). Routes are entry-points in Kong and define rules to match client requests. Once a Route is matched, Kong proxies the request to its associated Service. See the [Proxy Reference][proxy-reference] for a detailed explanation of how Kong proxies traffic.


## Supported Types

### ServiceRequest1

```go
serviceRequest := components.CreateServiceRequestServiceRequest1(components.ServiceRequest1{/* values here */})
```

### ServiceRequest2

```go
serviceRequest := components.CreateServiceRequestServiceRequest2(components.ServiceRequest2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch serviceRequest.Type {
	case components.ServiceRequestTypeServiceRequest1:
		// serviceRequest.ServiceRequest1 is populated
	case components.ServiceRequestTypeServiceRequest2:
		// serviceRequest.ServiceRequest2 is populated
}
```
