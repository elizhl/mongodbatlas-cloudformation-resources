package resource

import "github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn/encoding"

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'generate' command.

Updates to this type are made my editing the schema file
and executing the 'generate' command
*/

// Model is autogenerated from the json schema
type Model struct {
	Name         *encoding.String `json:"Name,omitempty"`
	OrgId        *encoding.String `json:"OrgId,omitempty"`
	Id           *encoding.String `json:"Id,omitempty"`
	Created      *encoding.String `json:"Created,omitempty"`
	ClusterCount *encoding.Int    `json:"ClusterCount,omitempty"`
}
