package handler

import (
	"github.com/admpub/godownloader/service"
	"github.com/webx-top/echo"
	mw "github.com/webx-top/echo/middleware"

	"github.com/admpub/nging/v4/application/library/route"
	dlconfig "github.com/nging-plugins/dlmanager/pkg/library/config"
)

var Server = &service.DServ{}

func RegisterRoute(r *route.Collection) {
	r.Backend.RegisterToGroup(`/download`, registerRoute)
}

func registerRoute(g echo.RouteRegister) {
	Server.Register(g, true)
	g.Route(`GET,POST`, `/file`, File, mw.CORS())
}

var downloadDir = func() string {
	if len(dlconfig.Get().SavePath) == 0 {
		return service.GetDownloadPath()
	}
	return dlconfig.Get().SavePath
}

func init() {
	Server.SetTmpl(`download/index`)
	Server.SetSavePath(downloadDir)
}
