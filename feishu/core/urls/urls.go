package urls

//API 飞书服务端api接口
const (
	APIAppAccessTokenInternal    = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/"    //获取 app_access_token（企业自建应用）
	APIAppAccessToken            = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/"             //获取 app_access_token（应用商店应用）
	APITenantAccessToken         = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/"          //获取 tenant_access_token（应用商店应用）
	APITenantAccessTokenInternal = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/" //获取 tenant_access_token（企业自建应用）
	APIAppTicketResend           = "https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend/"            //重新推送 app_ticket
	APIOAuth2AccessToken         = "https://open.feishu.cn/connect/qrconnect/oauth2/access_token/"          //获取登录用户身份
	APITokenLoginValidate        = "https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate"            //code2session
	APIRefreshAccessToken        = "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token"        //刷新access_token

	////////群信息和群管理
	APICreateChatGroup = "https://open.feishu.cn/open-apis/chat/v4/create/"         //创建群
	APIChatList        = "https://open.feishu.cn/open-apis/chat/v4/list"            //获取群列表
	APIChatInfo        = "https://open.feishu.cn/open-apis/chat/v4/info"            //获取群信息
	APIUpdateChat      = "https://open.feishu.cn/open-apis/chat/v4/update/"         //更新群信息
	APIAddChatUser     = "https://open.feishu.cn/open-apis/chat/v4/chatter/add/"    //拉用户进群
	APIRemoveChatUser  = "https://open.feishu.cn/open-apis/chat/v4/chatter/delete/" //移除用户出群
	APIDisbandChat     = "https://open.feishu.cn/open-apis/chat/v4/disband"         //解散群
	APIAddBot          = "https://open.feishu.cn/open-apis/bot/v4/add"              //拉机器人进群

	//////////////////部门和用户
	APIScope                      = "https://open.feishu.cn/open-apis/contact/v1/scope/get"                   //获取通讯录授权范围
	APIScopeV2                    = "https://open.feishu.cn/open-apis/contact/v2/scope/get"                   //获取通讯录授权范围v2
	APIDepartmentSimpleList       = "https://open.feishu.cn/open-apis/contact/v1/department/simple/list"      //获取部门列表
	APIDepartmentSimpleListV2     = "https://open.feishu.cn/open-apis/contact/v2/department/simple/list"      //获取部门列表 v2
	APIDepartmentInfoGet          = "https://open.feishu.cn/open-apis/contact/v1/department/info/get"         //获取部门详情
	APIDepartmentInfoBatchGet     = "https://open.feishu.cn/open-apis/contact/v2/department/detail/batch_get" //批量获取部门详情
	APIDepartmentUserList         = "https://open.feishu.cn/open-apis/contact/v1/department/user/list"        //获取部门用户列表
	APIDepartmentUserListV2       = "https://open.feishu.cn/open-apis/contact/v2/department/user/list"        //获取部门用户列表v2
	APIDepartmentUserDetailList   = "https://open.feishu.cn/open-apis/contact/v1/department/user/detail/list" //获取用户详情
	APIDepartmentUserDetailListV2 = "https://open.feishu.cn/open-apis/contact/v2/department/user/detail/list" //获取用户详情v2
	APIGetUserFromPhone           = "https://open.feishu.cn/open-apis/user/v1/batch_get_id"                   //通过手机号获取用户信息
	APIUserBatchGet               = "https://open.feishu.cn/open-apis/contact/v1/user/batch_get"              //批量获取用户信息
	APIUserBatchGetV2             = "https://open.feishu.cn/open-apis/contact/v2/user/batch_get"              //批量获取用户信息v2
	APISearchUser                 = "https://open.feishu.cn/open-apis/search/v1/user"                         //搜索用户F
	APIIsUserAdmin                = "https://open.feishu.cn/open-apis/application/v3/is_user_admin"           //检验应用管理员
	APIUserGroupLIst              = "https://open.feishu.cn/open-apis/user/v4/group_list"                     //获取用户所在的群列表
	APIChatMembers                = "https://open.feishu.cn/open-apis/chat/v4/members"                        //获取群成员列表
	APIChatSearch                 = "https://open.feishu.cn/open-apis/chat/v4/search"                         //搜索用户所在的群列表

	//////////////////机器人发送消息
	APIRobotSendMessage      = "https://open.feishu.cn/open-apis/message/v4/send/"       //机器人发送消息
	APIRobotSendBatchMessage = "https://open.feishu.cn/open-apis/message/v4/batch_send/" //机器人批量发送消息

	//////////////////角色
	APIRoleList       = "https://open.feishu.cn/open-apis/contact/v2/role/list"    //获取角色列表
	APIRoleMemberList = "https://open.feishu.cn/open-apis/contact/v2/role/members" //获取角色成员列表

	//////////////////日历与行程
	APICalendarCreate               = "https://open.feishu.cn/open-apis/calendar/v3/calendars"                        //创建日历
	APICalendarGet                  = "https://open.feishu.cn/open-apis/calendar/v3/calendar_list/%s"                 //获取日历
	APICalendarListGet              = "https://open.feishu.cn/open-apis/calendar/v3/calendar_list"                    //获取日历列表
	APICalendarUpdate               = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s"                     //更新日历
	APICalendarEventCreate          = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/events"              //创建日程
	APICalendarEventDelete          = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/events/%s"           //删除日程
	APICalendarEventAttendeesUpdate = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/events/%s/attendees" //邀请/移除日程参与者
	APICalendarAttendeesGet         = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/acl"                 //获取访问控制列表
	APICalendarAttendeesDelete      = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/acl/%s"              //删除访问空值
)
