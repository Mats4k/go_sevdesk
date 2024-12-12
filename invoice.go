//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of sevdesk.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package go_sevdesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Invoice for the data that the function uses
type Invoice struct {
	ContactID       string
	Address         string
	InvoiceDate     string
	Status          string
	InvoiceType     string
	ContactPerson   string
	Subject         string
	Headtext        string
	FootText        string
	Token           string
	IsEInvoice      bool
	PaymentMethodId string
}

// InvoicesReturn to return the data from invoices
type InvoicesReturn struct {
	Objects []InvoiceObjects `json:"objects"`
}

// NewInvoiceReturn is for returning the data
type NewInvoiceReturn struct {
	Objects InvoiceObjects `json:"objects"`
}

type InvoiceObjects struct {
	ID                                  string     `json:"id"`
	ObjectName                          string     `json:"objectName"`
	AdditionalInformation               string     `json:"additionalInformation"`
	Contact                             ObjectName `json:"contact"`
	Create                              string     `json:"create"`
	Update                              string     `json:"update"`
	SevClient                           ObjectName `json:"sevClient"`
	InvoiceDate                         string     `json:"invoiceDate"`
	Header                              string     `json:"header"`
	HeadText                            string     `json:"headText"`
	FootText                            string     `json:"footText"`
	TimeToPlay                          string     `json:"timeToPlay"`
	DiscountTime                        string     `json:"discountTime"`
	Discount                            string     `json:"discount"`
	AddressName                         string     `json:"addressName"`
	AddressStreet                       string     `json:"addressStreet"`
	AddressZip                          string     `json:"addressZip"`
	AddressCity                         string     `json:"addressCity"`
	PayDate                             string     `json:"payDate"`
	CreateUser                          ObjectName `json:"createUser"`
	DeliveryDate                        string     `json:"deliveryDate"`
	Status                              string     `json:"status"`
	SmallSettlement                     string     `json:"smallSettlement"`
	ContactPerson                       ObjectName `json:"contactPerson"`
	TaxRate                             string     `json:"taxRate"`
	TaxText                             string     `json:"taxText"`
	DunningLevel                        string     `json:"dunningLevel"`
	AddressParentName                   string     `json:"addressParentName"`
	TaxType                             string     `json:"taxType"`
	SendDate                            string     `json:"sendDate"`
	OriginLastInvoice                   string     `json:"originLastInvoice"`
	InvoiceType                         string     `json:"invoiceType"`
	AccountIntervall                    string     `json:"accountIntervall"`
	AccountLastInvoice                  string     `json:"accountLastInvoice"`
	AccountNextInvoice                  string     `json:"accountNextInvoice"`
	ReminderTotal                       string     `json:"reminderTotal"`
	ReminderDebit                       string     `json:"reminderDebit"`
	ReminderDeadline                    string     `json:"reminderDeadline"`
	ReminderCharge                      string     `json:"reminderCharge"`
	AddressParentName2                  string     `json:"addressParentName2"`
	AddressName2                        string     `json:"addressName2"`
	AddressGender                       string     `json:"addressGender"`
	AccountStartDate                    string     `json:"accountStartDate"`
	AccountEndDate                      string     `json:"accountEndDate"`
	Address                             string     `json:"address"`
	Currency                            string     `json:"currency"`
	SumNet                              string     `json:"sumNet"`
	SumTax                              string     `json:"sumTax"`
	SumGross                            string     `json:"sumGross"`
	SumDiscounts                        string     `json:"sumDiscounts"`
	SumNetForeignCurrency               string     `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency               string     `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency             string     `json:"sumGrossForeignCurrency"`
	SumDiscountsForeignCurrency         string     `json:"sumDiscountsForeignCurrency"`
	SumNetAccounting                    string     `json:"sumNetAccounting"`
	SumTaxAccounting                    string     `json:"sumTaxAccounting"`
	SumGrossAccounting                  string     `json:"sumGrossAccounting"`
	PaidAmount                          float64    `json:"paidAmount"`
	CustomerInternalNote                string     `json:"customerInternalNote"`
	ShowNet                             string     `json:"showNet"`
	Enshrined                           string     `json:"enshrined"`
	SendType                            string     `json:"sendType"`
	DeliveryDateUntil                   string     `json:"deliveryDateUntil"`
	SendPaymentReceivedNotificationDate string     `json:"sendPaymentReceivedNotificationDate"`
	SumDiscountNet                      string     `json:"sumDiscountNet"`
	SumDiscountGross                    string     `json:"sumDiscountGross"`
	SumDiscountNetForeignCurrency       string     `json:"sumDiscountNetForeignCurrency"`
	SumDiscountGrossForeignCurrency     string     `json:"sumDiscountGrossForeignCurrency"`
}

