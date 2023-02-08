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
	"time"
)

// CustomDomain struct for CustomDomain
type CustomDomain struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DomainType *string `json:"domainType,omitempty"`
	PublicSuffix *string `json:"publicSuffix,omitempty"`
	RedirectForName *string `json:"redirectForName,omitempty"`
	VerificationStatus *string `json:"verificationStatus,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Server *StaticSiteDetailsParentServer `json:"server,omitempty"`
}

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomDomain) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomDomain) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomDomain) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomDomain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomDomain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomDomain) SetName(v string) {
	o.Name = &v
}

// GetDomainType returns the DomainType field value if set, zero value otherwise.
func (o *CustomDomain) GetDomainType() string {
	if o == nil || isNil(o.DomainType) {
		var ret string
		return ret
	}
	return *o.DomainType
}

// GetDomainTypeOk returns a tuple with the DomainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetDomainTypeOk() (*string, bool) {
	if o == nil || isNil(o.DomainType) {
    return nil, false
	}
	return o.DomainType, true
}

// HasDomainType returns a boolean if a field has been set.
func (o *CustomDomain) HasDomainType() bool {
	if o != nil && !isNil(o.DomainType) {
		return true
	}

	return false
}

// SetDomainType gets a reference to the given string and assigns it to the DomainType field.
func (o *CustomDomain) SetDomainType(v string) {
	o.DomainType = &v
}

// GetPublicSuffix returns the PublicSuffix field value if set, zero value otherwise.
func (o *CustomDomain) GetPublicSuffix() string {
	if o == nil || isNil(o.PublicSuffix) {
		var ret string
		return ret
	}
	return *o.PublicSuffix
}

// GetPublicSuffixOk returns a tuple with the PublicSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetPublicSuffixOk() (*string, bool) {
	if o == nil || isNil(o.PublicSuffix) {
    return nil, false
	}
	return o.PublicSuffix, true
}

// HasPublicSuffix returns a boolean if a field has been set.
func (o *CustomDomain) HasPublicSuffix() bool {
	if o != nil && !isNil(o.PublicSuffix) {
		return true
	}

	return false
}

// SetPublicSuffix gets a reference to the given string and assigns it to the PublicSuffix field.
func (o *CustomDomain) SetPublicSuffix(v string) {
	o.PublicSuffix = &v
}

// GetRedirectForName returns the RedirectForName field value if set, zero value otherwise.
func (o *CustomDomain) GetRedirectForName() string {
	if o == nil || isNil(o.RedirectForName) {
		var ret string
		return ret
	}
	return *o.RedirectForName
}

// GetRedirectForNameOk returns a tuple with the RedirectForName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetRedirectForNameOk() (*string, bool) {
	if o == nil || isNil(o.RedirectForName) {
    return nil, false
	}
	return o.RedirectForName, true
}

// HasRedirectForName returns a boolean if a field has been set.
func (o *CustomDomain) HasRedirectForName() bool {
	if o != nil && !isNil(o.RedirectForName) {
		return true
	}

	return false
}

// SetRedirectForName gets a reference to the given string and assigns it to the RedirectForName field.
func (o *CustomDomain) SetRedirectForName(v string) {
	o.RedirectForName = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *CustomDomain) GetVerificationStatus() string {
	if o == nil || isNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetVerificationStatusOk() (*string, bool) {
	if o == nil || isNil(o.VerificationStatus) {
    return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *CustomDomain) HasVerificationStatus() bool {
	if o != nil && !isNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *CustomDomain) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomDomain) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomDomain) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomDomain) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *CustomDomain) GetServer() StaticSiteDetailsParentServer {
	if o == nil || isNil(o.Server) {
		var ret StaticSiteDetailsParentServer
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetServerOk() (*StaticSiteDetailsParentServer, bool) {
	if o == nil || isNil(o.Server) {
    return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *CustomDomain) HasServer() bool {
	if o != nil && !isNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given StaticSiteDetailsParentServer and assigns it to the Server field.
func (o *CustomDomain) SetServer(v StaticSiteDetailsParentServer) {
	o.Server = &v
}

func (o CustomDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DomainType) {
		toSerialize["domainType"] = o.DomainType
	}
	if !isNil(o.PublicSuffix) {
		toSerialize["publicSuffix"] = o.PublicSuffix
	}
	if !isNil(o.RedirectForName) {
		toSerialize["redirectForName"] = o.RedirectForName
	}
	if !isNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

