package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"

	ethUnit "github.com/DeOne4eg/eth-unit-converter"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	crypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient(chainid ChainIds) (*ethclient.Client, error) {
	url := chainid.getRPC()

	provider, err := ethclient.Dial(string(url))
	if err != nil {
		log.Fatalf("Unable to connect to the network url %s \n", url)
	} else {
		fmt.Printf("Success! You are connected to %s \n", url)
	}

	return provider, err
}

func getKeysFromEnv() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	key := os.Getenv("RELAYER_PRIVATEKEY")

	privatekey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("Invalid key")
	}

	return privatekey, publicKeyECDSA
}

func CreateKs(password string) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New Keystore created! Address:")
	fmt.Println(account.Address.Hex())
}

func OpenKs(password string, filename string) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	accounts := ks.Accounts()
	fmt.Println(accounts[0].Address)
}

func parseEther(input string) *big.Int {
	value, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic(err)
	}
	unit := ethUnit.NewEther(big.NewFloat(value))
	return unit.Wei()
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
