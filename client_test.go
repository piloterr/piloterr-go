package piloterr

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type recorded struct {
	request *http.Request
	body    []byte
}

func newTestClient(t *testing.T, respond func(*http.Request) *http.Response) (*Client, *[]recorded) {
	t.Helper()
	var calls []recorded
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		var body []byte
		if r.Body != nil {
			body, _ = io.ReadAll(r.Body)
		}
		calls = append(calls, recorded{request: r, body: body})
		return respond(r), nil
	})
	client, err := NewClient("test-key", &ClientOptions{HTTPClient: &http.Client{Transport: transport}})
	if err != nil {
		t.Fatal(err)
	}
	return client, &calls
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json; charset=utf-8"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func okJSON(*http.Request) *http.Response { return jsonResponse(200, `{"ok":true}`) }

func TestNewClientRequiresAPIKey(t *testing.T) {
	t.Setenv(APIKeyEnvironmentVariable, "")
	var configurationError *ConfigurationError
	if _, err := NewClient("", nil); !errors.As(err, &configurationError) {
		t.Fatalf("expected ConfigurationError, got %v", err)
	}
	if _, err := NewClient("   ", nil); !errors.As(err, &configurationError) {
		t.Fatalf("expected ConfigurationError, got %v", err)
	}
	if _, err := NewClient("k", &ClientOptions{BaseURL: "ftp://x"}); !errors.As(err, &configurationError) {
		t.Fatalf("expected ConfigurationError for BaseURL, got %v", err)
	}
}

func TestNewClientReadsEnvironment(t *testing.T) {
	t.Setenv(APIKeyEnvironmentVariable, "env-key")
	client, err := NewClient("", nil)
	if err != nil || client.apiKey != "env-key" {
		t.Fatalf("unexpected client %v %v", client, err)
	}
}

func TestFeaturedOperation(t *testing.T) {
	client, calls := newTestClient(t, func(*http.Request) *http.Response {
		return jsonResponse(200, `{"hello":"world"}`)
	})
	ctx := context.Background()
	response, err := client.AmazonProduct(ctx, &AmazonProductParams{Query: "B09DFCB66S"})
	if err != nil {
		t.Fatal(err)
	}
	data, err := response.Map()
	if err != nil || data["hello"] != "world" {
		t.Fatalf("unexpected data %v %v", data, err)
	}
	request := (*calls)[0].request
	if request.Method != "GET" || request.URL.Scheme+"://"+request.URL.Host+request.URL.Path != "https://api.piloterr.com/v2/amazon/product" {
		t.Fatalf("unexpected request %s %s", request.Method, request.URL)
	}
	if request.Header.Get("x-api-key") != "test-key" || !strings.HasPrefix(request.Header.Get("User-Agent"), "piloterr-go/") {
		t.Fatalf("unexpected headers %v", request.Header)
	}
}

