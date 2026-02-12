# VirtualClusterAuthenticationOauthBearerMediation

Methods to mediate authentication:
* passthrough - pass authentication from the client through proxy to the backend cluster without any kind of


  validation
* validate_forward - pass authentication from the client through proxy to the backend cluster.


  Proxy does the validation before forwarding it to the client.
* terminate - terminate authentication at the proxy level and originate authentication to the backend cluster


  using the configuration defined at BackendCluster's authentication.
  SASL auth is not originated if authentication on the backend_cluster is not configured.



## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `VirtualClusterAuthenticationOauthBearerMediationPassthrough`     | passthrough                                                       |
| `VirtualClusterAuthenticationOauthBearerMediationValidateForward` | validate_forward                                                  |
| `VirtualClusterAuthenticationOauthBearerMediationTerminate`       | terminate                                                         |