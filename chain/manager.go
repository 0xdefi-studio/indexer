package chain

import (
	"context"
	"crypto/ecdsa"
	"github.com/0xdefi-studio/indexer/dice"
	"github.com/0xdefi-studio/indexer/fomo3d"
	models2 "github.com/0xdefi-studio/indexer/models"
	cmd "github.com/0xdefi-studio/indexer/notification"
	"github.com/ethereum/go-ethereum"
	ABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go.uber.org/zap"
	"log"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var Fomo3dEndTxV1SigHex = crypto.Keccak256Hash([]byte("onEndTx(uint256,uint256,bytes32,address,uint256,uint256,address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256)"))
var DicePlayEventSig = crypto.Keccak256Hash([]byte("Dice_Play_Event(address,uint256)"))

type Manager struct {
	Sugar       *zap.SugaredLogger
	Client      *ethclient.Client
	Fomo3dV1ABI ABI.ABI
	DiceABI     ABI.ABI
	DB          *pg.DB
	FromAddress common.Address
	PrivateKey  *ecdsa.PrivateKey
	ChainId     *big.Int
	NextNonce   uint64
	mu          sync.Mutex
}

func NewWalker(sugar *zap.SugaredLogger, client *ethclient.Client, db *pg.DB) *Manager {
	return &Manager{
		Sugar:  sugar,
		Client: client,
		DB:     db,
	}
}

func (m *Manager) Run(startBlock int64, step int64, addresses []common.Address, market string, discord_web_hook string, telegram_key string, chat_id int64) {
	m.Sugar.Infof("Running 0xdefi-studio indexer for %v from block: %d", market, startBlock)
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				m.Sugar.Infof("Querying latest block number at: %s", t.String())
				//Get the latest block number
				header, err := m.GetHeader()
				if err != nil {
					m.Sugar.Errorf("Get header error: %s", err.Error())
				}

				m.Sugar.Infof("Latest Ethereum Block Number: %d", header)
				if header > startBlock {
					m.Sugar.Infof("Start handling block between %d - %d", startBlock, header)

					i := startBlock
					for i < header {
						i += step
						if i > header {
							i = header
						}

						m.Sugar.Debugf("Getting block %d\n", i)
						err = m.ParseEvent(startBlock, i, addresses, market, discord_web_hook, telegram_key, chat_id)
						if err != nil {
							m.Sugar.Errorf("Parse event error: %s block: %v", err.Error(), startBlock)
						}
						startBlock = i
					}
				} else {
					m.Sugar.Infoln("There is no new block generated")
				}
			}
		}
	}()
	select {}
}

func (m *Manager) InitFomo3dV1ABI(abi string) {
	m.Fomo3dV1ABI, _ = ABI.JSON(strings.NewReader(abi))
}

func (m *Manager) InitDiceABI(abi string) {
	m.DiceABI, _ = ABI.JSON(strings.NewReader(abi))
}

func (m *Manager) SetDicePrivateKey(privateKeyStr string) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err, "Failed to get private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	chainId, err := m.Client.ChainID(context.Background())
	if err != nil {
		log.Fatalln(err, "Failed to get chainID")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	m.FromAddress = fromAddress
	m.PrivateKey = privateKey
	m.ChainId = chainId

	log.Println("From Address: ", m.FromAddress.Hex())
}

func (m *Manager) getNonce() (float64, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

	nonceOnSite, err := m.Client.PendingNonceAt(context.Background(), m.FromAddress)
	if err != nil {
		return 0, err
	}
	nonce := math.Max(float64(nonceOnSite), float64(m.NextNonce))
	m.NextNonce = uint64(nonce) + 1
	return nonce, nil
}

