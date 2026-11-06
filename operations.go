// Code generated from the Piloterr OpenAPI document. DO NOT EDIT.

package piloterr

import "context"

// AboutyouProduct — About You Product. Read one About You product from a product URL or a keyword. A keyword opens the first search hit. Set region when the query does not already name the shop. Credit cost: 2 credits per call.
//
// GET / POST /v2/aboutyou/product
func (c *Client) AboutyouProduct(ctx context.Context, params *AboutyouProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aboutyou_product"], params.params(), options)
}

// AboutyouSearch — About You Search. Search an About You shop with a keyword or a search URL. Pick the country with region, or let a full link choose the shop. Cards include brand, price, and the product link. Credit cost: 2 credits per call.
//
// GET / POST /v2/aboutyou/search
func (c *Client) AboutyouSearch(ctx context.Context, params *AboutyouSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aboutyou_search"], params.params(), options)
}

// ActionProduct — Action Product. Action Product API: fetch full product page data from a product ID or URL on Action.com. Returns pricing, specs, images, breadcrumbs and availability via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/action/product
func (c *Client) ActionProduct(ctx context.Context, params *ActionProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["action_product"], params.params(), options)
}

// ActionSearch — Action Search. Action Search API: search product listings by keyword or URL on Action.com across European locales. Returns prices, images, locale, currency and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/action/search
func (c *Client) ActionSearch(ctx context.Context, params *ActionSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["action_search"], params.params(), options)
}

// AldiProduct — Aldi Product. Aldi Product API: fetch full product page data from an Aldi URL or path slug. Returns pricing, images, description, category and cooling flags. Credit cost: 1 credit per call.
//
// GET / POST /v2/aldi/product
func (c *Client) AldiProduct(ctx context.Context, params *AldiProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aldi_product"], params.params(), options)
}

// AldiSearch — Aldi Search. Aldi Search API: list products from an Aldi category page URL or path across FR, DE, NL and Canarias shops. Returns prices, thumbnails and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/aldi/search
func (c *Client) AldiSearch(ctx context.Context, params *AldiSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aldi_search"], params.params(), options)
}

// AlibabaProduct — Alibaba Product. Alibaba Product API. Fetch a full B2B product page by URL or product ID with optional locale subdomain. Browser rendering returns tiered pricing, seller profile and lead times. Credit cost: 2 credits per call.
//
// GET / POST /v2/alibaba/product
func (c *Client) AlibabaProduct(ctx context.Context, params *AlibabaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["alibaba_product"], params.params(), options)
}

// AlibabaSearch — Alibaba Search. Alibaba Search API. Query B2B product listings by keyword or search URL across 17 locale subdomains. Request mode returns prices, MOQ, seller data and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/alibaba/search
func (c *Client) AlibabaSearch(ctx context.Context, params *AlibabaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["alibaba_search"], params.params(), options)
}

// AliexpressProduct — AliExpress Product. AliExpress Product API: scrape live /item/ product pages from a full URL. Get title, images, seller identity, pricing, specs, delivery, and FAQs as JSON via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/aliexpress/product
func (c *Client) AliexpressProduct(ctx context.Context, params *AliexpressProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aliexpress_product"], params.params(), options)
}

// AliexpressProductReviews — AliExpress Reviews. AliExpress Reviews API: list buyer reviews for a product id or item URL. Each review includes author, star rating, text, date, and photos when the buyer attached them. Credit cost: 1 credit per call.
//
// GET / POST /v2/aliexpress/product/reviews
func (c *Client) AliexpressProductReviews(ctx context.Context, params *AliexpressProductReviewsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aliexpress_product_reviews"], params.params(), options)
}

// AliexpressSearch — AliExpress Search. AliExpress Search API: search by keyword or wholesale URL and get product listings with prices, ratings, sold counts, and pagination as JSON via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/aliexpress/search
func (c *Client) AliexpressSearch(ctx context.Context, params *AliexpressSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aliexpress_search"], params.params(), options)
}

// AliexpressStoreProduct — AliExpress Store Product. AliExpress Store Product API: list products from an AliExpress seller store by URL or store ID. Returns store metadata and product cards (~40 SSR items) as JSON via browser rendering. Pagination is not supported. Credit cost: 2 credits per call.
//
// GET / POST /v2/aliexpress/store/product
func (c *Client) AliexpressStoreProduct(ctx context.Context, params *AliexpressStoreProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["aliexpress_store_product"], params.params(), options)
}

// AllegroProduct — Allegro Product. Allegro Product API. Extract full product page data from allegro.pl by URL (/produkt/, /oferta/, /ogloszenie/). Browser rendering, 3 credits per call. Pair with Allegro Search (3 credits) for listing discovery. Credit cost: 3 credits per call.
//
// GET / POST /v2/allegro/product
func (c *Client) AllegroProduct(ctx context.Context, params *AllegroProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["allegro_product"], params.params(), options)
}

// AllegroSearch — Allegro Search. Allegro Search API. Query product listings by keyword, optional page number, or full allegro.pl listing URL. Returns results with prices, sellers, sponsorship flags and pagination. 3 credits per call. Credit cost: 3 credits per call.
//
// GET / POST /v2/allegro/search
func (c *Client) AllegroSearch(ctx context.Context, params *AllegroSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["allegro_search"], params.params(), options)
}

// AmazonProduct — Amazon Product. Efficiently extract product details from Amazon for dynamic retail insights, including price, stock, features, and seller data. Credit cost: 1 credit per call.
//
// GET / POST /v2/amazon/product
func (c *Client) AmazonProduct(ctx context.Context, params *AmazonProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["amazon_product"], params.params(), options)
}

// AmazonProductOffer — Amazon Product Offer. Efficiently extract details of third-party seller offers for any Amazon product across multiple marketplaces. Credit cost: 1 credit per call.
//
// GET / POST /v2/amazon/product/offer
func (c *Client) AmazonProductOffer(ctx context.Context, params *AmazonProductOfferParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["amazon_product_offer"], params.params(), options)
}

// AmazonSearch — Amazon Search. Efficiently scrape Amazon search results for real-time product data and insights across multiple Amazon marketplaces. Credit cost: 1 credit per call.
//
// GET / POST /v2/amazon/search
func (c *Client) AmazonSearch(ctx context.Context, params *AmazonSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["amazon_search"], params.params(), options)
}

// ApartmentsProperty — Apartments Property. Extract a full Apartments.com rental listing from a property URL, including amenities, gallery, description, geolocation, and unit details via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/apartments/property
func (c *Client) ApartmentsProperty(ctx context.Context, params *ApartmentsPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["apartments_property"], params.params(), options)
}

// ApartmentsSearch — Apartments Search. Search Apartments.com rentals from a full city or browse URL and retrieve prices, photos, addresses, phone numbers, and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/apartments/search
func (c *Client) ApartmentsSearch(ctx context.Context, params *ApartmentsSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["apartments_search"], params.params(), options)
}

// ArgenpropProperty — Argenprop Property. Read one Argenprop listing from its URL or id. You get price, address, rooms, expenses, and the agent when the page shows them. Credit cost: 1 credit per call.
//
// GET / POST /v2/argenprop/property
func (c *Client) ArgenpropProperty(ctx context.Context, params *ArgenpropPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["argenprop_property"], params.params(), options)
}

// ArgenpropSearch — Argenprop Search. Search Argenprop by neighbourhood or city. Set operation to venta or alquiler, and property_type such as departamentos or casas. Page with pagina. Credit cost: 1 credit per call.
//
// GET / POST /v2/argenprop/search
func (c *Client) ArgenpropSearch(ctx context.Context, params *ArgenpropSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["argenprop_search"], params.params(), options)
}

// ArgosProduct — Argos Product. Argos Product API: one UK product by id or URL. Browser rendering (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/argos/product
func (c *Client) ArgosProduct(ctx context.Context, params *ArgosProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["argos_product"], params.params(), options)
}

// ArgosSearch — Argos Search. Search the Argos UK catalogue by keyword and page the results. Prices are in pounds. A one-word query that Argos turns into a brand page with no product cards returns no results. Credit cost: 2 credits per call.
//
// GET / POST /v2/argos/search
func (c *Client) ArgosSearch(ctx context.Context, params *ArgosSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["argos_search"], params.params(), options)
}

// ArxivPaper — arXiv Paper. Fetch one arXiv paper by id or by an abs, pdf, or html URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/arxiv/paper
func (c *Client) ArxivPaper(ctx context.Context, params *ArxivPaperParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["arxiv_paper"], params.params(), options)
}

// ArxivSearch — arXiv Search. Search arXiv papers by keyword, author query such as au:einstein, or arXiv id. A paper URL returns 400. Page through results when the match count is large. Credit cost: 1 credit per call.
//
// GET / POST /v2/arxiv/search
func (c *Client) ArxivSearch(ctx context.Context, params *ArxivSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["arxiv_search"], params.params(), options)
}

// AuchanProduct — Auchan Product. Effortlessly extract product details from Auchan for dynamic business insights, including nutritional information, ingredients, ratings, and reviews. Credit cost: 2 credits per call.
//
// GET / POST /v2/auchan/product
func (c *Client) AuchanProduct(ctx context.Context, params *AuchanProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["auchan_product"], params.params(), options)
}

// AuchanSearch — Auchan Search. Extract Auchan product listings from keyword searches for insightful market analysis and competitive edge. Credit cost: 2 credits per call.
//
// GET / POST /v2/auchan/search
func (c *Client) AuchanSearch(ctx context.Context, params *AuchanSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["auchan_search"], params.params(), options)
}

// AutodocProduct — Autodoc Product. Scrape Autodoc product pages: price, specs, EAN, images, compatibility, and TecDoc-related vehicle data via the Related endpoint. Credit cost: 2 credits per call.
//
// GET / POST /v2/autodoc/product
func (c *Client) AutodocProduct(ctx context.Context, params *AutodocProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autodoc_product"], params.params(), options)
}

// AutodocProductRelated — Autodoc Product Related. Vehicle fitment lines per maker for an Autodoc product. All parameters from Product fetch.params are required. Credit cost: 2 credits per call.
//
// GET / POST /v2/autodoc/product/related
func (c *Client) AutodocProductRelated(ctx context.Context, params *AutodocProductRelatedParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autodoc_product_related"], params.params(), options)
}

// AutodocSearch — Autodoc Search. Search Autodoc by URL: keyword pages use /search?keyword=… (e.g. frein on auto-doc.fr). Extract product cards, prices, and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/autodoc/search
func (c *Client) AutodocSearch(ctx context.Context, params *AutodocSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autodoc_search"], params.params(), options)
}

// AutoeuropeResult — Auto Europe Result. Read one vehicle from an Auto Europe quote. Pass the offer id from Search, with the same pickup and dates. You get the supplier, the SIPP code, and the package prices. Credit cost: 1 credit per call.
//
// GET / POST /v2/autoeurope/result
func (c *Client) AutoeuropeResult(ctx context.Context, params *AutoeuropeResultParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autoeurope_result"], params.params(), options)
}

// AutoeuropeSearch — Auto Europe Search. Quote a rental car on Auto Europe. Pass pickup, pickup date, and dropoff date. Region chooses the country site and its currency. The whole quote comes back on one page. Credit cost: 1 credit per call.
//
// GET / POST /v2/autoeurope/search
func (c *Client) AutoeuropeSearch(ctx context.Context, params *AutoeuropeSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autoeurope_search"], params.params(), options)
}

// Autoscout24Ad — AutoScout24 Ad. Read one AutoScout24 listing, including an inactive ad that still exposes the listing. Price, vehicle, seller, and location. (1 credit) Credit cost: 1 credit per call.
//
// GET / POST /v2/autoscout24/ad
func (c *Client) Autoscout24Ad(ctx context.Context, params *Autoscout24AdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autoscout24_ad"], params.params(), options)
}

// Autoscout24Search — AutoScout24 Search. Search AutoScout24 listings from a search URL. Each card has make, model, numeric price, mileage, and the listing URL. (1 credit) Credit cost: 1 credit per call.
//
// GET / POST /v2/autoscout24/search
func (c *Client) Autoscout24Search(ctx context.Context, params *Autoscout24SearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autoscout24_search"], params.params(), options)
}

// AutotraderAd — Autotrader Ad. Extract full vehicle details from an AutoTrader UK car-details URL, including price, specs, gallery, seller, and description via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/autotrader/ad
func (c *Client) AutotraderAd(ctx context.Context, params *AutotraderAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autotrader_ad"], params.params(), options)
}

// AutotraderSearch — Autotrader Search. Search AutoTrader UK listings from a car-search URL and retrieve prices, specs, photos, seller info, and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/autotrader/search
func (c *Client) AutotraderSearch(ctx context.Context, params *AutotraderSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["autotrader_search"], params.params(), options)
}

// BestbuyProduct — Best Buy Product. Get a Best Buy US or Canada product page: title, price, brand, stock, images, categories and specs via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/bestbuy/product
func (c *Client) BestbuyProduct(ctx context.Context, params *BestbuyProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bestbuy_product"], params.params(), options)
}

// BestbuyProductReviews — Best Buy Reviews. Get customer reviews for a Best Buy product URL: rating, review text, author, photos and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/bestbuy/product/reviews
func (c *Client) BestbuyProductReviews(ctx context.Context, params *BestbuyProductReviewsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bestbuy_product_reviews"], params.params(), options)
}

// BestbuySearch — Best Buy Search. Search Best Buy US or Canada by keyword or search URL. Returns product listings with price, rating and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/bestbuy/search
func (c *Client) BestbuySearch(ctx context.Context, params *BestbuySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bestbuy_search"], params.params(), options)
}

// BhphotovideoProduct — B&H Photo Video Product. Read a B&H Photo Video product from its URL or id. Returns price, rating, specifications, and box contents. Credit cost: 1 credit per call.
//
// GET / POST /v2/bhphotovideo/product
func (c *Client) BhphotovideoProduct(ctx context.Context, params *BhphotovideoProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bhphotovideo_product"], params.params(), options)
}

// BhphotovideoSearch — B&H Photo Video Search. Search B&H Photo Video by keyword or search URL. Returns BH numbers, dollar prices, and a next page URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/bhphotovideo/search
func (c *Client) BhphotovideoSearch(ctx context.Context, params *BhphotovideoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bhphotovideo_search"], params.params(), options)
}

// BingSearch — Bing Search. Scrape real-time Bing search results with full control over location, language, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/bing/search
func (c *Client) BingSearch(ctx context.Context, params *BingSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bing_search"], params.params(), options)
}

// BitibaProduct — Bitiba Product. Read one product from a live Bitiba shop, including the selected variant price and the pack list. Credit cost: 1 credit per call.
//
// GET / POST /v2/bitiba/product
func (c *Client) BitibaProduct(ctx context.Context, params *BitibaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bitiba_product"], params.params(), options)
}

// BitibaSearch — Bitiba Search. Search a live Bitiba country shop. Germany is the default. www.bitiba.com is closed and returns 400. Credit cost: 1 credit per call.
//
// GET / POST /v2/bitiba/search
func (c *Client) BitibaSearch(ctx context.Context, params *BitibaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bitiba_search"], params.params(), options)
}

// BolProduct — Bol Product. Read one bol.com product by URL or product id, in nl, be-nl, or be-fr. You get the selling price, the list price, and the seller. Credit cost: 2 credits per call.
//
// GET / POST /v2/bol/product
func (c *Client) BolProduct(ctx context.Context, params *BolProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bol_product"], params.params(), options)
}

// BolSearch — Bol Search. Search bol.com by keyword. Region is nl, be-nl, or be-fr. Page the product cards. Sponsored ads can repeat across pages. Credit cost: 2 credits per call.
//
// GET / POST /v2/bol/search
func (c *Client) BolSearch(ctx context.Context, params *BolSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bol_search"], params.params(), options)
}

// BoulangerProduct — Boulanger Product. Read one Boulanger product from its URL or product id. You get price, brand, rating, seller, and specifications. Credit cost: 2 credits per call.
//
// GET / POST /v2/boulanger/product
func (c *Client) BoulangerProduct(ctx context.Context, params *BoulangerProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["boulanger_product"], params.params(), options)
}

// BoulangerSearch — Boulanger Search. Search Boulanger by keyword and page the product grid. Cards include price, brand, seller, and the product URL. Credit cost: 2 credits per call.
//
// GET / POST /v2/boulanger/search
func (c *Client) BoulangerSearch(ctx context.Context, params *BoulangerSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["boulanger_search"], params.params(), options)
}

// BraveSearch — Brave Search. Access the Brave search engine results directly through our efficient Piloterr service for reliable data scraping. Credit cost: 1 credit per call.
//
// GET / POST /v2/brave/search
func (c *Client) BraveSearch(ctx context.Context, params *BraveSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["brave_search"], params.params(), options)
}

// BricodepotProduct — Brico Dépôt Product. Read one Brico Dépôt product from its product URL. You get price, brand, SKU, and availability. Credit cost: 1 credit per call.
//
// GET / POST /v2/bricodepot/product
func (c *Client) BricodepotProduct(ctx context.Context, params *BricodepotProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bricodepot_product"], params.params(), options)
}

// BricodepotSearch — Brico Dépôt Search. Search Brico Dépôt with a /produits/ listing URL, not a keyword box. Page 2 is the same path with /2 at the end. Credit cost: 1 credit per call.
//
// GET / POST /v2/bricodepot/search
func (c *Client) BricodepotSearch(ctx context.Context, params *BricodepotSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bricodepot_search"], params.params(), options)
}

// BricoramaProduct — Bricorama Product. Read one Bricorama product from its URL. You get price, brand, and the product details. Credit cost: 1 credit per call.
//
// GET / POST /v2/bricorama/product
func (c *Client) BricoramaProduct(ctx context.Context, params *BricoramaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bricorama_product"], params.params(), options)
}

// BricoramaSearch — Bricorama Search. Search Bricorama by keyword or search URL and page the home-improvement grid. Credit cost: 1 credit per call.
//
// GET / POST /v2/bricorama/search
func (c *Client) BricoramaSearch(ctx context.Context, params *BricoramaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bricorama_search"], params.params(), options)
}

// CarrefourProduct — Carrefour Product. Read one Carrefour.fr product from its page URL. Every product returns the same keys: price, brand, pack size, Nutriscore, categories, unit price, and customer reviews. Credit cost: 2 credits per call.
//
// GET / POST /v2/carrefour/product
func (c *Client) CarrefourProduct(ctx context.Context, params *CarrefourProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["carrefour_product"], params.params(), options)
}

// CarrefourSearch — Carrefour Search. Search for products on Carrefour by keyword or URL, returning structured listings with prices, ratings, and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/carrefour/search
func (c *Client) CarrefourSearch(ctx context.Context, params *CarrefourSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["carrefour_search"], params.params(), options)
}

// CarrefourSuggest — Carrefour Suggest. Get product suggestions from Carrefour by keyword or EAN, returning matching products with pricing, availability, and product details. Credit cost: 1 credit per call.
//
// GET / POST /v2/carrefour/suggest
func (c *Client) CarrefourSuggest(ctx context.Context, params *CarrefourSuggestParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["carrefour_suggest"], params.params(), options)
}

// CastoramaProduct — Castorama Product. Read one Castorama product from its URL. You get price, brand, and specifications. Credit cost: 1 credit per call.
//
// GET / POST /v2/castorama/product
func (c *Client) CastoramaProduct(ctx context.Context, params *CastoramaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["castorama_product"], params.params(), options)
}

