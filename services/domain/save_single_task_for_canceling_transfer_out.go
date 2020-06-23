package domain

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

// SaveSingleTaskForCancelingTransferOut invokes the domain.SaveSingleTaskForCancelingTransferOut API synchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforcancelingtransferout.html
func (client *Client) SaveSingleTaskForCancelingTransferOut(request *SaveSingleTaskForCancelingTransferOutRequest) (response *SaveSingleTaskForCancelingTransferOutResponse, err error) {
	response = CreateSaveSingleTaskForCancelingTransferOutResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForCancelingTransferOutWithChan invokes the domain.SaveSingleTaskForCancelingTransferOut API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforcancelingtransferout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForCancelingTransferOutWithChan(request *SaveSingleTaskForCancelingTransferOutRequest) (<-chan *SaveSingleTaskForCancelingTransferOutResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForCancelingTransferOutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForCancelingTransferOut(request)
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

// SaveSingleTaskForCancelingTransferOutWithCallback invokes the domain.SaveSingleTaskForCancelingTransferOut API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforcancelingtransferout.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForCancelingTransferOutWithCallback(request *SaveSingleTaskForCancelingTransferOutRequest, callback func(response *SaveSingleTaskForCancelingTransferOutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForCancelingTransferOutResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForCancelingTransferOut(request)
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

// SaveSingleTaskForCancelingTransferOutRequest is the request struct for api SaveSingleTaskForCancelingTransferOut
type SaveSingleTaskForCancelingTransferOutRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// SaveSingleTaskForCancelingTransferOutResponse is the response struct for api SaveSingleTaskForCancelingTransferOut
type SaveSingleTaskForCancelingTransferOutResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForCancelingTransferOutRequest creates a request to invoke SaveSingleTaskForCancelingTransferOut API
func CreateSaveSingleTaskForCancelingTransferOutRequest() (request *SaveSingleTaskForCancelingTransferOutRequest) {
	request = &SaveSingleTaskForCancelingTransferOutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForCancelingTransferOut", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForCancelingTransferOutResponse creates a response to parse from SaveSingleTaskForCancelingTransferOut response
func CreateSaveSingleTaskForCancelingTransferOutResponse() (response *SaveSingleTaskForCancelingTransferOutResponse) {
	response = &SaveSingleTaskForCancelingTransferOutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
