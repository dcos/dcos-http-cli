/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type IamUser struct {
	Uid          string `json:"uid"`
	Url          string `json:"url"`
	Description  string `json:"description"`
	IsRemote     bool   `json:"is_remote"`
	IsService    bool   `json:"is_service,omitempty"`
	PublicKey    string `json:"public_key,omitempty"`
	ProviderType string `json:"provider_type"`
	ProviderId   string `json:"provider_id"`
}
