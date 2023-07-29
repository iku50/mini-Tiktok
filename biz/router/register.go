// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	chat "mini-Tiktok/biz/router/chat"
	comment "mini-Tiktok/biz/router/comment"
	favorite "mini-Tiktok/biz/router/favorite"
	feed "mini-Tiktok/biz/router/feed"
	publish "mini-Tiktok/biz/router/publish"
	relation "mini-Tiktok/biz/router/relation"
	user "mini-Tiktok/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user.Register(r)

	relation.Register(r)

	publish.Register(r)

	feed.Register(r)

	comment.Register(r)

	chat.Register(r)

	favorite.Register(r)
}
