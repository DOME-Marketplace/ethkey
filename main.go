// newnodekey generates a new nodekey to be used in Ethereum-compatible networks
package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		utils.Fatalf("could not generate key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Ethereum Private Key (keep this only for you!): %v\n", hexutil.Encode(privateKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Ethereum Address:", address)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Ethereum Public Key:", hexutil.Encode(publicKeyBytes))

	fmt.Printf("enode: %x\n", crypto.FromECDSAPub(&privateKey.PublicKey)[1:])

}
