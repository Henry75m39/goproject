package projects

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

func CreateForTask(resp entity.Projects) (status []string, err error) {
	var c configs.WSConfig
	var contents []byte

	logger := log.Instance()
	token := helper.GetToken()

	contents = helper.GetJsonContents()
	//unmarshal the json to struct object
	err = json.Unmarshal([]byte(contents), &c)
	if err != nil {
		return
	}
	//construct login Restful API path
	apiPath := c.WSProjects.WSCreateForTask

	if apiPath == "" {
		logger.Error("Error", zap.Any("createProjectForTask", "REST API cannot read from JSON config file "))
		return
	}
	jsonValue, err := json.Marshal(&resp)
	if err != nil {
		logger.Error("Error", zap.Any("createProjectForTask", "json data input error"))
		return
	}

	req, _ := http.NewRequest("POST", apiPath, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("token", *token)

	client := &http.Client{}
	rsp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	data, _ := ioutil.ReadAll(rsp.Body)
	status = helper.ParseCommonJSONResult(data)
	return status, err
}
