package section

// PreAuthInfo 预授权码获取
type PreAuthInfo struct {
	Pre_auth_code string `json:"pre_auth_code"` //预授权码
	ExpiresIn     int64  `json:"expires_in"`
}

// AccessTokenInfo 获取第三平台accesstoken
type AccessTokenInfo struct {
	ErrCode                int64  `json:"errcode"`
	ErrMsg                 string `json:"errmsg"`
	Component_access_token string `json:"component_access_token"`
	ExpiresIn              int64  `json:"expires_in"` // 有效时间, seconds
}

// AuthorizationInfos 授权成功返回的信息
type AuthorizationInfos struct {
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

// AuthorizationInfo 授权方信息返回
type AuthorizationInfo struct {
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"` //刷新token
	OriginalID             string `json:"original_id"`
}

// AuthorizerInfos 获取授权方的基础信息
type AuthorizerInfos struct {
	AuthorizerInfo AuthorizerInfo `json:"authorizer_info"`
}

// AuthorizerInfo AuthorizerInfos具体信息
type AuthorizerInfo struct {
	NickName        string `json:"nick_name"`
	HeadImg         string `json:"head_img"`
	ServiceTypeInfo struct {
		ID int64 `json:"id"` //0代表订阅号，1代表升级后的订阅号，2代表服务号
	} `json:"service_type_info"`
	VerifyTypeInfo struct {
		ID int64 `json:"id"`
	} `json:"verify_type_info"`
	UserName      string `json:"user_name"`      //公众号原始id
	PrincipalName string `json:"principal_name"` //公众号名称
	QrcodeURL     string `json:"qrcode_url"`
	Appid         string `json:"appid"`
}

// RespData 返回给客户端数据
type RespData struct {
	AuthorizerInfo *AuthorizerInfo `json:"authorizer_info"`
}
