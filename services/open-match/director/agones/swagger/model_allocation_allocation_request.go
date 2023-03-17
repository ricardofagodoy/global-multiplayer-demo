// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * proto/allocation/allocation.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AllocationAllocationRequest struct {
	Namespace string `json:"namespace,omitempty"`
	// If specified, multi-cluster policies are applied. Otherwise, allocation will happen locally.
	MultiClusterSetting *AllocationMultiClusterSetting `json:"multiClusterSetting,omitempty"`
	// Deprecated: Please use gameServerSelectors instead. This field is ignored if the gameServerSelectors field is set The required allocation. Defaults to all GameServers.
	RequiredGameServerSelector *AllocationGameServerSelector `json:"requiredGameServerSelector,omitempty"`
	// Deprecated: Please use gameServerSelectors instead. This field is ignored if the gameServerSelectors field is set The ordered list of preferred allocations out of the `required` set. If the first selector is not matched, the selection attempts the second selector, and so on.
	PreferredGameServerSelectors []AllocationGameServerSelector `json:"preferredGameServerSelectors,omitempty"`
	// Scheduling strategy. Defaults to \"Packed\".
	Scheduling *AllocationRequestSchedulingStrategy `json:"scheduling,omitempty"`
	MetaPatch *AllocationMetaPatch `json:"metaPatch,omitempty"`
	Metadata *AllocationMetaPatch `json:"metadata,omitempty"`
	// Ordered list of GameServer label selectors. If the first selector is not matched, the selection attempts the second selector, and so on. This is useful for things like smoke testing of new game servers. Note: This field can only be set if neither Required or Preferred is set.
	GameServerSelectors []AllocationGameServerSelector `json:"gameServerSelectors,omitempty"`
}