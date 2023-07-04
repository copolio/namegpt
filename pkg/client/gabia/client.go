package gabia

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RegistCheckResult struct {
	Domain        string `json:"domain"`
	PunyDomain    string `json:"puny_domain"`
	Flag          string `json:"flag"`
	Status        string `json:"status"`
	Check         string `json:"check"`
	Event         string `json:"event"`
	IsNew         string `json:"is_new"`
	Result        string `json:"result"`
	Reserve       string `json:"reserve"`
	Cart          string `json:"cart"`
	Whois         string `json:"whois"`
	ReserveGubun  string `json:"reserve_gubun"`
	BackorderDate string `json:"backorder_date"`
	WhoisDomain   string `json:"whois_domain"`
	Price         []struct {
		OfferYear      int    `json:"offer_year"`
		DiscountPrice  int    `json:"discount_price"`
		FixedPrice     int    `json:"fixed_price"`
		FixedPriceVat  int    `json:"fixed_price_vat"`
		SavePoint      int    `json:"save_point"`
		VatPrice       int    `json:"vat_price"`
		ViewPriceValue string `json:"view_price_value"`
		AppPrice       int    `json:"app_price"`
		ViewAppPrice   int    `json:"view_app_price"`
	} `json:"price"`
	AppPrice           string        `json:"app_price"`
	PremiumExtendPrice []interface{} `json:"premium_extend_price"`
	PriceDetail        string        `json:"price_detail"`
	DomainTld          string        `json:"domain_tld"`
	Description        string        `json:"description"`
	ParentIndex        bool          `json:"parent_index"`
	Managed            string        `json:"managed"`
	CategoryPath       string        `json:"category_path"`
	CategoryEnd        string        `json:"category_end"`
	EventFlag          string        `json:"event_flag"`
	DisplayTooltip     string        `json:"display_tooltip"`
	ExtraInfo          interface{}   `json:"extra_info"`
	Tab                string        `json:"tab"`
	NewFlag            string        `json:"new_flag"`
	SpecialFlag        string        `json:"special_flag"`
	IsPremium          string        `json:"is_premium"`
	InputSource        string        `json:"input_source"`
	Stype              int           `json:"stype"`
}

func CheckDomainRegist(domain string) (registCheckResult *RegistCheckResult, err error) {
	res, err := http.PostForm("https://domain.gabia.com/ajax_lib/check/regist.php", url.Values{"domain": {domain}, "tab": {"recommend"}})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result RegistCheckResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
