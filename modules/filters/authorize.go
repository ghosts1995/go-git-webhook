package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init()  {
	Authorize := func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/member/*",beego.BeforeRouter,Authorize);
	beego.InsertFilter("/",beego.BeforeRouter,Authorize);
}