type ObjectName struct {
	Id         string `json:"id"`
	ObjectName string `json:"objectName"`
}

// InvoiceEmail is to send invoice by email
type InvoiceEmail struct {
	ID          string
	Email       string
	Subject     string
	Text        string
	CC          string
	BCC         string
	Token       string
	SendOnlyXML bool
}

// SendInvoiceEmailReturn is to decode json data
type SendInvoiceEmailReturn struct {
	Objects SendInvoiceEmailObjects `json:"objects"`
}

type SendInvoiceEmailObjects struct {
	ID                    string                 `json:"id"`
	ObjectName            string                 `json:"objectName"`
	AdditionalInformation interface{}            `json:"additionalInformation"`
	Create                string                 `json:"create"`
	Update                string                 `json:"update"`
	Object                SendInvoiceEmailObject `json:"object"`
	From                  string                 `json:"from"`
	To                    string                 `json:"to"`
	Subject               string                 `json:"subject"`
	Text                  string                 `json:"text"`
	sevClient             ObjectName             `json:"sevClient"`
}

type SendInvoiceEmailObject struct {
	ID                                  string      `json:"id"`
	ObjectName                          string      `json:"objectName"`
	AdditionalInformation               interface{} `json:"additionalInformation"`
	InvoiceNumber                       string      `json:"invoiceNumber"`
	Contact                             ObjectName  `json:"contact"`
	Create                              string      `json:"create"`
	Update                              string      `json:"update"`
	SevClient                           ObjectName  `json:"sevClient"`
	InvoiceDate                         string      `json:"invoice_date"`
	Header                              string      `json:"header"`
	HeadText                            interface{} `json:"headText"`
	FootText                            interface{} `json:"footText"`
	TimeToPay                           interface{} `json:"timeToPay"`
	DiscountTime                        interface{} `json:"discountTime"`
	Discount                            string      `json:"discount"`
	AddressName                         interface{} `json:"addressName"`
	AddressStreet                       interface{} `json:"addressStreet"`
	AddressZip                          interface{} `json:"addressZip"`
	AddressCity                         interface{} `json:"addressCity"`
	AddressCountry                      ObjectName  `json:"addressCountry"`
	PayDate                             interface{} `json:"payDate"`
	CreateUser                          ObjectName  `json:"createUser"`
	DeliveryDate                        string      `json:"deliveryDate"`
	Status                              string      `json:"status"`
	SmallSettlement                     string      `json:"smallSettlement"`
	ContactPerson                       ObjectName  `json:"contactPerson"`
	TaxRate                             string      `json:"taxRate"`
	TaxRule                             ObjectName  `json:"taxRule"`
	TaxText                             string      `json:"taxText"`
	DunningLevel                        interface{} `json:"dunningLevel"`
	AddressParentName                   interface{} `json:"addressParentName"`
	TaxType                             string      `json:"taxType"`
	SendDate                            interface{} `json:"sendDate"`
	OriginLastInvoice                   interface{} `json:"originLastInvoice"`
	InvoiceType                         string      `json:"invoiceType"`
	AccountIntervall                    interface{} `json:"accountIntervall"`
	AccountLastInvoice                  interface{} `json:"accountLastInvoice"`
	AccountNextInvoice                  interface{} `json:"accountNextInvoice"`
	ReminderTotal                       interface{} `json:"reminderTotal"`
	ReminderDebit                       interface{} `json:"reminderDebit"`
	ReminderDeadline                    interface{} `json:"reminderDeadline"`
	ReminderCharge                      interface{} `json:"reminderCharge"`
	AddressParentName2                  interface{} `json:"addressParentName2"`
	AddressName2                        interface{} `json:"addressName2"`
	AddressGender                       interface{} `json:"addressGender"`
	AccountStartDate                    interface{} `json:"accountStartDate"`
	AccountEndDate                      interface{} `json:"accountEndDate"`
	Address                             interface{} `json:"address"`
	Currency                            string      `json:"currency"`
	SumNet                              string      `json:"sumNet"`
	SumTax                              string      `json:"sumTax"`
	SumGross                            string      `json:"sumGross"`
	SumDiscounts                        string      `json:"sumDiscounts"`
	SumNetForeignCurrency               string      `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency               string      `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency             string      `json:"sumGrossForeignCurrency"`
	SumDiscountsForeignCurrency         string      `json:"sumDiscountsForeignCurrency"`
	SumNetAccounting                    string      `json:"sumNetAccounting"`
	SumTaxAccounting                    string      `json:"sumTaxAccounting"`
	SumGrossAccounting                  string      `json:"sumGrossAccounting"`
	PaidAmount                          string      `json:"paidAmount"`
	CustomerInternalNote                interface{} `json:"customerInternalNote"`
	ShowNet                             string      `json:"showNet"`
	Enshrined                           interface{} `json:"enshrined"`
	SendType                            string      `json:"sendType"`
	DeliveryDateUntil                   interface{} `json:"deliveryDateUntil"`
	SendPaymentReceivedNotificationDate interface{} `json:"sendPaymentReceivedNotificationDate"`
	SumDiscountNet                      string      `json:"sumDiscountNet"`
	SumDiscountGross                    string      `json:"sumDiscountGross"`
	SumDiscountNetForeignCurrency       string      `json:"sumDiscountNetForeignCurrency"`
	SumDiscountGrossForeignCurrency     string      `json:"sumDiscountGrossForeignCurrency"`
}

