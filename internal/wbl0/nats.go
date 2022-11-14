package wbl0

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	pkg "github.com/vladqstrn/l0/pkg/const"
)

var validate = validator.New()

type DeliveryMessageHandler struct {
	cache  *CacheManager
	dbConn *pgx.Conn
}

func CreateDeliveryMessageHandler(cache *CacheManager, dbConn *pgx.Conn) *DeliveryMessageHandler {
	return &DeliveryMessageHandler{
		cache,
		dbConn,
	}
}

func (dmh *DeliveryMessageHandler) Handle(m *stan.Msg) {
	fmt.Println("Received a new order")
	data := Order{}
	json.Unmarshal(m.Data, &data)
	_, ok := dmh.cache.Get(data.Order_uid)
	if !ok {
		validationErr := validate.Struct(data)
		if validationErr != nil {
			log.Print("validation error: ", validationErr)
			return
		}
		dmh.cache.Set(data.Order_uid, data)
		MakeTransaction(&data, dmh.dbConn)

	} else {
		fmt.Println("Order already exsist")
	}
}

func NatsConnection() *stan.Conn {
	sc, _ := stan.Connect(pkg.ClusterID, pkg.ClientID)
	return &sc
}

// channelSubscription - подписка на канал
// по callback`у данные записываются в кеш и выполняется транзакция в бд
func channelSubscription(sc stan.Conn, cache *CacheManager) stan.Subscription {
	msgHandler := CreateDeliveryMessageHandler(cache, cache.dbConn)
	sub, _ := sc.Subscribe(pkg.Channel, msgHandler.Handle)
	return sub
}
