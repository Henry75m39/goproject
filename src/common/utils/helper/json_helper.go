package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strconv"
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

//Global variable define for storing REST response
// results without knowing JSON schema .
var Runes []string

func ParseJSONResult(v interface{}) []string {
	switch vv := v.(type) {
	case string:
		fmt.Println("Status:", vv)
		Runes = append(Runes, "Status:")
		Runes = append(Runes, vv)
	case float64:
		fmt.Println("is: ", vv)
		Runes = append(Runes, "Error Code:")
		s1 := strconv.FormatInt(int64(vv), 10)
		Runes = append(Runes, s1)
	case []interface{}:
		fmt.Println("is an array")
		for i, u := range vv {
			fmt.Print(i, " ")
			ParseJSONResult(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			ParseJSONResult(u)
		}
	default:
		fmt.Println("Unknown type")
		s2 := "Unknown type error"
		Runes = append(Runes, s2)
		return Runes
	}
	return Runes
}

func ParseCommonJSONResult(jsonResult []byte) (status []string) {
	var f interface{}
	err := json.Unmarshal(jsonResult, &f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := f.(map[string]interface{})
	status = ParseJSONResult(m)
	return status
}
