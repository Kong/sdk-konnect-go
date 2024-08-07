//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package components

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateControlPlaneRequest) DeepCopyInto(out *CreateControlPlaneRequest) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ClusterType != nil {
		in, out := &in.ClusterType, &out.ClusterType
		*out = new(ClusterType)
		**out = **in
	}
	if in.AuthType != nil {
		in, out := &in.AuthType, &out.AuthType
		*out = new(AuthType)
		**out = **in
	}
	if in.CloudGateway != nil {
		in, out := &in.CloudGateway, &out.CloudGateway
		*out = new(bool)
		**out = **in
	}
	if in.ProxyUrls != nil {
		in, out := &in.ProxyUrls, &out.ProxyUrls
		*out = make([]ProxyURL, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateControlPlaneRequest.
func (in *CreateControlPlaneRequest) DeepCopy() *CreateControlPlaneRequest {
	if in == nil {
		return nil
	}
	out := new(CreateControlPlaneRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destinations) DeepCopyInto(out *Destinations) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destinations.
func (in *Destinations) DeepCopy() *Destinations {
	if in == nil {
		return nil
	}
	out := new(Destinations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
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
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
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
		*out = make([]RouteProtocols, len(*in))
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
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RouteService)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteInput) DeepCopyInto(out *RouteInput) {
	*out = *in
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]Destinations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
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
		*out = make([]RouteProtocols, len(*in))
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
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(RouteService)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteInput.
func (in *RouteInput) DeepCopy() *RouteInput {
	if in == nil {
		return nil
	}
	out := new(RouteInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteService) DeepCopyInto(out *RouteService) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteService.
func (in *RouteService) DeepCopy() *RouteService {
	if in == nil {
		return nil
	}
	out := new(RouteService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteWithoutParents) DeepCopyInto(out *RouteWithoutParents) {
	*out = *in
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]RouteWithoutParentsDestinations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HTTPSRedirectStatusCode != nil {
		in, out := &in.HTTPSRedirectStatusCode, &out.HTTPSRedirectStatusCode
		*out = new(RouteWithoutParentsHTTPSRedirectStatusCode)
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
		*out = new(RouteWithoutParentsPathHandling)
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
		*out = make([]RouteWithoutParentsProtocols, len(*in))
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
	if in.Snis != nil {
		in, out := &in.Snis, &out.Snis
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]RouteWithoutParentsSources, len(*in))
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteWithoutParents.
func (in *RouteWithoutParents) DeepCopy() *RouteWithoutParents {
	if in == nil {
		return nil
	}
	out := new(RouteWithoutParents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteWithoutParentsDestinations) DeepCopyInto(out *RouteWithoutParentsDestinations) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteWithoutParentsDestinations.
func (in *RouteWithoutParentsDestinations) DeepCopy() *RouteWithoutParentsDestinations {
	if in == nil {
		return nil
	}
	out := new(RouteWithoutParentsDestinations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteWithoutParentsSources) DeepCopyInto(out *RouteWithoutParentsSources) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteWithoutParentsSources.
func (in *RouteWithoutParentsSources) DeepCopy() *RouteWithoutParentsSources {
	if in == nil {
		return nil
	}
	out := new(RouteWithoutParentsSources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sources) DeepCopyInto(out *Sources) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sources.
func (in *Sources) DeepCopy() *Sources {
	if in == nil {
		return nil
	}
	out := new(Sources)
	in.DeepCopyInto(out)
	return out
}
