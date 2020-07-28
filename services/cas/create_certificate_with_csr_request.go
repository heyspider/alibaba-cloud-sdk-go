package cas

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

// CreateCertificateWithCsrRequest invokes the cas.CreateCertificateWithCsrRequest API synchronously
// api document: https://help.aliyun.com/api/cas/createcertificatewithcsrrequest.html
func (client *Client) CreateCertificateWithCsrRequest(request *CreateCertificateWithCsrRequestRequest) (response *CreateCertificateWithCsrRequestResponse, err error) {
	response = CreateCreateCertificateWithCsrRequestResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCertificateWithCsrRequestWithChan invokes the cas.CreateCertificateWithCsrRequest API asynchronously
// api document: https://help.aliyun.com/api/cas/createcertificatewithcsrrequest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCertificateWithCsrRequestWithChan(request *CreateCertificateWithCsrRequestRequest) (<-chan *CreateCertificateWithCsrRequestResponse, <-chan error) {
	responseChan := make(chan *CreateCertificateWithCsrRequestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCertificateWithCsrRequest(request)
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

// CreateCertificateWithCsrRequestWithCallback invokes the cas.CreateCertificateWithCsrRequest API asynchronously
// api document: https://help.aliyun.com/api/cas/createcertificatewithcsrrequest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCertificateWithCsrRequestWithCallback(request *CreateCertificateWithCsrRequestRequest, callback func(response *CreateCertificateWithCsrRequestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCertificateWithCsrRequestResponse
		var err error
		defer close(result)
		response, err = client.CreateCertificateWithCsrRequest(request)
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

// CreateCertificateWithCsrRequestRequest is the request struct for api CreateCertificateWithCsrRequest
type CreateCertificateWithCsrRequestRequest struct {
	*requests.RpcRequest
	ProductCode  string `position:"Query" name:"ProductCode"`
	Csr          string `position:"Query" name:"Csr"`
	ValidateType string `position:"Query" name:"ValidateType"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	Email        string `position:"Query" name:"Email"`
	Phone        string `position:"Query" name:"Phone"`
	Username     string `position:"Query" name:"Username"`
}

// CreateCertificateWithCsrRequestResponse is the response struct for api CreateCertificateWithCsrRequest
type CreateCertificateWithCsrRequestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   int64  `json:"OrderId" xml:"OrderId"`
}

// CreateCreateCertificateWithCsrRequestRequest creates a request to invoke CreateCertificateWithCsrRequest API
func CreateCreateCertificateWithCsrRequestRequest() (request *CreateCertificateWithCsrRequestRequest) {
	request = &CreateCertificateWithCsrRequestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "CreateCertificateWithCsrRequest", "cas-pack", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCertificateWithCsrRequestResponse creates a response to parse from CreateCertificateWithCsrRequest response
func CreateCreateCertificateWithCsrRequestResponse() (response *CreateCertificateWithCsrRequestResponse) {
	response = &CreateCertificateWithCsrRequestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
