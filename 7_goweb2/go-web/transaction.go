package main

import (
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id              int    `json:"id"`
	TransactionCode string `json:"transaction_code"`
	Money           int    `json:"money"`
	Emitter         string `json:"emitter"`
	Receiver        string `json:"receiver"`
	Date            string `json:"date"`
}

func Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto Transaction
		token := ctx.Request.Header.Get("token")
		if token != "123456789" {
			ctx.JSON(401, gin.H{
				"error": "invalid token",
			})
			return
		}
		if err := ctx.Bind(&dto); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if missingField := checkDto(dto); missingField != "" {
			ctx.JSON(400, "The field "+missingField+" is required")
			return
		}
		if len(Transactions) > 0 {
			id := Transactions[len(Transactions)-1].Id + 1
			dto.Id = id
		} else {
			dto.Id = 1
		}
		Transactions = append(Transactions, dto)
		ctx.JSON(200, dto)
	}
}

func checkDto(dto Transaction) (field string) {
	if dto.TransactionCode == "" {
		field = "transaction_code"
	}
	if dto.Emitter == "" {
		field = "emmiter"
	}
	if dto.Receiver == "" {
		field = "receiver"
	}
	if dto.Money == 0 {
		field = "money"
	}
	if dto.Date == "" {
		field = "date"
	}
	return field
}
