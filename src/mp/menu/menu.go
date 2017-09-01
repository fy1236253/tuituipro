package menu

const (
	MenuButtonCountLimit    = 3 // 一级菜单最多包含 3 个按钮
	SubMenuButtonCountLimit = 5 // 二级菜单最多包含 5 个按钮
)

const (
	MenuButtonNameLenLimit    = 16 // 菜单标题不超过16个字节
	SubMenuButtonNameLenLimit = 40 // 子菜单标题不超过40个字节
)

const (
	ButtonKeyLenLimit = 128 // 菜单KEY值不能超过128字节
	ButtonURLLenLimit = 256 // 网页链接不能超过256字节
)

const (
	// 下面六個類型(包括view)是在公众平台官网通过网站功能发布菜单的按鈕類型
	ButtonTypeText  = "text"
	ButtonTypeImage = "img"
	ButtonTypePhoto = "photo"
	ButtonTypeVideo = "video"
	ButtonTypeVoice = "voice"

	// NOTE: 上面的按鈕類型不能通過API設置

	ButtonTypeView  = "view"  // 跳转URL
	ButtonTypeClick = "click" // 点击推事件

	// 仅支持微信iPhone5.4.1以上版本, 和Android5.4以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	ButtonTypeScanCodePush    = "scancode_push"      // 扫码推事件
	ButtonTypeScanCodeWaitMsg = "scancode_waitmsg"   // 扫码带提示
	ButtonTypePicSysPhoto     = "pic_sysphoto"       // 系统拍照发图
	ButtonTypePicPhotoOrAlbum = "pic_photo_or_album" // 拍照或者相册发图
	ButtonTypePicWeixin       = "pic_weixin"         // 微信相册发图
	ButtonTypeLocationSelect  = "location_select"    // 发送位置

	// 专门给第三方平台旗下未微信认证(具体而言, 是资质认证未通过)的订阅号准备的事件类型,
	// 它们是没有事件推送的, 能力相对受限, 其他类型的公众号不必使用.
	ButtonTypeMediaId     = "media_id"     // 下发消息
	ButtonTypeViewLimited = "view_limited" // 跳转图文消息URL
)

// type SubButton struct {
// 	Type    string   `json:"type"` // 非必须; 菜单的响应动作类型
// 	Name    string   `json:"name"` // 必须;  菜单标题, 不超过16个字节, 子菜单不超过40个字节
// 	Key     string   `json:"key"`  // 非必须; 菜单KEY值, 用于消息接口推送, 不超过128字节
// 	URL     string   `json:"url"`  // 非必须; 网页链接, 用户点击菜单可打开链接, 不超过256字节
// 	List    []Button `json:"list"`
// 	MediaId string   `json:"media_id"` // 非必须; 调用新增永久素材接口返回的合法media_id
// }

// NewsList 图文消息列表
type NewsList struct {
	Title      string `json:"title,omitempty"`
	Author     string `json:"author,omitempty"`
	Digest     string `json:"digest,omitempty"`     //摘要
	ShowCover  int    `json:"show_cover,omitempty"` //是否显示封面 0 不显示 1 显示
	CoverURL   string `json:"cover_url,omitempty"`
	ContentURL string `json:"content_url,omitempty"`
	SourceURL  string `json:"source_url,omitempty"`
}

// MenuJSON 菜单json结构体
type MenuJSON struct {
	Menu Menu `json:"menu"`
}

type SearchMenuJSON struct {
	IsMenuOpen   int  `json:"is_menu_open"`
	SelfMenuInfo Menu `json:"selfmenu_info"`
}

type Menu struct {
	Buttons []Button `json:"button"` // 一级菜单数组, 个数应为1~3个
}

//Button 菜单的按钮
type Button struct {
	Type     string `json:"type,omitempty"` // 非必须; 菜单的响应动作类型
	Name     string `json:"name,omitempty"` // 必须;  菜单标题, 不超过16个字节, 子菜单不超过40个字节
	Key      string `json:"key,omitempty"`  // 非必须; 菜单KEY值, 用于消息接口推送, 不超过128字节
	URL      string `json:"url,omitempty"`  // 非必须; 网页链接, 用户点击菜单可打开链接, 不超过256字节
	NewsInfo struct {
		List []NewsList `json:"list,omitempty"`
	} `json:"news_info,omitempty"`
	MediaId    string `json:"media_id,omitempty"` // 非必须; 调用新增永久素材接口返回的合法media_id
	SubButtons struct {
		List []SubButton `json:"list,omitempty"`
	} `json:"sub_button,omitempty"` // 非必须; 二级菜单数组, 个数应为1~5个
}

type SubButton struct {
	Type     string `json:"type,omitempty"` // 非必须; 菜单的响应动作类型
	Name     string `json:"name,omitempty"` // 必须;  菜单标题, 不超过16个字节, 子菜单不超过40个字节
	Key      string `json:"key,omitempty"`  // 非必须; 菜单KEY值, 用于消息接口推送, 不超过128字节
	URL      string `json:"url,omitempty"`  // 非必须; 网页链接, 用户点击菜单可打开链接, 不超过256字节
	NewsInfo struct {
		List []NewsList `json:"list,omitempty"`
	} `json:"news_info,omitempty"`
}
