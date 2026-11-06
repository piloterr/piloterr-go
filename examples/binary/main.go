package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/piloterr/piloterr-go"
)

func main() {
	client, err := piloterr.NewClient("", nil) // reads PILOTERR_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.WebsiteScreenshot(ctx, &piloterr.WebsiteScreenshotParams{URL: "https://example.com"})
	if err != nil {
		log.Fatal(err)
	}
	if response.IsJSON {
		fmt.Println(string(response.Body))
		return
	}
	if err := os.WriteFile("screenshot.png", response.Body, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Saved screenshot.png (%s, %d bytes)\n", response.ContentType, len(response.Body))
}
