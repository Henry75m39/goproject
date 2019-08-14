package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func GetJsonContents() (contents []byte) {

	const dataFile = "../../../common/configs/config.json"

	_, filename, _, _ := runtime.Caller(1)

	dataPath := path.Join(path.Dir(filename), dataFile)

	fmt.Println(dataPath)

	f, err := os.Open(dataPath)
	if err != nil {
		fmt.Println("open JSON config file error: " + err.Error())
		return nil
	}

	contents, err = ioutil.ReadAll(f)

	if err != nil {
		fmt.Println("Read JSON config file error: " + err.Error())
		return nil
	}
	return contents
}