func TestQuerySerialization(t *testing.T) {
	client, calls := newTestClient(t, okJSON)
	_, err := client.Request(context.Background(), "GET", "/v1/anything", map[string]any{
		"text": "a b&c", "flag": false, "count": 0, "ratio": 1.5, "whole": 2.0, "list": []string{"x", "y"}, "nothing": nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	got := (*calls)[0].request.URL.RawQuery
	want := "count=0&flag=false&list=x&list=y&ratio=1.5&text=a%20b%26c&whole=2"
	if got != want {
		t.Fatalf("query = %q, want %q", got, want)
	}
}

func TestPostSendsJSONBody(t *testing.T) {
	client, calls := newTestClient(t, okJSON)
	_, err := client.Request(context.Background(), "POST", "/v1/anything", map[string]any{"query": "x", "page": 2, "gone": nil})
	if err != nil {
		t.Fatal(err)
	}
	call := (*calls)[0]
	var body map[string]any
	if err := json.Unmarshal(call.body, &body); err != nil {
		t.Fatal(err)
	}
	if call.request.Method != "POST" || call.request.Header.Get("Content-Type") != "application/json" || body["query"] != "x" || body["page"] != 2.0 || len(body) != 2 {
		t.Fatalf("unexpected POST %s %v", call.request.Method, body)
	}
}

func TestPostKeepsGlobalParametersInQuery(t *testing.T) {
	client, calls := newTestClient(t, okJSON)
	_, err := client.Request(context.Background(), "POST", "/v1/anything", map[string]any{"query": "x", "is_async": true, "fields": "a"})
	if err != nil {
		t.Fatal(err)
	}
	call := (*calls)[0]
	var body map[string]any
	if err := json.Unmarshal(call.body, &body); err != nil {
		t.Fatal(err)
	}
	query := call.request.URL.Query()
	if len(body) != 1 || body["query"] != "x" || query.Get("is_async") != "true" || query.Get("fields") != "a" {
		t.Fatalf("unexpected POST %v ?%s", body, call.request.URL.RawQuery)
	}
}

func TestPathParameters(t *testing.T) {
	client, calls := newTestClient(t, okJSON)
	if _, err := client.Request(context.Background(), "GET", "/v1/items/{id}", map[string]any{"id": "a/b c", "other": 1}); err != nil {
		t.Fatal(err)
	}
	if got := (*calls)[0].request.URL.RawPath; got != "/v1/items/a%2Fb%20c" {
		t.Fatalf("path = %q", got)
	}
	var serializationError *SerializationError
	if _, err := client.Request(context.Background(), "GET", "/v1/items/{id}", nil); !errors.As(err, &serializationError) {
		t.Fatalf("expected SerializationError, got %v", err)
	}
}

func TestRequiredParameters(t *testing.T) {
	client, calls := newTestClient(t, okJSON)
	ctx := context.Background()
	_, err := client.AmazonProduct(ctx, &AmazonProductParams{})
	var serializationError *SerializationError
	if !errors.As(err, &serializationError) || len(*calls) != 0 {
		t.Fatalf("expected SerializationError before sending, got %v", err)
	}
}

func TestUnsupportedMethod(t *testing.T) {
	client, _ := newTestClient(t, okJSON)
	ctx := context.Background()
	_, err := client.Jobs(ctx, &JobsParams{ID: "test"}, WithMethod("POST"))
	var configurationError *ConfigurationError
	if !errors.As(err, &configurationError) {
		t.Fatalf("expected ConfigurationError, got %v", err)
	}
}

func TestEveryEndpoint(t *testing.T) {
	client, calls := newTestClient(t, func(*http.Request) *http.Response { return jsonResponse(200, `{}`) })
	ctx := context.Background()
	checks := []struct {
		name string
		call func() (*Response, error)
	}{
		{"aboutyou_product", func() (*Response, error) { return client.AboutyouProduct(ctx, &AboutyouProductParams{Query: "test"}) }},
		{"aboutyou_search", func() (*Response, error) {
			return client.AboutyouSearch(ctx, &AboutyouSearchParams{Query: "suede bloom"})
		}},
		{"action_product", func() (*Response, error) { return client.ActionProduct(ctx, &ActionProductParams{Query: "3216159"}) }},
		{"action_search", func() (*Response, error) { return client.ActionSearch(ctx, &ActionSearchParams{Query: "meuble"}) }},
		{"aldi_product", func() (*Response, error) {
			return client.AldiProduct(ctx, &AldiProductParams{Query: "fiches-produits/filets-danchois-marines-6806.html"})
		}},
		{"aldi_search", func() (*Response, error) {
			return client.AldiSearch(ctx, &AldiSearchParams{Query: "produits/viande-poisson/poisson-crustace.html"})
		}},
		{"alibaba_product", func() (*Response, error) {
			return client.AlibabaProduct(ctx, &AlibabaProductParams{Query: "https://www.alibaba.com/product-detail/Custom-Neon-Light-Acrylic-Neon-Sign_1601175379813.html"})
		}},
		{"alibaba_search", func() (*Response, error) {
			return client.AlibabaSearch(ctx, &AlibabaSearchParams{Query: "https://www.alibaba.com/trade/search?SearchText=neon"})
		}},
		{"aliexpress_product", func() (*Response, error) {
			return client.AliexpressProduct(ctx, &AliexpressProductParams{Query: "https://www.aliexpress.com/item/1005009517624640.html"})
		}},
		{"aliexpress_product_reviews", func() (*Response, error) {
			return client.AliexpressProductReviews(ctx, &AliexpressProductReviewsParams{Query: "1005001388303149"})
		}},
		{"aliexpress_search", func() (*Response, error) { return client.AliexpressSearch(ctx, &AliexpressSearchParams{Query: "meta"}) }},
		{"aliexpress_store_product", func() (*Response, error) {
			return client.AliexpressStoreProduct(ctx, &AliexpressStoreProductParams{Query: "https://www.aliexpress.com/store/1101705614"})
		}},
		{"allegro_product", func() (*Response, error) {
			return client.AllegroProduct(ctx, &AllegroProductParams{Query: "https://allegro.pl/produkt/slug-uuid?offerId=123"})
		}},
		{"allegro_search", func() (*Response, error) { return client.AllegroSearch(ctx, &AllegroSearchParams{Query: "mufowa"}) }},
		{"amazon_product", func() (*Response, error) { return client.AmazonProduct(ctx, &AmazonProductParams{Query: "B09DFCB66S"}) }},
		{"amazon_product_offer", func() (*Response, error) {
			return client.AmazonProductOffer(ctx, &AmazonProductOfferParams{Query: "B093C9B1HK"})
		}},
		{"amazon_search", func() (*Response, error) {
			return client.AmazonSearch(ctx, &AmazonSearchParams{Query: "playstation 5"})
		}},
		{"apartments_property", func() (*Response, error) {
			return client.ApartmentsProperty(ctx, &ApartmentsPropertyParams{Query: "https://www.apartments.com/141-e-33rd-st-new-york-ny-unit-si-fl16-id43/p6b4mes/"})
		}},
		{"apartments_search", func() (*Response, error) {
			return client.ApartmentsSearch(ctx, &ApartmentsSearchParams{Query: "https://www.apartments.com/new-york-ny/3"})
		}},
		{"argenprop_property", func() (*Response, error) {
			return client.ArgenpropProperty(ctx, &ArgenpropPropertyParams{Query: "https://www.argenprop.com/departamento-en-venta-en-palermo-chico-5-ambientes--20724993"})
		}},
		{"argenprop_search", func() (*Response, error) {
			return client.ArgenpropSearch(ctx, &ArgenpropSearchParams{Query: "palermo"})
		}},
		{"argos_product", func() (*Response, error) { return client.ArgosProduct(ctx, &ArgosProductParams{Query: "test"}) }},
		{"argos_search", func() (*Response, error) { return client.ArgosSearch(ctx, &ArgosSearchParams{Query: "laptop"}) }},
		{"arxiv_paper", func() (*Response, error) {
			return client.ArxivPaper(ctx, &ArxivPaperParams{Query: "2201.00978 or https://arxiv.org/abs/cond-mat/0011267v1."})
		}},
		{"arxiv_search", func() (*Response, error) {
			return client.ArxivSearch(ctx, &ArxivSearchParams{Query: "electron. A paper URL such as https://arxiv.org/abs/2201.00978 returns 400."})
		}},
		{"auchan_product", func() (*Response, error) {
			return client.AuchanProduct(ctx, &AuchanProductParams{Query: "https://www.auchan.fr/charles-alice-specialite-pommes-mangues-sans-sucres-ajoutes/pr-C1177497"})
		}},
		{"auchan_search", func() (*Response, error) { return client.AuchanSearch(ctx, &AuchanSearchParams{Query: "pomme"}) }},
		{"autodoc_product", func() (*Response, error) {
			return client.AutodocProduct(ctx, &AutodocProductParams{Query: "https://www.auto-doc.fr/ridex/7999541"})
		}},
		{"autodoc_product_related", func() (*Response, error) {
			return client.AutodocProductRelated(ctx, &AutodocProductRelatedParams{Query: "https://www.auto-doc.fr/ridex/7999541", MakerID: 93, VehicleType: "pkw", MakerName: "RENAULT", ProductID: 7999541})
		}},
		{"autodoc_search", func() (*Response, error) {
			return client.AutodocSearch(ctx, &AutodocSearchParams{Query: "https://www.auto-doc.fr/search?keyword=frein"})
		}},
		{"autoeurope_result", func() (*Response, error) {
			return client.AutoeuropeResult(ctx, &AutoeuropeResultParams{OfferID: "3CD30DE2D9BE4CE8FC28E20D2702FCC7", Pickup: "CDG", PickupDate: "2026-11-06", DropoffDate: "2026-11-09"})
		}},
		{"autoeurope_search", func() (*Response, error) {
			return client.AutoeuropeSearch(ctx, &AutoeuropeSearchParams{Pickup: "CDG", PickupDate: "2026-11-06", DropoffDate: "2026-11-09"})
		}},
		{"autoscout24_ad", func() (*Response, error) {
			return client.Autoscout24Ad(ctx, &Autoscout24AdParams{Query: "https://www.autoscout24.fr/offres/porsche-macan-s-diesel-pano-1-hand-diesel-or-cat_ma57mo20311-547632b2-e4fc-47f3-8143-1a2d7b276366"})
		}},
		{"autoscout24_search", func() (*Response, error) {
			return client.Autoscout24Search(ctx, &Autoscout24SearchParams{Query: "https://www.autoscout24.fr/lst/volkswagen/golf?atype=C"})
		}},
		{"autotrader_ad", func() (*Response, error) {
			return client.AutotraderAd(ctx, &AutotraderAdParams{Query: "https://www.autotrader.co.uk/car-details/202606053030491"})
		}},
		{"autotrader_search", func() (*Response, error) {
			return client.AutotraderSearch(ctx, &AutotraderSearchParams{Query: "https://www.autotrader.co.uk/car-search?channel=cars&postcode=W1F%208EB&price_search_type=total&sort=relevance&page=1"})
		}},
		{"bestbuy_product", func() (*Response, error) {
			return client.BestbuyProduct(ctx, &BestbuyProductParams{Query: "https://www.bestbuy.com/product/gigabyte-aero-x16-copilot-pc-165hz-wqxga-rtx-5070-ryzen-ai-7-350-1tb-ssd-16gb-ddr5-windows-11-lunar-white/J3GWPQC2RQ"})
		}},
		{"bestbuy_product_reviews", func() (*Response, error) {
			return client.BestbuyProductReviews(ctx, &BestbuyProductReviewsParams{Query: "test"})
		}},
		{"bestbuy_search", func() (*Response, error) { return client.BestbuySearch(ctx, &BestbuySearchParams{Query: "machine"}) }},
		{"bhphotovideo_product", func() (*Response, error) {
			return client.BhphotovideoProduct(ctx, &BhphotovideoProductParams{Query: "1948568-REG"})
		}},
		{"bhphotovideo_search", func() (*Response, error) {
			return client.BhphotovideoSearch(ctx, &BhphotovideoSearchParams{Query: "sandisk extreme pro"})
		}},
		{"bing_search", func() (*Response, error) { return client.BingSearch(ctx, &BingSearchParams{Query: "openai"}) }},
		{"bitiba_product", func() (*Response, error) {
			return client.BitibaProduct(ctx, &BitibaProductParams{Query: "75036. A search link returns 400."})
		}},
		{"bitiba_search", func() (*Response, error) { return client.BitibaSearch(ctx, &BitibaSearchParams{Query: "test"}) }},
		{"bol_product", func() (*Response, error) { return client.BolProduct(ctx, &BolProductParams{Query: "9300000240171676"}) }},
		{"bol_search", func() (*Response, error) { return client.BolSearch(ctx, &BolSearchParams{Query: "lego"}) }},
		{"boulanger_product", func() (*Response, error) {
			return client.BoulangerProduct(ctx, &BoulangerProductParams{Query: "https://www.boulanger.com/ref/1193303"})
		}},
		{"boulanger_search", func() (*Response, error) {
			return client.BoulangerSearch(ctx, &BoulangerSearchParams{Query: "aspirateur"})
		}},
		{"brave_search", func() (*Response, error) { return client.BraveSearch(ctx, &BraveSearchParams{Query: "anthropic"}) }},
		{"bricodepot_product", func() (*Response, error) {
			return client.BricodepotProduct(ctx, &BricodepotProductParams{Query: "https://www.bricodepot.fr/p/3373270014082/compresseur-coaxial-50l-vertical-10-bars-accessoires-1"})
		}},
		{"bricodepot_search", func() (*Response, error) {
			return client.BricodepotSearch(ctx, &BricodepotSearchParams{Query: "https://www.bricodepot.fr/produits/outillage/outillage-electroportatif-a-main-et-accessoire/perceuse-visseuse-et-outillage-de-percage/perceuse-visseuse-et-perceuse-a-percussion-sans-fil"})
		}},
		{"bricorama_product", func() (*Response, error) {
			return client.BricoramaProduct(ctx, &BricoramaProductParams{Query: "https://www.bricorama.fr/p/perceuse-visseuse-sans-fil-18v/3999991691267"})
		}},
		{"bricorama_search", func() (*Response, error) {
			return client.BricoramaSearch(ctx, &BricoramaSearchParams{Query: "perceuse"})
		}},
		{"carrefour_product", func() (*Response, error) {
			return client.CarrefourProduct(ctx, &CarrefourProductParams{Query: "https://www.carrefour.fr/p/oeufs-de-caille-le-gaulois-3266980075934"})
		}},
		{"carrefour_search", func() (*Response, error) { return client.CarrefourSearch(ctx, &CarrefourSearchParams{Query: "pomme"}) }},
		{"carrefour_suggest", func() (*Response, error) { return client.CarrefourSuggest(ctx, &CarrefourSuggestParams{Query: "coca"}) }},
		{"castorama_product", func() (*Response, error) {
			return client.CastoramaProduct(ctx, &CastoramaProductParams{Query: "https://www.castorama.fr/mastic-bain-et-cuisine-rubson-blanc-280ml/5410091628840_CAFR.prd"})
		}},
		{"castorama_search", func() (*Response, error) {
			return client.CastoramaSearch(ctx, &CastoramaSearchParams{Query: "joint-silicone"})
		}},
		{"cdiscount_product", func() (*Response, error) {
			return client.CdiscountProduct(ctx, &CdiscountProductParams{Query: "https://www.cdiscount.com/informatique/ordinateurs-pc-portables/acemagic-pc-portables-gaming/f-1070992-aaaph06865.html"})
		}},
		{"cdiscount_product_offer", func() (*Response, error) {
			return client.CdiscountProductOffer(ctx, &CdiscountProductOfferParams{Query: "https://www.cdiscount.com/mp-352880-auc3283981039048.html"})
		}},
		{"cdiscount_search", func() (*Response, error) {
			return client.CdiscountSearch(ctx, &CdiscountSearchParams{Query: "https://www.cdiscount.com/search/10/laptop.html"})
		}},
		{"centris_property", func() (*Response, error) {
			return client.CentrisProperty(ctx, &CentrisPropertyParams{Query: "https://www.centris.ca/en/condos~for-sale~montreal-le-sud-ouest/17796540"})
		}},
		{"centris_search", func() (*Response, error) { return client.CentrisSearch(ctx, &CentrisSearchParams{Query: "montreal"}) }},
		{"chewy_product", func() (*Response, error) { return client.ChewyProduct(ctx, &ChewyProductParams{Query: "test"}) }},
		{"chewy_search", func() (*Response, error) { return client.ChewySearch(ctx, &ChewySearchParams{Query: "test"}) }},
		{"clearedjobs_job", func() (*Response, error) {
			return client.ClearedjobsJob(ctx, &ClearedjobsJobParams{Query: "analyst-washington-district-of-columbia-1902823."})
		}},
		{"clearedjobs_search", func() (*Response, error) {
			return client.ClearedjobsSearch(ctx, &ClearedjobsSearchParams{Query: "analyst. A job URL returns 400."})
		}},
		{"company", func() (*Response, error) { return client.Company(ctx, &CompanyParams{Query: "gucci.com"}) }},
		{"coolblue_product", func() (*Response, error) {
			return client.CoolblueProduct(ctx, &CoolblueProductParams{Query: "https://www.coolblue.be/nl/product/905648/sony-wh-1000xm5-zwart.html"})
		}},
		{"coolblue_search", func() (*Response, error) {
			return client.CoolblueSearch(ctx, &CoolblueSearchParams{Query: "sony wh-1000xm5"})
		}},
		{"copart_lot", func() (*Response, error) { return client.CopartLot(ctx, &CopartLotParams{Query: "55295666"}) }},
		{"copart_search", func() (*Response, error) { return client.CopartSearch(ctx, &CopartSearchParams{Query: "chevrolet"}) }},
		{"costco_product", func() (*Response, error) { return client.CostcoProduct(ctx, &CostcoProductParams{Query: "4000266385"}) }},
		{"costco_search", func() (*Response, error) { return client.CostcoSearch(ctx, &CostcoSearchParams{Query: "home"}) }},
		{"crunchbase_company_info", func() (*Response, error) { return client.CrunchbaseCompanyInfo(ctx, &CrunchbaseCompanyInfoParams{}) }},
		{"crunchbase_event", func() (*Response, error) {
			return client.CrunchbaseEvent(ctx, &CrunchbaseEventParams{Query: "london-tech-week-2025"})
		}},
		{"crunchbase_funding_round", func() (*Response, error) {
			return client.CrunchbaseFundingRound(ctx, &CrunchbaseFundingRoundParams{Query: "openai-secondary-market--3bab6e98"})
		}},
		{"crunchbase_funding_rounds", func() (*Response, error) {
			return client.CrunchbaseFundingRounds(ctx, &CrunchbaseFundingRoundsParams{})
		}},
		{"crunchbase_people_info", func() (*Response, error) {
			return client.CrunchbasePeopleInfo(ctx, &CrunchbasePeopleInfoParams{Query: "alex-martin"})
		}},
		{"crunchbase_search", func() (*Response, error) {
			return client.CrunchbaseSearch(ctx, &CrunchbaseSearchParams{Query: "openai"})
		}},
		{"cultura_product", func() (*Response, error) {
			return client.CulturaProduct(ctx, &CulturaProductParams{Query: "https://www.cultura.com/p-harry-potter-et-le-prisonnier-d-azkaban-illustre-par-minalima-9782075193993.html"})
		}},
		{"cultura_search", func() (*Response, error) {
			return client.CulturaSearch(ctx, &CulturaSearchParams{Query: "harry potter"})
		}},
		{"darty_product", func() (*Response, error) {
			return client.DartyProduct(ctx, &DartyProductParams{Query: "https://www.darty.com/nav/achat/petit_electromenager/aspirateur-balai_main/aspirateur_balai/dyson_piston_animal.html"})
		}},
		{"darty_search", func() (*Response, error) { return client.DartySearch(ctx, &DartySearchParams{Query: "aspirateur"}) }},
		{"datasets_linkedin_company", func() (*Response, error) {
			return client.DatasetsLinkedinCompany(ctx, &DatasetsLinkedinCompanyParams{})
		}},
		{"decathlon_product", func() (*Response, error) {
			return client.DecathlonProduct(ctx, &DecathlonProductParams{Query: "https://www.decathlon.fr/p/gourde-alu-1l-avec-bouchon-instantane-composee-a-100percent-dalu-recycle-bleu/330041/c98m8612536"})
		}},
		{"decathlon_search", func() (*Response, error) { return client.DecathlonSearch(ctx, &DecathlonSearchParams{Query: "gourde"}) }},
		{"digikey_product", func() (*Response, error) {
			return client.DigikeyProduct(ctx, &DigikeyProductParams{Query: "https://www.digikey.fr/fr/products/detail/phoenix-contact/2981965/3604293"})
		}},
		{"digikey_search", func() (*Response, error) {
			return client.DigikeySearch(ctx, &DigikeySearchParams{Query: "phoenix contact"})
		}},
		{"doaj_article", func() (*Response, error) {
			return client.DoajArticle(ctx, &DoajArticleParams{Query: "000004687a3a411cb466625281f2ceb0."})
		}},
		{"doaj_journal", func() (*Response, error) { return client.DoajJournal(ctx, &DoajJournalParams{Query: "2747-173X."}) }},
		{"doaj_search", func() (*Response, error) {
			return client.DoajSearch(ctx, &DoajSearchParams{Query: "climate. An article or journal URL returns 400."})
		}},
		{"domain_dnsbl", func() (*Response, error) { return client.DomainDnsbl(ctx, &DomainDnsblParams{Query: "185.122.10.10"}) }},
		{"domain_malicious", func() (*Response, error) {
			return client.DomainMalicious(ctx, &DomainMaliciousParams{Query: "197.33.189.70"})
		}},
		{"domain_whois", func() (*Response, error) { return client.DomainWhois(ctx, &DomainWhoisParams{Query: "piloterr.com"}) }},
		{"domaincomau_property", func() (*Response, error) {
			return client.DomaincomauProperty(ctx, &DomaincomauPropertyParams{Query: "https://www.domain.com.au/7d-27-31-ocean-street-n-bondi-nsw-2026-2021178032"})
		}},
		{"domaincomau_search", func() (*Response, error) {
			return client.DomaincomauSearch(ctx, &DomaincomauSearchParams{Query: "Bondi NSW 2026"})
		}},
		{"duckduckgo_search", func() (*Response, error) { return client.DuckduckgoSearch(ctx, &DuckduckgoSearchParams{Query: "test"}) }},
		{"ebay_product", func() (*Response, error) { return client.EbayProduct(ctx, &EbayProductParams{Query: "257758922911"}) }},
		{"ebay_search", func() (*Response, error) { return client.EbaySearch(ctx, &EbaySearchParams{Query: "laptop"}) }},
		{"elcorteingles_product", func() (*Response, error) {
			return client.ElcorteinglesProduct(ctx, &ElcorteinglesProductParams{Query: "A200910175"})
		}},
		{"elcorteingles_search", func() (*Response, error) {
			return client.ElcorteinglesSearch(ctx, &ElcorteinglesSearchParams{Query: "marino"})
		}},
		{"eleclerc_product", func() (*Response, error) {
			return client.EleclercProduct(ctx, &EleclercProductParams{Query: "3564700201442"})
		}},
		{"eleclerc_search", func() (*Response, error) { return client.EleclercSearch(ctx, &EleclercSearchParams{Query: "pomme"}) }},
		{"eleclerc_suggest", func() (*Response, error) { return client.EleclercSuggest(ctx, &EleclercSuggestParams{Query: "pomme"}) }},
		{"email_analyzes", func() (*Response, error) { return client.EmailAnalyzes(ctx, &EmailAnalyzesParams{Query: "airbus.com"}) }},
		{"email_finder", func() (*Response, error) { return client.EmailFinder(ctx, &EmailFinderParams{Query: "Jane Doe"}) }},
		{"email_verify", func() (*Response, error) {
			return client.EmailVerify(ctx, &EmailVerifyParams{Query: "jane.doe@example.com"})
		}},
		{"europages_company", func() (*Response, error) {
			return client.EuropagesCompany(ctx, &EuropagesCompanyParams{Query: "https://www.europages.fr/fr/company/synyco-sas-22106635"})
		}},
		{"europages_search", func() (*Response, error) {
			return client.EuropagesSearch(ctx, &EuropagesSearchParams{Query: "emballage"})
		}},
		{"eventbrite_event", func() (*Response, error) {
			return client.EventbriteEvent(ctx, &EventbriteEventParams{Query: "2002625586037"})
		}},
		{"eventbrite_search", func() (*Response, error) { return client.EventbriteSearch(ctx, &EventbriteSearchParams{}) }},
		{"expedia_ad", func() (*Response, error) {
			return client.ExpediaAd(ctx, &ExpediaAdParams{Query: "https://www.expedia.com/New-York-Hotels-Hilton-New-York-Fashion-District.h3117215.Hotel-Information"})
		}},
		{"expedia_flight", func() (*Response, error) { return client.ExpediaFlight(ctx, &ExpediaFlightParams{Query: "JFK-LAX"}) }},
		{"expedia_search", func() (*Response, error) {
			return client.ExpediaSearch(ctx, &ExpediaSearchParams{Query: "https://euro.expedia.net/Hotel-Search?currency=EUR&siteid=4400&destination=Toulouse"})
		}},
		{"fashionphile_product", func() (*Response, error) {
			return client.FashionphileProduct(ctx, &FashionphileProductParams{Query: "celine-supple-shiny-lambskin-little-new-luggage-black-1978804"})
		}},
		{"fashionphile_search", func() (*Response, error) {
			return client.FashionphileSearch(ctx, &FashionphileSearchParams{Query: "celine"})
		}},
		{"finder_chewy_products", func() (*Response, error) { return client.FinderChewyProducts(ctx, &FinderChewyProductsParams{}) }},
		{"finder_eleclerc_stores", func() (*Response, error) {
			return client.FinderEleclercStores(ctx, &FinderEleclercStoresParams{Query: "paris"})
		}},
		{"finder_google_countries", func() (*Response, error) {
			return client.FinderGoogleCountries(ctx, &FinderGoogleCountriesParams{Query: "fr"})
		}},
		{"finder_google_languages", func() (*Response, error) {
			return client.FinderGoogleLanguages(ctx, &FinderGoogleLanguagesParams{Query: "french"})
		}},
		{"finder_google_locations", func() (*Response, error) {
			return client.FinderGoogleLocations(ctx, &FinderGoogleLocationsParams{Query: "Paris"})
		}},
		{"finder_leroymerlin_products", func() (*Response, error) {
			return client.FinderLeroymerlinProducts(ctx, &FinderLeroymerlinProductsParams{})
		}},
		{"finder_leroymerlin_stores", func() (*Response, error) {
			return client.FinderLeroymerlinStores(ctx, &FinderLeroymerlinStoresParams{Query: "balma"})
		}},
		{"finder_linkedin_industries", func() (*Response, error) {
			return client.FinderLinkedinIndustries(ctx, &FinderLinkedinIndustriesParams{Query: "software development"})
		}},
		{"finder_mrbricolage_stores", func() (*Response, error) {
			return client.FinderMrbricolageStores(ctx, &FinderMrbricolageStoresParams{Query: "matignon"})
		}},
		{"fiverr_profile_info", func() (*Response, error) {
			return client.FiverrProfileInfo(ctx, &FiverrProfileInfoParams{Query: "https://www.fiverr.com/username"})
		}},
		{"fiverr_search", func() (*Response, error) {
			return client.FiverrSearch(ctx, &FiverrSearchParams{Query: "https://www.fiverr.com/search/gigs?query=website+development"})
		}},
		{"fnac_product", func() (*Response, error) { return client.FnacProduct(ctx, &FnacProductParams{Query: "23428278"}) }},
		{"fnac_search", func() (*Response, error) { return client.FnacSearch(ctx, &FnacSearchParams{Query: "iphone"}) }},
		{"fotocasa_ad", func() (*Response, error) {
			return client.FotocasaAd(ctx, &FotocasaAdParams{Query: "https://www.fotocasa.es/es/comprar/vivienda/madrid-capital/aire-acondicionado-calefaccion-terraza-ascensor-amueblado/190589697/d"})
		}},
		{"fotocasa_search", func() (*Response, error) { return client.FotocasaSearch(ctx, &FotocasaSearchParams{Query: "madrid"}) }},
		{"franchisedirect_franchise", func() (*Response, error) {
			return client.FranchisedirectFranchise(ctx, &FranchisedirectFranchiseParams{Query: "https://www.franchisedirect.com/coffee-franchises/coffeenow"})
		}},
		{"franchisedirect_search", func() (*Response, error) {
			return client.FranchisedirectSearch(ctx, &FranchisedirectSearchParams{Query: "coffee"})
		}},
		{"fressnapf_product", func() (*Response, error) {
			return client.FressnapfProduct(ctx, &FressnapfProductParams{Query: "https://www.fressnapf.de/p/royal-canin-maxi-adult-15-kg-1003112004/"})
		}},
		{"fressnapf_search", func() (*Response, error) {
			return client.FressnapfSearch(ctx, &FressnapfSearchParams{Query: "halsband"})
		}},
		{"funda_property", func() (*Response, error) {
			return client.FundaProperty(ctx, &FundaPropertyParams{Query: "https://www.funda.nl/detail/koop/amsterdam/appartement-hillegomstraat-57-1/44654106/"})
		}},
		{"funda_search", func() (*Response, error) { return client.FundaSearch(ctx, &FundaSearchParams{Query: "amsterdam"}) }},
		{"g2_product_info", func() (*Response, error) { return client.G2ProductInfo(ctx, &G2ProductInfoParams{Query: "postman"}) }},
		{"github_user_info", func() (*Response, error) {
			return client.GithubUserInfo(ctx, &GithubUserInfoParams{Query: "https://github.com/example-user"})
		}},
		{"glassdoor_company_info", func() (*Response, error) {
			return client.GlassdoorCompanyInfo(ctx, &GlassdoorCompanyInfoParams{Query: "1692882"})
		}},
		{"glassdoor_company_search", func() (*Response, error) {
			return client.GlassdoorCompanySearch(ctx, &GlassdoorCompanySearchParams{Query: "airbus"})
		}},
		{"glassdoor_job_info", func() (*Response, error) {
			return client.GlassdoorJobInfo(ctx, &GlassdoorJobInfoParams{Query: "1010283388161"})
		}},
		{"glassdoor_job_search", func() (*Response, error) {
			return client.GlassdoorJobSearch(ctx, &GlassdoorJobSearchParams{Query: "security"})
		}},
		{"goodreads_book", func() (*Response, error) { return client.GoodreadsBook(ctx, &GoodreadsBookParams{Query: "44767458"}) }},
		{"goodreads_search", func() (*Response, error) { return client.GoodreadsSearch(ctx, &GoodreadsSearchParams{Query: "dune"}) }},
		{"google_images", func() (*Response, error) {
			return client.GoogleImages(ctx, &GoogleImagesParams{Query: "mountain landscape"})
		}},
		{"google_jobs", func() (*Response, error) {
			return client.GoogleJobs(ctx, &GoogleJobsParams{Query: "software engineer"})
		}},
		{"google_maps", func() (*Response, error) {
			return client.GoogleMaps(ctx, &GoogleMapsParams{Query: "coffee shop paris"})
		}},
		{"google_news", func() (*Response, error) {
			return client.GoogleNews(ctx, &GoogleNewsParams{Query: "artificial intelligence news"})
		}},
		{"google_search", func() (*Response, error) {
			return client.GoogleSearch(ctx, &GoogleSearchParams{Query: "best electric cars 2024"})
		}},
		{"google_search_ai", func() (*Response, error) {
			return client.GoogleSearchAI(ctx, &GoogleSearchAIParams{Query: "what is mozilla firefox"})
		}},
		{"google_search_autocomplete", func() (*Response, error) {
			return client.GoogleSearchAutocomplete(ctx, &GoogleSearchAutocompleteParams{Query: "coffee"})
		}},
		{"google_shopping", func() (*Response, error) {
			return client.GoogleShopping(ctx, &GoogleShoppingParams{Query: "playstation 5"})
		}},
		{"google_videos", func() (*Response, error) {
			return client.GoogleVideos(ctx, &GoogleVideosParams{Query: "paris travel guide"})
		}},
		{"googleadstransparency_ad", func() (*Response, error) {
			return client.GoogleadstransparencyAd(ctx, &GoogleadstransparencyAdParams{Query: "https://adstransparency.google.com/advertiser/AR10904289436120383489/creative/CR10927354991047868417?region=FR"})
		}},
		{"googleadstransparency_search", func() (*Response, error) {
			return client.GoogleadstransparencySearch(ctx, &GoogleadstransparencySearchParams{Query: "webapi.group"})
		}},
		{"grainger_product", func() (*Response, error) { return client.GraingerProduct(ctx, &GraingerProductParams{Query: "3FRU1"}) }},
		{"grainger_search", func() (*Response, error) { return client.GraingerSearch(ctx, &GraingerSearchParams{Query: "carolina"}) }},
		{"greenhouse_job", func() (*Response, error) {
			return client.GreenhouseJob(ctx, &GreenhouseJobParams{Query: "discord/8806482002 or https://job-boards.greenhouse.io/discord/jobs/8806482002."})
		}},
		{"greenhouse_search", func() (*Response, error) {
			return client.GreenhouseSearch(ctx, &GreenhouseSearchParams{Query: "discord. A job URL returns 400."})
		}},
		{"gsaelibrary_contract", func() (*Response, error) {
			return client.GsaelibraryContract(ctx, &GsaelibraryContractParams{Query: "https://www.gsaelibrary.gsa.gov/ElibMain/contractorInfo.do?contractNumber=GS-07F-0239T&contractorName=US+RADAR+INC&executeQuery=YES"})
		}},
		{"gsaelibrary_search", func() (*Response, error) {
			return client.GsaelibrarySearch(ctx, &GsaelibrarySearchParams{Query: "radar"})
		}},
		{"homestra_ad", func() (*Response, error) {
			return client.HomestraAd(ctx, &HomestraAdParams{Query: "https://homestra.com/property/charming-4-bedroom-stone-house-with-historic-towers-in-idyllic-french-village/"})
		}},
		{"homestra_search", func() (*Response, error) {
			return client.HomestraSearch(ctx, &HomestraSearchParams{Query: "https://homestra.com/list/houses-for-sale/?maximum-price=900000&minimum-price=75000&page=1"})
		}},
		{"huggingface_model", func() (*Response, error) {
			return client.HuggingfaceModel(ctx, &HuggingfaceModelParams{Query: "google-bert/bert-base-uncased"})
		}},
		{"huggingface_search", func() (*Response, error) {
			return client.HuggingfaceSearch(ctx, &HuggingfaceSearchParams{Query: "bert"})
		}},
		{"idealista_property", func() (*Response, error) {
			return client.IdealistaProperty(ctx, &IdealistaPropertyParams{Query: "https://www.idealista.com/fr/inmueble/111463799/"})
		}},
		{"idealista_provinces", func() (*Response, error) { return client.IdealistaProvinces(ctx, &IdealistaProvincesParams{}) }},
		{"idealista_search", func() (*Response, error) {
			return client.IdealistaSearch(ctx, &IdealistaSearchParams{Query: "https://www.idealista.com/venta-viviendas/barcelona/ciutat-vella/la-barceloneta/"})
		}},
		{"idealo_product", func() (*Response, error) {
			return client.IdealoProduct(ctx, &IdealoProductParams{Query: "https://www.idealo.es/precios/205227870/asics-novablast-5-1011b974.html"})
		}},
		{"idealo_search", func() (*Response, error) { return client.IdealoSearch(ctx, &IdealoSearchParams{Query: "maison"}) }},
		{"ikea_product", func() (*Response, error) {
			return client.IkeaProduct(ctx, &IkeaProductParams{Query: "https://www.ikea.com/fr/fr/p/gladhoejden-bureau-assis-debout-blanc-30493939/"})
		}},
		{"ikea_search", func() (*Response, error) {
			return client.IkeaSearch(ctx, &IkeaSearchParams{Query: "https://www.ikea.com/fr/fr/search/?q=meuble"})
		}},
		{"immobilienscout24_property", func() (*Response, error) {
			return client.Immobilienscout24Property(ctx, &Immobilienscout24PropertyParams{Query: "168052908"})
		}},
		{"immobilienscout24_search", func() (*Response, error) {
			return client.Immobilienscout24Search(ctx, &Immobilienscout24SearchParams{Query: "berlin/berlin"})
		}},
		{"immoscout24ch_search", func() (*Response, error) {
			return client.Immoscout24chSearch(ctx, &Immoscout24chSearchParams{Query: "zurich"})
		}},
		{"immowelt_property", func() (*Response, error) {
			return client.ImmoweltProperty(ctx, &ImmoweltPropertyParams{Query: "https://www.immowelt.de/expose/65de0309-bf41-48ce-86fd-6a3441fd0487"})
		}},
		{"immowelt_search", func() (*Response, error) {
			return client.ImmoweltSearch(ctx, &ImmoweltSearchParams{Query: "https://www.immowelt.de/suche/mieten/wohnung/deutschland/ad02de1"})
		}},
		{"imoveisweb_property", func() (*Response, error) {
			return client.ImoveiswebProperty(ctx, &ImoveiswebPropertyParams{Query: "https://www.imoveisweb.com.br/imoveis/casas/3151/"})
		}},
		{"imoveisweb_search", func() (*Response, error) {
			return client.ImoveiswebSearch(ctx, &ImoveiswebSearchParams{Query: "casas"})
		}},
		{"indeed_company_info", func() (*Response, error) {
			return client.IndeedCompanyInfo(ctx, &IndeedCompanyInfoParams{Query: "Proman-1"})
		}},
		{"indeed_company_search", func() (*Response, error) {
			return client.IndeedCompanySearch(ctx, &IndeedCompanySearchParams{Query: "airbus"})
		}},
		{"indeed_job_info", func() (*Response, error) {
			return client.IndeedJobInfo(ctx, &IndeedJobInfoParams{Query: "585181841cb618e0"})
		}},
		{"indeed_job_salary", func() (*Response, error) {
			return client.IndeedJobSalary(ctx, &IndeedJobSalaryParams{Query: "ingénieur-etudes-et-développement"})
		}},
		{"indeed_job_search", func() (*Response, error) {
			return client.IndeedJobSearch(ctx, &IndeedJobSearchParams{Query: "developpeur python"})
		}},
		{"inet_product", func() (*Response, error) {
			return client.InetProduct(ctx, &InetProductParams{Query: "4306287. A search link returns 400."})
		}},
		{"inet_search", func() (*Response, error) { return client.InetSearch(ctx, &InetSearchParams{Query: "test"}) }},
		{"instagram_post_info", func() (*Response, error) {
			return client.InstagramPostInfo(ctx, &InstagramPostInfoParams{Query: "https://www.instagram.com/p/CshTQQlMGMl"})
		}},
		{"instagram_user_info", func() (*Response, error) {
			return client.InstagramUserInfo(ctx, &InstagramUserInfoParams{Query: "https://www.instagram.com/example-user/"})
		}},
		{"intermarche_product", func() (*Response, error) {
			return client.IntermarcheProduct(ctx, &IntermarcheProductParams{Query: "https://www.intermarche.com/produit/flan-patissier-cuit-&-dore-au-four/3250391522828", ShopID: "11301"})
		}},
		{"intermarche_search", func() (*Response, error) {
			return client.IntermarcheSearch(ctx, &IntermarcheSearchParams{Query: "lait", ShopID: "11301"})
		}},
		{"intersport_product", func() (*Response, error) {
			return client.IntersportProduct(ctx, &IntersportProductParams{Query: "https://www.intersport.fr/blancetvert-chaussures_de_running_homme_pegasus_plus2-nike-p-IM2541~VZE/"})
		}},
		{"intersport_search", func() (*Response, error) { return client.IntersportSearch(ctx, &IntersportSearchParams{Query: "nike"}) }},
		{"jamesedition_property", func() (*Response, error) {
			return client.JameseditionProperty(ctx, &JameseditionPropertyParams{Query: "test"})
		}},
		{"jamesedition_search", func() (*Response, error) {
			return client.JameseditionSearch(ctx, &JameseditionSearchParams{Query: "villa"})
		}},
		{"jobs", func() (*Response, error) { return client.Jobs(ctx, &JobsParams{ID: "test"}) }},
		{"kick_user_info", func() (*Response, error) {
			return client.KickUserInfo(ctx, &KickUserInfoParams{Query: "https://kick.com/example-user"})
		}},
		{"kijiji_ad", func() (*Response, error) {
			return client.KijijiAd(ctx, &KijijiAdParams{Query: "https://www.kijiji.ca/v-cell-phone-accessories/city-of-toronto/nateroso-iphone-16-case-cute-bear-honey-design/1745477275"})
		}},
		{"kijiji_search", func() (*Response, error) {
			return client.KijijiSearch(ctx, &KijijiSearchParams{Query: "https://www.kijiji.ca/b-bikes/city-of-toronto/c644l1700273"})
		}},
		{"kroger_product", func() (*Response, error) {
			return client.KrogerProduct(ctx, &KrogerProductParams{Query: "0001111040101"})
		}},
		{"kroger_search", func() (*Response, error) { return client.KrogerSearch(ctx, &KrogerSearchParams{Query: "milk"}) }},
		{"lamudi_ad", func() (*Response, error) {
			return client.LamudiAd(ctx, &LamudiAdParams{Query: "https://www.lamudi.com.ph/property/41032-73-626c638217d1-1d2a-1985488-907f-7802"})
		}},
		{"lamudi_search", func() (*Response, error) { return client.LamudiSearch(ctx, &LamudiSearchParams{}) }},
		{"leboncoin_ad", func() (*Response, error) { return client.LeboncoinAd(ctx, &LeboncoinAdParams{Query: "3012186547"}) }},
		{"leboncoin_search", func() (*Response, error) {
			return client.LeboncoinSearch(ctx, &LeboncoinSearchParams{Query: "https://www.leboncoin.fr/c/offres_d_emploi?locations=Paris_75015__48.84105_2.29926_2747"})
		}},
		{"leboncoin_search_api", func() (*Response, error) { return client.LeboncoinSearchAPI(ctx, &LeboncoinSearchAPIParams{}) }},
		{"lenovo_product", func() (*Response, error) {
			return client.LenovoProduct(ctx, &LenovoProductParams{Query: "https://www.lenovo.com/fr/fr/p/workstations/thinkstationp/thinkstation-p5-gen-2-intel-tower-workstation/len102s0024"})
		}},
		{"lenovo_search", func() (*Response, error) { return client.LenovoSearch(ctx, &LenovoSearchParams{Query: "thinkpad"}) }},
		{"leroymerlin_product", func() (*Response, error) {
			return client.LeroymerlinProduct(ctx, &LeroymerlinProductParams{Query: "https://www.leroymerlin.fr/produits/bobine-soudure-63378273.html"})
		}},
		{"leroymerlin_search", func() (*Response, error) {
			return client.LeroymerlinSearch(ctx, &LeroymerlinSearchParams{Query: "perceuse"})
		}},
		{"lidl_product", func() (*Response, error) { return client.LidlProduct(ctx, &LidlProductParams{Query: "100397447"}) }},
		{"lidl_search", func() (*Response, error) { return client.LidlSearch(ctx, &LidlSearchParams{Query: "perceuse"}) }},
		{"linkedin_company_info", func() (*Response, error) { return client.LinkedinCompanyInfo(ctx, &LinkedinCompanyInfoParams{}) }},
		{"linkedin_job_count", func() (*Response, error) { return client.LinkedinJobCount(ctx, &LinkedinJobCountParams{}) }},
		{"linkedin_job_info", func() (*Response, error) {
			return client.LinkedinJobInfo(ctx, &LinkedinJobInfoParams{Query: "https://www.linkedin.com/jobs/view/4390885746"})
		}},
		{"linkedin_job_search", func() (*Response, error) { return client.LinkedinJobSearch(ctx, &LinkedinJobSearchParams{}) }},
		{"linkedin_job_suggest", func() (*Response, error) {
			return client.LinkedinJobSuggest(ctx, &LinkedinJobSuggestParams{Query: "Air"})
		}},
		{"linkedin_post_info", func() (*Response, error) {
			return client.LinkedinPostInfo(ctx, &LinkedinPostInfoParams{Query: "https://www.linkedin.com/posts/tanyalindsay_we-have-been-building-a-web3-gaming-platform-activity-7027826235833520128-m0GW"})
		}},
		{"linkedin_product_info", func() (*Response, error) {
			return client.LinkedinProductInfo(ctx, &LinkedinProductInfoParams{Query: "https://www.linkedin.com/products/15five"})
		}},
		{"linkedin_profile_info", func() (*Response, error) {
			return client.LinkedinProfileInfo(ctx, &LinkedinProfileInfoParams{Query: "https://www.linkedin.com/in/example-user"})
		}},
		{"linkedinadlibrary_ad", func() (*Response, error) {
			return client.LinkedinadlibraryAd(ctx, &LinkedinadlibraryAdParams{Query: "https://www.linkedin.com/ad-library/detail/1585536544"})
		}},
		{"linkedinadlibrary_search", func() (*Response, error) {
			return client.LinkedinadlibrarySearch(ctx, &LinkedinadlibrarySearchParams{})
		}},
		{"loopnet_property", func() (*Response, error) { return client.LoopnetProperty(ctx, &LoopnetPropertyParams{Query: "test"}) }},
		{"loopnet_search", func() (*Response, error) {
			return client.LoopnetSearch(ctx, &LoopnetSearchParams{Query: "New York, NY"})
		}},
		{"maersk_tracking", func() (*Response, error) {
			return client.MaerskTracking(ctx, &MaerskTrackingParams{Query: "MMAU1155207"})
		}},
		{"manomano_product", func() (*Response, error) {
			return client.ManomanoProduct(ctx, &ManomanoProductParams{Query: "https://www.manomano.fr/p/perceuse-visseuse-lxt-makita-18-v-3-ah-84523637?product_id=158588345"})
		}},
		{"manomano_product_offer", func() (*Response, error) {
			return client.ManomanoProductOffer(ctx, &ManomanoProductOfferParams{Query: "158588345"})
		}},
		{"manomano_search", func() (*Response, error) {
			return client.ManomanoSearch(ctx, &ManomanoSearchParams{Query: "https://www.manomano.fr/perceuse-1146"})
		}},
		{"mascus_ad", func() (*Response, error) {
			return client.MascusAd(ctx, &MascusAdParams{Query: "https://www.mascus.fr/travaux-publics/broyeur-dechets/doppstadt-ak-435-profi/hocmgmfi.html"})
		}},
		{"mascus_search", func() (*Response, error) {
			return client.MascusSearch(ctx, &MascusSearchParams{Query: "https://www.mascus.be/bouw/graafmachines/be,country.html"})
		}},
		{"mediamarkt_product", func() (*Response, error) {
			return client.MediamarktProduct(ctx, &MediamarktProductParams{Query: "https://www.mediamarkt.ch/fr/product/_apple-iphone-15-2203371.html"})
		}},
		{"mediamarkt_search", func() (*Response, error) {
			return client.MediamarktSearch(ctx, &MediamarktSearchParams{Query: "iphone"})
		}},
		{"micromania_product", func() (*Response, error) {
			return client.MicromaniaProduct(ctx, &MicromaniaProductParams{Query: "162114"})
		}},
		{"micromania_search", func() (*Response, error) {
			return client.MicromaniaSearch(ctx, &MicromaniaSearchParams{Query: "zelda"})
		}},
		{"mobile_ad", func() (*Response, error) { return client.MobileAd(ctx, &MobileAdParams{Query: "455111824"}) }},
		{"mobile_search", func() (*Response, error) { return client.MobileSearch(ctx, &MobileSearchParams{Query: "20100"}) }},
		{"morele_product", func() (*Response, error) {
			return client.MoreleProduct(ctx, &MoreleProductParams{Query: "https://www.morele.net/laptop-lenovo-thinkbook-15-iil-20sm000fpb-6721152/"})
		}},
		{"morele_search", func() (*Response, error) { return client.MoreleSearch(ctx, &MoreleSearchParams{Query: "laptop"}) }},
		{"mrbricolage_suggest", func() (*Response, error) {
			return client.MrbricolageSuggest(ctx, &MrbricolageSuggestParams{Query: "marteau"})
		}},
		{"newegg_product", func() (*Response, error) {
			return client.NeweggProduct(ctx, &NeweggProductParams{Query: "N82E16823126188"})
		}},
		{"newegg_search", func() (*Response, error) {
			return client.NeweggSearch(ctx, &NeweggSearchParams{Query: "logitech keyboard"})
		}},
		{"northdata_company", func() (*Response, error) {
			return client.NorthdataCompany(ctx, &NorthdataCompanyParams{Query: "Maison Boinaud"})
		}},
		{"northdata_search", func() (*Response, error) { return client.NorthdataSearch(ctx, &NorthdataSearchParams{Query: "BMW"}) }},
		{"npi_provider", func() (*Response, error) { return client.NPIProvider(ctx, &NPIProviderParams{Query: "test"}) }},
		{"npi_search", func() (*Response, error) { return client.NPISearch(ctx, &NPISearchParams{}) }},
		{"oceanio_search", func() (*Response, error) { return client.OceanioSearch(ctx, &OceanioSearchParams{Query: "google.com"}) }},
		{"olx_ad", func() (*Response, error) {
			return client.OlxAd(ctx, &OlxAdParams{Query: "https://www.olx.pl/d/oferta/apple-iphone-17-pro-256-gb-nowy-szczecin-zaplombowany-karta-sim-ul-baluki-4-szczecin-CID99-ID18Acsd.html"})
		}},
		{"olx_product", func() (*Response, error) {
			return client.OlxProduct(ctx, &OlxProductParams{Query: "https://www.olx.com.br/eletronicos-e-celulares/celulares-e-smartphones/iphone-14-128gb-como-novo-IDJ4mBb.html"})
		}},
		{"olx_search", func() (*Response, error) { return client.OlxSearch(ctx, &OlxSearchParams{Query: "iphone 14"}) }},
		{"oscaro_product", func() (*Response, error) {
			return client.OscaroProduct(ctx, &OscaroProductParams{Query: "https://www.oscaro.com/huile-moteur-motul-102208-15213891-1862-p"})
		}},
		{"oscaro_search", func() (*Response, error) { return client.OscaroSearch(ctx, &OscaroSearchParams{Query: "moteur ai"}) }},
		{"ovoko_ad", func() (*Response, error) {
			return client.OvokoAd(ctx, &OvokoAdParams{Query: "https://ovoko.fr/les-pieces/lgn634714-1780a036-mitsubishi-asx-filtre-d-habitacle"})
		}},
		{"ovoko_search", func() (*Response, error) {
			return client.OvokoSearch(ctx, &OvokoSearchParams{Query: "https://ovoko.pl/szukaj?q=Porsche+Macan+uszczelka+klapy+bagaznika&exact=1"})
		}},
		{"owler_company_info", func() (*Response, error) {
			return client.OwlerCompanyInfo(ctx, &OwlerCompanyInfoParams{Query: "https://owler.com/company/tesla"})
		}},
		{"owler_search", func() (*Response, error) { return client.OwlerSearch(ctx, &OwlerSearchParams{Query: "tesla"}) }},
		{"pagesjaunes_page_info", func() (*Response, error) {
			return client.PagesjaunesPageInfo(ctx, &PagesjaunesPageInfoParams{Query: "https://www.pagesjaunes.fr/pros/56824689"})
		}},
		{"pagesjaunes_search", func() (*Response, error) {
			return client.PagesjaunesSearch(ctx, &PagesjaunesSearchParams{Query: "https://www.pagesjaunes.fr/annuaire/chercherlespros?quoiqui=Dentiste&ou=Toulouse+(31000)&univers=pagesjaunes&idOu="})
		}},
		{"patreon_profile_info", func() (*Response, error) {
			return client.PatreonProfileInfo(ctx, &PatreonProfileInfoParams{Query: "https://www.patreon.com/example-user"})
		}},
		{"patreon_search", func() (*Response, error) { return client.PatreonSearch(ctx, &PatreonSearchParams{Query: "creator"}) }},
		{"petsmart_product", func() (*Response, error) {
			return client.PetsmartProduct(ctx, &PetsmartProductParams{Query: "https://www.petsmart.com/dog/food/dry-food/purina-pro-plan-sensitive-skin-and-stomach-dry-dog-food-adult-salmon-and-rice-formula-digestive-health-36648.html"})
		}},
		{"petsmart_search", func() (*Response, error) { return client.PetsmartSearch(ctx, &PetsmartSearchParams{Query: "dog food"}) }},
		{"picard_product", func() (*Response, error) { return client.PicardProduct(ctx, &PicardProductParams{Query: "050314"}) }},
		{"picard_search", func() (*Response, error) { return client.PicardSearch(ctx, &PicardSearchParams{Query: "patate douce"}) }},
		{"pinterest_user_info", func() (*Response, error) {
			return client.PinterestUserInfo(ctx, &PinterestUserInfoParams{Query: "https://www.pinterest.com/username"})
		}},
		{"pitchbook_company_info", func() (*Response, error) {
			return client.PitchbookCompanyInfo(ctx, &PitchbookCompanyInfoParams{Query: "11062-99"})
		}},
		{"pitchbook_search", func() (*Response, error) { return client.PitchbookSearch(ctx, &PitchbookSearchParams{Query: "airbus"}) }},
		{"pointp_product", func() (*Response, error) {
			return client.PointpProduct(ctx, &PointpProductParams{Query: "https://www.pointp.fr/p/platre-isolation-ite/laine-de-verre-a-souffler-thermo-loft-knauf-insulation-sac-16-A3334230"})
		}},
		{"pointp_search", func() (*Response, error) { return client.PointpSearch(ctx, &PointpSearchParams{Query: "colle"}) }},
		{"prnewswire_article", func() (*Response, error) {
			return client.PrnewswireArticle(ctx, &PrnewswireArticleParams{Query: "test"})
		}},
		{"prnewswire_search", func() (*Response, error) {
			return client.PrnewswireSearch(ctx, &PrnewswireSearchParams{Query: "apple"})
		}},
		{"producthunt_product_info", func() (*Response, error) {
			return client.ProducthuntProductInfo(ctx, &ProducthuntProductInfoParams{Query: "https://www.producthunt.com/products/miro"})
		}},
		{"proshop_product", func() (*Response, error) { return client.ProshopProduct(ctx, &ProshopProductParams{Query: "3071602"}) }},
		{"proshop_search", func() (*Response, error) { return client.ProshopSearch(ctx, &ProshopSearchParams{Query: "ssd"}) }},
		{"rappi_product", func() (*Response, error) { return client.RappiProduct(ctx, &RappiProductParams{Query: "test"}) }},
		{"rappi_search", func() (*Response, error) { return client.RappiSearch(ctx, &RappiSearchParams{Query: "leche"}) }},
		{"ratemyprofessors_professor", func() (*Response, error) {
			return client.RatemyprofessorsProfessor(ctx, &RatemyprofessorsProfessorParams{Query: "1064107"})
		}},
		{"ratemyprofessors_search", func() (*Response, error) {
			return client.RatemyprofessorsSearch(ctx, &RatemyprofessorsSearchParams{Query: "smith"})
		}},
		{"realestatecomau_search", func() (*Response, error) {
			return client.RealestatecomauSearch(ctx, &RealestatecomauSearchParams{Query: "Sydney, NSW"})
		}},
		{"realtor_property", func() (*Response, error) {
			return client.RealtorProperty(ctx, &RealtorPropertyParams{Query: "https://www.realtor.com/realestateandhomes-detail/1492-E-57th-St_Brooklyn_NY_11234_M46117-61190"})
		}},
		{"realtor_search", func() (*Response, error) {
			return client.RealtorSearch(ctx, &RealtorSearchParams{Query: "https://www.realtor.com/realestateandhomes-search/New-York_NY"})
		}},
		{"redfin_property", func() (*Response, error) {
			return client.RedfinProperty(ctx, &RedfinPropertyParams{Query: "https://www.redfin.com/WA/Seattle/1120-8th-Ave-98101/unit-1603/home/14602?region_id=16163&region_type=6"})
		}},
		{"redfin_search", func() (*Response, error) {
			return client.RedfinSearch(ctx, &RedfinSearchParams{Query: "https://www.redfin.com/city/16163/WA/Seattle"})
		}},
		{"reichelt_product", func() (*Response, error) { return client.ReicheltProduct(ctx, &ReicheltProductParams{Query: "test"}) }},
		{"reichelt_search", func() (*Response, error) { return client.ReicheltSearch(ctx, &ReicheltSearchParams{Query: "test"}) }},
		{"reverb_product", func() (*Response, error) { return client.ReverbProduct(ctx, &ReverbProductParams{Query: "96933467"}) }},
		{"reverb_search", func() (*Response, error) { return client.ReverbSearch(ctx, &ReverbSearchParams{Query: "guitar"}) }},
		{"rexel_product", func() (*Response, error) {
			return client.RexelProduct(ctx, &RexelProductParams{Query: "https://www.rexel.fr/frx/HAGWXF122/p/71742449"})
		}},
		{"rexel_search", func() (*Response, error) { return client.RexelSearch(ctx, &RexelSearchParams{Query: "prise"}) }},
		{"rightmove_property", func() (*Response, error) {
			return client.RightmoveProperty(ctx, &RightmovePropertyParams{Query: "https://www.rightmove.co.uk/properties/172438235"})
		}},
		{"rightmove_search", func() (*Response, error) {
			return client.RightmoveSearch(ctx, &RightmoveSearchParams{Query: "Bristol"})
		}},
		{"rsonline_product", func() (*Response, error) {
			return client.RsonlineProduct(ctx, &RsonlineProductParams{Query: "https://uk.rs-online.com/web/p/splice-connectors/0303327"})
		}},
		{"rsonline_search", func() (*Response, error) {
			return client.RsonlineSearch(ctx, &RsonlineSearchParams{Query: "connector"})
		}},
		{"seek_job", func() (*Response, error) { return client.SeekJob(ctx, &SeekJobParams{Query: "test"}) }},
		{"seek_search", func() (*Response, error) { return client.SeekSearch(ctx, &SeekSearchParams{Query: "nurse"}) }},
		{"seloger_agency", func() (*Response, error) {
			return client.SelogerAgency(ctx, &SelogerAgencyParams{Query: "https://www.seloger.com/professionnels/agences-immobilieres/castanet-tolosan-31320/agence-21719"})
		}},
		{"seloger_agency_rent", func() (*Response, error) {
			return client.SelogerAgencyRent(ctx, &SelogerAgencyRentParams{Query: "18317"})
		}},
		{"seloger_agency_sale", func() (*Response, error) {
			return client.SelogerAgencySale(ctx, &SelogerAgencySaleParams{Query: "18317"})
		}},
		{"seloger_property", func() (*Response, error) {
			return client.SelogerProperty(ctx, &SelogerPropertyParams{Query: "https://www.seloger.com/annonces/achat/appartement/bordeaux-33/capucins-saint-michel-nansouty-saint-genes/234462129.htm"})
		}},
		{"seloger_search", func() (*Response, error) {
			return client.SelogerSearch(ctx, &SelogerSearchParams{Query: "https://www.seloger.com/annonces/achat/appartement/bordeaux-33/"})
		}},
		{"sephora_product", func() (*Response, error) {
			return client.SephoraProduct(ctx, &SephoraProductParams{Query: "https://www.sephora.fr/p/midnight-espresso---eau-de-parfum-au-cafe-831955.html"})
		}},
		{"sephora_search", func() (*Response, error) { return client.SephoraSearch(ctx, &SephoraSearchParams{Query: "mascara"}) }},
		{"shophouzz_product", func() (*Response, error) {
			return client.ShophouzzProduct(ctx, &ShophouzzProductParams{Query: "https://shophouzz.com/products/burke-counter-stool-cognac-vegan-leather-white-oak-veneer-metal-dowel-prvw-vr-186041078"})
		}},
		{"shopify_apps", func() (*Response, error) {
			return client.ShopifyApps(ctx, &ShopifyAppsParams{Query: "https://apps.shopify.com/judgeme"})
		}},
		{"shopify_catalog", func() (*Response, error) {
			return client.ShopifyCatalog(ctx, &ShopifyCatalogParams{Query: "https://www.decathlon.com"})
		}},
		{"shopify_collections", func() (*Response, error) {
			return client.ShopifyCollections(ctx, &ShopifyCollectionsParams{Query: "https://www.decathlon.com"})
		}},
		{"shopify_product", func() (*Response, error) {
			return client.ShopifyProduct(ctx, &ShopifyProductParams{Query: "https://www.decathlon.com/products/simond-backpacking-sleeping-bag-mt500-41-f-synthetic-346446"})
		}},
		{"similarweb_domain", func() (*Response, error) {
			return client.SimilarwebDomain(ctx, &SimilarwebDomainParams{Query: "google.com"})
		}},
		{"similarweb_search", func() (*Response, error) {
			return client.SimilarwebSearch(ctx, &SimilarwebSearchParams{Query: "google.com"})
		}},
		{"snapeda_product", func() (*Response, error) {
			return client.SnapedaProduct(ctx, &SnapedaProductParams{Query: "CYPD3175-24LQXQ/Cypress Semiconductor"})
		}},
		{"snapeda_search", func() (*Response, error) { return client.SnapedaSearch(ctx, &SnapedaSearchParams{Query: "USB Type C"}) }},
		{"starbucks_menu", func() (*Response, error) { return client.StarbucksMenu(ctx, &StarbucksMenuParams{}) }},
		{"starbucks_product", func() (*Response, error) {
			return client.StarbucksProduct(ctx, &StarbucksProductParams{Query: "418/hot"})
		}},
		{"starbucks_search", func() (*Response, error) { return client.StarbucksSearch(ctx, &StarbucksSearchParams{}) }},
		{"stats", func() (*Response, error) { return client.Stats(ctx, &StatsParams{}) }},
		{"stockx_product", func() (*Response, error) {
			return client.StockxProduct(ctx, &StockxProductParams{Query: "nike-air-max-1-travis-scott-wheat"})
		}},
		{"stockx_search", func() (*Response, error) {
			return client.StockxSearch(ctx, &StockxSearchParams{Query: "Air Jordan Retro"})
		}},
		{"stockx_trends", func() (*Response, error) { return client.StockxTrends(ctx, &StockxTrendsParams{Query: "sneakers"}) }},
		{"stubhub_event", func() (*Response, error) { return client.StubhubEvent(ctx, &StubhubEventParams{Query: "test"}) }},
		{"stubhub_search", func() (*Response, error) { return client.StubhubSearch(ctx, &StubhubSearchParams{Query: "nba"}) }},
		{"target_product", func() (*Response, error) { return client.TargetProduct(ctx, &TargetProductParams{Query: "17457624"}) }},
		{"target_search", func() (*Response, error) { return client.TargetSearch(ctx, &TargetSearchParams{Query: "shampoo"}) }},
		{"tdlr_license", func() (*Response, error) { return client.TdlrLicense(ctx, &TdlrLicenseParams{Query: "test"}) }},
		{"tdlr_search", func() (*Response, error) { return client.TdlrSearch(ctx, &TdlrSearchParams{}) }},
		{"ted_notice", func() (*Response, error) { return client.TedNotice(ctx, &TedNoticeParams{Query: "30-2026"}) }},
		{"ted_search", func() (*Response, error) { return client.TedSearch(ctx, &TedSearchParams{}) }},
		{"thalia_product", func() (*Response, error) {
			return client.ThaliaProduct(ctx, &ThaliaProductParams{Query: "A1074644377"})
		}},
		{"thalia_search", func() (*Response, error) { return client.ThaliaSearch(ctx, &ThaliaSearchParams{Query: "Harry Potter"}) }},
		{"thefork_restaurant", func() (*Response, error) {
			return client.TheforkRestaurant(ctx, &TheforkRestaurantParams{Query: "https://www.thefork.fr/restaurant/salamboo-amilkar-r10892#date=2026-06-13&rankedBy=SEARCH_ENGINE"})
		}},
		{"thefork_search", func() (*Response, error) {
			return client.TheforkSearch(ctx, &TheforkSearchParams{Query: "https://www.thefork.fr/search?cityId=529135&date=2026-06-13&hour=660&partySize=8&timezone=Europe%2FParis"})
		}},
		{"thomasnet_company", func() (*Response, error) {
			return client.ThomasnetCompany(ctx, &ThomasnetCompanyParams{Query: "https://www.thomasnet.com/company/zepole-supply-co-20088400/profile"})
		}},
		{"thomasnet_search", func() (*Response, error) {
			return client.ThomasnetSearch(ctx, &ThomasnetSearchParams{Query: "https://www.thomasnet.com/suppliers/search?searchterm=Table+Forks"})
		}},
		{"ticketmaster_event", func() (*Response, error) {
			return client.TicketmasterEvent(ctx, &TicketmasterEventParams{Query: "test"})
		}},
		{"ticketmaster_search", func() (*Response, error) {
			return client.TicketmasterSearch(ctx, &TicketmasterSearchParams{Query: "jazz"})
		}},
		{"tiktoktopads_ad", func() (*Response, error) {
			return client.TiktoktopadsAd(ctx, &TiktoktopadsAdParams{Query: "https://ads.tiktok.com/business/creativecenter/topads/7686267188656750612/pc/en?period=30"})
		}},
		{"tiktoktopads_search", func() (*Response, error) { return client.TiktoktopadsSearch(ctx, &TiktoktopadsSearchParams{}) }},
		{"tradeaplane_listing", func() (*Response, error) {
			return client.TradeaplaneListing(ctx, &TradeaplaneListingParams{Query: "https://www.trade-a-plane.com/search?category_level1=Piston+Engine&make=ROLLS+ROYCE&model=PACKARD+V-1650&listing_id=2419556&s-type=engine"})
		}},
		{"tradeaplane_search", func() (*Response, error) {
			return client.TradeaplaneSearch(ctx, &TradeaplaneSearchParams{Query: "rol"})
		}},
		{"trustpilot_company_info", func() (*Response, error) {
			return client.TrustpilotCompanyInfo(ctx, &TrustpilotCompanyInfoParams{Query: "https://www.trustpilot.com/review/nordvpn.com"})
		}},
		{"untappd_beer", func() (*Response, error) { return client.UntappdBeer(ctx, &UntappdBeerParams{Query: "4509"}) }},
		{"untappd_search", func() (*Response, error) { return client.UntappdSearch(ctx, &UntappdSearchParams{Query: "ipa"}) }},
		{"upwork_freelancer", func() (*Response, error) {
			return client.UpworkFreelancer(ctx, &UpworkFreelancerParams{Query: "~0123456789abcdef01"})
		}},
		{"upwork_jobs", func() (*Response, error) { return client.UpworkJobs(ctx, &UpworkJobsParams{Query: "api-development"}) }},
		{"upwork_services", func() (*Response, error) {
			return client.UpworkServices(ctx, &UpworkServicesParams{Query: "mobile-apps"})
		}},
		{"upwork_services_info", func() (*Response, error) {
			return client.UpworkServicesInfo(ctx, &UpworkServicesInfoParams{Query: "https://www.upwork.com/services/product/development-it-mobile-app-developer-android-ios-flutter-application-hybrid-apps-1403673770703212544"})
		}},
		{"usage", func() (*Response, error) { return client.Usage(ctx) }},
		{"usaspending_award", func() (*Response, error) {
			return client.UsaspendingAward(ctx, &UsaspendingAwardParams{Query: "CONT_AWD_GSFC0198106DNAS526555_8000_-NONE-_-NONE-."})
		}},
		{"usaspending_search", func() (*Response, error) {
			return client.UsaspendingSearch(ctx, &UsaspendingSearchParams{Query: "software. An award URL returns 400."})
		}},
		{"vestiairecollective_product", func() (*Response, error) {
			return client.VestiairecollectiveProduct(ctx, &VestiairecollectiveProductParams{Query: "https://www.vestiairecollective.com/vetements-femme/vestes/chanel/vesteblouson-chanel-la-petite-veste-noire-en-tweed-rose-73024488.shtml"})
		}},
		{"vestiairecollective_search", func() (*Response, error) {
			return client.VestiairecollectiveSearch(ctx, &VestiairecollectiveSearchParams{Query: "chanel"})
		}},
		{"vinted_item", func() (*Response, error) {
			return client.VintedItem(ctx, &VintedItemParams{Query: "https://www.vinted.com/items/9437210920-purple-prom-dress"})
		}},
		{"vinted_search", func() (*Response, error) { return client.VintedSearch(ctx, &VintedSearchParams{Query: "nike"}) }},
		{"vinted_user_info", func() (*Response, error) {
			return client.VintedUserInfo(ctx, &VintedUserInfoParams{Query: "https://www.vinted.fr/member/123456789-username"})
		}},
		{"vivareal_property", func() (*Response, error) {
			return client.VivarealProperty(ctx, &VivarealPropertyParams{Query: "https://www.vivareal.com.br/imovel/apartamento-3-quartos-brooklin-zona-sul-sao-paulo-com-garagem-186m2-venda-RS3710000-id-2776423622/"})
		}},
		{"vivareal_search", func() (*Response, error) {
			return client.VivarealSearch(ctx, &VivarealSearchParams{Query: "São Paulo"})
		}},
		{"walmart_product", func() (*Response, error) {
			return client.WalmartProduct(ctx, &WalmartProductParams{Query: "https://www.walmart.com/ip/Fresh-Rockit-Crisp-Sweet-Miniature-Apples-3lb-Tub/309762096"})
		}},
		{"walmart_search", func() (*Response, error) {
			return client.WalmartSearch(ctx, &WalmartSearchParams{Query: "https://www.walmart.com/search?q=apples"})
		}},
		{"walmart_suggest", func() (*Response, error) { return client.WalmartSuggest(ctx, &WalmartSuggestParams{Query: "apples"}) }},
		{"website_antibot", func() (*Response, error) {
			return client.WebsiteAntibot(ctx, &WebsiteAntibotParams{Query: "https://www.example.com"})
		}},
		{"website_crawler", func() (*Response, error) {
			return client.WebsiteCrawler(ctx, &WebsiteCrawlerParams{Query: "https://example.com"})
		}},
		{"website_email_phone_extractor", func() (*Response, error) {
			return client.WebsiteEmailPhoneExtractor(ctx, &WebsiteEmailPhoneExtractorParams{Query: "https://www.hexa.cc"})
		}},
		{"website_rendering", func() (*Response, error) {
			return client.WebsiteRendering(ctx, &WebsiteRenderingParams{Query: "https://example.com"})
		}},
		{"website_screenshot", func() (*Response, error) {
			return client.WebsiteScreenshot(ctx, &WebsiteScreenshotParams{URL: "https://www.amazon.com/s?k=pokemon"})
		}},
		{"website_technology", func() (*Response, error) {
			return client.WebsiteTechnology(ctx, &WebsiteTechnologyParams{Query: "https://example.com"})
		}},
		{"website_webunlocker", func() (*Response, error) {
			return client.WebsiteWebunlocker(ctx, &WebsiteWebunlockerParams{Query: "https://www.norauto.fr/t/pneu/w-185-h-55-r-14/ete-s/80-l-h-q.html"})
		}},
		{"weibo_post_info", func() (*Response, error) { return client.WeiboPostInfo(ctx, &WeiboPostInfoParams{Query: "test"}) }},
		{"weibo_search", func() (*Response, error) { return client.WeiboSearch(ctx, &WeiboSearchParams{Query: "test"}) }},
		{"welcometothejungle_company", func() (*Response, error) {
			return client.WelcometothejungleCompany(ctx, &WelcometothejungleCompanyParams{Query: "alan"})
		}},
		{"welcometothejungle_search", func() (*Response, error) {
			return client.WelcometothejungleSearch(ctx, &WelcometothejungleSearchParams{Query: "python"})
		}},
		{"wellfound_company_info", func() (*Response, error) {
			return client.WellfoundCompanyInfo(ctx, &WellfoundCompanyInfoParams{Query: "https://wellfound.com/company/dataiku"})
		}},
		{"x_post_info", func() (*Response, error) {
			return client.XPostInfo(ctx, &XPostInfoParams{Query: "https://x.com/example-user/status/0000000000000000001"})
		}},
		{"x_profile_info", func() (*Response, error) { return client.XProfileInfo(ctx, &XProfileInfoParams{Query: "elonmusk"}) }},
		{"yahoo_search", func() (*Response, error) { return client.YahooSearch(ctx, &YahooSearchParams{Query: "test"}) }},
		{"yellowpages_business", func() (*Response, error) {
			return client.YellowpagesBusiness(ctx, &YellowpagesBusinessParams{Query: "https://www.yellowpages.com/los-angeles-ca/mip/anawalt-lumber-13945594?lid=1002108313121"})
		}},
		{"yellowpages_search", func() (*Response, error) {
			return client.YellowpagesSearch(ctx, &YellowpagesSearchParams{Query: "https://www.yellowpages.com/los-angeles-ca/plumbers"})
		}},
		{"yelp_business", func() (*Response, error) {
			return client.YelpBusiness(ctx, &YelpBusinessParams{Query: "https://www.yelp.com/biz/novela-san-francisco-2"})
		}},
		{"yelp_reviews", func() (*Response, error) {
			return client.YelpReviews(ctx, &YelpReviewsParams{Query: "cX5dOMkwZ5rqWbeIJpw_Ig"})
		}},
		{"yelp_search", func() (*Response, error) {
			return client.YelpSearch(ctx, &YelpSearchParams{Query: "https://www.yelp.com/search?find_desc=Cocktail+Bars&find_loc=San+Francisco%2C+CA"})
		}},
		{"zapimoveis_ad", func() (*Response, error) {
			return client.ZapimoveisAd(ctx, &ZapimoveisAdParams{Query: "https://www.zapimoveis.com.br/imovel/venda-apartamento-3-quartos-mobiliado-jardim-paulista-zona-sul-sao-paulo-sp-170m2-id-2912096343/"})
		}},
		{"zapimoveis_search", func() (*Response, error) {
			return client.ZapimoveisSearch(ctx, &ZapimoveisSearchParams{Query: "São Paulo"})
		}},
		{"zazzle_product", func() (*Response, error) {
			return client.ZazzleProduct(ctx, &ZazzleProductParams{Query: "https://www.zazzle.com/branded_logo_two_tone_coffee_mug-256528848440338952."})
		}},
		{"zazzle_search", func() (*Response, error) {
			return client.ZazzleSearch(ctx, &ZazzleSearchParams{Query: "coffee mug. A product URL returns 400."})
		}},
		{"zillow_property", func() (*Response, error) {
			return client.ZillowProperty(ctx, &ZillowPropertyParams{Query: "https://www.zillow.com/homedetails/20-Kyle-Ct-Staten-Island-NY-10312/32361983_zpid/"})
		}},
		{"zillow_search", func() (*Response, error) {
			return client.ZillowSearch(ctx, &ZillowSearchParams{Query: "https://www.zillow.com/staten-island-ny/"})
		}},
		{"zonaprop_property", func() (*Response, error) {
			return client.ZonapropProperty(ctx, &ZonapropPropertyParams{Query: "https://www.zonaprop.com.ar/propiedades/clasificado/veclapin-venta-departamento-2-ambientes-villa-crespo-59096735.html"})
		}},
		{"zonaprop_search", func() (*Response, error) {
			return client.ZonapropSearch(ctx, &ZonapropSearchParams{Query: "capital federal"})
		}},
		{"zoopla_property", func() (*Response, error) {
			return client.ZooplaProperty(ctx, &ZooplaPropertyParams{Query: "https://www.zoopla.co.uk/to-rent/details/73119609/"})
		}},
		{"zoopla_search", func() (*Response, error) {
			return client.ZooplaSearch(ctx, &ZooplaSearchParams{Query: "https://www.zoopla.co.uk/for-sale/property/manchester/"})
		}},
		{"zooplus_product", func() (*Response, error) { return client.ZooplusProduct(ctx, &ZooplusProductParams{Query: "705321"}) }},
		{"zooplus_search", func() (*Response, error) { return client.ZooplusSearch(ctx, &ZooplusSearchParams{Query: "croquettes"}) }},
		{"zoro_product", func() (*Response, error) {
			return client.ZoroProduct(ctx, &ZoroProductParams{Query: "https://www.zoro.com/3m-safety-glasses-securefit-200-series-scotchgard-anti-fog-frameless-blue-temples-clear-lens-sf201sgaf-blu/i/G1937045/"})
		}},
		{"zoro_search", func() (*Response, error) { return client.ZoroSearch(ctx, &ZoroSearchParams{Query: "safety glasses"}) }},
	}
	for i, check := range checks {
		if _, err := check.call(); err != nil {
			t.Fatalf("%s: %v", check.name, err)
		}
		ep := endpoints[check.name]
		request := (*calls)[i].request
		if request.URL.Path != strings.ReplaceAll(ep.path, "{id}", "test") || request.Method != ep.methods[0] {
			t.Fatalf("%s: unexpected %s %s", check.name, request.Method, request.URL.Path)
		}
	}
	if len(checks) != len(endpoints) {
		t.Fatalf("%d checks for %d endpoints", len(checks), len(endpoints))
	}
}

func TestBinaryResponse(t *testing.T) {
	client, _ := newTestClient(t, func(*http.Request) *http.Response {
		return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"image/png"}}, Body: io.NopCloser(strings.NewReader("\x89PNG"))}
	})
	response, err := client.Request(context.Background(), "GET", "/v1/anything", nil)
	if err != nil {
		t.Fatal(err)
	}
	if response.IsJSON || string(response.Body) != "\x89PNG" || response.ContentType != "image/png" {
		t.Fatalf("unexpected response %+v", response)
	}
}

