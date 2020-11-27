package feishu

import (
	"asura/feishu/core/body"
	"asura/feishu/core/urls"

	"github.com/franela/goreq"
	jsoniter "github.com/json-iterator/go"
)

//"github.com/garudaen/asura/goreq"
// @Title  飞书发送消息模块
// @Description  飞书发送消息模块
// @Author  张振华  2020-10-31
// @Update  张振华  2020-10-31

var json = jsoniter.ConfigCompatibleWithStandardLibrary //jsoniter 初始化
//Auth 授权
type Auth struct {
	AppID             string `json:"app_id"`
	UserID            string `json:"user_id"`
	AppSecret         string `json:"app_secret"`
	AppAccessToken    string `json:"app_access_token"`
	TenantAccessToken string `json:"tenant_access_token"`
	Authorization     string `json:"authorization"`
}

//GetAccessToken 获取 access token
func (auth *Auth) GetAccessToken() *body.AppAccessTokenInternalRespBody {
	postBody, _ := json.Marshal(auth)
	res := &body.AppAccessTokenInternalRespBody{}
	req, _ := goreq.Request{
		Method:      "POST",
		Uri:         urls.APIAppAccessTokenInternal,
		Body:        postBody,
		ContentType: "application/json",
	}.Do()
	req.Body.FromJsonTo(&res)
	auth.AppAccessToken = res.AppAccessToken
	auth.TenantAccessToken = res.TenantAccessToken
	return res
}
