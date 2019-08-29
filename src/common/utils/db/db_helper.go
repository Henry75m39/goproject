package db

import (
	"GRM/src/common/utils/log"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
)

//Global db instance variable
var Instance *leveldb.DB

//LevelDB is designed that way and it doesn't allow more than a single instance of the database to be open.
// All of the options are for a single process.

//The init function (which is called automatically for every package) to initialize it.

func init() {
	var err error
	logger := log.Instance()
	//initialize levelDB storage files under leveldb directory.
	//Todo: Move the data persist path can be configurable in JSON file.

	//Instance = new(leveldb.DB)  this may not necessary, it only for invalid memory error
	Instance, err = leveldb.OpenFile("data/goapp/leveldb", nil)
	if err != nil {
		fmt.Println("LevelDB errors: %s", err)
		logger.Error("ERROR:", zap.Any("DBERROR:", err))
		panic(err)
	}
	defer Instance.Close()
	//return Instance
}

func ValidateKeyInDb(key string) bool {
	isKeyExist, err := Instance.Has([]byte(key), nil)
	if err != nil {
		fmt.Println("Validate Key from LevelDB errors: %s", err)
	}
	return isKeyExist
}