func TestNonJSONSuccessForJSONOperation(t *testing.T) {
	client, _ := newTestClient(t, func(*http.Request) *http.Response {
		return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": []string{"text/html"}}, Body: io.NopCloser(strings.NewReader("<html>"))}
	})
	ctx := context.Background()
	_, err := client.AmazonProduct(ctx, &AmazonProductParams{Query: "B09DFCB66S"})
	var decodingError *ResponseDecodingError
	if !errors.As(err, &decodingError) {
		t.Fatalf("expected ResponseDecodingError, got %v", err)
	}
}

func TestAPIError(t *testing.T) {
	client, _ := newTestClient(t, func(*http.Request) *http.Response {
		return jsonResponse(401, `{"error":"Invalid API Key"}`)
	})
	_, err := client.Request(context.Background(), "GET", "/v1/usage", nil)
	apiError, ok := IsAPIError(err)
	if !ok || apiError.StatusCode != 401 || apiError.Message != "Invalid API Key" {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestTransportError(t *testing.T) {
	client, err := NewClient("k", &ClientOptions{MaxRetries: -1, HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("network down")
	})}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Request(context.Background(), "GET", "/v1/usage", nil)
	var transportError *TransportError
	if !errors.As(err, &transportError) {
		t.Fatalf("expected TransportError, got %v", err)
	}
}

func TestTimeout(t *testing.T) {
	client, err := NewClient("k", &ClientOptions{Timeout: time.Millisecond, HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		<-r.Context().Done()
		return nil, r.Context().Err()
	})}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Request(context.Background(), "GET", "/v1/usage", nil)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
}