func (m *Manager) CallDiceCallback(toAddress common.Address, requestId *big.Int) (string, error) {
	// max between two uint64

	nonce, err := m.getNonce()
	if err != nil {
		m.Sugar.Errorf("Get nonce error: %s", err.Error())
		return "", err
	}

	gasPrice, err := m.Client.SuggestGasPrice(context.Background())
	if err != nil {
		m.Sugar.Errorf("Get gas price error: %s", err.Error())
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(m.PrivateKey, m.ChainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	instance, err := dice.NewDice(toAddress, m.Client)
	if err != nil {
		m.Sugar.Errorf("New dice error: %s", err.Error())
		return "", err
	}

	var randomBytes [32]byte
	_, err = rand.Read(randomBytes[:])
	if err != nil {
		m.Sugar.Errorf("Random error: %s", err.Error())
		return "", err
	}
	m.Sugar.Infof("Random bytes: %s", hexutil.Encode(randomBytes[:]))

	tx, err := instance.RandomizerCallback(auth, requestId, randomBytes)
	if err != nil {
		m.Sugar.Errorf("RandomizerCallback error: %s", err.Error())
		return "", err
	}

	m.Sugar.Debugf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	return tx.Hash().String(), nil
}

func (m *Manager) GetHeader() (int64, error) {
	header, err := m.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Int64(), nil
}

func (m *Manager) GetBlockByNum(blockNumber int64) (*types.Block, error) {
	bn := new(big.Int).SetInt64(blockNumber)
	block, err := m.Client.BlockByNumber(context.Background(), bn)
	return block, err
}

func (m *Manager) GetBlockTimestamp(blockNumber uint64, blockTimestamps map[int64]uint64) (uint64, error) {
	BlockNumber := int64(blockNumber)
	timestamp, ok := blockTimestamps[BlockNumber]
	if !ok {
		block, err := m.GetBlockByNum(BlockNumber)
		if err != nil {
			m.Sugar.Errorf("Get block %d error\n", BlockNumber)
			return 0, err
		}
		timestamp = block.Header().Time
		blockTimestamps[BlockNumber] = timestamp
		m.Sugar.Debugf("Query block %d!!!!!!", BlockNumber)
	} else {
		m.Sugar.Debugf("Get block from cache!!!!!!")
	}
	return timestamp, nil
}

func (m *Manager) GetFomo3dCurrentRoundInfo(address common.Address) (*big.Int, *big.Int) {
	contract, err := fomo3d.NewFomo3dCaller(address, m.Client)
	if err != nil {
		m.Sugar.Errorf("Get GetFomo3dCurrentRoundInfo %v error\n", err)
		return big.NewInt(0), big.NewInt(0)
	}
	_, _, _, roundEndTime, _, currentPot, _, _, _, _, _, _, _, _, _ := contract.GetCurrentRoundInfo(nil)
	return roundEndTime, currentPot
}

func (m *Manager) ParseEvent(fromBlock int64, toBlock int64, to []common.Address, project string, discord_web_hook string, telegram_key string, chat_id int64) error {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: to,
	}
	m.Sugar.Debugf("Processing %v event from block %d to block %d\n", project, fromBlock, toBlock)
	logs, err := m.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	blockTimestamps := map[int64]uint64{}

	for _, log := range logs {
		timestamp, err := m.GetBlockTimestamp(log.BlockNumber, blockTimestamps)
		if err != nil {
			m.Sugar.Errorf("Error in getting block %s\n", err.Error())
			continue
		}
		if project == fomo3d.FOMO3DV1 {
			switch log.Topics[0].Hex() {
			case Fomo3dEndTxV1SigHex.Hex():
				m.Sugar.Debugf("Get Fomo3d onEndTx event\n")
				onEndTx, err := m.ParseFomo3dOnEndTx(log)
				if err != nil {
					m.Sugar.Errorf("Unpack Fomo3dV1 onEndTx event error: %s\n", err.Error())
					continue
				}
				m.Sugar.Infof("Get Fomo3dV1 onEndTx event, PlayerAddress: %v, KeysBought: %v \n", onEndTx.PlayerAddress, onEndTx.KeysBought)
				_, err = m.InsertEndTxV1(onEndTx, log.BlockNumber, log.TxHash.String(), timestamp)
				if err != nil {
					m.Sugar.Errorf("Insert Fomo3dV1 onEndTx event error: %s\n", err.Error())
					continue
				}
				// current time in seconds

				current := uint64(time.Now().Unix())
				if current-3600 <= timestamp {
					m.Sugar.Infof("In notification discord start")
					cmd.SendDiscordEndTx(discord_web_hook, onEndTx.PlayerAddress.String(), onEndTx.KeysBought, onEndTx.EthIn, timestamp)
					m.Sugar.Infof("In notification telegram start")
					roundEndTime, currentPot := m.GetFomo3dCurrentRoundInfo(to[0])
					restTimeBeforeEnd := big.NewInt(0).Sub(roundEndTime, big.NewInt(int64(current)))
					cmd.SendTelegramEndTx(telegram_key, chat_id, onEndTx.PlayerAddress.String(), onEndTx.KeysBought, onEndTx.EthIn, timestamp, currentPot, restTimeBeforeEnd)
				}

			}
		} else if project == dice.DICE {
			switch log.Topics[0].Hex() {
			case DicePlayEventSig.Hex():
				m.Sugar.Debugf("Get Dice Play event\n")
				playTx, err := m.ParseDicePlayTx(log)
				if err != nil {
					m.Sugar.Errorf("Unpack Dice Play event error: %s\n", err.Error())
					continue
				}
				m.Sugar.Infof("Get Dice Play event, PlayerAddress: %v, RequestId: %v \n", playTx.PlayerAddress, playTx.RequestId)
				playTxInstance, err := m.InsertDicePlayTx(playTx, log.Address.String(), log.BlockNumber, log.TxHash.String(), timestamp)
				if err != nil {
					m.Sugar.Errorf("Insert Dice Play onEndTx event error: %s\n", err.Error())
					continue
				}
				callbackTxHash, err := m.CallDiceCallback(log.Address, playTx.RequestId)
				if err != nil {
					m.Sugar.Errorf("Call Dice callback error in parser: %s\n", err.Error())
					continue
				}
				m.Sugar.Debugf("Call Dice callback success in parser\n")
				// business logic of liquidated
				playTxInstance.IsSuccess = 1
				playTxInstance.CallbackTransactionHash = callbackTxHash
				_, err = m.DB.Model(playTxInstance).Column("is_success", "callback_transaction_hash").WherePK().Update()
				if err != nil {
					m.Sugar.Errorf("Update Dice Play playTxInstance is_success update error: %s\n", err.Error())
				}

				// current time in seconds
			}
		}

	}
	return nil
}

func (m *Manager) CreateFomo3dV1Schema() error {
	m.Sugar.Infof("Creating Fomo3dV1Schema schema...\n")
	models := []interface{}{
		(*models2.Fomo3dOnEndTx)(nil),
	}

	for _, model := range models {
		err := m.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) CreateDiceSchema() error {
	m.Sugar.Infof("Creating CreateDiceSchema schema...\n")
	models := []interface{}{
		(*models2.DicePlayTx)(nil),
	}

	for _, model := range models {
		err := m.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) CreateTelegramSchema() error {
	m.Sugar.Infof("Creating CreateTelegramSchema schema...\n")
	models := []interface{}{
		(*models2.TelegramChat)(nil),
	}

	for _, model := range models {
		err := m.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
