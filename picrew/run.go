package picrew

import (
	"math/rand"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	zero.OnFullMatch("今天是什么少女").SetBlock(true).
		SetPriority(-99).Handle(func(ctx *zero.Ctx) {

		ctx.SendChain(message.At(ctx.Event.UserID), message.Text("正在给你打扮中..."))
		var url string
		if rand.Intn(3) == 0 {
			url = ToPic0()
		} else if rand.Intn(2) == 0 {
			url = ToPic1()
		} else {
			url = ToPic2()
		}
		ctx.SendChain(message.At(ctx.Event.UserID), message.Image(url))

	})

}
