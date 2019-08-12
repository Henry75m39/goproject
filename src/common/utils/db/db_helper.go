package db

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

var err error

var Instance *leveldb.DB

func InitLevelDB() *leveldb.DB {
	//initialize levelDB storage files under leveldb directory.
	Instance, err = leveldb.OpenFile("../goapp/persistent/leveldb", nil)
	if err != nil {
		fmt.Println("LevelDB errors: %s", err)
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
