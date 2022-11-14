package wbl0

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
)

// структура приложенеия
type App struct {
	dbConn   *pgx.Conn
	natsConn *stan.Conn
	httpSrv  *http.Server
	cache    *CacheManager
}

// конструктор приложения
func CreateApp(dbConn *pgx.Conn, natsConn *stan.Conn, srv *http.Server, cm *CacheManager) *App {
	return &App{
		dbConn,
		natsConn,
		srv,
		cm,
	}
}

func (app *App) Run() {

	sub := channelSubscription(*app.natsConn, app.cache)
	defer sub.Unsubscribe()
	log.Fatal(app.httpSrv.ListenAndServe())

}
