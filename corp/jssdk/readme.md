### 获取 jsapi_ticket 示例
```Go
package main

import (
	"fmt"

	"github.com/gogap/wechat/corp"
	"github.com/gogap/wechat/corp/jssdk"
)

var AccessTokenServer = corp.NewDefaultAccessTokenServer("corpId", "corpSecret", nil)
var TicketServer = jssdk.NewDefaultTicketServer(AccessTokenServer, nil)

func main() {
	fmt.Println(TicketServer.Ticket())
}
```