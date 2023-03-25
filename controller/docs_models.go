package controller

import "ezTikTok/models"

// _ResponseCommunityList 社区列表接口响应数据
type _ResponseCommunityList struct {
	Code    ResCode             `json:"code"`    // 业务响应状态码
	Message string              `json:"message"` // 提示信息
	Data    []*models.Community `json:"data"`    // 数据
}

// _ResponseCommunityDetail 社区详情接口响应数据
type _ResponseCommunityDetail struct {
	Code    ResCode                   `json:"code"`    // 业务响应状态码
	Message string                    `json:"message"` // 提示信息
	Data    []*models.CommunityDetail `json:"data"`    // 数据
}

// _SignUp 社区详情接口响应数据
type _SignUp struct {
	Code    ResCode               `json:"code"`    // 业务响应状态码
	Message string                `json:"message"` // 提示信息
	Data    []*models.ParamSignUp `json:"data"`    // 数据
}

// _Login 社区详情接口响应数据
type _Login struct {
	Code    ResCode              `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    []*models.ParamLogin `json:"data"`    // 数据
}

// _CreatePost 创建帖子接口响应数据
type _CreatePost struct {
	Code    ResCode        `json:"code"`    // 业务响应状态码
	Message string         `json:"message"` // 提示信息
	Data    []*models.Post `json:"data"`    // 数据
}

// _GetPostDetail 获取帖子详情接口响应数据
type _GetPostDetail struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

// _PostVote 帖子列表接口响应数据
type _PostVote struct {
	Code    ResCode             `json:"code"`    // 业务响应状态码
	Message string              `json:"message"` // 提示信息
	Data    []*models.ParamVote `json:"data"`    // 数据
}
