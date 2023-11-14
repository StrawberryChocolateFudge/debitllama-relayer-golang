package workers

import (
	"encoding/json"
	"fmt"
)

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
func GetAssignmentCount(url string, xrelayer string, authkey string) (*[]AssignmentsPerNetwork, *[]AssignmentsPerNetwork, error) {
	resp, err := GETRequest(url, xrelayer, authkey)

	if err != nil {
		return nil, nil, err
	}

	//We Read the response body on the line below.
	body, err := ReadResponseBody(resp)

	if err != nil {
		return nil, nil, err
	}

	// Unmarshal the json string
	var assignments Assignments

	err = json.Unmarshal(body, &assignments)
	if err != nil {
		return nil, nil, err
	}

	// fmt.Printf("Created Assignments %+v\n", assignments.Created)
	// fmt.Printf("Assigned Assignments %+v\n", assignments.Assigned)

	return assignments.Created, assignments.Assigned, nil
}

func SolveAnAssignment(xrelayer string, authkey string, devenv bool, password string, chainid ChainIds) error {
	fmt.Println("Starting to solve the assignment")
	url := GetUrl(devenv)
	reqBody, err := EncodeGetJobAssignmentReqBody(chainid)
	if err != nil {
		return err
	}
	resp, err := POSTRequestWithBody(url, xrelayer, authkey, reqBody)
	if err != nil {
		return err
	}
	body, err := ReadResponseBody(resp)
	if err != nil {
		return err
	}
	var assignment DirectDebitAssignment
	err = json.Unmarshal(body, &assignment)
	if err != nil {
		return err
	}

	// _, err = assignment.RelayTx(password)

	// if err != nil {
	// 	return err
	// }

	return nil
}
