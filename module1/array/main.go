package main

import "fmt"

var 	array = [5]string{"I","am","stupid","and","weak"}


func main() {
	// array := [5]string{"I","am","stupid","and","weak"}

	// array[2] = "smart"
	// array[4] = "strong"




    for i := 0 ; i < len(array) ; i++ {

		if array[i] == "stupid" {
		   array[i] = "smart"}
	
		if array[i] == "weak" {
		   array[i] = "strong" }

		fmt.Println(array[i])
	}
   
}