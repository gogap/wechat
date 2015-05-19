// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/gogap/wechat for the canonical source repository
// @license     https://github.com/gogap/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mmpaymkttransfers

import (
	"github.com/gogap/wechat/mch"
)

// 查询代金券批次信息.
func QueryCouponStock(proxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return proxy.PostXML("https://api.mch.weixin.qq.com/mmpaymkttransfers/query_coupon_stock", req)
}
