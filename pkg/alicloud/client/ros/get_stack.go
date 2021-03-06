// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetStack invokes the ros.GetStack API synchronously
// api document: https://www.alibabacloud.com/help/doc-detail/132088.htm
func (client *Client) GetStack(request *GetStackRequest) (response *GetStackResponse, err error) {
	response = CreateGetStackResponse()
	err = client.DoAction(request, response)
	return
}

// GetStackRequest is the request struct for api GetStack
type GetStackRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	StackId     string `position:"Query" name:"StackId"`
}

// GetStackResponse is the response struct for api GetStack
type GetStackResponse struct {
	*responses.BaseResponse
	CreateTime          string              `json:"CreateTime" xml:"CreateTime"`
	Description         string              `json:"Description" xml:"Description"`
	DisableRollback     bool                `json:"DisableRollback" xml:"DisableRollback"`
	RegionId            string              `json:"RegionId" xml:"RegionId"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	StackId             string              `json:"StackId" xml:"StackId"`
	StackName           string              `json:"StackName" xml:"StackName"`
	Status              string              `json:"Status" xml:"Status"`
	StatusReason        string              `json:"StatusReason" xml:"StatusReason"`
	TemplateDescription string              `json:"TemplateDescription" xml:"TemplateDescription"`
	TimeoutInMinutes    int                 `json:"TimeoutInMinutes" xml:"TimeoutInMinutes"`
	UpdateTime          string              `json:"UpdateTime" xml:"UpdateTime"`
	ParentStackId       string              `json:"ParentStackId" xml:"ParentStackId"`
	Outputs             []map[string]string `json:"Outputs" xml:"Outputs"`
	NotificationURLs    []string            `json:"NotificationURLs" xml:"NotificationURLs"`
	Parameters          []Parameter         `json:"Parameters" xml:"Parameters"`
}

// CreateGetStackRequest creates a request to invoke GetStack API
func CreateGetStackRequest() (request *GetStackRequest) {
	request = &GetStackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "GetStack", "ROS", "openAPI")
	return
}

// CreateGetStackResponse creates a response to parse from GetStack response
func CreateGetStackResponse() (response *GetStackResponse) {
	response = &GetStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
