package workers

import (
	"math/big"
)

func PackToSolidityProof(proof [8]string) [8]*big.Int {
	pia0, _ := new(big.Int).SetString(proof[0], 10)
	pia1, _ := new(big.Int).SetString(proof[1], 10)
	pib01, _ := new(big.Int).SetString(proof[2], 10)
	pib00, _ := new(big.Int).SetString(proof[3], 10)
	pib11, _ := new(big.Int).SetString(proof[4], 10)
	pib10, _ := new(big.Int).SetString(proof[5], 10)
	pic0, _ := new(big.Int).SetString(proof[6], 10)
	pic1, _ := new(big.Int).SetString(proof[7], 10)

	return [8]*big.Int{pia0, pia1, pib01, pib00, pib11, pib10, pic0, pic1}
}
