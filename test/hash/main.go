package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 示例：生成一个address类型indexed参数的topic hash
	address := common.HexToAddress("0x1234567890123456789012345678901234567890")
	topic := crypto.Keccak256Hash(address.Bytes()).Hex()
	fmt.Println("Address topic:", topic)

	// 示例：生成一个uint256类型indexed参数的topic hash
	number := uint256Example("1234567890")
	topic = crypto.Keccak256Hash(number).Hex()
	fmt.Println("Uint256 topic:", topic)

	// 示例：生成一个string类型indexed参数的topic hash
	str := "bafkreih6n5g5w4y6u7uvc4mh7jhjm7gidmkrbbpi7phyiyg54gplvngcpn"
	topic = crypto.Keccak256Hash([]byte(str)).Hex()
	fmt.Println("String topic:", topic)
}

// 对于uint256类型，需要转换为32字节的大端表示
func uint256Example(numStr string) []byte {
	// 这里简化处理，实际应用中需要使用big.Int正确处理大数
	// 这里假设输入是十进制字符串
	bytes := make([]byte, 32)
	copy(bytes[32-len(numStr):], numStr)
	return bytes
}
