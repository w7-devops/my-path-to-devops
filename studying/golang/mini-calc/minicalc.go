package main
import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Print("enter a: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	a, err := strconv.ParseFloat(input,64)
	if err!=nil {
		log.Fatal(err)
	}

	fmt.Print("enter b: ")
	input, err = reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	b, err := strconv.ParseFloat(input,64)
	if err!=nil{
		log.Fatal(err)
	}
	// c - operation
	fmt.Print("enter operation ( + | - | * | / ): ")
	input, err = reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	c := input
	if c=="+"{
		fmt.Println("Result a+b =>",a+b)
	} else if c=="-"{
		fmt.Println("Result a-b =>",a-b)
	} else if c=="*"{
		fmt.Println("Result a*b =>",a*b)
	} else if c=="/"{
		fmt.Println("Result a/b =>",a/b)
	} else{
		fmt.Println("there is no such operation")
	}
}
