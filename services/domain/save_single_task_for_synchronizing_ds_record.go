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

// SaveSingleTaskForSynchronizingDSRecord invokes the domain.SaveSingleTaskForSynchronizingDSRecord API synchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforsynchronizingdsrecord.html
func (client *Client) SaveSingleTaskForSynchronizingDSRecord(request *SaveSingleTaskForSynchronizingDSRecordRequest) (response *SaveSingleTaskForSynchronizingDSRecordResponse, err error) {
	response = CreateSaveSingleTaskForSynchronizingDSRecordResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForSynchronizingDSRecordWithChan invokes the domain.SaveSingleTaskForSynchronizingDSRecord API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforsynchronizingdsrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForSynchronizingDSRecordWithChan(request *SaveSingleTaskForSynchronizingDSRecordRequest) (<-chan *SaveSingleTaskForSynchronizingDSRecordResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForSynchronizingDSRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForSynchronizingDSRecord(request)
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

// SaveSingleTaskForSynchronizingDSRecordWithCallback invokes the domain.SaveSingleTaskForSynchronizingDSRecord API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskforsynchronizingdsrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForSynchronizingDSRecordWithCallback(request *SaveSingleTaskForSynchronizingDSRecordRequest, callback func(response *SaveSingleTaskForSynchronizingDSRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForSynchronizingDSRecordResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForSynchronizingDSRecord(request)
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

// SaveSingleTaskForSynchronizingDSRecordRequest is the request struct for api SaveSingleTaskForSynchronizingDSRecord
type SaveSingleTaskForSynchronizingDSRecordRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// SaveSingleTaskForSynchronizingDSRecordResponse is the response struct for api SaveSingleTaskForSynchronizingDSRecord
type SaveSingleTaskForSynchronizingDSRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForSynchronizingDSRecordRequest creates a request to invoke SaveSingleTaskForSynchronizingDSRecord API
func CreateSaveSingleTaskForSynchronizingDSRecordRequest() (request *SaveSingleTaskForSynchronizingDSRecordRequest) {
	request = &SaveSingleTaskForSynchronizingDSRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForSynchronizingDSRecord", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForSynchronizingDSRecordResponse creates a response to parse from SaveSingleTaskForSynchronizingDSRecord response
func CreateSaveSingleTaskForSynchronizingDSRecordResponse() (response *SaveSingleTaskForSynchronizingDSRecordResponse) {
	response = &SaveSingleTaskForSynchronizingDSRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
