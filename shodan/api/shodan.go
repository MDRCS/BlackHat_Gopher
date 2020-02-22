package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

type APIInfo struct {
	ScanCredits int `json:scan_credits`
	UsageLimits UsageLimits `json:usage_limits`
	Plan string `json:plan`
	Https bool `json:https`
	Unlocked bool `json:unlocked`
	QueryCredits int `json:query_credits`
	MonitoredIps int `json:monitored_ips`
	UnlockedLeft int `json:unlocked_left`
	Telnet bool `json:telnet`
}

type UsageLimits struct {
	ScanCredits int `json:scan_credits`
	QueryCredits int `json:query_credits`
	MonitoredIps int `json:monitored_ips`
}

type HostLocation struct {
	City 		 string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode	 int     `json:"area_code"`
	Longitude	 float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}
type Host struct {
	OS        string 		`json:"os"`
	Timestamp string 		`json:"timestamp"`
	ISP       string 		`json:"isp"`
	ASN       string 		`json:"asn"`
	Hostnames []string		`json:"hostnames"`
	Location  HostLocation  `json:"location"`
	IP        int64			`json:"ip"`
	Domains   []string  	`json:"domains"`
	Org       string    	`json:"org"`
	Data      string    	`json:"data"`
	Port      int			`json:"port"`
	IPString  string        `json:"ip_str"`
}

type HostSearch struct {
	Matches []Host `json:"matches"`
}

func New(apiKey string) *Client {
	return &Client{apiKey:apiKey}
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Client) HostSearch(query string) (*HostSearch,error){
	res, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseURL, s.apiKey, query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ret HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}