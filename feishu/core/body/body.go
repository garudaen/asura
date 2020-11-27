package body

//CommonBody 默认返回值
type CommonBody struct {
	Code int    `json:"code"` //状态码
	Msg  string `json:"msg"`  //错误描述
}

//AppAccessTokenInternalRespBody 获取access_token返回值
type AppAccessTokenInternalRespBody struct {
	CommonBody
	TenantAccessToken string `json:"tenant_access_token"`
	AppAccessToken    string `json:"app_access_token"`
	Expire            int64  `json:"expire"`
}

//AppAccessTokenRespBody 获取access_token返回值
type AppAccessTokenRespBody struct {
	CommonBody
	AppAccessToken string `json:"app_access_token"`
	Expire         int64  `json:"expire"`
}

//TenantAccessTokenRespBody 获取企业授权 token返回值
type TenantAccessTokenRespBody struct {
	CommonBody
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}

type userID struct {
	OpenID string `json:"open_id"` //成员 user_id
	UserID string `json:"user_id"` //成员 open_id
}

//UserInfo 用户信息，用于发送消息
type UserInfo struct {
	CommonBody
	userID
}

//SendMessageRespBody 发送消息返回值
//{'code': 0, 'data': {'message_id': 'om_de7ee73d07ae2ad153b34bd96fa4194b'}, 'msg': 'ok'}
type SendMessageRespBody struct {
	CommonBody
	Data messageID `json:"data"`
}
type messageID struct {
	MessageID string `json:"message_id"`
}

//SendTextMessageToUserPostBody 发送文本消息结构体,POST请求
type SendTextMessageToUserPostBody struct {
	UserID  string         `json:"user_id"`  //给用户发私聊消息，只需要填 open_id、email、user_id 中的一个即可，这里统一用 user_id
	MsgType string         `json:"msg_type"` //消息类型，此处固定填 "text"
	Content messageContent `json:"content"`  //消息内容
}

//SendTextMessageToChatGroupPostBody 发送文本消息结构体,POST请求
//chat_id ：会话（包括单聊、群聊）的唯一标识。群列表接口返回。
type SendTextMessageToChatGroupPostBody struct {
	ChatID  string         `json:"chat_id"`  //给用户发私聊消息，只需要填 open_id、email、user_id 中的一个即可，这里统一用 user_id
	MsgType string         `json:"msg_type"` //消息类型，此处固定填 "text"
	Content messageContent `json:"content"`  //消息内容
}

type messageContent struct {
	Text string `json:"text"`
}

//CreateChatGroupPostBody 创建聊天群接口，POST 请求
//例： {"name":"group name","user_ids":["33417745","cb93bdca"],"only_owner_add":false,"share_allowed":true,"only_owner_at_all":false,"only_owner_edit":false}
type CreateChatGroupPostBody struct {
	GroupName      string   `json:"name"`              //群的名称
	UserIDs        []string `json:"user_ids"`          //成员 user_id 列表，最多可以传200个
	OnlyOwnerAdd   bool     `json:"only_owner_add"`    //是否仅群主可以添加人
	ShareAllowed   bool     `json:"share_allowed"`     //是否允许分享群
	OnlyOwnerAtAll bool     `json:"only_owner_at_all"` //是否仅群主@all
	OnlyOwnerEdit  bool     `json:"only_owner_edit"`   //是否仅群主可编辑群信息，群信息包括头像、名称、描述、公告
}

//CreateChatGroupPostBodyDefault 设置默认值
func CreateChatGroupPostBodyDefault() CreateChatGroupPostBody {
	return CreateChatGroupPostBody{
		OnlyOwnerAdd:   false,
		ShareAllowed:   true,
		OnlyOwnerAtAll: false,
		OnlyOwnerEdit:  false,
	}
}

//CreateChatGroupRespBody 创建群组返回值结构体
type CreateChatGroupRespBody struct {
	CommonBody
	Data chatID `json:"data"`
}
type chatID struct {
	ChatID string `json:"chat_id"`
}
