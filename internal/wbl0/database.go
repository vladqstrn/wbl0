package wbl0

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

func StartDatabaseConnection() *pgx.Conn {
	db := viper.GetString("DBserver.db")
	port := viper.GetString("DBserver.port")
	dbName := viper.GetString("DBserver.name")
	host := viper.GetString("DBserver.host")
	ssl := viper.GetString("DBserver.ssl")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")

	// connection URL
	url := db + "://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName + "?sslmode=" + ssl

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

// GetOrderStatement возвращает утверждение для записи в таблицу order
func GetOrderStatement() string {
	sqlStatement := "INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	return sqlStatement
}

// GetDeliveryStatement возвращает утверждение для записи в таблицу delivery
func GetDeliveryStatement() string {
	sqlStatement := "INSERT INTO delivery (name, phone, zip, city, address, region, email, order_uid) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	return sqlStatement
}

// GetPaymentsStatement возвращает утверждение для записи в таблицу payments
func GetPaymentsStatement() string {
	sqlStatement := "INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	return sqlStatement
}

// GetItemStatement возвращает утверждение для записи в таблицу payments
func GetItemStatement() string {
	sqlStatement := "INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"
	return sqlStatement
}

// MakeTransaction выполняет запись данных в бд атомарно
func MakeTransaction(data *Order, dbConn *pgx.Conn) {
	tx, err := dbConn.Begin(context.Background())
	if err != nil {
		log.Fatal("Start transaction error: ", err.Error())
	}
	defer tx.Rollback(context.Background())

	orderStatement := GetOrderStatement()
	_, err = tx.Exec(context.Background(), orderStatement, data.Order_uid, data.Track_number, data.Entry, data.Locale, data.Internal_signature, data.Customer_id, data.Delivery_service, data.Shardkey, data.Sm_id, data.Date_created, data.Oof_shard)
	if err != nil {
		log.Fatal()
	}

	deliveryStatement := GetDeliveryStatement()
	_, err = tx.Exec(context.Background(), deliveryStatement, data.Delivery.Name, data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City, data.Delivery.Address, data.Delivery.Region, data.Delivery.Email, data.Order_uid)
	if err != nil {
		log.Print("Exec error", err.Error())
	}

	paymentStatement := GetPaymentsStatement()
	_, err = tx.Exec(context.Background(), paymentStatement, data.Payment.Transaction, data.Payment.Request_id, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount, data.Payment.Payment_dt, data.Payment.Bank, data.Payment.Delivery_cost, data.Payment.Goods_total, data.Payment.Custom_fee, data.Order_uid)
	if err != nil {
		log.Fatal()
	}

	itemStatement := GetItemStatement()
	for i := range data.Item {
		_, err = tx.Exec(context.Background(), itemStatement, data.Item[i].Chrt_id, data.Item[i].Track_number, data.Item[i].Price, data.Item[i].Rid, data.Item[i].Name, data.Item[i].Sale, data.Item[i].Size, data.Item[i].Total_price, data.Item[i].Nm_id, data.Item[i].Brand, data.Item[i].Status, data.Order_uid)
		if err != nil {
			log.Fatal()
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("Commit transaction error: ", err.Error())
	}
}

func PushCache(dbConn *pgx.Conn, key string, property interface{}) {
	sqlStatement := "INSERT INTO cache (order_uid, property) VALUES ($1,$2)"
	_, err := dbConn.Exec(context.Background(), sqlStatement, key, property)
	if err != nil {
		log.Println("unable to send object: ", err.Error())
	}

}

func GetCache(dbConn *pgx.Conn) map[string]interface{} {
	sqlStatement := "SELECT property from cache"
	rows, err := dbConn.Query(context.Background(), sqlStatement)
	if err != nil {
		log.Fatal("The request failed: ", err.Error())
	}
	cache := make(map[string]interface{})
	model := Order{}
	var property string

	for rows.Next() {
		err = rows.Scan(&property)
		if err != nil {
			log.Fatal("scan error: ", err, err.Error())
		}
		json.Unmarshal([]byte(property), &model)
		cache[model.Order_uid] = model
	}
	return cache

}