// SendInvoice is to config
type SendInvoice struct {
	ID        string
	SendType  string
	SendDraft string
	Token     string
}

// SendInvoiceReturn is to decode json data
type SendInvoiceReturn struct {
	Objects SendInvoiceObjects `json:"objects"`
}

type SendInvoiceObjects struct {
	ID                                  string               `json:"id"`
	ObjectName                          string               `json:"objectName"`
	AdditionalInformation               interface{}          `json:"additionalInformation"`
	InvoiceNumber                       string               `json:"invoiceNumber"`
	Contact                             ObjectName           `json:"contact"`
	Create                              string               `json:"create"`
	Update                              string               `json:"update"`
	SevClient                           SendInvoiceSevClient `json:"sevClient"`
	InvoiceDate                         string               `json:"invoiceDate"`
	Header                              string               `json:"header"`
	HeadText                            string               `json:"headText"`
	FootText                            string               `json:"footText"`
	TimeToPlay                          interface{}          `json:"timeToPlay"`
	DiscountTime                        string               `json:"discountTime"`
	Discount                            string               `json:"discount"`
	AddressName                         string               `json:"addressName"`
	AddressStreet                       interface{}          `json:"addressStreet"`
	AddressZip                          interface{}          `json:"addressZip"`
	AddressCity                         interface{}          `json:"addressCity"`
	AddressCountry                      ObjectName           `json:"addressCountry"`
	PayDate                             interface{}          `json:"payDate"`
	CreateUser                          ObjectName           `json:"createUser"`
	DeliveryDate                        string               `json:"deliveryDate"`
	Status                              string               `json:"status"`
	SmallSettlement                     string               `json:"smallSettlement"`
	ContactPerson                       ObjectName           `json:"contactPerson"`
	TaxRate                             string               `json:"taxRate"`
	TaxText                             string               `json:"taxText"`
	DunningLevel                        interface{}          `json:"dunningLevel"`
	AddressParentName                   interface{}          `json:"addressParentName"`
	TaxType                             string               `json:"taxType"`
	SendDate                            string               `json:"sendDate"`
	OriginLastInvoice                   interface{}          `json:"originLastInvoice"`
	InvoiceType                         string               `json:"invoiceType"`
	AccountIntervall                    interface{}          `json:"accountIntervall"`
	AccountLastInvoice                  interface{}          `json:"accountLastInvoice"`
	AccountNextInvoice                  interface{}          `json:"accountNextInvoice"`
	ReminderTotal                       interface{}          `json:"reminderTotal"`
	ReminderDebit                       interface{}          `json:"reminderDebit"`
	ReminderDeadline                    interface{}          `json:"reminderDeadline"`
	ReminderCharge                      interface{}          `json:"reminderCharge"`
	AddressParentName2                  interface{}          `json:"addressParentName2"`
	AddressName2                        interface{}          `json:"addressName2"`
	AddressGender                       interface{}          `json:"addressGender"`
	AccountStartDate                    interface{}          `json:"accountStartDate"`
	AccountEndDate                      interface{}          `json:"accountEndDate"`
	Address                             interface{}          `json:"address"`
	Currency                            string               `json:"currency"`
	SumNet                              string               `json:"sumNet"`
	SumTax                              string               `json:"sumTax"`
	SumGross                            string               `json:"sumGross"`
	SumDiscounts                        string               `json:"sumDiscounts"`
	SumNetForeignCurrency               string               `json:"sumNetForeignCurrency"`
	SumTaxForeignCurrency               string               `json:"sumTaxForeignCurrency"`
	SumGrossForeignCurrency             string               `json:"sumGrossForeignCurrency"`
	SumDiscountsForeignCurrency         string               `json:"sumDiscountsForeignCurrency"`
	SumNetAccounting                    string               `json:"sumNetAccounting"`
	SumTaxAccounting                    string               `json:"sumTaxAccounting"`
	SumGrossAccounting                  string               `json:"sumGrossAccounting"`
	PaidAmount                          int                  `json:"paidAmount"`
	CustomerInternalNote                interface{}          `json:"customerInternalNote"`
	ShowNet                             string               `json:"showNet"`
	Enshrined                           interface{}          `json:"enshrined"`
	SendType                            string               `json:"sendType"`
	DeliveryDateUntil                   interface{}          `json:"deliveryDateUntil"`
	SendPaymentReceivedNotificationDate interface{}          `json:"sendPaymentReceivedNotificationDate"`
	SumDiscountNet                      string               `json:"sumDiscountNet"`
	SumDiscountGross                    string               `json:"sumDiscountGross"`
	SumDiscountNetForeignCurrency       string               `json:"sumDiscountNetForeignCurrency"`
	SumDiscountGrossForeignCurrency     string               `json:"sumDiscountGrossForeignCurrency"`
}