// CastoramaSearch — Castorama Search. Search Castorama by keyword or search URL and page the DIY listings. Cards include title, price, and product link. Credit cost: 1 credit per call.
//
// GET / POST /v2/castorama/search
func (c *Client) CastoramaSearch(ctx context.Context, params *CastoramaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["castorama_search"], params.params(), options)
}

// CdiscountProduct — Cdiscount Product. Extract full product details from a Cdiscount product page, including price, seller, ratings, specifications, and offers. Credit cost: 2 credits per call.
//
// GET / POST /v2/cdiscount/product
func (c *Client) CdiscountProduct(ctx context.Context, params *CdiscountProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["cdiscount_product"], params.params(), options)
}

// CdiscountProductOffer — Cdiscount Product Offer. Retrieve all marketplace seller offers for a Cdiscount product, including prices, seller ratings, shipping costs, and availability. Credit cost: 2 credits per call.
//
// GET / POST /v2/cdiscount/product/offer
func (c *Client) CdiscountProductOffer(ctx context.Context, params *CdiscountProductOfferParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["cdiscount_product_offer"], params.params(), options)
}

// CdiscountSearch — Cdiscount Search. Extract product listings from Cdiscount search results or category pages for market intelligence and competitive analysis. Credit cost: 2 credits per call.
//
// GET / POST /v2/cdiscount/search
func (c *Client) CdiscountSearch(ctx context.Context, params *CdiscountSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["cdiscount_search"], params.params(), options)
}

// CentrisProperty — Centris Property. Read one Centris property from its URL or listing id. You get price, address, rooms, and the broker. Credit cost: 1 credit per call.
//
// GET / POST /v2/centris/property
func (c *Client) CentrisProperty(ctx context.Context, params *CentrisPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["centris_property"], params.params(), options)
}

// CentrisSearch — Centris Search. Search Centris listings by city or by a full results URL. Use the English or French path. Page through properties for sale. Credit cost: 1 credit per call.
//
// GET / POST /v2/centris/search
func (c *Client) CentrisSearch(ctx context.Context, params *CentrisSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["centris_search"], params.params(), options)
}

// ChewyProduct — Chewy Product. Read one Chewy product by URL or numeric id. The default mode is expert and returns the historical product cache. simple returns one snake_case product. Credit cost: 3 credits per call.
//
// GET / POST /v2/chewy/product
func (c *Client) ChewyProduct(ctx context.Context, params *ChewyProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["chewy_product"], params.params(), options)
}

// ChewySearch — Chewy Search. Search Chewy by keyword or listing URL. The default mode is expert and returns the historical listing objects. simple returns short snake_case cards. Credit cost: 3 credits per call.
//
// GET / POST /v2/chewy/search
func (c *Client) ChewySearch(ctx context.Context, params *ChewySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["chewy_search"], params.params(), options)
}

// ClearedjobsJob — ClearedJobs Job. Fetch one cleared job by its URL or by a slug that ends in the numeric id. Credit cost: 1 credit per call.
//
// GET / POST /v2/clearedjobs/job
func (c *Client) ClearedjobsJob(ctx context.Context, params *ClearedjobsJobParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["clearedjobs_job"], params.params(), options)
}

// ClearedjobsSearch — ClearedJobs Search. Search cleared jobs by keyword. Only the first HTML page is available, so pagination.next is null and page 2 returns 404. Credit cost: 1 credit per call.
//
// GET / POST /v2/clearedjobs/search
func (c *Client) ClearedjobsSearch(ctx context.Context, params *ClearedjobsSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["clearedjobs_search"], params.params(), options)
}

// Company — Company Database. Retrieve comprehensive company data including financials, technologies, social profiles, and location by domain name. Credit cost: 1 credit per call.
//
// GET / POST /v2/company
func (c *Client) Company(ctx context.Context, params *CompanyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["company"], params.params(), options)
}

// CoolblueProduct — Coolblue Product. Read one Coolblue product by article id or product URL, in the region of that shop. Credit cost: 1 credit per call.
//
// GET / POST /v2/coolblue/product
func (c *Client) CoolblueProduct(ctx context.Context, params *CoolblueProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["coolblue_product"], params.params(), options)
}

// CoolblueSearch — Coolblue Search. Search Coolblue with a product keyword. Region defaults to Belgium. Also use nl or de. Page through product cards. Credit cost: 1 credit per call.
//
// GET / POST /v2/coolblue/search
func (c *Client) CoolblueSearch(ctx context.Context, params *CoolblueSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["coolblue_search"], params.params(), options)
}

// CopartLot — Copart Lot. Copart Lot API: retrieve structured auction lot data by lot number or URL. Specs, damage, bids, photos, yard location and sale status across regional Copart sites. Credit cost: 2 credits per call.
//
// GET / POST /v2/copart/lot
func (c *Client) CopartLot(ctx context.Context, params *CopartLotParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["copart_lot"], params.params(), options)
}

// CopartSearch — Copart Search. Copart Search API: search auction lots by keyword or URL across Copart regional sites. Returns lot listings and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/copart/search
func (c *Client) CopartSearch(ctx context.Context, params *CopartSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["copart_search"], params.params(), options)
}

// CostcoProduct — Costco Product. Costco Product API: fetch full product page data from a product ID or URL across regional Costco sites. Returns pricing, availability, ratings, images and breadcrumbs via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/costco/product
func (c *Client) CostcoProduct(ctx context.Context, params *CostcoProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["costco_product"], params.params(), options)
}

// CostcoSearch — Costco Search. Costco Search API: search product listings by keyword or URL across 13 regional Costco sites. Returns prices, availability, review counts and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/costco/search
func (c *Client) CostcoSearch(ctx context.Context, params *CostcoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["costco_search"], params.params(), options)
}

// CrunchbaseCompanyInfo — Crunchbase Company Info. Convert URLs into detailed Crunchbase profiles for immediate company insights. Credit cost: 2 credits per call.
//
// GET / POST /v2/crunchbase/company/info
func (c *Client) CrunchbaseCompanyInfo(ctx context.Context, params *CrunchbaseCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["crunchbase_company_info"], params.params(), options)
}

// CrunchbaseEvent — Crunchbase Event. Retrieve detailed Crunchbase event profiles with speakers, sponsors, exhibitors, and press timelines. Credit cost: 2 credits per call.
//
// GET / POST /v2/crunchbase/event
func (c *Client) CrunchbaseEvent(ctx context.Context, params *CrunchbaseEventParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["crunchbase_event"], params.params(), options)
}

// CrunchbaseFundingRound — Crunchbase Funding Round. Retrieve full details of a specific Crunchbase funding round including investors and press timeline. Credit cost: 2 credits per call.
//
// GET / POST /v2/crunchbase/funding_round
func (c *Client) CrunchbaseFundingRound(ctx context.Context, params *CrunchbaseFundingRoundParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["crunchbase_funding_round"], params.params(), options)
}

// CrunchbaseFundingRounds — Crunchbase Funding Rounds. Monitor funding history and financial milestones of companies with Crunchbase data. Credit cost: 2 credits per call.
//
// GET / POST /v2/crunchbase/funding_rounds
func (c *Client) CrunchbaseFundingRounds(ctx context.Context, params *CrunchbaseFundingRoundsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["crunchbase_funding_rounds"], params.params(), options)
}

// CrunchbasePeopleInfo — Crunchbase People Info. Retrieve detailed Crunchbase profiles for individuals including career history, investments, and board positions. Credit cost: 2 credits per call.
//
// GET / POST /v2/crunchbase/people/info
func (c *Client) CrunchbasePeopleInfo(ctx context.Context, params *CrunchbasePeopleInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["crunchbase_people_info"], params.params(), options)
}

// CrunchbaseSearch — Crunchbase Search. Search Crunchbase for companies and people by keyword and retrieve their profiles and Crunchbase identifiers. Credit cost: 2 credits per call.
//
// GET / POST /v2/crunchbase/search
func (c *Client) CrunchbaseSearch(ctx context.Context, params *CrunchbaseSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["crunchbase_search"], params.params(), options)
}

// CulturaProduct — Cultura Product. Read one Cultura product from its URL or SKU. You get price, EAN, seller, condition, and stock wording. Credit cost: 2 credits per call.
//
// GET / POST /v2/cultura/product
func (c *Client) CulturaProduct(ctx context.Context, params *CulturaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["cultura_product"], params.params(), options)
}

// CulturaSearch — Cultura Search. Search Cultura by keyword and page through books, games, and culture products. Cards include title, price, availability, and the product link. Credit cost: 2 credits per call.
//
// GET / POST /v2/cultura/search
func (c *Client) CulturaSearch(ctx context.Context, params *CulturaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["cultura_search"], params.params(), options)
}

// DartyProduct — Darty Product. Read one Darty product sheet from a product URL or a Darty product id. You get price, brand, availability, rating, and specifications. Credit cost: 2 credits per call.
//
// GET / POST /v2/darty/product
func (c *Client) DartyProduct(ctx context.Context, params *DartyProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["darty_product"], params.params(), options)
}

// DartySearch — Darty Search. Search Darty by keyword or search URL and page through the listing. Each card has title, price, brand, seller, and product link. Credit cost: 2 credits per call.
//
// GET / POST /v2/darty/search
func (c *Client) DartySearch(ctx context.Context, params *DartySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["darty_search"], params.params(), options)
}

// DatasetsLinkedinCompany — Dataset Linkedin Company. Query the LinkedIn company dataset with Elasticsearch Query DSL. Filter by name, domain, industry, headcount, and country. POST-only, 1 credit per row returned. Credit cost: 1 credit per row.
//
// POST /v2/datasets/linkedin/company
func (c *Client) DatasetsLinkedinCompany(ctx context.Context, params *DatasetsLinkedinCompanyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["datasets_linkedin_company"], params.params(), options)
}

// DecathlonProduct — Decathlon Product. Read one Decathlon product from a product URL, a product path, or a keyword. A keyword uses the first search hit. Set region when the query does not already name the shop. Credit cost: 1 credit per call.
//
// GET / POST /v2/decathlon/product
func (c *Client) DecathlonProduct(ctx context.Context, params *DecathlonProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["decathlon_product"], params.params(), options)
}

// DecathlonSearch — Decathlon Search. Search a Decathlon shop with a keyword or a search URL. Choose the country with region, or let a full link pick the shop, then page through titles, prices, and product links. Credit cost: 1 credit per call.
//
// GET / POST /v2/decathlon/search
func (c *Client) DecathlonSearch(ctx context.Context, params *DecathlonSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["decathlon_search"], params.params(), options)
}

// DigikeyProduct — DigiKey Product. Get a DigiKey product page: part numbers, manufacturer, price tiers, stock, images, datasheet and specs via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/digikey/product
func (c *Client) DigikeyProduct(ctx context.Context, params *DigikeyProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["digikey_product"], params.params(), options)
}

// DigikeySearch — DigiKey Search. Search DigiKey country sites by keyword or result/filter URL. Returns category branches or product listings with pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/digikey/search
func (c *Client) DigikeySearch(ctx context.Context, params *DigikeySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["digikey_search"], params.params(), options)
}

// DoajArticle — DOAJ Article. Fetch one DOAJ article by its id or by a doaj.org article URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/doaj/article
func (c *Client) DoajArticle(ctx context.Context, params *DoajArticleParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["doaj_article"], params.params(), options)
}

// DoajJournal — DOAJ Journal. Fetch one DOAJ journal by ISSN, journal id, or a doaj.org toc URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/doaj/journal
func (c *Client) DoajJournal(ctx context.Context, params *DoajJournalParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["doaj_journal"], params.params(), options)
}

// DoajSearch — DOAJ Search. Search DOAJ articles or journals. kind is article (default) or journal. Article ids and journal ISSNs are the detail keys. Credit cost: 1 credit per call.
//
// GET / POST /v2/doaj/search
func (c *Client) DoajSearch(ctx context.Context, params *DoajSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["doaj_search"], params.params(), options)
}

// DomainDnsbl — Domain DNSBL. Check if a domain or IP address is listed in DNS-based blacklists (DNSBL), with detailed provider-level detection and category information. Credit cost: 1 credit per call.
//
// GET / POST /v2/domain/dnsbl
func (c *Client) DomainDnsbl(ctx context.Context, params *DomainDnsblParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["domain_dnsbl"], params.params(), options)
}

// DomainMalicious — Domain Malicious. Check if a domain or IP address has been flagged as malicious across multiple threat intelligence feeds, with source attribution. Credit cost: 1 credit per call.
//
// GET / POST /v2/domain/malicious
func (c *Client) DomainMalicious(ctx context.Context, params *DomainMaliciousParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["domain_malicious"], params.params(), options)
}

// DomainWhois — Domain Whois. Retrieve WHOIS registration data for any domain, including registration dates, registrar details, name servers, and contact information. Credit cost: 1 credit per call.
//
// GET / POST /v2/domain/whois
func (c *Client) DomainWhois(ctx context.Context, params *DomainWhoisParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["domain_whois"], params.params(), options)
}

// DomaincomauProperty — Domain.com.au Property. Read one Domain.com.au listing from its URL. The response is a flat object with price, address, agency, and photos. Credit cost: 2 credits per call.
//
// GET / POST /v2/domaincomau/property
func (c *Client) DomaincomauProperty(ctx context.Context, params *DomaincomauPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["domaincomau_property"], params.params(), options)
}

// DomaincomauSearch — Domain.com.au Search. Search Domain.com.au by suburb. mode is sale, rent, or sold. page starts at 1. Each card has a dollar price and a price label rebuilt from that amount. Credit cost: 2 credits per call.
//
// GET / POST /v2/domaincomau/search
func (c *Client) DomaincomauSearch(ctx context.Context, params *DomaincomauSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["domaincomau_search"], params.params(), options)
}

// DuckduckgoSearch — DuckDuckGo Search. Query DuckDuckGo via API and retrieve structured organic results with region and language. Credit cost: 1 credit per call.
//
// GET / POST /v2/duckduckgo/search
func (c *Client) DuckduckgoSearch(ctx context.Context, params *DuckduckgoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["duckduckgo_search"], params.params(), options)
}

// EbayProduct — eBay Product. eBay Product API: one listing on a supported eBay site, by URL or item id (3 credits). Credit cost: 3 credits per call.
//
// GET / POST /v2/ebay/product
func (c *Client) EbayProduct(ctx context.Context, params *EbayProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ebay_product"], params.params(), options)
}

// EbaySearch — eBay Search. eBay Search API: search listings on supported eBay sites. Returns price, seller, and pagination (3 credits). Credit cost: 3 credits per call.
//
// GET / POST /v2/ebay/search
func (c *Client) EbaySearch(ctx context.Context, params *EbaySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ebay_search"], params.params(), options)
}

// ElcorteinglesProduct — El Corte Inglés Product. El Corte Inglés Product API. Fetch a full product page from elcorteingles.es by URL or product id. Returns pricing, stock, images and identifiers via browser rendering. Credit cost: 3 credits per call.
//
// GET / POST /v2/elcorteingles/product
func (c *Client) ElcorteinglesProduct(ctx context.Context, params *ElcorteinglesProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["elcorteingles_product"], params.params(), options)
}

// ElcorteinglesSearch — El Corte Inglés Search. El Corte Inglés Search API. Search product listings on elcorteingles.es by keyword or URL. Returns prices, brands, images and pagination via webunlocker. Credit cost: 3 credits per call.
//
// GET / POST /v2/elcorteingles/search
func (c *Client) ElcorteinglesSearch(ctx context.Context, params *ElcorteinglesSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["elcorteingles_search"], params.params(), options)
}

// EleclercProduct — E.Leclerc Product. Read one E.Leclerc product by EAN or product URL. Pass store_id to get that store's price and stock. Another store returns a different price and stock. Credit cost: 1 credit per call.
//
// GET / POST /v2/eleclerc/product
func (c *Client) EleclercProduct(ctx context.Context, params *EleclercProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["eleclerc_product"], params.params(), options)
}

// EleclercSearch — E.Leclerc Search. Search for products on E.Leclerc by keyword with optional store-specific filtering and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/eleclerc/search
func (c *Client) EleclercSearch(ctx context.Context, params *EleclercSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["eleclerc_search"], params.params(), options)
}

// EleclercSuggest — E.Leclerc Suggest. Get instant product autocomplete suggestions from E.Leclerc for a given keyword or EAN. Credit cost: 1 credit per call.
//
// GET / POST /v2/eleclerc/suggest
func (c *Client) EleclercSuggest(ctx context.Context, params *EleclercSuggestParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["eleclerc_suggest"], params.params(), options)
}

// EmailAnalyzes — Email Analyzes. Analyze a domain to retrieve its email infrastructure. Microsoft 365 and Google Workspace domains, SPF records, and DNS configurations. Credit cost: 1 credit per call.
//
// GET / POST /v2/email/analyzes
func (c *Client) EmailAnalyzes(ctx context.Context, params *EmailAnalyzesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["email_analyzes"], params.params(), options)
}

// EmailFinder — Email Finder. Find verified professional email addresses by combining a person's full name with their company domain or company name. Credit cost: 1 credit per call.
//
// GET / POST /v2/email/finder
func (c *Client) EmailFinder(ctx context.Context, params *EmailFinderParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["email_finder"], params.params(), options)
}

// EmailVerify — Email Verify. Verify whether an email address is valid and deliverable, returning a clear status to prevent bounces and maintain list hygiene. Credit cost: 1 credit per call.
//
// GET / POST /v2/email/verify
func (c *Client) EmailVerify(ctx context.Context, params *EmailVerifyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["email_verify"], params.params(), options)
}

// EuropagesCompany — Europages Company. Europages company profile with address, supplier type and the products listed on the page (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/europages/company
func (c *Client) EuropagesCompany(ctx context.Context, params *EuropagesCompanyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["europages_company"], params.params(), options)
}

// EuropagesSearch — Europages Search. Search Europages suppliers by keyword or category URL. Returns company cards and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/europages/search
func (c *Client) EuropagesSearch(ctx context.Context, params *EuropagesSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["europages_search"], params.params(), options)
}

// EventbriteEvent — Eventbrite Event. Get one Eventbrite event by detail URL or numeric id: venue, organizer, schedule, and description (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/eventbrite/event
func (c *Client) EventbriteEvent(ctx context.Context, params *EventbriteEventParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["eventbrite_event"], params.params(), options)
}

// EventbriteSearch — Eventbrite Search. Search Eventbrite events by keyword or paste a city/category URL. Returns event cards and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/eventbrite/search
func (c *Client) EventbriteSearch(ctx context.Context, params *EventbriteSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["eventbrite_search"], params.params(), options)
}

// ExpediaAd — Expedia Ad. Fetch one Expedia hotel from its information URL: name, postal address, and amenity groups. Credit cost: 2 credits per call.
//
// GET / POST /v2/expedia/ad
func (c *Client) ExpediaAd(ctx context.Context, params *ExpediaAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["expedia_ad"], params.params(), options)
}

// ExpediaFlight — Expedia Flight. Search Expedia flights by airport route and dates, or by a flight search URL. Each offer includes the airline, times, stops, and displayed price. Credit cost: 2 credits per call.
//
// GET / POST /v2/expedia/flight
func (c *Client) ExpediaFlight(ctx context.Context, params *ExpediaFlightParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["expedia_flight"], params.params(), options)
}

// ExpediaSearch — Expedia Search. Execute a hotel search on Expedia using a direct search URL and receive structured hotel listings with prices, ratings, and availability. Credit cost: 2 credits per call.
//
// GET / POST /v2/expedia/search
func (c *Client) ExpediaSearch(ctx context.Context, params *ExpediaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["expedia_search"], params.params(), options)
}