func TestRetries(t *testing.T) {
	statuses := []int{503, 429, 200}
	var calls int
	client, err := NewClient("k", &ClientOptions{RetryDelay: time.Millisecond, HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		if calls == 1 {
			return nil, errors.New("connection reset")
		}
		status := statuses[0]
		statuses = statuses[1:]
		response := jsonResponse(status, `{}`)
		if status == 429 {
			response.Header.Set("Retry-After", "0")
		}
		return response, nil
	})}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Request(context.Background(), "POST", "/v1/usage", map[string]any{"q": "x"}); err != nil || calls != 4 {
		t.Fatalf("expected success after 4 calls, got %v after %d", err, calls)
	}
}

func TestRetryLimitsAndClientErrors(t *testing.T) {
	newClient := func(status int, options ...RequestOption) (int, error) {
		var calls int
		client, _ := NewClient("k", &ClientOptions{MaxRetries: 2, RetryDelay: time.Millisecond, HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			calls++
			return jsonResponse(status, `{"error":"x"}`), nil
		})}})
		_, err := client.Request(context.Background(), "GET", "/v1/usage", nil, options...)
		return calls, err
	}
	if calls, err := newClient(503); calls != 3 || err == nil {
		t.Fatalf("503: %d calls, %v", calls, err)
	}
	if calls, _ := newClient(402); calls != 1 {
		t.Fatalf("402 should not be retried, got %d calls", calls)
	}
	if calls, _ := newClient(500, WithMaxRetries(0)); calls != 1 {
		t.Fatalf("WithMaxRetries(0) should disable retries, got %d calls", calls)
	}
}

