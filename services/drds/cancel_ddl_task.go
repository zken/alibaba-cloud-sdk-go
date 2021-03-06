package drds

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

// CancelDDLTask invokes the drds.CancelDDLTask API synchronously
func (client *Client) CancelDDLTask(request *CancelDDLTaskRequest) (response *CancelDDLTaskResponse, err error) {
	response = CreateCancelDDLTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CancelDDLTaskWithChan invokes the drds.CancelDDLTask API asynchronously
func (client *Client) CancelDDLTaskWithChan(request *CancelDDLTaskRequest) (<-chan *CancelDDLTaskResponse, <-chan error) {
	responseChan := make(chan *CancelDDLTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelDDLTask(request)
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

// CancelDDLTaskWithCallback invokes the drds.CancelDDLTask API asynchronously
func (client *Client) CancelDDLTaskWithCallback(request *CancelDDLTaskRequest, callback func(response *CancelDDLTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelDDLTaskResponse
		var err error
		defer close(result)
		response, err = client.CancelDDLTask(request)
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

// CancelDDLTaskRequest is the request struct for api CancelDDLTask
type CancelDDLTaskRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	TaskId         string `position:"Query" name:"TaskId"`
}

// CancelDDLTaskResponse is the response struct for api CancelDDLTask
type CancelDDLTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelDDLTaskRequest creates a request to invoke CancelDDLTask API
func CreateCancelDDLTaskRequest() (request *CancelDDLTaskRequest) {
	request = &CancelDDLTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2015-04-13", "CancelDDLTask", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelDDLTaskResponse creates a response to parse from CancelDDLTask response
func CreateCancelDDLTaskResponse() (response *CancelDDLTaskResponse) {
	response = &CancelDDLTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