// FashionphileProduct — Fashionphile Product. Load one Fashionphile product from a handle or a product URL. Keywords and search URLs belong on the search endpoint. Credit cost: 1 credit per call.
//
// GET / POST /v2/fashionphile/product
func (c *Client) FashionphileProduct(ctx context.Context, params *FashionphileProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fashionphile_product"], params.params(), options)
}

// FashionphileSearch — Fashionphile Search. Search Fashionphile by keywords, a search URL, a collection URL, or a products.json URL. A product handle or product URL belongs on the product endpoint. Credit cost: 1 credit per call.
//
// GET / POST /v2/fashionphile/search
func (c *Client) FashionphileSearch(ctx context.Context, params *FashionphileSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fashionphile_search"], params.params(), options)
}

// FinderChewyProducts — Chewy Products Finder. Find Chewy product URLs by product name, ID, or any combination of filters to use in product detail and rendering API calls.
//
// GET / POST /v2/finder/chewy-products
func (c *Client) FinderChewyProducts(ctx context.Context, params *FinderChewyProductsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_chewy_products"], params.params(), options)
}

// FinderEleclercStores — Eleclerc Stores Finder. Search Eleclerc stores by name or city to retrieve contact details, opening hours, and geolocation.
//
// GET / POST /v2/finder/eleclerc-stores
func (c *Client) FinderEleclercStores(ctx context.Context, params *FinderEleclercStoresParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_eleclerc_stores"], params.params(), options)
}

// FinderGoogleCountries — Google Countries Finder. Search the list of Google-supported countries by name or country code to use in Google Search API requests.
//
// GET / POST /v2/finder/google-countries
func (c *Client) FinderGoogleCountries(ctx context.Context, params *FinderGoogleCountriesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_google_countries"], params.params(), options)
}

// FinderGoogleLanguages — Google Languages Finder. Search the list of Google-supported languages by name to retrieve language codes for use in Google Search API requests.
//
// GET / POST /v2/finder/google-languages
func (c *Client) FinderGoogleLanguages(ctx context.Context, params *FinderGoogleLanguagesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_google_languages"], params.params(), options)
}

// FinderGoogleLocations — Google Locations Finder. Search Google-supported locations by name to retrieve location IDs and GPS coordinates for geo-targeted search API calls.
//
// GET / POST /v2/finder/google-locations
func (c *Client) FinderGoogleLocations(ctx context.Context, params *FinderGoogleLocationsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_google_locations"], params.params(), options)
}

// FinderLeroymerlinProducts — Leroy Merlin Products Finder. Find Leroy Merlin product URLs by product name, ID, or any combination of filters to use in product detail and rendering API calls.
//
// GET / POST /v2/finder/leroymerlin-products
func (c *Client) FinderLeroymerlinProducts(ctx context.Context, params *FinderLeroymerlinProductsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_leroymerlin_products"], params.params(), options)
}

// FinderLeroymerlinStores — Leroy Merlin Stores Finder. Search Leroy Merlin stores by name, city or postal code to retrieve their store ID, slug, address and store locator URL.
//
// GET / POST /v2/finder/leroymerlin-stores
func (c *Client) FinderLeroymerlinStores(ctx context.Context, params *FinderLeroymerlinStoresParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_leroymerlin_stores"], params.params(), options)
}

// FinderLinkedinIndustries — LinkedIn Industries Finder. Search LinkedIn industry taxonomy to retrieve industry IDs and hierarchies for use in LinkedIn API calls and data enrichment.
//
// GET / POST /v2/finder/linkedin-industries
func (c *Client) FinderLinkedinIndustries(ctx context.Context, params *FinderLinkedinIndustriesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_linkedin_industries"], params.params(), options)
}

// FinderMrbricolageStores — Mr.Bricolage Stores Finder. Search Mr.Bricolage stores by name, city or postal code to retrieve their store ID, slug, contact information and geolocation.
//
// GET / POST /v2/finder/mrbricolage-stores
func (c *Client) FinderMrbricolageStores(ctx context.Context, params *FinderMrbricolageStoresParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_mrbricolage_stores"], params.params(), options)
}

// FiverrProfileInfo — Fiverr Profile Info. Extract complete public profile data from any Fiverr seller, including gigs, ratings, reviews, skills, portfolio, and consultation pricing. Credit cost: 1 credit per call.
//
// GET / POST /v2/fiverr/profile/info
func (c *Client) FiverrProfileInfo(ctx context.Context, params *FiverrProfileInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fiverr_profile_info"], params.params(), options)
}

// FiverrSearch — Fiverr Search. Search Fiverr gigs using a direct search URL and receive paginated listings with seller information, gig titles, images, and seller levels. Credit cost: 1 credit per call.
//
// GET / POST /v2/fiverr/search
func (c *Client) FiverrSearch(ctx context.Context, params *FiverrSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fiverr_search"], params.params(), options)
}

// FnacProduct — Fnac Product. Read one Fnac product from its URL or product id. You get the price, brand, EAN, seller, and rating. Credit cost: 2 credits per call.
//
// GET / POST /v2/fnac/product
func (c *Client) FnacProduct(ctx context.Context, params *FnacProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fnac_product"], params.params(), options)
}

// FnacSearch — Fnac Search. Search Fnac for a keyword or a SearchResult URL. Each card has the product id, name, price, seller, and stock. Credit cost: 2 credits per call.
//
// GET / POST /v2/fnac/search
func (c *Client) FnacSearch(ctx context.Context, params *FnacSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fnac_search"], params.params(), options)
}

// FotocasaAd — Fotocasa Property Ad. Get a full Fotocasa listing by URL: price, features, coordinates, energy certificate, images, and agency (browser rendering, 2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/fotocasa/ad
func (c *Client) FotocasaAd(ctx context.Context, params *FotocasaAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fotocasa_ad"], params.params(), options)
}

// FotocasaSearch — Fotocasa Property Search. Search Fotocasa listings by city or search URL. Returns prices, rooms, area, coordinates, and agency data via browser rendering (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/fotocasa/search
func (c *Client) FotocasaSearch(ctx context.Context, params *FotocasaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fotocasa_search"], params.params(), options)
}

// FranchisedirectFranchise — Franchise Direct Franchise. Read one Franchise Direct profile: cash requirement, franchise fee, investment range, and training. Credit cost: 1 credit per call.
//
// GET / POST /v2/franchisedirect/franchise
func (c *Client) FranchisedirectFranchise(ctx context.Context, params *FranchisedirectFranchiseParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["franchisedirect_franchise"], params.params(), options)
}

// FranchisedirectSearch — Franchise Direct Search. Search Franchise Direct listings by keyword or listing URL. Each card includes a USD cash figure and a numeric minimum. Credit cost: 1 credit per call.
//
// GET / POST /v2/franchisedirect/search
func (c *Client) FranchisedirectSearch(ctx context.Context, params *FranchisedirectSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["franchisedirect_search"], params.params(), options)
}

// FressnapfProduct — Fressnapf Product. Read a Fressnapf Germany product sheet from its URL. Returns sku, GTIN, euro price, unit price, and pack variants. Credit cost: 1 credit per call.
//
// GET / POST /v2/fressnapf/product
func (c *Client) FressnapfProduct(ctx context.Context, params *FressnapfProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fressnapf_product"], params.params(), options)
}

// FressnapfSearch — Fressnapf Search. Search Fressnapf Germany by keyword or category URL. Returns product cards with euro prices and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/fressnapf/search
func (c *Client) FressnapfSearch(ctx context.Context, params *FressnapfSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["fressnapf_search"], params.params(), options)
}

// FundaProperty — Funda Property. Read one Funda listing from its URL. You get the asking price, living area, and the feature groups. Credit cost: 2 credits per call.
//
// GET / POST /v2/funda/property
func (c *Client) FundaProperty(ctx context.Context, params *FundaPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["funda_property"], params.params(), options)
}

// FundaSearch — Funda Search. Search Funda by place. Set offering to koop for sale or huur for rent, then page the listings. Credit cost: 2 credits per call.
//
// GET / POST /v2/funda/search
func (c *Client) FundaSearch(ctx context.Context, params *FundaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["funda_search"], params.params(), options)
}

// G2ProductInfo — G2 Product Info. Read a G2 product reviews page: rating, pricing, categories, alternatives, screenshots, and the reviews shown on the page. Credit cost: 2 credits per call.
//
// GET / POST /v2/g2/product/info
func (c *Client) G2ProductInfo(ctx context.Context, params *G2ProductInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["g2_product_info"], params.params(), options)
}

// GithubUserInfo — GitHub User Info. Retrieve public GitHub user profile data including bio, company, location, follower count, repositories, and associated email addresses. Credit cost: 1 credit per call.
//
// GET / POST /v2/github/user/info
func (c *Client) GithubUserInfo(ctx context.Context, params *GithubUserInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["github_user_info"], params.params(), options)
}

// GlassdoorCompanyInfo — Glassdoor Company Info. Get one Glassdoor company profile by URL or employer id (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/glassdoor/company/info
func (c *Client) GlassdoorCompanyInfo(ctx context.Context, params *GlassdoorCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["glassdoor_company_info"], params.params(), options)
}

// GlassdoorCompanySearch — Glassdoor Company Search. Search Glassdoor employers by keyword across country sites (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/glassdoor/company/search
func (c *Client) GlassdoorCompanySearch(ctx context.Context, params *GlassdoorCompanySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["glassdoor_company_search"], params.params(), options)
}

// GlassdoorJobInfo — Glassdoor Job Info. Get one Glassdoor job listing by URL or jl id (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/glassdoor/job/info
func (c *Client) GlassdoorJobInfo(ctx context.Context, params *GlassdoorJobInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["glassdoor_job_info"], params.params(), options)
}

// GlassdoorJobSearch — Glassdoor Job Search. Search Glassdoor jobs by keyword or SEO URL, with domain and filters (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/glassdoor/job/search
func (c *Client) GlassdoorJobSearch(ctx context.Context, params *GlassdoorJobSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["glassdoor_job_search"], params.params(), options)
}

// GoodreadsBook — Goodreads Book. Read one Goodreads book by the numeric id from search, such as 44767458 for Dune. Credit cost: 2 credits per call.
//
// GET / POST /v2/goodreads/book
func (c *Client) GoodreadsBook(ctx context.Context, params *GoodreadsBookParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["goodreads_book"], params.params(), options)
}

// GoodreadsSearch — Goodreads Search. Search Goodreads books by keyword. pagination.next is a Goodreads search URL for the following page. Credit cost: 2 credits per call.
//
// GET / POST /v2/goodreads/search
func (c *Client) GoodreadsSearch(ctx context.Context, params *GoodreadsSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["goodreads_search"], params.params(), options)
}

// GoogleImages — Google Images. Google Images Scraper API – bringing the power of visual search to your fingertips. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/images
func (c *Client) GoogleImages(ctx context.Context, params *GoogleImagesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_images"], params.params(), options)
}

// GoogleJobs — Google Jobs. Extract job listings from Google Careers for enhanced recruitment and market analysis. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/jobs
func (c *Client) GoogleJobs(ctx context.Context, params *GoogleJobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_jobs"], params.params(), options)
}

// GoogleMaps — Google Maps. Extract comprehensive company data from Google Maps for enhanced business insights. Credit cost: 3 credits per call.
//
// GET / POST /v2/google/maps
func (c *Client) GoogleMaps(ctx context.Context, params *GoogleMapsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_maps"], params.params(), options)
}

// GoogleNews — Google News. Extract real-time news articles from Google News with advanced filtering by topic, location, language, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/news
func (c *Client) GoogleNews(ctx context.Context, params *GoogleNewsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_news"], params.params(), options)
}

// GoogleSearch — Google Search. Scrape real-time Google search results with organic listings, knowledge graph, related searches, and advanced filters. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/search
func (c *Client) GoogleSearch(ctx context.Context, params *GoogleSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search"], params.params(), options)
}

// GoogleSearchAI — Google AI Mode. Scrape real-time Google AI Mode answers: text_blocks, references, and shopping_results. Credit cost: 3 credits per call.
//
// GET / POST /v2/google/search/ai
func (c *Client) GoogleSearchAI(ctx context.Context, params *GoogleSearchAIParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search_ai"], params.params(), options)
}

// GoogleSearchAutocomplete — Google Search Autocomplete. Retrieve real-time Google Search autocomplete suggestions for any query, with support for language and country localization. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/search/autocomplete
func (c *Client) GoogleSearchAutocomplete(ctx context.Context, params *GoogleSearchAutocompleteParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search_autocomplete"], params.params(), options)
}

// GoogleShopping — Google Shopping. Search Google Shopping and retrieve structured product listings with prices, ratings, merchants, and product IDs. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/shopping
func (c *Client) GoogleShopping(ctx context.Context, params *GoogleShoppingParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_shopping"], params.params(), options)
}

// GoogleVideos — Google Videos. Scrape video results from Google Videos search including YouTube, Reddit, and other video platforms, with location and language support. Credit cost: 1 credit per call.
//
// GET / POST /v2/google/videos
func (c *Client) GoogleVideos(ctx context.Context, params *GoogleVideosParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_videos"], params.params(), options)
}

// GoogleadstransparencyAd — Google Ads Transparency Ad. Get one Google Ads Transparency creative by URL: dates, impressions, platforms, regions, and the text read from the preview image (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/googleadstransparency/ad
func (c *Client) GoogleadstransparencyAd(ctx context.Context, params *GoogleadstransparencyAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["googleadstransparency_ad"], params.params(), options)
}

// GoogleadstransparencySearch — Google Ads Transparency Search. Search the Google Ads Transparency Center by domain, advertiser, region, format, platform, and date. Returns creative cards and pagination (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/googleadstransparency/search
func (c *Client) GoogleadstransparencySearch(ctx context.Context, params *GoogleadstransparencySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["googleadstransparency_search"], params.params(), options)
}

// GraingerProduct — Grainger Product. Get a Grainger US or Canada product page: SKU, brand, price, images, description and specs via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/grainger/product
func (c *Client) GraingerProduct(ctx context.Context, params *GraingerProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["grainger_product"], params.params(), options)
}

// GraingerSearch — Grainger Search. Search Grainger US or Canada by keyword, category URL or SKU. Returns product listings, category branches and suggestions via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/grainger/search
func (c *Client) GraingerSearch(ctx context.Context, params *GraingerSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["grainger_search"], params.params(), options)
}

// GreenhouseJob — Greenhouse Job. Fetch one Greenhouse job by board/id, such as discord/8806482002, or by a job-boards.greenhouse.io job URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/greenhouse/job
func (c *Client) GreenhouseJob(ctx context.Context, params *GreenhouseJobParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["greenhouse_job"], params.params(), options)
}

// GreenhouseSearch — Greenhouse Search. List jobs on a Greenhouse board. query is the board token, such as discord. keyword filters title and location. Credit cost: 1 credit per call.
//
// GET / POST /v2/greenhouse/search
func (c *Client) GreenhouseSearch(ctx context.Context, params *GreenhouseSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["greenhouse_search"], params.params(), options)
}

// GsaelibraryContract — GSA eLibrary Contract. Read one GSA eLibrary contractor record from the contract URL returned by search. Credit cost: 1 credit per call.
//
// GET / POST /v2/gsaelibrary/contract
func (c *Client) GsaelibraryContract(ctx context.Context, params *GsaelibraryContractParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["gsaelibrary_contract"], params.params(), options)
}

// GsaelibrarySearch — GSA eLibrary Search. Search GSA eLibrary for schedule, SIN, and contract rows by keyword or by a full search URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/gsaelibrary/search
func (c *Client) GsaelibrarySearch(ctx context.Context, params *GsaelibrarySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["gsaelibrary_search"], params.params(), options)
}

// HomestraAd — Homestra Ad. Retrieve detailed property listing data from Homestra, including price, size, location, features, photos, and nearby properties. Credit cost: 1 credit per call.
//
// GET / POST /v2/homestra/ad
func (c *Client) HomestraAd(ctx context.Context, params *HomestraAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["homestra_ad"], params.params(), options)
}

// HomestraSearch — Homestra Search. Search Homestra property listings using a direct search URL with filters, returning structured results with prices, sizes, locations, and photos. Credit cost: 1 credit per call.
//
// GET / POST /v2/homestra/search
func (c *Client) HomestraSearch(ctx context.Context, params *HomestraSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["homestra_search"], params.params(), options)
}

// HuggingfaceModel — Hugging Face Model. Read one public Hugging Face model by an id such as google-bert/bert-base-uncased. Credit cost: 1 credit per call.
//
// GET / POST /v2/huggingface/model
func (c *Client) HuggingfaceModel(ctx context.Context, params *HuggingfaceModelParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["huggingface_model"], params.params(), options)
}

// HuggingfaceSearch — Hugging Face Search. Search the public Hugging Face models API by keyword and follow pagination.next to the next cursor page. Credit cost: 1 credit per call.
//
// GET / POST /v2/huggingface/search
func (c *Client) HuggingfaceSearch(ctx context.Context, params *HuggingfaceSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["huggingface_search"], params.params(), options)
}

// IdealistaProperty — Idealista Property. Extract full property details from an Idealista listing URL, including price, specs, photos by room, features, energy rating, and advertiser data. Credit cost: 3 credits per call.
//
// GET / POST /v2/idealista/property
func (c *Client) IdealistaProperty(ctx context.Context, params *IdealistaPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["idealista_property"], params.params(), options)
}

// IdealistaProvinces — Idealista Provinces. List the 50 Spanish provinces Idealista shows for homes for sale. Each row includes the sale search URL to pass as query to Idealista Search. Madrid is madrid-provincia. Credit cost: 3 credits per call.
//
// GET / POST /v2/idealista/provinces
func (c *Client) IdealistaProvinces(ctx context.Context, params *IdealistaProvincesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["idealista_provinces"], params.params(), options)
}

// IdealistaSearch — Idealista Search. Search Idealista listings from a search URL and retrieve prices, specs, photos, agency info, and pagination. Pass a province URL from /v2/idealista/provinces (results[].url) as query. Madrid is `https://www.idealista.com/venta-vivien… Credit cost: 3 credits per call.
//
// GET / POST /v2/idealista/search
func (c *Client) IdealistaSearch(ctx context.Context, params *IdealistaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["idealista_search"], params.params(), options)
}

// IdealoProduct — Idealo Product. Get an Idealo product from Spain, France, or Germany. URL or product id. Offers, shipping, specs, and images via WebUnlocker (3 credits). Credit cost: 3 credits per call.
//
// GET / POST /v2/idealo/product
func (c *Client) IdealoProduct(ctx context.Context, params *IdealoProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["idealo_product"], params.params(), options)
}

// IdealoSearch — Idealo Search. Search Idealo Spain, France, and Germany. Keyword or search URL. Prices, shipping, ratings, and pagination via WebUnlocker (3 credits). Credit cost: 3 credits per call.
//
// GET / POST /v2/idealo/search
func (c *Client) IdealoSearch(ctx context.Context, params *IdealoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["idealo_search"], params.params(), options)
}

// IkeaProduct — IKEA Product. Extract full IKEA product details from a product URL, including price, gallery, variants, dimensions, and breadcrumb. Credit cost: 1 credit per call.
//
// GET / POST /v2/ikea/product
func (c *Client) IkeaProduct(ctx context.Context, params *IkeaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ikea_product"], params.params(), options)
}

// IkeaSearch — IKEA Search. Search IKEA products from a search URL and retrieve prices, ratings, dimensions, images, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/ikea/search
func (c *Client) IkeaSearch(ctx context.Context, params *IkeaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ikea_search"], params.params(), options)
}

