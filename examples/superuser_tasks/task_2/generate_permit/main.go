package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func main() {
	privateKey, err := crypto.HexToECDSA(
		"0e5ff7fe55905e6f12ae87aa05703d4b4747c8251b025c8e7e07b257fdbdd557",
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	publicKey := privateKey.Public()
	publicKeyOk, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	owner := crypto.PubkeyToAddress(*publicKeyOk)
	spender := common.HexToAddress("0x655328270123049AE1337C05B640213b1Ad081aB")
	verifyingContract := common.HexToAddress("0xA8AA3b44Afb4aA219ee12Ca4CE93C2081767BfDb")
	nonce := 1       // Must fetch from blockchain (using polygonscan for example)
	chainId := 80002 // 80002 for Amoy, 137 for Polygon
	d := time.Now().Add(60 * 60 * time.Second).Unix()
	value := 8000 // Two decimal places

	val := big.NewInt(0)
	val.SetString(fmt.Sprintf("%d0000000000000000", value), 10)

	log.Println(val)

	r, s, v, deadline, h, err := GenerateSignedPermit(
		"BRLA Token",
		"1",
		false,
		owner,
		spender,
		verifyingContract,
		int64(chainId),
		val,
		int64(nonce),
		int64(d),
		privateKey,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===========> permit order -> r, s, v, dealine")
	log.Println(r, s, v, deadline)

	rByte, err := hexutil.Decode(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	sByte, err := hexutil.Decode(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	sig := append(rByte, sByte...)
	sig = append(sig, v-27)

	pKeyRes, err := crypto.SigToPub(h.Bytes(), sig)
	if err != nil {
		log.Fatal(err.Error())
	}

	addr := crypto.PubkeyToAddress(*pKeyRes)

	log.Println(addr.String() == owner.String())

	log.Println(owner.String())
}

func GeneratePermitHash(
	contractName string,
	contractVersion string,
	legacyPermit bool,
	owner common.Address,
	spender common.Address,
	verifyingContract common.Address,
	chainId int64,
	value *big.Int,
	nonce int64,
	deadline int64,
) (common.Hash, error) {
	val := math.HexOrDecimal256(*value)

	var domain apitypes.TypedDataDomain
	var typesPermit apitypes.Types

	if !legacyPermit {
		domain = apitypes.TypedDataDomain{
			Name:              contractName,
			Version:           contractVersion,
			ChainId:           math.NewHexOrDecimal256(chainId),
			VerifyingContract: verifyingContract.String(),
		}
		typesPermit = apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Permit": []apitypes.Type{
				{Name: "owner", Type: "address"},
				{Name: "spender", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
			},
		}
	} else {
		domain = apitypes.TypedDataDomain{
			Name:              contractName,
			Version:           contractVersion,
			Salt:              common.HexToHash(strconv.FormatInt(chainId, 16)).String(),
			VerifyingContract: verifyingContract.String(),
		}
		typesPermit = apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "verifyingContract", Type: "address"},
				{Name: "salt", Type: "bytes32"},
			},
			"Permit": []apitypes.Type{
				{Name: "owner", Type: "address"},
				{Name: "spender", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "deadline", Type: "uint256"},
			},
		}
	}

	signerData := apitypes.TypedData{
		Types:       typesPermit,
		PrimaryType: "Permit",
		Domain:      domain,
		Message: apitypes.TypedDataMessage{
			"owner":    owner.String(),
			"spender":  spender.String(),
			"value":    &val,
			"nonce":    math.NewHexOrDecimal256(nonce),
			"deadline": math.NewHexOrDecimal256(deadline),
		},
	}

	log.Println(signerData.Map())
	log.Println(nonce, deadline, val)

	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}

	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return common.Hash{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := common.BytesToHash(crypto.Keccak256(rawData))

	log.Println(hash.String())
	return hash, nil
}

func GenerateSignedPermit(
	contractName string,
	contractVersion string,
	legacyPermit bool,
	owner common.Address,
	spender common.Address,
	verifyingContract common.Address,
	chainId int64,
	value *big.Int,
	nonce int64,
	deadline int64,
	privateKey *ecdsa.PrivateKey,
) (r string, s string, v uint8, dl int64, hash common.Hash, err error) {
	hash, err = GeneratePermitHash(
		contractName,
		contractVersion,
		legacyPermit,
		owner,
		spender,
		verifyingContract,
		chainId,
		value,
		nonce,
		deadline,
	)
	if err != nil {
		return
	}

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	r = hexutil.Encode(signatureBytes[:32])
	s = hexutil.Encode(signatureBytes[32:64])
	v = uint8(int(signatureBytes[64])) + 27
	dl = deadline

	return
}
