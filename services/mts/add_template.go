package mts

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

func (client *Client) AddTemplate(request *AddTemplateRequest) (response *AddTemplateResponse, err error) {
	response = CreateAddTemplateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddTemplateWithChan(request *AddTemplateRequest) (<-chan *AddTemplateResponse, <-chan error) {
	responseChan := make(chan *AddTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTemplate(request)
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

func (client *Client) AddTemplateWithCallback(request *AddTemplateRequest, callback func(response *AddTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddTemplate(request)
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

type AddTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	Container            string           `position:"Query" name:"Container"`
	Video                string           `position:"Query" name:"Video"`
	Audio                string           `position:"Query" name:"Audio"`
	TransConfig          string           `position:"Query" name:"TransConfig"`
	MuxConfig            string           `position:"Query" name:"MuxConfig"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type AddTemplateResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Template  Template `json:"Template" xml:"Template"`
}

func CreateAddTemplateRequest() (request *AddTemplateRequest) {
	request = &AddTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddTemplate", "mts", "openAPI")
	return
}

func CreateAddTemplateResponse() (response *AddTemplateResponse) {
	response = &AddTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}