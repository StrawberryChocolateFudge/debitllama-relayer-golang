package workers

import "fmt"

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

func (chain ChainIds) getFlag() string {
	switch chain {
	case BTT_TESTNET_ID:
		return "btt_testnet"
	case BTT_MAINNET_ID:
		return "btt"
	}
	return ""
}

func (chain ChainIds) getName() string {
	switch chain {
	case BTT_TESTNET_ID:
		return "BitTorrent Donau Testnet"
	case BTT_MAINNET_ID:
		return "BitTorrent Chain"
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

func PrintChainData(chid ChainIds) {
	fmt.Println("######################################################################")
	fmt.Println(chid.getName())
	fmt.Println("Flag: ", chid.getFlag())
	fmt.Println("ChainId: ", chid)
	fmt.Println("RPC:", chid.getRPC())
	fmt.Println("Connected Wallets Contract: ", chid.getConnectedWalletsContractAddress())
	fmt.Println("Virtual Accounts Contract: ", chid.getVirtualAccountsContractAddress())
}

func PrintChainIdsHelp() {
	fmt.Println("Available Chains:")
	PrintChainData(BTT_TESTNET_ID)
	PrintChainData(BTT_MAINNET_ID)
}

func GetAssignmentsCountByNetwork(network ChainIds, assignments *[]AssignmentsPerNetwork) int {
	for _, v := range *assignments {
		if ChainIds(v.Network) == network {
			return v.Jobs
		}
	}
	return 0
}

func SumAssignmentsCountFromFlags(chids []string, assignments *[]AssignmentsPerNetwork) (int, []ChainIds) {
	// Calculate how many jobs I can do based on the assignments count per network
	// Chain ordering is an array of instructions to be passed off to the goroutines about what network they should be using
	sum := 0
	var chain_ordering []ChainIds
	for _, v := range chids {
		switch v {
		case BTT_TESTNET_ID.getFlag():
			{
				count := GetAssignmentsCountByNetwork(BTT_TESTNET_ID, assignments)
				sum += count

				for i := 0; i < count; i++ {
					chain_ordering = append(chain_ordering, BTT_TESTNET_ID)
				}

			}
		case BTT_MAINNET_ID.getFlag():
			{
				count := GetAssignmentsCountByNetwork(BTT_MAINNET_ID, assignments)
				sum += count
				for i := 0; i < count; i++ {
					chain_ordering = append(chain_ordering, BTT_MAINNET_ID)
				}

			}
		}
	}
	return sum, chain_ordering
}
