package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func GetJsonContents() (contents []byte) {

	const dataFile = "../../common/configs/config.json"

	_, filename, _, _ := runtime.Caller(1)

	dataPath := path.Join(path.Dir(filename), dataFile)

	fmt.Println(dataPath)

	f, err := os.Open(dataPath)

	contents, err = ioutil.ReadAll(f)

	if err != nil {
		fmt.Println("open config file error: " + err.Error())
		return nil
	}
	return contents
}