func TestRateLimitRetries(t *testing.T) {
	cases := []struct {
		message string
		calls   int
	}{
		{"Rate limit exceeded: You've exceeded the rate limit on your subscription", 2},
		{"Rate limit exceeded for the API key: quota daily", 1},
		{"Invalid API Key", 1},
	}
	for _, c := range cases {
		calls := 0
		client, _ := NewClient("k", &ClientOptions{RetryDelay: time.Millisecond, HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			calls++
			if calls == 1 {
				body, _ := json.Marshal(map[string]string{"error": c.message})
				return jsonResponse(401, string(body)), nil
			}
			return jsonResponse(200, `{}`), nil
		})}})
		_, _ = client.Request(context.Background(), "GET", "/v1/usage", nil)
		if calls != c.calls {
			t.Fatalf("%q: %d calls, want %d", c.message, calls, c.calls)
		}
	}
}

func TestGlobalParameters(t *testing.T) {
	client, calls := newTestClient(t, func(*http.Request) *http.Response { return jsonResponse(200, `{"job_id":"j1"}`) })
	ctx := context.Background()
	if _, err := client.AmazonProduct(ctx, &AmazonProductParams{Query: "B09DFCB66S", Fields: String("a,b.c"), IsAsync: Bool(true), SoftFail: Bool(false)}); err != nil {
		t.Fatal(err)
	}
	query := (*calls)[0].request.URL.Query()
	if query.Get("fields") != "a,b.c" || query.Get("is_async") != "true" || query.Get("soft_fail") != "false" {
		t.Fatalf("unexpected query %v", query)
	}
}

