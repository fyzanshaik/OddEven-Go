package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filepath := "joke.go"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	defer file.Close()

	code := `package main
import (
	"fmt"
)
func oddEven(){
	var input int64
	fmt.Print("Enter a number to check: ")
	_, err := fmt.Scanf("%d", &input)
	
	if err != nil {
		fmt.Printf("Some error")
	}
	
	if input == 0 {
		fmt.Println("Even")
	}`
	file.Write([]byte(code))

	var tempCode strings.Builder

	for i := 1; i < 1000; i = i + 2 {
		tempCode.WriteString(fmt.Sprintf(
			`else if(input == %d){
			fmt.Println("Odd");
		}else if(input == %d){
			fmt.Println("Even");
		}`, i, i+1))
	}

	endingCode := `else {
		fmt.Println("randomValue")
	}
	}`

	file.Write([]byte(tempCode.String()))
	file.Write([]byte(endingCode))

	fmt.Println("Data written to", filepath)
}
