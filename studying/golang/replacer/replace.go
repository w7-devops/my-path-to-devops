package main
import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)
func main(){
	fmt.Print("enter str: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	str := input

	fmt.Print("enter a (replace: a -> b): ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	a := input

	fmt.Print("enter b: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	b := input

	replacer := strings.NewReplacer(a,b)
	new_str := replacer.Replace(str)
	fmt.Println("\n---\n","result:",new_str)
}
