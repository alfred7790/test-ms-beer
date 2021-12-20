package currencylayer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-ms-beer/pkg/utilities"
	"time"
)

const (
	converURL = "%s/convert?access_key=%s&from=%s&to=%s&amount=%v"
	liveURL   = "%s/live?access_key=%s&currencies=AUD,CAD,PLN,MXN&format=1"
)

type CurrencyResponse struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		USDUSD int     `json:"USDUSD"`
		USDAUD float64 `json:"USDAUD"`
		USDCAD float64 `json:"USDCAD"`
		USDPLN float64 `json:"USDPLN"`
		USDMXN float64 `json:"USDMXN"`
	} `json:"quotes"`
	Query struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	} `json:"query"`
	Info struct {
		Timestamp int     `json:"timestamp"`
		Quote     float64 `json:"quote"`
	} `json:"info"`
	Result float64 `json:"result"`
}

type CurrencyErrReponse struct {
	Success bool `json:"success"`
	Error   struct {
		Code int    `json:"code"`
		Info string `json:"info"`
	} `json:"error"`
}

// Service interface contains main functions
type Service interface {
	Live() (*CurrencyResponse, error)
	Convert(from, to string, amount float64) (*CurrencyResponse, error)
	SetHostAndKey(host, accessKey string)
}

//Handle struct for vcc handle for url
type Handle struct {
	CurrencyHost           string
	AccessKey              string
	lastSuccessfulSyncTime time.Time
}

//NewHandle function to create new Handle object
func NewHandle() *Handle {
	return &Handle{}
}

func (v *Handle) SetHostAndKey(host, accessKey string) {
	v.CurrencyHost = host
	v.AccessKey = accessKey
}

func (v *Handle) Live() (*CurrencyResponse, error) {
	url := fmt.Sprintf(liveURL, v.CurrencyHost, v.AccessKey)
	return v.requestCurrency(url)
}

func (v *Handle) Convert(from, to string, amount float64) (*CurrencyResponse, error) {
	url := fmt.Sprintf(converURL, v.CurrencyHost, v.AccessKey, from, to, amount)
	return v.requestCurrency(url)
}

func (v *Handle) requestCurrency(url string) (*CurrencyResponse, error) {
	options := &utilities.Options{
		Method: http.MethodGet,
		URL:    url,
	}

	result, err := utilities.RequestJSON(options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	if result.StatusCode != 200 {
		str := fmt.Sprintf("Attempt to query failed with status code %d", result.StatusCode)
		fmt.Println(str)
		return nil, errors.New(str)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	var currencyBody CurrencyResponse
	var currencyBodyErr CurrencyErrReponse
	if err = json.Unmarshal(body, &currencyBody); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	if !currencyBody.Success {
		if err = json.Unmarshal(body, &currencyBodyErr); err != nil {
			fmt.Printf("Error: %v\n", err)
			return nil, err
		}

		fmt.Println(currencyBodyErr.Error)
		return nil, errors.New(currencyBodyErr.Error.Info)
	}

	return &currencyBody, nil
}
