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

// SyncRulesItem is a nested struct in cr response
type SyncRulesItem struct {
	SyncRuleId          string `json:"SyncRuleId" xml:"SyncRuleId"`
	LocalRegionId       string `json:"LocalRegionId" xml:"LocalRegionId"`
	LocalInstanceId     string `json:"LocalInstanceId" xml:"LocalInstanceId"`
	LocalNamespaceName  string `json:"LocalNamespaceName" xml:"LocalNamespaceName"`
	LocalRepoName       string `json:"LocalRepoName" xml:"LocalRepoName"`
	TargetRegionId      string `json:"TargetRegionId" xml:"TargetRegionId"`
	TargetInstanceId    string `json:"TargetInstanceId" xml:"TargetInstanceId"`
	TargetNamespaceName string `json:"TargetNamespaceName" xml:"TargetNamespaceName"`
	TargetRepoName      string `json:"TargetRepoName" xml:"TargetRepoName"`
	TagFilter           string `json:"TagFilter" xml:"TagFilter"`
	SyncScope           string `json:"SyncScope" xml:"SyncScope"`
	SyncDirection       string `json:"SyncDirection" xml:"SyncDirection"`
	CreateTime          int64  `json:"CreateTime" xml:"CreateTime"`
	ModifiedTime        int64  `json:"ModifiedTime" xml:"ModifiedTime"`
	SyncRuleName        string `json:"SyncRuleName" xml:"SyncRuleName"`
	SyncTrigger         string `json:"SyncTrigger" xml:"SyncTrigger"`
}
