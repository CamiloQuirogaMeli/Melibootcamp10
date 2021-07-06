package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func customerExists(id int, name string, dni string, phone string, address string) (bool, error) {

	defer func() {
		err := recover()

		fmt.Println(err)
	}()

	_, dataErr := dataValidate(id)

	if dataErr != nil {
		return false, dataErr
	}

	_, dataErr = dataValidate(name)

	if dataErr != nil {
		return false, dataErr
	}

	_, dataErr = dataValidate(dni)

	if dataErr != nil {
		return false, dataErr
	}

	_, dataErr = dataValidate(phone)

	if dataErr != nil {
		return false, dataErr
	}

	_, dataErr = dataValidate(address)

	if dataErr != nil {
		return false, dataErr
	}

	_, err := ioutil.ReadFile("./customers.txt")

	if err != nil {
		panic("the indicated file was not found or is corrupted")
	} else {
		//Search of customer simulated
		exists := true
		if exists {
			return true, nil
		} else {
			return false, nil
		}
	}

}

func dataValidate(data interface{}) (bool, error) {
	switch data.(type) {
	case int:
		if data == 0 {
			return false, errors.New("data error")
		}
	case string:
		if data == "" {
			return false, errors.New("data error")
		}
	}
	return true, nil
}

func main() {

	exists, err := customerExists(123, "Jhon Doe", "11222333", "123123123", "street 123")

	if err == nil {
		if exists {
			fmt.Println("Customer found")
		} else {
			fmt.Println("Customer not found")
		}
	}

	fmt.Println("Execution finished")
	fmt.Println("Several errors were detected in time of execution")
	fmt.Println("There are no files open")

}
