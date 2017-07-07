package main

import (
	"fmt"

	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/wxweb"

	"github.com/laike9m/WechatShadowClone/receiver"
)

func main() {
	// 创建session, 一个session对应一个机器人
	// 二维码显示在终端上
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		fmt.Println(err)
		return
	}

	switcher.Register(session)
	session.HandlerRegister.EnableByName("switcher")
	receiver.Register(session)
	session.HandlerRegister.EnableByName("receiver")

	// 登录并接收消息
	if err := session.LoginAndServe(false); err != nil {
		fmt.Printf("session exit, %s", err)
	}
}
