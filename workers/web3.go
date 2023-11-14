package workers

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient(chainid ChainIds) (*ethclient.Client, error) {
	url := chainid.getRPC()

	client, err := ethclient.Dial(string(url))
	if err != nil {
		log.Fatalf("Unable to connect to the network url %s \n", url)
	} else {
		fmt.Printf("Success! You are connected to %s \n", url)
	}

	return client, err
}

func CreateKs(password string) {
	ks := GetKS()
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New Keystore created! Address:")
	fmt.Println(account.Address.Hex())
}

func GetKS() *keystore.KeyStore {
	return keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
}

func GetAccountFromKs(keystore *keystore.KeyStore) accounts.Account {
	accounts := keystore.Accounts()
	return accounts[0]
}

func formatEther(wei *big.Int) string {
	// This will accept wei and format it as eth
	//Formatted as explained in https://goethereumbook.org/account-balance/
	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue.String()
}

func GetRelayerBalances() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()
	address := accounts[0].Address

	fmt.Println("Address: ", address.String())

	// get the secret key and create a new ethers wallet with the ethClient
	BTTTestnetClient, err := getClient(BTT_TESTNET_ID)

	if err != nil {
		panic("BTT Testnet client  can't connect")
	}

	BTTMainnetClient, err := getClient(BTT_MAINNET_ID)

	if err != nil {
		panic("BTT Mainnet client can't connect")
	}

	BTTTestnetBalance, err := BTTTestnetClient.BalanceAt(context.Background(), address, nil)

	if err != nil {
		panic("Btt testnet balance error")
	}
	fmt.Println("BTT Testnet Balance: ", formatEther(BTTTestnetBalance), "BTT")

	BTTMainnetBalance, err := BTTMainnetClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		panic("BTT mainnet balance error")
	}
	fmt.Println("BTT Mainnet Balance", formatEther(BTTMainnetBalance), "BTT")

}

func (a *DirectDebitAssignment) RelayTx(password string) (int, error) {
	client, err := getClient(ChainIds(a.Network))

	if err != nil {
		fmt.Println("Client error")
		return 0, err
	}

	ks := GetKS()

	senderAccount := GetAccountFromKs(ks)
	senderAddress := senderAccount.Address

	ksJson, err := os.ReadFile(senderAccount.URL.Path)
	if err != nil {
		fmt.Println("read file error")
		return 0, err
	}

	contractAddress := a.ContractAddress

	fmt.Println("contract address ", contractAddress)

	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, eof := new(big.Int).SetString(a.GasPrice, 10)
	if !eof {
		return 0, fmt.Errorf("unable to parse a bigint")
	}
	fmt.Println("gas price", gasPrice)
	fmt.Println(a.Network)

	bigintChainId, err := hexutil.DecodeBig(a.Network)
	fmt.Println("bigintchainid ")
	fmt.Println(bigintChainId)
	if err != nil {
		return 0, err
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(ksJson)), password, bigintChainId)

	if err != nil {
		return 0, err
	}
	fmt.Println("Auth transactor with chain id")

	gasLimit, err := strconv.ParseInt(a.GasLimit, 10, 64)

	if err != nil {
		return 0, err
	}

	gasPrice, eof = new(big.Int).SetString(a.GasPrice, 10)
	if !eof {
		return 0, fmt.Errorf("unable to parse a bigint")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // In WEI
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = gasPrice         // In WEI

	ddcontract, err := NewDirectDebit(common.HexToAddress(a.ContractAddress), client)
	if err != nil {
		return 0, err
	}
	fmt.Println("Direct debit contract", ddcontract)

	var hashes [2][32]byte
	hashes[0] = [32]byte([]byte(a.Inputs.Hashes[0]))
	hashes[1] = [32]byte([]byte(a.Inputs.Hashes[1]))
	fmt.Println("hashes")
	var debit [4]*big.Int
	debit[0], eof = new(big.Int).SetString(a.Inputs.Debit[0], 10)

	if !eof {
		return 0, fmt.Errorf("unable to parse string to bigint")
	}

	debit[1], eof = new(big.Int).SetString(a.Inputs.Debit[1], 10)
	if !eof {
		return 0, fmt.Errorf("unable to parse string to bigint")
	}

	debit[2], eof = new(big.Int).SetString(a.Inputs.Debit[2], 10)
	if !eof {
		return 0, fmt.Errorf("unable to parse string to bigint")
	}

	debit[3], eof = new(big.Int).SetString(a.Inputs.Debit[3], 10)
	if !eof {
		return 0, fmt.Errorf("unable to parse string to bigint")
	}
	fmt.Println("sending tx")
	tx, err := ddcontract.Directdebit(
		auth, PackToSolidityProof(a.Inputs.Proof),
		hashes,
		common.HexToAddress(a.Inputs.Payee),
		debit,
	)
	fmt.Println("sent tx")
	if err != nil {
		return 0, err
	}

	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
	return 0, nil
}
