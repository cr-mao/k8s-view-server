package k8sadmin

import (
	"net/http"
	"time"

	"github.com/cr-mao/k8s-view-server/app/k8sadmin/routers"
	"github.com/cr-mao/k8s-view-server/infra/conf"
	"github.com/cr-mao/k8s-view-server/infra/console"
	"github.com/cr-mao/k8s-view-server/infra/helpers"
)

func NewServe(port string) *http.Server {
	router := routers.NewRouter()
	addr := conf.GetString("app.http_host") + ":" + port
	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	console.Success(time.Now().Format(helpers.CSTLayout) + "  http listening on " + addr)
	return s
}
