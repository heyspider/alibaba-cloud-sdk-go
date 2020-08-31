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

// DeleteCorpGroup invokes the vcs.DeleteCorpGroup API synchronously
// api document: https://help.aliyun.com/api/vcs/deletecorpgroup.html
func (client *Client) DeleteCorpGroup(request *DeleteCorpGroupRequest) (response *DeleteCorpGroupResponse, err error) {
	response = CreateDeleteCorpGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCorpGroupWithChan invokes the vcs.DeleteCorpGroup API asynchronously
// api document: https://help.aliyun.com/api/vcs/deletecorpgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCorpGroupWithChan(request *DeleteCorpGroupRequest) (<-chan *DeleteCorpGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteCorpGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCorpGroup(request)
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

// DeleteCorpGroupWithCallback invokes the vcs.DeleteCorpGroup API asynchronously
// api document: https://help.aliyun.com/api/vcs/deletecorpgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCorpGroupWithCallback(request *DeleteCorpGroupRequest, callback func(response *DeleteCorpGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCorpGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteCorpGroup(request)
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

// DeleteCorpGroupRequest is the request struct for api DeleteCorpGroup
type DeleteCorpGroupRequest struct {
	*requests.RpcRequest
	CorpId  string `position:"Body" name:"CorpId"`
	GroupId string `position:"Body" name:"GroupId"`
}

// DeleteCorpGroupResponse is the response struct for api DeleteCorpGroup
type DeleteCorpGroupResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteCorpGroupRequest creates a request to invoke DeleteCorpGroup API
func CreateDeleteCorpGroupRequest() (request *DeleteCorpGroupRequest) {
	request = &DeleteCorpGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "DeleteCorpGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCorpGroupResponse creates a response to parse from DeleteCorpGroup response
func CreateDeleteCorpGroupResponse() (response *DeleteCorpGroupResponse) {
	response = &DeleteCorpGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
