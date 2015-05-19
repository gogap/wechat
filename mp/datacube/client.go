// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/gogap/wechat for the canonical source repository
// @license     https://github.com/gogap/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package datacube

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
