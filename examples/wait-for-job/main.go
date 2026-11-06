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
	submitted, err := client.AmazonProduct(ctx, &piloterr.AmazonProductParams{Query: "B09DFCB66S", IsAsync: piloterr.Bool(true)})
	if err != nil {
		log.Fatal(err)
	}
	var accepted struct {
		JobID string `json:"job_id"`
	}
	if err := submitted.Decode(&accepted); err != nil {
		log.Fatal(err)
	}
	job, err := client.WaitForJob(ctx, accepted.JobID, &piloterr.WaitForJobOptions{
		OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(job.Body))
}