// Immobilienscout24Property — ImmobilienScout24 Property. Read one ImmobilienScout24 expose by URL or numeric id. You get cold rent or purchase price, address, and area. Credit cost: 2 credits per call.
//
// GET / POST /v2/immobilienscout24/property
func (c *Client) Immobilienscout24Property(ctx context.Context, params *Immobilienscout24PropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["immobilienscout24_property"], params.params(), options)
}

// Immobilienscout24Search — ImmobilienScout24 Search. Search ImmobilienScout24 by region and property type. real_estate_type covers rent, buy, house, and land. Page with pagenumber. Credit cost: 2 credits per call.
//
// GET / POST /v2/immobilienscout24/search
func (c *Client) Immobilienscout24Search(ctx context.Context, params *Immobilienscout24SearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["immobilienscout24_search"], params.params(), options)
}

// Immoscout24chSearch — ImmoScout24 Switzerland Search. Search ImmoScout24.ch listings by city or search URL and retrieve the CHF price, rooms, living area, and address. Property pages are not available. Credit cost: 2 credits per call.
//
// GET / POST /v2/immoscout24ch/search
func (c *Client) Immoscout24chSearch(ctx context.Context, params *Immoscout24chSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["immoscout24ch_search"], params.params(), options)
}

// ImmoweltProperty — Immowelt Property. Extract full property details from an Immowelt listing URL, including price, specs, photo gallery, energy certificate, location, and contact data. Credit cost: 1 credit per call.
//
// GET / POST /v2/immowelt/property
func (c *Client) ImmoweltProperty(ctx context.Context, params *ImmoweltPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["immowelt_property"], params.params(), options)
}

// ImmoweltSearch — Immowelt Search. Search Immowelt listings from a search URL and retrieve prices, specs, photos, agency info, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/immowelt/search
func (c *Client) ImmoweltSearch(ctx context.Context, params *ImmoweltSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["immowelt_search"], params.params(), options)
}

// ImoveiswebProperty — ImoveisWeb Property. Read one Imobiliária Aquários listing from its imoveisweb.com.br URL. price_display is rebuilt from the asking amount. Credit cost: 1 credit per call.
//
// GET / POST /v2/imoveisweb/property
func (c *Client) ImoveiswebProperty(ctx context.Context, params *ImoveiswebPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["imoveisweb_property"], params.params(), options)
}

// ImoveiswebSearch — ImoveisWeb Search. Search the Imobiliária Aquários catalog on imoveisweb.com.br. Pass a catalog name such as casas. price_display is rebuilt from the asking amount. Credit cost: 1 credit per call.
//
// GET / POST /v2/imoveisweb/search
func (c *Client) ImoveiswebSearch(ctx context.Context, params *ImoveiswebSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["imoveisweb_search"], params.params(), options)
}

// IndeedCompanyInfo — Indeed Company Info. Indeed Company Info API. Full employer profile by slug or URL. Browser rendering returns headcount, revenue, industry and FAQ or review blocks. Credit cost: 2 credits per call.
//
// GET / POST /v2/indeed/company/info
func (c *Client) IndeedCompanyInfo(ctx context.Context, params *IndeedCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["indeed_company_info"], params.params(), options)
}

// IndeedCompanySearch — Indeed Company Search. Indeed Company Search API. Search fr.indeed.com employer profiles by keyword or URL. Browser rendering returns ratings, industry and profile links. Credit cost: 2 credits per call.
//
// GET / POST /v2/indeed/company/search
func (c *Client) IndeedCompanySearch(ctx context.Context, params *IndeedCompanySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["indeed_company_search"], params.params(), options)
}

// IndeedJobInfo — Indeed Job Info. Indeed Job Info API. Fetch a full job posting by viewjob URL, sponsored click URL or job key. Browser rendering returns description and structured salary. Credit cost: 2 credits per call.
//
// GET / POST /v2/indeed/job/info
func (c *Client) IndeedJobInfo(ctx context.Context, params *IndeedJobInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["indeed_job_info"], params.params(), options)
}

// IndeedJobSalary — Indeed Job Salary. Indeed Job Salary API. Career salary aggregates by job title slug or URL. Browser rendering returns median, percentiles and top paying cities. Credit cost: 2 credits per call.
//
// GET / POST /v2/indeed/job/salary
func (c *Client) IndeedJobSalary(ctx context.Context, params *IndeedJobSalaryParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["indeed_job_salary"], params.params(), options)
}

// IndeedJobSearch — Indeed Job Search. Indeed Job Search API. Search fr.indeed.com job listings by keyword or URL. Browser rendering returns cards with salary snippets and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/indeed/job/search
func (c *Client) IndeedJobSearch(ctx context.Context, params *IndeedJobSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["indeed_job_search"], params.params(), options)
}

// InetProduct — Inet Product. Read one Inet.se product by article number or URL, including SEK price, specs, and warehouse stock. Credit cost: 1 credit per call.
//
// GET / POST /v2/inet/product
func (c *Client) InetProduct(ctx context.Context, params *InetProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["inet_product"], params.params(), options)
}

// InetSearch — Inet Search. Search Inet.se by keyword or category. Prices are inc-VAT SEK, with the next page URL when more results exist. Credit cost: 1 credit per call.
//
// GET / POST /v2/inet/search
func (c *Client) InetSearch(ctx context.Context, params *InetSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["inet_search"], params.params(), options)
}

// InstagramPostInfo — Instagram Post Info. Monitor competitors and market trends by scraping Instagram posts for valuable data. Credit cost: 1 credit per call.
//
// GET / POST /v2/instagram/post/info
func (c *Client) InstagramPostInfo(ctx context.Context, params *InstagramPostInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["instagram_post_info"], params.params(), options)
}

// InstagramUserInfo — Instagram User Info. Gather user insights efficiently from Instagram for targeted marketing strategies. Credit cost: 1 credit per call.
//
// GET / POST /v2/instagram/user/info
func (c *Client) InstagramUserInfo(ctx context.Context, params *InstagramUserInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["instagram_user_info"], params.params(), options)
}

// IntermarcheProduct — Intermarché Product. Get an Intermarché product sheet for a store (shop_id). Try an EAN or product URL with shop_id=11301. 3 credits per call. Credit cost: 3 credits per call.
//
// GET / POST /v2/intermarche/product
func (c *Client) IntermarcheProduct(ctx context.Context, params *IntermarcheProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["intermarche_product"], params.params(), options)
}

// IntermarcheSearch — Intermarché Search. Search Intermarché products for a store (shop_id). Try query=lait and shop_id=11301. 3 credits per call. Credit cost: 3 credits per call.
//
// GET / POST /v2/intermarche/search
func (c *Client) IntermarcheSearch(ctx context.Context, params *IntermarcheSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["intermarche_search"], params.params(), options)
}

// IntersportProduct — Intersport Product. Read one Intersport product from its URL. You get price, brand, and the product details. Credit cost: 2 credits per call.
//
// GET / POST /v2/intersport/product
func (c *Client) IntersportProduct(ctx context.Context, params *IntersportProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["intersport_product"], params.params(), options)
}

// IntersportSearch — Intersport Search. Search Intersport by keyword and page the sports grid. Cards include title, price, and product link. Credit cost: 2 credits per call.
//
// GET / POST /v2/intersport/search
func (c *Client) IntersportSearch(ctx context.Context, params *IntersportSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["intersport_search"], params.params(), options)
}

// JameseditionProperty — JamesEdition Property. Read one JamesEdition listing from its URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/jamesedition/property
func (c *Client) JameseditionProperty(ctx context.Context, params *JameseditionPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jamesedition_property"], params.params(), options)
}

// JameseditionSearch — JamesEdition Search. Search JamesEdition luxury homes by keyword or by a results URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/jamesedition/search
func (c *Client) JameseditionSearch(ctx context.Context, params *JameseditionSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jamesedition_search"], params.params(), options)
}

// Jobs — Jobs. Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: CompletedJob, Error, NotFound or BadRequest.
//
// GET /v1/jobs/{id}
func (c *Client) Jobs(ctx context.Context, params *JobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jobs"], params.params(), options)
}

// KickUserInfo — Kick User Info. Retrieve complete public channel data from Kick including streamer profile, follower count, live status, and recent categories. Credit cost: 1 credit per call.
//
// GET / POST /v2/kick/user/info
func (c *Client) KickUserInfo(ctx context.Context, params *KickUserInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["kick_user_info"], params.params(), options)
}

// KijijiAd — Kijiji Ad. Read one Kijiji ad from its id or URL. You get the asking price, the title, and the city. Credit cost: 1 credit per call.
//
// GET / POST /v2/kijiji/ad
func (c *Client) KijijiAd(ctx context.Context, params *KijijiAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["kijiji_ad"], params.params(), options)
}

// KijijiSearch — Kijiji Search. Search Kijiji with a results URL or a city and category. Language can be English or French. Page through classified ads. Credit cost: 1 credit per call.
//
// GET / POST /v2/kijiji/search
func (c *Client) KijijiSearch(ctx context.Context, params *KijijiSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["kijiji_search"], params.params(), options)
}

// KrogerProduct — Kroger Product. Read one Kroger product by UPC. The sheet uses the same default store as search. You get the shelf price and the previous price when Kroger shows one. Credit cost: 2 credits per call.
//
// GET / POST /v2/kroger/product
func (c *Client) KrogerProduct(ctx context.Context, params *KrogerProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["kroger_product"], params.params(), options)
}

// KrogerSearch — Kroger Search. Search Kroger groceries by keyword and page the results. Prices come from the default pickup store returned in the response. A location id you pass is not applied. Credit cost: 2 credits per call.
//
// GET / POST /v2/kroger/search
func (c *Client) KrogerSearch(ctx context.Context, params *KrogerSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["kroger_search"], params.params(), options)
}

// LamudiAd — Lamudi Ad. Get a full Lamudi listing by URL: price, area, amenities, coordinates, images, and agent (browser rendering, 2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/lamudi/ad
func (c *Client) LamudiAd(ctx context.Context, params *LamudiAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["lamudi_ad"], params.params(), options)
}

// LamudiSearch — Lamudi Search. Search Lamudi Philippines and Indonesia listings. Pass a search URL in query, or use location and filters. Prices, area, rooms, tags, and pagination via browser rendering (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/lamudi/search
func (c *Client) LamudiSearch(ctx context.Context, params *LamudiSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["lamudi_search"], params.params(), options)
}

// LeboncoinAd — Leboncoin Ad. Extract full details from a Leboncoin ad including price, location, attributes, seller info, and images. Returns structured JSON by default, or raw HTML on demand. Credit cost: 1 credit per call.
//
// GET / POST /v2/leboncoin/ad
func (c *Client) LeboncoinAd(ctx context.Context, params *LeboncoinAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["leboncoin_ad"], params.params(), options)
}

// LeboncoinSearch — Leboncoin Search. Scrape Leboncoin search results or category pages by URL, returning paginated listings with price, location, and attributes. Returns structured JSON. return_page_source is deprecated and returns 400. Credit cost: 1 credit per call.
//
// GET / POST /v2/leboncoin/search
func (c *Client) LeboncoinSearch(ctx context.Context, params *LeboncoinSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["leboncoin_search"], params.params(), options)
}

// LeboncoinSearchAPI — Leboncoin Search API. Search Leboncoin listings via structured POST API with advanced filters for category, location, price range, seller type, and sorting. Credit cost: 1 credit per call.
//
// POST /v2/leboncoin/search_api
func (c *Client) LeboncoinSearchAPI(ctx context.Context, params *LeboncoinSearchAPIParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["leboncoin_search_api"], params.params(), options)
}

// LenovoProduct — Lenovo Product. Extract a Lenovo product page by URL: id, name, price, currency, availability, images, and description (browser rendering, 2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/lenovo/product
func (c *Client) LenovoProduct(ctx context.Context, params *LenovoProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["lenovo_product"], params.params(), options)
}

// LenovoSearch — Lenovo Search. Search Lenovo products by keyword or listing URL. Returns product id, price, currency, series, and availability via browser rendering (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/lenovo/search
func (c *Client) LenovoSearch(ctx context.Context, params *LenovoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["lenovo_search"], params.params(), options)
}

// LeroymerlinProduct — Leroy Merlin Product. Live Leroy Merlin product detail for France and Romania from a URL or identifier: offers, currency, images, features and store context. Credit cost: 3 credits per call.
//
// GET / POST /v2/leroymerlin/product
func (c *Client) LeroymerlinProduct(ctx context.Context, params *LeroymerlinProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["leroymerlin_product"], params.params(), options)
}

// LeroymerlinSearch — Leroy Merlin Search. Live Leroy Merlin search across France and Romania: title, URL, price, currency, brand, seller, rating and pagination. Credit cost: 3 credits per call.
//
// GET / POST /v2/leroymerlin/search
func (c *Client) LeroymerlinSearch(ctx context.Context, params *LeroymerlinSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["leroymerlin_search"], params.params(), options)
}

// LidlProduct — Lidl Product. Lidl Product API: fetch full product page data from a Lidl URL or ERP id. Returns pricing, images, description, availability and category via Nuxt SSR parsing. Credit cost: 1 credit per call.
//
// GET / POST /v2/lidl/product
func (c *Client) LidlProduct(ctx context.Context, params *LidlProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["lidl_product"], params.params(), options)
}

// LidlSearch — Lidl Search. Lidl Search API: search product listings by keyword or URL across Lidl online shops in 27 countries. Parses Nuxt SSR payload from HTML; returns prices, ratings, brand and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/lidl/search
func (c *Client) LidlSearch(ctx context.Context, params *LidlSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["lidl_search"], params.params(), options)
}

// LinkedinCompanyInfo — LinkedIn Company Info. LinkedIn Company Info API. Fetch firmographics by LinkedIn slug, URL, company ID, or reverse domain lookup (guideflow.com). Industry, headcount, HQ, posts, and more. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/company/info
func (c *Client) LinkedinCompanyInfo(ctx context.Context, params *LinkedinCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_company_info"], params.params(), options)
}

// LinkedinJobCount — LinkedIn Job Count. Obtain real-time job counts on LinkedIn by industry, region, and role for strategic planning. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/job/count
func (c *Client) LinkedinJobCount(ctx context.Context, params *LinkedinJobCountParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_job_count"], params.params(), options)
}

// LinkedinJobInfo — LinkedIn Job Info. Dive into LinkedIn job data and trends for in-depth labor market insights. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/job/info
func (c *Client) LinkedinJobInfo(ctx context.Context, params *LinkedinJobInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_job_info"], params.params(), options)
}

// LinkedinJobSearch — LinkedIn Job Search. Experience advanced job data extraction from LinkedIn, enhancing your recruitment strategy. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/job/search
func (c *Client) LinkedinJobSearch(ctx context.Context, params *LinkedinJobSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_job_search"], params.params(), options)
}

// LinkedinJobSuggest — LinkedIn Job Suggest. Get personalized job suggestions based on keywords or locations to streamline your job search. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/job/suggest
func (c *Client) LinkedinJobSuggest(ctx context.Context, params *LinkedinJobSuggestParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_job_suggest"], params.params(), options)
}

// LinkedinPostInfo — LinkedIn Post Info. Simplify LinkedIn post retrieval with our efficient API for enhanced content analysis. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/post/info
func (c *Client) LinkedinPostInfo(ctx context.Context, params *LinkedinPostInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_post_info"], params.params(), options)
}

// LinkedinProductInfo — LinkedIn Product Info. Gain competitive insights with detailed LinkedIn product information for strategic advantage. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/product/info
func (c *Client) LinkedinProductInfo(ctx context.Context, params *LinkedinProductInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_product_info"], params.params(), options)
}

// LinkedinProfileInfo — LinkedIn Profile Info. Streamline recruitment and business development by extracting valuable LinkedIn profile insights. Credit cost: 1 credit per call.
//
// GET / POST /v2/linkedin/profile/info
func (c *Client) LinkedinProfileInfo(ctx context.Context, params *LinkedinProfileInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedin_profile_info"], params.params(), options)
}

// LinkedinadlibraryAd — LinkedIn Ad Library Ad. Get one LinkedIn Ad Library creative by URL or id: copy, CTA, payer, impressions, and targeting (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/linkedinadlibrary/ad
func (c *Client) LinkedinadlibraryAd(ctx context.Context, params *LinkedinadlibraryAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedinadlibrary_ad"], params.params(), options)
}

// LinkedinadlibrarySearch — LinkedIn Ad Library Search. Search the LinkedIn Ad Library by advertiser, keyword, country, and date. Returns ad cards and pagination (2 credits). Credit cost: 2 credits per call.
//
// GET / POST /v2/linkedinadlibrary/search
func (c *Client) LinkedinadlibrarySearch(ctx context.Context, params *LinkedinadlibrarySearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["linkedinadlibrary_search"], params.params(), options)
}

// LoopnetProperty — LoopNet Property. Read one LoopNet commercial listing from its URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/loopnet/property
func (c *Client) LoopnetProperty(ctx context.Context, params *LoopnetPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["loopnet_property"], params.params(), options)
}

// LoopnetSearch — LoopNet Search. Search LoopNet commercial listings by city or by a results URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/loopnet/search
func (c *Client) LoopnetSearch(ctx context.Context, params *LoopnetSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["loopnet_search"], params.params(), options)
}

// MaerskTracking — Maersk Tracking. Track a Maersk container or BOL: origin, destination, current milestone and full transport-plan events via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/maersk/tracking
func (c *Client) MaerskTracking(ctx context.Context, params *MaerskTrackingParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["maersk_tracking"], params.params(), options)
}

// ManomanoProduct — ManoMano Product. Read a ManoMano product URL. The body is the product page: brand and seller are objects, and the current price is price.primary_price.amount_with_vat. Credit cost: 2 credits per call.
//
// GET / POST /v2/manomano/product
func (c *Client) ManomanoProduct(ctx context.Context, params *ManomanoProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["manomano_product"], params.params(), options)
}

// ManomanoProductOffer — ManoMano Product Offer. List the other sellers for a ManoMano product id. Each offer adds price and price_display on top of the site offer. Credit cost: 1 credit per call.
//
// GET / POST /v2/manomano/product/offer
func (c *Client) ManomanoProductOffer(ctx context.Context, params *ManomanoProductOfferParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["manomano_product_offer"], params.params(), options)
}

// ManomanoSearch — ManoMano Search. Read a ManoMano category URL or keyword. Each card keeps the site fields and adds prices.main, brand_name, and image_fullpath. pagination.next is a boolean. Credit cost: 2 credits per call.
//
// GET / POST /v2/manomano/search
func (c *Client) ManomanoSearch(ctx context.Context, params *ManomanoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["manomano_search"], params.params(), options)
}

// MascusAd — Mascus Ad. Get full details for a Mascus used equipment listing: specs, price, seller and description. Credit cost: 1 credit per call.
//
// GET / POST /v2/mascus/ad
func (c *Client) MascusAd(ctx context.Context, params *MascusAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mascus_ad"], params.params(), options)
}

// MascusSearch — Mascus Search. Search used equipment listings on Mascus from a category URL, with pagination that keeps country and category filters. Credit cost: 1 credit per call.
//
// GET / POST /v2/mascus/search
func (c *Client) MascusSearch(ctx context.Context, params *MascusSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mascus_search"], params.params(), options)
}

