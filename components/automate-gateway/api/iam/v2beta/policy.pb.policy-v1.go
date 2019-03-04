// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/policy.proto

package v2beta

import (
	policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/CreatePolicy", "auth:policies", "create", "POST", "/iam/v2beta/policies", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreatePolicyReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/GetPolicy", "auth:policies:{id}", "get", "GET", "/iam/v2beta/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetPolicyReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/ListPolicies", "auth:policies", "read", "GET", "/iam/v2beta/policies", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/DeletePolicy", "auth:policies:{id}", "delete", "DELETE", "/iam/v2beta/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeletePolicyReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/UpdatePolicy", "auth:policies:{id}", "update", "PUT", "/iam/v2beta/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdatePolicyReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/GetPolicyVersion", "auth:policies", "read", "GET", "/iam/v2beta/policy_version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/ListPolicyMembers", "auth:policies:{id}", "read", "GET", "/iam/v2beta/policies/{id}/members", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ListPolicyMembersReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/ReplacePolicyMembers", "auth:policies:{id}", "update", "PUT", "/iam/v2beta/policies/{id}/members", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ReplacePolicyMembersReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/RemovePolicyMembers", "auth:policies:{id}", "delete", "POST", "/iam/v2beta/policies/{id}/members:remove", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RemovePolicyMembersReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/AddPolicyMembers", "auth:policies:{id}", "create", "POST", "/iam/v2beta/policies/{id}/members:add", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AddPolicyMembersReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/CreateRole", "auth:roles", "create", "POST", "/iam/v2beta/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateRoleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/ListRoles", "auth:roles", "read", "GET", "/iam/v2beta/roles", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/GetRole", "auth:roles:{id}", "get", "GET", "/iam/v2beta/roles/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetRoleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/DeleteRole", "auth:roles:{id}", "delete", "DELETE", "/iam/v2beta/roles/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteRoleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/UpdateRole", "auth:roles:{id}", "update", "PUT", "/iam/v2beta/roles/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateRoleReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/CreateProject", "auth:projects", "create", "POST", "/iam/v2beta/projects", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateProjectReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/UpdateProject", "auth:projects:{id}", "update", "PUT", "/iam/v2beta/projects/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateProjectReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/GetProject", "auth:projects:{id}", "get", "GET", "/iam/v2beta/projects/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetProjectReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/ListProjects", "auth:projects", "read", "GET", "/iam/v2beta/projects", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/DeleteProject", "auth:projects:{id}", "delete", "DELETE", "/iam/v2beta/projects/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteProjectReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/UpgradeToV2", "system:iam:upgrade_to_v2", "upgrade", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2beta.Policies/ResetToV1", "system:iam:reset_to_v1", "reset", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}