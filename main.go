package main

import (
	"fmt"
	"time"

	"github.com/laike9m/wechat-go/wxweb"
	"github.com/songtianyi/rrframework/logs"

	"github.com/laike9m/WechatShadowClone/receiver"
)

func main() {
	// 创建session, 一个session对应一个机器人
	// 二维码显示在终端上
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	fmt.Println("here")
	if err != nil {
		fmt.Println(err)
		return
	}

	// switcher.Register(session)
	// session.HandlerRegister.EnableByName("switcher")
	receiver.Register(session)
	session.HandlerRegister.EnableByName("receiver")

	// 登录并接收消息
	for {
		if err := session.LoginAndServe(false); err != nil {
			fmt.Printf("session exit, %s", err)
			for i := 0; i < 3; i++ {
				logs.Info("trying re-login with cache")
				if err := session.LoginAndServe(true); err != nil {
					logs.Error("re-login error, %s", err)
				}
				time.Sleep(3 * time.Second)
			}
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				logs.Error("create new sesion failed, %s", err)
				break
			}
		}
	}
}
