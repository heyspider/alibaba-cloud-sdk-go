package ddoscoo

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

// DescribeBackSourceCidr invokes the ddoscoo.DescribeBackSourceCidr API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describebacksourcecidr.html
func (client *Client) DescribeBackSourceCidr(request *DescribeBackSourceCidrRequest) (response *DescribeBackSourceCidrResponse, err error) {
	response = CreateDescribeBackSourceCidrResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackSourceCidrWithChan invokes the ddoscoo.DescribeBackSourceCidr API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describebacksourcecidr.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackSourceCidrWithChan(request *DescribeBackSourceCidrRequest) (<-chan *DescribeBackSourceCidrResponse, <-chan error) {
	responseChan := make(chan *DescribeBackSourceCidrResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackSourceCidr(request)
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

// DescribeBackSourceCidrWithCallback invokes the ddoscoo.DescribeBackSourceCidr API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describebacksourcecidr.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackSourceCidrWithCallback(request *DescribeBackSourceCidrRequest, callback func(response *DescribeBackSourceCidrResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackSourceCidrResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackSourceCidr(request)
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

// DescribeBackSourceCidrRequest is the request struct for api DescribeBackSourceCidr
type DescribeBackSourceCidrRequest struct {
	*requests.RpcRequest
	Line            string `position:"Query" name:"Line"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

// DescribeBackSourceCidrResponse is the response struct for api DescribeBackSourceCidr
type DescribeBackSourceCidrResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Cidrs     []string `json:"Cidrs" xml:"Cidrs"`
}

// CreateDescribeBackSourceCidrRequest creates a request to invoke DescribeBackSourceCidr API
func CreateDescribeBackSourceCidrRequest() (request *DescribeBackSourceCidrRequest) {
	request = &DescribeBackSourceCidrRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeBackSourceCidr", "ddoscoo", "openAPI")
	return
}

// CreateDescribeBackSourceCidrResponse creates a response to parse from DescribeBackSourceCidr response
func CreateDescribeBackSourceCidrResponse() (response *DescribeBackSourceCidrResponse) {
	response = &DescribeBackSourceCidrResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
