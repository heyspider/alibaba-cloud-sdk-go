package yundun_dbaudit

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

// DescribeInstanceAttribue invokes the yundun_dbaudit.DescribeInstanceAttribue API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstanceattribue.html
func (client *Client) DescribeInstanceAttribue(request *DescribeInstanceAttribueRequest) (response *DescribeInstanceAttribueResponse, err error) {
	response = CreateDescribeInstanceAttribueResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceAttribueWithChan invokes the yundun_dbaudit.DescribeInstanceAttribue API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstanceattribue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceAttribueWithChan(request *DescribeInstanceAttribueRequest) (<-chan *DescribeInstanceAttribueResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAttribueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAttribue(request)
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

// DescribeInstanceAttribueWithCallback invokes the yundun_dbaudit.DescribeInstanceAttribue API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstanceattribue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceAttribueWithCallback(request *DescribeInstanceAttribueRequest, callback func(response *DescribeInstanceAttribueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAttribueResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAttribue(request)
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

// DescribeInstanceAttribueRequest is the request struct for api DescribeInstanceAttribue
type DescribeInstanceAttribueRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
}

// DescribeInstanceAttribueResponse is the response struct for api DescribeInstanceAttribue
type DescribeInstanceAttribueResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	InstanceAttribue InstanceAttribue `json:"InstanceAttribue" xml:"InstanceAttribue"`
}

// CreateDescribeInstanceAttribueRequest creates a request to invoke DescribeInstanceAttribue API
func CreateDescribeInstanceAttribueRequest() (request *DescribeInstanceAttribueRequest) {
	request = &DescribeInstanceAttribueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "DescribeInstanceAttribue", "dbaudit", "openAPI")
	return
}

// CreateDescribeInstanceAttribueResponse creates a response to parse from DescribeInstanceAttribue response
func CreateDescribeInstanceAttribueResponse() (response *DescribeInstanceAttribueResponse) {
	response = &DescribeInstanceAttribueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}