// MediamarktProduct — MediaMarkt Product. Read one MediaMarkt product by id or product URL, in the region you choose. You get price, currency, EAN, and availability. Credit cost: 2 credits per call.
//
// GET / POST /v2/mediamarkt/product
func (c *Client) MediamarktProduct(ctx context.Context, params *MediamarktProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mediamarkt_product"], params.params(), options)
}

// MediamarktSearch — MediaMarkt Search. Search a MediaMarkt shop by keyword. Set region for Switzerland, Germany, Austria, and the other MediaMarktSaturn countries, then page the results. Credit cost: 2 credits per call.
//
// GET / POST /v2/mediamarkt/search
func (c *Client) MediamarktSearch(ctx context.Context, params *MediamarktSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mediamarkt_search"], params.params(), options)
}

// MicromaniaProduct — Micromania Product. Read one Micromania product from its URL or product id. You get the price, brand, EAN, platform, and stock. Credit cost: 2 credits per call.
//
// GET / POST /v2/micromania/product
func (c *Client) MicromaniaProduct(ctx context.Context, params *MicromaniaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["micromania_product"], params.params(), options)
}

// MicromaniaSearch — Micromania Search. Search Micromania for a keyword or a Search-Show URL. Each card has the product id, name, brand, EAN, price, and stock. Credit cost: 2 credits per call.
//
// GET / POST /v2/micromania/search
func (c *Client) MicromaniaSearch(ctx context.Context, params *MicromaniaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["micromania_search"], params.params(), options)
}

// MobileAd — Mobile Ad. Mobile Ad API. Fetch full car listing data from a mobile.de ad URL or numeric id. Returns pricing, specs, images and features via webunlocker. Credit cost: 3 credits per call.
//
// GET / POST /v2/mobile/ad
func (c *Client) MobileAd(ctx context.Context, params *MobileAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mobile_ad"], params.params(), options)
}

// MobileSearch — Mobile Search. Mobile Search API. Same mobile.de inventory for every locale. locale sets the language label and ad links, not which cars are returned. (3 credits) Credit cost: 3 credits per call.
//
// GET / POST /v2/mobile/search
func (c *Client) MobileSearch(ctx context.Context, params *MobileSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mobile_search"], params.params(), options)
}

// MoreleProduct — Morele Product. Read one morele.net product page. Flat JSON with PLN price, EAN, specs, and availability. Credit cost: 1 credit per call.
//
// GET / POST /v2/morele/product
func (c *Client) MoreleProduct(ctx context.Context, params *MoreleProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["morele_product"], params.params(), options)
}

// MoreleSearch — Morele Search. Search morele.net by keyword or search URL. Returns product cards with PLN prices and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/morele/search
func (c *Client) MoreleSearch(ctx context.Context, params *MoreleSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["morele_search"], params.params(), options)
}

// MrbricolageSuggest — Mr.Bricolage Suggest. Live Mr.Bricolage product suggestions for a search keyword or EAN, including title, URL, image, price and currency. Credit cost: 3 credits per call.
//
// GET / POST /v2/mrbricolage/suggest
func (c *Client) MrbricolageSuggest(ctx context.Context, params *MrbricolageSuggestParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["mrbricolage_suggest"], params.params(), options)
}

// NeweggProduct — Newegg Product. Read a Newegg product from an item number, catalog id, or URL. Returns price, specs, images, and seller. Credit cost: 2 credits per call.
//
// GET / POST /v2/newegg/product
func (c *Client) NeweggProduct(ctx context.Context, params *NeweggProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["newegg_product"], params.params(), options)
}

// NeweggSearch — Newegg Search. Search Newegg by keyword or search URL. Returns item numbers, dollar prices, stock, and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/newegg/search
func (c *Client) NeweggSearch(ctx context.Context, params *NeweggSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["newegg_search"], params.params(), options)
}

// NorthdataCompany — North Data Company. North Data company profile: register, address, financials and history (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/northdata/company
func (c *Client) NorthdataCompany(ctx context.Context, params *NorthdataCompanyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["northdata_company"], params.params(), options)
}

// NorthdataSearch — North Data Search. Search North Data companies and people by name. Returns suggestion cards and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/northdata/search
func (c *Client) NorthdataSearch(ctx context.Context, params *NorthdataSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["northdata_search"], params.params(), options)
}

// NPIProvider — NPI Provider. Fetch one NPPES provider or organization by 10-digit NPI or provider-view URL, including addresses and taxonomies. Credit cost: 1 credit per call.
//
// GET / POST /v2/npi/provider
func (c *Client) NPIProvider(ctx context.Context, params *NPIProviderParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["npi_provider"], params.params(), options)
}

// NPISearch — NPI Search. Search the NPPES NPI registry by person name, organization, or 10-digit NPI. With NPI-2, a text query is the organization name. total_results is null, and next is set only when the page is full. Credit cost: 1 credit per call.
//
// GET / POST /v2/npi/search
func (c *Client) NPISearch(ctx context.Context, params *NPISearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["npi_search"], params.params(), options)
}

// OceanioSearch — Oceanio Search. Search Ocean.io by domain or company name to find the matching company, similar companies, contacts, and people. Credit cost: 1 credit per call.
//
// GET / POST /v2/oceanio/search
func (c *Client) OceanioSearch(ctx context.Context, params *OceanioSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["oceanio_search"], params.params(), options)
}

// OlxAd — OLX Ad. Read one OLX ad from its id or URL. You get the asking price, the city, and the seller name. Credit cost: 1 credit per call.
//
// GET / POST /v2/olx/ad
func (c *Client) OlxAd(ctx context.Context, params *OlxAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["olx_ad"], params.params(), options)
}

// OlxProduct — OLX Product. Read one OLX ad from its id or URL. You get the asking price, the city, and the seller name. Credit cost: 1 credit per call.
//
// GET / POST /v2/olx/product
func (c *Client) OlxProduct(ctx context.Context, params *OlxProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["olx_product"], params.params(), options)
}

// OlxSearch — OLX Search. Search OLX Brazil classifieds by keyword and get title, price, location, seller type, and condition as structured JSON in real time. Credit cost: 1 credit per call.
//
// GET / POST /v2/olx/search
func (c *Client) OlxSearch(ctx context.Context, params *OlxSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["olx_search"], params.params(), options)
}

// OscaroProduct — Oscaro Product. Oscaro Product API. Fetch full auto parts product data from an Oscaro product URL. Returns pricing, brand, reference, images and availability via browser rendering and JSON-LD parsing. Credit cost: 2 credits per call.
//
// GET / POST /v2/oscaro/product
func (c *Client) OscaroProduct(ctx context.Context, params *OscaroProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["oscaro_product"], params.params(), options)
}

// OscaroSearch — Oscaro Search. Oscaro Search API. Search auto parts by keyword, search URL or category listing on oscaro.com, oscaro.es and oscaro.pt. Returns prices and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/oscaro/search
func (c *Client) OscaroSearch(ctx context.Context, params *OscaroSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["oscaro_search"], params.params(), options)
}

// OvokoAd — Ovoko Ad. Extract an Ovoko listing from a SKU URL: price, currency, images, seller, quality, vehicle make, part numbers, and warranty. Credit cost: 1 credit per call.
//
// GET / POST /v2/ovoko/ad
func (c *Client) OvokoAd(ctx context.Context, params *OvokoAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ovoko_ad"], params.params(), options)
}

// OvokoSearch — Ovoko Search. Search Ovoko used-parts listings from a search URL and retrieve title, price, OEM codes, seller, vehicle text, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/ovoko/search
func (c *Client) OvokoSearch(ctx context.Context, params *OvokoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ovoko_search"], params.params(), options)
}

// OwlerCompanyInfo — Owler Company Info. Retrieve detailed company intelligence from Owler, including revenue, employee count, CEO info, competitors, funding, and industry sectors. Credit cost: 1 credit per call.
//
// GET / POST /v2/owler/company/info
func (c *Client) OwlerCompanyInfo(ctx context.Context, params *OwlerCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["owler_company_info"], params.params(), options)
}

// OwlerSearch — Owler Search. Search for companies on Owler by name or keyword, returning a list of matching companies with slugs, logos, domains, and profile URLs. Credit cost: 1 credit per call.
//
// GET / POST /v2/owler/search
func (c *Client) OwlerSearch(ctx context.Context, params *OwlerSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["owler_search"], params.params(), options)
}

// PagesjaunesPageInfo — PagesJaunes Page Info. Retrieve detailed business information from a PagesJaunes page, including contact details, opening hours, services, payment methods, and SIREN number. Credit cost: 1 credit per call.
//
// GET / POST /v2/pagesjaunes/page/info
func (c *Client) PagesjaunesPageInfo(ctx context.Context, params *PagesjaunesPageInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pagesjaunes_page_info"], params.params(), options)
}

// PagesjaunesSearch — PagesJaunes Search. Search PagesJaunes for French local businesses by category and location, returning paginated results with contact details and ratings. Credit cost: 1 credit per call.
//
// GET / POST /v2/pagesjaunes/search
func (c *Client) PagesjaunesSearch(ctx context.Context, params *PagesjaunesSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pagesjaunes_search"], params.params(), options)
}

// PatreonProfileInfo — Patreon Profile Info. Retrieve a Patreon creator profile with patron counts, membership tiers, campaign summary, and creator metadata. Credit cost: 1 credit per call.
//
// GET / POST /v2/patreon/profile/info
func (c *Client) PatreonProfileInfo(ctx context.Context, params *PatreonProfileInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["patreon_profile_info"], params.params(), options)
}

// PatreonSearch — Patreon Search. Search Patreon creators by keyword and retrieve campaign cards with patron counts, post counts, and profile URLs. Credit cost: 1 credit per call.
//
// GET / POST /v2/patreon/search
func (c *Client) PatreonSearch(ctx context.Context, params *PatreonSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["patreon_search"], params.params(), options)
}

// PetsmartProduct — PetSmart Product. Read one PetSmart product URL, including the selected size, dollar price, flavor, and item number. Credit cost: 1 credit per call.
//
// GET / POST /v2/petsmart/product
func (c *Client) PetsmartProduct(ctx context.Context, params *PetsmartProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["petsmart_product"], params.params(), options)
}

// PetsmartSearch — PetSmart Search. Search PetSmart by keyword or search URL. Page through pet food, litter, and supplies with dollar prices. Credit cost: 1 credit per call.
//
// GET / POST /v2/petsmart/search
func (c *Client) PetsmartSearch(ctx context.Context, params *PetsmartSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["petsmart_search"], params.params(), options)
}

// PicardProduct — Picard Product. Picard Product API. Fetch full frozen food product data from a Picard URL or reference. Returns pricing, images, description, Nutri-Score, Planet-Score, breadcrumb and reviews. Credit cost: 1 credit per call.
//
// GET / POST /v2/picard/product
func (c *Client) PicardProduct(ctx context.Context, params *PicardProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["picard_product"], params.params(), options)
}

// PicardSearch — Picard Search. Picard Search API. Search frozen food listings on picard.fr by keyword or URL. Parses SFCC Search-UpdateGrid HTML; returns prices, Nutri-Score, ratings and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/picard/search
func (c *Client) PicardSearch(ctx context.Context, params *PicardSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["picard_search"], params.params(), options)
}

// PinterestUserInfo — Pinterest User Info. Access Pinterest user profiles and analytics for targeted marketing and audience insights. Credit cost: 1 credit per call.
//
// GET / POST /v2/pinterest/user/info
func (c *Client) PinterestUserInfo(ctx context.Context, params *PinterestUserInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pinterest_user_info"], params.params(), options)
}

// PitchbookCompanyInfo — PitchBook Company Info. Get a PitchBook company profile: overview, funding, competitors and more via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/pitchbook/company/info
func (c *Client) PitchbookCompanyInfo(ctx context.Context, params *PitchbookCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pitchbook_company_info"], params.params(), options)
}

// PitchbookSearch — PitchBook Search. Search PitchBook for companies, news and profiles by keyword or search URL via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/pitchbook/search
func (c *Client) PitchbookSearch(ctx context.Context, params *PitchbookSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pitchbook_search"], params.params(), options)
}

// PointpProduct — Point.P Product. Point.P Product API: scrape live /p/ product pages from a full URL. Get title, images, EAN, pricing, specifications, and breadcrumbs as JSON via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/pointp/product
func (c *Client) PointpProduct(ctx context.Context, params *PointpProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pointp_product"], params.params(), options)
}

// PointpSearch — Point.P Search. Point.P Search API: search by keyword or URL and get building-material listings with prices, brands, ratings, units, and pagination as JSON via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/pointp/search
func (c *Client) PointpSearch(ctx context.Context, params *PointpSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["pointp_search"], params.params(), options)
}

// PrnewswireArticle — PR Newswire Article. Read one PR Newswire press release from its URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/prnewswire/article
func (c *Client) PrnewswireArticle(ctx context.Context, params *PrnewswireArticleParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["prnewswire_article"], params.params(), options)
}

// PrnewswireSearch — PR Newswire Search. Search PR Newswire press releases by keyword or by a news-search URL. Credit cost: 1 credit per call.
//
// GET / POST /v2/prnewswire/search
func (c *Client) PrnewswireSearch(ctx context.Context, params *PrnewswireSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["prnewswire_search"], params.params(), options)
}

// ProducthuntProductInfo — ProductHunt Product Info. Retrieve product data from ProductHunt, including tagline, categories, screenshots, reviews rating, and social links. Credit cost: 1 credit per call.
//
// GET / POST /v2/producthunt/product/info
func (c *Client) ProducthuntProductInfo(ctx context.Context, params *ProducthuntProductInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["producthunt_product_info"], params.params(), options)
}

// ProshopProduct — Proshop Product. Read one proshop.dk product page. Flat JSON with DKK price, EAN, specs, and stock text. Credit cost: 1 credit per call.
//
// GET / POST /v2/proshop/product
func (c *Client) ProshopProduct(ctx context.Context, params *ProshopProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["proshop_product"], params.params(), options)
}

// ProshopSearch — Proshop Search. Search proshop.dk by keyword or search URL. Returns product cards with DKK prices and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/proshop/search
func (c *Client) ProshopSearch(ctx context.Context, params *ProshopSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["proshop_search"], params.params(), options)
}

// RappiProduct — Rappi Product. Fetch one Rappi Mexico product, including the primary store price and the other store offers. The primary price can differ from a search card. Credit cost: 1 credit per call.
//
// GET / POST /v2/rappi/product
func (c *Client) RappiProduct(ctx context.Context, params *RappiProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rappi_product"], params.params(), options)
}

// RappiSearch — Rappi Search. Search Rappi Mexico on rappi.com.mx. The city resolves to Mexico City. Each row is one store offer, and next is null. Credit cost: 1 credit per call.
//
// GET / POST /v2/rappi/search
func (c *Client) RappiSearch(ctx context.Context, params *RappiSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rappi_search"], params.params(), options)
}

// RatemyprofessorsProfessor — Rate My Professors Professor. Read one Rate My Professors profile, including difficulty and would-take-again when the page shows them. Credit cost: 2 credits per call.
//
// GET / POST /v2/ratemyprofessors/professor
func (c *Client) RatemyprofessorsProfessor(ctx context.Context, params *RatemyprofessorsProfessorParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ratemyprofessors_professor"], params.params(), options)
}

// RatemyprofessorsSearch — Rate My Professors Search. Search Rate My Professors by name. Cards include school, department, difficulty, and would-take-again when the card shows them. Credit cost: 2 credits per call.
//
// GET / POST /v2/ratemyprofessors/search
func (c *Client) RatemyprofessorsSearch(ctx context.Context, params *RatemyprofessorsSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ratemyprofessors_search"], params.params(), options)
}

// RealestatecomauSearch — realestate.com.au Search. Search realestate.com.au buy, rent, or sold results. A place such as Sydney, NSW defaults to homes for sale. Each card has price, address, rooms, and a property URL. Credit cost: 2 credits per call.
//
// GET / POST /v2/realestatecomau/search
func (c *Client) RealestatecomauSearch(ctx context.Context, params *RealestatecomauSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["realestatecomau_search"], params.params(), options)
}

// RealtorProperty — Realtor Property. Extract full property details from a Realtor.com listing URL, including price, specs, photos, agents, features, and location data. Credit cost: 2 credits per call.
//
// GET / POST /v2/realtor/property
func (c *Client) RealtorProperty(ctx context.Context, params *RealtorPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["realtor_property"], params.params(), options)
}

// RealtorSearch — Realtor Search. Search Realtor.com listings from a search URL and retrieve prices, beds, baths, sqft, photos, and property metadata. Credit cost: 2 credits per call.
//
// GET / POST /v2/realtor/search
func (c *Client) RealtorSearch(ctx context.Context, params *RealtorSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["realtor_search"], params.params(), options)
}

// RedfinProperty — Redfin Property. Read one for-sale Redfin home. Pass the url from search, which already includes the market id. The response is a flat listing, the same fields as one search result. Credit cost: 1 credit per call.
//
// GET / POST /v2/redfin/property
func (c *Client) RedfinProperty(ctx context.Context, params *RedfinPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["redfin_property"], params.params(), options)
}

// RedfinSearch — Redfin Search. Search for-sale homes on Redfin from a city, neighborhood, or county link. Each listing url includes the market id to pass to the property endpoint. Credit cost: 1 credit per call.
//
// GET / POST /v2/redfin/search
func (c *Client) RedfinSearch(ctx context.Context, params *RedfinSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["redfin_search"], params.params(), options)
}

// ReicheltProduct — Reichelt Product. Read one reichelt.de product URL, including article number, EAN, and the euro price. Credit cost: 1 credit per call.
//
// GET / POST /v2/reichelt/product
func (c *Client) ReicheltProduct(ctx context.Context, params *ReicheltProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["reichelt_product"], params.params(), options)
}

// ReicheltSearch — Reichelt Search. Search reichelt.de by keyword or category. Euro prices are rebuilt from the number, including whole-euro amounts. Credit cost: 1 credit per call.
//
// GET / POST /v2/reichelt/search
func (c *Client) ReicheltSearch(ctx context.Context, params *ReicheltSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["reichelt_search"], params.params(), options)
}

// ReverbProduct — Reverb Product. Reverb Product API: fetch full listing data from a Reverb URL or listing id. Returns pricing, specs, images, description and shop info via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/reverb/product
func (c *Client) ReverbProduct(ctx context.Context, params *ReverbProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["reverb_product"], params.params(), options)
}

// ReverbSearch — Reverb Search. Reverb Search API: search music gear listings by keyword, query string or marketplace URL on Reverb. Returns prices, thumbnails and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/reverb/search
func (c *Client) ReverbSearch(ctx context.Context, params *ReverbSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["reverb_search"], params.params(), options)
}

// RexelProduct — Rexel Product. Get a Rexel France product page: references, EAN, descriptions, images and specs via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/rexel/product
func (c *Client) RexelProduct(ctx context.Context, params *RexelProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rexel_product"], params.params(), options)
}

// RexelSearch — Rexel Search. Search Rexel France by keyword, EAN or search URL. Returns product cards and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/rexel/search
func (c *Client) RexelSearch(ctx context.Context, params *RexelSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rexel_search"], params.params(), options)
}

// RightmoveProperty — Rightmove Property. Read one Rightmove property from its URL or numeric id. The response is a flat object with price, tenure, floor area, stations, and the agent. Credit cost: 2 credits per call.
//
// GET / POST /v2/rightmove/property
func (c *Client) RightmoveProperty(ctx context.Context, params *RightmovePropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rightmove_property"], params.params(), options)
}

