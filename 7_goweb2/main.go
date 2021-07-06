package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	usersSlice := Users{}
	getAll("./users.json", &usersSlice)

	productsSlice := Products{}
	getAll("./products.json", &productsSlice)

	transactionsSlice := Transactions{}
	getAll("./transactions.json", &transactionsSlice)

	router.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Nacho!",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		users := Users{}
		edad, isExist := c.GetQuery("edad")

		if err := getAll("./users.json", &users); err != nil {
			c.JSON(500, gin.H{"error": "No se puede procesar esta solicitud"})
		} else {
			if isExist {
				filterSlice := []User{}
				for _, value := range users.Users {
					if e, err := strconv.Atoi(edad); value.Edad > e && err == nil {
						filterSlice = append(filterSlice, value)
					}
				}
				c.JSON(200, filterSlice)
			} else {
				c.JSON(200, users)
			}
		}
	})

	router.GET("/users/:ID", func(c *gin.Context) {
		users := Users{}
		id, isExist := c.Params.Get("ID")
		found := false

		if err := getAll("./users.json", &users); err != nil {
			c.JSON(500, gin.H{"error": "No se puede procesar esta solicitud"})
		} else {
			if isExist {
				for _, user := range users.Users {
					if i, err := strconv.Atoi(id); user.Id == i && err == nil {
						c.JSON(200, user)
						found = true
						break
					}
				}
				if !found {
					c.JSON(404, gin.H{"Error": "No se encontro el recurso que esta solicitando"})
				}
			} else {
				c.JSON(500, gin.H{"Error": "Vuelva a intentarlo mas tarde"})
			}
		}
	})

	router.GET("/products", func(c *gin.Context) {
		products := Products{}
		stock, isExist := c.GetQuery("stock")

		if err := getAll("./products.json", &products); err != nil {
			c.JSON(500, gin.H{"error": "No se puede procesar esta solicitud"})
		} else {
			if isExist {
				filterSlice := []Product{}
				for _, value := range products.Products {
					if e, err := strconv.Atoi(stock); value.Stock > e && err == nil {
						filterSlice = append(filterSlice, value)
					}
				}
				c.JSON(200, filterSlice)
			} else {
				c.JSON(200, products)
			}
		}
	})

	router.GET("/products/:ID", func(c *gin.Context) {
		products := Products{}
		id, isExist := c.Params.Get("ID")
		found := false

		if err := getAll("./products.json", &products); err != nil {
			c.JSON(500, gin.H{"error": "No se puede procesar esta solicitud"})
		} else {
			if isExist {
				for _, product := range products.Products {
					if i, err := strconv.Atoi(id); product.Id == i && err == nil {
						c.JSON(200, product)
						found = true
						break
					}
				}
				if !found {
					c.JSON(404, gin.H{"Error": "No se encontro el recurso que esta solicitando"})
				}
			} else {
				c.JSON(500, gin.H{"Error": "Vuelva a intentarlo mas tarde"})
			}
		}
	})

	router.POST("/productos", func(c *gin.Context) {
		var req Product
		if c.Request.Header.Get("token") == "123456" {
			if err := c.Bind(&req); err != nil {
				fmt.Println("error", err.Error())
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			req.Id = productsSlice.Products[len(productsSlice.Products)-1].Id
			productsSlice.Products = append(productsSlice.Products, req)
			c.JSON(200, req)
		} else {
			c.JSON(401, "no tiene permisos para realizar la petición solicitada")
		}
	})

	router.GET("/transactions", func(c *gin.Context) {
		transactions := Transactions{}
		monto, isExist := c.GetQuery("monto")

		if err := getAll("./transactions.json", &transactions); err != nil {
			c.JSON(500, gin.H{"error": "No se puede procesar esta solicitud"})
		} else {
			if isExist {
				filterSlice := []Transaction{}
				for _, value := range transactions.Transactions {
					if e, err := strconv.Atoi(monto); value.Monto > e && err == nil {
						filterSlice = append(filterSlice, value)
					}
				}
				c.JSON(200, filterSlice)
			} else {
				c.JSON(200, transactions)
			}
		}
	})

	router.GET("/transactions/:ID", func(c *gin.Context) {
		transactions := Transactions{}
		id, isExist := c.Params.Get("ID")
		found := false

		if err := getAll("./transactions.json", &transactions); err != nil {
			c.JSON(500, gin.H{"error": "No se puede procesar esta solicitud"})
		} else {
			if isExist {
				for _, transaction := range transactions.Transactions {
					if i, err := strconv.Atoi(id); transaction.Id == i && err == nil {
						c.JSON(200, transaction)
						found = true
						break
					}
				}
				if !found {
					c.JSON(404, gin.H{"Error": "No se encontro el recurso que esta solicitando"})
				}
			} else {
				c.JSON(500, gin.H{"Error": "Vuelva a intentarlo mas tarde"})
			}
		}
	})

	router.Run()
}

func getAll(route string, s interface{}) error {
	dataSlice, err := ioutil.ReadFile(route)
	if err != nil {
		return err
	} else {
		err = json.Unmarshal(dataSlice, s)
	}
	return err
}
