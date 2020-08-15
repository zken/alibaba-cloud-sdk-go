package vcs

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

// GetFaceModelResult invokes the vcs.GetFaceModelResult API synchronously
// api document: https://help.aliyun.com/api/vcs/getfacemodelresult.html
func (client *Client) GetFaceModelResult(request *GetFaceModelResultRequest) (response *GetFaceModelResultResponse, err error) {
	response = CreateGetFaceModelResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetFaceModelResultWithChan invokes the vcs.GetFaceModelResult API asynchronously
// api document: https://help.aliyun.com/api/vcs/getfacemodelresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceModelResultWithChan(request *GetFaceModelResultRequest) (<-chan *GetFaceModelResultResponse, <-chan error) {
	responseChan := make(chan *GetFaceModelResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFaceModelResult(request)
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

// GetFaceModelResultWithCallback invokes the vcs.GetFaceModelResult API asynchronously
// api document: https://help.aliyun.com/api/vcs/getfacemodelresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceModelResultWithCallback(request *GetFaceModelResultRequest, callback func(response *GetFaceModelResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFaceModelResultResponse
		var err error
		defer close(result)
		response, err = client.GetFaceModelResult(request)
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

// GetFaceModelResultRequest is the request struct for api GetFaceModelResult
type GetFaceModelResultRequest struct {
	*requests.RpcRequest
	PictureUrl     string `position:"Body" name:"PictureUrl"`
	PictureContent string `position:"Body" name:"PictureContent"`
	PictureId      string `position:"Body" name:"PictureId"`
}

// GetFaceModelResultResponse is the response struct for api GetFaceModelResult
type GetFaceModelResultResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetFaceModelResultRequest creates a request to invoke GetFaceModelResult API
func CreateGetFaceModelResultRequest() (request *GetFaceModelResultRequest) {
	request = &GetFaceModelResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetFaceModelResult", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetFaceModelResultResponse creates a response to parse from GetFaceModelResult response
func CreateGetFaceModelResultResponse() (response *GetFaceModelResultResponse) {
	response = &GetFaceModelResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