func TestBatch(t *testing.T) {
	var mu sync.Mutex
	running, peak := 0, 0
	results := Batch(context.Background(), []int{1, 2, 3, 4, 5, 6}, 2, func(_ context.Context, n int) (int, error) {
		mu.Lock()
		running++
		peak = max(peak, running)
		mu.Unlock()
		time.Sleep(5 * time.Millisecond)
		mu.Lock()
		running--
		mu.Unlock()
		if n == 4 {
			return 0, errors.New("boom")
		}
		return n * 10, nil
	})
	if peak != 2 {
		t.Fatalf("peak concurrency %d", peak)
	}
	for i, r := range results {
		if r.Index != i || (i == 3) != (r.Err != nil) || (r.Err == nil && r.Value != (i+1)*10) {
			t.Fatalf("unexpected result %d: %+v", i, r)
		}
	}
}

func TestDualOperationCanUsePost(t *testing.T) {
	client, calls := newTestClient(t, okJSON)
	ctx := context.Background()
	if _, err := client.AboutyouProduct(ctx, &AboutyouProductParams{Query: "test"}, WithMethod("POST")); err != nil {
		t.Fatal(err)
	}
	call := (*calls)[0]
	if call.request.Method != "POST" || len(call.body) == 0 || call.request.URL.RawQuery != "" {
		t.Fatalf("unexpected request %s %q", call.request.Method, call.body)
	}
}

