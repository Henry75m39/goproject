package wrapper

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/helper"
	"GRM/src/common/utils/log"
	"GRM/src/tms-srv/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Status   string     `json:"status"`
	Response []Response `json:"response"`
}
type Response struct {
	Status   string `json:"status"`
	Response int    `json:"response"`
}

func CreateProjectGroup(resp entity.ProjectGroup) (status string, err error) {
	var c configs.WSConfig
	var contents []byte
	contents = helper.GetJsonContents()
	//unmarshal the json to struct object
	err = json.Unmarshal([]byte(contents), &c)
	if err != nil {
		return
	}
	//construct login Restful API path
	apiPath := c.WSTMSProjectGroupsCreateAPI
	logger := log.Instance()

	jsonValue, err := json.Marshal(&resp)
	if err != nil {
		logger.Error("Error", zap.Any("createProjectGroup", "json data input error"))
		return
	}
	req, _ := http.NewRequest("POST", apiPath, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := ioutil.ReadAll(rsp.Body)

	status = parseResult(data)

	fmt.Println(status)

	return status, err
}

func parseResult(jsonResult []byte) (status string) {
	var r Result
	_ = json.Unmarshal(jsonResult, &r)
	status = r.Status
	return status
}
