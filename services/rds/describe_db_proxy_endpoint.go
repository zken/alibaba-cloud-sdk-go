package rds

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBProxyEndpoint invokes the rds.DescribeDBProxyEndpoint API synchronously
// api document: https://help.aliyun.com/api/rds/describedbproxyendpoint.html
func (client *Client) DescribeDBProxyEndpoint(request *DescribeDBProxyEndpointRequest) (response *DescribeDBProxyEndpointResponse, err error) {
	response = CreateDescribeDBProxyEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBProxyEndpointWithChan invokes the rds.DescribeDBProxyEndpoint API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbproxyendpoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBProxyEndpointWithChan(request *DescribeDBProxyEndpointRequest) (<-chan *DescribeDBProxyEndpointResponse, <-chan error) {
	responseChan := make(chan *DescribeDBProxyEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBProxyEndpoint(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDBProxyEndpointWithCallback invokes the rds.DescribeDBProxyEndpoint API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbproxyendpoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBProxyEndpointWithCallback(request *DescribeDBProxyEndpointRequest, callback func(response *DescribeDBProxyEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBProxyEndpointResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBProxyEndpoint(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDBProxyEndpointRequest is the request struct for api DescribeDBProxyEndpoint
type DescribeDBProxyEndpointRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBProxyConnectString string           `position:"Query" name:"DBProxyConnectString"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBProxyEndpointId    string           `position:"Query" name:"DBProxyEndpointId"`
}

// DescribeDBProxyEndpointResponse is the response struct for api DescribeDBProxyEndpoint
type DescribeDBProxyEndpointResponse struct {
	*responses.BaseResponse
	RequestId                        string `json:"RequestId" xml:"RequestId"`
	DBProxyEndpointId                string `json:"DBProxyEndpointId" xml:"DBProxyEndpointId"`
	DBProxyConnectString             string `json:"DBProxyConnectString" xml:"DBProxyConnectString"`
	DBProxyConnectStringPort         string `json:"DBProxyConnectStringPort" xml:"DBProxyConnectStringPort"`
	DBProxyConnectStringNetType      string `json:"DBProxyConnectStringNetType" xml:"DBProxyConnectStringNetType"`
	DBProxyFeatures                  string `json:"DBProxyFeatures" xml:"DBProxyFeatures"`
	ReadOnlyInstanceMaxDelayTime     string `json:"ReadOnlyInstanceMaxDelayTime" xml:"ReadOnlyInstanceMaxDelayTime"`
	ReadOnlyInstanceDistributionType string `json:"ReadOnlyInstanceDistributionType" xml:"ReadOnlyInstanceDistributionType"`
	ReadOnlyInstanceWeight           string `json:"ReadOnlyInstanceWeight" xml:"ReadOnlyInstanceWeight"`
}

// CreateDescribeDBProxyEndpointRequest creates a request to invoke DescribeDBProxyEndpoint API
func CreateDescribeDBProxyEndpointRequest() (request *DescribeDBProxyEndpointRequest) {
	request = &DescribeDBProxyEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBProxyEndpoint", "Rds", "openAPI")
	return
}

// CreateDescribeDBProxyEndpointResponse creates a response to parse from DescribeDBProxyEndpoint response
func CreateDescribeDBProxyEndpointResponse() (response *DescribeDBProxyEndpointResponse) {
	response = &DescribeDBProxyEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}