// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ApplicationDeveloperDetailed - List of developers that have a registered application.
type ApplicationDeveloperDetailed struct {
	ID       string  `json:"id"`
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"full_name,omitempty"`
}

func (o *ApplicationDeveloperDetailed) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ApplicationDeveloperDetailed) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ApplicationDeveloperDetailed) GetFullName() *string {
	if o == nil {
		return nil
	}
	return o.FullName
}
