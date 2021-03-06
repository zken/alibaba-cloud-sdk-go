package arms

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

// CreateWehook invokes the arms.CreateWehook API synchronously
func (client *Client) CreateWehook(request *CreateWehookRequest) (response *CreateWehookResponse, err error) {
	response = CreateCreateWehookResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWehookWithChan invokes the arms.CreateWehook API asynchronously
func (client *Client) CreateWehookWithChan(request *CreateWehookRequest) (<-chan *CreateWehookResponse, <-chan error) {
	responseChan := make(chan *CreateWehookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWehook(request)
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

// CreateWehookWithCallback invokes the arms.CreateWehook API asynchronously
func (client *Client) CreateWehookWithCallback(request *CreateWehookRequest, callback func(response *CreateWehookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWehookResponse
		var err error
		defer close(result)
		response, err = client.CreateWehook(request)
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

// CreateWehookRequest is the request struct for api CreateWehook
type CreateWehookRequest struct {
	*requests.RpcRequest
	HttpHeaders string `position:"Query" name:"HttpHeaders"`
	Method      string `position:"Query" name:"Method"`
	HttpParams  string `position:"Query" name:"HttpParams"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
	Body        string `position:"Query" name:"Body"`
	Url         string `position:"Query" name:"Url"`
	ContactName string `position:"Query" name:"ContactName"`
}

// CreateWehookResponse is the response struct for api CreateWehook
type CreateWehookResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ContactId string `json:"ContactId" xml:"ContactId"`
}

// CreateCreateWehookRequest creates a request to invoke CreateWehook API
func CreateCreateWehookRequest() (request *CreateWehookRequest) {
	request = &CreateWehookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateWehook", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateWehookResponse creates a response to parse from CreateWehook response
func CreateCreateWehookResponse() (response *CreateWehookResponse) {
	response = &CreateWehookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
