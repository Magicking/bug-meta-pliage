package internal

import (
	"context"
	"crypto/ecdsa"
	"log"
	"strings"

	ethtk "github.com/Magicking/gethitihteg"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var contractKey key = 4242
var _key *ecdsa.PrivateKey
var contractAddr common.Address
var DAO *DAOTracabiliteSession

func InitContract(ctx context.Context, auth *bind.TransactOpts) (common.Address, interface{}) {
	rawContract := common.FromHex(DAOTracabiliteBin)
	client := CCFromContext(ctx)
	contractABI, err := abi.JSON(strings.NewReader(DAOTracabiliteBin))
	addr, tx, c, err := ethtk.DeployContract(auth, contractABI, rawContract, client)
	if err != nil {
		log.Fatalf("InitContract: %v", err)
	}
	log.Printf("%v\nDeployed contract at %s", tx, addr.String())
	return addr, c
}

func Init(ctx context.Context, contract_address, private_key string) {
	nc := CCFromContext(ctx)
	key, err := crypto.HexToECDSA(private_key)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	_key = key
	auth := bind.NewKeyedTransactor(key)
	ctct_addr := common.HexToAddress(contract_address)
	if contract_address == "" {
		ctct_addr, _ = InitContract(ctx, auth)
	}
	ctct, err := NewDAOTracabilite(ctct_addr, nc)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	setContextValue(ctx, contractKey, ctct)
	contractAddr = ctct_addr
	DAO = &DAOTracabiliteSession{Contract: ctct,
		CallOpts:     bind.CallOpts{},
		TransactOpts: *auth,
	}
}
