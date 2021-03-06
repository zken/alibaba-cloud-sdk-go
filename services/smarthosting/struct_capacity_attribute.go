package smarthosting

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CapacityAttribute is a nested struct in smarthosting response
type CapacityAttribute struct {
	AvailableHostCount int `json:"AvailableHostCount" xml:"AvailableHostCount"`
	TotalHostCount     int `json:"TotalHostCount" xml:"TotalHostCount"`
	AvailableMemories  int `json:"AvailableMemories" xml:"AvailableMemories"`
	ReservedMemories   int `json:"ReservedMemories" xml:"ReservedMemories"`
	AvailableVcpus     int `json:"AvailableVcpus" xml:"AvailableVcpus"`
	TotalVcpus         int `json:"TotalVcpus" xml:"TotalVcpus"`
	ReservedVcpus      int `json:"ReservedVcpus" xml:"ReservedVcpus"`
	ReservedHostCount  int `json:"ReservedHostCount" xml:"ReservedHostCount"`
	TotalMemories      int `json:"TotalMemories" xml:"TotalMemories"`
}
