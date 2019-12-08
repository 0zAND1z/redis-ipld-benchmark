package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

// SampleStruct defines the sample payload used in benchmarking
type SampleStruct struct {
	ID    string
	Name  string
	Value string
}

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter value for ID: ")
	scanner.Scan()
	inputID := scanner.Text()

	fmt.Println("Enter value for Name: ")
	scanner.Scan()
	inputName := scanner.Text()

	fmt.Println("Enter value for Value: ")
	scanner.Scan()
	inputValue := scanner.Text()

	json, err := json.Marshal(SampleStruct{inputID, inputName, inputValue})

	// SET OPERATION
	start := time.Now()
	err = client.Set(inputID, json, 0).Err()
	elapsed := time.Since(start)
	log.Printf("REDIS SET OPERATION TOOK %s", elapsed)
	if err != nil {
		fmt.Println(err)
	}
	// GET OPERATION
	start = time.Now()
	val, err := client.Get(inputID).Result()
	elapsed = time.Since(start)
	log.Printf("REDIS GET OPERATION TOOK %s", elapsed)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

}
