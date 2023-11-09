package eth

type ChainIds string

const (
	BTT_TESTNET_ID ChainIds = "0x405"
	BTT_MAINNET_ID ChainIds = "0xc7"
)

type VirtualAccountsContractAddress string

const (
	VirtualAccounts_BTT_TESTNET VirtualAccountsContractAddress = "0xF75515Df5AC843a8B261E232bB890dc2F75A4066"
	VirtualAccounts_BTT_MAINNET VirtualAccountsContractAddress = "0xc4Cf42D5a6F4F061cf5F98d0338FC5913b6fF581"
)

type ConnectedWalletsContractAddress string

const (
	ConnectedWallets_BTT_TESTNET ConnectedWalletsContractAddress = "0x9c85da9E45126Fd45BC62656026A2E7226bba239"
	COnnectedWallets_BTT_MAINNET ConnectedWalletsContractAddress = "0xF9962f3C23De4e864E56ef29125D460c785905c6"
)

type RPCURLS string

const (
	RPCURLS_BTT_TESTNET RPCURLS = "https://pre-rpc.bt.io/"
	RPCURLS_BTT_MAINNET RPCURLS = "https://rpc.bittorrentchain.io"
)

func (chain ChainIds) getVirtualAccountsContractAddress() VirtualAccountsContractAddress {
	switch chain {
	case BTT_TESTNET_ID:
		return VirtualAccounts_BTT_TESTNET
	case BTT_MAINNET_ID:
		return VirtualAccounts_BTT_MAINNET
	}
	return ""
}

func (chain ChainIds) getConnectedWalletsContractAddress() ConnectedWalletsContractAddress {
	switch chain {
	case BTT_TESTNET_ID:
		return ConnectedWallets_BTT_TESTNET
	case BTT_MAINNET_ID:
		return COnnectedWallets_BTT_MAINNET
	}
	return ""
}

func (chain ChainIds) getRPC() RPCURLS {
	switch chain {
	case BTT_TESTNET_ID:
		return RPCURLS_BTT_TESTNET
	case BTT_MAINNET_ID:
		return RPCURLS_BTT_MAINNET
	}
	return ""
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

type Account struct {
	id          int
	created_at  string
	user_id     string
	network_id  string
	commitment  string
	name        string
	closed      bool
	currency    string
	balance     string
	accountType AccountTypes
}

type RelayerBalance struct {
	id                                int
	created_at                        string
	BTT_Donau_Testnet_Balance         string
	Missing_BTT_Donau_Testnet_Balance string
	BTT_Mainnet_Balance               string
	Missing_BTT_Mainnet_Balance       string
	user_id                           string
	last_topup                        string
}

type PaymentIntentRow struct {
	id                int
	created_at        string
	creator_user_id   string
	payee_user_id     string
	account_id        string
	payee_address     string
	maxDebitAmount    string
	debitTimes        int
	debitInterval     int
	paymentIntent     string
	commitment        string
	estimatedGas      string
	statusText        string
	lastPaymentDate   string
	nextPaymentDate   string
	pricing           Pricing
	currency          string
	network           string
	debit_item_id     int
	used_for          int
	proof             string
	publicSignals     string
	relayerBalance_id RelayerBalance
}

type DynamicPaymentRequestJobRow struct {
	id                 int
	created_at         string
	paymentIntent_id   PaymentIntentRow
	requestedAmount    string
	status             string
	request_creator_id string
	allocatedGas       string
	relayerBalance_id  int
}
