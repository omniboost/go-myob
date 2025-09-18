package myob

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
	TaxCode            *AccountTaxCode  `json:"TaxCode"`
	Level              int              `json:"Level"`
	OpeningBalance     Number           `json:"OpeningBalance"`
	CurrentBalance     Number           `json:"CurrentBalance"`
	BankingDetails     *BankingDetails  `json:"BankingDetails"`
	IsHeader           bool             `json:"IsHeader"`
	LastReconciledDate *DateTime        `json:"LastReconciledDate"`
	ForeignCurrency    *ForeignCurrency `json:"ForeignCurrency"`
	LastModified       DateTime         `json:"LastModified"`
	URI                string           `json:"URI"`
	RowVersion         string           `json:"RowVersion"`
}

type ParentAccount struct {
	UID       string `json:"UID"`
	Name      string `json:"Name"`
	DisplayID string `json:"DisplayID"`
	URI       string `json:"URI"`
}

type AccountTaxCode struct {
	UID  string `json:"UID"`
	Code string `json:"Code"`
	URI  string `json:"URI"`
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
	UID          string `json:"UID"`
	Code         string `json:"Code"`
	CurrencyName string `json:"CurrencyName"`
	URI          string `json:"URI"`
}

type Categories []Category

type Category struct {
	UID         string `json:"UID"`
	DisplayID   string `json:"DisplayID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	IsActive    bool   `json:"IsActive"`
	URI         string `json:"URI"`
	RowVersion  string `json:"RowVersion"`
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
	URI                string          `json:"URI"`
	RowVersion         string          `json:"RowVersion"`
}

type ParentJob struct {
	UID    string `json:"UID"`
	Number string `json:"Number"`
	Name   string `json:"Name"`
	URI    string `json:"URI"`
}

type LinkedCustomer struct {
	UID       string `json:"UID"`
	Name      string `json:"Name"`
	DisplayID string `json:"DisplayID"`
	URI       string `json:"URI"`
}

type GeneralJournals []GeneralJournal

type GeneralJournal struct {
	UID                  string                         `json:"UID,omitempty"`
	DisplayID            string                         `json:"DisplayID"`
	DateOccurred         DateTime                       `json:"DateOccurred"`
	IsTaxInclusive       bool                           `json:"IsTaxInclusive"`
	Memo                 string                         `json:"Memo"`
	GSTReportingMethod   string                         `json:"GSTReportingMethod"`
	IsYearEndAdjustment  bool                           `json:"IsYearEndAdjustment"`
	Category             *GeneralJournalCategory        `json:"Category,omitempty"`
	Lines                []GeneralJournalLine           `json:"Lines"`
	ForeignCurrency      *GeneralJournalForeignCurrency `json:"ForeignCurrency"`
	URI                  string                         `json:"URI,omitempty,omitzero"`
	CurrencyExchangeRate *Number                        `json:"CurrencyExchangeRate,omitempty"`
	RowVersion           string                         `json:"RowVersion,omitempty"`
}

type GeneralJournalCategory struct {
	UID       string `json:"UID"`
	Name      string `json:"Name,omitempty"`
	DisplayID string `json:"DisplayID,omitempty"`
	URI       string `json:"URI,omitempty,omitzero"`
}

type GeneralJournalForeignCurrency struct {
	UID          string `json:"UID"`
	Code         string `json:"Code,omitempty"`
	CurrencyName string `json:"CurrencyName,omitempty"`
	URI          string `json:"URI,omitempty,omitzero"`
}

type GeneralJournalLine struct {
	RowID                 *int                       `json:"RowID,omitempty"`
	Account               GeneralJournalLineAccount  `json:"Account"`
	Job                   *GeneralJournalLineJob     `json:"Job"`
	Memo                  string                     `json:"Memo"`
	TaxCode               *GeneralJournalLineTaxCode `json:"TaxCode"`
	Amount                Number                     `json:"Amount"`
	AmountForeign         *Number                    `json:"AmountForeign,omitempty"`
	IsCredit              bool                       `json:"IsCredit"`
	TaxAmount             Number                     `json:"TaxAmount"`
	TaxAmountForeign      *Number                    `json:"TaxAmountForeign,omitempty"`
	IsOverriddenTaxAmount bool                       `json:"IsOverriddenTaxAmount"`
	UnitCount             *Number                    `json:"UnitCount"`
	RowVersion            string                     `json:"RowVersion,omitempty"`
}

type GeneralJournalLineAccount struct {
	UID       string `json:"UID"`
	Name      string `json:"Name,omitempty"`
	DisplayID string `json:"DisplayID,omitempty"`
	URI       string `json:"URI,omitempty,omitzero"`
}

type GeneralJournalLineJob struct {
	UID    string `json:"UID"`
	Number string `json:"Number,omitempty"`
	Name   string `json:"Name,omitempty"`
	URI    string `json:"URI,omitempty,omitzero"`
}

type GeneralJournalLineTaxCode struct {
	UID  string `json:"UID"`
	Code string `json:"Code,omitempty"`
	URI  string `json:"URI,omitempty,omitzero"`
}

type TaxCodes []TaxCode

type TaxCode struct {
	UID                       string                 `json:"UID"`
	Code                      string                 `json:"Code"`
	Description               string                 `json:"Description"`
	Type                      string                 `json:"Type"`
	Rate                      Number                 `json:"Rate"`
	IsRateNegative            bool                   `json:"IsRateNegative"`
	TaxCollectedAccount       *TaxCodeAccount        `json:"TaxCollectedAccount"`
	TaxPaidAccount            *TaxCodeAccount        `json:"TaxPaidAccount"`
	WithholdingCreditAccount  *TaxCodeAccount        `json:"WithholdingCreditAccount"`
	WithholdingPayableAccount *TaxCodeAccount        `json:"WithholdingPayableAccount"`
	ImportDutyPayableAccount  *TaxCodeAccount        `json:"ImportDutyPayableAccount"`
	LinkedSupplier            *TaxCodeLinkedSupplier `json:"LinkedSupplier"`
	LuxuryCarTaxThreshold     *Number                `json:"LuxuryCarTaxThreshold"`
	LastModified              DateTime               `json:"LastModified"`
	URI                       string                 `json:"URI"`
	RowVersion                string                 `json:"RowVersion"`
}

type TaxCodeAccount struct {
	UID       string `json:"UID"`
	Name      string `json:"Name"`
	DisplayID string `json:"DisplayID"`
	URI       string `json:"URI"`
}

type TaxCodeLinkedSupplier struct {
	UID       string `json:"UID"`
	Name      string `json:"Name"`
	DisplayID string `json:"DisplayID"`
	URI       string `json:"URI"`
}
