package cr_2018_12_01

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

// DeleteChartNamespace invokes the cr.DeleteChartNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/deletechartnamespace.html
func (client *Client) DeleteChartNamespace(request *DeleteChartNamespaceRequest) (response *DeleteChartNamespaceResponse, err error) {
	response = CreateDeleteChartNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteChartNamespaceWithChan invokes the cr.DeleteChartNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/deletechartnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChartNamespaceWithChan(request *DeleteChartNamespaceRequest) (<-chan *DeleteChartNamespaceResponse, <-chan error) {
	responseChan := make(chan *DeleteChartNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteChartNamespace(request)
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

// DeleteChartNamespaceWithCallback invokes the cr.DeleteChartNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/deletechartnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChartNamespaceWithCallback(request *DeleteChartNamespaceRequest, callback func(response *DeleteChartNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteChartNamespaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteChartNamespace(request)
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

// DeleteChartNamespaceRequest is the request struct for api DeleteChartNamespace
type DeleteChartNamespaceRequest struct {
	*requests.RpcRequest
	NamespaceName string `position:"Query" name:"NamespaceName"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// DeleteChartNamespaceResponse is the response struct for api DeleteChartNamespace
type DeleteChartNamespaceResponse struct {
	*responses.BaseResponse
	DeleteChartNamespaceIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                          string `json:"Code" xml:"Code"`
	RequestId                     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteChartNamespaceRequest creates a request to invoke DeleteChartNamespace API
func CreateDeleteChartNamespaceRequest() (request *DeleteChartNamespaceRequest) {
	request = &DeleteChartNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "DeleteChartNamespace", "cr", "openAPI")
	return
}

// CreateDeleteChartNamespaceResponse creates a response to parse from DeleteChartNamespace response
func CreateDeleteChartNamespaceResponse() (response *DeleteChartNamespaceResponse) {
	response = &DeleteChartNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