// RightmoveSearch — Rightmove Search. Search Rightmove by a UK place name or a search URL. Each card has a numeric price, a price label rebuilt from that number, the address, the agent, and a property URL. Credit cost: 2 credits per call.
//
// GET / POST /v2/rightmove/search
func (c *Client) RightmoveSearch(ctx context.Context, params *RightmoveSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rightmove_search"], params.params(), options)
}

// RsonlineProduct — RS Online Product. Read one RS Online UK product URL, including the RS stock number, manufacturer part, and pound price. Credit cost: 2 credits per call.
//
// GET / POST /v2/rsonline/product
func (c *Client) RsonlineProduct(ctx context.Context, params *RsonlineProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rsonline_product"], params.params(), options)
}

// RsonlineSearch — RS Online Search. Search RS Online UK by keyword or category. Pound prices keep the amount printed on the page. Credit cost: 2 credits per call.
//
// GET / POST /v2/rsonline/search
func (c *Client) RsonlineSearch(ctx context.Context, params *RsonlineSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["rsonline_search"], params.params(), options)
}

// SeekJob — SEEK Job. Read one SEEK Australia job from its URL or numeric id. Credit cost: 2 credits per call.
//
// GET / POST /v2/seek/job
func (c *Client) SeekJob(ctx context.Context, params *SeekJobParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seek_job"], params.params(), options)
}

// SeekSearch — SEEK Search. Search SEEK Australia jobs by keyword or by a results URL. Credit cost: 2 credits per call.
//
// GET / POST /v2/seek/search
func (c *Client) SeekSearch(ctx context.Context, params *SeekSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seek_search"], params.params(), options)
}

// SelogerAgency — SeLoger Agency. Retrieve detailed profile information for a SeLoger real estate agency including contact details, ratings, reviews, and address. Credit cost: 2 credits per call.
//
// GET / POST /v2/seloger/agency
func (c *Client) SelogerAgency(ctx context.Context, params *SelogerAgencyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seloger_agency"], params.params(), options)
}

// SelogerAgencyRent — SeLoger Agency Rent. Retrieve paginated rental property listings published by a specific SeLoger agency using its numeric agency ID. Credit cost: 2 credits per call.
//
// GET / POST /v2/seloger/agency/rent
func (c *Client) SelogerAgencyRent(ctx context.Context, params *SelogerAgencyRentParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seloger_agency_rent"], params.params(), options)
}

// SelogerAgencySale — SeLoger Agency Sale. Retrieve paginated sale property listings published by a specific SeLoger agency using its numeric agency ID. Credit cost: 2 credits per call.
//
// GET / POST /v2/seloger/agency/sale
func (c *Client) SelogerAgencySale(ctx context.Context, params *SelogerAgencySaleParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seloger_agency_sale"], params.params(), options)
}

// SelogerProperty — SeLoger Property. Efficiently extract detailed data from a SeLoger real estate listing including agency info, pricing, photos, description, and property criteria. Credit cost: 2 credits per call.
//
// GET / POST /v2/seloger/property
func (c *Client) SelogerProperty(ctx context.Context, params *SelogerPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seloger_property"], params.params(), options)
}

// SelogerSearch — SeLoger Search. Efficiently extract real estate listings from SeLoger search pages for market insights, property management, and investment analysis. Credit cost: 2 credits per call.
//
// GET / POST /v2/seloger/search
func (c *Client) SelogerSearch(ctx context.Context, params *SelogerSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["seloger_search"], params.params(), options)
}

// SephoraProduct — Sephora Product. Read one European Sephora product by id or product URL. Set region when you are not on the French shop. Credit cost: 2 credits per call.
//
// GET / POST /v2/sephora/product
func (c *Client) SephoraProduct(ctx context.Context, params *SephoraProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["sephora_product"], params.params(), options)
}

// SephoraSearch — Sephora Search. Search a European Sephora shop by keyword. Region defaults to France. Page through makeup and skincare cards. The US site is a different page and is not included. Credit cost: 2 credits per call.
//
// GET / POST /v2/sephora/search
func (c *Client) SephoraSearch(ctx context.Context, params *SephoraSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["sephora_search"], params.params(), options)
}

// ShophouzzProduct — Shophouzz Product. Retrieve full product details from Shophouzz, including title, price, variants, images, vendor, and structured metadata. Credit cost: 1 credit per call.
//
// GET / POST /v2/shophouzz/product
func (c *Client) ShophouzzProduct(ctx context.Context, params *ShophouzzProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["shophouzz_product"], params.params(), options)
}

// ShopifyApps — Shopify Apps. Retrieve structured details about any Shopify App Store listing, including title, logo, screenshots, pricing plans, categories, and description. Credit cost: 1 credit per call.
//
// GET / POST /v2/shopify/apps
func (c *Client) ShopifyApps(ctx context.Context, params *ShopifyAppsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["shopify_apps"], params.params(), options)
}

// ShopifyCatalog — Shopify Catalog. List products from a Shopify shop. query is the shop domain or a URL on that shop. page and limit choose the slice. Credit cost: 1 credit per call.
//
// GET / POST /v2/shopify/catalog
func (c *Client) ShopifyCatalog(ctx context.Context, params *ShopifyCatalogParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["shopify_catalog"], params.params(), options)
}

// ShopifyCollections — Shopify Collections. List collections from a Shopify shop. query is the shop domain or a URL on that shop. page and limit choose the slice. Credit cost: 1 credit per call.
//
// GET / POST /v2/shopify/collections
func (c *Client) ShopifyCollections(ctx context.Context, params *ShopifyCollectionsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["shopify_collections"], params.params(), options)
}

// ShopifyProduct — Shopify Product. Extract full product details from any Shopify store product page, including title, variants, pricing, images, inventory, and options. Credit cost: 1 credit per call.
//
// GET / POST /v2/shopify/product
func (c *Client) ShopifyProduct(ctx context.Context, params *ShopifyProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["shopify_product"], params.params(), options)
}

// SimilarwebDomain — SimilarWeb Domain. Retrieve website traffic and engagement metrics for a domain from the public Similarweb page, including visits, bounce rate, rank, keywords, and traffic sources. Credit cost: 2 credits per call.
//
// GET / POST /v2/similarweb/domain
func (c *Client) SimilarwebDomain(ctx context.Context, params *SimilarwebDomainParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["similarweb_domain"], params.params(), options)
}

// SimilarwebSearch — Similarweb Search. Search the public Similarweb page for websites, companies, and mobile apps related to a keyword. Credit cost: 2 credits per call.
//
// GET / POST /v2/similarweb/search
func (c *Client) SimilarwebSearch(ctx context.Context, params *SimilarwebSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["similarweb_search"], params.params(), options)
}

// SnapedaProduct — SnapEDA Product. SnapEDA Product API: fetch full part page data from a SnapEDA URL or MPN/manufacturer slug. Returns CAD models, datasheet, RoHS, distributor pricing and availability via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/snapeda/product
func (c *Client) SnapedaProduct(ctx context.Context, params *SnapedaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["snapeda_product"], params.params(), options)
}

// SnapedaSearch — SnapEDA Search. SnapEDA Search API: search electronic parts by keyword or URL on SnapMagic Search (ex-SnapEDA). Returns MPN, manufacturer, pricing, CAD availability and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/snapeda/search
func (c *Client) SnapedaSearch(ctx context.Context, params *SnapedaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["snapeda_search"], params.params(), options)
}

// StarbucksMenu — Starbucks Menu. Full Starbucks menu for one café, with the default-size price of each item. Prices use browser rendering. Credit cost: 3 credits per call.
//
// GET / POST /v2/starbucks/menu
func (c *Client) StarbucksMenu(ctx context.Context, params *StarbucksMenuParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["starbucks_menu"], params.params(), options)
}

// StarbucksProduct — Starbucks Product. One Starbucks item at one café, with a price for every size. Get the café from Search first. Prices use browser rendering. Credit cost: 3 credits per call.
//
// GET / POST /v2/starbucks/product
func (c *Client) StarbucksProduct(ctx context.Context, params *StarbucksProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["starbucks_product"], params.params(), options)
}

// StarbucksSearch — Starbucks Search. Nearest Starbucks cafés from coordinates or a store-locator URL. Returns address, hours, amenities, and the store number used by Menu and Product. Credit cost: 3 credits per call.
//
// GET / POST /v2/starbucks/search
func (c *Client) StarbucksSearch(ctx context.Context, params *StarbucksSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["starbucks_search"], params.params(), options)
}

// Stats — Stats. Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 means contact support to activate it.
//
// GET /v2/stats
func (c *Client) Stats(ctx context.Context, params *StatsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stats"], params.params(), options)
}

// StockxProduct — StockX Product. Retrieve detailed product information from StockX, including market prices, bid/ask data, size variants, and product metadata. Credit cost: 2 credits per call.
//
// GET / POST /v2/stockx/product
func (c *Client) StockxProduct(ctx context.Context, params *StockxProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stockx_product"], params.params(), options)
}

// StockxSearch — StockX Search. Search StockX for products by name, brand, SKU, or slug and receive a list of matching items with product IDs, slugs, images, and colorways. Credit cost: 2 credits per call.
//
// GET / POST /v2/stockx/search
func (c *Client) StockxSearch(ctx context.Context, params *StockxSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stockx_search"], params.params(), options)
}

// StockxTrends — StockX Trends. Retrieve trending products on StockX by category, with options to sort by featured, most active, or release date. Credit cost: 2 credits per call.
//
// GET / POST /v2/stockx/trends
func (c *Client) StockxTrends(ctx context.Context, params *StockxTrendsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stockx_trends"], params.params(), options)
}

// StubhubEvent — StubHub Event. Read one StubHub event from its URL or numeric id, including the lowest printed ticket price. Credit cost: 2 credits per call.
//
// GET / POST /v2/stubhub/event
func (c *Client) StubhubEvent(ctx context.Context, params *StubhubEventParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stubhub_event"], params.params(), options)
}

// StubhubSearch — StubHub Search. Search StubHub events by keyword or by a StubHub search URL. The response is the first embedded page. Credit cost: 1 credit per call.
//
// GET / POST /v2/stubhub/search
func (c *Client) StubhubSearch(ctx context.Context, params *StubhubSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stubhub_search"], params.params(), options)
}

// TargetProduct — Target Product. Get a Target.com product page: price, description, images and variations via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/target/product
func (c *Client) TargetProduct(ctx context.Context, params *TargetProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["target_product"], params.params(), options)
}

// TargetSearch — Target Search. Search Target.com by keyword or search URL. Returns product cards with price, rating and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/target/search
func (c *Client) TargetSearch(ctx context.Context, params *TargetSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["target_search"], params.params(), options)
}

// TdlrLicense — TDLR License. Fetch one TDLR license by detail id or SearchResultDetail URL, including status, subtype, DBA, and notes. Credit cost: 1 credit per call.
//
// GET / POST /v2/tdlr/license
func (c *Client) TdlrLicense(ctx context.Context, params *TdlrLicenseParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["tdlr_license"], params.params(), options)
}

// TdlrSearch — TDLR License Search. Search Texas TDLR licenses by name, city, county, or ZIP. license_type such as AIRREF is required for a useful search. 25 rows per page. Credit cost: 1 credit per call.
//
// GET / POST /v2/tdlr/search
func (c *Client) TdlrSearch(ctx context.Context, params *TdlrSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["tdlr_search"], params.params(), options)
}

// TedNotice — TED Notice. TED notice profile: buyer, lots, deadline, value and documents (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/ted/notice
func (c *Client) TedNotice(ctx context.Context, params *TedNoticeParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ted_notice"], params.params(), options)
}

// TedSearch — TED Search. Search TED procurement notices by keyword or filters. Returns notice cards and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/ted/search
func (c *Client) TedSearch(ctx context.Context, params *TedSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ted_search"], params.params(), options)
}

// ThaliaProduct — Thalia Product. Read one Thalia product from its URL. You get price, author or brand, and the product details. Credit cost: 2 credits per call.
//
// GET / POST /v2/thalia/product
func (c *Client) ThaliaProduct(ctx context.Context, params *ThaliaProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["thalia_product"], params.params(), options)
}

// ThaliaSearch — Thalia Search. Search Thalia by keyword and page books, media, and gifts. Cards include title, price, and the product link. Credit cost: 2 credits per call.
//
// GET / POST /v2/thalia/search
func (c *Client) ThaliaSearch(ctx context.Context, params *ThaliaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["thalia_search"], params.params(), options)
}

// TheforkRestaurant — TheFork Restaurant. Extract a full TheFork restaurant profile from a restaurant page URL, including ratings, tags, photo gallery, opening hours, and AI review summaries via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/thefork/restaurant
func (c *Client) TheforkRestaurant(ctx context.Context, params *TheforkRestaurantParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["thefork_restaurant"], params.params(), options)
}

// TheforkSearch — TheFork Search. Search TheFork restaurants from a full search or city URL and retrieve ratings, cuisine, prices, availability, review summaries, and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/thefork/search
func (c *Client) TheforkSearch(ctx context.Context, params *TheforkSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["thefork_search"], params.params(), options)
}

// ThomasnetCompany — Thomasnet Company. Get a full Thomasnet supplier profile from a company URL, including contact details, product categories, and business information. Credit cost: 3 credits per call.
//
// GET / POST /v2/thomasnet/company
func (c *Client) ThomasnetCompany(ctx context.Context, params *ThomasnetCompanyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["thomasnet_company"], params.params(), options)
}

// ThomasnetSearch — Thomasnet Search. Search Thomasnet suppliers from a search or category URL and get company listings, filters, name suggestions, and pagination. Credit cost: 3 credits per call.
//
// GET / POST /v2/thomasnet/search
func (c *Client) ThomasnetSearch(ctx context.Context, params *ThomasnetSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["thomasnet_search"], params.params(), options)
}

// TicketmasterEvent — Ticketmaster Event. Fetch one Ticketmaster event from a slugged event URL. The slug is searched and the matching event id is returned with the venue address. Credit cost: 1 credit per call.
//
// GET / POST /v2/ticketmaster/event
func (c *Client) TicketmasterEvent(ctx context.Context, params *TicketmasterEventParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ticketmaster_event"], params.params(), options)
}

// TicketmasterSearch — Ticketmaster Search. Search Ticketmaster events by keyword. Returns event cards, including Ticketweb partner rows, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/ticketmaster/search
func (c *Client) TicketmasterSearch(ctx context.Context, params *TicketmasterSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["ticketmaster_search"], params.params(), options)
}

// TiktoktopadsAd — TikTok Top Ads Ad. Get one TikTok Creative Center Top Ad by URL or id: creative, landing page, CTR, and engagement (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/tiktoktopads/ad
func (c *Client) TiktoktopadsAd(ctx context.Context, params *TiktoktopadsAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["tiktoktopads_ad"], params.params(), options)
}

// TiktoktopadsSearch — TikTok Top Ads Search. Search TikTok Creative Center Top Ads by region, objective, period, and keyword. Returns ad cards and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/tiktoktopads/search
func (c *Client) TiktoktopadsSearch(ctx context.Context, params *TiktoktopadsSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["tiktoktopads_search"], params.params(), options)
}

// TradeaplaneListing — Trade-A-Plane Listing. Read one Trade-A-Plane aircraft listing from its URL. You get the asking price, the airframe details, and the seller. Credit cost: 2 credits per call.
//
// GET / POST /v2/tradeaplane/listing
func (c *Client) TradeaplaneListing(ctx context.Context, params *TradeaplaneListingParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["tradeaplane_listing"], params.params(), options)
}

// TradeaplaneSearch — Trade-A-Plane Search. Search Trade-A-Plane aircraft listings by keyword and page the results. Cards include title, price, and the listing link. Credit cost: 2 credits per call.
//
// GET / POST /v2/tradeaplane/search
func (c *Client) TradeaplaneSearch(ctx context.Context, params *TradeaplaneSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["tradeaplane_search"], params.params(), options)
}

// TrustpilotCompanyInfo — Trustpilot Company Info. Retrieve Trustpilot company profile data, including trust score, star rating, categories, contact info, and review response behavior. Credit cost: 1 credit per call.
//
// GET / POST /v2/trustpilot/company/info
func (c *Client) TrustpilotCompanyInfo(ctx context.Context, params *TrustpilotCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["trustpilot_company_info"], params.params(), options)
}

// UntappdBeer — Untappd Beer. Read one Untappd beer, including ABV, IBU, and brewery. Credit cost: 2 credits per call.
//
// GET / POST /v2/untappd/beer
func (c *Client) UntappdBeer(ctx context.Context, params *UntappdBeerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["untappd_beer"], params.params(), options)
}

// UntappdSearch — Untappd Search. Search Untappd beers by keyword. Cards include brewery, ABV, and IBU. next is null when Untappd does not expose a next page. Credit cost: 2 credits per call.
//
// GET / POST /v2/untappd/search
func (c *Client) UntappdSearch(ctx context.Context, params *UntappdSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["untappd_search"], params.params(), options)
}

// UpworkFreelancer — Upwork Freelancer. Upwork Freelancer API: fetch public freelancer profile data from a profile ID, username or URL. Returns bio, hourly rate, skills, ratings and agency info via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/upwork/freelancer
func (c *Client) UpworkFreelancer(ctx context.Context, params *UpworkFreelancerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["upwork_freelancer"], params.params(), options)
}

// UpworkJobs — Upwork Jobs. Upwork Jobs API: search freelance job listings by skill slug or URL. Returns budgets, skills, experience level and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/upwork/jobs
func (c *Client) UpworkJobs(ctx context.Context, params *UpworkJobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["upwork_jobs"], params.params(), options)
}

// UpworkServices — Upwork Services. Upwork Services API: browse Project Catalog listings by category slug or URL. Returns prices, product IDs and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/upwork/services
func (c *Client) UpworkServices(ctx context.Context, params *UpworkServicesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["upwork_services"], params.params(), options)
}

// UpworkServicesInfo — Upwork Services Info. Upwork Services Info API: fetch Project Catalog product details from a product URL. Returns packages, pricing, seller, media and breadcrumbs via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/upwork/services/info
func (c *Client) UpworkServicesInfo(ctx context.Context, params *UpworkServicesInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["upwork_services_info"], params.params(), options)
}

// Usage — Usage. Retrieve real-time usage statistics and credit balance for your Piloterr account.
//
// GET /v2/usage
func (c *Client) Usage(ctx context.Context, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usage"], nil, options)
}

// UsaspendingAward — USAspending Award. Fetch one US federal award by its generated id or by a usaspending.gov award URL. Amounts are a number plus a display string. Credit cost: 1 credit per call.
//
// GET / POST /v2/usaspending/award
func (c *Client) UsaspendingAward(ctx context.Context, params *UsaspendingAwardParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usaspending_award"], params.params(), options)
}

// UsaspendingSearch — USAspending Search. Search US federal awards by keyword. award_type selects one group: contract (default), grant, idv, or loan. Amounts are a number plus a display string. Credit cost: 1 credit per call.
//
// GET / POST /v2/usaspending/search
func (c *Client) UsaspendingSearch(ctx context.Context, params *UsaspendingSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usaspending_search"], params.params(), options)
}

// VestiairecollectiveProduct — Vestiaire Collective Product. Fetch one Vestiaire Collective listing by product URL. The response is a flat object with price, size, condition, and seller. Credit cost: 2 credits per call.
//
// GET / POST /v2/vestiairecollective/product
func (c *Client) VestiairecollectiveProduct(ctx context.Context, params *VestiairecollectiveProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vestiairecollective_product"], params.params(), options)
}

