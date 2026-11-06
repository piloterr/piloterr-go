// Code generated from the Piloterr OpenAPI document. DO NOT EDIT.

package piloterr

// AboutyouProductParams holds the parameters of AboutyouProduct (GET / POST /v2/aboutyou/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AboutyouProductParams struct {
	// Product URL or keyword. A keyword opens the first search result.
	Query string `json:"query" api:"query,required"`
	// Shop code. Default is fr. A full link in query picks the shop, and a conflicting region returns 400.
	Region *string `json:"region" api:"region"`
	// Return the rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AboutyouProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AboutyouSearchParams holds the parameters of AboutyouSearch (GET / POST /v2/aboutyou/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AboutyouSearchParams struct {
	// Keyword or About You search, category, or brand URL.
	Query string `json:"query" api:"query,required"`
	// Shop code. Default is fr. A full link in query picks the shop, and a conflicting region returns 400.
	Region *string `json:"region" api:"region"`
	// Only page 1 is available.
	Page *float64 `json:"page" api:"page"`
	// Return the rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AboutyouSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ActionProductParams holds the parameters of ActionProduct (GET / POST /v2/action/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ActionProductParams struct {
	// Action product ID or full product URL.
	Query string `json:"query" api:"query,required"`
	// Locale (fr-fr, de-de, pt-pt, …) or country code (fr, de, pt). Default: fr-fr.
	Region *string `json:"region" api:"region"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ActionProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ActionSearchParams holds the parameters of ActionSearch (GET / POST /v2/action/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ActionSearchParams struct {
	// Search keyword or full Action.com search URL.
	Query string `json:"query" api:"query,required"`
	// Locale (fr-fr, de-de, pt-pt, …) or country code (fr, de, pt). Default: fr-fr.
	Region *string `json:"region" api:"region"`
	// 1-based page number (default: 1).
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ActionSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AldiProductParams holds the parameters of AldiProduct (GET / POST /v2/aldi/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AldiProductParams struct {
	// Aldi product page URL or site path slug.
	Query string `json:"query" api:"query,required"`
	// Target country when query is a path only (default: fr). Ignored for full product URLs.
	Region *string `json:"region" api:"region"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AldiProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AldiSearchParams holds the parameters of AldiSearch (GET / POST /v2/aldi/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AldiSearchParams struct {
	// Aldi category page URL or site path.
	Query string `json:"query" api:"query,required"`
	// Target country when query is a path only (default: fr).
	Region *string `json:"region" api:"region"`
	// 1-based page number override (default: read from URL or 1).
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AldiSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AlibabaProductParams holds the parameters of AlibabaProduct (GET / POST /v2/alibaba/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AlibabaProductParams struct {
	// Full Alibaba product URL or numeric product ID.
	Query string `json:"query" api:"query,required"`
	// Locale subdomain when query is a bare product ID, or to override the host on a product URL.
	Subdomain *string `json:"subdomain" api:"subdomain"`
	// When true, return the rendered HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AlibabaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Subdomain != nil {
		values = append(values, param{"subdomain", *p.Subdomain})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AlibabaSearchParams holds the parameters of AlibabaSearch (GET / POST /v2/alibaba/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AlibabaSearchParams struct {
	// Search keyword or full Alibaba search URL.
	Query string `json:"query" api:"query,required"`
	// Locale subdomain for keyword searches when query is not a full URL.
	Subdomain *string `json:"subdomain" api:"subdomain"`
	// 1-based page number (20 results per page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw open-s JSON as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AlibabaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Subdomain != nil {
		values = append(values, param{"subdomain", *p.Subdomain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AliexpressProductParams holds the parameters of AliexpressProduct (GET / POST /v2/aliexpress/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AliexpressProductParams struct {
	// Full AliExpress product page URL (/item/{id}.html).
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of parsed JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AliexpressProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AliexpressProductReviewsParams holds the parameters of AliexpressProductReviews (GET / POST /v2/aliexpress/product/reviews).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AliexpressProductReviewsParams struct {
	// AliExpress product id or item URL.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Defaults to 1. The feed returns up to 20 reviews per page. pagination.next is the next page URL, or null on the last page.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw feed body as {"result": "..."} instead of parsed JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AliexpressProductReviewsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AliexpressSearchParams holds the parameters of AliexpressSearch (GET / POST /v2/aliexpress/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AliexpressSearchParams struct {
	// Search keyword or full AliExpress search results URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AliexpressSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AliexpressStoreProductParams holds the parameters of AliexpressStoreProduct (GET / POST /v2/aliexpress/store/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AliexpressStoreProductParams struct {
	// AliExpress store URL or numeric store ID.
	Query string `json:"query" api:"query,required"`
	// Optional sort alias for the store all-items page.
	Sort *string `json:"sort" api:"sort"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AliexpressStoreProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Sort != nil {
		values = append(values, param{"sort", *p.Sort})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AllegroProductParams holds the parameters of AllegroProduct (GET / POST /v2/allegro/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AllegroProductParams struct {
	// Full allegro.pl product or offer URL to extract structured PDP data.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AllegroProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AllegroSearchParams holds the parameters of AllegroSearch (GET / POST /v2/allegro/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AllegroSearchParams struct {
	// Search keyword or full Allegro listing URL.
	Query string `json:"query" api:"query,required"`
	// Page number for keyword searches (default 1).
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AllegroSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AmazonProductParams holds the parameters of AmazonProduct (GET / POST /v2/amazon/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AmazonProductParams struct {
	// Amazon product ASIN (Amazon Standard Identification Number) to retrieve full product details for.
	Query string `json:"query" api:"query,required"`
	// Amazon domain extension to query. Controls which Amazon marketplace is targeted. Default: "com".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AmazonProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AmazonProductOfferParams holds the parameters of AmazonProductOffer (GET / POST /v2/amazon/product/offer).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AmazonProductOfferParams struct {
	// Amazon product ASIN (Amazon Standard Identification Number) to retrieve third-party seller offers for.
	Query string `json:"query" api:"query,required"`
	// Amazon domain extension to query. Controls which Amazon marketplace is targeted. Default: "com".
	Domain *string `json:"domain" api:"domain"`
	// Page number for paginated offer results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AmazonProductOfferParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AmazonSearchParams holds the parameters of AmazonSearch (GET / POST /v2/amazon/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AmazonSearchParams struct {
	// Keyword or phrase to search for products on Amazon.
	Query string `json:"query" api:"query,required"`
	// Amazon domain extension to query. Controls which Amazon marketplace is searched. Default: "com".
	Domain *string `json:"domain" api:"domain"`
	// Page number for paginated search results. Each page returns up to 60 results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AmazonSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ApartmentsPropertyParams holds the parameters of ApartmentsProperty (GET / POST /v2/apartments/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ApartmentsPropertyParams struct {
	// Full Apartments.com property listing URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ApartmentsPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ApartmentsSearchParams holds the parameters of ApartmentsSearch (GET / POST /v2/apartments/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ApartmentsSearchParams struct {
	// Full Apartments.com search or city browse URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ApartmentsSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ArgenpropPropertyParams holds the parameters of ArgenpropProperty (GET / POST /v2/argenprop/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ArgenpropPropertyParams struct {
	// The full listing page on argenprop.com. Copy it from results[].url on [Argenprop Search](https://www.piloterr.com/library/argenprop-search). The listing id is the number after the double dash.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ArgenpropPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ArgenpropSearchParams holds the parameters of ArgenpropSearch (GET / POST /v2/argenprop/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ArgenpropSearchParams struct {
	// The place to search, or a full Argenprop results page.
	Query string `json:"query" api:"query,required"`
	// Which page of the list to open. The first page is 1. The placeholder is 2. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Sale or rent, when query is a place name. The placeholder is venta, which means for sale. alquiler is a long-term rental. alquiler-temporal is a short stay. sale, buy, and rent are accepted as aliases. Default: "venta".
	Operation *string `json:"operation" api:"operation"`
	// The kind of home, when query is a place name. The placeholder is departamentos, which means apartments. Also accepted: casas (houses), ph, terrenos, locales, oficinas, and quintas. casa and house map to casas. Default: "departamentos".
	PropertyType *string `json:"property_type" api:"property_type"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ArgenpropSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Operation != nil {
		values = append(values, param{"operation", *p.Operation})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ArgosProductParams holds the parameters of ArgosProduct (GET / POST /v2/argos/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ArgosProductParams struct {
	// Argos product id or product URL.
	Query string `json:"query" api:"query,required"`
	// When true, return raw HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ArgosProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ArgosSearchParams holds the parameters of ArgosSearch (GET / POST /v2/argos/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ArgosSearchParams struct {
	// Keyword or Argos UK search URL.
	Query string `json:"query" api:"query,required"`
	// Page number from 1.
	Page *float64 `json:"page" api:"page"`
	// When true, return raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ArgosSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ArxivPaperParams holds the parameters of ArxivPaper (GET / POST /v2/arxiv/paper).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ArxivPaperParams struct {
	// Required. arXiv id or an abs, pdf, or html URL. Example: 2201.00978 or https://arxiv.org/abs/cond-mat/0011267v1.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ArxivPaperParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ArxivSearchParams holds the parameters of ArxivSearch (GET / POST /v2/arxiv/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ArxivSearchParams struct {
	// Required. Keyword, fielded query such as au:einstein, arXiv id, or an arxiv.org search URL. Example: electron. A paper URL such as https://arxiv.org/abs/2201.00978 returns 400.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number from 1. Example: 2. Defaults to 1. Each page has 10 papers.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ArxivSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AuchanProductParams holds the parameters of AuchanProduct (GET / POST /v2/auchan/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AuchanProductParams struct {
	// URL of the Auchan product page to extract data from.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AuchanProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AuchanSearchParams holds the parameters of AuchanSearch (GET / POST /v2/auchan/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AuchanSearchParams struct {
	// Search keyword or Auchan search URL to retrieve product listings.
	Query string `json:"query" api:"query,required"`
	// Page number for paginated search results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AuchanSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutodocProductParams holds the parameters of AutodocProduct (GET / POST /v2/autodoc/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutodocProductParams struct {
	// HTTPS URL of an Autodoc product page.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutodocProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutodocProductRelatedParams holds the parameters of AutodocProductRelated (GET / POST /v2/autodoc/product/related).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutodocProductRelatedParams struct {
	// Same HTTPS product page URL as Product.
	Query string `json:"query" api:"query,required"`
	// TecDoc maker id from compatibility.cars[] / compatibility.trucks[], e.g. 93 for RENAULT.
	MakerID float64 `json:"maker_id" api:"maker_id,required"`
	// Required. pkw (passenger cars) or lkw (commercial). Must match the compatibility row. Default: "pkw".
	VehicleType string `json:"vehicle_type" api:"vehicle_type,required"`
	// Required maker label from compatibility, e.g. RENAULT. Used to parse vehicle lines.
	MakerName string `json:"maker_name" api:"maker_name,required"`
	// Required article id from Product response or URL path, e.g. 7999541.
	ProductID float64 `json:"product_id" api:"product_id,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutodocProductRelatedParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	values = append(values, param{"maker_id", p.MakerID})
	if p.VehicleType != "" {
		values = append(values, param{"vehicle_type", p.VehicleType})
	}
	if p.MakerName != "" {
		values = append(values, param{"maker_name", p.MakerName})
	}
	values = append(values, param{"product_id", p.ProductID})
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutodocSearchParams holds the parameters of AutodocSearch (GET / POST /v2/autodoc/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutodocSearchParams struct {
	// Full Autodoc search/listing URL, or path only when domain sets the shop.
	Query string `json:"query" api:"query,required"`
	// Regional shop hostname without https:// or www. (e.g. auto-doc.fr, autodoc.de). Omit to default to autodoc.de.
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutodocSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutoeuropeResultParams holds the parameters of AutoeuropeResult (GET / POST /v2/autoeurope/result).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutoeuropeResultParams struct {
	// Vehicle id from search results[].id, or a package id of that vehicle.
	OfferID string `json:"offer_id" api:"offer_id,required"`
	// Same pickup place as the search. Example: CDG
	Pickup string `json:"pickup" api:"pickup,required"`
	// Same pickup date as the search, YYYY-MM-DD. Example: 2026-11-06
	PickupDate string `json:"pickup_date" api:"pickup_date,required"`
	// Same return date as the search, YYYY-MM-DD, after pickup_date. Example: 2026-11-09
	DropoffDate string `json:"dropoff_date" api:"dropoff_date,required"`
	// Country site used for the search. Default fr. Example: fr. There is no page parameter.
	Region *string `json:"region" api:"region"`
	// Return place when it differs from pickup.
	Dropoff *string `json:"dropoff" api:"dropoff"`
	// Pickup time HH:MM. Default 10:00. Default: "10:00".
	PickupTime *string `json:"pickup_time" api:"pickup_time"`
	// Return time HH:MM. Default 10:00. Default: "10:00".
	DropoffTime *string `json:"dropoff_time" api:"dropoff_time"`
	// Driver age from 18 to 99. Default 40.
	DriverAge *float64 `json:"driver_age" api:"driver_age"`
	// When true, return the raw quote HTML instead of JSON. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutoeuropeResultParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 13)
	if p.OfferID != "" {
		values = append(values, param{"offer_id", p.OfferID})
	}
	if p.Pickup != "" {
		values = append(values, param{"pickup", p.Pickup})
	}
	if p.PickupDate != "" {
		values = append(values, param{"pickup_date", p.PickupDate})
	}
	if p.DropoffDate != "" {
		values = append(values, param{"dropoff_date", p.DropoffDate})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Dropoff != nil {
		values = append(values, param{"dropoff", *p.Dropoff})
	}
	if p.PickupTime != nil {
		values = append(values, param{"pickup_time", *p.PickupTime})
	}
	if p.DropoffTime != nil {
		values = append(values, param{"dropoff_time", *p.DropoffTime})
	}
	if p.DriverAge != nil {
		values = append(values, param{"driver_age", *p.DriverAge})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutoeuropeSearchParams holds the parameters of AutoeuropeSearch (GET / POST /v2/autoeurope/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutoeuropeSearchParams struct {
	// Pickup place or airport code.
	Pickup string `json:"pickup" api:"pickup,required"`
	// Pickup date, YYYY-MM-DD.
	PickupDate string `json:"pickup_date" api:"pickup_date,required"`
	// Return date, YYYY-MM-DD, after pickup_date.
	DropoffDate string `json:"dropoff_date" api:"dropoff_date,required"`
	// Country site. Default fr. Example: fr. A code such as de or uk, or a full Auto Europe URL. There is no page parameter.
	Region *string `json:"region" api:"region"`
	// Return place when it differs from pickup.
	Dropoff *string `json:"dropoff" api:"dropoff"`
	// Pickup time HH:MM. Default 10:00. Default: "10:00".
	PickupTime *string `json:"pickup_time" api:"pickup_time"`
	// Return time HH:MM. Default 10:00. Default: "10:00".
	DropoffTime *string `json:"dropoff_time" api:"dropoff_time"`
	// Driver age from 18 to 99. Default 40.
	DriverAge *float64 `json:"driver_age" api:"driver_age"`
	// When true, return the raw quote HTML instead of JSON. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutoeuropeSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Pickup != "" {
		values = append(values, param{"pickup", p.Pickup})
	}
	if p.PickupDate != "" {
		values = append(values, param{"pickup_date", p.PickupDate})
	}
	if p.DropoffDate != "" {
		values = append(values, param{"dropoff_date", p.DropoffDate})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Dropoff != nil {
		values = append(values, param{"dropoff", *p.Dropoff})
	}
	if p.PickupTime != nil {
		values = append(values, param{"pickup_time", *p.PickupTime})
	}
	if p.DropoffTime != nil {
		values = append(values, param{"dropoff_time", *p.DropoffTime})
	}
	if p.DriverAge != nil {
		values = append(values, param{"driver_age", *p.DriverAge})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// Autoscout24AdParams holds the parameters of Autoscout24Ad (GET / POST /v2/autoscout24/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type Autoscout24AdParams struct {
	// Full AutoScout24 listing URL. An inactive listing that still contains script#__NEXT_DATA__ is returned with status Inactive. A page without that selector returns 404.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *Autoscout24AdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// Autoscout24SearchParams holds the parameters of Autoscout24Search (GET / POST /v2/autoscout24/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type Autoscout24SearchParams struct {
	// Full AutoScout24 search URL. Filters and page in the URL are kept. page starts at 1. A page with no cards returns 404.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *Autoscout24SearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutotraderAdParams holds the parameters of AutotraderAd (GET / POST /v2/autotrader/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutotraderAdParams struct {
	// Full AutoTrader UK car details URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutotraderAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// AutotraderSearchParams holds the parameters of AutotraderSearch (GET / POST /v2/autotrader/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type AutotraderSearchParams struct {
	// Full AutoTrader UK car search URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AutotraderSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BestbuyProductParams holds the parameters of BestbuyProduct (GET / POST /v2/bestbuy/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BestbuyProductParams struct {
	// Full Best Buy product URL.
	Query string `json:"query" api:"query,required"`
	// When true, return rendered HTML as {"result": "..."} instead of JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BestbuyProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BestbuyProductReviewsParams holds the parameters of BestbuyProductReviews (GET / POST /v2/bestbuy/product/reviews).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BestbuyProductReviewsParams struct {
	// Best Buy product URL or reviews URL.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Default 1. A value above 1 overrides page on the Best Buy URL. If this stays at 1, a page already on the URL is kept.
	Page *float64 `json:"page" api:"page"`
	// When true, return the rendered HTML instead of JSON. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BestbuyProductReviewsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BestbuySearchParams holds the parameters of BestbuySearch (GET / POST /v2/bestbuy/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BestbuySearchParams struct {
	// Keyword or full Best Buy search URL.
	Query string `json:"query" api:"query,required"`
	// Best Buy country when query is not a full URL (default: com).
	Region *string `json:"region" api:"region"`
	// Page number (default: 1).
	Page *float64 `json:"page" api:"page"`
	// Sort order (default: best_selling). Default: "best_selling".
	Sort *string `json:"sort" api:"sort"`
	// When true, return rendered HTML as {"result": "..."} instead of JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BestbuySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Sort != nil {
		values = append(values, param{"sort", *p.Sort})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BhphotovideoProductParams holds the parameters of BhphotovideoProduct (GET / POST /v2/bhphotovideo/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BhphotovideoProductParams struct {
	// Product URL or id.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML. The string false stays false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BhphotovideoProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BhphotovideoSearchParams holds the parameters of BhphotovideoSearch (GET / POST /v2/bhphotovideo/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BhphotovideoSearchParams struct {
	// Keyword or a bhphotovideo.com search URL.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Example: 2.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML. The string false stays false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BhphotovideoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BingSearchParams holds the parameters of BingSearch (GET / POST /v2/bing/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BingSearchParams struct {
	// The search query to execute on Bing. Supports all Bing advanced search operators.
	Query string `json:"query" api:"query,required"`
	// Page number of results to retrieve. Used for pagination.
	Page *float64 `json:"page" api:"page"`
	// Number of results per page. Minimum: 10, Maximum: 50.
	Num *float64 `json:"num" api:"num"`
	// City or region from which the search should originate. If multiple locations match, the most popular one is selected.
	Location *string `json:"location" api:"location"`
	// GPS latitude coordinate for the search origin.
	Lat *float64 `json:"lat" api:"lat"`
	// GPS longitude coordinate for the search origin.
	Lon *float64 `json:"lon" api:"lon"`
	// Market code combining language and country, in the form <language>-<country>. Defines where the results come from.
	Mkt *string `json:"mkt" api:"mkt"`
	// Two-letter ISO 3166-1 country code defining the country to search from.
	Cc *string `json:"cc" api:"cc"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BingSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Lat != nil {
		values = append(values, param{"lat", *p.Lat})
	}
	if p.Lon != nil {
		values = append(values, param{"lon", *p.Lon})
	}
	if p.Mkt != nil {
		values = append(values, param{"mkt", *p.Mkt})
	}
	if p.Cc != nil {
		values = append(values, param{"cc", *p.Cc})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BitibaProductParams holds the parameters of BitibaProduct (GET / POST /v2/bitiba/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BitibaProductParams struct {
	// Product URL, shop path, or product number. Example: 75036. A search link returns 400.
	Query string `json:"query" api:"query,required"`
	// Shop used when query is a number or a path. de is Germany (default). Default: "de".
	Region *string `json:"region" api:"region"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BitibaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BitibaSearchParams holds the parameters of BitibaSearch (GET / POST /v2/bitiba/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BitibaSearchParams struct {
	// What to look for. A word such as hundefutter, or a Bitiba search link. A product link returns 400. www.bitiba.com returns 400.
	Query string `json:"query" api:"query,required"`
	// Country shop. de is Germany (default), fr is France, uk is the United Kingdom. gb means uk. Default: "de".
	Region *string `json:"region" api:"region"`
	// Page number, starting at 1. Page 2 uses p=2.
	Page *float64 `json:"page" api:"page"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BitibaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BolProductParams holds the parameters of BolProduct (GET / POST /v2/bol/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BolProductParams struct {
	// Product URL or numeric product id.
	Query string `json:"query" api:"query,required"`
	// Shop. Use nl for the Netherlands in Dutch, be-nl for Belgium in Dutch, or be-fr for Belgium in French. Leave empty for nl. A full URL picks the shop from the path.
	Region *string `json:"region" api:"region"`
	// When true, return the raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BolProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BolSearchParams holds the parameters of BolSearch (GET / POST /v2/bol/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BolSearchParams struct {
	// Keyword or full bol.com search URL. A product URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Shop. Use nl for the Netherlands in Dutch, be-nl for Belgium in Dutch, or be-fr for Belgium in French. Leave empty for nl. A full URL picks the shop from /nl/nl/, /be/nl/, or /be/fr/.
	Region *string `json:"region" api:"region"`
	// Page number, starting at 1. Page 2 asks bol for the next screen of the same search.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BolSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BoulangerProductParams holds the parameters of BoulangerProduct (GET / POST /v2/boulanger/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BoulangerProductParams struct {
	// Product URL /ref/{id}, or the numeric product id.
	Query string `json:"query" api:"query,required"`
	// When true, returns the raw HTML instead of the parsed JSON. The response shape is {"result": "<html>..."}. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BoulangerProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BoulangerSearchParams holds the parameters of BoulangerSearch (GET / POST /v2/boulanger/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BoulangerSearchParams struct {
	// Keyword, a Boulanger /resultats URL, or a /c/ category URL.
	Query string `json:"query" api:"query,required"`
	// Page number, from 1. Default is 1. Page 2 is sent as numPage=2 and returns different product ids.
	Page *float64 `json:"page" api:"page"`
	// When true, returns the raw HTML instead of the parsed JSON. The response shape is {"result": "<html>..."}. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BoulangerSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BraveSearchParams holds the parameters of BraveSearch (GET / POST /v2/brave/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BraveSearchParams struct {
	// The search query to execute on Brave Search. Supports all standard Brave search operators.
	Query string `json:"query" api:"query,required"`
	// Page number of results to retrieve. Used for pagination. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code, for example us, fr, gb, de. Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language code, for example en, fr, de. Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BraveSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BricodepotProductParams holds the parameters of BricodepotProduct (GET / POST /v2/bricodepot/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BricodepotProductParams struct {
	// Product URL /p/{ean}/{slug}.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BricodepotProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BricodepotSearchParams holds the parameters of BricodepotSearch (GET / POST /v2/bricodepot/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BricodepotSearchParams struct {
	// Brico Dépôt listing URL under /produits/.
	Query string `json:"query" api:"query,required"`
	// Page from 1. Page 2 is the /2 path suffix.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BricodepotSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BricoramaProductParams holds the parameters of BricoramaProduct (GET / POST /v2/bricorama/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BricoramaProductParams struct {
	// Product URL /p/{slug}/{ean}.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BricoramaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BricoramaSearchParams holds the parameters of BricoramaSearch (GET / POST /v2/bricorama/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type BricoramaSearchParams struct {
	// Keyword, or a Bricorama /search URL.
	Query string `json:"query" api:"query,required"`
	// Page from 1. Sent as ?page= on the listing.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BricoramaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CarrefourProductParams holds the parameters of CarrefourProduct (GET / POST /v2/carrefour/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CarrefourProductParams struct {
	// Product page URL on www.carrefour.fr. The path starts with /p/.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CarrefourProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CarrefourSearchParams holds the parameters of CarrefourSearch (GET / POST /v2/carrefour/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CarrefourSearchParams struct {
	// Search keyword or full Carrefour search page URL.
	Query string `json:"query" api:"query,required"`
	// Carrefour domain to search on. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Page number to retrieve. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CarrefourSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CarrefourSuggestParams holds the parameters of CarrefourSuggest (GET / POST /v2/carrefour/suggest).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CarrefourSuggestParams struct {
	// Search keyword or EAN code to get product suggestions from Carrefour.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CarrefourSuggestParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CastoramaProductParams holds the parameters of CastoramaProduct (GET / POST /v2/castorama/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CastoramaProductParams struct {
	// Product URL ending in _CAFR.prd.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CastoramaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CastoramaSearchParams holds the parameters of CastoramaSearch (GET / POST /v2/castorama/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CastoramaSearchParams struct {
	// Keyword, or a Castorama listing URL (/search or a .cat category).
	Query string `json:"query" api:"query,required"`
	// Page from 1. Sent as ?page= on the listing.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CastoramaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CdiscountProductParams holds the parameters of CdiscountProduct (GET / POST /v2/cdiscount/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CdiscountProductParams struct {
	// Full Cdiscount product page URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CdiscountProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CdiscountProductOfferParams holds the parameters of CdiscountProductOffer (GET / POST /v2/cdiscount/product/offer).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CdiscountProductOfferParams struct {
	// Cdiscount product offer page URL. Pass the multiple_offer_url value returned by the Cdiscount Product API.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CdiscountProductOfferParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CdiscountSearchParams holds the parameters of CdiscountSearch (GET / POST /v2/cdiscount/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CdiscountSearchParams struct {
	// Cdiscount search results page URL. Use any URL from a search or category browse on cdiscount.com.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CdiscountSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CentrisPropertyParams holds the parameters of CentrisProperty (GET / POST /v2/centris/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CentrisPropertyParams struct {
	// Full Centris listing URL, in English or French.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CentrisPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CentrisSearchParams holds the parameters of CentrisSearch (GET / POST /v2/centris/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CentrisSearchParams struct {
	// Place name or a full Centris results URL.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Default is 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Results language when query is a place name. Default is en. Default: "en".
	Language *string `json:"language" api:"language"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CentrisSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Language != nil {
		values = append(values, param{"language", *p.Language})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ChewyProductParams holds the parameters of ChewyProduct (GET / POST /v2/chewy/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ChewyProductParams struct {
	// Which product to read. A Chewy product URL such as https://www.chewy.com/purina-pro-plan-adult-sensitive-skin/dp/128666, or just the numeric id 128666. A search link, a keyword, or another website returns 400.
	Query string `json:"query" api:"query,required"`
	// Which JSON to return. Default expert. Default: "expert".
	Mode *string `json:"mode" api:"mode"`
	// Set to true to get the rendered HTML in result instead of JSON. Default false. The string false stays false and still returns JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ChewyProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Mode != nil {
		values = append(values, param{"mode", *p.Mode})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ChewySearchParams holds the parameters of ChewySearch (GET / POST /v2/chewy/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ChewySearchParams struct {
	// What to search for. A keyword such as salmon kibble, or a Chewy search or category link such as https://www.chewy.com/s?query=dog+food. A product link (a URL that contains /dp/) returns 400. Another website returns 400.
	Query string `json:"query" api:"query,required"`
	// Which page to fetch, starting at 1. Leave this empty for page 1. If query is a listing link that already has page=2, this value replaces that page. In simple mode, pagination.page is this page and pagination.next is the next page URL or null.
	Page *float64 `json:"page" api:"page"`
	// Which JSON to return. Default expert. Default: "expert".
	Mode *string `json:"mode" api:"mode"`
	// Set to true to get the rendered HTML in result instead of JSON. Default false. The string false stays false and still returns JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ChewySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Mode != nil {
		values = append(values, param{"mode", *p.Mode})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ClearedjobsJobParams holds the parameters of ClearedjobsJob (GET / POST /v2/clearedjobs/job).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ClearedjobsJobParams struct {
	// Required. Job URL or slug ending in the numeric id. Example: analyst-washington-district-of-columbia-1902823.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ClearedjobsJobParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ClearedjobsSearchParams holds the parameters of ClearedjobsSearch (GET / POST /v2/clearedjobs/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ClearedjobsSearchParams struct {
	// Required. Keyword or a clearedjobs.net jobs URL. Example: analyst. A job URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Optional. Only page 1 is available. Example: 1. page=2 returns 404. pagination.next is always null.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ClearedjobsSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CompanyParams holds the parameters of Company (GET / POST /v2/company).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CompanyParams struct {
	// The domain name of the company to look up. Accepts multiple formats:
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CompanyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CoolblueProductParams holds the parameters of CoolblueProduct (GET / POST /v2/coolblue/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CoolblueProductParams struct {
	// Product link, or the article number.
	Query string `json:"query" api:"query,required"`
	// Shop for a bare article number.
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CoolblueProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CoolblueSearchParams holds the parameters of CoolblueSearch (GET / POST /v2/coolblue/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CoolblueSearchParams struct {
	// Search words, or a Coolblue listing link.
	Query string `json:"query" api:"query,required"`
	// Page number, from 1.
	Page *float64 `json:"page" api:"page"`
	// Shop, when query is words rather than a full link.
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CoolblueSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CopartLotParams holds the parameters of CopartLot (GET / POST /v2/copart/lot).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CopartLotParams struct {
	// Copart lot number or full lot URL.
	Query string `json:"query" api:"query,required"`
	// Regional Copart site when query is a lot number only. Defaults to US.
	Region *string `json:"region" api:"region"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CopartLotParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CopartSearchParams holds the parameters of CopartSearch (GET / POST /v2/copart/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CopartSearchParams struct {
	// Search keyword or full Copart search URL.
	Query string `json:"query" api:"query,required"`
	// Copart regional site when query is a keyword or lot number only. Defaults to US.
	Region *string `json:"region" api:"region"`
	// Page number (default: 1).
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CopartSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CostcoProductParams holds the parameters of CostcoProduct (GET / POST /v2/costco/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CostcoProductParams struct {
	// Costco product ID or full product URL.
	Query string `json:"query" api:"query,required"`
	// Regional Costco site. Accepts a region code (us, fr, ca, uk, …) or full host (www.costco.fr). Defaults to US (www.costco.com) when omitted.
	Region *string `json:"region" api:"region"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CostcoProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CostcoSearchParams holds the parameters of CostcoSearch (GET / POST /v2/costco/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CostcoSearchParams struct {
	// Search keyword or full regional search URL.
	Query string `json:"query" api:"query,required"`
	// Regional Costco site. Accepts a region code (us, fr, ca, uk, …) or full host (www.costco.fr). Defaults to US (www.costco.com) when omitted.
	Region *string `json:"region" api:"region"`
	// 1-based page number override. Uses currentPage (most regions) or page (UK) when rebuilding pagination URLs.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CostcoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CrunchbaseCompanyInfoParams holds the parameters of CrunchbaseCompanyInfo (GET / POST /v2/crunchbase/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CrunchbaseCompanyInfoParams struct {
	// Search for a company by its name, permalink, or full Crunchbase URL.
	Query *string `json:"query" api:"query"`
	// Perform a reverse lookup by domain name to find the associated Crunchbase company profile.
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CrunchbaseCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CrunchbaseEventParams holds the parameters of CrunchbaseEvent (GET / POST /v2/crunchbase/event).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CrunchbaseEventParams struct {
	// Retrieve event data using its Crunchbase permalink or full URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CrunchbaseEventParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CrunchbaseFundingRoundParams holds the parameters of CrunchbaseFundingRound (GET / POST /v2/crunchbase/funding_round).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CrunchbaseFundingRoundParams struct {
	// The permalink of the specific funding round to retrieve. This is the unique identifier found in the Crunchbase funding round URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CrunchbaseFundingRoundParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CrunchbaseFundingRoundsParams holds the parameters of CrunchbaseFundingRounds (GET / POST /v2/crunchbase/funding_rounds).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CrunchbaseFundingRoundsParams struct {
	// Filter funding rounds announced within the last N days. Default: 1.
	DaysSinceAnnouncement *float64 `json:"days_since_announcement" api:"days_since_announcement"`
	// Filter by funding round type. Accepted values:
	InvestmentType *string `json:"investment_type" api:"investment_type"`
	// Filter rounds by a specific investor using their Crunchbase UUID.
	InvestorIdentifiers *string `json:"investor_identifiers" api:"investor_identifiers"`
	// Filter rounds by a specific funded organization using their Crunchbase UUID.
	FundedOrganizationIdentifier *string `json:"funded_organization_identifier" api:"funded_organization_identifier"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CrunchbaseFundingRoundsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.DaysSinceAnnouncement != nil {
		values = append(values, param{"days_since_announcement", *p.DaysSinceAnnouncement})
	}
	if p.InvestmentType != nil {
		values = append(values, param{"investment_type", *p.InvestmentType})
	}
	if p.InvestorIdentifiers != nil {
		values = append(values, param{"investor_identifiers", *p.InvestorIdentifiers})
	}
	if p.FundedOrganizationIdentifier != nil {
		values = append(values, param{"funded_organization_identifier", *p.FundedOrganizationIdentifier})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CrunchbasePeopleInfoParams holds the parameters of CrunchbasePeopleInfo (GET / POST /v2/crunchbase/people/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CrunchbasePeopleInfoParams struct {
	// Search for a person by their name, Crunchbase permalink, or full Crunchbase profile URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CrunchbasePeopleInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CrunchbaseSearchParams holds the parameters of CrunchbaseSearch (GET / POST /v2/crunchbase/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CrunchbaseSearchParams struct {
	// The search keyword or phrase to query across Crunchbase. Searches both organizations and people.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CrunchbaseSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CulturaProductParams holds the parameters of CulturaProduct (GET / POST /v2/cultura/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CulturaProductParams struct {
	// Link of the Cultura product page.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CulturaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// CulturaSearchParams holds the parameters of CulturaSearch (GET / POST /v2/cultura/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type CulturaSearchParams struct {
	// Words to search for on Cultura, or the search page link.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. For example, 2 returns the next page of results.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *CulturaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DartyProductParams holds the parameters of DartyProduct (GET / POST /v2/darty/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DartyProductParams struct {
	// Product URL on www.darty.com. The path is under /nav/achat/ and ends in .html.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DartyProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DartySearchParams holds the parameters of DartySearch (GET / POST /v2/darty/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DartySearchParams struct {
	// Keyword or a Darty search URL on www.darty.com.
	Query string `json:"query" api:"query,required"`
	// 1-based page number. Defaults to 1. Page 2 returns the next 30 products.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DartySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DatasetsLinkedinCompanyParams holds the parameters of DatasetsLinkedinCompany (POST /v2/datasets/linkedin/company).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DatasetsLinkedinCompanyParams struct {
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DatasetsLinkedinCompanyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 3)
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DecathlonProductParams holds the parameters of DecathlonProduct (GET / POST /v2/decathlon/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DecathlonProductParams struct {
	// Product link, a /p/... or /products/... path, or a keyword. A keyword opens the first search hit. A search link returns 400.
	Query string `json:"query" api:"query,required"`
	// Country shop when query is a keyword or a path. Default is fr. A full product link picks the shop. Default: "fr".
	Region *string `json:"region" api:"region"`
	// Set to true to get the raw page body in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DecathlonProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DecathlonSearchParams holds the parameters of DecathlonSearch (GET / POST /v2/decathlon/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DecathlonSearchParams struct {
	// Keyword, or a Decathlon search link.
	Query string `json:"query" api:"query,required"`
	// Country shop. Default is fr. A full link in query picks the shop, and a conflicting region returns 400. Default: "fr".
	Region *string `json:"region" api:"region"`
	// Page number starting at 1.
	Page *float64 `json:"page" api:"page"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DecathlonSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DigikeyProductParams holds the parameters of DigikeyProduct (GET / POST /v2/digikey/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DigikeyProductParams struct {
	// Full DigiKey product detail URL, or relative path.
	Query string `json:"query" api:"query,required"`
	// DigiKey country TLD when query is a relative path.
	Region *string `json:"region" api:"region"`
	// When true, return rendered HTML as {"result": "..."} instead of JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DigikeyProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DigikeySearchParams holds the parameters of DigikeySearch (GET / POST /v2/digikey/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DigikeySearchParams struct {
	// Keyword, DigiKey result/filter URL, or relative path.
	Query string `json:"query" api:"query,required"`
	// DigiKey country TLD when query is not a full URL.
	Region *string `json:"region" api:"region"`
	// When true, return rendered HTML as {"result": "..."} instead of JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DigikeySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DoajArticleParams holds the parameters of DoajArticle (GET / POST /v2/doaj/article).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DoajArticleParams struct {
	// Required. Article id or a doaj.org/article URL. Example: 000004687a3a411cb466625281f2ceb0.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DoajArticleParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DoajJournalParams holds the parameters of DoajJournal (GET / POST /v2/doaj/journal).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DoajJournalParams struct {
	// Required. Journal ISSN, journal id, or a doaj.org/toc URL. Example: 2747-173X.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DoajJournalParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DoajSearchParams holds the parameters of DoajSearch (GET / POST /v2/doaj/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DoajSearchParams struct {
	// Required. Keyword or a doaj.org search URL. Example: climate. An article or journal URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Optional. article or journal. Example: journal. Defaults to article. Default: "article".
	Kind *string `json:"kind" api:"kind"`
	// Optional. Page number from 1. Example: 2. Defaults to 1. Each page has 10 records.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DoajSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Kind != nil {
		values = append(values, param{"kind", *p.Kind})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DomainDnsblParams holds the parameters of DomainDnsbl (GET / POST /v2/domain/dnsbl).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DomainDnsblParams struct {
	// A domain name or IP address to check against DNS-based blacklists (DNSBL).
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DomainDnsblParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DomainMaliciousParams holds the parameters of DomainMalicious (GET / POST /v2/domain/malicious).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DomainMaliciousParams struct {
	// A domain name or IP address to check for malicious activity across multiple threat intelligence sources.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DomainMaliciousParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DomainWhoisParams holds the parameters of DomainWhois (GET / POST /v2/domain/whois).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DomainWhoisParams struct {
	// A domain name for which you want to retrieve WHOIS registration information.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DomainWhoisParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DomaincomauPropertyParams holds the parameters of DomaincomauProperty (GET / POST /v2/domaincomau/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DomaincomauPropertyParams struct {
	// Domain listing URL. The path ends with the listing id.
	Query string `json:"query" api:"query,required"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DomaincomauPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DomaincomauSearchParams holds the parameters of DomaincomauSearch (GET / POST /v2/domaincomau/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DomaincomauSearchParams struct {
	// Suburb with state and postcode, a location slug, or a Domain search URL.
	Query string `json:"query" api:"query,required"`
	// sale, rent, or sold when query is a place. Default: "sale".
	Mode *string `json:"mode" api:"mode"`
	// Page from 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DomaincomauSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Mode != nil {
		values = append(values, param{"mode", *p.Mode})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DuckduckgoSearchParams holds the parameters of DuckduckgoSearch (GET / POST /v2/duckduckgo/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type DuckduckgoSearchParams struct {
	// Search query sent to DuckDuckGo Search.
	Query string `json:"query" api:"query,required"`
	// Page number. Page 1 is the first page of results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code, for example us, fr, gb, de, jp.
	Gl *string `json:"gl" api:"gl"`
	// Language code, for example en, fr, de, ja, zh.
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DuckduckgoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EbayProductParams holds the parameters of EbayProduct (GET / POST /v2/ebay/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EbayProductParams struct {
	// Item id, or an item URL on a supported eBay site.
	Query string `json:"query" api:"query,required"`
	// When true, return raw HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EbayProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EbaySearchParams holds the parameters of EbaySearch (GET / POST /v2/ebay/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EbaySearchParams struct {
	// Keyword on www.ebay.com, or a search URL on a supported eBay site.
	Query string `json:"query" api:"query,required"`
	// Page number, sent as _pgn. Default 1 for a keyword. On a URL, the page already in the URL is kept when this is empty. Below 1 returns 400.
	Page *float64 `json:"page" api:"page"`
	// When true, return raw HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EbaySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ElcorteinglesProductParams holds the parameters of ElcorteinglesProduct (GET / POST /v2/elcorteingles/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ElcorteinglesProductParams struct {
	// Product URL or product id (A… / MP_…).
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ElcorteinglesProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ElcorteinglesSearchParams holds the parameters of ElcorteinglesSearch (GET / POST /v2/elcorteingles/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ElcorteinglesSearchParams struct {
	// Search keyword or full El Corte Inglés search URL.
	Query string `json:"query" api:"query,required"`
	// 1-based page number. Overrides /search-nwx/{page}/ when present in the URL (default: from URL or 1).
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ElcorteinglesSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EleclercProductParams holds the parameters of EleclercProduct (GET / POST /v2/eleclerc/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EleclercProductParams struct {
	// E.Leclerc product EAN or product page URL.
	Query string `json:"query" api:"query,required"`
	// E.Leclerc store sign code. This selects that store's price and stock.
	StoreID *string `json:"store_id" api:"store_id"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EleclercProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.StoreID != nil {
		values = append(values, param{"store_id", *p.StoreID})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EleclercSearchParams holds the parameters of EleclercSearch (GET / POST /v2/eleclerc/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EleclercSearchParams struct {
	// Search keyword or phrase to query on E.Leclerc.
	Query string `json:"query" api:"query,required"`
	// ID of a specific E.Leclerc store to search in.
	StoreID *string `json:"store_id" api:"store_id"`
	// Page number for paginated results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EleclercSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.StoreID != nil {
		values = append(values, param{"store_id", *p.StoreID})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EleclercSuggestParams holds the parameters of EleclercSuggest (GET / POST /v2/eleclerc/suggest).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EleclercSuggestParams struct {
	// Search term or EAN to get product suggestions for.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EleclercSuggestParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EmailAnalyzesParams holds the parameters of EmailAnalyzes (GET / POST /v2/email/analyzes).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EmailAnalyzesParams struct {
	// The domain name you want to analyze for email infrastructure.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EmailAnalyzesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EmailFinderParams holds the parameters of EmailFinder (GET / POST /v2/email/finder).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EmailFinderParams struct {
	// Full name of the person whose professional email address you want to find.
	Query string `json:"query" api:"query,required"`
	// The company's domain name. Multiple formats are accepted:
	CompanyDomain *string `json:"company_domain" api:"company_domain"`
	// The company name of the person you are looking for.
	CompanyName *string `json:"company_name" api:"company_name"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EmailFinderParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.CompanyDomain != nil {
		values = append(values, param{"company_domain", *p.CompanyDomain})
	}
	if p.CompanyName != nil {
		values = append(values, param{"company_name", *p.CompanyName})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EmailVerifyParams holds the parameters of EmailVerify (GET / POST /v2/email/verify).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EmailVerifyParams struct {
	// The email address you want to verify.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EmailVerifyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EuropagesCompanyParams holds the parameters of EuropagesCompany (GET / POST /v2/europages/company).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EuropagesCompanyParams struct {
	// Profile URL copied from a search card, or a slug.
	Query string `json:"query" api:"query,required"`
	// Site when query is a slug. Same values as search. Default fr. Ignored when query is a URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EuropagesCompanyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EuropagesSearchParams holds the parameters of EuropagesSearch (GET / POST /v2/europages/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EuropagesSearchParams struct {
	// Keyword, or a category / search URL on Europages.
	Query string `json:"query" api:"query,required"`
	// Site of a keyword: fr (default), de, co.uk (en, uk, gb), es, it, nl, and the other Europages TLDs. Ignored when query is a URL. German uses /suche/. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Page from 1. Page 2 and after use /pg-N/. 30 companies per page.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EuropagesSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EventbriteEventParams holds the parameters of EventbriteEvent (GET / POST /v2/eventbrite/event).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EventbriteEventParams struct {
	// Event detail URL or numeric id.
	Query string `json:"query" api:"query,required"`
	// Country TLD when query is a bare id (fr, com, es, com.mx, …). Ignored if query is a full URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EventbriteEventParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// EventbriteSearchParams holds the parameters of EventbriteSearch (GET / POST /v2/eventbrite/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type EventbriteSearchParams struct {
	// Keyword, or paste a full Eventbrite city / category / search URL.
	Query *string `json:"query" api:"query"`
	// Eventbrite TLD when query is a keyword (fr, com, es, com.mx, co.uk, …). Ignored if query is a full URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Location slug such as france--paris (keyword searches only).
	Location *string `json:"location" api:"location"`
	// Page number.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *EventbriteSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ExpediaAdParams holds the parameters of ExpediaAd (GET / POST /v2/expedia/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ExpediaAdParams struct {
	// An Expedia hotel information URL. The response is one flat hotel. hotel_id comes from .h{id}.Hotel-Information when that pattern is in the URL, and is null otherwise. A destination, a non-Expedia URL, or an empty value returns 400.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ExpediaAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ExpediaFlightParams holds the parameters of ExpediaFlight (GET / POST /v2/expedia/flight).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ExpediaFlightParams struct {
	// A route or a full Expedia Flights-Search URL. Empty, a route without two airport codes, or a non-Expedia URL returns 400. A full URL is opened as given. price stays the label on the offer. pagination.next is null.
	Query string `json:"query" api:"query,required"`
	// Departure as YYYY-MM-DD for a route. Omit it and the route uses 2026-11-15. A value that is not YYYY-MM-DD returns 400. Ignored when query is a full URL.
	DepartureDate *string `json:"departure_date" api:"departure_date"`
	// Return date as YYYY-MM-DD. Omit it for a one-way trip. A value that is not YYYY-MM-DD returns 400. Ignored when query is a full URL.
	ReturnDate *string `json:"return_date" api:"return_date"`
	// Adult passengers for a route. Omit it and the request uses 1. Ignored when query is a full URL.
	Adults *float64 `json:"adults" api:"adults"`
	// Label copied onto every result as cabin_class. Omit it and each result uses economy. It does not add a cabin filter to the Expedia URL.
	CabinClass *string `json:"cabin_class" api:"cabin_class"`
	// Host for a route. Omit it for www.expedia.com. Allowed: us, ca, uk, au, fr, de, es, it, nl, mx. Any other value returns 400. Ignored when query is a full URL.
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ExpediaFlightParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 9)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.DepartureDate != nil {
		values = append(values, param{"departure_date", *p.DepartureDate})
	}
	if p.ReturnDate != nil {
		values = append(values, param{"return_date", *p.ReturnDate})
	}
	if p.Adults != nil {
		values = append(values, param{"adults", *p.Adults})
	}
	if p.CabinClass != nil {
		values = append(values, param{"cabin_class", *p.CabinClass})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ExpediaSearchParams holds the parameters of ExpediaSearch (GET / POST /v2/expedia/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ExpediaSearchParams struct {
	// A valid Expedia hotel search URL. Copy the URL directly from your Expedia search results page.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ExpediaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FashionphileProductParams holds the parameters of FashionphileProduct (GET / POST /v2/fashionphile/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FashionphileProductParams struct {
	// Product handle or a fashionphile.com product URL. The request reads products/{handle}.js and returns one flat product. price is the USD amount and price_display is that amount rebuilt as a dollar string.
	Query string `json:"query" api:"query,required"`
	// Default false. true replaces the product object with result, the raw Shopify JSON text.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FashionphileProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FashionphileSearchParams holds the parameters of FashionphileSearch (GET / POST /v2/fashionphile/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FashionphileSearchParams struct {
	// Keywords, a Fashionphile search URL, a collection URL, or a products.json URL. Empty returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number starting at 1. Omit it for page 1.
	Page *float64 `json:"page" api:"page"`
	// Default false. true replaces results and pagination with result, the raw Shopify JSON text.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FashionphileSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderChewyProductsParams holds the parameters of FinderChewyProducts (GET / POST /v2/finder/chewy-products).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderChewyProductsParams struct {
	// Free-text search across product titles and IDs. Optional if at least one filter is provided.
	Search *string `json:"search" api:"search"`
	// Legacy alias of search. Kept for backward compatibility.
	Query *string `json:"query" api:"query"`
	// Filter by one or several internal product IDs. Provide a comma-separated list to match multiple IDs.
	ID *string `json:"id" api:"id"`
	// Filter by website domain. Comma-separated values are accepted.
	Domain *string `json:"domain" api:"domain"`
	// Raw filter expression for advanced filtering. Combined with the shorthand filters using AND.
	Filter *string `json:"filter" api:"filter"`
	// Maximum number of results to return. Default is 10, maximum is 100. Default: 10.
	Limit *float64 `json:"limit" api:"limit"`
	// Number of results to skip for pagination. Default is 0. Default: 0.
	Offset *float64 `json:"offset" api:"offset"`
	// Sort the results. Provide one or several attributes with :asc or :desc, comma-separated.
	Sort *string `json:"sort" api:"sort"`
	// Comma-separated list of fields to return. Defaults to all fields.
	Fields *string `json:"fields" api:"fields"`
	// When true, the response is wrapped in an object containing the total estimated count plus the hits: { count, hits }. When false or omitted, returns the raw hits array. Default: false.
	Count *bool `json:"count" api:"count"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderChewyProductsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Search != nil {
		values = append(values, param{"search", *p.Search})
	}
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.ID != nil {
		values = append(values, param{"id", *p.ID})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Filter != nil {
		values = append(values, param{"filter", *p.Filter})
	}
	if p.Limit != nil {
		values = append(values, param{"limit", *p.Limit})
	}
	if p.Offset != nil {
		values = append(values, param{"offset", *p.Offset})
	}
	if p.Sort != nil {
		values = append(values, param{"sort", *p.Sort})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.Count != nil {
		values = append(values, param{"count", *p.Count})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderEleclercStoresParams holds the parameters of FinderEleclercStores (GET / POST /v2/finder/eleclerc-stores).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderEleclercStoresParams struct {
	// Filters stores based on a search string. The search is non-case-sensitive and returns all stores whose names contain the specified string.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderEleclercStoresParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderGoogleCountriesParams holds the parameters of FinderGoogleCountries (GET / POST /v2/finder/google-countries).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderGoogleCountriesParams struct {
	// Filters countries based on a search string. The search is non-case-sensitive and returns all countries whose names contain the specified string.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderGoogleCountriesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderGoogleLanguagesParams holds the parameters of FinderGoogleLanguages (GET / POST /v2/finder/google-languages).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderGoogleLanguagesParams struct {
	// Filters languages based on a search string. The search is non-case-sensitive and returns all languages whose names contain the specified string.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderGoogleLanguagesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderGoogleLocationsParams holds the parameters of FinderGoogleLocations (GET / POST /v2/finder/google-locations).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderGoogleLocationsParams struct {
	// Restricts your search to locations that contain the supplied string.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderGoogleLocationsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderLeroymerlinProductsParams holds the parameters of FinderLeroymerlinProducts (GET / POST /v2/finder/leroymerlin-products).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderLeroymerlinProductsParams struct {
	// Free-text search across product titles and IDs. Optional if at least one filter is provided.
	Search *string `json:"search" api:"search"`
	// Legacy alias of search. Kept for backward compatibility.
	Query *string `json:"query" api:"query"`
	// Filter by one or several internal product IDs. Provide a comma-separated list to match multiple IDs.
	ID *string `json:"id" api:"id"`
	// Filter by website domain. Comma-separated values are accepted.
	Domain *string `json:"domain" api:"domain"`
	// Raw filter expression for advanced filtering. Combined with the shorthand filters using AND.
	Filter *string `json:"filter" api:"filter"`
	// Maximum number of results to return. Default is 10, maximum is 100. Default: 10.
	Limit *float64 `json:"limit" api:"limit"`
	// Number of results to skip for pagination. Default is 0. Default: 0.
	Offset *float64 `json:"offset" api:"offset"`
	// Sort the results. Provide one or several attributes with :asc or :desc, comma-separated.
	Sort *string `json:"sort" api:"sort"`
	// Comma-separated list of fields to return. Defaults to all fields.
	Fields *string `json:"fields" api:"fields"`
	// When true, the response is wrapped in an object containing the total estimated count plus the hits: { count, hits }. When false or omitted, returns the raw hits array. Default: false.
	Count *bool `json:"count" api:"count"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderLeroymerlinProductsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Search != nil {
		values = append(values, param{"search", *p.Search})
	}
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.ID != nil {
		values = append(values, param{"id", *p.ID})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Filter != nil {
		values = append(values, param{"filter", *p.Filter})
	}
	if p.Limit != nil {
		values = append(values, param{"limit", *p.Limit})
	}
	if p.Offset != nil {
		values = append(values, param{"offset", *p.Offset})
	}
	if p.Sort != nil {
		values = append(values, param{"sort", *p.Sort})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.Count != nil {
		values = append(values, param{"count", *p.Count})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderLeroymerlinStoresParams holds the parameters of FinderLeroymerlinStores (GET / POST /v2/finder/leroymerlin-stores).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderLeroymerlinStoresParams struct {
	// Filters Leroy Merlin stores based on a search string. The search is non-case-sensitive and matches store names, cities, postal codes and store slugs.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderLeroymerlinStoresParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderLinkedinIndustriesParams holds the parameters of FinderLinkedinIndustries (GET / POST /v2/finder/linkedin-industries).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderLinkedinIndustriesParams struct {
	// Restricts your search to industries that contain the supplied string.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderLinkedinIndustriesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderMrbricolageStoresParams holds the parameters of FinderMrbricolageStores (GET / POST /v2/finder/mrbricolage-stores).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FinderMrbricolageStoresParams struct {
	// Filters Mr.Bricolage stores based on a search string. The search is non-case-sensitive and matches store names, cities and postal codes.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderMrbricolageStoresParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FiverrProfileInfoParams holds the parameters of FiverrProfileInfo (GET / POST /v2/fiverr/profile/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FiverrProfileInfoParams struct {
	// Fiverr profile URL of the seller you want to extract information from.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FiverrProfileInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FiverrSearchParams holds the parameters of FiverrSearch (GET / POST /v2/fiverr/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FiverrSearchParams struct {
	// A Fiverr search URL for gigs.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FiverrSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FnacProductParams holds the parameters of FnacProduct (GET / POST /v2/fnac/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FnacProductParams struct {
	// Product URL (/a{id} or /mp{id}) or the numeric Fnac id. A keyword or a search URL returns 400.
	Query string `json:"query" api:"query,required"`
	// When true, returns the raw HTML instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FnacProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FnacSearchParams holds the parameters of FnacSearch (GET / POST /v2/fnac/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FnacSearchParams struct {
	// Keyword or a Fnac SearchResult URL. A product URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number, from 1. Default is 1. Page 2 is sent as PageIndex=2.
	Page *float64 `json:"page" api:"page"`
	// When true, returns the raw HTML instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FnacSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FotocasaAdParams holds the parameters of FotocasaAd (GET / POST /v2/fotocasa/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FotocasaAdParams struct {
	// Full Fotocasa ad URL (/comprar/ or /alquiler/, ending with /d).
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FotocasaAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FotocasaSearchParams holds the parameters of FotocasaSearch (GET / POST /v2/fotocasa/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FotocasaSearchParams struct {
	// City name or full Fotocasa search URL.
	Query string `json:"query" api:"query,required"`
	// Used when query is a city name (ignored for full URLs). Default: "buy".
	Operation *string `json:"operation" api:"operation"`
	// Page number (1-based). Default: 1. Page 2 maps to /l/2. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FotocasaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Operation != nil {
		values = append(values, param{"operation", *p.Operation})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FranchisedirectFranchiseParams holds the parameters of FranchisedirectFranchise (GET / POST /v2/franchisedirect/franchise).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FranchisedirectFranchiseParams struct {
	// Required. A Franchise Direct profile URL from search, or the same path beginning with /. A keyword or a listing URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FranchisedirectFranchiseParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FranchisedirectSearchParams holds the parameters of FranchisedirectSearch (GET / POST /v2/franchisedirect/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FranchisedirectSearchParams struct {
	// Required. A keyword, or a Franchise Direct listing URL (search, directory, or category). A profile URL such as /coffee-franchises/coffeenow returns 400 on this endpoint. Send that URL to /v2/franchisedirect/franchise.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FranchisedirectSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FressnapfProductParams holds the parameters of FressnapfProduct (GET / POST /v2/fressnapf/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FressnapfProductParams struct {
	// Product URL on www.fressnapf.de.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML. The string false stays false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FressnapfProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FressnapfSearchParams holds the parameters of FressnapfSearch (GET / POST /v2/fressnapf/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FressnapfSearchParams struct {
	// Keyword or a fressnapf.de search or category URL.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Example: 2.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML. The string false stays false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FressnapfSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FundaPropertyParams holds the parameters of FundaProperty (GET / POST /v2/funda/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FundaPropertyParams struct {
	// Listing URL on www.funda.nl, under /detail/koop or /detail/huur.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FundaPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FundaSearchParams holds the parameters of FundaSearch (GET / POST /v2/funda/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type FundaSearchParams struct {
	// City, area, or a Funda search URL.
	Query string `json:"query" api:"query,required"`
	// Sale or rent when query is a place.
	Offering *string `json:"offering" api:"offering"`
	// Page from 1. Example: 2. Funda receives this as search_result=2. Page 1 and page 2 return different listings.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FundaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Offering != nil {
		values = append(values, param{"offering", *p.Offering})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// G2ProductInfoParams holds the parameters of G2ProductInfo (GET / POST /v2/g2/product/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type G2ProductInfoParams struct {
	// Product slug or a www.g2.com product URL.
	Query string `json:"query" api:"query,required"`
	// Keep reviews of one star rating, from 1 to 5. Omit it to read the unfiltered reviews page.
	NpsScore *float64 `json:"nps_score" api:"nps_score"`
	// Return the rendered HTML instead of the parsed product.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *G2ProductInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.NpsScore != nil {
		values = append(values, param{"nps_score", *p.NpsScore})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GithubUserInfoParams holds the parameters of GithubUserInfo (GET / POST /v2/github/user/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GithubUserInfoParams struct {
	// A GitHub user identifier. Multiple formats are accepted:
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GithubUserInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GlassdoorCompanyInfoParams holds the parameters of GlassdoorCompanyInfo (GET / POST /v2/glassdoor/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GlassdoorCompanyInfoParams struct {
	// Company URL or employer id.\n\nExample:\n\n 1692882
	Query string `json:"query" api:"query,required"`
	// Glassdoor TLD when query is not a full URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GlassdoorCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GlassdoorCompanySearchParams holds the parameters of GlassdoorCompanySearch (GET / POST /v2/glassdoor/company/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GlassdoorCompanySearchParams struct {
	// Employer keyword, or a full Glassdoor autocomplete URL.\n\nExample:\n\n airbus
	Query string `json:"query" api:"query,required"`
	// Glassdoor TLD (fr, com, co.uk, …). Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GlassdoorCompanySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GlassdoorJobInfoParams holds the parameters of GlassdoorJobInfo (GET / POST /v2/glassdoor/job/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GlassdoorJobInfoParams struct {
	// Job listing URL or jl id.\n\nExample:\n\n 1010283388161
	Query string `json:"query" api:"query,required"`
	// Glassdoor TLD when query is not a full URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GlassdoorJobInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GlassdoorJobSearchParams holds the parameters of GlassdoorJobSearch (GET / POST /v2/glassdoor/job/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GlassdoorJobSearchParams struct {
	// Keyword or full Glassdoor job search URL.
	Query string `json:"query" api:"query,required"`
	// Glassdoor TLD when query is not a full URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Page number.
	Page *float64 `json:"page" api:"page"`
	// Glassdoor sortBy value (e.g. date_desc).
	SortBy *string `json:"sort_by" api:"sort_by"`
	// When true, set remoteWorkType=1.
	Remote *bool `json:"remote" api:"remote"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GlassdoorJobSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.SortBy != nil {
		values = append(values, param{"sort_by", *p.SortBy})
	}
	if p.Remote != nil {
		values = append(values, param{"remote", *p.Remote})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoodreadsBookParams holds the parameters of GoodreadsBook (GET / POST /v2/goodreads/book).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoodreadsBookParams struct {
	// Required. The numeric Goodreads book id from search, or a /book/show/ URL. A title is not accepted.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoodreadsBookParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoodreadsSearchParams holds the parameters of GoodreadsSearch (GET / POST /v2/goodreads/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoodreadsSearchParams struct {
	// Required. A book keyword, or a Goodreads search URL (the path contains /search or the query contains q=).
	Query string `json:"query" api:"query,required"`
	// Optional. Page number, starting at 1. Omitted or blank uses page 1. A value below 1 returns 400. When query is already a Goodreads search URL, this number replaces the page parameter in that URL. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoodreadsSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleImagesParams holds the parameters of GoogleImages (GET / POST /v2/google/images).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleImagesParams struct {
	// The image search query to execute on Google Images.
	Query string `json:"query" api:"query,required"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Page number for pagination. Images are returned in batches of 100 per page. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleImagesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleJobsParams holds the parameters of GoogleJobs (GET / POST /v2/google/jobs).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleJobsParams struct {
	// Job search term or keyword for querying Google job listings.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleJobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleMapsParams holds the parameters of GoogleMaps (GET / POST /v2/google/maps).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleMapsParams struct {
	// Business name or keyword to search for relevant data.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleMapsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleNewsParams holds the parameters of GoogleNews (GET / POST /v2/google/news).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleNewsParams struct {
	// The news search query to execute on Google News. Supports standard Google search operators.
	Query string `json:"query" api:"query,required"`
	// City or region from which the news search should originate. Simulates a search performed from the specified location.
	Location *string `json:"location" api:"location"`
	// Google-encoded location string for precise geographic targeting. Use this instead of location when you need exact control.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Page number of news results to retrieve. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Maximum number of results to return per page. Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleNewsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 13)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchParams holds the parameters of GoogleSearch (GET / POST /v2/google/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleSearchParams struct {
	// The search query to execute on Google. Supports all standard Google search operators.
	Query string `json:"query" api:"query,required"`
	// Advanced time-range and filter string for Google Search. Used to restrict results by recency.
	Tbs *string `json:"tbs" api:"tbs"`
	// City or region from which the search should originate. Simulates a real user's search from the specified location.
	Location *string `json:"location" api:"location"`
	// Google-encoded location string for precise geographic targeting.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Page number of search results to retrieve. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Maximum number of results to return per page (e.g. 10, 40, 100). Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 15)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Tbs != nil {
		values = append(values, param{"tbs", *p.Tbs})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchAIParams holds the parameters of GoogleSearchAI (GET / POST /v2/google/search/ai).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleSearchAIParams struct {
	// The question or search query to send to Google AI Mode. Natural-language questions work best; advanced operators are also accepted.
	Query string `json:"query" api:"query,required"`
	// City or region from which the AI Mode search should originate. Simulates a real user searching from that location.
	Location *string `json:"location" api:"location"`
	// Google-encoded location string for precise geographic targeting. Mutually exclusive with location.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchAIParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchAutocompleteParams holds the parameters of GoogleSearchAutocomplete (GET / POST /v2/google/search/autocomplete).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleSearchAutocompleteParams struct {
	// The partial or complete search query for which you want to retrieve autocomplete suggestions from Google.
	Query string `json:"query" api:"query,required"`
	// Cursor position in the query string where autocomplete suggestions should be generated. Determines which part of the query is being completed. Default: 0.
	Cp *float64 `json:"cp" api:"cp"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchAutocompleteParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 10)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Cp != nil {
		values = append(values, param{"cp", *p.Cp})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleShoppingParams holds the parameters of GoogleShopping (GET / POST /v2/google/shopping).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleShoppingParams struct {
	// Product search query on Google Shopping. Use a product name, brand, or category.
	Query string `json:"query" api:"query,required"`
	// Page number of Shopping results. Each page returns about 20 products. There is no num parameter. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// City or region used to localize Shopping results. Simulates a search from that location and generates the UULE parameter automatically.
	Location *string `json:"location" api:"location"`
	// Google-encoded location string (UULE) for precise geo-targeting. Use location instead if you do not already have an encoded value.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleShoppingParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 13)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleVideosParams holds the parameters of GoogleVideos (GET / POST /v2/google/videos).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleVideosParams struct {
	// The video search query to execute on Google Videos. Supports standard Google search operators.
	Query string `json:"query" api:"query,required"`
	// City or region from which the video search should originate.
	Location *string `json:"location" api:"location"`
	// Google-encoded location string for precise geographic targeting.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Page number of video results to retrieve. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Maximum number of video results to return per page. Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleVideosParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 14)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleadstransparencyAdParams holds the parameters of GoogleadstransparencyAd (GET / POST /v2/googleadstransparency/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleadstransparencyAdParams struct {
	// Creative URL from search.
	Query string `json:"query" api:"query,required"`
	// ISO alpha-2. The matching region is returned first.
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleadstransparencyAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleadstransparencySearchParams holds the parameters of GoogleadstransparencySearch (GET / POST /v2/googleadstransparency/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GoogleadstransparencySearchParams struct {
	// Domain name, advertiser name, advertiser id, or a full Ads Transparency URL.
	Query string `json:"query" api:"query,required"`
	// Country. ISO alpha-2 (FR), anywhere, or a Transparency region code (2250 for France, 250 is also accepted).
	Region *string `json:"region" api:"region"`
	// Domain filter. Used when query is not already a domain.
	Domain *string `json:"domain" api:"domain"`
	// Advertiser id, AR followed by digits.
	AdvertiserID *string `json:"advertiser_id" api:"advertiser_id"`
	// Creative format.
	Format *string `json:"format" api:"format"`
	// Comma-separated surfaces.
	Platform *string `json:"platform" api:"platform"`
	// Relative window.
	PresetDate *string `json:"preset_date" api:"preset_date"`
	// Inclusive start date, YYYY-MM-DD or YYYYMMDD.
	StartDate *string `json:"start_date" api:"start_date"`
	// Inclusive end date, YYYY-MM-DD or YYYYMMDD. A single day uses the same value as start_date.
	EndDate *string `json:"end_date" api:"end_date"`
	// Ad topic. Default: commercial. Default: "commercial".
	Topic *string `json:"topic" api:"topic"`
	// Page number, from 1 to 20. Default: 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Page size, from 1 to 100. Default: 40. Default: 40.
	PerPage *float64 `json:"per_page" api:"per_page"`
	// Opaque cursor copied from pagination.next.
	Cursor *string `json:"cursor" api:"cursor"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleadstransparencySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 16)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.AdvertiserID != nil {
		values = append(values, param{"advertiser_id", *p.AdvertiserID})
	}
	if p.Format != nil {
		values = append(values, param{"format", *p.Format})
	}
	if p.Platform != nil {
		values = append(values, param{"platform", *p.Platform})
	}
	if p.PresetDate != nil {
		values = append(values, param{"preset_date", *p.PresetDate})
	}
	if p.StartDate != nil {
		values = append(values, param{"start_date", *p.StartDate})
	}
	if p.EndDate != nil {
		values = append(values, param{"end_date", *p.EndDate})
	}
	if p.Topic != nil {
		values = append(values, param{"topic", *p.Topic})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.PerPage != nil {
		values = append(values, param{"per_page", *p.PerPage})
	}
	if p.Cursor != nil {
		values = append(values, param{"cursor", *p.Cursor})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GraingerProductParams holds the parameters of GraingerProduct (GET / POST /v2/grainger/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GraingerProductParams struct {
	// SKU, full Grainger product URL, or relative /product/... path.
	Query string `json:"query" api:"query,required"`
	// Grainger country when query is not a full URL (default: com).
	Region *string `json:"region" api:"region"`
	// When true, return rendered HTML as {"result": "..."} instead of JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GraingerProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GraingerSearchParams holds the parameters of GraingerSearch (GET / POST /v2/grainger/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GraingerSearchParams struct {
	// Keyword, full Grainger search/category URL, category path, or SKU.
	Query string `json:"query" api:"query,required"`
	// Grainger country when query is not a full URL (default: com).
	Region *string `json:"region" api:"region"`
	// When true, return rendered HTML as {"result": "..."} instead of JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GraingerSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GreenhouseJobParams holds the parameters of GreenhouseJob (GET / POST /v2/greenhouse/job).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GreenhouseJobParams struct {
	// Required. board/id or a job URL. Example: discord/8806482002 or https://job-boards.greenhouse.io/discord/jobs/8806482002.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GreenhouseJobParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GreenhouseSearchParams holds the parameters of GreenhouseSearch (GET / POST /v2/greenhouse/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GreenhouseSearchParams struct {
	// Required. Board token or a job-boards.greenhouse.io board URL. Example: discord. A job URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Optional. Case-insensitive filter on title and location. Example: policy. Omit it to list the whole board.
	Keyword *string `json:"keyword" api:"keyword"`
	// Optional. Page number from 1. Example: 2. Defaults to 1. Page size is 20.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GreenhouseSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Keyword != nil {
		values = append(values, param{"keyword", *p.Keyword})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GsaelibraryContractParams holds the parameters of GsaelibraryContract (GET / POST /v2/gsaelibrary/contract).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GsaelibraryContractParams struct {
	// Required. The contractorInfo.do URL from a search contract row. A contract number is also accepted, but GS-07F-0239T alone currently returns 404 because GSA needs the contractor name that search includes in the URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GsaelibraryContractParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GsaelibrarySearchParams holds the parameters of GsaelibrarySearch (GET / POST /v2/gsaelibrary/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type GsaelibrarySearchParams struct {
	// Required. A keyword, or a full search URL on gsaelibrary.gsa.gov. The URL must contain searchResults.do or searchText=.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GsaelibrarySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// HomestraAdParams holds the parameters of HomestraAd (GET / POST /v2/homestra/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type HomestraAdParams struct {
	// A Homestra property URL or slug.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *HomestraAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// HomestraSearchParams holds the parameters of HomestraSearch (GET / POST /v2/homestra/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type HomestraSearchParams struct {
	// A valid Homestra search URL with optional filters.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *HomestraSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// HuggingfaceModelParams holds the parameters of HuggingfaceModel (GET / POST /v2/huggingface/model).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type HuggingfaceModelParams struct {
	// Required. A public model id (owner/name) or a huggingface.co model URL. Reserved site paths such as models, datasets, or spaces return 400.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *HuggingfaceModelParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// HuggingfaceSearchParams holds the parameters of HuggingfaceSearch (GET / POST /v2/huggingface/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type HuggingfaceSearchParams struct {
	// Required. A model keyword, a huggingface.co/models URL, or the full pagination.next cursor URL from a previous page.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *HuggingfaceSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IdealistaPropertyParams holds the parameters of IdealistaProperty (GET / POST /v2/idealista/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IdealistaPropertyParams struct {
	// Full Idealista property detail URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IdealistaPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IdealistaProvincesParams holds the parameters of IdealistaProvinces (GET / POST /v2/idealista/provinces).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IdealistaProvincesParams struct {
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IdealistaProvincesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 3)
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IdealistaSearchParams holds the parameters of IdealistaSearch (GET / POST /v2/idealista/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IdealistaSearchParams struct {
	// Full Idealista search results URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IdealistaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IdealoProductParams holds the parameters of IdealoProduct (GET / POST /v2/idealo/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IdealoProductParams struct {
	// Full Idealo product URL, or a numeric product id.
	Query string `json:"query" api:"query,required"`
	// Region when query is a numeric product id.
	Country *string `json:"country" api:"country"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IdealoProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Country != nil {
		values = append(values, param{"country", *p.Country})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IdealoSearchParams holds the parameters of IdealoSearch (GET / POST /v2/idealo/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IdealoSearchParams struct {
	// Keyword or full Idealo search URL for Spain, France, or Germany.
	Query string `json:"query" api:"query,required"`
	// Region when query is a keyword.
	Country *string `json:"country" api:"country"`
	// Search page. 2 appends ?page=2.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IdealoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Country != nil {
		values = append(values, param{"country", *p.Country})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IkeaProductParams holds the parameters of IkeaProduct (GET / POST /v2/ikea/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IkeaProductParams struct {
	// Full IKEA product page URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IkeaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IkeaSearchParams holds the parameters of IkeaSearch (GET / POST /v2/ikea/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IkeaSearchParams struct {
	// Full IKEA search URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IkeaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// Immobilienscout24PropertyParams holds the parameters of Immobilienscout24Property (GET / POST /v2/immobilienscout24/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type Immobilienscout24PropertyParams struct {
	// The expose id, or the expose page URL. Default: "168052908".
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML in result instead of the parsed expose.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *Immobilienscout24PropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// Immobilienscout24SearchParams holds the parameters of Immobilienscout24Search (GET / POST /v2/immobilienscout24/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type Immobilienscout24SearchParams struct {
	// Where to search. Default: "berlin/berlin".
	Query string `json:"query" api:"query,required"`
	// Rent or buy.
	RealEstateType *string `json:"real_estate_type" api:"real_estate_type"`
	// Page number, starting at 1.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML in result instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *Immobilienscout24SearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.RealEstateType != nil {
		values = append(values, param{"real_estate_type", *p.RealEstateType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// Immoscout24chSearchParams holds the parameters of Immoscout24chSearch (GET / POST /v2/immoscout24ch/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type Immoscout24chSearchParams struct {
	// Swiss city or a full ImmoScout24.ch search URL.
	Query string `json:"query" api:"query,required"`
	// rent or buy.
	OfferType *string `json:"offer_type" api:"offer_type"`
	// apartment, house, or all.
	PropertyType *string `json:"property_type" api:"property_type"`
	// Page from 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *Immoscout24chSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.OfferType != nil {
		values = append(values, param{"offer_type", *p.OfferType})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ImmoweltPropertyParams holds the parameters of ImmoweltProperty (GET / POST /v2/immowelt/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ImmoweltPropertyParams struct {
	// Full Immowelt property detail URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ImmoweltPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ImmoweltSearchParams holds the parameters of ImmoweltSearch (GET / POST /v2/immowelt/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ImmoweltSearchParams struct {
	// Full Immowelt search URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ImmoweltSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ImoveiswebPropertyParams holds the parameters of ImoveiswebProperty (GET / POST /v2/imoveisweb/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ImoveiswebPropertyParams struct {
	// Listing URL from search.
	Query string `json:"query" api:"query,required"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ImoveiswebPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ImoveiswebSearchParams holds the parameters of ImoveiswebSearch (GET / POST /v2/imoveisweb/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ImoveiswebSearchParams struct {
	// Catalog name or an ImoveisWeb results URL. This is Imobiliária Aquários, not ZonaProp.
	Query string `json:"query" api:"query,required"`
	// Optional catalog when query is not already a catalog name.
	PropertyType *string `json:"property_type" api:"property_type"`
	// Page from 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ImoveiswebSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IndeedCompanyInfoParams holds the parameters of IndeedCompanyInfo (GET / POST /v2/indeed/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IndeedCompanyInfoParams struct {
	// Indeed company slug or full /cmp/ URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {\"result\": \"...\"} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IndeedCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IndeedCompanySearchParams holds the parameters of IndeedCompanySearch (GET / POST /v2/indeed/company/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IndeedCompanySearchParams struct {
	// Company keyword or full Indeed companies search URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {\"result\": \"...\"} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IndeedCompanySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IndeedJobInfoParams holds the parameters of IndeedJobInfo (GET / POST /v2/indeed/job/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IndeedJobInfoParams struct {
	// Indeed job URL, sponsored SERP click URL or 16-char jk job key.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {\"result\": \"...\"} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IndeedJobInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IndeedJobSalaryParams holds the parameters of IndeedJobSalary (GET / POST /v2/indeed/job/salary).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IndeedJobSalaryParams struct {
	// Career salary slug, full /career/.../salaries URL or slug|location pair.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {\"result\": \"...\"} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IndeedJobSalaryParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IndeedJobSearchParams holds the parameters of IndeedJobSearch (GET / POST /v2/indeed/job/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IndeedJobSearchParams struct {
	// Keyword, Indeed /jobs URL or SEO search URL on fr.indeed.com (default domain).
	Query string `json:"query" api:"query,required"`
	// 1-based page number (10 results per page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML as {\"result\": \"...\"} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IndeedJobSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// InetProductParams holds the parameters of InetProduct (GET / POST /v2/inet/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type InetProductParams struct {
	// Article number or product URL. Example: 4306287. A search link returns 400.
	Query string `json:"query" api:"query,required"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *InetProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// InetSearchParams holds the parameters of InetSearch (GET / POST /v2/inet/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type InetSearchParams struct {
	// What to look for. A word such as ssd, or an Inet search or category link. A product link returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Page 2 uses page=2.
	Page *float64 `json:"page" api:"page"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *InetSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// InstagramPostInfoParams holds the parameters of InstagramPostInfo (GET / POST /v2/instagram/post/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type InstagramPostInfoParams struct {
	// The Instagram post or reel to retrieve. Accepts a full URL or a shortcode.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *InstagramPostInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// InstagramUserInfoParams holds the parameters of InstagramUserInfo (GET / POST /v2/instagram/user/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type InstagramUserInfoParams struct {
	// The Instagram user to look up. Accepts a full profile URL or a username.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *InstagramUserInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IntermarcheProductParams holds the parameters of IntermarcheProduct (GET / POST /v2/intermarche/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IntermarcheProductParams struct {
	// Full Intermarché product URL, or an EAN barcode.
	Query string `json:"query" api:"query,required"`
	// Intermarché store id. Resolve a city with the [Stores Finder](https://www.piloterr.com/library/finder-intermarche-stores) (for example query=Paris).
	ShopID string `json:"shop_id" api:"shop_id,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IntermarcheProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ShopID != "" {
		values = append(values, param{"shop_id", p.ShopID})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IntermarcheSearchParams holds the parameters of IntermarcheSearch (GET / POST /v2/intermarche/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IntermarcheSearchParams struct {
	// Search keyword or full Intermarché search URL.
	Query string `json:"query" api:"query,required"`
	// Intermarché store id. Resolve a city with the [Stores Finder](https://www.piloterr.com/library/finder-intermarche-stores) (for example query=Paris).
	ShopID string `json:"shop_id" api:"shop_id,required"`
	// Page number starting at 1. Use pagination.total_pages from a previous response to iterate.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IntermarcheSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ShopID != "" {
		values = append(values, param{"shop_id", p.ShopID})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IntersportProductParams holds the parameters of IntersportProduct (GET / POST /v2/intersport/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IntersportProductParams struct {
	// Product URL on intersport.fr, or a root-relative path ending in -p-{id}.
	Query string `json:"query" api:"query,required"`
	// When true, returns the raw HTML instead of the parsed JSON. The response shape is {"result": "<html>..."}. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IntersportProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// IntersportSearchParams holds the parameters of IntersportSearch (GET / POST /v2/intersport/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type IntersportSearchParams struct {
	// Keyword or an intersport.fr listing URL. A product URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number, from 1. Default is 1. Intersport's page query is zero-based, so page 2 is sent as page=1.
	Page *float64 `json:"page" api:"page"`
	// When true, returns the raw HTML instead of the parsed JSON. The response shape is {"result": "<html>..."}. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *IntersportSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// JameseditionPropertyParams holds the parameters of JameseditionProperty (GET / POST /v2/jamesedition/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type JameseditionPropertyParams struct {
	// Required. JamesEdition listing URL, for example https://www.jamesedition.com/real_estate/munich-germany/alpine-panorama-above-lake-starnberg-new-build-project-near-munich-17793477. A keyword is rejected.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *JameseditionPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// JameseditionSearchParams holds the parameters of JameseditionSearch (GET / POST /v2/jamesedition/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type JameseditionSearchParams struct {
	// Required. Keyword, for example villa, or a JamesEdition results URL such as https://www.jamesedition.com/real_estate/france. A single listing URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number starting at 1. Example: 2. Omitted means page 1. A page already present in a results URL is kept when this is omitted or 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *JameseditionSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// JobsParams holds the parameters of Jobs (GET /v1/jobs/{id}).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type JobsParams struct {
	// The job_id returned by an endpoint called with is_async=true.
	ID string `json:"id" api:"id,required"`
}

func (p *JobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 1)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	return values
}

// KickUserInfoParams holds the parameters of KickUserInfo (GET / POST /v2/kick/user/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type KickUserInfoParams struct {
	// The Kick channel to look up. Accepts a full channel URL or a username/slug.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *KickUserInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// KijijiAdParams holds the parameters of KijijiAd (GET / POST /v2/kijiji/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type KijijiAdParams struct {
	// Full Kijiji listing URL or the numeric listing id.
	Query string `json:"query" api:"query,required"`
	// Locale sent to Kijiji as the siteLocale cookie. en sets en-CA and prints $20. fr sets fr-CA and prints 20 $ on the same ad. Default: "en".
	Language *string `json:"language" api:"language"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *KijijiAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Language != nil {
		values = append(values, param{"language", *p.Language})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// KijijiSearchParams holds the parameters of KijijiSearch (GET / POST /v2/kijiji/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type KijijiSearchParams struct {
	// Keyword or full Kijiji results URL.
	Query string `json:"query" api:"query,required"`
	// Results page, starting at 1. Page 2 of the City of Toronto bikes search returns a different set of listing ids. When page stays at 1, a page-N segment already present in the URL is kept.
	Page *float64 `json:"page" api:"page"`
	// Locale sent to Kijiji as the siteLocale cookie. en sets en-CA. fr sets fr-CA. French cards print prices like 10,00 $. Default: "en".
	Language *string `json:"language" api:"language"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *KijijiSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Language != nil {
		values = append(values, param{"language", *p.Language})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// KrogerProductParams holds the parameters of KrogerProduct (GET / POST /v2/kroger/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type KrogerProductParams struct {
	// Kroger product URL or UPC.
	Query string `json:"query" api:"query,required"`
	// Ignored. Prices stay on the default Rolla store, 02100537 (Kroger, 605 W 4th St, Rolla, MO 65401). The response location_id is the store that was actually priced.
	LocationID *string `json:"location_id" api:"location_id"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *KrogerProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.LocationID != nil {
		values = append(values, param{"location_id", *p.LocationID})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// KrogerSearchParams holds the parameters of KrogerSearch (GET / POST /v2/kroger/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type KrogerSearchParams struct {
	// Keyword or a full Kroger search URL.
	Query string `json:"query" api:"query,required"`
	// Page number, 1-based.
	Page *float64 `json:"page" api:"page"`
	// Ignored. Prices stay on the default Rolla store, 02100537 (Kroger, 605 W 4th St, Rolla, MO 65401). The response location_id is the store that was actually priced.
	LocationID *string `json:"location_id" api:"location_id"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *KrogerSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.LocationID != nil {
		values = append(values, param{"location_id", *p.LocationID})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LamudiAdParams holds the parameters of LamudiAd (GET / POST /v2/lamudi/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LamudiAdParams struct {
	// Full Lamudi listing URL. Accepts lamudi.com.ph (/property/) and lamudi.co.id (/properti/).
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LamudiAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LamudiSearchParams holds the parameters of LamudiSearch (GET / POST /v2/lamudi/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LamudiSearchParams struct {
	// Full Lamudi search URL. Used first when set. Location and path filters are ignored.
	Query *string `json:"query" api:"query"`
	// City or district. Used when query is empty. Builds a Philippines URL (lamudi.com.ph). For Indonesia, pass a search URL in query.
	Location *string `json:"location" api:"location"`
	// Transaction type. Used when query is empty. Default: "sale".
	Operation *string `json:"operation" api:"operation"`
	// Property type. Used when query is empty. Maps to Lamudi slugs (condominium → /condo/, land → /lot/).
	PropertyType *string `json:"property_type" api:"property_type"`
	// Minimum asking price in PHP or IDR.
	MinPrice *float64 `json:"min_price" api:"min_price"`
	// Maximum asking price in PHP or IDR.
	MaxPrice *float64 `json:"max_price" api:"max_price"`
	// Minimum bedroom count.
	MinBedrooms *float64 `json:"min_bedrooms" api:"min_bedrooms"`
	// Maximum bedroom count.
	MaxBedrooms *float64 `json:"max_bedrooms" api:"max_bedrooms"`
	// Minimum bathroom count.
	MinBathrooms *float64 `json:"min_bathrooms" api:"min_bathrooms"`
	// Minimum area in m².
	MinArea *float64 `json:"min_area" api:"min_area"`
	// Maximum area in m².
	MaxArea *float64 `json:"max_area" api:"max_area"`
	// Sort order. Default: "relevance".
	Sort *string `json:"sort" api:"sort"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LamudiSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 15)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Operation != nil {
		values = append(values, param{"operation", *p.Operation})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.MinPrice != nil {
		values = append(values, param{"min_price", *p.MinPrice})
	}
	if p.MaxPrice != nil {
		values = append(values, param{"max_price", *p.MaxPrice})
	}
	if p.MinBedrooms != nil {
		values = append(values, param{"min_bedrooms", *p.MinBedrooms})
	}
	if p.MaxBedrooms != nil {
		values = append(values, param{"max_bedrooms", *p.MaxBedrooms})
	}
	if p.MinBathrooms != nil {
		values = append(values, param{"min_bathrooms", *p.MinBathrooms})
	}
	if p.MinArea != nil {
		values = append(values, param{"min_area", *p.MinArea})
	}
	if p.MaxArea != nil {
		values = append(values, param{"max_area", *p.MaxArea})
	}
	if p.Sort != nil {
		values = append(values, param{"sort", *p.Sort})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LeboncoinAdParams holds the parameters of LeboncoinAd (GET / POST /v2/leboncoin/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LeboncoinAdParams struct {
	// Leboncoin ad ID or full ad URL. Both formats are accepted regardless of return_page_source.
	Query string `json:"query" api:"query,required"`
	// Controls the response format. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LeboncoinAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LeboncoinSearchParams holds the parameters of LeboncoinSearch (GET / POST /v2/leboncoin/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LeboncoinSearchParams struct {
	// Full Leboncoin search or category URL. Both keyword searches and category browse URLs are supported.
	Query string `json:"query" api:"query,required"`
	// Deprecated. Sending true returns 400. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LeboncoinSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LeboncoinSearchAPIParams holds the parameters of LeboncoinSearchAPI (POST /v2/leboncoin/search_api).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LeboncoinSearchAPIParams struct {
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LeboncoinSearchAPIParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 3)
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LenovoProductParams holds the parameters of LenovoProduct (GET / POST /v2/lenovo/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LenovoProductParams struct {
	// Full Lenovo product URL (must contain /p/). Bare part numbers are rejected.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LenovoProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LenovoSearchParams holds the parameters of LenovoSearch (GET / POST /v2/lenovo/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LenovoSearchParams struct {
	// Keyword or full Lenovo search / category URL.
	Query string `json:"query" api:"query,required"`
	// Store locale when query is a keyword (ignored for full URLs). Default: "fr/fr".
	Locale *string `json:"locale" api:"locale"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LenovoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Locale != nil {
		values = append(values, param{"locale", *p.Locale})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LeroymerlinProductParams holds the parameters of LeroymerlinProduct (GET / POST /v2/leroymerlin/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LeroymerlinProductParams struct {
	// Full Leroy Merlin product URL, or an identifier (SKU / EAN / text).
	Query string `json:"query" api:"query,required"`
	// Leroy Merlin country when query is a keyword or identifier (default: fr). Ignored when query is a full Leroy Merlin URL (country is detected from the domain).
	Region *string `json:"region" api:"region"`
	// Optional store id for local offer context (prices, stock, deliveries). Must exist for the active region. Returns 400 when unknown.
	StoreID *string `json:"store_id" api:"store_id"`
	// When set to true, returns the raw HTML of the product page instead of the parsed JSON payload. Useful for debugging or custom parsing. The response shape is {"return": "<html>..."}. Defaults to false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LeroymerlinProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.StoreID != nil {
		values = append(values, param{"store_id", *p.StoreID})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LeroymerlinSearchParams holds the parameters of LeroymerlinSearch (GET / POST /v2/leroymerlin/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LeroymerlinSearchParams struct {
	// Search keyword, SKU/EAN, or full Leroy Merlin search URL.
	Query string `json:"query" api:"query,required"`
	// Leroy Merlin country when query is a keyword or identifier (default: fr). Ignored when query is a full Leroy Merlin URL (country is detected from the domain).
	Region *string `json:"region" api:"region"`
	// 1-based page number for pagination. Defaults to 1 when omitted. Use the pagination.total_pages returned by a previous call to iterate.
	Page *float64 `json:"page" api:"page"`
	// Optional store id for local prices and availability. Must exist for the active region (FR stores with region=fr, RO stores with region=ro). Returns 400 when unknown.
	StoreID *string `json:"store_id" api:"store_id"`
	// When set to true, returns the raw HTML of the page instead of the parsed JSON payload. Useful for debugging or custom parsing. The response shape is {"return": "<html>..."}. Defaults to false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LeroymerlinSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.StoreID != nil {
		values = append(values, param{"store_id", *p.StoreID})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LidlProductParams holds the parameters of LidlProduct (GET / POST /v2/lidl/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LidlProductParams struct {
	// Lidl product URL or ERP id.
	Query string `json:"query" api:"query,required"`
	// Target country when query is an ERP id only (default: fr). Ignored for full product URLs.
	Region *string `json:"region" api:"region"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LidlProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LidlSearchParams holds the parameters of LidlSearch (GET / POST /v2/lidl/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LidlSearchParams struct {
	// Search keyword or full Lidl search URL.
	Query string `json:"query" api:"query,required"`
	// Target country when query is a keyword (default: fr).
	Region *string `json:"region" api:"region"`
	// 1-based page number (default: 1). Lidl paginates by offset (48 items per page).
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LidlSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinCompanyInfoParams holds the parameters of LinkedinCompanyInfo (GET / POST /v2/linkedin/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinCompanyInfoParams struct {
	// LinkedIn company identifier when you already know the page.
	Query *string `json:"query" api:"query"`
	// Company website domain for reverse lookup when you do not have the LinkedIn URL.
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinJobCountParams holds the parameters of LinkedinJobCount (GET / POST /v2/linkedin/job/count).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinJobCountParams struct {
	// Job title or keyword to count matching postings.
	Keyword *string `json:"keyword" api:"keyword"`
	// Filter by seniority level.
	ExperienceLevel *string `json:"experience_level" api:"experience_level"`
	// Filter by employment type.
	JobType *string `json:"job_type" api:"job_type"`
	// Filter by posting recency.
	When *string `json:"when" api:"when"`
	// Filter by work arrangement.
	Flexibility *string `json:"flexibility" api:"flexibility"`
	// Maximum distance radius (in miles) from the target location (requires geo_id).
	Distance *float64 `json:"distance" api:"distance"`
	// LinkedIn geo ID of the target location.
	GeoID *string `json:"geo_id" api:"geo_id"`
	// LinkedIn numeric company ID to restrict the count to jobs posted by a specific company.
	CompanyID *string `json:"company_id" api:"company_id"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinJobCountParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Keyword != nil {
		values = append(values, param{"keyword", *p.Keyword})
	}
	if p.ExperienceLevel != nil {
		values = append(values, param{"experience_level", *p.ExperienceLevel})
	}
	if p.JobType != nil {
		values = append(values, param{"job_type", *p.JobType})
	}
	if p.When != nil {
		values = append(values, param{"when", *p.When})
	}
	if p.Flexibility != nil {
		values = append(values, param{"flexibility", *p.Flexibility})
	}
	if p.Distance != nil {
		values = append(values, param{"distance", *p.Distance})
	}
	if p.GeoID != nil {
		values = append(values, param{"geo_id", *p.GeoID})
	}
	if p.CompanyID != nil {
		values = append(values, param{"company_id", *p.CompanyID})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinJobInfoParams holds the parameters of LinkedinJobInfo (GET / POST /v2/linkedin/job/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinJobInfoParams struct {
	// The LinkedIn job posting to retrieve. Accepts a numeric job ID or a full job posting URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinJobInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinJobSearchParams holds the parameters of LinkedinJobSearch (GET / POST /v2/linkedin/job/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinJobSearchParams struct {
	// Job title or keyword to search for.
	Keyword *string `json:"keyword" api:"keyword"`
	// Filter by seniority level.
	ExperienceLevel *string `json:"experience_level" api:"experience_level"`
	// Filter by employment type.
	JobType *string `json:"job_type" api:"job_type"`
	// Filter by posting date recency.
	When *string `json:"when" api:"when"`
	// Filter by work arrangement.
	Flexibility *string `json:"flexibility" api:"flexibility"`
	// Maximum radius in miles from the target location (requires geo_id).
	Distance *float64 `json:"distance" api:"distance"`
	// LinkedIn geo ID for the target location. Default is 92000000 (worldwide). Default: "92000000".
	GeoID *string `json:"geo_id" api:"geo_id"`
	// LinkedIn numeric company ID to filter results to jobs posted by a specific company.
	CompanyID *string `json:"company_id" api:"company_id"`
	// Page number for paginating through results. Default is 1. Each page returns up to 25 job listings. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinJobSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Keyword != nil {
		values = append(values, param{"keyword", *p.Keyword})
	}
	if p.ExperienceLevel != nil {
		values = append(values, param{"experience_level", *p.ExperienceLevel})
	}
	if p.JobType != nil {
		values = append(values, param{"job_type", *p.JobType})
	}
	if p.When != nil {
		values = append(values, param{"when", *p.When})
	}
	if p.Flexibility != nil {
		values = append(values, param{"flexibility", *p.Flexibility})
	}
	if p.Distance != nil {
		values = append(values, param{"distance", *p.Distance})
	}
	if p.GeoID != nil {
		values = append(values, param{"geo_id", *p.GeoID})
	}
	if p.CompanyID != nil {
		values = append(values, param{"company_id", *p.CompanyID})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinJobSuggestParams holds the parameters of LinkedinJobSuggest (GET / POST /v2/linkedin/job/suggest).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinJobSuggestParams struct {
	// The keyword prefix to search for job title suggestions or location names.
	Query string `json:"query" api:"query,required"`
	// Set to true to receive location (geo) suggestions instead of keyword suggestions. Default: false.
	Location *bool `json:"location" api:"location"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinJobSuggestParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinPostInfoParams holds the parameters of LinkedinPostInfo (GET / POST /v2/linkedin/post/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinPostInfoParams struct {
	// The LinkedIn post to retrieve. Accepts a numeric post ID or a full post URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinPostInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinProductInfoParams holds the parameters of LinkedinProductInfo (GET / POST /v2/linkedin/product/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinProductInfoParams struct {
	// The full LinkedIn product page URL to retrieve.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinProductInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinProfileInfoParams holds the parameters of LinkedinProfileInfo (GET / POST /v2/linkedin/profile/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinProfileInfoParams struct {
	// The LinkedIn profile to retrieve. Accepts a username, full profile URL, or numeric member ID.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinProfileInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinadlibraryAdParams holds the parameters of LinkedinadlibraryAd (GET / POST /v2/linkedinadlibrary/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinadlibraryAdParams struct {
	// Ad id or detail URL from search.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinadlibraryAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LinkedinadlibrarySearchParams holds the parameters of LinkedinadlibrarySearch (GET / POST /v2/linkedinadlibrary/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LinkedinadlibrarySearchParams struct {
	// Full LinkedIn Ad Library search URL, or a keyword when account_owner is empty.
	Query *string `json:"query" api:"query"`
	// Advertiser / company name filter (accountOwner).
	AccountOwner *string `json:"account_owner" api:"account_owner"`
	// Keyword filter on ad copy.
	Keyword *string `json:"keyword" api:"keyword"`
	// Comma-separated ISO country codes, or ALL.
	Countries *string `json:"countries" api:"countries"`
	// Date window: last-30-days, current-month, current-year, last-year, custom-date-range.
	DateOption *string `json:"date_option" api:"date_option"`
	// Start date YYYY-MM-DD when date_option=custom-date-range.
	StartDate *string `json:"start_date" api:"start_date"`
	// End date YYYY-MM-DD when date_option=custom-date-range.
	EndDate *string `json:"end_date" api:"end_date"`
	// Opaque cursor from pagination.cursor / pagination.next.
	Cursor *string `json:"cursor" api:"cursor"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LinkedinadlibrarySearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.AccountOwner != nil {
		values = append(values, param{"account_owner", *p.AccountOwner})
	}
	if p.Keyword != nil {
		values = append(values, param{"keyword", *p.Keyword})
	}
	if p.Countries != nil {
		values = append(values, param{"countries", *p.Countries})
	}
	if p.DateOption != nil {
		values = append(values, param{"date_option", *p.DateOption})
	}
	if p.StartDate != nil {
		values = append(values, param{"start_date", *p.StartDate})
	}
	if p.EndDate != nil {
		values = append(values, param{"end_date", *p.EndDate})
	}
	if p.Cursor != nil {
		values = append(values, param{"cursor", *p.Cursor})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LoopnetPropertyParams holds the parameters of LoopnetProperty (GET / POST /v2/loopnet/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LoopnetPropertyParams struct {
	// Required. LoopNet listing URL from search, for example https://www.loopnet.com/Listing/484-Amsterdam-Ave-New-York-NY/31379002/. A keyword is rejected.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LoopnetPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// LoopnetSearchParams holds the parameters of LoopnetSearch (GET / POST /v2/loopnet/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type LoopnetSearchParams struct {
	// Required. City and state, for example New York, NY, or a LoopNet search URL such as https://www.loopnet.com/search/commercial-real-estate/new-york-ny/for-sale/. A listing URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number starting at 1. Example: 2. Omitted means page 1. A page already present in a search URL is kept when this is omitted or 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *LoopnetSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MaerskTrackingParams holds the parameters of MaerskTracking (GET / POST /v2/maersk/tracking).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MaerskTrackingParams struct {
	// Container / BOL tracking number, or full Maersk tracking URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the rendered HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MaerskTrackingParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ManomanoProductParams holds the parameters of ManomanoProduct (GET / POST /v2/manomano/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ManomanoProductParams struct {
	// Full ManoMano product page URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ManomanoProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ManomanoProductOfferParams holds the parameters of ManomanoProductOffer (GET / POST /v2/manomano/product/offer).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ManomanoProductOfferParams struct {
	// Numeric ManoMano product id from the product endpoint.
	Query string `json:"query" api:"query,required"`
	// ManoMano country. FR, GB, DE, ES, or IT. Default FR. Default: "FR".
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ManomanoProductOfferParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ManomanoSearchParams holds the parameters of ManomanoSearch (GET / POST /v2/manomano/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ManomanoSearchParams struct {
	// ManoMano category URL or keyword. A keyword is opened on manomano.fr.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ManomanoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MascusAdParams holds the parameters of MascusAd (GET / POST /v2/mascus/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MascusAdParams struct {
	// Full Mascus listing URL. Usually taken from results[].url in search.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MascusAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MascusSearchParams holds the parameters of MascusSearch (GET / POST /v2/mascus/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MascusSearchParams struct {
	// Full Mascus category or search URL. The country site comes from the host (mascus.be, mascus.nl, mascus.de, mascus.com, …).
	Query string `json:"query" api:"query,required"`
	// Page number (default: 1). About 40 results per page.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MascusSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MediamarktProductParams holds the parameters of MediamarktProduct (GET / POST /v2/mediamarkt/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MediamarktProductParams struct {
	// Product URL or article id.  https://www.mediamarkt.ch/fr/product/_apple-iphone-15-2203371.html
	Query string `json:"query" api:"query,required"`
	// Shop code when query is an article id. Example: ch-fr. Default ch-fr.
	Region *string `json:"region" api:"region"`
	// When true, return raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MediamarktProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MediamarktSearchParams holds the parameters of MediamarktSearch (GET / POST /v2/mediamarkt/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MediamarktSearchParams struct {
	// Keyword or MediaMarkt URL. Example: iphone
	Query string `json:"query" api:"query,required"`
	// Shop code. Example: ch-fr or de. Default ch-fr. Ignored when query is a full URL.
	Region *string `json:"region" api:"region"`
	// Page number from 1. Example: 1
	Page *float64 `json:"page" api:"page"`
	// When true, return raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MediamarktSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MicromaniaProductParams holds the parameters of MicromaniaProduct (GET / POST /v2/micromania/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MicromaniaProductParams struct {
	// Product URL or the numeric Micromania id. A keyword or a search URL returns 400.
	Query string `json:"query" api:"query,required"`
	// When true, returns the raw HTML instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MicromaniaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MicromaniaSearchParams holds the parameters of MicromaniaSearch (GET / POST /v2/micromania/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MicromaniaSearchParams struct {
	// Keyword or a Micromania Search-Show URL. A product URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number, from 1. Default is 1. Page 2 is sent as page=2.
	Page *float64 `json:"page" api:"page"`
	// When true, returns the raw HTML instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MicromaniaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MobileAdParams holds the parameters of MobileAd (GET / POST /v2/mobile/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MobileAdParams struct {
	// Mobile ad URL or numeric ad id.
	Query string `json:"query" api:"query,required"`
	// Language for this ad. The same numeric id stays reachable for fr, es, de, and en.
	Locale *string `json:"locale" api:"locale"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MobileAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Locale != nil {
		values = append(values, param{"locale", *p.Locale})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MobileSearchParams holds the parameters of MobileSearch (GET / POST /v2/mobile/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MobileSearchParams struct {
	// Search URL or make/model code (ms=).
	Query string `json:"query" api:"query,required"`
	// Locale label for the results. The cars themselves stay the same for de, fr, and en.
	Locale *string `json:"locale" api:"locale"`
	// Page number, starting at 1. Same effect as page or pageNumber in the Mobile.de URL.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MobileSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Locale != nil {
		values = append(values, param{"locale", *p.Locale})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MoreleProductParams holds the parameters of MoreleProduct (GET / POST /v2/morele/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MoreleProductParams struct {
	// Morele product URL, or a numeric product id.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MoreleProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MoreleSearchParams holds the parameters of MoreleSearch (GET / POST /v2/morele/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MoreleSearchParams struct {
	// Keyword or a full Morele search URL. A product URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Page number starting at 1. Page 2 asks Morele for the next screen of the same search.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MoreleSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// MrbricolageSuggestParams holds the parameters of MrbricolageSuggest (GET / POST /v2/mrbricolage/suggest).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type MrbricolageSuggestParams struct {
	// Search keyword or EAN used to query the Mr.Bricolage live suggest endpoint. Accepts product names, partial titles or numeric identifiers such as EAN barcodes.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MrbricolageSuggestParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// NeweggProductParams holds the parameters of NeweggProduct (GET / POST /v2/newegg/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type NeweggProductParams struct {
	// Item number, catalog id, or product URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML. The string false stays false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *NeweggProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// NeweggSearchParams holds the parameters of NeweggSearch (GET / POST /v2/newegg/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type NeweggSearchParams struct {
	// Keyword or a newegg.com search URL.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Example: 2.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML. The string false stays false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *NeweggSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// NorthdataCompanyParams holds the parameters of NorthdataCompany (GET / POST /v2/northdata/company).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type NorthdataCompanyParams struct {
	// Company name, North Data id, or a profile URL.
	Query string `json:"query" api:"query,required"`
	// ISO alpha-2, used only when query is a name. Same form as search: FR or FR,DE.
	Countries *string `json:"countries" api:"countries"`
	// fr (default), com or en, de. Ignored when query is a URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *NorthdataCompanyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Countries != nil {
		values = append(values, param{"countries", *p.Countries})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// NorthdataSearchParams holds the parameters of NorthdataSearch (GET / POST /v2/northdata/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type NorthdataSearchParams struct {
	// Company or person name, or a North Data suggest URL.
	Query string `json:"query" api:"query,required"`
	// ISO alpha-2 codes, separated by comma, pipe or space (FR or FR,DE). Empty searches every country. A URL that already has countries wins.
	Countries *string `json:"countries" api:"countries"`
	// Host of a name: fr (default), com or en, de. Ignored when query is a URL. Default: "fr".
	Domain *string `json:"domain" api:"domain"`
	// Only 1. Below 1 returns 400. Above 1 returns 404. At most 15 rows.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *NorthdataSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Countries != nil {
		values = append(values, param{"countries", *p.Countries})
	}
	if p.Domain != nil {
		values = append(values, param{"domain", *p.Domain})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// NPIProviderParams holds the parameters of NPIProvider (GET / POST /v2/npi/provider).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type NPIProviderParams struct {
	// Required. 10-digit NPI, example 1972813608, or a provider-view URL such as https://npiregistry.cms.hhs.gov/provider-view/1972813608.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *NPIProviderParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// NPISearchParams holds the parameters of NPISearch (GET / POST /v2/npi/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type NPISearchParams struct {
	// Optional when another criterion is set. Example: Anthony Fauci. Also a 10-digit NPI such as 1972813608, or an NPPES search URL. With enumeration_type=NPI-2 this text is the organization name, for example Mayo Clinic. A provider-view URL is rejected here.
	Query *string `json:"query" api:"query"`
	// Optional. Example: Anthony. Used as the given name when you do not want query to be split.
	FirstName *string `json:"first_name" api:"first_name"`
	// Optional. Example: Fauci. A single-word query such as Smith is used as the last name.
	LastName *string `json:"last_name" api:"last_name"`
	// Optional. Example: Mayo Clinic. You can also pass that name in query with enumeration_type=NPI-2.
	OrganizationName *string `json:"organization_name" api:"organization_name"`
	// Optional. Practice city. Example: Washington.
	City *string `json:"city" api:"city"`
	// Optional. Two-letter state code. Example: DC. Any other value is rejected.
	State *string `json:"state" api:"state"`
	// Optional. Postal code. Example: 20016.
	PostalCode *string `json:"postal_code" api:"postal_code"`
	// Optional. Taxonomy description. Example: Internal Medicine.
	Taxonomy *string `json:"taxonomy" api:"taxonomy"`
	// Optional. NPI-1 for a person or NPI-2 for an organization. Example: NPI-2. NPI-2 treats a text query as an organization name.
	EnumerationType *string `json:"enumeration_type" api:"enumeration_type"`
	// Optional. Page number, starting at 1. Example: 2. Defaults to 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Page size. Example: 10. Defaults to 10 and cannot exceed 200. next is a URL only when the page contains this many rows.
	Limit *float64 `json:"limit" api:"limit"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *NPISearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 14)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.FirstName != nil {
		values = append(values, param{"first_name", *p.FirstName})
	}
	if p.LastName != nil {
		values = append(values, param{"last_name", *p.LastName})
	}
	if p.OrganizationName != nil {
		values = append(values, param{"organization_name", *p.OrganizationName})
	}
	if p.City != nil {
		values = append(values, param{"city", *p.City})
	}
	if p.State != nil {
		values = append(values, param{"state", *p.State})
	}
	if p.PostalCode != nil {
		values = append(values, param{"postal_code", *p.PostalCode})
	}
	if p.Taxonomy != nil {
		values = append(values, param{"taxonomy", *p.Taxonomy})
	}
	if p.EnumerationType != nil {
		values = append(values, param{"enumeration_type", *p.EnumerationType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Limit != nil {
		values = append(values, param{"limit", *p.Limit})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OceanioSearchParams holds the parameters of OceanioSearch (GET / POST /v2/oceanio/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OceanioSearchParams struct {
	// Domain or company name.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OceanioSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OlxAdParams holds the parameters of OlxAd (GET / POST /v2/olx/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OlxAdParams struct {
	// OLX listing URL.
	Query string `json:"query" api:"query,required"`
	// Ignored when query is a full URL. The host selects the country.
	Region *string `json:"region" api:"region"`
	// When true, return raw HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OlxAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OlxProductParams holds the parameters of OlxProduct (GET / POST /v2/olx/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OlxProductParams struct {
	// Full URL of the OLX listing page to retrieve. Must be a valid olx.com.br listing URL. Example 1:
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OlxProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OlxSearchParams holds the parameters of OlxSearch (GET / POST /v2/olx/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OlxSearchParams struct {
	// Search keyword or phrase, as you would type it in the OLX search bar. Supports Portuguese terms for Brazilian inventory.
	Query string `json:"query" api:"query,required"`
	// OLX category slug to restrict results to a specific vertical. Common values:
	Category *string `json:"category" api:"category"`
	// Brazilian state abbreviation to filter results geographically.
	State *string `json:"state" api:"state"`
	// City slug to narrow results to a specific municipality. Use lowercase with hyphens.
	City *string `json:"city" api:"city"`
	// Minimum listing price in BRL.
	MinPrice *float64 `json:"min_price" api:"min_price"`
	// Maximum listing price in BRL.
	MaxPrice *float64 `json:"max_price" api:"max_price"`
	// Filter by seller type.
	SellerType *string `json:"seller_type" api:"seller_type"`
	// Sort order for results. Default: "relevance".
	Sort *string `json:"sort" api:"sort"`
	// Page number for paginated results. Default: 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OlxSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Category != nil {
		values = append(values, param{"category", *p.Category})
	}
	if p.State != nil {
		values = append(values, param{"state", *p.State})
	}
	if p.City != nil {
		values = append(values, param{"city", *p.City})
	}
	if p.MinPrice != nil {
		values = append(values, param{"min_price", *p.MinPrice})
	}
	if p.MaxPrice != nil {
		values = append(values, param{"max_price", *p.MaxPrice})
	}
	if p.SellerType != nil {
		values = append(values, param{"seller_type", *p.SellerType})
	}
	if p.Sort != nil {
		values = append(values, param{"sort", *p.Sort})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OscaroProductParams holds the parameters of OscaroProduct (GET / POST /v2/oscaro/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OscaroProductParams struct {
	// Oscaro product page URL (*-{sku}-p suffix).
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OscaroProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OscaroSearchParams holds the parameters of OscaroSearch (GET / POST /v2/oscaro/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OscaroSearchParams struct {
	// Search keyword, search URL, or category listing URL (*-{id}-g).
	Query string `json:"query" api:"query,required"`
	// Target shop when query is a keyword (default: com).
	Region *string `json:"region" api:"region"`
	// 1-based page number override for category listing URLs (default: read from URL or 1). Uses ?page=2 on *-g pages.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OscaroSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OvokoAdParams holds the parameters of OvokoAd (GET / POST /v2/ovoko/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OvokoAdParams struct {
	// Full Ovoko listing URL. The last path segment must be a SKU slug (example: qbb17557-95c827460c-porsche-macan-...) or a known product prefix.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OvokoAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OvokoSearchParams holds the parameters of OvokoSearch (GET / POST /v2/ovoko/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OvokoSearchParams struct {
	// Full Ovoko search URL. Only search pages are accepted (not category, catalog, or offer pages).
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OvokoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OwlerCompanyInfoParams holds the parameters of OwlerCompanyInfo (GET / POST /v2/owler/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OwlerCompanyInfoParams struct {
	// The company to look up on Owler. Multiple formats are accepted:
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OwlerCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// OwlerSearchParams holds the parameters of OwlerSearch (GET / POST /v2/owler/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type OwlerSearchParams struct {
	// The company name or keyword to search for on Owler.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *OwlerSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PagesjaunesPageInfoParams holds the parameters of PagesjaunesPageInfo (GET / POST /v2/pagesjaunes/page/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PagesjaunesPageInfoParams struct {
	// A PagesJaunes business page URL or business ID.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PagesjaunesPageInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PagesjaunesSearchParams holds the parameters of PagesjaunesSearch (GET / POST /v2/pagesjaunes/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PagesjaunesSearchParams struct {
	// A valid PagesJaunes search URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PagesjaunesSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PatreonProfileInfoParams holds the parameters of PatreonProfileInfo (GET / POST /v2/patreon/profile/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PatreonProfileInfoParams struct {
	// Patreon profile URL or creator vanity name.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PatreonProfileInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PatreonSearchParams holds the parameters of PatreonSearch (GET / POST /v2/patreon/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PatreonSearchParams struct {
	// Search keyword to find Patreon creators.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PatreonSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PetsmartProductParams holds the parameters of PetsmartProduct (GET / POST /v2/petsmart/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PetsmartProductParams struct {
	// A PetSmart product link.
	Query string `json:"query" api:"query,required"`
	// Set to true to get the raw HTML in result instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PetsmartProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PetsmartSearchParams holds the parameters of PetsmartSearch (GET / POST /v2/petsmart/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PetsmartSearchParams struct {
	// What to look for. A word, or a PetSmart search link.
	Query string `json:"query" api:"query,required"`
	// Which page of results, starting at 1. Page 2 returns a different set of products.
	Page *float64 `json:"page" api:"page"`
	// Set to true to get the raw HTML in result instead of the parsed JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PetsmartSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PicardProductParams holds the parameters of PicardProduct (GET / POST /v2/picard/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PicardProductParams struct {
	// Picard product URL or numeric reference.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PicardProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PicardSearchParams holds the parameters of PicardSearch (GET / POST /v2/picard/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PicardSearchParams struct {
	// Search keyword or full Picard search URL.
	Query string `json:"query" api:"query,required"`
	// 1-based page number (24 results per page via SFCC Search-UpdateGrid). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PicardSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PinterestUserInfoParams holds the parameters of PinterestUserInfo (GET / POST /v2/pinterest/user/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PinterestUserInfoParams struct {
	// The Pinterest profile to retrieve. Accepts a full profile URL or a username.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PinterestUserInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PitchbookCompanyInfoParams holds the parameters of PitchbookCompanyInfo (GET / POST /v2/pitchbook/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PitchbookCompanyInfoParams struct {
	// PitchBook company id or company profile URL.
	Query string `json:"query" api:"query,required"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PitchbookCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PitchbookSearchParams holds the parameters of PitchbookSearch (GET / POST /v2/pitchbook/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PitchbookSearchParams struct {
	// Keyword or PitchBook search URL.
	Query string `json:"query" api:"query,required"`
	// Page number for content search (/search?q=...).
	Page *string `json:"page" api:"page"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PitchbookSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PointpProductParams holds the parameters of PointpProduct (GET / POST /v2/pointp/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PointpProductParams struct {
	// Full Point.P product page URL (/p/...-A{id}).
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of parsed JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PointpProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PointpSearchParams holds the parameters of PointpSearch (GET / POST /v2/pointp/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PointpSearchParams struct {
	// Search keyword or full Point.P search results URL.
	Query string `json:"query" api:"query,required"`
	// 1-based page number for pagination. Overrides the page inferred from the URL. Defaults to 1 when omitted.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PointpSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PrnewswireArticleParams holds the parameters of PrnewswireArticle (GET / POST /v2/prnewswire/article).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PrnewswireArticleParams struct {
	// Required. PR Newswire release URL, for example https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html. A keyword is rejected.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PrnewswireArticleParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// PrnewswireSearchParams holds the parameters of PrnewswireSearch (GET / POST /v2/prnewswire/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type PrnewswireSearchParams struct {
	// Required. Keyword, for example apple, or a PR Newswire news-search URL such as https://www.prnewswire.com/search/news/?keyword=apple. A release URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number starting at 1. Example: 2. Omitted means page 1. A page already present in a search URL is kept when this is omitted or 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *PrnewswireSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ProducthuntProductInfoParams holds the parameters of ProducthuntProductInfo (GET / POST /v2/producthunt/product/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ProducthuntProductInfoParams struct {
	// A ProductHunt product URL or product slug.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ProducthuntProductInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ProshopProductParams holds the parameters of ProshopProduct (GET / POST /v2/proshop/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ProshopProductParams struct {
	// Proshop product id, or a product URL. The id is the last path segment.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ProshopProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ProshopSearchParams holds the parameters of ProshopSearch (GET / POST /v2/proshop/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ProshopSearchParams struct {
	// Keyword or a full Proshop search URL with s=. A product URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Page number starting at 1. Proshop sends this as pn.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML as {"result": "..."}.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ProshopSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RappiProductParams holds the parameters of RappiProduct (GET / POST /v2/rappi/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RappiProductParams struct {
	// Required. Product URL on rappi.com.mx, example https://www.rappi.com.mx/p/santa-clara-leche-deslactosada-semidescremada-1-l-14069, or a numeric master id such as 14069.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RappiProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RappiSearchParams holds the parameters of RappiSearch (GET / POST /v2/rappi/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RappiSearchParams struct {
	// Required. Keyword, example leche, or a rappi.com.mx search URL such as https://www.rappi.com.mx/search?query=leche. A product URL or another country host is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Only page 1 returns rows. Example: 1. Defaults to 1. Page 2 is rejected because the payload has no next page.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RappiSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RatemyprofessorsProfessorParams holds the parameters of RatemyprofessorsProfessor (GET / POST /v2/ratemyprofessors/professor).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RatemyprofessorsProfessorParams struct {
	// Required. The numeric professor id from search, or a /professor/{id} URL. A name is not accepted.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RatemyprofessorsProfessorParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RatemyprofessorsSearchParams holds the parameters of RatemyprofessorsSearch (GET / POST /v2/ratemyprofessors/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RatemyprofessorsSearchParams struct {
	// Required. A professor name, or a Rate My Professors search URL whose path contains /search/professors or whose query contains q=.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RatemyprofessorsSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RealestatecomauSearchParams holds the parameters of RealestatecomauSearch (GET / POST /v2/realestatecomau/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RealestatecomauSearchParams struct {
	// Place or a realestate.com.au results URL.
	Query string `json:"query" api:"query,required"`
	// buy, rent, or sold. Used when query is a place.
	Channel *string `json:"channel" api:"channel"`
	// Page from 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RealestatecomauSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Channel != nil {
		values = append(values, param{"channel", *p.Channel})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RealtorPropertyParams holds the parameters of RealtorProperty (GET / POST /v2/realtor/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RealtorPropertyParams struct {
	// Full Realtor.com property detail URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RealtorPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RealtorSearchParams holds the parameters of RealtorSearch (GET / POST /v2/realtor/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RealtorSearchParams struct {
	// Full Realtor.com search results URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RealtorSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RedfinPropertyParams holds the parameters of RedfinProperty (GET / POST /v2/redfin/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RedfinPropertyParams struct {
	// Redfin home link. Use the url returned by search.
	Query string `json:"query" api:"query,required"`
	// Market id. Required only when the home link does not already include region_id.
	RegionID *float64 `json:"region_id" api:"region_id"`
	// Required with region_id when the home link has no region_type.
	RegionType *float64 `json:"region_type" api:"region_type"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RedfinPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.RegionID != nil {
		values = append(values, param{"region_id", *p.RegionID})
	}
	if p.RegionType != nil {
		values = append(values, param{"region_type", *p.RegionType})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RedfinSearchParams holds the parameters of RedfinSearch (GET / POST /v2/redfin/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RedfinSearchParams struct {
	// Redfin for-sale link for a city, neighborhood, or county.
	Query string `json:"query" api:"query,required"`
	// Page of the for-sale feed, from 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RedfinSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ReicheltProductParams holds the parameters of ReicheltProduct (GET / POST /v2/reichelt/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ReicheltProductParams struct {
	// A reichelt.de or reichelt.com product link. A search link returns 400.
	Query string `json:"query" api:"query,required"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ReicheltProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ReicheltSearchParams holds the parameters of ReicheltSearch (GET / POST /v2/reichelt/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ReicheltSearchParams struct {
	// What to look for. A word such as 1n4148, or a reichelt.de search or category link. A product link returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Page 2 is PAGE=1 on the shop.
	Page *float64 `json:"page" api:"page"`
	// Set to true to get the raw HTML in result.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ReicheltSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ReverbProductParams holds the parameters of ReverbProduct (GET / POST /v2/reverb/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ReverbProductParams struct {
	// Reverb listing URL or numeric listing id.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ReverbProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ReverbSearchParams holds the parameters of ReverbSearch (GET / POST /v2/reverb/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ReverbSearchParams struct {
	// Search keyword, query string, or full Reverb marketplace URL.
	Query string `json:"query" api:"query,required"`
	// 1-based page number override (default: read from URL or 1).
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ReverbSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RexelProductParams holds the parameters of RexelProduct (GET / POST /v2/rexel/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RexelProductParams struct {
	// Rexel France product URL containing /p/{id}.
	Query string `json:"query" api:"query,required"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RexelProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RexelSearchParams holds the parameters of RexelSearch (GET / POST /v2/rexel/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RexelSearchParams struct {
	// Keyword, EAN, or Rexel France search URL.
	Query string `json:"query" api:"query,required"`
	// Page number (1-based). Page 1 omits page in the URL.
	Page *string `json:"page" api:"page"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RexelSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RightmovePropertyParams holds the parameters of RightmoveProperty (GET / POST /v2/rightmove/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RightmovePropertyParams struct {
	// Rightmove property URL or the numeric id.
	Query string `json:"query" api:"query,required"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RightmovePropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RightmoveSearchParams holds the parameters of RightmoveSearch (GET / POST /v2/rightmove/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RightmoveSearchParams struct {
	// UK place name or a Rightmove search URL.
	Query string `json:"query" api:"query,required"`
	// Results page, from 1.
	Page *float64 `json:"page" api:"page"`
	// Optional. Default false. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RightmoveSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RsonlineProductParams holds the parameters of RsonlineProduct (GET / POST /v2/rsonline/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RsonlineProductParams struct {
	// A uk.rs-online.com product link. The response is one flat product. price is the numeric pound amount and price_display is the printed string. A search link, a bare stock number, or another host returns 400.
	Query string `json:"query" api:"query,required"`
	// Default false. true replaces the product object with result, the rendered HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RsonlineProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// RsonlineSearchParams holds the parameters of RsonlineSearch (GET / POST /v2/rsonline/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type RsonlineSearchParams struct {
	// A keyword, or a uk.rs-online.com search or category link. Empty, another host, or a product link returns 400. A keyword builds /web/c/?searchTerm=... and page 1 omits pn. Each card has price as a number and price_display as the pound string.
	Query string `json:"query" api:"query,required"`
	// Page number starting at 1. Omit it for page 1. Page 2 sets pn=2, changes pagination.page, and returns the next cards. pagination.next is that next URL, or null on the last page. A value below 1 returns 400.
	Page *float64 `json:"page" api:"page"`
	// Default false. true replaces results and pagination with result, the rendered HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *RsonlineSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SeekJobParams holds the parameters of SeekJob (GET / POST /v2/seek/job).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SeekJobParams struct {
	// Required. SEEK job URL or numeric id, for example 95568355 or https://www.seek.com.au/job/95568355. A keyword is rejected.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SeekJobParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SeekSearchParams holds the parameters of SeekSearch (GET / POST /v2/seek/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SeekSearchParams struct {
	// Required. Job keyword, for example nurse, or a SEEK results URL such as https://www.seek.com.au/nurse-jobs. A /job/<id> URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number starting at 1. Example: 2. Omitted means page 1. A page already present in a results URL is kept when this is omitted or 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SeekSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SelogerAgencyParams holds the parameters of SelogerAgency (GET / POST /v2/seloger/agency).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SelogerAgencyParams struct {
	// A valid SeLoger agency profile URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SelogerAgencyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SelogerAgencyRentParams holds the parameters of SelogerAgencyRent (GET / POST /v2/seloger/agency/rent).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SelogerAgencyRentParams struct {
	// A valid SeLoger agency ID.
	Query string `json:"query" api:"query,required"`
	// Page number for paginating through the agency's rental listings. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SelogerAgencyRentParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SelogerAgencySaleParams holds the parameters of SelogerAgencySale (GET / POST /v2/seloger/agency/sale).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SelogerAgencySaleParams struct {
	// A valid SeLoger agency ID.
	Query string `json:"query" api:"query,required"`
	// Page number for paginating through the agency's sale listings. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SelogerAgencySaleParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SelogerPropertyParams holds the parameters of SelogerProperty (GET / POST /v2/seloger/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SelogerPropertyParams struct {
	// A valid SeLoger property listing URL.
	Query string `json:"query" api:"query,required"`
	// Whether to return the raw HTML source of the property page instead of parsed JSON data. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SelogerPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SelogerSearchParams holds the parameters of SelogerSearch (GET / POST /v2/seloger/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SelogerSearchParams struct {
	// A valid SeLoger search results URL. Any SeLoger search page URL with filters is supported: location, property type, price range, surface, and number of rooms can all be embedded in the URL.
	Query string `json:"query" api:"query,required"`
	// Whether to return the raw HTML source of the search results page instead of parsed JSON data. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SelogerSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SephoraProductParams holds the parameters of SephoraProduct (GET / POST /v2/sephora/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SephoraProductParams struct {
	// A Sephora product URL. A product id such as 831955 also works.  https://www.sephora.fr/p/midnight-espresso---eau-de-parfum-au-cafe-831955.html
	Query string `json:"query" api:"query,required"`
	// Shop when query is only an id. Use fr for France or de for Germany. Leave empty for France.
	Region *string `json:"region" api:"region"`
	// Return the rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SephoraProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SephoraSearchParams holds the parameters of SephoraSearch (GET / POST /v2/sephora/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SephoraSearchParams struct {
	// A product word, for example mascara. Or paste a full Sephora search URL.
	Query string `json:"query" api:"query,required"`
	// Shop code. Use fr for France or de for Germany. Leave empty for France. A full URL picks the shop.
	Region *string `json:"region" api:"region"`
	// Page number, starting at 1. Use 2 for the next page of products.
	Page *float64 `json:"page" api:"page"`
	// Return the rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SephoraSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ShophouzzProductParams holds the parameters of ShophouzzProduct (GET / POST /v2/shophouzz/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ShophouzzProductParams struct {
	// Full URL of the Shophouzz product page to retrieve details for.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ShophouzzProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ShopifyAppsParams holds the parameters of ShopifyApps (GET / POST /v2/shopify/apps).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ShopifyAppsParams struct {
	// URL of the Shopify App Store listing to scrape.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ShopifyAppsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ShopifyCatalogParams holds the parameters of ShopifyCatalog (GET / POST /v2/shopify/catalog).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ShopifyCatalogParams struct {
	// Shop domain or a URL on that shop.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1.
	Page *float64 `json:"page" api:"page"`
	// How many items to return on this page.
	Limit *float64 `json:"limit" api:"limit"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ShopifyCatalogParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Limit != nil {
		values = append(values, param{"limit", *p.Limit})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ShopifyCollectionsParams holds the parameters of ShopifyCollections (GET / POST /v2/shopify/collections).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ShopifyCollectionsParams struct {
	// Shop domain or a URL on that shop.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1.
	Page *float64 `json:"page" api:"page"`
	// How many items to return on this page.
	Limit *float64 `json:"limit" api:"limit"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ShopifyCollectionsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Limit != nil {
		values = append(values, param{"limit", *p.Limit})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ShopifyProductParams holds the parameters of ShopifyProduct (GET / POST /v2/shopify/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ShopifyProductParams struct {
	// URL of a Shopify product page to scrape.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ShopifyProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SimilarwebDomainParams holds the parameters of SimilarwebDomain (GET / POST /v2/similarweb/domain).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SimilarwebDomainParams struct {
	// A domain name to retrieve traffic and engagement metrics for.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SimilarwebDomainParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SimilarwebSearchParams holds the parameters of SimilarwebSearch (GET / POST /v2/similarweb/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SimilarwebSearchParams struct {
	// Domain name or keyword to search for in the Similarweb database.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SimilarwebSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SnapedaProductParams holds the parameters of SnapedaProduct (GET / POST /v2/snapeda/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SnapedaProductParams struct {
	// SnapEDA part page URL or MPN/manufacturer slug.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw rendered HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SnapedaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// SnapedaSearchParams holds the parameters of SnapedaSearch (GET / POST /v2/snapeda/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type SnapedaSearchParams struct {
	// Search keyword or full SnapEDA / SnapMagic search URL.
	Query string `json:"query" api:"query,required"`
	// 1-based page number override (default: read from URL or 1). SnapEDA returns 20 results per page.
	Page *float64 `json:"page" api:"page"`
	// When true, return the raw rendered HTML as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SnapedaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StarbucksMenuParams holds the parameters of StarbucksMenu (GET / POST /v2/starbucks/menu).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StarbucksMenuParams struct {
	// Café number from Search, short (9230) or full (9230-154245).
	StoreNumber *string `json:"store_number" api:"store_number"`
	// Same café number, or a menu URL that contains storeNumber. Used when store_number is empty.
	Query *string `json:"query" api:"query"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StarbucksMenuParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.StoreNumber != nil {
		values = append(values, param{"store_number", *p.StoreNumber})
	}
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StarbucksProductParams holds the parameters of StarbucksProduct (GET / POST /v2/starbucks/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StarbucksProductParams struct {
	// Product and form, or a Starbucks product URL.
	Query string `json:"query" api:"query,required"`
	// Café number from Search, short (9230) or full (9230-154245). Required when the product URL has no storeNumber.
	StoreNumber *string `json:"store_number" api:"store_number"`
	// Form when query is only a product number. Usually hot or iced.
	FormCode *string `json:"form_code" api:"form_code"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StarbucksProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.StoreNumber != nil {
		values = append(values, param{"store_number", *p.StoreNumber})
	}
	if p.FormCode != nil {
		values = append(values, param{"form_code", *p.FormCode})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StarbucksSearchParams holds the parameters of StarbucksSearch (GET / POST /v2/starbucks/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StarbucksSearchParams struct {
	// Store-locator URL, or a lat,lng pair. Required when lat and lng are empty.
	Query *string `json:"query" api:"query"`
	// Latitude. Send it with lng. If both are set, they replace any coordinates found in query.
	Lat *float64 `json:"lat" api:"lat"`
	// Longitude. Send it with lat.
	Lng *float64 `json:"lng" api:"lng"`
	// Optional place label sent with the coordinates, for example New York, NY. Coordinates are still required.
	Place *string `json:"place" api:"place"`
	// Optional amenity filter. Separate codes with commas. ON keeps cafés that offer mobile order.
	Features *string `json:"features" api:"features"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StarbucksSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Lat != nil {
		values = append(values, param{"lat", *p.Lat})
	}
	if p.Lng != nil {
		values = append(values, param{"lng", *p.Lng})
	}
	if p.Place != nil {
		values = append(values, param{"place", *p.Place})
	}
	if p.Features != nil {
		values = append(values, param{"features", *p.Features})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StatsParams holds the parameters of Stats (GET /v2/stats).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StatsParams struct {
	// API family to check, for example website. Provide at least engine or path.
	Engine *string `json:"engine" api:"engine"`
	// Specific endpoint to check, for example /v2/website/crawler. If you also pass engine, use an endpoint from that API.
	Path *string `json:"path" api:"path"`
}

func (p *StatsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Engine != nil {
		values = append(values, param{"engine", *p.Engine})
	}
	if p.Path != nil {
		values = append(values, param{"path", *p.Path})
	}
	return values
}

// StockxProductParams holds the parameters of StockxProduct (GET / POST /v2/stockx/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StockxProductParams struct {
	// A StockX product slug.
	Query string `json:"query" api:"query,required"`
	// Country code for localized pricing and availability. Default: "US".
	Country *string `json:"country" api:"country"`
	// Currency code for price display. Default: "USD".
	Currency *string `json:"currency" api:"currency"`
	// Language code for localized content. Default: "en-US".
	Language *string `json:"language" api:"language"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StockxProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Country != nil {
		values = append(values, param{"country", *p.Country})
	}
	if p.Currency != nil {
		values = append(values, param{"currency", *p.Currency})
	}
	if p.Language != nil {
		values = append(values, param{"language", *p.Language})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StockxSearchParams holds the parameters of StockxSearch (GET / POST /v2/stockx/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StockxSearchParams struct {
	// Search query for StockX products. Can be a product name, brand, SKU, or slug.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StockxSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StockxTrendsParams holds the parameters of StockxTrends (GET / POST /v2/stockx/trends).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StockxTrendsParams struct {
	// StockX product category to retrieve trending items for.
	Query string `json:"query" api:"query,required"`
	// Sort order for trending results. Default: "featured".
	SortBy *string `json:"sort_by" api:"sort_by"`
	// Country code for localized trending data. Default: "US".
	Country *string `json:"country" api:"country"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StockxTrendsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.SortBy != nil {
		values = append(values, param{"sort_by", *p.SortBy})
	}
	if p.Country != nil {
		values = append(values, param{"country", *p.Country})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StubhubEventParams holds the parameters of StubhubEvent (GET / POST /v2/stubhub/event).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StubhubEventParams struct {
	// Required. StubHub event URL or numeric id, for example 162387399 or https://www.stubhub.com/harlem-globetrotters-marseille-tickets-3-31-2027/event/162387399/. A keyword is rejected.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StubhubEventParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StubhubSearchParams holds the parameters of StubhubSearch (GET / POST /v2/stubhub/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type StubhubSearchParams struct {
	// Required. Keyword, for example nba, or a StubHub search URL such as https://www.stubhub.com/secure/search/?q=nba. An event URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Only page 1 returns events. Example: 1. Page 2 and above return 404 because the search document has no further page. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *StubhubSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TargetProductParams holds the parameters of TargetProduct (GET / POST /v2/target/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TargetProductParams struct {
	// Target product URL or bare TCIN.
	Query string `json:"query" api:"query,required"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TargetProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TargetSearchParams holds the parameters of TargetSearch (GET / POST /v2/target/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TargetSearchParams struct {
	// Keyword or Target search URL.
	Query string `json:"query" api:"query,required"`
	// Page number (1-based).
	Page *string `json:"page" api:"page"`
	// When true, return rendered HTML instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TargetSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TdlrLicenseParams holds the parameters of TdlrLicense (GET / POST /v2/tdlr/license).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TdlrLicenseParams struct {
	// Required. Detail id, example AIRREF00000564, or the SearchResultDetail URL from search.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TdlrLicenseParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TdlrSearchParams holds the parameters of TdlrSearch (GET / POST /v2/tdlr/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TdlrSearchParams struct {
	// Optional when city, county, or postal_code is set. Person or business name, example SMITH. A detail URL is rejected.
	Query *string `json:"query" api:"query"`
	// Optional on the request, required for a useful search. License type code, example AIRREF for air conditioning contractors. ALL or blank searches every type.
	LicenseType *string `json:"license_type" api:"license_type"`
	// Optional. Business city. Example: Austin. One of query, city, county, or postal_code is required.
	City *string `json:"city" api:"city"`
	// Optional. Business county. Example: TRAVIS.
	County *string `json:"county" api:"county"`
	// Optional. Business ZIP. Example: 78701.
	PostalCode *string `json:"postal_code" api:"postal_code"`
	// Optional. Page number starting at 1. Example: 2. Defaults to 1. Each page has 25 rows.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TdlrSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 9)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.LicenseType != nil {
		values = append(values, param{"license_type", *p.LicenseType})
	}
	if p.City != nil {
		values = append(values, param{"city", *p.City})
	}
	if p.County != nil {
		values = append(values, param{"county", *p.County})
	}
	if p.PostalCode != nil {
		values = append(values, param{"postal_code", *p.PostalCode})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TedNoticeParams holds the parameters of TedNotice (GET / POST /v2/ted/notice).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TedNoticeParams struct {
	// Publication number, or a TED notice URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TedNoticeParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TedSearchParams holds the parameters of TedSearch (GET / POST /v2/ted/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TedSearchParams struct {
	// Keyword, publication number, or a TED search URL.
	Query *string `json:"query" api:"query"`
	// Notice code such as cn-standard or can-standard. Combined with AND.
	NoticeType *string `json:"notice_type" api:"notice_type"`
	// Buyer country, ISO alpha-2 or alpha-3. FR becomes FRA. A country name returns 400.
	Country *string `json:"country" api:"country"`
	// Place of performance: country code or NUTS code such as FRK24.
	Place *string `json:"place" api:"place"`
	// CPV code or prefix. Digits and an optional star, for example 72000000 or 30*.
	Cpv *string `json:"cpv" api:"cpv"`
	// Buyer name, matched as a phrase.
	Buyer *string `json:"buyer" api:"buyer"`
	// Page from 1. next is always null: pass page + 1. A page past the last one returns 404.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TedSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 10)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.NoticeType != nil {
		values = append(values, param{"notice_type", *p.NoticeType})
	}
	if p.Country != nil {
		values = append(values, param{"country", *p.Country})
	}
	if p.Place != nil {
		values = append(values, param{"place", *p.Place})
	}
	if p.Cpv != nil {
		values = append(values, param{"cpv", *p.Cpv})
	}
	if p.Buyer != nil {
		values = append(values, param{"buyer", *p.Buyer})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ThaliaProductParams holds the parameters of ThaliaProduct (GET / POST /v2/thalia/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ThaliaProductParams struct {
	// Product URL or Thalia article id.
	Query string `json:"query" api:"query,required"`
	// When true, returns {"result": "<html>..."} instead of the parsed JSON. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ThaliaProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ThaliaSearchParams holds the parameters of ThaliaSearch (GET / POST /v2/thalia/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ThaliaSearchParams struct {
	// Keyword or a thalia.de search URL (/suche or a category URL). A product URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Page number, from 1. Sent to Thalia as the query parameter p. Default 1.
	Page *float64 `json:"page" api:"page"`
	// When true, returns {"result": "<html>..."} instead of the parsed JSON. Default false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ThaliaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TheforkRestaurantParams holds the parameters of TheforkRestaurant (GET / POST /v2/thefork/restaurant).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TheforkRestaurantParams struct {
	// Full TheFork restaurant page URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TheforkRestaurantParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TheforkSearchParams holds the parameters of TheforkSearch (GET / POST /v2/thefork/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TheforkSearchParams struct {
	// Full TheFork search or city listing URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TheforkSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ThomasnetCompanyParams holds the parameters of ThomasnetCompany (GET / POST /v2/thomasnet/company).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ThomasnetCompanyParams struct {
	// Full Thomasnet company profile URL from your browser.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ThomasnetCompanyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ThomasnetSearchParams holds the parameters of ThomasnetSearch (GET / POST /v2/thomasnet/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ThomasnetSearchParams struct {
	// Full Thomasnet search URL from your browser.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ThomasnetSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TicketmasterEventParams holds the parameters of TicketmasterEvent (GET / POST /v2/ticketmaster/event).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TicketmasterEventParams struct {
	// Required. Slugged event URL, example https://www.ticketmaster.com/suns-vs-jazz-kids-day-phoenix-arizona-01-24-2027/event/19006507A7257EEA. A short /event/{discoveryId} link has no slug and cannot be resolved.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TicketmasterEventParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TicketmasterSearchParams holds the parameters of TicketmasterSearch (GET / POST /v2/ticketmaster/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TicketmasterSearchParams struct {
	// Required. Keyword, example jazz, or a ticketmaster.com search URL such as https://www.ticketmaster.com/search?q=jazz. An event URL is rejected.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number starting at 1. Example: 2. Defaults to 1. On ticketmaster.com the page query is zero-based, so API page 2 is page=1 on the site.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TicketmasterSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TiktoktopadsAdParams holds the parameters of TiktoktopadsAd (GET / POST /v2/tiktoktopads/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TiktoktopadsAdParams struct {
	// Ad id or detail URL from search.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TiktoktopadsAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TiktoktopadsSearchParams holds the parameters of TiktoktopadsSearch (GET / POST /v2/tiktoktopads/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TiktoktopadsSearchParams struct {
	// Full TikTok Top Ads search URL, or a keyword.
	Query *string `json:"query" api:"query"`
	// Brand or product keyword.
	Keyword *string `json:"keyword" api:"keyword"`
	// ISO country code (country_code).
	Region *string `json:"region" api:"region"`
	// Campaign objective id (URL param object).
	Objective *string `json:"objective" api:"objective"`
	// Lookback window: 7, 30, or 180. Default: "30".
	Period *string `json:"period" api:"period"`
	// Sort: for_you, reach/impression, ctr/like.
	OrderBy *string `json:"order_by" api:"order_by"`
	// Industry key filter.
	Industry *string `json:"industry" api:"industry"`
	// Page number.
	Page *float64 `json:"page" api:"page"`
	// Page size (max 20).
	PerPage *float64 `json:"per_page" api:"per_page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TiktoktopadsSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Query != nil {
		values = append(values, param{"query", *p.Query})
	}
	if p.Keyword != nil {
		values = append(values, param{"keyword", *p.Keyword})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Objective != nil {
		values = append(values, param{"objective", *p.Objective})
	}
	if p.Period != nil {
		values = append(values, param{"period", *p.Period})
	}
	if p.OrderBy != nil {
		values = append(values, param{"order_by", *p.OrderBy})
	}
	if p.Industry != nil {
		values = append(values, param{"industry", *p.Industry})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.PerPage != nil {
		values = append(values, param{"per_page", *p.PerPage})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TradeaplaneListingParams holds the parameters of TradeaplaneListing (GET / POST /v2/tradeaplane/listing).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TradeaplaneListingParams struct {
	// Listing URL that includes listing_id.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TradeaplaneListingParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TradeaplaneSearchParams holds the parameters of TradeaplaneSearch (GET / POST /v2/tradeaplane/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TradeaplaneSearchParams struct {
	// Keyword, or a Trade-A-Plane search URL without listing_id.
	Query string `json:"query" api:"query,required"`
	// aircraft or engine when query is a keyword. Ignored when query is a URL.
	ListingType *string `json:"listing_type" api:"listing_type"`
	// Page from 1. Sent as s-page.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TradeaplaneSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ListingType != nil {
		values = append(values, param{"listing_type", *p.ListingType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// TrustpilotCompanyInfoParams holds the parameters of TrustpilotCompanyInfo (GET / POST /v2/trustpilot/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type TrustpilotCompanyInfoParams struct {
	// A Trustpilot company URL or domain name.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *TrustpilotCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UntappdBeerParams holds the parameters of UntappdBeer (GET / POST /v2/untappd/beer).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UntappdBeerParams struct {
	// Required. The numeric Untappd beer id from search, or a beer URL under /b/ or /beer/. A keyword is not accepted.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UntappdBeerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UntappdSearchParams holds the parameters of UntappdSearch (GET / POST /v2/untappd/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UntappdSearchParams struct {
	// Required. A beer keyword, or an Untappd search URL whose path is /search.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UntappdSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UpworkFreelancerParams holds the parameters of UpworkFreelancer (GET / POST /v2/upwork/freelancer).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UpworkFreelancerParams struct {
	// Upwork freelancer profile ID, username, or full profile URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UpworkFreelancerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UpworkJobsParams holds the parameters of UpworkJobs (GET / POST /v2/upwork/jobs).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UpworkJobsParams struct {
	// Upwork skill slug or full freelance-jobs URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UpworkJobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UpworkServicesParams holds the parameters of UpworkServices (GET / POST /v2/upwork/services).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UpworkServicesParams struct {
	// Upwork Project Catalog category slug or full services URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UpworkServicesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UpworkServicesInfoParams holds the parameters of UpworkServicesInfo (GET / POST /v2/upwork/services/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UpworkServicesInfoParams struct {
	// Full Upwork Project Catalog product URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UpworkServicesInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UsaspendingAwardParams holds the parameters of UsaspendingAward (GET / POST /v2/usaspending/award).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UsaspendingAwardParams struct {
	// Required. Generated award id or a usaspending.gov/award URL. Example: CONT_AWD_GSFC0198106DNAS526555_8000_-NONE-_-NONE-.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UsaspendingAwardParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// UsaspendingSearchParams holds the parameters of UsaspendingSearch (GET / POST /v2/usaspending/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type UsaspendingSearchParams struct {
	// Required. Keyword or a usaspending.gov search URL. Example: software. An award URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Optional. One group: contract (default when omitted), grant, idv, or loan. Example: grant.
	AwardType *string `json:"award_type" api:"award_type"`
	// Optional. Page number from 1. Example: 2. Defaults to 1. Each page has 10 awards. total_results stays null.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *UsaspendingSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.AwardType != nil {
		values = append(values, param{"award_type", *p.AwardType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VestiairecollectiveProductParams holds the parameters of VestiairecollectiveProduct (GET / POST /v2/vestiairecollective/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VestiairecollectiveProductParams struct {
	// A vestiairecollective.com product URL ending in the item id and .shtml. The response is one flat listing. price is the number on the page and price_display is that amount as written. A keyword or a search URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Default false. true replaces the listing with result, the rendered HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VestiairecollectiveProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VestiairecollectiveSearchParams holds the parameters of VestiairecollectiveSearch (GET / POST /v2/vestiairecollective/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VestiairecollectiveSearchParams struct {
	// A keyword or a vestiairecollective.com search or category URL. Empty, another host, or a product URL returns 400. A keyword searches www.vestiairecollective.com. A full URL keeps its host. Each card has price as a number and price_display as the label on the card.
	Query string `json:"query" api:"query,required"`
	// Page number starting at 1. Omit it for page 1, which has no /p-N/ segment. Page 2 requests /p-2/ and sets pagination.page to 2. pagination.next is the next catalog URL, or null on the last page. A value below 1 returns 400.
	Page *float64 `json:"page" api:"page"`
	// Default false. true replaces results and pagination with result, the rendered HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VestiairecollectiveSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VintedItemParams holds the parameters of VintedItem (GET / POST /v2/vinted/item).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VintedItemParams struct {
	// Vinted item URL or numeric item ID.
	Query string `json:"query" api:"query,required"`
	// Force the Vinted domain when query is a bare item ID only.
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VintedItemParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VintedSearchParams holds the parameters of VintedSearch (GET / POST /v2/vinted/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VintedSearchParams struct {
	// Search keyword or full Vinted catalog URL. Filters from the URL query string are forwarded to the internal Vinted API.
	Query string `json:"query" api:"query,required"`
	// Page number. Default 1.
	Page *float64 `json:"page" api:"page"`
	// Results per page. Default 24.
	PerPage *float64 `json:"per_page" api:"per_page"`
	// Optional sort order.
	Order *string `json:"order" api:"order"`
	// Force the Vinted domain when query is a keyword only.
	Region *string `json:"region" api:"region"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VintedSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.PerPage != nil {
		values = append(values, param{"per_page", *p.PerPage})
	}
	if p.Order != nil {
		values = append(values, param{"order", *p.Order})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VintedUserInfoParams holds the parameters of VintedUserInfo (GET / POST /v2/vinted/user/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VintedUserInfoParams struct {
	// Vinted member profile URL, direct API URL, or numeric user ID.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VintedUserInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VivarealPropertyParams holds the parameters of VivarealProperty (GET / POST /v2/vivareal/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VivarealPropertyParams struct {
	// Viva Real listing URL under /imovel/ with an id- number.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VivarealPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// VivarealSearchParams holds the parameters of VivarealSearch (GET / POST /v2/vivareal/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type VivarealSearchParams struct {
	// City, City plus UF, or a Viva Real search URL.
	Query string `json:"query" api:"query,required"`
	// Sale or rent when query is a city. Default: "sale".
	Operation *string `json:"operation" api:"operation"`
	// Optional type when query is a city.
	PropertyType *string `json:"property_type" api:"property_type"`
	// Page from 1. Placeholder 2. Viva Real receives this as pagina=2. Page 1 and page 2 return different listings. A few featured ids can repeat. Most ids on page 2 are new.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VivarealSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Operation != nil {
		values = append(values, param{"operation", *p.Operation})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WalmartProductParams holds the parameters of WalmartProduct (GET / POST /v2/walmart/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WalmartProductParams struct {
	// Full Walmart product page URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WalmartProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WalmartSearchParams holds the parameters of WalmartSearch (GET / POST /v2/walmart/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WalmartSearchParams struct {
	// Full Walmart search results URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WalmartSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WalmartSuggestParams holds the parameters of WalmartSuggest (GET / POST /v2/walmart/suggest).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WalmartSuggestParams struct {
	// Search keyword for Walmart autocomplete suggestions.
	Query string `json:"query" api:"query,required"`
	// Optional US ZIP code to localize Walmart store and inventory context for suggestions.
	Zipcode *string `json:"zipcode" api:"zipcode"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WalmartSuggestParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Zipcode != nil {
		values = append(values, param{"zipcode", *p.Zipcode})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteAntibotParams holds the parameters of WebsiteAntibot (GET /v2/website/antibot).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteAntibotParams struct {
	// Full URL or domain name to analyze for anti-bot protection. Must be a valid hostname.
	Query string `json:"query" api:"query,required"`
	// Maximum request timeout in seconds. Defaults to 15.
	Timeout *float64 `json:"timeout" api:"timeout"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteAntibotParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Timeout != nil {
		values = append(values, param{"timeout", *p.Timeout})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteCrawlerParams holds the parameters of WebsiteCrawler (GET / POST /v2/website/crawler).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteCrawlerParams struct {
	// A public website URL to fetch with the Crawler engine (HTTP only, no JavaScript). Must include http:// or https://.
	Query string `json:"query" api:"query,required"`
	// Controls whether the crawler should follow HTTP redirects. Default: true.
	AllowRedirects *bool `json:"allow_redirects" api:"allow_redirects"`
	// Controls whether the response returns the raw, unprocessed HTML source of the page. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteCrawlerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.AllowRedirects != nil {
		values = append(values, param{"allow_redirects", *p.AllowRedirects})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteEmailPhoneExtractorParams holds the parameters of WebsiteEmailPhoneExtractor (GET / POST /v2/website/email_phone_extractor).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteEmailPhoneExtractorParams struct {
	// A website URL to crawl for contact information. Must include the http:// or https:// protocol. The crawler will follow internal links ("a" anchors) to find contact details across the site.
	Query string `json:"query" api:"query,required"`
	// An ISO 3166-1 alpha-2 country code used to improve the accuracy of phone number validation and parsing.
	CountryCode *string `json:"country_code" api:"country_code"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteEmailPhoneExtractorParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.CountryCode != nil {
		values = append(values, param{"country_code", *p.CountryCode})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteRenderingParams holds the parameters of WebsiteRendering (GET / POST /v2/website/rendering).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteRenderingParams struct {
	// A public website URL to render in a headless browser. Must include http:// or https://.
	Query string `json:"query" api:"query,required"`
	// Number of seconds to wait after page load before returning the rendered HTML. Useful for pages with delayed rendering or animations. Default: 0.
	WaitInSeconds *float64 `json:"wait_in_seconds" api:"wait_in_seconds"`
	// A CSS or XPath selector to wait for before returning the rendered HTML. The headless browser will hold the response until the specified element appears in the DOM.
	WaitFor *string `json:"wait_for" api:"wait_for"`
	// Controls whether the headless browser should block ad requests during rendering. Default: false.
	BlockAds *bool `json:"block_ads" api:"block_ads"`
	// Maximum number of seconds to wait for the page to fully load before returning whatever has been rendered so far. Default: 30.
	Timeout *float64 `json:"timeout" api:"timeout"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteRenderingParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.WaitInSeconds != nil {
		values = append(values, param{"wait_in_seconds", *p.WaitInSeconds})
	}
	if p.WaitFor != nil {
		values = append(values, param{"wait_for", *p.WaitFor})
	}
	if p.BlockAds != nil {
		values = append(values, param{"block_ads", *p.BlockAds})
	}
	if p.Timeout != nil {
		values = append(values, param{"timeout", *p.Timeout})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteScreenshotParams holds the parameters of WebsiteScreenshot (GET / POST /v2/website/screenshot).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteScreenshotParams struct {
	// The URL of the webpage to capture. Examples:  https://www.amazon.com/s?k=pokemon
	URL string `json:"url" api:"url,required"`
	// The output format of the screenshot. Allowed values: jpeg, jpg, pdf, png, webp. Default: "png".
	Format *string `json:"format" api:"format"`
	// The device type to emulate for the screenshot. Allowed values: galaxy_fold4, galaxy_s23, galaxy_s23_ultra, galaxy_s8, huawei_mate_50_pro, huawei_p60_pro, ipad, iphone_12, iphone_13_mini, iphone_13_pro_max, iphone_14_pro, iphone_14_pro_max, iphone_x, pixel_5, pixel_6a, pixel_7_pro, redmi_note_11, re…
	Device *string `json:"device" api:"device"`
	// Capture the entire page instead of just the visible viewport. Default: false.
	FullPage *bool `json:"full_page" api:"full_page"`
	// Scroll the page to fully load lazy-loaded elements before capturing a full-page screenshot. Default: false.
	FullPageScroll *bool `json:"full_page_scroll" api:"full_page_scroll"`
	// Time in milliseconds for scrolling before capturing the full page. Default: 400.
	FullPageScrollDuration *float64 `json:"full_page_scroll_duration" api:"full_page_scroll_duration"`
	// The width of the browser viewport in pixels. Default: 1280.
	ViewportWidth *float64 `json:"viewport_width" api:"viewport_width"`
	// The height of the browser viewport in pixels. Default: 1024.
	ViewportHeight *float64 `json:"viewport_height" api:"viewport_height"`
	// The scale factor for high-resolution screenshots. Default: 1.
	ScaleFactor *float64 `json:"scale_factor" api:"scale_factor"`
	// Generate images with custom quality settings. Default: "80".
	ImageQuality *string `json:"image_quality" api:"image_quality"`
	// Capture a specific element on the page instead of the full viewport.
	Selector *string `json:"selector" api:"selector"`
	// Cache the response. Default: false.
	Cache *bool `json:"cache" api:"cache"`
	// Cache the response for a custom TTL (in seconds). Default: 2592000.
	CacheTTL *float64 `json:"cache_ttl" api:"cache_ttl"`
	// A comma-separated list of elements to hide before capturing (e.g., ads, popups).
	RemoveSelectors *string `json:"remove_selectors" api:"remove_selectors"`
	// Automatically remove cookie banners before capturing. Default: false.
	RemoveCookieBanners *bool `json:"remove_cookie_banners" api:"remove_cookie_banners"`
	// Automatically remove ads before capturing. Default: false.
	RemoveAds *bool `json:"remove_ads" api:"remove_ads"`
	// Comma-separated list of resource types to block (e.g., “image,stylesheet,font”). Allowed values: document, eventsource, fetch, font, image, manifest, media, other, script, stylesheet, texttrack, websocket, xhr.
	BlockResources *string `json:"block_resources" api:"block_resources"`
	// Comma-separated list of URL patterns to block (e.g., “analytics,tracking,advertisement”).
	BlockUrls *string `json:"block_urls" api:"block_urls"`
	// Delay in seconds before capturing the screenshot. Default: 0.
	Delay *float64 `json:"delay" api:"delay"`
	// Define when to capture. Allowed values: domcontentloaded, load, networkidle0, networkidle2.
	WaitUntil []string `json:"wait_until" api:"wait_until"`
	// Wait for a specific element to appear before taking the screenshot.
	WaitForSelector *string `json:"wait_for_selector" api:"wait_for_selector"`
	// Specify a proxy server to route your request through.
	Proxy *string `json:"proxy" api:"proxy"`
	// AWS S3 secret access key.
	S3SecretKey *string `json:"s3_secret_key" api:"s3_secret_key"`
	// AWS S3 access key ID.
	S3AccessKeyID *string `json:"s3_access_key_id" api:"s3_access_key_id"`
	// The S3 bucket name where the screenshot will be uploaded.
	S3Bucket *string `json:"s3_bucket" api:"s3_bucket"`
	// The filename for the S3 object (auto-generated if not provided).
	S3ObjectKey *string `json:"s3_object_key" api:"s3_object_key"`
	// AWS S3 region for storage.
	S3Region *string `json:"s3_region" api:"s3_region"`
	// Leave this field empty for AWS S3. Only specify a value if required. S3-compatible storage services are supported, for example, use https://<accountId>.r2.cloudflarestorage.com for Cloudflare R2.
	StorageEndpoint *string `json:"storage_endpoint" api:"storage_endpoint"`
	// Use S3 storage for screenshots. The S3 URL will be returned in the response. Default: false.
	S3URL *bool `json:"s3_url" api:"s3_url"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteScreenshotParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 32)
	if p.URL != "" {
		values = append(values, param{"url", p.URL})
	}
	if p.Format != nil {
		values = append(values, param{"format", *p.Format})
	}
	if p.Device != nil {
		values = append(values, param{"device", *p.Device})
	}
	if p.FullPage != nil {
		values = append(values, param{"full_page", *p.FullPage})
	}
	if p.FullPageScroll != nil {
		values = append(values, param{"full_page_scroll", *p.FullPageScroll})
	}
	if p.FullPageScrollDuration != nil {
		values = append(values, param{"full_page_scroll_duration", *p.FullPageScrollDuration})
	}
	if p.ViewportWidth != nil {
		values = append(values, param{"viewport_width", *p.ViewportWidth})
	}
	if p.ViewportHeight != nil {
		values = append(values, param{"viewport_height", *p.ViewportHeight})
	}
	if p.ScaleFactor != nil {
		values = append(values, param{"scale_factor", *p.ScaleFactor})
	}
	if p.ImageQuality != nil {
		values = append(values, param{"image_quality", *p.ImageQuality})
	}
	if p.Selector != nil {
		values = append(values, param{"selector", *p.Selector})
	}
	if p.Cache != nil {
		values = append(values, param{"cache", *p.Cache})
	}
	if p.CacheTTL != nil {
		values = append(values, param{"cache_ttl", *p.CacheTTL})
	}
	if p.RemoveSelectors != nil {
		values = append(values, param{"remove_selectors", *p.RemoveSelectors})
	}
	if p.RemoveCookieBanners != nil {
		values = append(values, param{"remove_cookie_banners", *p.RemoveCookieBanners})
	}
	if p.RemoveAds != nil {
		values = append(values, param{"remove_ads", *p.RemoveAds})
	}
	if p.BlockResources != nil {
		values = append(values, param{"block_resources", *p.BlockResources})
	}
	if p.BlockUrls != nil {
		values = append(values, param{"block_urls", *p.BlockUrls})
	}
	if p.Delay != nil {
		values = append(values, param{"delay", *p.Delay})
	}
	if p.WaitUntil != nil {
		values = append(values, param{"wait_until", p.WaitUntil})
	}
	if p.WaitForSelector != nil {
		values = append(values, param{"wait_for_selector", *p.WaitForSelector})
	}
	if p.Proxy != nil {
		values = append(values, param{"proxy", *p.Proxy})
	}
	if p.S3SecretKey != nil {
		values = append(values, param{"s3_secret_key", *p.S3SecretKey})
	}
	if p.S3AccessKeyID != nil {
		values = append(values, param{"s3_access_key_id", *p.S3AccessKeyID})
	}
	if p.S3Bucket != nil {
		values = append(values, param{"s3_bucket", *p.S3Bucket})
	}
	if p.S3ObjectKey != nil {
		values = append(values, param{"s3_object_key", *p.S3ObjectKey})
	}
	if p.S3Region != nil {
		values = append(values, param{"s3_region", *p.S3Region})
	}
	if p.StorageEndpoint != nil {
		values = append(values, param{"storage_endpoint", *p.StorageEndpoint})
	}
	if p.S3URL != nil {
		values = append(values, param{"s3_url", *p.S3URL})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteTechnologyParams holds the parameters of WebsiteTechnology (GET / POST /v2/website/technology).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteTechnologyParams struct {
	// A website URL to analyze for technology stack detection. Must include the http:// or https:// protocol.
	Query string `json:"query" api:"query,required"`
	// Controls the level of detail returned in the technology detection response. Default: "expert".
	Mode *string `json:"mode" api:"mode"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteTechnologyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Mode != nil {
		values = append(values, param{"mode", *p.Mode})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteWebunlockerParams holds the parameters of WebsiteWebunlocker (GET / POST /v2/website/webunlocker).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WebsiteWebunlockerParams struct {
	// A public website URL to unlock with the WebUnlocker engine. Must include http:// or https://. The target domain must be on Piloterr's approved whitelist.
	Query string `json:"query" api:"query,required"`
	// Controls whether the WebUnlocker engine should follow HTTP redirects. Default: true.
	AllowRedirects *bool `json:"allow_redirects" api:"allow_redirects"`
	// Controls whether the response returns the raw, unprocessed HTML source of the page. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteWebunlockerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.AllowRedirects != nil {
		values = append(values, param{"allow_redirects", *p.AllowRedirects})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WeiboPostInfoParams holds the parameters of WeiboPostInfo (GET / POST /v2/weibo/post/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WeiboPostInfoParams struct {
	// Weibo post id.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WeiboPostInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WeiboSearchParams holds the parameters of WeiboSearch (GET / POST /v2/weibo/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WeiboSearchParams struct {
	// Keywords to search, such as AI.
	Query string `json:"query" api:"query,required"`
	// Page number, starting at 1. Default 1. Page 2 is the next slice of the same keywords. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WeiboSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WelcometothejungleCompanyParams holds the parameters of WelcometothejungleCompany (GET / POST /v2/welcometothejungle/company).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WelcometothejungleCompanyParams struct {
	// Slug, company name, or a Welcome to the Jungle company URL.
	Query string `json:"query" api:"query,required"`
	// fr (default) or en. en-gb and en-us count as en. Chooses the language of the profile text and URL. Default: "fr".
	Language *string `json:"language" api:"language"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WelcometothejungleCompanyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Language != nil {
		values = append(values, param{"language", *p.Language})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WelcometothejungleSearchParams holds the parameters of WelcometothejungleSearch (GET / POST /v2/welcometothejungle/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WelcometothejungleSearchParams struct {
	// Job keyword or a Welcome to the Jungle jobs-search URL.
	Query string `json:"query" api:"query,required"`
	// fr (default) or en. en-gb and en-us count as en. A /fr/ or /en/ path in the URL overrides this. Default: "fr".
	Language *string `json:"language" api:"language"`
	// Page from 1, 30 offers. A page value already in the URL overrides this.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WelcometothejungleSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Language != nil {
		values = append(values, param{"language", *p.Language})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WellfoundCompanyInfoParams holds the parameters of WellfoundCompanyInfo (GET / POST /v2/wellfound/company/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type WellfoundCompanyInfoParams struct {
	// A Wellfound company URL or company slug.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WellfoundCompanyInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// XPostInfoParams holds the parameters of XPostInfo (GET / POST /v2/x/post/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type XPostInfoParams struct {
	// X tweet to retrieve. Accepts a full status URL or numeric tweet ID.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *XPostInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// XProfileInfoParams holds the parameters of XProfileInfo (GET / POST /v2/x/profile/info).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type XProfileInfoParams struct {
	// X profile to look up. Accepts a username, @handle, or profile URL.
	Query string `json:"query" api:"query,required"`
	// When true, return the raw HTML page source as {"result": "..."} instead of structured JSON. Defaults to false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *XProfileInfoParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// YahooSearchParams holds the parameters of YahooSearch (GET / POST /v2/yahoo/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type YahooSearchParams struct {
	// Search query sent to Yahoo Search.
	Query string `json:"query" api:"query,required"`
	// Page number. Page 1 is the first page of results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code, for example us, fr, gb, de, jp.
	Gl *string `json:"gl" api:"gl"`
	// Language code, for example en, fr, de, ja, zh.
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YahooSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// YellowpagesBusinessParams holds the parameters of YellowpagesBusiness (GET / POST /v2/yellowpages/business).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type YellowpagesBusinessParams struct {
	// Full Yellow Pages business listing URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YellowpagesBusinessParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// YellowpagesSearchParams holds the parameters of YellowpagesSearch (GET / POST /v2/yellowpages/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type YellowpagesSearchParams struct {
	// Full Yellow Pages search URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YellowpagesSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// YelpBusinessParams holds the parameters of YelpBusiness (GET / POST /v2/yelp/business).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type YelpBusinessParams struct {
	// Full Yelp business URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YelpBusinessParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// YelpReviewsParams holds the parameters of YelpReviews (GET / POST /v2/yelp/reviews).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type YelpReviewsParams struct {
	// Yelp business ID / encid.
	Query string `json:"query" api:"query,required"`
	// Page number. Each page returns up to 10 reviews. Default: 1. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YelpReviewsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// YelpSearchParams holds the parameters of YelpSearch (GET / POST /v2/yelp/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type YelpSearchParams struct {
	// Full Yelp search URL with find_desc and find_loc query parameters.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YelpSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZapimoveisAdParams holds the parameters of ZapimoveisAd (GET / POST /v2/zapimoveis/ad).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZapimoveisAdParams struct {
	// Full Zap Imóveis listing URL. Accepts /imovel/ and /lancamentos/.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZapimoveisAdParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZapimoveisSearchParams holds the parameters of ZapimoveisSearch (GET / POST /v2/zapimoveis/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZapimoveisSearchParams struct {
	// City, City, UF, neighborhood, or full Zap Imóveis search URL.
	Query string `json:"query" api:"query,required"`
	// Transaction type. Used when query is a city. Ignored as a path builder when query is a full URL. Default: "sale".
	Operation *string `json:"operation" api:"operation"`
	// Property type. Used when query is a city.
	PropertyType *string `json:"property_type" api:"property_type"`
	// Page number (1-based). Default: 1. Values above 1 set pagina on the URL. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Minimum asking price in BRL.
	MinPrice *float64 `json:"min_price" api:"min_price"`
	// Maximum asking price in BRL.
	MaxPrice *float64 `json:"max_price" api:"max_price"`
	// Minimum area in m².
	MinArea *float64 `json:"min_area" api:"min_area"`
	// Maximum area in m².
	MaxArea *float64 `json:"max_area" api:"max_area"`
	// Minimum bedroom count.
	MinBedrooms *float64 `json:"min_bedrooms" api:"min_bedrooms"`
	// Minimum bathroom count.
	MinBathrooms *float64 `json:"min_bathrooms" api:"min_bathrooms"`
	// Minimum parking spaces.
	MinParking *float64 `json:"min_parking" api:"min_parking"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZapimoveisSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 14)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Operation != nil {
		values = append(values, param{"operation", *p.Operation})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.MinPrice != nil {
		values = append(values, param{"min_price", *p.MinPrice})
	}
	if p.MaxPrice != nil {
		values = append(values, param{"max_price", *p.MaxPrice})
	}
	if p.MinArea != nil {
		values = append(values, param{"min_area", *p.MinArea})
	}
	if p.MaxArea != nil {
		values = append(values, param{"max_area", *p.MaxArea})
	}
	if p.MinBedrooms != nil {
		values = append(values, param{"min_bedrooms", *p.MinBedrooms})
	}
	if p.MinBathrooms != nil {
		values = append(values, param{"min_bathrooms", *p.MinBathrooms})
	}
	if p.MinParking != nil {
		values = append(values, param{"min_parking", *p.MinParking})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZazzleProductParams holds the parameters of ZazzleProduct (GET / POST /v2/zazzle/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZazzleProductParams struct {
	// Required. Product URL or slug ending in the numeric product id. Example: https://www.zazzle.com/branded_logo_two_tone_coffee_mug-256528848440338952.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZazzleProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZazzleSearchParams holds the parameters of ZazzleSearch (GET / POST /v2/zazzle/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZazzleSearchParams struct {
	// Required. Keyword or a zazzle.com/s search URL. Example: coffee mug. A product URL returns 400.
	Query string `json:"query" api:"query,required"`
	// Optional. Page number from 1. Example: 2. Maps to the Zazzle pg parameter. Defaults to 1.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZazzleSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZillowPropertyParams holds the parameters of ZillowProperty (GET / POST /v2/zillow/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZillowPropertyParams struct {
	// Full Zillow homedetails URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZillowPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZillowSearchParams holds the parameters of ZillowSearch (GET / POST /v2/zillow/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZillowSearchParams struct {
	// Full Zillow search or map results URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZillowSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZonapropPropertyParams holds the parameters of ZonapropProperty (GET / POST /v2/zonaprop/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZonapropPropertyParams struct {
	// ZonaProp listing URL under /propiedades/clasificado/.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZonapropPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZonapropSearchParams holds the parameters of ZonapropSearch (GET / POST /v2/zonaprop/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZonapropSearchParams struct {
	// Place or a ZonaProp search URL.
	Query string `json:"query" api:"query,required"`
	// Sale or rent when query is a place.
	Operation *string `json:"operation" api:"operation"`
	// Kind of property. Default departamentos.
	PropertyType *string `json:"property_type" api:"property_type"`
	// Page from 1. ZonaProp writes this into the path as -pagina-N.html.
	Page *float64 `json:"page" api:"page"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZonapropSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Operation != nil {
		values = append(values, param{"operation", *p.Operation})
	}
	if p.PropertyType != nil {
		values = append(values, param{"property_type", *p.PropertyType})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZooplaPropertyParams holds the parameters of ZooplaProperty (GET / POST /v2/zoopla/property).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZooplaPropertyParams struct {
	// Full Zoopla property detail URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZooplaPropertyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZooplaSearchParams holds the parameters of ZooplaSearch (GET / POST /v2/zoopla/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZooplaSearchParams struct {
	// Full Zoopla search results URL.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZooplaSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZooplusProductParams holds the parameters of ZooplusProduct (GET / POST /v2/zooplus/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZooplusProductParams struct {
	// A product link, a /shop/.../{id} path, or the product number.
	Query string `json:"query" api:"query,required"`
	// Which country shop when query is a number or a path. fr is France (default). de is Germany. uk is the United Kingdom. Default: "fr".
	Region *string `json:"region" api:"region"`
	// Set to true to get the raw HTML in result instead of the product.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZooplusProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZooplusSearchParams holds the parameters of ZooplusSearch (GET / POST /v2/zooplus/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZooplusSearchParams struct {
	// What to look for. A word, or a Zooplus search link.
	Query string `json:"query" api:"query,required"`
	// Which country shop. fr is France (default, www.zooplus.fr). de is Germany (www.zooplus.de). uk is the United Kingdom (www.zooplus.co.uk). Default: "fr".
	Region *string `json:"region" api:"region"`
	// Which page of results, starting at 1. Page 2 uses p=2 and returns different products.
	Page *float64 `json:"page" api:"page"`
	// Set to true to get the raw HTML in result instead of the product list.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZooplusSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Region != nil {
		values = append(values, param{"region", *p.Region})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZoroProductParams holds the parameters of ZoroProduct (GET / POST /v2/zoro/product).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZoroProductParams struct {
	// Zoro product URL, or a path that contains /i/G....
	Query string `json:"query" api:"query,required"`
	// When true, return the rendered HTML as {"result": "..."} instead of JSON.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZoroProductParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// ZoroSearchParams holds the parameters of ZoroSearch (GET / POST /v2/zoro/search).
// Required fields are plain values; optional fields are pointers (use piloterr.String, piloterr.Float, piloterr.Bool) or nil slices and maps.
type ZoroSearchParams struct {
	// Keyword, or a Zoro search or category URL. Empty, another host, or a product URL returns 400. A keyword builds /search?q=.... Each card has price as a USD number and price_display as the dollar string.
	Query string `json:"query" api:"query,required"`
	// Page number starting at 1. Omit it for page 1. Page 2 adds page=2, sets pagination.page to 2, and returns the next cards. pagination.next is the next page URL, or null on the last page. A value below 1 returns 400.
	Page *float64 `json:"page" api:"page"`
	// Default false. true replaces results and pagination with result, the rendered HTML.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *ZoroSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}