type SendInvoiceSevClient struct {
	ID                        string                      `json:"id"`
	ObjectName                string                      `json:"objectName"`
	AdditionalInformation     interface{}                 `json:"additionalInformation"`
	Create                    string                      `json:"create"`
	Update                    string                      `json:"update"`
	Name                      string                      `json:"name"`
	TemplateMainColor         string                      `json:"templateMainColor"`
	TemplateSubColor          string                      `json:"templateSubColor"`
	Status                    string                      `json:"status"`
	AddressStreet             string                      `json:"addressStreet"`
	AddressCity               string                      `json:"addressCity"`
	AddressZip                string                      `json:"addressZip"`
	MuncipalityKey            interface{}                 `json:"muncipalityKey"`
	ContactPhone              string                      `json:"contactPhone"`
	ContactEmail              string                      `json:"contactEmail"`
	PaypalEmail               interface{}                 `json:"paypalEmail"`
	VatNumber                 string                      `json:"vatNumber"`
	CeoName                   string                      `json:"ceoName"`
	ContactFax                interface{}                 `json:"contactFax"`
	Domain                    string                      `json:"domain"`
	TemplateHeadlineColor     string                      `json:"templateHeadlineColor"`
	Website                   string                      `json:"website"`
	Bank                      string                      `json:"bank"`
	BankNumber                string                      `json:"bankNumber"`
	BankAccountNumber         string                      `json:"bankAccountNumber"`
	DistrictCourt             string                      `json:"districtCourt"`
	SmallSettlement           string                      `json:"smallSettlement"`
	HasSeenBasicSettingsModal string                      `json:"hasSeenBasicSettingsModal"`
	Bank2                     interface{}                 `json:"bank2"`
	BankIban                  string                      `json:"bankIban"`
	BankBic                   string                      `json:"bankBic"`
	OrderPaymentTerms         interface{}                 `json:"orderPaymentTerms"`
	OrderDeliveryTerms        interface{}                 `json:"orderDeliveryTerms"`
	InvoiceTimeToPay          interface{}                 `json:"invoiceTimeToPay"`
	Type                      string                      `json:"type"`
	ForeignId                 interface{}                 `json:"foreignId"`
	DefaultBillingTime        interface{}                 `json:"defaultBillingTime"`
	BillingAccountNumber      interface{}                 `json:"billingAccountNumber"`
	BillingBankCode           interface{}                 `json:"billingBankCode"`
	CompanyRegister           string                      `json:"companyRegister"`
	NameAddition              string                      `json:"nameDddition"`
	ShowNet                   string                      `json:"showNet"`
	PrintShowQr               string                      `json:"printShowQr"`
	PrintShowPayPal           string                      `json:"printShowPayPal"`
	PrintDeliveryDate         string                      `json:"printDeliveryDate"`
	PrintContactPerson        string                      `json:"printContactPerson"`
	TaxNumber                 string                      `json:"taxNumber"`
	PrintDeliveryReturn       string                      `json:"printDeliveryReturn"`
	PrintPartNumber           string                      `json:"printPartNumber"`
	PrintPosDescription       string                      `json:"printPosDescription"`
	ChartOfAccounts           string                      `json:"chartOfAccounts"`
	AccountingChart           ObjectName                  `json:"accountingChart"`
	AccountingSystem          SendInvoiceAccountingSystem `json:"accountingSystem"`
	ContractNotePaymentTerms  interface{}                 `json:"contractNotePaymentTerms"`
	ContractNoteDeliveryTerms interface{}                 `json:"contractNoteDeliveryTerms"`
	PackingListDeliveryTerms  interface{}                 `json:"packingListDeliveryTerms"`
	InvitedBy                 interface{}                 `json:"invitedBy"`
	PartnerID                 string                      `json:"partnerId"`
	InviterRewarded           string                      `json:"inviterRewarded"`
	InvoiceReminderDuration   interface{}                 `json:"invoiceReminderDuration"`
	InvoiceReminderCharge     interface{}                 `json:"invoiceReminderCharge"`
	AutoDeliveryDate          string                      `json:"autoDeliveryDate"`
	AddressCountry            SendInvoiceAddressCountry   `json:"addressCountry"`
	PrintPdfDeliveryReturn    string                      `json:"printPdfDeliveryReturn"`
	PrintPdfFooter            string                      `json:"printPdfFooter"`
	DefaultCurrency           string                      `json:"defaultCurrency"`
	PrintPageNumbers          string                      `json:"printPageNumbers"`
	FormOfCompany             string                      `json:"formOfCompany"`
	ComapnySize               string                      `json:"comapnySize"`
	ReferralProgram           ObjectName                  `json:"referralProgram"`
	onBoardingStatus          interface{}                 `json:"onBoardingStatus"`
	PactasID                  string                      `json:"pactasId"`
	DocServer                 string                      `json:"docServer"`
	CreateInvoiceReminder     string                      `json:"createInvoiceReminder"`
	AccountantNumber          string                      `json:"accountantNumber"`
	AccountantClientNumber    string                      `json:"accountantClientNumber"`
	AccountingYearBegin       string                      `json:"accountingYearBegin"`
	CancelationDate           interface{}                 `json:"cancelationDate"`
	FigoUsername              string                      `json:"figoUsername"`
	SubscriptionStartDate     string                      `json:"subscriptionStartDate"`
	SubscriptionEndDate       interface{}                 `json:"subscriptionEndDate"`
	SubscriptionCycle         string                      `json:"subscriptionCycle"`
	PrintFoldLines            string                      `json:"printFoldLines"`
	DiscoveredAt              string                      `json:"discoveredAt"`
	FormerBookkeepingTool     interface{}                 `json:"formerBookkeepingTool"`
	HasAccountant             string                      `json:"hasAccountant"`
	PartnerBrand              interface{}                 `json:"partnerBrand"`
	PartnerCustomerID         interface{}                 `json:"partnerCustomerId"`
	PartnerProvisioningID     interface{}                 `json:"partnerProvisioningId"`
	PartnerWhitelabeled       string                      `json:"partnerWhitelabeled"`
	Tenant                    string                      `json:"tenant"`
	state                     string                      `json:"state"`
}

