package workers

type CliArgs struct {
	Devenv   bool
	Xrelayer string
	Authkey  string
	Password string
	Chains   []string
}

type AssignmentsPerNetwork struct {
	Jobs    int    `json:"jobs"`
	Network string `json:"network"`
}

type Assignments struct {
	Created  *[]AssignmentsPerNetwork `json:"created"`
	Assigned *[]AssignmentsPerNetwork `json:"assigned"`
}

type DirectDebitAssignment struct {
	Assignment           string                   `json:"assignment"`
	Abi                  DirectDebitAssignmentAbi `json:"abi"`
	Network              string                   `json:"network"`
	ContractAddress      string                   `json:"contractAddress"`
	Inputs               DirectDebitInput         `json:"inputs"`
	GasPrice             string                   `json:"gasPrice"`
	GasLimit             string                   `json:"gasLimit"`
	MaxFeePerGas         string                   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string                   `json:"maxPriorityFeePerGas"`
}

type DirectDebitAssignmentAbi struct {
	Inputs          [4]AbiVariable `json:"inputs"`
	Name            string         `json:"name"`
	Outputs         []AbiVariable  `json:"outputs"`
	StateMutability string         `json:"stateMutability"`
	Type            string         `json:"type"`
}

type AbiVariable struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type DirectDebitInput struct {
	Proof  [8]string `json:"proof"`
	Hashes [2]string `json:"hashes"`
	Payee  string    `json:"payee"`
	Debit  [4]string `json:"debit"`
}