// VestiairecollectiveSearch — Vestiaire Collective Search. Search Vestiaire Collective by keyword or catalog URL. Each card has brand, name, price, size, and a product link. Credit cost: 2 credits per call.
//
// GET / POST /v2/vestiairecollective/search
func (c *Client) VestiairecollectiveSearch(ctx context.Context, params *VestiairecollectiveSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vestiairecollective_search"], params.params(), options)
}

// VintedItem — Vinted Item. Fetch a Vinted listing by item URL or numeric ID. Parses the public item page (Next.js payload), returns a normalized snake_case object with photos, attributes, shipping, pricing, and optional seller profile. Credit cost: 1 credit per call.
//
// GET / POST /v2/vinted/item
func (c *Client) VintedItem(ctx context.Context, params *VintedItemParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vinted_item"], params.params(), options)
}

// VintedSearch — Vinted Search. Search Vinted second-hand listings by keyword or catalog URL across vinted.fr, vinted.com and other domains. Returns results and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v2/vinted/search
func (c *Client) VintedSearch(ctx context.Context, params *VintedSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vinted_search"], params.params(), options)
}

// VintedUserInfo — Vinted User Info. Retrieve public Vinted seller profile by member URL or user ID. Returns the raw /api/v2/users/{id} object: login, photo, feedback stats, followers, verification, bundle discounts, and listing counts. Credit cost: 1 credit per call.
//
// GET / POST /v2/vinted/user/info
func (c *Client) VintedUserInfo(ctx context.Context, params *VintedUserInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vinted_user_info"], params.params(), options)
}

// VivarealProperty — Viva Real Property. Read one Viva Real property from its URL. The price follows the operation of that URL, so a sale link stays a sale price. Credit cost: 1 credit per call.
//
// GET / POST /v2/vivareal/property
func (c *Client) VivarealProperty(ctx context.Context, params *VivarealPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vivareal_property"], params.params(), options)
}

// VivarealSearch — Viva Real Search. Search Viva Real by city and operation, sale or rent. Page with pagina. Development cards and 'from' prices are left out. Credit cost: 1 credit per call.
//
// GET / POST /v2/vivareal/search
func (c *Client) VivarealSearch(ctx context.Context, params *VivarealSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["vivareal_search"], params.params(), options)
}

// WalmartProduct — Walmart Product. Extract full Walmart product page (PDP) data from a /ip/ URL: pricing, ratings, images, fulfillment, specs, badges, and availability via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/walmart/product
func (c *Client) WalmartProduct(ctx context.Context, params *WalmartProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["walmart_product"], params.params(), options)
}

// WalmartSearch — Walmart Search. Scrape Walmart.com search result pages from a full URL and retrieve product cards with pricing, ratings, availability, and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/walmart/search
func (c *Client) WalmartSearch(ctx context.Context, params *WalmartSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["walmart_search"], params.params(), options)
}

// WalmartSuggest — Walmart Suggest. Get Walmart.com search autocomplete suggestions for a keyword via browser rendering, with optional US ZIP code localization. Credit cost: 2 credits per call.
//
// GET / POST /v2/walmart/suggest
func (c *Client) WalmartSuggest(ctx context.Context, params *WalmartSuggestParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["walmart_suggest"], params.params(), options)
}

// WebsiteAntibot — Website Antibot. Detect which anti-bot protection protects a website before you scrape. Get the vendor, confidence level, and matching clues from a single URL check. Credit cost: 1 credit per call.
//
// GET /v2/website/antibot
func (c *Client) WebsiteAntibot(ctx context.Context, params *WebsiteAntibotParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_antibot"], params.params(), options)
}

// WebsiteCrawler — Website Crawler. Fetch public pages with fast HTTP crawling (no JavaScript). Best first choice for static or server-rendered HTML; use Rendering or WebUnlocker when JS or anti-bot blocks you. Credit cost: 1 credit per call.
//
// GET / POST /v2/website/crawler
func (c *Client) WebsiteCrawler(ctx context.Context, params *WebsiteCrawlerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_crawler"], params.params(), options)
}

// WebsiteEmailPhoneExtractor — Website Email Phone Extractor. Extract emails and phone numbers from websites for comprehensive contact details, including social media profiles across 12+ platforms. Credit cost: 1 credit per call.
//
// GET / POST /v2/website/email_phone_extractor
func (c *Client) WebsiteEmailPhoneExtractor(ctx context.Context, params *WebsiteEmailPhoneExtractorParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_email_phone_extractor"], params.params(), options)
}

// WebsiteRendering — Website Rendering. Render JavaScript-heavy pages in a headless browser and return post-render HTML. Use when Crawler returns empty DOM; escalate to WebUnlocker if bot protection blocks the session. Credit cost: 2 credits per call.
//
// GET / POST /v2/website/rendering
func (c *Client) WebsiteRendering(ctx context.Context, params *WebsiteRenderingParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_rendering"], params.params(), options)
}

// WebsiteScreenshot — Website Screenshot. Browser-like Screenshot API that captures webpages as PNG/JPEG/WebP or PDF for previews, QA, reporting, monitoring, and archiving no headless setup. Credit cost: 2 credits per call.
//
// GET / POST /v2/website/screenshot
func (c *Client) WebsiteScreenshot(ctx context.Context, params *WebsiteScreenshotParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_screenshot"], params.params(), options)
}

// WebsiteTechnology — Website Technology. Identify the technologies behind any website CMS, frameworks, analytics, CDN, hosting, and more, for competitive market analysis and technological insight. Credit cost: 1 credit per call.
//
// GET / POST /v2/website/technology
func (c *Client) WebsiteTechnology(ctx context.Context, params *WebsiteTechnologyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_technology"], params.params(), options)
}

// WebsiteWebunlocker — Website WebUnlocker. Bypass advanced anti-bot systems (Cloudflare, DataDome, Akamai, PerimeterX) on whitelisted domains. Combines browser rendering and unlock tooling; use after Crawler and Rendering fail. Credit cost: 3 credits per call.
//
// GET / POST /v2/website/webunlocker
func (c *Client) WebsiteWebunlocker(ctx context.Context, params *WebsiteWebunlockerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_webunlocker"], params.params(), options)
}

// WeiboPostInfo — Weibo Post Info. Retrieve one public Weibo post by post id, including text, images, publish time, engagement counts, and author profile. Credit cost: 2 credits per call.
//
// GET / POST /v2/weibo/post/info
func (c *Client) WeiboPostInfo(ctx context.Context, params *WeiboPostInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["weibo_post_info"], params.params(), options)
}

// WeiboSearch — Weibo Search. Search public Weibo posts by keywords. Each post includes the author, text, publish time, engagement counts, image, and video link. Credit cost: 2 credits per call.
//
// GET / POST /v2/weibo/search
func (c *Client) WeiboSearch(ctx context.Context, params *WeiboSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["weibo_search"], params.params(), options)
}

// WelcometothejungleCompany — Welcome to the Jungle Company. Welcome to the Jungle company profile: headcount, offices, sectors and website (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/welcometothejungle/company
func (c *Client) WelcometothejungleCompany(ctx context.Context, params *WelcometothejungleCompanyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["welcometothejungle_company"], params.params(), options)
}

// WelcometothejungleSearch — Welcome to the Jungle Search. Search Welcome to the Jungle jobs. Returns offer cards and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/welcometothejungle/search
func (c *Client) WelcometothejungleSearch(ctx context.Context, params *WelcometothejungleSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["welcometothejungle_search"], params.params(), options)
}

// WellfoundCompanyInfo — Wellfound Company Info. Retrieve startup and company data from Wellfound (formerly AngelList), including description, funding, team size, hiring status, and social links. Credit cost: 1 credit per call.
//
// GET / POST /v2/wellfound/company/info
func (c *Client) WellfoundCompanyInfo(ctx context.Context, params *WellfoundCompanyInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["wellfound_company_info"], params.params(), options)
}

// XPostInfo — X Post Info (twitter). X Post Info API: fetch public tweet data (text, likes, retweets, views, media) from a status URL or tweet ID via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/x/post/info
func (c *Client) XPostInfo(ctx context.Context, params *XPostInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["x_post_info"], params.params(), options)
}

// XProfileInfo — X Profile Info (twitter). X Profile Info API: fetch public X/Twitter profile data (bio, followers, verification, avatars) from a username or URL via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/x/profile/info
func (c *Client) XProfileInfo(ctx context.Context, params *XProfileInfoParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["x_profile_info"], params.params(), options)
}

// YahooSearch — Yahoo Search. Query Yahoo via API and retrieve structured organic results with country and language. Credit cost: 1 credit per call.
//
// GET / POST /v2/yahoo/search
func (c *Client) YahooSearch(ctx context.Context, params *YahooSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yahoo_search"], params.params(), options)
}

// YellowpagesBusiness — Yellowpages Business. Extract a full Yellow Pages business profile from a listing URL, including phone, website, categories, and ratings via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/yellowpages/business
func (c *Client) YellowpagesBusiness(ctx context.Context, params *YellowpagesBusinessParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yellowpages_business"], params.params(), options)
}

// YellowpagesSearch — Yellowpages Search. Search Yellow Pages businesses from a full search URL and retrieve names, phones, addresses, ratings, opening hours, and pagination via browser rendering. Credit cost: 2 credits per call.
//
// GET / POST /v2/yellowpages/search
func (c *Client) YellowpagesSearch(ctx context.Context, params *YellowpagesSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yellowpages_search"], params.params(), options)
}

// YelpBusiness — Yelp Business. Fetch a Yelp business profile from a full business URL with phone, address, website, logo, hours, and claim status. Credit cost: 2 credits per call.
//
// GET / POST /v2/yelp/business
func (c *Client) YelpBusiness(ctx context.Context, params *YelpBusinessParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yelp_business"], params.params(), options)
}

// YelpReviews — Yelp Reviews. Fetch paginated Yelp reviews for a business with rating, text, author profile, reactions, and attached photos or videos. Credit cost: 2 credits per call.
//
// GET / POST /v2/yelp/reviews
func (c *Client) YelpReviews(ctx context.Context, params *YelpReviewsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yelp_reviews"], params.params(), options)
}

// YelpSearch — Yelp Search. Search Yelp by keyword and location and retrieve paginated business cards with ratings, categories, photos, and sponsored flags. Credit cost: 2 credits per call.
//
// GET / POST /v2/yelp/search
func (c *Client) YelpSearch(ctx context.Context, params *YelpSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yelp_search"], params.params(), options)
}

// ZapimoveisAd — Zap Imóveis Ad. Get a full Zap Imóveis listing by URL: price, IPTU, features, coordinates, images, and publisher (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/zapimoveis/ad
func (c *Client) ZapimoveisAd(ctx context.Context, params *ZapimoveisAdParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zapimoveis_ad"], params.params(), options)
}

// ZapimoveisSearch — Zap Imóveis Search. Search Zap Imóveis listings by city or search URL. Returns prices, IPTU, condominium fees, rooms, and pagination (1 credit). Credit cost: 1 credit per call.
//
// GET / POST /v2/zapimoveis/search
func (c *Client) ZapimoveisSearch(ctx context.Context, params *ZapimoveisSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zapimoveis_search"], params.params(), options)
}

// ZazzleProduct — Zazzle Product. Fetch one Zazzle product by its product URL or by a slug that ends in the numeric product id. Credit cost: 2 credits per call.
//
// GET / POST /v2/zazzle/product
func (c *Client) ZazzleProduct(ctx context.Context, params *ZazzleProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zazzle_product"], params.params(), options)
}

// ZazzleSearch — Zazzle Search. Search Zazzle products by keyword. page maps to the pg parameter. Prices are a number plus a display string. Credit cost: 1 credit per call.
//
// GET / POST /v2/zazzle/search
func (c *Client) ZazzleSearch(ctx context.Context, params *ZazzleSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zazzle_search"], params.params(), options)
}

// ZillowProperty — Zillow Property. Extract full property details from a Zillow homedetails URL, including specs, photos, schools, tax history, and agent attribution. Credit cost: 3 credits per call.
//
// GET / POST /v2/zillow/property
func (c *Client) ZillowProperty(ctx context.Context, params *ZillowPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zillow_property"], params.params(), options)
}

// ZillowSearch — Zillow Search. Search Zillow listings from a search or map URL and retrieve prices, beds, baths, coordinates, and listing status. Credit cost: 3 credits per call.
//
// GET / POST /v2/zillow/search
func (c *Client) ZillowSearch(ctx context.Context, params *ZillowSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zillow_search"], params.params(), options)
}

// ZonapropProperty — ZonaProp Property. Read one ZonaProp listing from its URL. You get the asking price in USD or ARS, the address, and the publisher. Credit cost: 2 credits per call.
//
// GET / POST /v2/zonaprop/property
func (c *Client) ZonapropProperty(ctx context.Context, params *ZonapropPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zonaprop_property"], params.params(), options)
}

// ZonapropSearch — ZonaProp Search. Search ZonaProp for apartments or houses in an Argentine city. Page with pagina. A few featured ads can repeat. Development projects priced 'from' are left out. Credit cost: 2 credits per call.
//
// GET / POST /v2/zonaprop/search
func (c *Client) ZonapropSearch(ctx context.Context, params *ZonapropSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zonaprop_search"], params.params(), options)
}

// ZooplaProperty — Zoopla Property. Extract full property details from a Zoopla listing URL, including price, specs, photo gallery, EPC rating, coordinates, and agent data. Credit cost: 2 credits per call.
//
// GET / POST /v2/zoopla/property
func (c *Client) ZooplaProperty(ctx context.Context, params *ZooplaPropertyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zoopla_property"], params.params(), options)
}

// ZooplaSearch — Zoopla Search. Search Zoopla listings from a search URL and retrieve prices, specs, photos, agency info, and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/zoopla/search
func (c *Client) ZooplaSearch(ctx context.Context, params *ZooplaSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zoopla_search"], params.params(), options)
}

// ZooplusProduct — Zooplus Product. Read one Zooplus product by id or URL, in the region of that shop. Variants keep their own price and EAN. Credit cost: 1 credit per call.
//
// GET / POST /v2/zooplus/product
func (c *Client) ZooplusProduct(ctx context.Context, params *ZooplusProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zooplus_product"], params.params(), options)
}

// ZooplusSearch — Zooplus Search. Search a Zooplus country shop by keyword. Region defaults to France. Page through food, litter, and accessory cards with local prices. Credit cost: 1 credit per call.
//
// GET / POST /v2/zooplus/search
func (c *Client) ZooplusSearch(ctx context.Context, params *ZooplusSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zooplus_search"], params.params(), options)
}

// ZoroProduct — Zoro Product. Get a Zoro product page: title, brand, USD price, Zoro number, manufacturer number, images, and specifications. Credit cost: 2 credits per call.
//
// GET / POST /v2/zoro/product
func (c *Client) ZoroProduct(ctx context.Context, params *ZoroProductParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zoro_product"], params.params(), options)
}

// ZoroSearch — Zoro Search. Search Zoro by keyword or category URL. Returns product cards with USD price, brand, Zoro number, and pagination. Credit cost: 2 credits per call.
//
// GET / POST /v2/zoro/search
func (c *Client) ZoroSearch(ctx context.Context, params *ZoroSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["zoro_search"], params.params(), options)
}

