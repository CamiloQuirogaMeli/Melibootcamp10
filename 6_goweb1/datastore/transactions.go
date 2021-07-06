package datastore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Datastore struct {
	jsonpath string
}

type Transaction struct {
	ID         int     `json:"id,omitempty"`
	Code       string  `json:"code_transaction,omitempty"`
	Currency   string  `json:"currency,omitempty"`
	Mount      float64 `json:"mount,omitempty"`
	Dispatcher string  `json:"dispatcher,omitempty"`
	Receiver   string  `json:"receiver,omitempty"`
	Date       string  `json:"transaction_date,omitempty"`
}

func (d *Datastore) GetAll() *[]Transaction {
	jsonFile, err := ioutil.ReadFile(d.jsonpath)
	if err != nil {
		fmt.Println(err)
	}

	transactions := []Transaction{}

	json.Unmarshal(jsonFile, &transactions)

	return &transactions
}

func (d *Datastore) GetByID(id string) Transaction {
	jsonFile, err := ioutil.ReadFile(d.jsonpath)
	if err != nil {
		fmt.Println(err)
	}
	transactions := []Transaction{}
	json.Unmarshal(jsonFile, &transactions)
	transaction := d.filterByID(transactions, id)
	return transaction
}

func (d *Datastore) GetByFilter(id, mount, code, currency, dispatcher, receiver, date string) []Transaction {
	jsonFile, err := ioutil.ReadFile(d.jsonpath)
	if err != nil {
		fmt.Println(err)
	}
	transactions := []Transaction{}
	json.Unmarshal(jsonFile, &transactions)
	transactions = d.filter(transactions, id, mount, code, currency, dispatcher, receiver, date)
	return transactions
}

func (d *Datastore) filterByID(transactions []Transaction, id string) Transaction {
	tx := Transaction{}
	var idtmp int
	if id != "" {
		idtmp, _ = strconv.Atoi(id)
	}
	for _, t := range transactions {
		if idtmp != t.ID {
			continue
		}
		tx = t
	}
	return tx
}

func (d *Datastore) filter(transactions []Transaction, id, mount, code, currency, dispatcher, receiver, date string) []Transaction {
	tx := []Transaction{}
	for _, t := range transactions {
		if id != "" {
			idtmp, _ := strconv.Atoi(id)
			if t.ID != idtmp {
				continue
			}
		}
		if mount != "" {
			mountmp, _ := strconv.ParseFloat(mount, 64)
			if t.Mount != mountmp {
				continue
			}
		}
		if code != "" {
			if t.Code != code {
				continue
			}
		}
		if currency != "" {
			if t.Currency != currency {
				continue
			}
		}
		if dispatcher != "" {
			if t.Dispatcher != dispatcher {
				continue
			}
		}
		if receiver != "" {
			if t.Receiver != receiver {
				continue
			}
		}
		if date != "" {
			if t.Date != date {
				continue
			}
		}
		tx = append(tx, t)
	}
	return tx
}

func NewDS(jsonPath string) *Datastore {
	return &Datastore{
		jsonpath: jsonPath,
	}
}
