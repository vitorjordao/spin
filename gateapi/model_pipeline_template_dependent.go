/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PipelineTemplateDependent struct {
	PipelineName string `json:"pipelineName,omitempty"`
	PipelineConfigId string `json:"pipelineConfigId,omitempty"`
	Application string `json:"application,omitempty"`
}
