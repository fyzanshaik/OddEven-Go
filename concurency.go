package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func generateCode1(wg *sync.WaitGroup, start, end int, file *os.File) {
	defer wg.Done()

	var tempCode strings.Builder

	for i := start; i < end; i = i + 2 {
		tempCode.WriteString(fmt.Sprintf(
			`else if(input == %d){
			fmt.Println("Odd");
		}else if(input == %d){
			fmt.Println("Even");
		}`, i, i+1))
	}

	file.Write([]byte(tempCode.String()))
}

func main() {
	filepath := "oddeven.go"

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	// var number int = 100000000 this is the limit 
	var number int = 100000;

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

	var wg sync.WaitGroup
	numGoroutines := 10
	for i := 1; i < number; i += numGoroutines * 2 {
		wg.Add(1)
		go generateCode1(&wg, i, i+numGoroutines*2, file)
	}

	wg.Wait()

	endingCode := `else{
		fmt.Println("randomValue")
	}
}`
	file.Write([]byte(endingCode))

	fmt.Println("Data written to", filepath)
}