// endpoints lists every documented operation by its snake_case name.
var endpoints = map[string]endpoint{
	"aboutyou_product":              {path: "/v2/aboutyou/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aboutyou_search":               {path: "/v2/aboutyou/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"action_product":                {path: "/v2/action/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"action_search":                 {path: "/v2/action/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aldi_product":                  {path: "/v2/aldi/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aldi_search":                   {path: "/v2/aldi/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"alibaba_product":               {path: "/v2/alibaba/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"alibaba_search":                {path: "/v2/alibaba/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aliexpress_product":            {path: "/v2/aliexpress/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aliexpress_product_reviews":    {path: "/v2/aliexpress/product/reviews", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aliexpress_search":             {path: "/v2/aliexpress/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"aliexpress_store_product":      {path: "/v2/aliexpress/store/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"allegro_product":               {path: "/v2/allegro/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"allegro_search":                {path: "/v2/allegro/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"amazon_product":                {path: "/v2/amazon/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"amazon_product_offer":          {path: "/v2/amazon/product/offer", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"amazon_search":                 {path: "/v2/amazon/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"apartments_property":           {path: "/v2/apartments/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"apartments_search":             {path: "/v2/apartments/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"argenprop_property":            {path: "/v2/argenprop/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"argenprop_search":              {path: "/v2/argenprop/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"argos_product":                 {path: "/v2/argos/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"argos_search":                  {path: "/v2/argos/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"arxiv_paper":                   {path: "/v2/arxiv/paper", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"arxiv_search":                  {path: "/v2/arxiv/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"auchan_product":                {path: "/v2/auchan/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"auchan_search":                 {path: "/v2/auchan/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"autodoc_product":               {path: "/v2/autodoc/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"autodoc_product_related":       {path: "/v2/autodoc/product/related", methods: []string{"GET", "POST"}, json: true, required: []string{"query", "maker_id", "vehicle_type", "maker_name", "product_id"}},
	"autodoc_search":                {path: "/v2/autodoc/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"autoeurope_result":             {path: "/v2/autoeurope/result", methods: []string{"GET", "POST"}, json: true, required: []string{"offer_id", "pickup", "pickup_date", "dropoff_date"}},
	"autoeurope_search":             {path: "/v2/autoeurope/search", methods: []string{"GET", "POST"}, json: true, required: []string{"pickup", "pickup_date", "dropoff_date"}},
	"autoscout24_ad":                {path: "/v2/autoscout24/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"autoscout24_search":            {path: "/v2/autoscout24/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"autotrader_ad":                 {path: "/v2/autotrader/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"autotrader_search":             {path: "/v2/autotrader/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bestbuy_product":               {path: "/v2/bestbuy/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bestbuy_product_reviews":       {path: "/v2/bestbuy/product/reviews", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bestbuy_search":                {path: "/v2/bestbuy/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bhphotovideo_product":          {path: "/v2/bhphotovideo/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bhphotovideo_search":           {path: "/v2/bhphotovideo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bing_search":                   {path: "/v2/bing/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bitiba_product":                {path: "/v2/bitiba/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bitiba_search":                 {path: "/v2/bitiba/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bol_product":                   {path: "/v2/bol/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bol_search":                    {path: "/v2/bol/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"boulanger_product":             {path: "/v2/boulanger/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"boulanger_search":              {path: "/v2/boulanger/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"brave_search":                  {path: "/v2/brave/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bricodepot_product":            {path: "/v2/bricodepot/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bricodepot_search":             {path: "/v2/bricodepot/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bricorama_product":             {path: "/v2/bricorama/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"bricorama_search":              {path: "/v2/bricorama/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"carrefour_product":             {path: "/v2/carrefour/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"carrefour_search":              {path: "/v2/carrefour/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"carrefour_suggest":             {path: "/v2/carrefour/suggest", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"castorama_product":             {path: "/v2/castorama/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"castorama_search":              {path: "/v2/castorama/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"cdiscount_product":             {path: "/v2/cdiscount/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"cdiscount_product_offer":       {path: "/v2/cdiscount/product/offer", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"cdiscount_search":              {path: "/v2/cdiscount/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"centris_property":              {path: "/v2/centris/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"centris_search":                {path: "/v2/centris/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"chewy_product":                 {path: "/v2/chewy/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"chewy_search":                  {path: "/v2/chewy/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"clearedjobs_job":               {path: "/v2/clearedjobs/job", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"clearedjobs_search":            {path: "/v2/clearedjobs/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"company":                       {path: "/v2/company", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"coolblue_product":              {path: "/v2/coolblue/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"coolblue_search":               {path: "/v2/coolblue/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"copart_lot":                    {path: "/v2/copart/lot", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"copart_search":                 {path: "/v2/copart/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"costco_product":                {path: "/v2/costco/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"costco_search":                 {path: "/v2/costco/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"crunchbase_company_info":       {path: "/v2/crunchbase/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"crunchbase_event":              {path: "/v2/crunchbase/event", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"crunchbase_funding_round":      {path: "/v2/crunchbase/funding_round", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"crunchbase_funding_rounds":     {path: "/v2/crunchbase/funding_rounds", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"crunchbase_people_info":        {path: "/v2/crunchbase/people/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"crunchbase_search":             {path: "/v2/crunchbase/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"cultura_product":               {path: "/v2/cultura/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"cultura_search":                {path: "/v2/cultura/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"darty_product":                 {path: "/v2/darty/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"darty_search":                  {path: "/v2/darty/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"datasets_linkedin_company":     {path: "/v2/datasets/linkedin/company", methods: []string{"POST"}, json: true, required: []string{}},
	"decathlon_product":             {path: "/v2/decathlon/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"decathlon_search":              {path: "/v2/decathlon/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"digikey_product":               {path: "/v2/digikey/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"digikey_search":                {path: "/v2/digikey/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"doaj_article":                  {path: "/v2/doaj/article", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"doaj_journal":                  {path: "/v2/doaj/journal", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"doaj_search":                   {path: "/v2/doaj/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"domain_dnsbl":                  {path: "/v2/domain/dnsbl", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"domain_malicious":              {path: "/v2/domain/malicious", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"domain_whois":                  {path: "/v2/domain/whois", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"domaincomau_property":          {path: "/v2/domaincomau/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"domaincomau_search":            {path: "/v2/domaincomau/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"duckduckgo_search":             {path: "/v2/duckduckgo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ebay_product":                  {path: "/v2/ebay/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ebay_search":                   {path: "/v2/ebay/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"elcorteingles_product":         {path: "/v2/elcorteingles/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"elcorteingles_search":          {path: "/v2/elcorteingles/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"eleclerc_product":              {path: "/v2/eleclerc/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"eleclerc_search":               {path: "/v2/eleclerc/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"eleclerc_suggest":              {path: "/v2/eleclerc/suggest", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"email_analyzes":                {path: "/v2/email/analyzes", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"email_finder":                  {path: "/v2/email/finder", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"email_verify":                  {path: "/v2/email/verify", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"europages_company":             {path: "/v2/europages/company", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"europages_search":              {path: "/v2/europages/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"eventbrite_event":              {path: "/v2/eventbrite/event", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"eventbrite_search":             {path: "/v2/eventbrite/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"expedia_ad":                    {path: "/v2/expedia/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"expedia_flight":                {path: "/v2/expedia/flight", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"expedia_search":                {path: "/v2/expedia/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fashionphile_product":          {path: "/v2/fashionphile/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fashionphile_search":           {path: "/v2/fashionphile/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_chewy_products":         {path: "/v2/finder/chewy-products", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"finder_eleclerc_stores":        {path: "/v2/finder/eleclerc-stores", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_google_countries":       {path: "/v2/finder/google-countries", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_google_languages":       {path: "/v2/finder/google-languages", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_google_locations":       {path: "/v2/finder/google-locations", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_leroymerlin_products":   {path: "/v2/finder/leroymerlin-products", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"finder_leroymerlin_stores":     {path: "/v2/finder/leroymerlin-stores", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_linkedin_industries":    {path: "/v2/finder/linkedin-industries", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_mrbricolage_stores":     {path: "/v2/finder/mrbricolage-stores", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fiverr_profile_info":           {path: "/v2/fiverr/profile/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fiverr_search":                 {path: "/v2/fiverr/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fnac_product":                  {path: "/v2/fnac/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fnac_search":                   {path: "/v2/fnac/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fotocasa_ad":                   {path: "/v2/fotocasa/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fotocasa_search":               {path: "/v2/fotocasa/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"franchisedirect_franchise":     {path: "/v2/franchisedirect/franchise", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"franchisedirect_search":        {path: "/v2/franchisedirect/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fressnapf_product":             {path: "/v2/fressnapf/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"fressnapf_search":              {path: "/v2/fressnapf/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"funda_property":                {path: "/v2/funda/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"funda_search":                  {path: "/v2/funda/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"g2_product_info":               {path: "/v2/g2/product/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"github_user_info":              {path: "/v2/github/user/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"glassdoor_company_info":        {path: "/v2/glassdoor/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"glassdoor_company_search":      {path: "/v2/glassdoor/company/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"glassdoor_job_info":            {path: "/v2/glassdoor/job/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"glassdoor_job_search":          {path: "/v2/glassdoor/job/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"goodreads_book":                {path: "/v2/goodreads/book", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"goodreads_search":              {path: "/v2/goodreads/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_images":                 {path: "/v2/google/images", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_jobs":                   {path: "/v2/google/jobs", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_maps":                   {path: "/v2/google/maps", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_news":                   {path: "/v2/google/news", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search":                 {path: "/v2/google/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search_ai":              {path: "/v2/google/search/ai", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search_autocomplete":    {path: "/v2/google/search/autocomplete", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_shopping":               {path: "/v2/google/shopping", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_videos":                 {path: "/v2/google/videos", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"googleadstransparency_ad":      {path: "/v2/googleadstransparency/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"googleadstransparency_search":  {path: "/v2/googleadstransparency/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"grainger_product":              {path: "/v2/grainger/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"grainger_search":               {path: "/v2/grainger/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"greenhouse_job":                {path: "/v2/greenhouse/job", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"greenhouse_search":             {path: "/v2/greenhouse/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"gsaelibrary_contract":          {path: "/v2/gsaelibrary/contract", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"gsaelibrary_search":            {path: "/v2/gsaelibrary/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"homestra_ad":                   {path: "/v2/homestra/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"homestra_search":               {path: "/v2/homestra/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"huggingface_model":             {path: "/v2/huggingface/model", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"huggingface_search":            {path: "/v2/huggingface/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"idealista_property":            {path: "/v2/idealista/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"idealista_provinces":           {path: "/v2/idealista/provinces", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"idealista_search":              {path: "/v2/idealista/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"idealo_product":                {path: "/v2/idealo/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"idealo_search":                 {path: "/v2/idealo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ikea_product":                  {path: "/v2/ikea/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ikea_search":                   {path: "/v2/ikea/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"immobilienscout24_property":    {path: "/v2/immobilienscout24/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"immobilienscout24_search":      {path: "/v2/immobilienscout24/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"immoscout24ch_search":          {path: "/v2/immoscout24ch/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"immowelt_property":             {path: "/v2/immowelt/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"immowelt_search":               {path: "/v2/immowelt/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"imoveisweb_property":           {path: "/v2/imoveisweb/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"imoveisweb_search":             {path: "/v2/imoveisweb/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"indeed_company_info":           {path: "/v2/indeed/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"indeed_company_search":         {path: "/v2/indeed/company/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"indeed_job_info":               {path: "/v2/indeed/job/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"indeed_job_salary":             {path: "/v2/indeed/job/salary", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"indeed_job_search":             {path: "/v2/indeed/job/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"inet_product":                  {path: "/v2/inet/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"inet_search":                   {path: "/v2/inet/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"instagram_post_info":           {path: "/v2/instagram/post/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"instagram_user_info":           {path: "/v2/instagram/user/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"intermarche_product":           {path: "/v2/intermarche/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query", "shop_id"}},
	"intermarche_search":            {path: "/v2/intermarche/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query", "shop_id"}},
	"intersport_product":            {path: "/v2/intersport/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"intersport_search":             {path: "/v2/intersport/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"jamesedition_property":         {path: "/v2/jamesedition/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"jamesedition_search":           {path: "/v2/jamesedition/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"jobs":                          {path: "/v1/jobs/{id}", methods: []string{"GET"}, json: true, required: []string{"id"}},
	"kick_user_info":                {path: "/v2/kick/user/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"kijiji_ad":                     {path: "/v2/kijiji/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"kijiji_search":                 {path: "/v2/kijiji/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"kroger_product":                {path: "/v2/kroger/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"kroger_search":                 {path: "/v2/kroger/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"lamudi_ad":                     {path: "/v2/lamudi/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"lamudi_search":                 {path: "/v2/lamudi/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"leboncoin_ad":                  {path: "/v2/leboncoin/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"leboncoin_search":              {path: "/v2/leboncoin/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"leboncoin_search_api":          {path: "/v2/leboncoin/search_api", methods: []string{"POST"}, json: true, required: []string{}},
	"lenovo_product":                {path: "/v2/lenovo/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"lenovo_search":                 {path: "/v2/lenovo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"leroymerlin_product":           {path: "/v2/leroymerlin/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"leroymerlin_search":            {path: "/v2/leroymerlin/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"lidl_product":                  {path: "/v2/lidl/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"lidl_search":                   {path: "/v2/lidl/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedin_company_info":         {path: "/v2/linkedin/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"linkedin_job_count":            {path: "/v2/linkedin/job/count", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"linkedin_job_info":             {path: "/v2/linkedin/job/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedin_job_search":           {path: "/v2/linkedin/job/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"linkedin_job_suggest":          {path: "/v2/linkedin/job/suggest", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedin_post_info":            {path: "/v2/linkedin/post/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedin_product_info":         {path: "/v2/linkedin/product/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedin_profile_info":         {path: "/v2/linkedin/profile/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedinadlibrary_ad":          {path: "/v2/linkedinadlibrary/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"linkedinadlibrary_search":      {path: "/v2/linkedinadlibrary/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"loopnet_property":              {path: "/v2/loopnet/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"loopnet_search":                {path: "/v2/loopnet/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"maersk_tracking":               {path: "/v2/maersk/tracking", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"manomano_product":              {path: "/v2/manomano/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"manomano_product_offer":        {path: "/v2/manomano/product/offer", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"manomano_search":               {path: "/v2/manomano/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mascus_ad":                     {path: "/v2/mascus/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mascus_search":                 {path: "/v2/mascus/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mediamarkt_product":            {path: "/v2/mediamarkt/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mediamarkt_search":             {path: "/v2/mediamarkt/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"micromania_product":            {path: "/v2/micromania/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"micromania_search":             {path: "/v2/micromania/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mobile_ad":                     {path: "/v2/mobile/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mobile_search":                 {path: "/v2/mobile/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"morele_product":                {path: "/v2/morele/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"morele_search":                 {path: "/v2/morele/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"mrbricolage_suggest":           {path: "/v2/mrbricolage/suggest", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"newegg_product":                {path: "/v2/newegg/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"newegg_search":                 {path: "/v2/newegg/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"northdata_company":             {path: "/v2/northdata/company", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"northdata_search":              {path: "/v2/northdata/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"npi_provider":                  {path: "/v2/npi/provider", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"npi_search":                    {path: "/v2/npi/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"oceanio_search":                {path: "/v2/oceanio/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"olx_ad":                        {path: "/v2/olx/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"olx_product":                   {path: "/v2/olx/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"olx_search":                    {path: "/v2/olx/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"oscaro_product":                {path: "/v2/oscaro/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"oscaro_search":                 {path: "/v2/oscaro/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ovoko_ad":                      {path: "/v2/ovoko/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ovoko_search":                  {path: "/v2/ovoko/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"owler_company_info":            {path: "/v2/owler/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"owler_search":                  {path: "/v2/owler/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pagesjaunes_page_info":         {path: "/v2/pagesjaunes/page/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pagesjaunes_search":            {path: "/v2/pagesjaunes/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"patreon_profile_info":          {path: "/v2/patreon/profile/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"patreon_search":                {path: "/v2/patreon/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"petsmart_product":              {path: "/v2/petsmart/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"petsmart_search":               {path: "/v2/petsmart/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"picard_product":                {path: "/v2/picard/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"picard_search":                 {path: "/v2/picard/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pinterest_user_info":           {path: "/v2/pinterest/user/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pitchbook_company_info":        {path: "/v2/pitchbook/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pitchbook_search":              {path: "/v2/pitchbook/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pointp_product":                {path: "/v2/pointp/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"pointp_search":                 {path: "/v2/pointp/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"prnewswire_article":            {path: "/v2/prnewswire/article", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"prnewswire_search":             {path: "/v2/prnewswire/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"producthunt_product_info":      {path: "/v2/producthunt/product/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"proshop_product":               {path: "/v2/proshop/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"proshop_search":                {path: "/v2/proshop/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rappi_product":                 {path: "/v2/rappi/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rappi_search":                  {path: "/v2/rappi/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ratemyprofessors_professor":    {path: "/v2/ratemyprofessors/professor", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ratemyprofessors_search":       {path: "/v2/ratemyprofessors/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"realestatecomau_search":        {path: "/v2/realestatecomau/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"realtor_property":              {path: "/v2/realtor/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"realtor_search":                {path: "/v2/realtor/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"redfin_property":               {path: "/v2/redfin/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"redfin_search":                 {path: "/v2/redfin/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"reichelt_product":              {path: "/v2/reichelt/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"reichelt_search":               {path: "/v2/reichelt/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"reverb_product":                {path: "/v2/reverb/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"reverb_search":                 {path: "/v2/reverb/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rexel_product":                 {path: "/v2/rexel/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rexel_search":                  {path: "/v2/rexel/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rightmove_property":            {path: "/v2/rightmove/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rightmove_search":              {path: "/v2/rightmove/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rsonline_product":              {path: "/v2/rsonline/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"rsonline_search":               {path: "/v2/rsonline/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seek_job":                      {path: "/v2/seek/job", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seek_search":                   {path: "/v2/seek/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seloger_agency":                {path: "/v2/seloger/agency", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seloger_agency_rent":           {path: "/v2/seloger/agency/rent", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seloger_agency_sale":           {path: "/v2/seloger/agency/sale", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seloger_property":              {path: "/v2/seloger/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"seloger_search":                {path: "/v2/seloger/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"sephora_product":               {path: "/v2/sephora/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"sephora_search":                {path: "/v2/sephora/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"shophouzz_product":             {path: "/v2/shophouzz/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"shopify_apps":                  {path: "/v2/shopify/apps", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"shopify_catalog":               {path: "/v2/shopify/catalog", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"shopify_collections":           {path: "/v2/shopify/collections", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"shopify_product":               {path: "/v2/shopify/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"similarweb_domain":             {path: "/v2/similarweb/domain", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"similarweb_search":             {path: "/v2/similarweb/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"snapeda_product":               {path: "/v2/snapeda/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"snapeda_search":                {path: "/v2/snapeda/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"starbucks_menu":                {path: "/v2/starbucks/menu", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"starbucks_product":             {path: "/v2/starbucks/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"starbucks_search":              {path: "/v2/starbucks/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"stats":                         {path: "/v2/stats", methods: []string{"GET"}, json: true, required: []string{}},
	"stockx_product":                {path: "/v2/stockx/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"stockx_search":                 {path: "/v2/stockx/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"stockx_trends":                 {path: "/v2/stockx/trends", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"stubhub_event":                 {path: "/v2/stubhub/event", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"stubhub_search":                {path: "/v2/stubhub/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"target_product":                {path: "/v2/target/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"target_search":                 {path: "/v2/target/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"tdlr_license":                  {path: "/v2/tdlr/license", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"tdlr_search":                   {path: "/v2/tdlr/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"ted_notice":                    {path: "/v2/ted/notice", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ted_search":                    {path: "/v2/ted/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"thalia_product":                {path: "/v2/thalia/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"thalia_search":                 {path: "/v2/thalia/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"thefork_restaurant":            {path: "/v2/thefork/restaurant", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"thefork_search":                {path: "/v2/thefork/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"thomasnet_company":             {path: "/v2/thomasnet/company", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"thomasnet_search":              {path: "/v2/thomasnet/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ticketmaster_event":            {path: "/v2/ticketmaster/event", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"ticketmaster_search":           {path: "/v2/ticketmaster/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"tiktoktopads_ad":               {path: "/v2/tiktoktopads/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"tiktoktopads_search":           {path: "/v2/tiktoktopads/search", methods: []string{"GET", "POST"}, json: true, required: []string{}},
	"tradeaplane_listing":           {path: "/v2/tradeaplane/listing", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"tradeaplane_search":            {path: "/v2/tradeaplane/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"trustpilot_company_info":       {path: "/v2/trustpilot/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"untappd_beer":                  {path: "/v2/untappd/beer", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"untappd_search":                {path: "/v2/untappd/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"upwork_freelancer":             {path: "/v2/upwork/freelancer", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"upwork_jobs":                   {path: "/v2/upwork/jobs", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"upwork_services":               {path: "/v2/upwork/services", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"upwork_services_info":          {path: "/v2/upwork/services/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"usage":                         {path: "/v2/usage", methods: []string{"GET"}, json: true, required: []string{}},
	"usaspending_award":             {path: "/v2/usaspending/award", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"usaspending_search":            {path: "/v2/usaspending/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vestiairecollective_product":   {path: "/v2/vestiairecollective/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vestiairecollective_search":    {path: "/v2/vestiairecollective/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vinted_item":                   {path: "/v2/vinted/item", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vinted_search":                 {path: "/v2/vinted/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vinted_user_info":              {path: "/v2/vinted/user/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vivareal_property":             {path: "/v2/vivareal/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"vivareal_search":               {path: "/v2/vivareal/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"walmart_product":               {path: "/v2/walmart/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"walmart_search":                {path: "/v2/walmart/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"walmart_suggest":               {path: "/v2/walmart/suggest", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_antibot":               {path: "/v2/website/antibot", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"website_crawler":               {path: "/v2/website/crawler", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_email_phone_extractor": {path: "/v2/website/email_phone_extractor", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_rendering":             {path: "/v2/website/rendering", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_screenshot":            {path: "/v2/website/screenshot", methods: []string{"GET", "POST"}, json: false, required: []string{"url"}},
	"website_technology":            {path: "/v2/website/technology", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_webunlocker":           {path: "/v2/website/webunlocker", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"weibo_post_info":               {path: "/v2/weibo/post/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"weibo_search":                  {path: "/v2/weibo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"welcometothejungle_company":    {path: "/v2/welcometothejungle/company", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"welcometothejungle_search":     {path: "/v2/welcometothejungle/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"wellfound_company_info":        {path: "/v2/wellfound/company/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"x_post_info":                   {path: "/v2/x/post/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"x_profile_info":                {path: "/v2/x/profile/info", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"yahoo_search":                  {path: "/v2/yahoo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"yellowpages_business":          {path: "/v2/yellowpages/business", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"yellowpages_search":            {path: "/v2/yellowpages/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"yelp_business":                 {path: "/v2/yelp/business", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"yelp_reviews":                  {path: "/v2/yelp/reviews", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"yelp_search":                   {path: "/v2/yelp/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zapimoveis_ad":                 {path: "/v2/zapimoveis/ad", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zapimoveis_search":             {path: "/v2/zapimoveis/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zazzle_product":                {path: "/v2/zazzle/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zazzle_search":                 {path: "/v2/zazzle/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zillow_property":               {path: "/v2/zillow/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zillow_search":                 {path: "/v2/zillow/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zonaprop_property":             {path: "/v2/zonaprop/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zonaprop_search":               {path: "/v2/zonaprop/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zoopla_property":               {path: "/v2/zoopla/property", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zoopla_search":                 {path: "/v2/zoopla/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zooplus_product":               {path: "/v2/zooplus/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zooplus_search":                {path: "/v2/zooplus/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zoro_product":                  {path: "/v2/zoro/product", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"zoro_search":                   {path: "/v2/zoro/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
}
