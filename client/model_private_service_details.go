/*
Render Public API

Manage everything about your Render services

API version: 1.0.0
Contact: support@render.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package render

import (
	"encoding/json"
)

// PrivateServiceDetails struct for PrivateServiceDetails
type PrivateServiceDetails struct {
	Disk *StaticSiteDetailsParentServer `json:"disk,omitempty"`
	Env *ServiceEnv `json:"env,omitempty"`
	EnvSpecificDetails *WebServiceDetailsEnvSpecificDetails `json:"envSpecificDetails,omitempty"`
	NumInstances *int32 `json:"numInstances,omitempty"`
	OpenPorts []ServerPort `json:"openPorts,omitempty"`
	ParentServer *StaticSiteDetailsParentServer `json:"parentServer,omitempty"`
	Plan *string `json:"plan,omitempty"`
	PullRequestPreviewsEnabled *string `json:"pullRequestPreviewsEnabled,omitempty"`
	Region *Region `json:"region,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewPrivateServiceDetails instantiates a new PrivateServiceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateServiceDetails() *PrivateServiceDetails {
	this := PrivateServiceDetails{}
	return &this
}

// NewPrivateServiceDetailsWithDefaults instantiates a new PrivateServiceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateServiceDetailsWithDefaults() *PrivateServiceDetails {
	this := PrivateServiceDetails{}
	return &this
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetDisk() StaticSiteDetailsParentServer {
	if o == nil || isNil(o.Disk) {
		var ret StaticSiteDetailsParentServer
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetDiskOk() (*StaticSiteDetailsParentServer, bool) {
	if o == nil || isNil(o.Disk) {
    return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasDisk() bool {
	if o != nil && !isNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given StaticSiteDetailsParentServer and assigns it to the Disk field.
func (o *PrivateServiceDetails) SetDisk(v StaticSiteDetailsParentServer) {
	o.Disk = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetEnv() ServiceEnv {
	if o == nil || isNil(o.Env) {
		var ret ServiceEnv
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetEnvOk() (*ServiceEnv, bool) {
	if o == nil || isNil(o.Env) {
    return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasEnv() bool {
	if o != nil && !isNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given ServiceEnv and assigns it to the Env field.
func (o *PrivateServiceDetails) SetEnv(v ServiceEnv) {
	o.Env = &v
}

// GetEnvSpecificDetails returns the EnvSpecificDetails field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetEnvSpecificDetails() WebServiceDetailsEnvSpecificDetails {
	if o == nil || isNil(o.EnvSpecificDetails) {
		var ret WebServiceDetailsEnvSpecificDetails
		return ret
	}
	return *o.EnvSpecificDetails
}

// GetEnvSpecificDetailsOk returns a tuple with the EnvSpecificDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetEnvSpecificDetailsOk() (*WebServiceDetailsEnvSpecificDetails, bool) {
	if o == nil || isNil(o.EnvSpecificDetails) {
    return nil, false
	}
	return o.EnvSpecificDetails, true
}

// HasEnvSpecificDetails returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasEnvSpecificDetails() bool {
	if o != nil && !isNil(o.EnvSpecificDetails) {
		return true
	}

	return false
}

// SetEnvSpecificDetails gets a reference to the given WebServiceDetailsEnvSpecificDetails and assigns it to the EnvSpecificDetails field.
func (o *PrivateServiceDetails) SetEnvSpecificDetails(v WebServiceDetailsEnvSpecificDetails) {
	o.EnvSpecificDetails = &v
}

// GetNumInstances returns the NumInstances field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetNumInstances() int32 {
	if o == nil || isNil(o.NumInstances) {
		var ret int32
		return ret
	}
	return *o.NumInstances
}

// GetNumInstancesOk returns a tuple with the NumInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetNumInstancesOk() (*int32, bool) {
	if o == nil || isNil(o.NumInstances) {
    return nil, false
	}
	return o.NumInstances, true
}

// HasNumInstances returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasNumInstances() bool {
	if o != nil && !isNil(o.NumInstances) {
		return true
	}

	return false
}

// SetNumInstances gets a reference to the given int32 and assigns it to the NumInstances field.
func (o *PrivateServiceDetails) SetNumInstances(v int32) {
	o.NumInstances = &v
}

// GetOpenPorts returns the OpenPorts field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetOpenPorts() []ServerPort {
	if o == nil || isNil(o.OpenPorts) {
		var ret []ServerPort
		return ret
	}
	return o.OpenPorts
}

// GetOpenPortsOk returns a tuple with the OpenPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetOpenPortsOk() ([]ServerPort, bool) {
	if o == nil || isNil(o.OpenPorts) {
    return nil, false
	}
	return o.OpenPorts, true
}

// HasOpenPorts returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasOpenPorts() bool {
	if o != nil && !isNil(o.OpenPorts) {
		return true
	}

	return false
}

// SetOpenPorts gets a reference to the given []ServerPort and assigns it to the OpenPorts field.
func (o *PrivateServiceDetails) SetOpenPorts(v []ServerPort) {
	o.OpenPorts = v
}

// GetParentServer returns the ParentServer field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetParentServer() StaticSiteDetailsParentServer {
	if o == nil || isNil(o.ParentServer) {
		var ret StaticSiteDetailsParentServer
		return ret
	}
	return *o.ParentServer
}

// GetParentServerOk returns a tuple with the ParentServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetParentServerOk() (*StaticSiteDetailsParentServer, bool) {
	if o == nil || isNil(o.ParentServer) {
    return nil, false
	}
	return o.ParentServer, true
}

// HasParentServer returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasParentServer() bool {
	if o != nil && !isNil(o.ParentServer) {
		return true
	}

	return false
}

// SetParentServer gets a reference to the given StaticSiteDetailsParentServer and assigns it to the ParentServer field.
func (o *PrivateServiceDetails) SetParentServer(v StaticSiteDetailsParentServer) {
	o.ParentServer = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetPlan() string {
	if o == nil || isNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetPlanOk() (*string, bool) {
	if o == nil || isNil(o.Plan) {
    return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasPlan() bool {
	if o != nil && !isNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *PrivateServiceDetails) SetPlan(v string) {
	o.Plan = &v
}

// GetPullRequestPreviewsEnabled returns the PullRequestPreviewsEnabled field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetPullRequestPreviewsEnabled() string {
	if o == nil || isNil(o.PullRequestPreviewsEnabled) {
		var ret string
		return ret
	}
	return *o.PullRequestPreviewsEnabled
}

// GetPullRequestPreviewsEnabledOk returns a tuple with the PullRequestPreviewsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetPullRequestPreviewsEnabledOk() (*string, bool) {
	if o == nil || isNil(o.PullRequestPreviewsEnabled) {
    return nil, false
	}
	return o.PullRequestPreviewsEnabled, true
}

// HasPullRequestPreviewsEnabled returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasPullRequestPreviewsEnabled() bool {
	if o != nil && !isNil(o.PullRequestPreviewsEnabled) {
		return true
	}

	return false
}

// SetPullRequestPreviewsEnabled gets a reference to the given string and assigns it to the PullRequestPreviewsEnabled field.
func (o *PrivateServiceDetails) SetPullRequestPreviewsEnabled(v string) {
	o.PullRequestPreviewsEnabled = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetRegion() Region {
	if o == nil || isNil(o.Region) {
		var ret Region
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetRegionOk() (*Region, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given Region and assigns it to the Region field.
func (o *PrivateServiceDetails) SetRegion(v Region) {
	o.Region = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PrivateServiceDetails) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateServiceDetails) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PrivateServiceDetails) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PrivateServiceDetails) SetUrl(v string) {
	o.Url = &v
}

func (o PrivateServiceDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !isNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !isNil(o.EnvSpecificDetails) {
		toSerialize["envSpecificDetails"] = o.EnvSpecificDetails
	}
	if !isNil(o.NumInstances) {
		toSerialize["numInstances"] = o.NumInstances
	}
	if !isNil(o.OpenPorts) {
		toSerialize["openPorts"] = o.OpenPorts
	}
	if !isNil(o.ParentServer) {
		toSerialize["parentServer"] = o.ParentServer
	}
	if !isNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.PullRequestPreviewsEnabled) {
		toSerialize["pullRequestPreviewsEnabled"] = o.PullRequestPreviewsEnabled
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateServiceDetails struct {
	value *PrivateServiceDetails
	isSet bool
}

func (v NullablePrivateServiceDetails) Get() *PrivateServiceDetails {
	return v.value
}

func (v *NullablePrivateServiceDetails) Set(val *PrivateServiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateServiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateServiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateServiceDetails(val *PrivateServiceDetails) *NullablePrivateServiceDetails {
	return &NullablePrivateServiceDetails{value: val, isSet: true}
}

func (v NullablePrivateServiceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateServiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

