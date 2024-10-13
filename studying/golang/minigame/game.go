package main
import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("mini game \"guess the number\": you are given a number from 1 to 100 and you have to guess it. You have 10 attempts. Also, you will have hints whether the number is less or more than you guessed.")
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	success := false
	fmt.Println("\n/-----------------------------------/\n")
	for i := 0; i < 10; i++ {
		fmt.Println("\n---\nAttemp #",i,"| you have",10-i,"attemps\n")
		fmt.Print("enter your number: ")
		input, err := reader.ReadString('\n')
		if err!=nil{
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		user_num, err := strconv.Atoi(input)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("\nyour number is",user_num)
		if user_num < target {
			fmt.Println(user_num,"< target")
		} else if user_num > target {
			fmt.Println(user_num,"> target")
		} else {
			fmt.Println("\nyou win :)")
			fmt.Println("target =",target,"| your_num =",user_num)
			success = true
			break
		}
	}
	if !success{
		fmt.Println("you lost :(")
		fmt.Println("target =",target)
	}
}
