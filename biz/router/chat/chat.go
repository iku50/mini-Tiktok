// Code generated by hertz generator. DO NOT EDIT.

package chat

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	chat "mini-Tiktok/biz/handler/chat"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_message := _douyin.Group("/message", _messageMw()...)
			_message.POST("/action/", append(_postmessageMw(), chat.PostMessage)...)
			_message.GET("/chat/", append(_getchatmessageMw(), chat.GetChatMessage)...)
		}
	}
}
