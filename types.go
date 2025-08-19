package myob

import "github.com/omniboost/go-myob/utils"

type Accounts []Account

type Account struct {
	UID                string           `json:"UID"`
	Name               string           `json:"Name"`
	DisplayID          string           `json:"DisplayID"`
	Classification     string           `json:"Classification"`
	Type               string           `json:"Type"`
	Number             int              `json:"Number"`
	Description        string           `json:"Description"`
	ParentAccount      *ParentAccount   `json:"ParentAccount"`
	IsActive           bool             `json:"IsActive"`
	TaxCode            *TaxCode         `json:"TaxCode"`
	Level              int              `json:"Level"`
	OpeningBalance     Number           `json:"OpeningBalance"`
	CurrentBalance     Number           `json:"CurrentBalance"`
	BankingDetails     *BankingDetails  `json:"BankingDetails"`
	IsHeader           bool             `json:"IsHeader"`
	LastReconciledDate *DateTime        `json:"LastReconciledDate"`
	ForeignCurrency    *ForeignCurrency `json:"ForeignCurrency"`
	LastModified       DateTime         `json:"LastModified"`
	URI                utils.URL        `json:"URI"`
	RowVersion         string           `json:"RowVersion"`
}

type ParentAccount struct {
	UID       string    `json:"UID"`
	Name      string    `json:"Name"`
	DisplayID string    `json:"DisplayID"`
	URI       utils.URL `json:"URI"`
}

type TaxCode struct {
	UID  string    `json:"UID"`
	Code string    `json:"Code"`
	URI  utils.URL `json:"URI"`
}

type BankingDetails struct {
	BSBNumber                       string `json:"BSBNumber"`
	BankAccountNumber               string `json:"BankAccountNumber"`
	BankAccountName                 string `json:"BankAccountName"`
	CompanyTradingName              string `json:"CompanyTradingName"`
	BankCode                        string `json:"BankCode"`
	CreateBankFiles                 bool   `json:"CreateBankFiles"`
	DirectEntryUserId               string `json:"DirectEntryUserId"`
	IncludeSelfBalancingTransaction bool   `json:"IncludeSelfBalancingTransaction"`
	StatementParticulars            string `json:"StatementParticulars"`
}

type ForeignCurrency struct {
	UID          string    `json:"UID"`
	Code         string    `json:"Code"`
	CurrencyName string    `json:"CurrencyName"`
	URI          utils.URL `json:"URI"`
}
