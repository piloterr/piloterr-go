# Piloterr Go SDK

The official Go SDK for the [Piloterr](https://www.piloterr.com) web scraping API.

Scrape 300+ websites (Amazon, LinkedIn, Google, Zillow, Crunchbase…) as structured JSON with one API.

The SDK exposes **392 typed methods**, one per documented endpoint, plus a generic `Request()`. It has no dependencies outside the standard library.

## Requirements

- Go 1.22 or newer

## Installation

```sh
go get github.com/piloterr/piloterr-go
```

## Get your API key

[Sign up for Piloterr](https://app.piloterr.com/register) or sign in, then copy your API key from the [dashboard](https://app.piloterr.com). `NewClient("", nil)` reads the `PILOTERR_API_KEY` environment variable.

## Quick start

```go
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
	data, err := response.Map()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
```

Parameter structs mirror the [API reference](https://app.piloterr.com/openapi). Required fields are plain values; optional fields are pointers — use `piloterr.String`, `piloterr.Float`, `piloterr.Int` and `piloterr.Bool`.

Decode into your own types with `response.Decode(&value)`.

## Binary responses

Endpoints that return a file (image, PDF…) give you the bytes in `Body`; `IsJSON` tells both cases apart:

```go
response, err := client.WebsiteScreenshot(ctx, &piloterr.WebsiteScreenshotParams{URL: "https://example.com"})
if err != nil {
	log.Fatal(err)
}
if !response.IsJSON {
	err = os.WriteFile("screenshot.png", response.Body, 0o644)
}
```

## Async jobs

Any call accepts `is_async: true` (useful past the 60-second server limit): the response is then a `job_id`. `WaitForJob` polls `GET /v1/jobs/{id}` until a terminal status (`CompletedJob`, `Error`, `NotFound`, `BadRequest`):

```go
submitted, err := client.AmazonProduct(ctx, &piloterr.AmazonProductParams{Query: "B09DFCB66S", IsAsync: piloterr.Bool(true)})
if err != nil {
	log.Fatal(err)
}
var accepted struct{ JobID string `json:"job_id"` }
if err := submitted.Decode(&accepted); err != nil {
	log.Fatal(err)
}
job, err := client.WaitForJob(ctx, accepted.JobID, &piloterr.WaitForJobOptions{
	OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
})
```

## GET or POST

Methods use GET when the endpoint supports it (parameters in the query string); `WithMethod("POST")` sends them as a JSON body.

```go
response, err := client.AboutyouProduct(ctx, &piloterr.AboutyouProductParams{Query: "test"}, piloterr.WithMethod("POST"))
```

## Any endpoint

```go
response, err := client.Request(ctx, "GET", "/v1/usage", nil)
```

## Errors

```go
response, err := client.AmazonProduct(ctx, &piloterr.AmazonProductParams{Query: "B09DFCB66S"})
if apiError, ok := piloterr.IsAPIError(err); ok {
	log.Printf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
}
```

| Error | When |
| --- | --- |
| `*ConfigurationError` | Missing API key, invalid options or unsupported method. |
| `*SerializationError` | A required parameter is missing or a value can't be sent. |
| `*TransportError` | Network failure, timeout or cancelled context (`errors.Is(err, context.DeadlineExceeded)`). |
| `*APIError` | Non-2xx status (`StatusCode`, `Message`, `Response`). |
| `*ResponseDecodingError` | A success response could not be decoded. |

## Retries

Requests are retried up to 3 times after a network error, a `408`, `500`, `502`, `503` or `504`, or a rate limit — which the API returns as `401` (or `429`) with a `Rate limit exceeded` message. Per-key quota errors (`…: quota daily`) are not retried, nor are `400`, `402`, `404` or client-side timeouts. The delay starts at 1 second and grows 1.5× per retry (capped at 30 seconds); a `Retry-After` header wins.

```go
client, err := piloterr.NewClient("", &piloterr.ClientOptions{
	MaxRetries: 5,                      // negative disables retries
	RetryDelay: 500 * time.Millisecond,
})

// or for one request
response, err := client.AmazonProduct(ctx, &piloterr.AmazonProductParams{Query: "B09DFCB66S"}, piloterr.WithMaxRetries(0))
```

## Batch requests

`piloterr.Batch` runs a function per item with at most 10 calls in flight and returns one `BatchResult` per item, in input order. A failing item never stops the others.

```go
inputs := []string{"B09DFCB66S", "B0BSHF7WHW", "B08N5WRWNW"}
results := piloterr.Batch(ctx, inputs, 5, func(ctx context.Context, value string) (*piloterr.Response, error) {
	return client.AmazonProduct(ctx, &piloterr.AmazonProductParams{Query: value})
})
for _, result := range results {
	if result.Err != nil {
		log.Printf("%s: %v", inputs[result.Index], result.Err)
		continue
	}
	fmt.Println(string(result.Value.Body))
}
```

| Constant | Default | Description |
| --- | --- | --- |
| `DefaultConcurrency` | `10` | Tasks run at the same time by batch |
| `DefaultTimeout` | `120 * time.Second` | Request timeout |
| `MaxRetries` | `3` | Retry attempts after a network error, a 408/5xx or a rate limit |
| `RetryBackoffFactor` | `1.5` | Exponential backoff multiplier |
| `DefaultRetryDelay` | `time.Second` | Delay before the first retry |

## Timeouts and custom HTTP clients

```go
client, err := piloterr.NewClient("", &piloterr.ClientOptions{
	Timeout:    60 * time.Second, // default 120 s; negative disables
	HTTPClient: &http.Client{Transport: myTransport},
})
```

Every method also honours the deadline and cancellation of the `context.Context` you pass.

## Endpoints

392 methods, generated from the [OpenAPI document](https://app.piloterr.com/openapi).

| Method | HTTP | Description |
| --- | --- | --- |
| `AboutyouProduct` | GET / POST `/v2/aboutyou/product` | Read one About You product from a product URL or a keyword. A keyword opens the first search hit. Set region when the query does not alread… |
| `AboutyouSearch` | GET / POST `/v2/aboutyou/search` | Search an About You shop with a keyword or a search URL. Pick the country with region, or let a full link choose the shop. Cards include br… |
| `ActionProduct` | GET / POST `/v2/action/product` | Action Product API: fetch full product page data from a product ID or URL on Action.com. Returns pricing, specs, images, breadcrumbs and av… |
| `ActionSearch` | GET / POST `/v2/action/search` | Action Search API: search product listings by keyword or URL on Action.com across European locales. Returns prices, images, locale, currenc… |
| `AldiProduct` | GET / POST `/v2/aldi/product` | Aldi Product API: fetch full product page data from an Aldi URL or path slug. Returns pricing, images, description, category and cooling fl… |
| `AldiSearch` | GET / POST `/v2/aldi/search` | Aldi Search API: list products from an Aldi category page URL or path across FR, DE, NL and Canarias shops. Returns prices, thumbnails and… |
| `AlibabaProduct` | GET / POST `/v2/alibaba/product` | Alibaba Product API. Fetch a full B2B product page by URL or product ID with optional locale subdomain. Browser rendering returns tiered pr… |
| `AlibabaSearch` | GET / POST `/v2/alibaba/search` | Alibaba Search API. Query B2B product listings by keyword or search URL across 17 locale subdomains. Request mode returns prices, MOQ, sell… |
| `AliexpressProduct` | GET / POST `/v2/aliexpress/product` | AliExpress Product API: scrape live /item/ product pages from a full URL. Get title, images, seller identity, pricing, specs, delivery, and… |
| `AliexpressProductReviews` | GET / POST `/v2/aliexpress/product/reviews` | AliExpress Reviews API: list buyer reviews for a product id or item URL. Each review includes author, star rating, text, date, and photos w… |
| `AliexpressSearch` | GET / POST `/v2/aliexpress/search` | AliExpress Search API: search by keyword or wholesale URL and get product listings with prices, ratings, sold counts, and pagination as JSO… |
| `AliexpressStoreProduct` | GET / POST `/v2/aliexpress/store/product` | AliExpress Store Product API: list products from an AliExpress seller store by URL or store ID. Returns store metadata and product cards (~… |
| `AllegroProduct` | GET / POST `/v2/allegro/product` | Allegro Product API. Extract full product page data from allegro.pl by URL (/produkt/, /oferta/, /ogloszenie/). Browser rendering, 3 credit… |
| `AllegroSearch` | GET / POST `/v2/allegro/search` | Allegro Search API. Query product listings by keyword, optional page number, or full allegro.pl listing URL. Returns results with prices, s… |
| `AmazonProduct` | GET / POST `/v2/amazon/product` | Efficiently extract product details from Amazon for dynamic retail insights, including price, stock, features, and seller data. |
| `AmazonProductOffer` | GET / POST `/v2/amazon/product/offer` | Efficiently extract details of third-party seller offers for any Amazon product across multiple marketplaces. |
| `AmazonSearch` | GET / POST `/v2/amazon/search` | Efficiently scrape Amazon search results for real-time product data and insights across multiple Amazon marketplaces. |
| `ApartmentsProperty` | GET / POST `/v2/apartments/property` | Extract a full Apartments.com rental listing from a property URL, including amenities, gallery, description, geolocation, and unit details… |
| `ApartmentsSearch` | GET / POST `/v2/apartments/search` | Search Apartments.com rentals from a full city or browse URL and retrieve prices, photos, addresses, phone numbers, and pagination via brow… |
| `ArgenpropProperty` | GET / POST `/v2/argenprop/property` | Read one Argenprop listing from its URL or id. You get price, address, rooms, expenses, and the agent when the page shows them. |
| `ArgenpropSearch` | GET / POST `/v2/argenprop/search` | Search Argenprop by neighbourhood or city. Set operation to venta or alquiler, and property_type such as departamentos or casas. Page with… |
| `ArgosProduct` | GET / POST `/v2/argos/product` | Argos Product API: one UK product by id or URL. Browser rendering (2 credits). |
| `ArgosSearch` | GET / POST `/v2/argos/search` | Search the Argos UK catalogue by keyword and page the results. Prices are in pounds. A one-word query that Argos turns into a brand page wi… |
| `ArxivPaper` | GET / POST `/v2/arxiv/paper` | Fetch one arXiv paper by id or by an abs, pdf, or html URL. |
| `ArxivSearch` | GET / POST `/v2/arxiv/search` | Search arXiv papers by keyword, author query such as au:einstein, or arXiv id. A paper URL returns 400. Page through results when the match… |
| `AuchanProduct` | GET / POST `/v2/auchan/product` | Effortlessly extract product details from Auchan for dynamic business insights, including nutritional information, ingredients, ratings, an… |
| `AuchanSearch` | GET / POST `/v2/auchan/search` | Extract Auchan product listings from keyword searches for insightful market analysis and competitive edge. |
| `AutodocProduct` | GET / POST `/v2/autodoc/product` | Scrape Autodoc product pages: price, specs, EAN, images, compatibility, and TecDoc-related vehicle data via the Related endpoint. |
| `AutodocProductRelated` | GET / POST `/v2/autodoc/product/related` | Vehicle fitment lines per maker for an Autodoc product. All parameters from Product fetch.params are required. |
| `AutodocSearch` | GET / POST `/v2/autodoc/search` | Search Autodoc by URL: keyword pages use /search?keyword=… (e.g. frein on auto-doc.fr). Extract product cards, prices, and pagination. |
| `AutoeuropeResult` | GET / POST `/v2/autoeurope/result` | Read one vehicle from an Auto Europe quote. Pass the offer id from Search, with the same pickup and dates. You get the supplier, the SIPP c… |
| `AutoeuropeSearch` | GET / POST `/v2/autoeurope/search` | Quote a rental car on Auto Europe. Pass pickup, pickup date, and dropoff date. Region chooses the country site and its currency. The whole… |
| `Autoscout24Ad` | GET / POST `/v2/autoscout24/ad` | Read one AutoScout24 listing, including an inactive ad that still exposes the listing. Price, vehicle, seller, and location. (1 credit) |
| `Autoscout24Search` | GET / POST `/v2/autoscout24/search` | Search AutoScout24 listings from a search URL. Each card has make, model, numeric price, mileage, and the listing URL. (1 credit) |
| `AutotraderAd` | GET / POST `/v2/autotrader/ad` | Extract full vehicle details from an AutoTrader UK car-details URL, including price, specs, gallery, seller, and description via browser re… |
| `AutotraderSearch` | GET / POST `/v2/autotrader/search` | Search AutoTrader UK listings from a car-search URL and retrieve prices, specs, photos, seller info, and pagination via browser rendering. |
| `BestbuyProduct` | GET / POST `/v2/bestbuy/product` | Get a Best Buy US or Canada product page: title, price, brand, stock, images, categories and specs via browser rendering. |
| `BestbuyProductReviews` | GET / POST `/v2/bestbuy/product/reviews` | Get customer reviews for a Best Buy product URL: rating, review text, author, photos and pagination. |
| `BestbuySearch` | GET / POST `/v2/bestbuy/search` | Search Best Buy US or Canada by keyword or search URL. Returns product listings with price, rating and pagination via browser rendering. |
| `BhphotovideoProduct` | GET / POST `/v2/bhphotovideo/product` | Read a B&H Photo Video product from its URL or id. Returns price, rating, specifications, and box contents. |
| `BhphotovideoSearch` | GET / POST `/v2/bhphotovideo/search` | Search B&H Photo Video by keyword or search URL. Returns BH numbers, dollar prices, and a next page URL. |
| `BingSearch` | GET / POST `/v2/bing/search` | Scrape real-time Bing search results with full control over location, language, and pagination. |
| `BitibaProduct` | GET / POST `/v2/bitiba/product` | Read one product from a live Bitiba shop, including the selected variant price and the pack list. |
| `BitibaSearch` | GET / POST `/v2/bitiba/search` | Search a live Bitiba country shop. Germany is the default. www.bitiba.com is closed and returns 400. |
| `BolProduct` | GET / POST `/v2/bol/product` | Read one bol.com product by URL or product id, in nl, be-nl, or be-fr. You get the selling price, the list price, and the seller. |
| `BolSearch` | GET / POST `/v2/bol/search` | Search bol.com by keyword. Region is nl, be-nl, or be-fr. Page the product cards. Sponsored ads can repeat across pages. |
| `BoulangerProduct` | GET / POST `/v2/boulanger/product` | Read one Boulanger product from its URL or product id. You get price, brand, rating, seller, and specifications. |
| `BoulangerSearch` | GET / POST `/v2/boulanger/search` | Search Boulanger by keyword and page the product grid. Cards include price, brand, seller, and the product URL. |
| `BraveSearch` | GET / POST `/v2/brave/search` | Access the Brave search engine results directly through our efficient Piloterr service for reliable data scraping. |
| `BricodepotProduct` | GET / POST `/v2/bricodepot/product` | Read one Brico Dépôt product from its product URL. You get price, brand, SKU, and availability. |
| `BricodepotSearch` | GET / POST `/v2/bricodepot/search` | Search Brico Dépôt with a /produits/ listing URL, not a keyword box. Page 2 is the same path with /2 at the end. |
| `BricoramaProduct` | GET / POST `/v2/bricorama/product` | Read one Bricorama product from its URL. You get price, brand, and the product details. |
| `BricoramaSearch` | GET / POST `/v2/bricorama/search` | Search Bricorama by keyword or search URL and page the home-improvement grid. |
| `CarrefourProduct` | GET / POST `/v2/carrefour/product` | Read one Carrefour.fr product from its page URL. Every product returns the same keys: price, brand, pack size, Nutriscore, categories, unit… |
| `CarrefourSearch` | GET / POST `/v2/carrefour/search` | Search for products on Carrefour by keyword or URL, returning structured listings with prices, ratings, and pagination. |
| `CarrefourSuggest` | GET / POST `/v2/carrefour/suggest` | Get product suggestions from Carrefour by keyword or EAN, returning matching products with pricing, availability, and product details. |
| `CastoramaProduct` | GET / POST `/v2/castorama/product` | Read one Castorama product from its URL. You get price, brand, and specifications. |
| `CastoramaSearch` | GET / POST `/v2/castorama/search` | Search Castorama by keyword or search URL and page the DIY listings. Cards include title, price, and product link. |
| `CdiscountProduct` | GET / POST `/v2/cdiscount/product` | Extract full product details from a Cdiscount product page, including price, seller, ratings, specifications, and offers. |
| `CdiscountProductOffer` | GET / POST `/v2/cdiscount/product/offer` | Retrieve all marketplace seller offers for a Cdiscount product, including prices, seller ratings, shipping costs, and availability. |
| `CdiscountSearch` | GET / POST `/v2/cdiscount/search` | Extract product listings from Cdiscount search results or category pages for market intelligence and competitive analysis. |
| `CentrisProperty` | GET / POST `/v2/centris/property` | Read one Centris property from its URL or listing id. You get price, address, rooms, and the broker. |
| `CentrisSearch` | GET / POST `/v2/centris/search` | Search Centris listings by city or by a full results URL. Use the English or French path. Page through properties for sale. |
| `ChewyProduct` | GET / POST `/v2/chewy/product` | Read one Chewy product by URL or numeric id. The default mode is expert and returns the historical product cache. simple returns one snake_… |
| `ChewySearch` | GET / POST `/v2/chewy/search` | Search Chewy by keyword or listing URL. The default mode is expert and returns the historical listing objects. simple returns short snake_c… |
| `ClearedjobsJob` | GET / POST `/v2/clearedjobs/job` | Fetch one cleared job by its URL or by a slug that ends in the numeric id. |
| `ClearedjobsSearch` | GET / POST `/v2/clearedjobs/search` | Search cleared jobs by keyword. Only the first HTML page is available, so pagination.next is null and page 2 returns 404. |
| `Company` | GET / POST `/v2/company` | Retrieve comprehensive company data including financials, technologies, social profiles, and location by domain name. |
| `CoolblueProduct` | GET / POST `/v2/coolblue/product` | Read one Coolblue product by article id or product URL, in the region of that shop. |
| `CoolblueSearch` | GET / POST `/v2/coolblue/search` | Search Coolblue with a product keyword. Region defaults to Belgium. Also use nl or de. Page through product cards. |
| `CopartLot` | GET / POST `/v2/copart/lot` | Copart Lot API: retrieve structured auction lot data by lot number or URL. Specs, damage, bids, photos, yard location and sale status acros… |
| `CopartSearch` | GET / POST `/v2/copart/search` | Copart Search API: search auction lots by keyword or URL across Copart regional sites. Returns lot listings and pagination. |
| `CostcoProduct` | GET / POST `/v2/costco/product` | Costco Product API: fetch full product page data from a product ID or URL across regional Costco sites. Returns pricing, availability, rati… |
| `CostcoSearch` | GET / POST `/v2/costco/search` | Costco Search API: search product listings by keyword or URL across 13 regional Costco sites. Returns prices, availability, review counts a… |
| `CrunchbaseCompanyInfo` | GET / POST `/v2/crunchbase/company/info` | Convert URLs into detailed Crunchbase profiles for immediate company insights. |
| `CrunchbaseEvent` | GET / POST `/v2/crunchbase/event` | Retrieve detailed Crunchbase event profiles with speakers, sponsors, exhibitors, and press timelines. |
| `CrunchbaseFundingRound` | GET / POST `/v2/crunchbase/funding_round` | Retrieve full details of a specific Crunchbase funding round including investors and press timeline. |
| `CrunchbaseFundingRounds` | GET / POST `/v2/crunchbase/funding_rounds` | Monitor funding history and financial milestones of companies with Crunchbase data. |
| `CrunchbasePeopleInfo` | GET / POST `/v2/crunchbase/people/info` | Retrieve detailed Crunchbase profiles for individuals including career history, investments, and board positions. |
| `CrunchbaseSearch` | GET / POST `/v2/crunchbase/search` | Search Crunchbase for companies and people by keyword and retrieve their profiles and Crunchbase identifiers. |
| `CulturaProduct` | GET / POST `/v2/cultura/product` | Read one Cultura product from its URL or SKU. You get price, EAN, seller, condition, and stock wording. |
| `CulturaSearch` | GET / POST `/v2/cultura/search` | Search Cultura by keyword and page through books, games, and culture products. Cards include title, price, availability, and the product li… |
| `DartyProduct` | GET / POST `/v2/darty/product` | Read one Darty product sheet from a product URL or a Darty product id. You get price, brand, availability, rating, and specifications. |
| `DartySearch` | GET / POST `/v2/darty/search` | Search Darty by keyword or search URL and page through the listing. Each card has title, price, brand, seller, and product link. |
| `DatasetsLinkedinCompany` | POST `/v2/datasets/linkedin/company` | Query the LinkedIn company dataset with Elasticsearch Query DSL. Filter by name, domain, industry, headcount, and country. POST-only, 1 cre… |
| `DecathlonProduct` | GET / POST `/v2/decathlon/product` | Read one Decathlon product from a product URL, a product path, or a keyword. A keyword uses the first search hit. Set region when the query… |
| `DecathlonSearch` | GET / POST `/v2/decathlon/search` | Search a Decathlon shop with a keyword or a search URL. Choose the country with region, or let a full link pick the shop, then page through… |
| `DigikeyProduct` | GET / POST `/v2/digikey/product` | Get a DigiKey product page: part numbers, manufacturer, price tiers, stock, images, datasheet and specs via browser rendering. |
| `DigikeySearch` | GET / POST `/v2/digikey/search` | Search DigiKey country sites by keyword or result/filter URL. Returns category branches or product listings with pagination via browser ren… |
| `DoajArticle` | GET / POST `/v2/doaj/article` | Fetch one DOAJ article by its id or by a doaj.org article URL. |
| `DoajJournal` | GET / POST `/v2/doaj/journal` | Fetch one DOAJ journal by ISSN, journal id, or a doaj.org toc URL. |
| `DoajSearch` | GET / POST `/v2/doaj/search` | Search DOAJ articles or journals. kind is article (default) or journal. Article ids and journal ISSNs are the detail keys. |
| `DomainDnsbl` | GET / POST `/v2/domain/dnsbl` | Check if a domain or IP address is listed in DNS-based blacklists (DNSBL), with detailed provider-level detection and category information. |
| `DomainMalicious` | GET / POST `/v2/domain/malicious` | Check if a domain or IP address has been flagged as malicious across multiple threat intelligence feeds, with source attribution. |
| `DomainWhois` | GET / POST `/v2/domain/whois` | Retrieve WHOIS registration data for any domain, including registration dates, registrar details, name servers, and contact information. |
| `DomaincomauProperty` | GET / POST `/v2/domaincomau/property` | Read one Domain.com.au listing from its URL. The response is a flat object with price, address, agency, and photos. |
| `DomaincomauSearch` | GET / POST `/v2/domaincomau/search` | Search Domain.com.au by suburb. mode is sale, rent, or sold. page starts at 1. Each card has a dollar price and a price label rebuilt from… |
| `DuckduckgoSearch` | GET / POST `/v2/duckduckgo/search` | Query DuckDuckGo via API and retrieve structured organic results with region and language. |
| `EbayProduct` | GET / POST `/v2/ebay/product` | eBay Product API: one listing on a supported eBay site, by URL or item id (3 credits). |
| `EbaySearch` | GET / POST `/v2/ebay/search` | eBay Search API: search listings on supported eBay sites. Returns price, seller, and pagination (3 credits). |
| `ElcorteinglesProduct` | GET / POST `/v2/elcorteingles/product` | El Corte Inglés Product API. Fetch a full product page from elcorteingles.es by URL or product id. Returns pricing, stock, images and ident… |
| `ElcorteinglesSearch` | GET / POST `/v2/elcorteingles/search` | El Corte Inglés Search API. Search product listings on elcorteingles.es by keyword or URL. Returns prices, brands, images and pagination vi… |
| `EleclercProduct` | GET / POST `/v2/eleclerc/product` | Read one E.Leclerc product by EAN or product URL. Pass store_id to get that store's price and stock. Another store returns a different pric… |
| `EleclercSearch` | GET / POST `/v2/eleclerc/search` | Search for products on E.Leclerc by keyword with optional store-specific filtering and pagination. |
| `EleclercSuggest` | GET / POST `/v2/eleclerc/suggest` | Get instant product autocomplete suggestions from E.Leclerc for a given keyword or EAN. |
| `EmailAnalyzes` | GET / POST `/v2/email/analyzes` | Analyze a domain to retrieve its email infrastructure. Microsoft 365 and Google Workspace domains, SPF records, and DNS configurations. |
| `EmailFinder` | GET / POST `/v2/email/finder` | Find verified professional email addresses by combining a person's full name with their company domain or company name. |
| `EmailVerify` | GET / POST `/v2/email/verify` | Verify whether an email address is valid and deliverable, returning a clear status to prevent bounces and maintain list hygiene. |
| `EuropagesCompany` | GET / POST `/v2/europages/company` | Europages company profile with address, supplier type and the products listed on the page (1 credit). |
| `EuropagesSearch` | GET / POST `/v2/europages/search` | Search Europages suppliers by keyword or category URL. Returns company cards and pagination (1 credit). |
| `EventbriteEvent` | GET / POST `/v2/eventbrite/event` | Get one Eventbrite event by detail URL or numeric id: venue, organizer, schedule, and description (1 credit). |
| `EventbriteSearch` | GET / POST `/v2/eventbrite/search` | Search Eventbrite events by keyword or paste a city/category URL. Returns event cards and pagination (1 credit). |
| `ExpediaAd` | GET / POST `/v2/expedia/ad` | Fetch one Expedia hotel from its information URL: name, postal address, and amenity groups. |
| `ExpediaFlight` | GET / POST `/v2/expedia/flight` | Search Expedia flights by airport route and dates, or by a flight search URL. Each offer includes the airline, times, stops, and displayed… |
| `ExpediaSearch` | GET / POST `/v2/expedia/search` | Execute a hotel search on Expedia using a direct search URL and receive structured hotel listings with prices, ratings, and availability. |
| `FashionphileProduct` | GET / POST `/v2/fashionphile/product` | Load one Fashionphile product from a handle or a product URL. Keywords and search URLs belong on the search endpoint. |
| `FashionphileSearch` | GET / POST `/v2/fashionphile/search` | Search Fashionphile by keywords, a search URL, a collection URL, or a products.json URL. A product handle or product URL belongs on the pro… |
| `FinderChewyProducts` | GET / POST `/v2/finder/chewy-products` | Find Chewy product URLs by product name, ID, or any combination of filters to use in product detail and rendering API calls. |
| `FinderEleclercStores` | GET / POST `/v2/finder/eleclerc-stores` | Search Eleclerc stores by name or city to retrieve contact details, opening hours, and geolocation. |
| `FinderGoogleCountries` | GET / POST `/v2/finder/google-countries` | Search the list of Google-supported countries by name or country code to use in Google Search API requests. |
| `FinderGoogleLanguages` | GET / POST `/v2/finder/google-languages` | Search the list of Google-supported languages by name to retrieve language codes for use in Google Search API requests. |
| `FinderGoogleLocations` | GET / POST `/v2/finder/google-locations` | Search Google-supported locations by name to retrieve location IDs and GPS coordinates for geo-targeted search API calls. |
| `FinderLeroymerlinProducts` | GET / POST `/v2/finder/leroymerlin-products` | Find Leroy Merlin product URLs by product name, ID, or any combination of filters to use in product detail and rendering API calls. |
| `FinderLeroymerlinStores` | GET / POST `/v2/finder/leroymerlin-stores` | Search Leroy Merlin stores by name, city or postal code to retrieve their store ID, slug, address and store locator URL. |
| `FinderLinkedinIndustries` | GET / POST `/v2/finder/linkedin-industries` | Search LinkedIn industry taxonomy to retrieve industry IDs and hierarchies for use in LinkedIn API calls and data enrichment. |
| `FinderMrbricolageStores` | GET / POST `/v2/finder/mrbricolage-stores` | Search Mr.Bricolage stores by name, city or postal code to retrieve their store ID, slug, contact information and geolocation. |
| `FiverrProfileInfo` | GET / POST `/v2/fiverr/profile/info` | Extract complete public profile data from any Fiverr seller, including gigs, ratings, reviews, skills, portfolio, and consultation pricing. |
| `FiverrSearch` | GET / POST `/v2/fiverr/search` | Search Fiverr gigs using a direct search URL and receive paginated listings with seller information, gig titles, images, and seller levels. |
| `FnacProduct` | GET / POST `/v2/fnac/product` | Read one Fnac product from its URL or product id. You get the price, brand, EAN, seller, and rating. |
| `FnacSearch` | GET / POST `/v2/fnac/search` | Search Fnac for a keyword or a SearchResult URL. Each card has the product id, name, price, seller, and stock. |
| `FotocasaAd` | GET / POST `/v2/fotocasa/ad` | Get a full Fotocasa listing by URL: price, features, coordinates, energy certificate, images, and agency (browser rendering, 2 credits). |
| `FotocasaSearch` | GET / POST `/v2/fotocasa/search` | Search Fotocasa listings by city or search URL. Returns prices, rooms, area, coordinates, and agency data via browser rendering (2 credits). |
| `FranchisedirectFranchise` | GET / POST `/v2/franchisedirect/franchise` | Read one Franchise Direct profile: cash requirement, franchise fee, investment range, and training. |
| `FranchisedirectSearch` | GET / POST `/v2/franchisedirect/search` | Search Franchise Direct listings by keyword or listing URL. Each card includes a USD cash figure and a numeric minimum. |
| `FressnapfProduct` | GET / POST `/v2/fressnapf/product` | Read a Fressnapf Germany product sheet from its URL. Returns sku, GTIN, euro price, unit price, and pack variants. |
| `FressnapfSearch` | GET / POST `/v2/fressnapf/search` | Search Fressnapf Germany by keyword or category URL. Returns product cards with euro prices and pagination. |
| `FundaProperty` | GET / POST `/v2/funda/property` | Read one Funda listing from its URL. You get the asking price, living area, and the feature groups. |
| `FundaSearch` | GET / POST `/v2/funda/search` | Search Funda by place. Set offering to koop for sale or huur for rent, then page the listings. |
| `G2ProductInfo` | GET / POST `/v2/g2/product/info` | Read a G2 product reviews page: rating, pricing, categories, alternatives, screenshots, and the reviews shown on the page. |
| `GithubUserInfo` | GET / POST `/v2/github/user/info` | Retrieve public GitHub user profile data including bio, company, location, follower count, repositories, and associated email addresses. |
| `GlassdoorCompanyInfo` | GET / POST `/v2/glassdoor/company/info` | Get one Glassdoor company profile by URL or employer id (2 credits). |
| `GlassdoorCompanySearch` | GET / POST `/v2/glassdoor/company/search` | Search Glassdoor employers by keyword across country sites (2 credits). |
| `GlassdoorJobInfo` | GET / POST `/v2/glassdoor/job/info` | Get one Glassdoor job listing by URL or jl id (2 credits). |
| `GlassdoorJobSearch` | GET / POST `/v2/glassdoor/job/search` | Search Glassdoor jobs by keyword or SEO URL, with domain and filters (2 credits). |
| `GoodreadsBook` | GET / POST `/v2/goodreads/book` | Read one Goodreads book by the numeric id from search, such as 44767458 for Dune. |
| `GoodreadsSearch` | GET / POST `/v2/goodreads/search` | Search Goodreads books by keyword. pagination.next is a Goodreads search URL for the following page. |
| `GoogleImages` | GET / POST `/v2/google/images` | Google Images Scraper API – bringing the power of visual search to your fingertips. |
| `GoogleJobs` | GET / POST `/v2/google/jobs` | Extract job listings from Google Careers for enhanced recruitment and market analysis. |
| `GoogleMaps` | GET / POST `/v2/google/maps` | Extract comprehensive company data from Google Maps for enhanced business insights. |
| `GoogleNews` | GET / POST `/v2/google/news` | Extract real-time news articles from Google News with advanced filtering by topic, location, language, and pagination. |
| `GoogleSearch` | GET / POST `/v2/google/search` | Scrape real-time Google search results with organic listings, knowledge graph, related searches, and advanced filters. |
| `GoogleSearchAI` | GET / POST `/v2/google/search/ai` | Scrape real-time Google AI Mode answers: text_blocks, references, and shopping_results. |
| `GoogleSearchAutocomplete` | GET / POST `/v2/google/search/autocomplete` | Retrieve real-time Google Search autocomplete suggestions for any query, with support for language and country localization. |
| `GoogleShopping` | GET / POST `/v2/google/shopping` | Search Google Shopping and retrieve structured product listings with prices, ratings, merchants, and product IDs. |
| `GoogleVideos` | GET / POST `/v2/google/videos` | Scrape video results from Google Videos search including YouTube, Reddit, and other video platforms, with location and language support. |
| `GoogleadstransparencyAd` | GET / POST `/v2/googleadstransparency/ad` | Get one Google Ads Transparency creative by URL: dates, impressions, platforms, regions, and the text read from the preview image (2 credit… |
| `GoogleadstransparencySearch` | GET / POST `/v2/googleadstransparency/search` | Search the Google Ads Transparency Center by domain, advertiser, region, format, platform, and date. Returns creative cards and pagination… |
| `GraingerProduct` | GET / POST `/v2/grainger/product` | Get a Grainger US or Canada product page: SKU, brand, price, images, description and specs via browser rendering. |
| `GraingerSearch` | GET / POST `/v2/grainger/search` | Search Grainger US or Canada by keyword, category URL or SKU. Returns product listings, category branches and suggestions via browser rende… |
| `GreenhouseJob` | GET / POST `/v2/greenhouse/job` | Fetch one Greenhouse job by board/id, such as discord/8806482002, or by a job-boards.greenhouse.io job URL. |
| `GreenhouseSearch` | GET / POST `/v2/greenhouse/search` | List jobs on a Greenhouse board. query is the board token, such as discord. keyword filters title and location. |
| `GsaelibraryContract` | GET / POST `/v2/gsaelibrary/contract` | Read one GSA eLibrary contractor record from the contract URL returned by search. |
| `GsaelibrarySearch` | GET / POST `/v2/gsaelibrary/search` | Search GSA eLibrary for schedule, SIN, and contract rows by keyword or by a full search URL. |
| `HomestraAd` | GET / POST `/v2/homestra/ad` | Retrieve detailed property listing data from Homestra, including price, size, location, features, photos, and nearby properties. |
| `HomestraSearch` | GET / POST `/v2/homestra/search` | Search Homestra property listings using a direct search URL with filters, returning structured results with prices, sizes, locations, and p… |
| `HuggingfaceModel` | GET / POST `/v2/huggingface/model` | Read one public Hugging Face model by an id such as google-bert/bert-base-uncased. |
| `HuggingfaceSearch` | GET / POST `/v2/huggingface/search` | Search the public Hugging Face models API by keyword and follow pagination.next to the next cursor page. |
| `IdealistaProperty` | GET / POST `/v2/idealista/property` | Extract full property details from an Idealista listing URL, including price, specs, photos by room, features, energy rating, and advertise… |
| `IdealistaProvinces` | GET / POST `/v2/idealista/provinces` | List the 50 Spanish provinces Idealista shows for homes for sale. Each row includes the sale search URL to pass as query to Idealista Searc… |
| `IdealistaSearch` | GET / POST `/v2/idealista/search` | Search Idealista listings from a search URL and retrieve prices, specs, photos, agency info, and pagination. Pass a province URL from /v2/i… |
| `IdealoProduct` | GET / POST `/v2/idealo/product` | Get an Idealo product from Spain, France, or Germany. URL or product id. Offers, shipping, specs, and images via WebUnlocker (3 credits). |
| `IdealoSearch` | GET / POST `/v2/idealo/search` | Search Idealo Spain, France, and Germany. Keyword or search URL. Prices, shipping, ratings, and pagination via WebUnlocker (3 credits). |
| `IkeaProduct` | GET / POST `/v2/ikea/product` | Extract full IKEA product details from a product URL, including price, gallery, variants, dimensions, and breadcrumb. |
| `IkeaSearch` | GET / POST `/v2/ikea/search` | Search IKEA products from a search URL and retrieve prices, ratings, dimensions, images, and pagination. |
| `Immobilienscout24Property` | GET / POST `/v2/immobilienscout24/property` | Read one ImmobilienScout24 expose by URL or numeric id. You get cold rent or purchase price, address, and area. |
| `Immobilienscout24Search` | GET / POST `/v2/immobilienscout24/search` | Search ImmobilienScout24 by region and property type. real_estate_type covers rent, buy, house, and land. Page with pagenumber. |
| `Immoscout24chSearch` | GET / POST `/v2/immoscout24ch/search` | Search ImmoScout24.ch listings by city or search URL and retrieve the CHF price, rooms, living area, and address. Property pages are not av… |
| `ImmoweltProperty` | GET / POST `/v2/immowelt/property` | Extract full property details from an Immowelt listing URL, including price, specs, photo gallery, energy certificate, location, and contac… |
| `ImmoweltSearch` | GET / POST `/v2/immowelt/search` | Search Immowelt listings from a search URL and retrieve prices, specs, photos, agency info, and pagination. |
| `ImoveiswebProperty` | GET / POST `/v2/imoveisweb/property` | Read one Imobiliária Aquários listing from its imoveisweb.com.br URL. price_display is rebuilt from the asking amount. |
| `ImoveiswebSearch` | GET / POST `/v2/imoveisweb/search` | Search the Imobiliária Aquários catalog on imoveisweb.com.br. Pass a catalog name such as casas. price_display is rebuilt from the asking a… |
| `IndeedCompanyInfo` | GET / POST `/v2/indeed/company/info` | Indeed Company Info API. Full employer profile by slug or URL. Browser rendering returns headcount, revenue, industry and FAQ or review blo… |
| `IndeedCompanySearch` | GET / POST `/v2/indeed/company/search` | Indeed Company Search API. Search fr.indeed.com employer profiles by keyword or URL. Browser rendering returns ratings, industry and profil… |
| `IndeedJobInfo` | GET / POST `/v2/indeed/job/info` | Indeed Job Info API. Fetch a full job posting by viewjob URL, sponsored click URL or job key. Browser rendering returns description and str… |
| `IndeedJobSalary` | GET / POST `/v2/indeed/job/salary` | Indeed Job Salary API. Career salary aggregates by job title slug or URL. Browser rendering returns median, percentiles and top paying citi… |
| `IndeedJobSearch` | GET / POST `/v2/indeed/job/search` | Indeed Job Search API. Search fr.indeed.com job listings by keyword or URL. Browser rendering returns cards with salary snippets and pagina… |
| `InetProduct` | GET / POST `/v2/inet/product` | Read one Inet.se product by article number or URL, including SEK price, specs, and warehouse stock. |
| `InetSearch` | GET / POST `/v2/inet/search` | Search Inet.se by keyword or category. Prices are inc-VAT SEK, with the next page URL when more results exist. |
| `InstagramPostInfo` | GET / POST `/v2/instagram/post/info` | Monitor competitors and market trends by scraping Instagram posts for valuable data. |
| `InstagramUserInfo` | GET / POST `/v2/instagram/user/info` | Gather user insights efficiently from Instagram for targeted marketing strategies. |
| `IntermarcheProduct` | GET / POST `/v2/intermarche/product` | Get an Intermarché product sheet for a store (shop_id). Try an EAN or product URL with shop_id=11301. 3 credits per call. |
| `IntermarcheSearch` | GET / POST `/v2/intermarche/search` | Search Intermarché products for a store (shop_id). Try query=lait and shop_id=11301. 3 credits per call. |
| `IntersportProduct` | GET / POST `/v2/intersport/product` | Read one Intersport product from its URL. You get price, brand, and the product details. |
| `IntersportSearch` | GET / POST `/v2/intersport/search` | Search Intersport by keyword and page the sports grid. Cards include title, price, and product link. |
| `JameseditionProperty` | GET / POST `/v2/jamesedition/property` | Read one JamesEdition listing from its URL. |
| `JameseditionSearch` | GET / POST `/v2/jamesedition/search` | Search JamesEdition luxury homes by keyword or by a results URL. |
| `Jobs` | GET `/v1/jobs/{id}` | Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: Completed… |
| `KickUserInfo` | GET / POST `/v2/kick/user/info` | Retrieve complete public channel data from Kick including streamer profile, follower count, live status, and recent categories. |
| `KijijiAd` | GET / POST `/v2/kijiji/ad` | Read one Kijiji ad from its id or URL. You get the asking price, the title, and the city. |
| `KijijiSearch` | GET / POST `/v2/kijiji/search` | Search Kijiji with a results URL or a city and category. Language can be English or French. Page through classified ads. |
| `KrogerProduct` | GET / POST `/v2/kroger/product` | Read one Kroger product by UPC. The sheet uses the same default store as search. You get the shelf price and the previous price when Kroger… |
| `KrogerSearch` | GET / POST `/v2/kroger/search` | Search Kroger groceries by keyword and page the results. Prices come from the default pickup store returned in the response. A location id… |
| `LamudiAd` | GET / POST `/v2/lamudi/ad` | Get a full Lamudi listing by URL: price, area, amenities, coordinates, images, and agent (browser rendering, 2 credits). |
| `LamudiSearch` | GET / POST `/v2/lamudi/search` | Search Lamudi Philippines and Indonesia listings. Pass a search URL in query, or use location and filters. Prices, area, rooms, tags, and p… |
| `LeboncoinAd` | GET / POST `/v2/leboncoin/ad` | Extract full details from a Leboncoin ad including price, location, attributes, seller info, and images. Returns structured JSON by default… |
| `LeboncoinSearch` | GET / POST `/v2/leboncoin/search` | Scrape Leboncoin search results or category pages by URL, returning paginated listings with price, location, and attributes. Returns struct… |
| `LeboncoinSearchAPI` | POST `/v2/leboncoin/search_api` | Search Leboncoin listings via structured POST API with advanced filters for category, location, price range, seller type, and sorting. |
| `LenovoProduct` | GET / POST `/v2/lenovo/product` | Extract a Lenovo product page by URL: id, name, price, currency, availability, images, and description (browser rendering, 2 credits). |
| `LenovoSearch` | GET / POST `/v2/lenovo/search` | Search Lenovo products by keyword or listing URL. Returns product id, price, currency, series, and availability via browser rendering (2 cr… |
| `LeroymerlinProduct` | GET / POST `/v2/leroymerlin/product` | Live Leroy Merlin product detail for France and Romania from a URL or identifier: offers, currency, images, features and store context. |
| `LeroymerlinSearch` | GET / POST `/v2/leroymerlin/search` | Live Leroy Merlin search across France and Romania: title, URL, price, currency, brand, seller, rating and pagination. |
| `LidlProduct` | GET / POST `/v2/lidl/product` | Lidl Product API: fetch full product page data from a Lidl URL or ERP id. Returns pricing, images, description, availability and category v… |
| `LidlSearch` | GET / POST `/v2/lidl/search` | Lidl Search API: search product listings by keyword or URL across Lidl online shops in 27 countries. Parses Nuxt SSR payload from HTML; ret… |
| `LinkedinCompanyInfo` | GET / POST `/v2/linkedin/company/info` | LinkedIn Company Info API. Fetch firmographics by LinkedIn slug, URL, company ID, or reverse domain lookup (guideflow.com). Industry, headc… |
| `LinkedinJobCount` | GET / POST `/v2/linkedin/job/count` | Obtain real-time job counts on LinkedIn by industry, region, and role for strategic planning. |
| `LinkedinJobInfo` | GET / POST `/v2/linkedin/job/info` | Dive into LinkedIn job data and trends for in-depth labor market insights. |
| `LinkedinJobSearch` | GET / POST `/v2/linkedin/job/search` | Experience advanced job data extraction from LinkedIn, enhancing your recruitment strategy. |
| `LinkedinJobSuggest` | GET / POST `/v2/linkedin/job/suggest` | Get personalized job suggestions based on keywords or locations to streamline your job search. |
| `LinkedinPostInfo` | GET / POST `/v2/linkedin/post/info` | Simplify LinkedIn post retrieval with our efficient API for enhanced content analysis. |
| `LinkedinProductInfo` | GET / POST `/v2/linkedin/product/info` | Gain competitive insights with detailed LinkedIn product information for strategic advantage. |
| `LinkedinProfileInfo` | GET / POST `/v2/linkedin/profile/info` | Streamline recruitment and business development by extracting valuable LinkedIn profile insights. |
| `LinkedinadlibraryAd` | GET / POST `/v2/linkedinadlibrary/ad` | Get one LinkedIn Ad Library creative by URL or id: copy, CTA, payer, impressions, and targeting (2 credits). |
| `LinkedinadlibrarySearch` | GET / POST `/v2/linkedinadlibrary/search` | Search the LinkedIn Ad Library by advertiser, keyword, country, and date. Returns ad cards and pagination (2 credits). |
| `LoopnetProperty` | GET / POST `/v2/loopnet/property` | Read one LoopNet commercial listing from its URL. |
| `LoopnetSearch` | GET / POST `/v2/loopnet/search` | Search LoopNet commercial listings by city or by a results URL. |
| `MaerskTracking` | GET / POST `/v2/maersk/tracking` | Track a Maersk container or BOL: origin, destination, current milestone and full transport-plan events via browser rendering. |
| `ManomanoProduct` | GET / POST `/v2/manomano/product` | Read a ManoMano product URL. The body is the product page: brand and seller are objects, and the current price is price.primary_price.amoun… |
| `ManomanoProductOffer` | GET / POST `/v2/manomano/product/offer` | List the other sellers for a ManoMano product id. Each offer adds price and price_display on top of the site offer. |
| `ManomanoSearch` | GET / POST `/v2/manomano/search` | Read a ManoMano category URL or keyword. Each card keeps the site fields and adds prices.main, brand_name, and image_fullpath. pagination.n… |
| `MascusAd` | GET / POST `/v2/mascus/ad` | Get full details for a Mascus used equipment listing: specs, price, seller and description. |
| `MascusSearch` | GET / POST `/v2/mascus/search` | Search used equipment listings on Mascus from a category URL, with pagination that keeps country and category filters. |
| `MediamarktProduct` | GET / POST `/v2/mediamarkt/product` | Read one MediaMarkt product by id or product URL, in the region you choose. You get price, currency, EAN, and availability. |
| `MediamarktSearch` | GET / POST `/v2/mediamarkt/search` | Search a MediaMarkt shop by keyword. Set region for Switzerland, Germany, Austria, and the other MediaMarktSaturn countries, then page the… |
| `MicromaniaProduct` | GET / POST `/v2/micromania/product` | Read one Micromania product from its URL or product id. You get the price, brand, EAN, platform, and stock. |
| `MicromaniaSearch` | GET / POST `/v2/micromania/search` | Search Micromania for a keyword or a Search-Show URL. Each card has the product id, name, brand, EAN, price, and stock. |
| `MobileAd` | GET / POST `/v2/mobile/ad` | Mobile Ad API. Fetch full car listing data from a mobile.de ad URL or numeric id. Returns pricing, specs, images and features via webunlock… |
| `MobileSearch` | GET / POST `/v2/mobile/search` | Mobile Search API. Same mobile.de inventory for every locale. locale sets the language label and ad links, not which cars are returned. (3… |
| `MoreleProduct` | GET / POST `/v2/morele/product` | Read one morele.net product page. Flat JSON with PLN price, EAN, specs, and availability. |
| `MoreleSearch` | GET / POST `/v2/morele/search` | Search morele.net by keyword or search URL. Returns product cards with PLN prices and pagination. |
| `MrbricolageSuggest` | GET / POST `/v2/mrbricolage/suggest` | Live Mr.Bricolage product suggestions for a search keyword or EAN, including title, URL, image, price and currency. |
| `NeweggProduct` | GET / POST `/v2/newegg/product` | Read a Newegg product from an item number, catalog id, or URL. Returns price, specs, images, and seller. |
| `NeweggSearch` | GET / POST `/v2/newegg/search` | Search Newegg by keyword or search URL. Returns item numbers, dollar prices, stock, and pagination. |
| `NorthdataCompany` | GET / POST `/v2/northdata/company` | North Data company profile: register, address, financials and history (1 credit). |
| `NorthdataSearch` | GET / POST `/v2/northdata/search` | Search North Data companies and people by name. Returns suggestion cards and pagination (1 credit). |
| `NPIProvider` | GET / POST `/v2/npi/provider` | Fetch one NPPES provider or organization by 10-digit NPI or provider-view URL, including addresses and taxonomies. |
| `NPISearch` | GET / POST `/v2/npi/search` | Search the NPPES NPI registry by person name, organization, or 10-digit NPI. With NPI-2, a text query is the organization name. total_resul… |
| `OceanioSearch` | GET / POST `/v2/oceanio/search` | Search Ocean.io by domain or company name to find the matching company, similar companies, contacts, and people. |
| `OlxAd` | GET / POST `/v2/olx/ad` | Read one OLX ad from its id or URL. You get the asking price, the city, and the seller name. |
| `OlxProduct` | GET / POST `/v2/olx/product` | Read one OLX ad from its id or URL. You get the asking price, the city, and the seller name. |
| `OlxSearch` | GET / POST `/v2/olx/search` | Search OLX Brazil classifieds by keyword and get title, price, location, seller type, and condition as structured JSON in real time. |
| `OscaroProduct` | GET / POST `/v2/oscaro/product` | Oscaro Product API. Fetch full auto parts product data from an Oscaro product URL. Returns pricing, brand, reference, images and availabili… |
| `OscaroSearch` | GET / POST `/v2/oscaro/search` | Oscaro Search API. Search auto parts by keyword, search URL or category listing on oscaro.com, oscaro.es and oscaro.pt. Returns prices and… |
| `OvokoAd` | GET / POST `/v2/ovoko/ad` | Extract an Ovoko listing from a SKU URL: price, currency, images, seller, quality, vehicle make, part numbers, and warranty. |
| `OvokoSearch` | GET / POST `/v2/ovoko/search` | Search Ovoko used-parts listings from a search URL and retrieve title, price, OEM codes, seller, vehicle text, and pagination. |
| `OwlerCompanyInfo` | GET / POST `/v2/owler/company/info` | Retrieve detailed company intelligence from Owler, including revenue, employee count, CEO info, competitors, funding, and industry sectors. |
| `OwlerSearch` | GET / POST `/v2/owler/search` | Search for companies on Owler by name or keyword, returning a list of matching companies with slugs, logos, domains, and profile URLs. |
| `PagesjaunesPageInfo` | GET / POST `/v2/pagesjaunes/page/info` | Retrieve detailed business information from a PagesJaunes page, including contact details, opening hours, services, payment methods, and SI… |
| `PagesjaunesSearch` | GET / POST `/v2/pagesjaunes/search` | Search PagesJaunes for French local businesses by category and location, returning paginated results with contact details and ratings. |
| `PatreonProfileInfo` | GET / POST `/v2/patreon/profile/info` | Retrieve a Patreon creator profile with patron counts, membership tiers, campaign summary, and creator metadata. |
| `PatreonSearch` | GET / POST `/v2/patreon/search` | Search Patreon creators by keyword and retrieve campaign cards with patron counts, post counts, and profile URLs. |
| `PetsmartProduct` | GET / POST `/v2/petsmart/product` | Read one PetSmart product URL, including the selected size, dollar price, flavor, and item number. |
| `PetsmartSearch` | GET / POST `/v2/petsmart/search` | Search PetSmart by keyword or search URL. Page through pet food, litter, and supplies with dollar prices. |
| `PicardProduct` | GET / POST `/v2/picard/product` | Picard Product API. Fetch full frozen food product data from a Picard URL or reference. Returns pricing, images, description, Nutri-Score,… |
| `PicardSearch` | GET / POST `/v2/picard/search` | Picard Search API. Search frozen food listings on picard.fr by keyword or URL. Parses SFCC Search-UpdateGrid HTML; returns prices, Nutri-Sc… |
| `PinterestUserInfo` | GET / POST `/v2/pinterest/user/info` | Access Pinterest user profiles and analytics for targeted marketing and audience insights. |
| `PitchbookCompanyInfo` | GET / POST `/v2/pitchbook/company/info` | Get a PitchBook company profile: overview, funding, competitors and more via browser rendering. |
| `PitchbookSearch` | GET / POST `/v2/pitchbook/search` | Search PitchBook for companies, news and profiles by keyword or search URL via browser rendering. |
| `PointpProduct` | GET / POST `/v2/pointp/product` | Point.P Product API: scrape live /p/ product pages from a full URL. Get title, images, EAN, pricing, specifications, and breadcrumbs as JSO… |
| `PointpSearch` | GET / POST `/v2/pointp/search` | Point.P Search API: search by keyword or URL and get building-material listings with prices, brands, ratings, units, and pagination as JSON… |
| `PrnewswireArticle` | GET / POST `/v2/prnewswire/article` | Read one PR Newswire press release from its URL. |
| `PrnewswireSearch` | GET / POST `/v2/prnewswire/search` | Search PR Newswire press releases by keyword or by a news-search URL. |
| `ProducthuntProductInfo` | GET / POST `/v2/producthunt/product/info` | Retrieve product data from ProductHunt, including tagline, categories, screenshots, reviews rating, and social links. |
| `ProshopProduct` | GET / POST `/v2/proshop/product` | Read one proshop.dk product page. Flat JSON with DKK price, EAN, specs, and stock text. |
| `ProshopSearch` | GET / POST `/v2/proshop/search` | Search proshop.dk by keyword or search URL. Returns product cards with DKK prices and pagination. |
| `RappiProduct` | GET / POST `/v2/rappi/product` | Fetch one Rappi Mexico product, including the primary store price and the other store offers. The primary price can differ from a search ca… |
| `RappiSearch` | GET / POST `/v2/rappi/search` | Search Rappi Mexico on rappi.com.mx. The city resolves to Mexico City. Each row is one store offer, and next is null. |
| `RatemyprofessorsProfessor` | GET / POST `/v2/ratemyprofessors/professor` | Read one Rate My Professors profile, including difficulty and would-take-again when the page shows them. |
| `RatemyprofessorsSearch` | GET / POST `/v2/ratemyprofessors/search` | Search Rate My Professors by name. Cards include school, department, difficulty, and would-take-again when the card shows them. |
| `RealestatecomauSearch` | GET / POST `/v2/realestatecomau/search` | Search realestate.com.au buy, rent, or sold results. A place such as Sydney, NSW defaults to homes for sale. Each card has price, address,… |
| `RealtorProperty` | GET / POST `/v2/realtor/property` | Extract full property details from a Realtor.com listing URL, including price, specs, photos, agents, features, and location data. |
| `RealtorSearch` | GET / POST `/v2/realtor/search` | Search Realtor.com listings from a search URL and retrieve prices, beds, baths, sqft, photos, and property metadata. |
| `RedfinProperty` | GET / POST `/v2/redfin/property` | Read one for-sale Redfin home. Pass the url from search, which already includes the market id. The response is a flat listing, the same fie… |
| `RedfinSearch` | GET / POST `/v2/redfin/search` | Search for-sale homes on Redfin from a city, neighborhood, or county link. Each listing url includes the market id to pass to the property… |
| `ReicheltProduct` | GET / POST `/v2/reichelt/product` | Read one reichelt.de product URL, including article number, EAN, and the euro price. |
| `ReicheltSearch` | GET / POST `/v2/reichelt/search` | Search reichelt.de by keyword or category. Euro prices are rebuilt from the number, including whole-euro amounts. |
| `ReverbProduct` | GET / POST `/v2/reverb/product` | Reverb Product API: fetch full listing data from a Reverb URL or listing id. Returns pricing, specs, images, description and shop info via… |
| `ReverbSearch` | GET / POST `/v2/reverb/search` | Reverb Search API: search music gear listings by keyword, query string or marketplace URL on Reverb. Returns prices, thumbnails and paginat… |
| `RexelProduct` | GET / POST `/v2/rexel/product` | Get a Rexel France product page: references, EAN, descriptions, images and specs via browser rendering. |
| `RexelSearch` | GET / POST `/v2/rexel/search` | Search Rexel France by keyword, EAN or search URL. Returns product cards and pagination via browser rendering. |
| `RightmoveProperty` | GET / POST `/v2/rightmove/property` | Read one Rightmove property from its URL or numeric id. The response is a flat object with price, tenure, floor area, stations, and the age… |
| `RightmoveSearch` | GET / POST `/v2/rightmove/search` | Search Rightmove by a UK place name or a search URL. Each card has a numeric price, a price label rebuilt from that number, the address, th… |
| `RsonlineProduct` | GET / POST `/v2/rsonline/product` | Read one RS Online UK product URL, including the RS stock number, manufacturer part, and pound price. |
| `RsonlineSearch` | GET / POST `/v2/rsonline/search` | Search RS Online UK by keyword or category. Pound prices keep the amount printed on the page. |
| `SeekJob` | GET / POST `/v2/seek/job` | Read one SEEK Australia job from its URL or numeric id. |
| `SeekSearch` | GET / POST `/v2/seek/search` | Search SEEK Australia jobs by keyword or by a results URL. |
| `SelogerAgency` | GET / POST `/v2/seloger/agency` | Retrieve detailed profile information for a SeLoger real estate agency including contact details, ratings, reviews, and address. |
| `SelogerAgencyRent` | GET / POST `/v2/seloger/agency/rent` | Retrieve paginated rental property listings published by a specific SeLoger agency using its numeric agency ID. |
| `SelogerAgencySale` | GET / POST `/v2/seloger/agency/sale` | Retrieve paginated sale property listings published by a specific SeLoger agency using its numeric agency ID. |
| `SelogerProperty` | GET / POST `/v2/seloger/property` | Efficiently extract detailed data from a SeLoger real estate listing including agency info, pricing, photos, description, and property crit… |
| `SelogerSearch` | GET / POST `/v2/seloger/search` | Efficiently extract real estate listings from SeLoger search pages for market insights, property management, and investment analysis. |
| `SephoraProduct` | GET / POST `/v2/sephora/product` | Read one European Sephora product by id or product URL. Set region when you are not on the French shop. |
| `SephoraSearch` | GET / POST `/v2/sephora/search` | Search a European Sephora shop by keyword. Region defaults to France. Page through makeup and skincare cards. The US site is a different pa… |
| `ShophouzzProduct` | GET / POST `/v2/shophouzz/product` | Retrieve full product details from Shophouzz, including title, price, variants, images, vendor, and structured metadata. |
| `ShopifyApps` | GET / POST `/v2/shopify/apps` | Retrieve structured details about any Shopify App Store listing, including title, logo, screenshots, pricing plans, categories, and descrip… |
| `ShopifyCatalog` | GET / POST `/v2/shopify/catalog` | List products from a Shopify shop. query is the shop domain or a URL on that shop. page and limit choose the slice. |
| `ShopifyCollections` | GET / POST `/v2/shopify/collections` | List collections from a Shopify shop. query is the shop domain or a URL on that shop. page and limit choose the slice. |
| `ShopifyProduct` | GET / POST `/v2/shopify/product` | Extract full product details from any Shopify store product page, including title, variants, pricing, images, inventory, and options. |
| `SimilarwebDomain` | GET / POST `/v2/similarweb/domain` | Retrieve website traffic and engagement metrics for a domain from the public Similarweb page, including visits, bounce rate, rank, keywords… |
| `SimilarwebSearch` | GET / POST `/v2/similarweb/search` | Search the public Similarweb page for websites, companies, and mobile apps related to a keyword. |
| `SnapedaProduct` | GET / POST `/v2/snapeda/product` | SnapEDA Product API: fetch full part page data from a SnapEDA URL or MPN/manufacturer slug. Returns CAD models, datasheet, RoHS, distributo… |
| `SnapedaSearch` | GET / POST `/v2/snapeda/search` | SnapEDA Search API: search electronic parts by keyword or URL on SnapMagic Search (ex-SnapEDA). Returns MPN, manufacturer, pricing, CAD ava… |
| `StarbucksMenu` | GET / POST `/v2/starbucks/menu` | Full Starbucks menu for one café, with the default-size price of each item. Prices use browser rendering. |
| `StarbucksProduct` | GET / POST `/v2/starbucks/product` | One Starbucks item at one café, with a price for every size. Get the café from Search first. Prices use browser rendering. |
| `StarbucksSearch` | GET / POST `/v2/starbucks/search` | Nearest Starbucks cafés from coordinates or a store-locator URL. Returns address, hours, amenities, and the store number used by Menu and P… |
| `Stats` | GET `/v2/stats` | Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 mea… |
| `StockxProduct` | GET / POST `/v2/stockx/product` | Retrieve detailed product information from StockX, including market prices, bid/ask data, size variants, and product metadata. |
| `StockxSearch` | GET / POST `/v2/stockx/search` | Search StockX for products by name, brand, SKU, or slug and receive a list of matching items with product IDs, slugs, images, and colorways. |
| `StockxTrends` | GET / POST `/v2/stockx/trends` | Retrieve trending products on StockX by category, with options to sort by featured, most active, or release date. |
| `StubhubEvent` | GET / POST `/v2/stubhub/event` | Read one StubHub event from its URL or numeric id, including the lowest printed ticket price. |
| `StubhubSearch` | GET / POST `/v2/stubhub/search` | Search StubHub events by keyword or by a StubHub search URL. The response is the first embedded page. |
| `TargetProduct` | GET / POST `/v2/target/product` | Get a Target.com product page: price, description, images and variations via browser rendering. |
| `TargetSearch` | GET / POST `/v2/target/search` | Search Target.com by keyword or search URL. Returns product cards with price, rating and pagination via browser rendering. |
| `TdlrLicense` | GET / POST `/v2/tdlr/license` | Fetch one TDLR license by detail id or SearchResultDetail URL, including status, subtype, DBA, and notes. |
| `TdlrSearch` | GET / POST `/v2/tdlr/search` | Search Texas TDLR licenses by name, city, county, or ZIP. license_type such as AIRREF is required for a useful search. 25 rows per page. |
| `TedNotice` | GET / POST `/v2/ted/notice` | TED notice profile: buyer, lots, deadline, value and documents (1 credit). |
| `TedSearch` | GET / POST `/v2/ted/search` | Search TED procurement notices by keyword or filters. Returns notice cards and pagination (1 credit). |
| `ThaliaProduct` | GET / POST `/v2/thalia/product` | Read one Thalia product from its URL. You get price, author or brand, and the product details. |
| `ThaliaSearch` | GET / POST `/v2/thalia/search` | Search Thalia by keyword and page books, media, and gifts. Cards include title, price, and the product link. |
| `TheforkRestaurant` | GET / POST `/v2/thefork/restaurant` | Extract a full TheFork restaurant profile from a restaurant page URL, including ratings, tags, photo gallery, opening hours, and AI review… |
| `TheforkSearch` | GET / POST `/v2/thefork/search` | Search TheFork restaurants from a full search or city URL and retrieve ratings, cuisine, prices, availability, review summaries, and pagina… |
| `ThomasnetCompany` | GET / POST `/v2/thomasnet/company` | Get a full Thomasnet supplier profile from a company URL, including contact details, product categories, and business information. |
| `ThomasnetSearch` | GET / POST `/v2/thomasnet/search` | Search Thomasnet suppliers from a search or category URL and get company listings, filters, name suggestions, and pagination. |
| `TicketmasterEvent` | GET / POST `/v2/ticketmaster/event` | Fetch one Ticketmaster event from a slugged event URL. The slug is searched and the matching event id is returned with the venue address. |
| `TicketmasterSearch` | GET / POST `/v2/ticketmaster/search` | Search Ticketmaster events by keyword. Returns event cards, including Ticketweb partner rows, and pagination. |
| `TiktoktopadsAd` | GET / POST `/v2/tiktoktopads/ad` | Get one TikTok Creative Center Top Ad by URL or id: creative, landing page, CTR, and engagement (1 credit). |
| `TiktoktopadsSearch` | GET / POST `/v2/tiktoktopads/search` | Search TikTok Creative Center Top Ads by region, objective, period, and keyword. Returns ad cards and pagination (1 credit). |
| `TradeaplaneListing` | GET / POST `/v2/tradeaplane/listing` | Read one Trade-A-Plane aircraft listing from its URL. You get the asking price, the airframe details, and the seller. |
| `TradeaplaneSearch` | GET / POST `/v2/tradeaplane/search` | Search Trade-A-Plane aircraft listings by keyword and page the results. Cards include title, price, and the listing link. |
| `TrustpilotCompanyInfo` | GET / POST `/v2/trustpilot/company/info` | Retrieve Trustpilot company profile data, including trust score, star rating, categories, contact info, and review response behavior. |
| `UntappdBeer` | GET / POST `/v2/untappd/beer` | Read one Untappd beer, including ABV, IBU, and brewery. |
| `UntappdSearch` | GET / POST `/v2/untappd/search` | Search Untappd beers by keyword. Cards include brewery, ABV, and IBU. next is null when Untappd does not expose a next page. |
| `UpworkFreelancer` | GET / POST `/v2/upwork/freelancer` | Upwork Freelancer API: fetch public freelancer profile data from a profile ID, username or URL. Returns bio, hourly rate, skills, ratings a… |
| `UpworkJobs` | GET / POST `/v2/upwork/jobs` | Upwork Jobs API: search freelance job listings by skill slug or URL. Returns budgets, skills, experience level and pagination via browser r… |
| `UpworkServices` | GET / POST `/v2/upwork/services` | Upwork Services API: browse Project Catalog listings by category slug or URL. Returns prices, product IDs and pagination via browser render… |
| `UpworkServicesInfo` | GET / POST `/v2/upwork/services/info` | Upwork Services Info API: fetch Project Catalog product details from a product URL. Returns packages, pricing, seller, media and breadcrumb… |
| `Usage` | GET `/v2/usage` | Retrieve real-time usage statistics and credit balance for your Piloterr account. |
| `UsaspendingAward` | GET / POST `/v2/usaspending/award` | Fetch one US federal award by its generated id or by a usaspending.gov award URL. Amounts are a number plus a display string. |
| `UsaspendingSearch` | GET / POST `/v2/usaspending/search` | Search US federal awards by keyword. award_type selects one group: contract (default), grant, idv, or loan. Amounts are a number plus a dis… |
| `VestiairecollectiveProduct` | GET / POST `/v2/vestiairecollective/product` | Fetch one Vestiaire Collective listing by product URL. The response is a flat object with price, size, condition, and seller. |
| `VestiairecollectiveSearch` | GET / POST `/v2/vestiairecollective/search` | Search Vestiaire Collective by keyword or catalog URL. Each card has brand, name, price, size, and a product link. |
| `VintedItem` | GET / POST `/v2/vinted/item` | Fetch a Vinted listing by item URL or numeric ID. Parses the public item page (Next.js payload), returns a normalized snake_case object wit… |
| `VintedSearch` | GET / POST `/v2/vinted/search` | Search Vinted second-hand listings by keyword or catalog URL across vinted.fr, vinted.com and other domains. Returns results and pagination. |
| `VintedUserInfo` | GET / POST `/v2/vinted/user/info` | Retrieve public Vinted seller profile by member URL or user ID. Returns the raw /api/v2/users/{id} object: login, photo, feedback stats, fo… |
| `VivarealProperty` | GET / POST `/v2/vivareal/property` | Read one Viva Real property from its URL. The price follows the operation of that URL, so a sale link stays a sale price. |
| `VivarealSearch` | GET / POST `/v2/vivareal/search` | Search Viva Real by city and operation, sale or rent. Page with pagina. Development cards and 'from' prices are left out. |
| `WalmartProduct` | GET / POST `/v2/walmart/product` | Extract full Walmart product page (PDP) data from a /ip/ URL: pricing, ratings, images, fulfillment, specs, badges, and availability via br… |
| `WalmartSearch` | GET / POST `/v2/walmart/search` | Scrape Walmart.com search result pages from a full URL and retrieve product cards with pricing, ratings, availability, and pagination via b… |
| `WalmartSuggest` | GET / POST `/v2/walmart/suggest` | Get Walmart.com search autocomplete suggestions for a keyword via browser rendering, with optional US ZIP code localization. |
| `WebsiteAntibot` | GET `/v2/website/antibot` | Detect which anti-bot protection protects a website before you scrape. Get the vendor, confidence level, and matching clues from a single U… |
| `WebsiteCrawler` | GET / POST `/v2/website/crawler` | Fetch public pages with fast HTTP crawling (no JavaScript). Best first choice for static or server-rendered HTML; use Rendering or WebUnloc… |
| `WebsiteEmailPhoneExtractor` | GET / POST `/v2/website/email_phone_extractor` | Extract emails and phone numbers from websites for comprehensive contact details, including social media profiles across 12+ platforms. |
| `WebsiteRendering` | GET / POST `/v2/website/rendering` | Render JavaScript-heavy pages in a headless browser and return post-render HTML. Use when Crawler returns empty DOM; escalate to WebUnlocke… |
| `WebsiteScreenshot` | GET / POST `/v2/website/screenshot` | Browser-like Screenshot API that captures webpages as PNG/JPEG/WebP or PDF for previews, QA, reporting, monitoring, and archiving no headle… |
| `WebsiteTechnology` | GET / POST `/v2/website/technology` | Identify the technologies behind any website CMS, frameworks, analytics, CDN, hosting, and more, for competitive market analysis and techno… |
| `WebsiteWebunlocker` | GET / POST `/v2/website/webunlocker` | Bypass advanced anti-bot systems (Cloudflare, DataDome, Akamai, PerimeterX) on whitelisted domains. Combines browser rendering and unlock t… |
| `WeiboPostInfo` | GET / POST `/v2/weibo/post/info` | Retrieve one public Weibo post by post id, including text, images, publish time, engagement counts, and author profile. |
| `WeiboSearch` | GET / POST `/v2/weibo/search` | Search public Weibo posts by keywords. Each post includes the author, text, publish time, engagement counts, image, and video link. |
| `WelcometothejungleCompany` | GET / POST `/v2/welcometothejungle/company` | Welcome to the Jungle company profile: headcount, offices, sectors and website (1 credit). |
| `WelcometothejungleSearch` | GET / POST `/v2/welcometothejungle/search` | Search Welcome to the Jungle jobs. Returns offer cards and pagination (1 credit). |
| `WellfoundCompanyInfo` | GET / POST `/v2/wellfound/company/info` | Retrieve startup and company data from Wellfound (formerly AngelList), including description, funding, team size, hiring status, and social… |
| `XPostInfo` | GET / POST `/v2/x/post/info` | X Post Info API: fetch public tweet data (text, likes, retweets, views, media) from a status URL or tweet ID via browser rendering. |
| `XProfileInfo` | GET / POST `/v2/x/profile/info` | X Profile Info API: fetch public X/Twitter profile data (bio, followers, verification, avatars) from a username or URL via browser renderin… |
| `YahooSearch` | GET / POST `/v2/yahoo/search` | Query Yahoo via API and retrieve structured organic results with country and language. |
| `YellowpagesBusiness` | GET / POST `/v2/yellowpages/business` | Extract a full Yellow Pages business profile from a listing URL, including phone, website, categories, and ratings via browser rendering. |
| `YellowpagesSearch` | GET / POST `/v2/yellowpages/search` | Search Yellow Pages businesses from a full search URL and retrieve names, phones, addresses, ratings, opening hours, and pagination via bro… |
| `YelpBusiness` | GET / POST `/v2/yelp/business` | Fetch a Yelp business profile from a full business URL with phone, address, website, logo, hours, and claim status. |
| `YelpReviews` | GET / POST `/v2/yelp/reviews` | Fetch paginated Yelp reviews for a business with rating, text, author profile, reactions, and attached photos or videos. |
| `YelpSearch` | GET / POST `/v2/yelp/search` | Search Yelp by keyword and location and retrieve paginated business cards with ratings, categories, photos, and sponsored flags. |
| `ZapimoveisAd` | GET / POST `/v2/zapimoveis/ad` | Get a full Zap Imóveis listing by URL: price, IPTU, features, coordinates, images, and publisher (1 credit). |
| `ZapimoveisSearch` | GET / POST `/v2/zapimoveis/search` | Search Zap Imóveis listings by city or search URL. Returns prices, IPTU, condominium fees, rooms, and pagination (1 credit). |
| `ZazzleProduct` | GET / POST `/v2/zazzle/product` | Fetch one Zazzle product by its product URL or by a slug that ends in the numeric product id. |
| `ZazzleSearch` | GET / POST `/v2/zazzle/search` | Search Zazzle products by keyword. page maps to the pg parameter. Prices are a number plus a display string. |
| `ZillowProperty` | GET / POST `/v2/zillow/property` | Extract full property details from a Zillow homedetails URL, including specs, photos, schools, tax history, and agent attribution. |
| `ZillowSearch` | GET / POST `/v2/zillow/search` | Search Zillow listings from a search or map URL and retrieve prices, beds, baths, coordinates, and listing status. |
| `ZonapropProperty` | GET / POST `/v2/zonaprop/property` | Read one ZonaProp listing from its URL. You get the asking price in USD or ARS, the address, and the publisher. |
| `ZonapropSearch` | GET / POST `/v2/zonaprop/search` | Search ZonaProp for apartments or houses in an Argentine city. Page with pagina. A few featured ads can repeat. Development projects priced… |
| `ZooplaProperty` | GET / POST `/v2/zoopla/property` | Extract full property details from a Zoopla listing URL, including price, specs, photo gallery, EPC rating, coordinates, and agent data. |
| `ZooplaSearch` | GET / POST `/v2/zoopla/search` | Search Zoopla listings from a search URL and retrieve prices, specs, photos, agency info, and pagination. |
| `ZooplusProduct` | GET / POST `/v2/zooplus/product` | Read one Zooplus product by id or URL, in the region of that shop. Variants keep their own price and EAN. |
| `ZooplusSearch` | GET / POST `/v2/zooplus/search` | Search a Zooplus country shop by keyword. Region defaults to France. Page through food, litter, and accessory cards with local prices. |
| `ZoroProduct` | GET / POST `/v2/zoro/product` | Get a Zoro product page: title, brand, USD price, Zoro number, manufacturer number, images, and specifications. |
| `ZoroSearch` | GET / POST `/v2/zoro/search` | Search Zoro by keyword or category URL. Returns product cards with USD price, brand, Zoro number, and pagination. |

## All Piloterr SDKs

| SDK | Package | Install |
| --- | --- | --- |
| Node.js | [`@piloterr/sdk`](https://www.npmjs.com/package/@piloterr/sdk) | `npm install @piloterr/sdk` |
| Python | [`piloterr`](https://pypi.org/project/piloterr/) | `pip install piloterr` |
| Go | [`piloterr-go`](https://pkg.go.dev/github.com/piloterr/piloterr-go) | `go get github.com/piloterr/piloterr-go` |
| .NET | [`Piloterr`](https://www.nuget.org/packages/Piloterr) | `dotnet add package Piloterr` |
| Ruby | [`piloterr`](https://rubygems.org/gems/piloterr) | `gem install piloterr` |
| PHP | [`piloterr/sdk`](https://packagist.org/packages/piloterr/sdk) | `composer require piloterr/sdk` |
| Java | [`piloterr-java`](https://github.com/piloterr/piloterr-java) | build from source, see its README |
| CLI | [`@piloterr/cli`](https://www.npmjs.com/package/@piloterr/cli) | `npm install -g @piloterr/cli` |
| MCP server | [`@piloterr/mcp`](https://www.npmjs.com/package/@piloterr/mcp) | `npx -y @piloterr/mcp` |

## License

MIT — see [LICENSE](LICENSE).
