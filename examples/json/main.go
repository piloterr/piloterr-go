package main

import (
	"context"
	"fmt"
	"log"

	"github.com/piloterr/piloterr-go"
)

func main() {
	client, err := piloterr.NewClient("", nil) // reads PILOTERR_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.AmazonProduct(ctx, &piloterr.AmazonProductParams{Query: "B09DFCB66S"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
