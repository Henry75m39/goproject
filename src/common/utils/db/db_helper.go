package db

import (
	"GRM/src/common/utils/log"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
)

var err error

var Instance *leveldb.DB

func InitLevelDB() *leveldb.DB {
	logger := log.Instance()
	//initialize levelDB storage files under leveldb directory.
	//Todo: Move the data persist path can be configurable in JSON file.
	Instance, err = leveldb.OpenFile("data/goapp/leveldb", nil)
	if err != nil {
		fmt.Println("LevelDB errors: %s", err)
		logger.Error("ERROR:", zap.Any("DBERROR:", err))
		panic(err)
	}
	return Instance
}

func ValidateKeyInDb(key string) bool {
	isKeyExist, err := Instance.Has([]byte(key), nil)
	if err != nil {
		fmt.Println("Validate Key from LevelDB errors: %s", err)
	}
	return isKeyExist
}
