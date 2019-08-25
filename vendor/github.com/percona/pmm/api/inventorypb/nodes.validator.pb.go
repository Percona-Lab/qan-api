// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inventorypb/nodes.proto

package inventorypb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GenericNode) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ContainerNode) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *RemoteNode) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *RemoteAmazonRDSNode) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ListNodesRequest) Validate() error {
	return nil
}
func (this *ListNodesResponse) Validate() error {
	for _, item := range this.Generic {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Generic", err)
			}
		}
	}
	for _, item := range this.Container {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Container", err)
			}
		}
	}
	for _, item := range this.Remote {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Remote", err)
			}
		}
	}
	for _, item := range this.RemoteAmazonRds {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RemoteAmazonRds", err)
			}
		}
	}
	return nil
}
func (this *GetNodeRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	return nil
}
func (this *GetNodeResponse) Validate() error {
	if oneOfNester, ok := this.GetNode().(*GetNodeResponse_Generic); ok {
		if oneOfNester.Generic != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Generic); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Generic", err)
			}
		}
	}
	if oneOfNester, ok := this.GetNode().(*GetNodeResponse_Container); ok {
		if oneOfNester.Container != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Container); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Container", err)
			}
		}
	}
	if oneOfNester, ok := this.GetNode().(*GetNodeResponse_Remote); ok {
		if oneOfNester.Remote != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Remote); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Remote", err)
			}
		}
	}
	if oneOfNester, ok := this.GetNode().(*GetNodeResponse_RemoteAmazonRds); ok {
		if oneOfNester.RemoteAmazonRds != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RemoteAmazonRds); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RemoteAmazonRds", err)
			}
		}
	}
	return nil
}
func (this *AddGenericNodeRequest) Validate() error {
	if this.NodeName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeName", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeName))
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	return nil
}
func (this *AddGenericNodeResponse) Validate() error {
	if this.Generic != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Generic); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Generic", err)
		}
	}
	return nil
}
func (this *AddContainerNodeRequest) Validate() error {
	if this.NodeName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeName", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeName))
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	return nil
}
func (this *AddContainerNodeResponse) Validate() error {
	if this.Container != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Container); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Container", err)
		}
	}
	return nil
}
func (this *AddRemoteNodeRequest) Validate() error {
	if this.NodeName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeName", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeName))
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	return nil
}
func (this *AddRemoteNodeResponse) Validate() error {
	if this.Remote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Remote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Remote", err)
		}
	}
	return nil
}
func (this *AddRemoteAmazonRDSNodeRequest) Validate() error {
	if this.NodeName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeName", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeName))
	}
	if this.Instance == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Instance", fmt.Errorf(`value '%v' must not be an empty string`, this.Instance))
	}
	if this.Region == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Region", fmt.Errorf(`value '%v' must not be an empty string`, this.Region))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddRemoteAmazonRDSNodeResponse) Validate() error {
	if this.RemoteAmazonRds != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RemoteAmazonRds); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RemoteAmazonRds", err)
		}
	}
	return nil
}
func (this *RemoveNodeRequest) Validate() error {
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	return nil
}
func (this *RemoveNodeResponse) Validate() error {
	return nil
}
