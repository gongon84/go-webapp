package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/go-webapp/cmd/config"
	"github.com/labstack/echo/v4"
)

type Company struct {
	Total   int      `json:"total"`
	Start   int      `json:"start"`
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

type Result struct {
	CorporationID            int    `json:"corporationId"`
	Name                     string `json:"name"`
	NameKana                 string `json:"nameKana"`
	NameEn                   string `json:"nameEn"`
	PostalCode               string `json:"postalCode"`
	LocalGovernmentCode      string `json:"localGovernmentCode"`
	Prefecture               string `json:"prefecture"`
	City                     string `json:"city"`
	Town                     string `json:"town"`
	Block                    string `json:"block"`
	Building                 string `json:"building"`
	PresidentPosition        string `json:"presidentPosition"`
	PresidentName            string `json:"presidentName"`
	Capital                  int    `json:"capital"`
	Employees                int    `json:"employees"`
	EstablishmentDate        string `json:"establishmentDate"`
	IndustryCode             string `json:"industryCode"`
	Listed                   int    `json:"listed"`
	StockCode                int    `json:"stockCode"`
	AverageAge               string `json:"averageAge"`
	FemaleRate               string `json:"femaleRate"`
	AverageAnnualIncome      string `json:"averageAnnualIncome"`
	PaidHolidayDigestibility string `json:"paidHolidayDigestibility"`
	TurnoverRate             string `json:"turnoverRate"`
	FemaleManagerRate        string `json:"femaleManagerRate"`
	HandicappedEmployeeRate  string `json:"handicappedEmployeeRate"`
	AverageDuration          string `json:"averageDuration"`
	Sales                    int    `json:"sales"`
	SalesDate                string `json:"salesDate"`
	CurrentEarnings          int    `json:"currentEarnings"`
	CurrentEarningsDate      string `json:"currentEarningsDate"`
	Branch                   string `json:"branch"`
	FacebookURL              string `json:"facebookUrl"`
	TwitterURL               string `json:"twitterUrl"`
	Tel                      string `json:"tel"`
	WebURL                   string `json:"webUrl"`
	LogoImgUrlPc             string `json:"logoImgUrlPc"`
	LogoImgUrlSp             string `json:"logoImgUrlSp"`
	PrText                   string `json:"prText"`
	Remarks                  string `json:"remarks"`
	CpName                   string `json:"cpName"`
	YCompanyID               string `json:"yCompanyId"`
}

func CompanyDetail(c echo.Context, cfg config.Configs) error {
	yahooApiKey := cfg.Env.YahooApiKey
	startNum := c.QueryParam("start")
	localGovernmentCode := c.QueryParam("localGovernmentCode")
	apiUrl := fmt.Sprintf("https://job.yahooapis.jp/v1/furusato/company/?appid=%s&results=10&localGovernmentCode=%s&start=%s", yahooApiKey, localGovernmentCode, startNum)

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data Company
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	return c.Render(http.StatusOK, "company.html", data)
}