type SendInvoiceAccountingSystem struct {
	ID                    string      `json:"id"`
	ObjectName            string      `json:"objectName"`
	AdditionalInformation interface{} `json:"additionalInformation"`
	Create                interface{} `json:"create"`
	Update                interface{} `json:"update"`
	Name                  string      `json:"name"`
	AccountingChart       ObjectName  `json:"accountingChart"`
}

type SendInvoiceAddressCountry struct {
	ID                    string      `json:"id"`
	ObjectName            string      `json:"objectName"`
	AdditionalInformation interface{} `json:"additionalInformation"`
	Code                  string      `json:"code"`
	Name                  string      `json:"name"`
	NameEn                string      `json:"nameEn"`
	TranslationCode       string      `json:"translationCode"`
	Locale                string      `json:"locale"`
	Priority              string      `json:"priority"`
}

// DownloadInvoice is to config
type DownloadInvoice struct {
	ID            string
	PreventSendBy string
	Download      string
	Token         string
}

// Invoices to check invoices
func Invoices(token string) (InvoicesReturn, error) {

	// Define client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", buildURL("Invoice"), nil)
	if err != nil {
		return InvoicesReturn{}, err
	}

	// Set headers
	request.Header.Set("Authorization", token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return InvoicesReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode InvoicesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return InvoicesReturn{}, err
	}

	// Return data
	return decode, nil

}

