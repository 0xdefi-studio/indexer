package cmd

import (
	"fmt"
	"github.com/0xdefi-studio/indexer/bankroll"
	"github.com/0xdefi-studio/indexer/chain"
	"github.com/0xdefi-studio/indexer/dice"
	"github.com/0xdefi-studio/indexer/fomo3d"
	models2 "github.com/0xdefi-studio/indexer/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-pg/pg/v10"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
	"reflect"
)

func getStartBlock(db *pg.DB, market string) int64 {
	switch market {
	case fomo3d.FOMO3DV1:
		var obj models2.Fomo3dOnEndTx
		err := db.Model(&obj).Order("block_num DESC").First()
		if err != nil {
			return FOMO3D_START_BLOCK
		}
		return int64(obj.BlockNum)
	case dice.DICE:
		var obj models2.DicePlayTx
		err := db.Model(&obj).Order("block_num DESC").First()
		if err != nil {
			return DICE_START_BLOCK
		}
		return int64(obj.BlockNum)
	case bankroll.BANKROLL:
		var obj models2.BankrollTx
		err := db.Model(&obj).Order("block_num DESC").First()
		if err != nil {
			return BANKROLL_START_BLOCK
		}
		return int64(obj.BlockNum)
	default:
		return 0
	}
}

func IndexBase(
	project string,
	createSchemeFuncName string,
	abiPath string,
	initABIFuncName string,
	API string,
) (*chain.Manager, int64, int64) {
	aa := zap.NewDevelopmentEncoderConfig()
	aa.EncodeLevel = zapcore.CapitalColorLevelEncoder

	DebugLevel := zapcore.DebugLevel
	if !IsDebug {
		DebugLevel = zapcore.InfoLevel
	}
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(aa),
		zapcore.AddSync(colorable.NewColorableStdout()),
		DebugLevel,
	))
	defer logger.Sync() // flushes buffer, if any

	// init log and eth client
	sugar := logger.Sugar()

	client, err := ethclient.Dial(API)
	if err != nil {
		sugar.Errorf("init ethereum client error: %s", err.Error())
		os.Exit(1)
	}

	// init database
	Addr := fmt.Sprintf("%v:%v", DBHost, DBPort)
	db := pg.Connect(&pg.Options{
		User:     DBUser,
		Password: DBPassword,
		Database: DBName,
		Addr:     Addr,
	})
	walker = chain.NewWalker(sugar, client, db)

	// init table
	createSchemeFunc := reflect.ValueOf(walker).MethodByName(createSchemeFuncName)
	value := createSchemeFunc.Call(nil)
	err, _ = value[0].Interface().(error)
	if err != nil {
		sugar.Errorf("init %v tables error: %s", project, err.Error())
		os.Exit(1)
	}

	// read abi
	ABI, err := ioutil.ReadFile(abiPath)
	if err != nil {
		sugar.Errorf("read %v abi error: %s", project, err.Error())
		os.Exit(1)
	}

	// walker.InitNFTFIABI(string(nftfiABI))
	InitAbiFunc := reflect.ValueOf(walker).MethodByName(initABIFuncName)
	in := []reflect.Value{reflect.ValueOf(string(ABI))}
	InitAbiFunc.Call(in)
	startBlock := getStartBlock(db, project)
	return walker, startBlock, Step
}
