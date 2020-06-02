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

// CreateAlertContact invokes the arms.CreateAlertContact API synchronously
// api document: https://help.aliyun.com/api/arms/createalertcontact.html
func (client *Client) CreateAlertContact(request *CreateAlertContactRequest) (response *CreateAlertContactResponse, err error) {
	response = CreateCreateAlertContactResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAlertContactWithChan invokes the arms.CreateAlertContact API asynchronously
// api document: https://help.aliyun.com/api/arms/createalertcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlertContactWithChan(request *CreateAlertContactRequest) (<-chan *CreateAlertContactResponse, <-chan error) {
	responseChan := make(chan *CreateAlertContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlertContact(request)
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

// CreateAlertContactWithCallback invokes the arms.CreateAlertContact API asynchronously
// api document: https://help.aliyun.com/api/arms/createalertcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlertContactWithCallback(request *CreateAlertContactRequest, callback func(response *CreateAlertContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlertContactResponse
		var err error
		defer close(result)
		response, err = client.CreateAlertContact(request)
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

// CreateAlertContactRequest is the request struct for api CreateAlertContact
type CreateAlertContactRequest struct {
	*requests.RpcRequest
	PhoneNum            string           `position:"Query" name:"PhoneNum"`
	ProxyUserId         string           `position:"Query" name:"ProxyUserId"`
	ContactName         string           `position:"Query" name:"ContactName"`
	DingRobotWebhookUrl string           `position:"Query" name:"DingRobotWebhookUrl"`
	Email               string           `position:"Query" name:"Email"`
	SystemNoc           requests.Boolean `position:"Query" name:"SystemNoc"`
}

// CreateAlertContactResponse is the response struct for api CreateAlertContact
type CreateAlertContactResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ContactId string `json:"ContactId" xml:"ContactId"`
}

// CreateCreateAlertContactRequest creates a request to invoke CreateAlertContact API
func CreateCreateAlertContactRequest() (request *CreateAlertContactRequest) {
	request = &CreateAlertContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateAlertContact", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAlertContactResponse creates a response to parse from CreateAlertContact response
func CreateCreateAlertContactResponse() (response *CreateAlertContactResponse) {
	response = &CreateAlertContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
