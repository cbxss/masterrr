package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

import "github.com/joho/godotenv"

var links []Links


func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	AUTH, exists := os.LookupEnv("AUTH")
	var musicID string
	var choice int
	fmt.Println("Enter the song ID:")
	fmt.Scanln(&musicID)
	if exists {
		put(musicID, AUTH)
		for get(musicID, AUTH) != 100{
			time.Sleep(time.Second*15)
			fmt.Println("Processing...")
		}
		fmt.Println("Enter the profile you would like to download : \n1)I\n2)M\n3)N\n4)G\n5)0\n6)Z\n7)dl all")
		fmt.Scanf("%d", &choice)
		if choice != 7 {
			downloadFile(links[choice-1])
		} else
		{
			dlAll()
		}

	}
	fmt.Println("Enter AUTH code.")
}

func dlAll(){
	for i := 0; i < len(links); i++ {
		downloadFile(links[i])
	}
}