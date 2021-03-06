package rtc

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

// UserListItem is a nested struct in rtc response
type UserListItem struct {
	User         string `json:"User" xml:"User"`
	SubAudio     int    `json:"SubAudio" xml:"SubAudio"`
	SubVideo1080 int    `json:"SubVideo1080" xml:"SubVideo1080"`
	PubVideo1080 int    `json:"PubVideo1080" xml:"PubVideo1080"`
	PubAudio     int    `json:"PubAudio" xml:"PubAudio"`
	ServiceArea  string `json:"ServiceArea" xml:"ServiceArea"`
	PubContent   int    `json:"PubContent" xml:"PubContent"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	EndTime      string `json:"EndTime" xml:"EndTime"`
	SubVideo360  int    `json:"SubVideo360" xml:"SubVideo360"`
	PubVideo720  int    `json:"PubVideo720" xml:"PubVideo720"`
	ChannelId    string `json:"ChannelId" xml:"ChannelId"`
	UserId       string `json:"UserId" xml:"UserId"`
	PubVideo360  int    `json:"PubVideo360" xml:"PubVideo360"`
	SubContent   int    `json:"SubContent" xml:"SubContent"`
	SubVideo720  int    `json:"SubVideo720" xml:"SubVideo720"`
}