// NewInvoice to create a new invoice
func NewInvoice(config Invoice, posConfigs []Position) (NewInvoiceReturn, error) {

	//log.Println("Create invoice")

	// Define client
	client := &http.Client{}

	// Define body data
	body := url.Values{}
	body.Set("invoiceNumber", "")
	body.Set("contact[id]", config.ContactID)
	body.Set("contact[objectName]", "Contact")
	body.Set("address", config.Address)
	body.Set("invoiceDate", config.InvoiceDate)
	body.Set("header", config.Subject)
	body.Set("status", config.Status)
	body.Set("headText", config.Headtext)
	body.Set("footText", config.FootText)
	body.Set("invoiceType", config.InvoiceType)
	body.Set("currency", "EUR")
	body.Set("mapAll", "true")
	body.Set("objectName", "Invoice")
	body.Set("discount", "0")
	body.Set("contactPerson[id]", config.ContactPerson)
	body.Set("contactPerson[objectName]", "SevUser")
	body.Set("taxType", "default")
	body.Set("taxRate", "")
	body.Set("taxRule[id]", "1")
	body.Set("taxRule[objectName]", "TaxRule")
	body.Set("taxText", "0")
	body.Set("showNet", "false")
	body.Set("paymentMethod[id]", config.PaymentMethodId)
	body.Set("paymentMethod[objectName]", "PaymentMethod")

	if config.IsEInvoice {
		body.Set("isEInvoice", "true")
	} else {
		body.Set("isEInvoice", "false")
	}

	for i, pos := range posConfigs {
		// Convert string to float64
		priceNet, err := strconv.ParseFloat(pos.PriceNet, 64)
		if err != nil {
			return NewInvoiceReturn{}, err
		}

		// Convert taxRate from string to float64
		taxRate, err := strconv.ParseFloat(pos.TaxRate, 64)
		if err != nil {
			return NewInvoiceReturn{}, err
		}

		// Calc gross
		priceGross := priceNet + (priceNet * taxRate / 100)

		body.Set(fmt.Sprintf("invoicePosSave[%d][price]", i), fmt.Sprintf("%.2f", priceGross))
		body.Set(fmt.Sprintf("invoicePosSave[%d][quantity]", i), pos.Quantity)
		body.Set(fmt.Sprintf("invoicePosSave[%d][taxRate]", i), pos.TaxRate)
		body.Set(fmt.Sprintf("invoicePosSave[%d][name]", i), pos.Name)
		body.Set(fmt.Sprintf("invoicePosSave[%d][text]", i), pos.Description)
		body.Set(fmt.Sprintf("invoicePosSave[%d][unity][id]", i), pos.UnityID)
		body.Set(fmt.Sprintf("invoicePosSave[%d][unity][objectName]", i), "Unity")
		body.Set(fmt.Sprintf("invoicePosSave[%d][objectName]", i), "InvoicePos")
		body.Set(fmt.Sprintf("invoicePosSave[%d][mapAll]", i), "true")
	}

	// New http request
	request, err := http.NewRequest("POST", buildURL("Invoice/Factory/saveInvoice"), strings.NewReader(body.Encode()))
	if err != nil {
		return NewInvoiceReturn{}, err
	}

	//log.Println("Request created")

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return NewInvoiceReturn{}, err
	}

	//log.Println("Response received")

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode NewInvoiceReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return NewInvoiceReturn{}, err
	}

	//log.Println("Decode: ", decode)

	// Return data
	return decode, nil

}

// SendInvoiceEmail to send an invoice by mail
func SendInvoiceEmail(config InvoiceEmail) (SendInvoiceEmailReturn, error) {

	// Define client
	client := &http.Client{}

	// Define body data
	body := url.Values{}
	body.Set("toEmail", config.Email)
	body.Set("subject", config.Subject)
	body.Set("text", config.Text)
	body.Set("copy", "false")
	body.Set("additionalAttachments", "null")
	body.Set("ccEmail", config.CC)
	body.Set("bccEmail", config.BCC)

	if config.SendOnlyXML {
		body.Set("sendXml", "true")
	} else {
		body.Set("sendXml", "false")
	}

	// New http request
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/Invoice/%s/sendViaEmail", ApiURL, config.ID), strings.NewReader(body.Encode()))
	if err != nil {
		return SendInvoiceEmailReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", config.Token)

	// Response to sevDesk
	response, err := client.Do(request)
	if err != nil {
		return SendInvoiceEmailReturn{}, err
	}

	// Close response
	defer response.Body.Close()

	// Decode data
	var decode SendInvoiceEmailReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SendInvoiceEmailReturn{}, err
	}

	// Return data
	return decode, nil

}
