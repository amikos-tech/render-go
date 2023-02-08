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
	"fmt"
)

// ServicePOSTServiceDetails - struct for ServicePOSTServiceDetails
type ServicePOSTServiceDetails struct {
	BackgroundWorkerDetailsPOST *BackgroundWorkerDetailsPOST
	CronJobDetailsPOST *CronJobDetailsPOST
	PrivateServiceDetailsPOST *PrivateServiceDetailsPOST
	StaticSiteDetailsPOST *StaticSiteDetailsPOST
	WebServiceDetailsPOST *WebServiceDetailsPOST
}

// BackgroundWorkerDetailsPOSTAsServicePOSTServiceDetails is a convenience function that returns BackgroundWorkerDetailsPOST wrapped in ServicePOSTServiceDetails
func BackgroundWorkerDetailsPOSTAsServicePOSTServiceDetails(v *BackgroundWorkerDetailsPOST) ServicePOSTServiceDetails {
	return ServicePOSTServiceDetails{
		BackgroundWorkerDetailsPOST: v,
	}
}

// CronJobDetailsPOSTAsServicePOSTServiceDetails is a convenience function that returns CronJobDetailsPOST wrapped in ServicePOSTServiceDetails
func CronJobDetailsPOSTAsServicePOSTServiceDetails(v *CronJobDetailsPOST) ServicePOSTServiceDetails {
	return ServicePOSTServiceDetails{
		CronJobDetailsPOST: v,
	}
}

// PrivateServiceDetailsPOSTAsServicePOSTServiceDetails is a convenience function that returns PrivateServiceDetailsPOST wrapped in ServicePOSTServiceDetails
func PrivateServiceDetailsPOSTAsServicePOSTServiceDetails(v *PrivateServiceDetailsPOST) ServicePOSTServiceDetails {
	return ServicePOSTServiceDetails{
		PrivateServiceDetailsPOST: v,
	}
}

// StaticSiteDetailsPOSTAsServicePOSTServiceDetails is a convenience function that returns StaticSiteDetailsPOST wrapped in ServicePOSTServiceDetails
func StaticSiteDetailsPOSTAsServicePOSTServiceDetails(v *StaticSiteDetailsPOST) ServicePOSTServiceDetails {
	return ServicePOSTServiceDetails{
		StaticSiteDetailsPOST: v,
	}
}

// WebServiceDetailsPOSTAsServicePOSTServiceDetails is a convenience function that returns WebServiceDetailsPOST wrapped in ServicePOSTServiceDetails
func WebServiceDetailsPOSTAsServicePOSTServiceDetails(v *WebServiceDetailsPOST) ServicePOSTServiceDetails {
	return ServicePOSTServiceDetails{
		WebServiceDetailsPOST: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServicePOSTServiceDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackgroundWorkerDetailsPOST
	err = newStrictDecoder(data).Decode(&dst.BackgroundWorkerDetailsPOST)
	if err == nil {
		jsonBackgroundWorkerDetailsPOST, _ := json.Marshal(dst.BackgroundWorkerDetailsPOST)
		if string(jsonBackgroundWorkerDetailsPOST) == "{}" { // empty struct
			dst.BackgroundWorkerDetailsPOST = nil
		} else {
			match++
		}
	} else {
		dst.BackgroundWorkerDetailsPOST = nil
	}

	// try to unmarshal data into CronJobDetailsPOST
	err = newStrictDecoder(data).Decode(&dst.CronJobDetailsPOST)
	if err == nil {
		jsonCronJobDetailsPOST, _ := json.Marshal(dst.CronJobDetailsPOST)
		if string(jsonCronJobDetailsPOST) == "{}" { // empty struct
			dst.CronJobDetailsPOST = nil
		} else {
			match++
		}
	} else {
		dst.CronJobDetailsPOST = nil
	}

	// try to unmarshal data into PrivateServiceDetailsPOST
	err = newStrictDecoder(data).Decode(&dst.PrivateServiceDetailsPOST)
	if err == nil {
		jsonPrivateServiceDetailsPOST, _ := json.Marshal(dst.PrivateServiceDetailsPOST)
		if string(jsonPrivateServiceDetailsPOST) == "{}" { // empty struct
			dst.PrivateServiceDetailsPOST = nil
		} else {
			match++
		}
	} else {
		dst.PrivateServiceDetailsPOST = nil
	}

	// try to unmarshal data into StaticSiteDetailsPOST
	err = newStrictDecoder(data).Decode(&dst.StaticSiteDetailsPOST)
	if err == nil {
		jsonStaticSiteDetailsPOST, _ := json.Marshal(dst.StaticSiteDetailsPOST)
		if string(jsonStaticSiteDetailsPOST) == "{}" { // empty struct
			dst.StaticSiteDetailsPOST = nil
		} else {
			match++
		}
	} else {
		dst.StaticSiteDetailsPOST = nil
	}

	// try to unmarshal data into WebServiceDetailsPOST
	err = newStrictDecoder(data).Decode(&dst.WebServiceDetailsPOST)
	if err == nil {
		jsonWebServiceDetailsPOST, _ := json.Marshal(dst.WebServiceDetailsPOST)
		if string(jsonWebServiceDetailsPOST) == "{}" { // empty struct
			dst.WebServiceDetailsPOST = nil
		} else {
			match++
		}
	} else {
		dst.WebServiceDetailsPOST = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackgroundWorkerDetailsPOST = nil
		dst.CronJobDetailsPOST = nil
		dst.PrivateServiceDetailsPOST = nil
		dst.StaticSiteDetailsPOST = nil
		dst.WebServiceDetailsPOST = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServicePOSTServiceDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServicePOSTServiceDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServicePOSTServiceDetails) MarshalJSON() ([]byte, error) {
	if src.BackgroundWorkerDetailsPOST != nil {
		return json.Marshal(&src.BackgroundWorkerDetailsPOST)
	}

	if src.CronJobDetailsPOST != nil {
		return json.Marshal(&src.CronJobDetailsPOST)
	}

	if src.PrivateServiceDetailsPOST != nil {
		return json.Marshal(&src.PrivateServiceDetailsPOST)
	}

	if src.StaticSiteDetailsPOST != nil {
		return json.Marshal(&src.StaticSiteDetailsPOST)
	}

	if src.WebServiceDetailsPOST != nil {
		return json.Marshal(&src.WebServiceDetailsPOST)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServicePOSTServiceDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BackgroundWorkerDetailsPOST != nil {
		return obj.BackgroundWorkerDetailsPOST
	}

	if obj.CronJobDetailsPOST != nil {
		return obj.CronJobDetailsPOST
	}

	if obj.PrivateServiceDetailsPOST != nil {
		return obj.PrivateServiceDetailsPOST
	}

	if obj.StaticSiteDetailsPOST != nil {
		return obj.StaticSiteDetailsPOST
	}

	if obj.WebServiceDetailsPOST != nil {
		return obj.WebServiceDetailsPOST
	}

	// all schemas are nil
	return nil
}

type NullableServicePOSTServiceDetails struct {
	value *ServicePOSTServiceDetails
	isSet bool
}

func (v NullableServicePOSTServiceDetails) Get() *ServicePOSTServiceDetails {
	return v.value
}

func (v *NullableServicePOSTServiceDetails) Set(val *ServicePOSTServiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePOSTServiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePOSTServiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePOSTServiceDetails(val *ServicePOSTServiceDetails) *NullableServicePOSTServiceDetails {
	return &NullableServicePOSTServiceDetails{value: val, isSet: true}
}

func (v NullableServicePOSTServiceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePOSTServiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

