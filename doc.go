// Package piloterr is the official Go SDK for the Piloterr API (https://www.piloterr.com).
//
// Scrape 300+ websites (Amazon, LinkedIn, Google, Zillow, Crunchbase…) as structured JSON with one API.
//
// Create one client and reuse it:
//
//	client, err := piloterr.NewClient("", nil) // reads PILOTERR_API_KEY
//	response, err := client.AmazonProduct(ctx, ...)
//
// Every documented endpoint has a typed method; Client.Request reaches any other path.
package piloterr
