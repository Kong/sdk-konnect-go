// This file contains manually maintained DeepCopy methods for types that have
// map[string]any fields, which controller-gen cannot process automatically.
// When new fields are added to RouteJSON or RouteExpression by Speakeasy, this
// file must be updated accordingly.

package components

// DeepCopyInto copies all properties from this object into another object of the
// same type that is provided as a pointer.
func (in *RouteJSON) DeepCopyInto(out *RouteJSON) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(int64)
		**out = **in
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]Destinations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HTTPSRedirectStatusCode != nil {
		in, out := &in.HTTPSRedirectStatusCode, &out.HTTPSRedirectStatusCode
		*out = new(HTTPSRedirectStatusCode)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PathHandling != nil {
		in, out := &in.PathHandling, &out.PathHandling
		*out = new(PathHandling)
		**out = **in
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreserveHost != nil {
		in, out := &in.PreserveHost, &out.PreserveHost
		*out = new(bool)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]RouteJSONProtocols, len(*in))
		copy(*out, *in)
	}
	if in.RegexPriority != nil {
		in, out := &in.RegexPriority, &out.RegexPriority
		*out = new(int64)
		**out = **in
	}
	if in.RequestBuffering != nil {
		in, out := &in.RequestBuffering, &out.RequestBuffering
		*out = new(bool)
		**out = **in
	}
	if in.ResponseBuffering != nil {
		in, out := &in.ResponseBuffering, &out.ResponseBuffering
		*out = new(bool)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RouteJSONService)
		(*in).DeepCopyInto(*out)
	}
	if in.Snis != nil {
		in, out := &in.Snis, &out.Snis
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]Sources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StripPath != nil {
		in, out := &in.StripPath, &out.StripPath
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy creates a deep copy of RouteJSON.
func (in *RouteJSON) DeepCopy() *RouteJSON {
	if in == nil {
		return nil
	}
	out := new(RouteJSON)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies all properties from this object into another object of the
// same type that is provided as a pointer.
func (in *RouteExpression) DeepCopyInto(out *RouteExpression) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(int64)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.HTTPSRedirectStatusCode != nil {
		in, out := &in.HTTPSRedirectStatusCode, &out.HTTPSRedirectStatusCode
		*out = new(RouteExpressionHTTPSRedirectStatusCode)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PathHandling != nil {
		in, out := &in.PathHandling, &out.PathHandling
		*out = new(RouteExpressionPathHandling)
		**out = **in
	}
	if in.PreserveHost != nil {
		in, out := &in.PreserveHost, &out.PreserveHost
		*out = new(bool)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]RouteExpressionProtocols, len(*in))
		copy(*out, *in)
	}
	if in.RequestBuffering != nil {
		in, out := &in.RequestBuffering, &out.RequestBuffering
		*out = new(bool)
		**out = **in
	}
	if in.ResponseBuffering != nil {
		in, out := &in.ResponseBuffering, &out.ResponseBuffering
		*out = new(bool)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RouteExpressionService)
		(*in).DeepCopyInto(*out)
	}
	if in.StripPath != nil {
		in, out := &in.StripPath, &out.StripPath
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy creates a deep copy of RouteExpression.
func (in *RouteExpression) DeepCopy() *RouteExpression {
	if in == nil {
		return nil
	}
	out := new(RouteExpression)
	in.DeepCopyInto(out)
	return out
}