func TestWaitForJob(t *testing.T) {
	statuses := []string{"Queued", "Running", "CompletedJob"}
	client, calls := newTestClient(t, func(*http.Request) *http.Response {
		status := statuses[0]
		statuses = statuses[1:]
		return jsonResponse(200, `{"status":"`+status+`","progress":50}`)
	})
	seen := 0
	response, err := client.WaitForJob(context.Background(), "abc", &WaitForJobOptions{Interval: time.Millisecond, OnProgress: func(map[string]any) { seen++ }})
	if err != nil {
		t.Fatal(err)
	}
	job, _ := response.Map()
	if job["status"] != "CompletedJob" || len(*calls) != 3 || seen != 3 || (*calls)[0].request.URL.Path != "/v1/jobs/abc" {
		t.Fatalf("unexpected job %v after %d calls", job, len(*calls))
	}
}

func TestWaitForJobTimesOut(t *testing.T) {
	client, _ := newTestClient(t, func(*http.Request) *http.Response { return jsonResponse(200, `{"status":"Running"}`) })
	_, err := client.WaitForJob(context.Background(), "abc", &WaitForJobOptions{Interval: 10 * time.Millisecond, Timeout: time.Millisecond})
	var timeoutError *JobTimeoutError
	if !errors.As(err, &timeoutError) {
		t.Fatalf("expected JobTimeoutError, got %v", err)
	}
}
