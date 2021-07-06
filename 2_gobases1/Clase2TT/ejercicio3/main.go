package main

import "fmt"

func main() {

	employees := map[string]interface{}{
		"Tomas": map[string]interface{}{
			"salary":     250000.0,
			"age":        32,
			"workingAge": 9,
			"isWorking":  true,
		},
		"Pedro": map[string]interface{}{
			"salary":     150000.0,
			"age":        19,
			"workingAge": 2,
			"isWorking":  true,
		},
		"Gabriela": map[string]interface{}{
			"salary":     77000.0,
			"age":        32,
			"workingAge": 1,
			"isWorking":  true,
		},
		"Florencia": map[string]interface{}{
			"salary":     0.0,
			"age":        19,
			"workingAge": 0,
			"isWorking":  false,
		},
	}

	for employeeName, employeeData := range employees {

		salary := employeeData.(map[string]interface{})["salary"].(float64)
		age := employeeData.(map[string]interface{})["age"].(int)
		workingAge := employeeData.(map[string]interface{})["workingAge"].(int)
		isWorking := employeeData.(map[string]interface{})["isWorking"].(bool)

		validEmployee := true

		if age < 22 {
			fmt.Printf("Employee %s does not have the required age. Employee age %d.\n", employeeName, age)
			validEmployee = false
		}

		if workingAge < 1 {

			if !(isWorking) {
				fmt.Printf("Employee %s is not currently working.\n", employeeName)
			} else {
				fmt.Printf("Employee %s does not have the required working age. Employee working age %d.\n", employeeName, workingAge)
			}
			validEmployee = false
		}

		if validEmployee {
			fmt.Printf("Employee %s can have the loan.\n", employeeName)

			if salary >= 100000 {
				fmt.Println("The loand should have interest.")
			}
		}

		fmt.Printf("\n\n")
	}
}
