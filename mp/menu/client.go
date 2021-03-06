// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/gogap/wechat for the canonical source repository
// @license     https://github.com/gogap/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

import (
	"net/http"

	"github.com/gogap/wechat/mp"
)

type Client struct {
	*mp.WechatClient
}

// 兼容保留, 建議實際項目全局維護一個 *mp.WechatClient
func NewClient(AccessTokenServer mp.AccessTokenServer, httpClient *http.Client) Client {
	return Client{
		WechatClient: mp.NewWechatClient(AccessTokenServer, httpClient),
	}
}

// 创建自定义菜单.
func (clt Client) CreateMenu(menu Menu) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	if err = clt.PostJSON(incompleteURL, &menu, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除自定义菜单
func (clt Client) DeleteMenu() (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取自定义菜单
func (clt Client) GetMenu() (menu Menu, err error) {
	var result struct {
		mp.Error
		Menu Menu `json:"menu"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = result.Menu
	return
}

// 获取自定义菜单配置接口
func (clt Client) GetCurrentSelfMenuInfo() (info MenuInfo, isMenuOpen bool, err error) {
	var result struct {
		mp.Error
		IsMenuOpen int      `json:"is_menu_open"`
		MenuInfo   MenuInfo `json:"selfmenu_info"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = result.MenuInfo
	if result.IsMenuOpen == 1 {
		isMenuOpen = true
	}
	return
}
