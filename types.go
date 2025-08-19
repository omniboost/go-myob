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

type Categories []Category

type Category struct {
	UID         string    `json:"UID"`
	DisplayID   string    `json:"DisplayID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	IsActive    bool      `json:"IsActive"`
	URI         utils.URL `json:"URI"`
	RowVersion  string    `json:"RowVersion"`
}

type Jobs []Job

type Job struct {
	UID                string          `json:"UID"`
	Number             string          `json:"Number"`
	IsHeader           bool            `json:"IsHeader"`
	Name               string          `json:"Name"`
	Description        string          `json:"Description"`
	ParentJob          *ParentJob      `json:"ParentJob"`
	LinkedCustomer     *LinkedCustomer `json:"LinkedCustomer"`
	PercentComplete    Number          `json:"PercentComplete"`
	StartDate          *DateTime       `json:"StartDate"`
	FinishDate         *DateTime       `json:"FinishDate"`
	Contact            *string         `json:"Contact"`
	Manager            *string         `json:"Manager"`
	IsActive           bool            `json:"IsActive"`
	TrackReimbursables bool            `json:"TrackReimbursables"`
	LastModified       DateTime        `json:"LastModified"`
	URI                utils.URL       `json:"URI"`
	RowVersion         string          `json:"RowVersion"`
}

type ParentJob struct {
	UID    string    `json:"UID"`
	Number string    `json:"Number"`
	Name   string    `json:"Name"`
	URI    utils.URL `json:"URI"`
}

type LinkedCustomer struct {
	UID       string    `json:"UID"`
	Name      string    `json:"Name"`
	DisplayID string    `json:"DisplayID"`
	URI       utils.URL `json:"URI"`
}
