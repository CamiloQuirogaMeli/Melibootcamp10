package transaction

import (
	"fmt"
)

type Repository interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	Store(transactionCode string, money int, emitter string, receiver string, date string) (Transaction, error)
	LastID() (int, error)
}

type repository struct {
}

type Transaction struct {
	Id              int    `json:"id"`
	TransactionCode string `json:"transaction_code"`
	Money           int    `json:"money"`
	Emitter         string `json:"emitter"`
	Receiver        string `json:"receiver"`
	Date            string `json:"date"`
}

var (
	transactionCollection []Transaction
	lastID                int
)

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaction, error) {
	return transactionCollection, nil
}

func (r *repository) GetByID(id int) (Transaction, error) {
	for _, transaction := range transactionCollection {
		if transaction.Id == id {
			return transaction, nil
		}
	}
	return Transaction{}, fmt.Errorf("error: id %d not exists", id)
}

func (r *repository) Store(transactionCode string, money int, emitter string, receiver string, date string) (Transaction, error) {
	if money == 0 || transactionCode == "" || emitter == "" || receiver == "" || date == "" {
		return Transaction{}, fmt.Errorf("error: some param is empty")
	}
	transaction := Transaction{
		Id:              lastID + 1,
		TransactionCode: transactionCode,
		Money:           money,
		Emitter:         emitter,
		Receiver:        receiver,
		Date:            date,
	}
	transactionCollection = append(transactionCollection, transaction)
	return transaction, nil
}

func (r *repository) LastID() (int, error) {
	if lastID < 0 {
		return 0, fmt.Errorf("error: last id is set below 0")
	}
	return lastID, nil
}
