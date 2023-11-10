package workers

import (
	"encoding/json"
	"fmt"
)

type Assignments struct {
	Created  int `json:"created"`
	Assigned int `json:"assigned"`
}

const (
	DEV_URL   string = "http://localhost:3000/api/v1/assignments"
	STAGE_URL string = "https://debitllama-stage.deno.dev/api/v1/assignments"
	PROD_URL  string = "https://debitllama.com/api/v1/assignments"
)

func GetUrl(devenv bool) string {
	// TODO: This is not ready for production, yet.
	if devenv {
		return DEV_URL
	} else {
		return STAGE_URL
	}
}

// If I can't connect to the server I just fail
func GetAssignmentCount(url string, xrelayer string, authkey string) (int, int, error) {
	resp, err := GETRequest(url, xrelayer, authkey)

	if err != nil {
		return 0, 0, err
	}

	//We Read the response body on the line below.
	body, err := ReadResponseBody(resp)

	if err != nil {
		return 0, 0, err
	}

	// Unmarshal the json string
	var assignments Assignments

	err = json.Unmarshal(body, &assignments)
	if err != nil {
		return 0, 0, err
	}

	fmt.Printf("Total Assignments %+v\n", assignments.Created+assignments.Assigned)
	fmt.Printf("Created Assignments %+v\n", assignments.Created)
	fmt.Printf("Assigned Assignments %+v\n", assignments.Assigned)

	return assignments.Created, assignments.Assigned, nil
}

func SolveAnAssignment(xrelayer string, authkey string, devenv bool) error {
	fmt.Println("Starting to solve the assignment")
	url := GetUrl(devenv)
	resp, err := POSTRequestWithoutBody(url, xrelayer, authkey)
	if err != nil {
		return err
	}
	body, err := ReadResponseBody(resp)
	if err != nil {
		return err
	}

	var job Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		return err
	}
	fmt.Println("Got job: ")
	fmt.Println(job)

	return nil
}

type Job struct {
	Assignment            string                      `json:"assignment"`
	PaymentIntent         PaymentIntentRow            `json:"paymentIntent"`
	Account               Account                     `json:"account"`
	RelayerBalance        RelayerBalance              `json:"relayerBalance"`
	DynamicPaymentRequest DynamicPaymentRequestJobRow `json:"dynamicpaymentrequest"`
}

type Pricing string

const (
	Fixed   Pricing = "Fixed"
	Dynamic Pricing = "Dynamic"
)

type AccountTypes string

const (
	VIRTUALACCOUNT  AccountTypes = "VIRTUALACCOUNT"
	CONNECTEDWALLET AccountTypes = "CONNECTEDWALLET"
)

type DynamicPaymentRequestJobsStatus string

const (
	CREATED_DPRJS   DynamicPaymentRequestJobsStatus = "Created"
	LOCKED_DPRJS    DynamicPaymentRequestJobsStatus = "Locked"
	COMPLETED_DPRJS DynamicPaymentRequestJobsStatus = "Completed"
	REJECTED_DPRJS  DynamicPaymentRequestJobsStatus = "Rejected"
)

type PaymentIntentStatus string

const (
	CREATED_PIS              PaymentIntentStatus = "CREATED"
	CANCELLED_PIS            PaymentIntentStatus = "CANCELLED"
	RECURRING_PIS            PaymentIntentStatus = "RECURRING"
	PAID_PIS                 PaymentIntentStatus = "PAID"
	BALANCETOOLOWTORELAY_PIS PaymentIntentStatus = "BALANCETOOLOWTORELAY"
	ACCOUNTBALANCETOOLOW_PIS PaymentIntentStatus = "ACCOUNTBALANCETOOLOW"
)

type DynamicPaymentRequestJobRow struct {
	Id                 int    `json:"id"`
	Created_at         string `json:"created_at"`
	PaymentIntent_id   int    `json:"paymentIntent_id"`
	RequestedAmount    string `json:"requestedAmount"`
	Status             string `json:"status"`
	Request_creator_id string `json:"request_creator_id"`
	AllocatedGas       string `json:"allocatedGas"`
	RelayerBalance_id  int    `json:"relayerBalance_id"`
}

type Account struct {
	Id          int          `json:"id"`
	Created_at  string       `json:"created_at"`
	User_id     string       `json:"user_id"`
	Network_id  string       `json:"network_id"`
	Commitment  string       `json:"commitment"`
	Name        string       `json:"name"`
	Closed      bool         `json:"closed"`
	Currency    string       `json:"Currency"`
	Balance     float64      `json:"balance"`
	AccountType AccountTypes `json:"accountType"`
}
type RelayerBalance struct {
	Id                                int     `json:"id"`
	Created_at                        string  `json:"created_at"`
	BTT_Donau_Testnet_Balance         float64 `json:"BTT_Donau_Testnet_Balance"`
	Missing_BTT_Donau_Testnet_Balance float64 `json:"Missing_BTT_Donau_Testnet_Balance"`
	BTT_Mainnet_Balance               float64 `json:"BTT_Mainnet_Balance"`
	Missing_BTT_Mainnet_Balance       float64 `json:"Missing_BTT_Mainnet_Balance"`
	User_id                           string  `json:"user_id"`
	Last_topup                        string  `json:"last_topupt"`
}

type PaymentIntentRow struct {
	Id                int                 `json:"id"`
	Created_at        string              `json:"created_at"`
	Creator_user_id   string              `json:"creator_user_id"`
	Payee_user_id     string              `json:"payee_user_id"`
	Account_id        int                 `json:"account_id"`
	Payee_address     string              `json:"payee_address"`
	MaxDebitAmount    float64             `json:"maxDebitAmount"`
	DebitTimes        int                 `json:"debitTimes"`
	DebitInterval     int                 `json:"debitInterval"`
	PaymentIntent     string              `json:"paymentIntent"`
	Commitment        string              `json:"commitment"`
	EstimatedGas      string              `json:"estimatedGas"`
	StatusText        PaymentIntentStatus `json:"statusText"`
	LastPaymentDate   string              `json:"lastPaymentDate"`
	NextPaymentDate   string              `json:"NextPaymentDate"`
	Pricing           Pricing             `json:"pricing"`
	Currency          string              `json:"currency"`
	Network           string              `json:"network"`
	Debit_item_id     int                 `json:"debit_item_id"`
	Used_for          int                 `json:"used_for"`
	Proof             string              `json:"proof"`
	PublicSignals     string              `json:"publicSignals"`
	RelayerBalance_id int                 `json:"relayerBalance_id"`
}
