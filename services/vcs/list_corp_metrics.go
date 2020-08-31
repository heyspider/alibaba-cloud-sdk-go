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

// ListCorpMetrics invokes the vcs.ListCorpMetrics API synchronously
// api document: https://help.aliyun.com/api/vcs/listcorpmetrics.html
func (client *Client) ListCorpMetrics(request *ListCorpMetricsRequest) (response *ListCorpMetricsResponse, err error) {
	response = CreateListCorpMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCorpMetricsWithChan invokes the vcs.ListCorpMetrics API asynchronously
// api document: https://help.aliyun.com/api/vcs/listcorpmetrics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCorpMetricsWithChan(request *ListCorpMetricsRequest) (<-chan *ListCorpMetricsResponse, <-chan error) {
	responseChan := make(chan *ListCorpMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCorpMetrics(request)
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

// ListCorpMetricsWithCallback invokes the vcs.ListCorpMetrics API asynchronously
// api document: https://help.aliyun.com/api/vcs/listcorpmetrics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCorpMetricsWithCallback(request *ListCorpMetricsRequest, callback func(response *ListCorpMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCorpMetricsResponse
		var err error
		defer close(result)
		response, err = client.ListCorpMetrics(request)
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

// ListCorpMetricsRequest is the request struct for api ListCorpMetrics
type ListCorpMetricsRequest struct {
	*requests.RpcRequest
	CorpId          string `position:"Body" name:"CorpId"`
	EndTime         string `position:"Body" name:"EndTime"`
	StartTime       string `position:"Body" name:"StartTime"`
	PageNumber      string `position:"Body" name:"PageNumber"`
	DeviceGroupList string `position:"Body" name:"DeviceGroupList"`
	TagCode         string `position:"Body" name:"TagCode"`
	UserGroupList   string `position:"Body" name:"UserGroupList"`
	PageSize        string `position:"Body" name:"PageSize"`
	DeviceIdList    string `position:"Body" name:"DeviceIdList"`
}

// ListCorpMetricsResponse is the response struct for api ListCorpMetrics
type ListCorpMetricsResponse struct {
	*responses.BaseResponse
	Code       string     `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Success    string     `json:"Success" xml:"Success"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateListCorpMetricsRequest creates a request to invoke ListCorpMetrics API
func CreateListCorpMetricsRequest() (request *ListCorpMetricsRequest) {
	request = &ListCorpMetricsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListCorpMetrics", "", "")
	request.Method = requests.POST
	return
}

// CreateListCorpMetricsResponse creates a response to parse from ListCorpMetrics response
func CreateListCorpMetricsResponse() (response *ListCorpMetricsResponse) {
	response = &ListCorpMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
