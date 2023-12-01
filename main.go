package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-dpw/connection"
	"go-dpw/entities"
	"go-dpw/model"
	"io"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	db := connection.InitMongoDB()

	r.Post("/api/v1/e-voucher/create", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		data, err := io.ReadAll(request.Body)

		if err != nil {
			panic(err)
		}
		var payload model.ModelVoucher
		_ = json.Unmarshal(data, &payload)

		var response entities.VoucherEntity = entities.VoucherEntity{
			Id:      payload.VoucherItem,
			Code:    "982137892173",
			Name:    payload.VoucherItem + " Rp.25.000",
			Expired: "15-10-2023",
		}

		//var data entities.VoucherEntity = entities.VoucherEntity{}
		_, err = db.Collection("vouchers").InsertOne(context.TODO(), response)

		if err != nil {
			panic(err)
		}

		rawRequest, _ := json.Marshal(response)

		writer.Write(rawRequest)
		//if err != nil {
		//	panic(err)
		//}
		//var results []entities.UserEntity
		//
		//if cursor.All(context.Background(), &results); err != nil {
		//	panic(err)
		//}
		//
		//rawRequest, _ := json.Marshal(results)
		//writer.Write(rawRequest)
	})

	http.ListenAndServe(":3000", r)
}
