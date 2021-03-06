/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type VirtualPMemController struct {

	Devices map[string]VirtualPMemDevice `json:"Devices,omitempty"`

	MaximumCount int32 `json:"MaximumCount,omitempty"`

	MaximumSizeBytes int32 `json:"MaximumSizeBytes,omitempty"`

	Backing string `json:"Backing,omitempty"`
}
