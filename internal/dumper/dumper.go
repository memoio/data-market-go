package dumper

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/data-market/internal/database"
	"github.com/data-market/internal/go-contracts/accountdid"
	"github.com/data-market/internal/go-contracts/filedid"
	"github.com/data-market/internal/logs"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	com "github.com/memoio/contractsv2/common"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

var (
	// blockNumber = big.NewInt(0)
	logger = logs.Logger("dumper")

	AccountDID_ABI = accountdid.AccountdidABI
	FileDID_ABI    = filedid.FiledidABI
)

type Dumper struct {
	// chain
	endpoint string

	// the instance contract address
	instance_ADDR common.Address

	// contract address
	accountdid_ADDR common.Address
	filedid_ADDR    common.Address

	// contract abi
	accountdid_ABI abi.ABI
	filedid_ABI    abi.ABI

	fromBlock *big.Int

	// store           MapStore
	eventNameMap map[common.Hash]string
	indexedMap   map[common.Hash]abi.Arguments
}

// init a dumper with the env
func NewDumper(env string) (dumper *Dumper, err error) {
	// new Dumper
	dumper = &Dumper{}

	// init dumper
	logger.Debug("init dumper..")
	dumper.Init(env)

	return dumper, nil
}

// sync db with block chain every 10 sec
func (d *Dumper) Subscribe(ctx context.Context) {
	// dial chain
	logger.Info("connect chain")
	client, err := ethclient.DialContext(context.TODO(), d.endpoint)
	if err != nil {
		log.Fatalf("connect chain failed when subscribe: %s", err)
	}
	defer client.Close()

	for {
		err := d.Dump(client)
		if err != nil {
			panic(err)
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(10 * time.Second):
		}
	}
}

// dump all event logs of blocks into db
func (d *Dumper) Dump(client *ethclient.Client) error {

	// get current chain block number
	chainBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		logger.Debug("get block number error:", err)
		return err
	}
	logger.Info("get current block number from chain: ", chainBlock)

	// if no new chain block, return
	if d.fromBlock.Cmp(new(big.Int).SetUint64(chainBlock)) >= 0 {
		logger.Info("no new chain block, waiting..")
		return nil
	}

	logger.Debug("dump from block: ", d.fromBlock)

	// filter event logs for all contracts
	logger.Debug("filter event logs")
	events, err := client.FilterLogs(context.TODO(), ethereum.FilterQuery{
		FromBlock: d.fromBlock,
		Addresses: []common.Address{d.accountdid_ADDR, d.filedid_ADDR},
	})
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	logger.Debug("event log number:", len(events))

	tmp := 0
	// parse each event
	for _, event := range events {
		// topic0 is the event name
		eventName, ok1 := d.eventNameMap[event.Topics[0]]
		if !ok1 {
			continue
		}

		// handle each event log
		switch eventName {
		case "CreateDID":
			logger.Debug("==== Handle CreateDID Event")
			err = d.HandleCreateDID(event)
			if err != nil {
				logger.Debug("handle createdid error: ", err.Error())
			}
		case "RegisterMfileDid":
			logger.Debug("==== Handle RegisterMfileDid Event")
		case "BuyRead":
			logger.Debug("==== Handle BuyRead Event")
			err = d.HandleBuyRead(event)
			if err != nil {
				logger.Debug("handle buyread error: ", err.Error())
			}

		default:
			continue
		}

		// for test, 10 records only
		tmp++
		if tmp > 10 {
			break
		}
	}

	// update from block to current chain block
	d.fromBlock = big.NewInt(int64(chainBlock))

	// update from block into db
	logger.Debug("update from block into db: ", d.fromBlock)
	err = database.SetBlockNumber(d.fromBlock.Int64())
	if err != nil {
		return err
	}

	return nil
}

// unpack a log
func (d *Dumper) unpack(log types.Log, ABI abi.ABI, out interface{}) error {
	// get event name from map with hash
	eventName := d.eventNameMap[log.Topics[0]]

	// get all topics
	indexed := d.indexedMap[log.Topics[0]]

	// parse data
	err := ABI.UnpackIntoInterface(out, eventName, log.Data)
	if err != nil {
		return err
	}

	// parse topic
	err = abi.ParseTopics(out, indexed, log.Topics[1:])
	if err != nil {
		return err
	}

	return nil
}

// get did contract address from instance, and get endpoint
func (d *Dumper) Init(env string) (err error) {
	// init map
	d.eventNameMap = make(map[common.Hash]string)
	d.indexedMap = make(map[common.Hash]abi.Arguments)

	// get instance address and chain ep
	instAddr, ep := com.GetInsEndPointByChain(env)
	logger.Debug("instance address:", instAddr)
	logger.Debug("endpoint:", ep)

	// save endpoint and instance
	d.endpoint = ep
	d.instance_ADDR = instAddr

	// get client
	client, err := ethclient.DialContext(context.Background(), ep)
	if err != nil {
		return err
	}

	// get instance
	instIns, err := inst.NewInstance(instAddr, client)
	if err != nil {
		return err
	}

	// get accountdid address
	d.accountdid_ADDR, err = instIns.Instances(&bind.CallOpts{}, 30)
	if err != nil {
		return err
	}
	logger.Debug("accountDID addr:", d.accountdid_ADDR)

	// get filedid address
	d.filedid_ADDR, err = instIns.Instances(&bind.CallOpts{}, 34)
	if err != nil {
		return err
	}
	logger.Debug("fileDID addr:", d.filedid_ADDR)

	// set accountdid abi
	d.accountdid_ABI, err = abi.JSON(strings.NewReader(AccountDID_ABI))
	if err != nil {
		return err
	}

	// set filedid abi
	d.filedid_ABI, err = abi.JSON(strings.NewReader(FileDID_ABI))
	if err != nil {
		return err
	}

	// group all abi together
	ABIs := []abi.ABI{d.accountdid_ABI, d.filedid_ABI}

	// parse all abi for event and topic
	logger.Debug("parse event and topics in all abi")
	for _, ABI := range ABIs {
		// each event
		for name, event := range ABI.Events {
			// save event in dumper
			d.eventNameMap[event.ID] = name
			var indexed abi.Arguments
			// each topic
			for _, arg := range ABI.Events[name].Inputs {
				if arg.Indexed {
					indexed = append(indexed, arg)
				}
			}
			// save topics for each event in dumper
			d.indexedMap[event.ID] = indexed
		}
	}

	// get block number from db
	logger.Debug("getting block number from db")
	blockNumber, err := database.GetBlockNumber()
	if err != nil {
		blockNumber = 0
	}
	logger.Debug("block number: ", blockNumber)

	// set from block number for dumper
	d.fromBlock = big.NewInt(blockNumber)

	logger.Debug("init complete")

	return nil
}
