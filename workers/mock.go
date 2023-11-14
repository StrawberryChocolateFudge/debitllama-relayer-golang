package workers

import (
	"encoding/json"
	"fmt"
)

func GetMockAssignment() (DirectDebitAssignment, error) {

	body := string(`{"assignment":"8450","abi":{"inputs":[{"internalType":"uint256[8]","name":"proof","type":"uint256[8]"},{"internalType":"bytes32[2]","name":"hashes","type":"bytes32[2]"},{"internalType":"address","name":"payee","type":"address"},{"internalType":"uint256[4]","name":"debit","type":"uint256[4]"}],"name":"directdebit","outputs":[],"stateMutability":"nonpayable","type":"function"},"network":"0x405","contractAddress":"0xF75515Df5AC843a8B261E232bB890dc2F75A4066","inputs":{"proof":["14863844320371167676309284296597351121006620090184411318728382870240010697078","10855371166930166232111125206361822205977298563671104411965703714580401874859","9444832039570675211829556348864010516391546185448332115979109057229632497471","12573972319946457624074445395367438491473342933385649874608246079908543175369","18544736756372325080629522962967934163715467926613942727764645010377433228512","10672401731655368778630590256406361206184111624869588032660847712252110378244","17813220734415559111106795881606531452972365167108131880550594239055866124027","19600474628273165249389890945496117850568974482433879548543250518739458370791"],"hashes":["0x1362e794b206fa1490399b33e55c17118ee371b884ff8ec7feeebd5c4f460db6","0x25f4251963adc4a13a1e58af056422cd5f934fb6c2fdc841c298366935270316"],"payee":"0xD97F13b8fd8a54434F7Bd7981F0D6C82EA1b59F3","debit":["10000000000000000000","10","10","10000000000000000000"]},"gasPrice":"9000000000000000","gasLimit":"512181","maxFeePerGas":"0","maxPriorityFeePerGas":"0"}`)

	var assignment DirectDebitAssignment

	err := json.Unmarshal([]byte(body), &assignment)
	if err != nil {
		return assignment, fmt.Errorf("unable to parse assignment")
	}
	return assignment, nil

}
