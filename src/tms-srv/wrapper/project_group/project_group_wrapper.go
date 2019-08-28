package project_group

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

func CreateProjectGroup(projectGroup entity.ProjectGroup) ([]byte, error) {

	var c configs.WSConfig
	var contents []byte

	logger := log.Instance()
	token := helper.GetToken()

	contents = helper.GetJsonContents()
	//unmarshal the json to struct object
	err := json.Unmarshal([]byte(contents), &c)
	if err != nil {
		logger.Error("Error", zap.Any("createProjectGroup", "json data input error"))
	}
	//construct login Restful API path
	apiPath := c.WSProjectGroup.ProjectGroupCreate
	if apiPath == "" {
		logger.Error("Error", zap.Any("createProjectGroup", "REST API cannot read from JSON config file "))
	}
	jsonValue, err := json.Marshal(&projectGroup)
	if err != nil {
		logger.Error("Error", zap.Any("createProjectGroup", "json data input error"))
	}
	req, _ := http.NewRequest("POST", apiPath, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("token", *token)

	client := &http.Client{}
	rsp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}
	//Reads the body of response and closes the body reader when done reading it.
	data, _ := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()

	dynamicStrutForJson := make(map[string]interface{})
	err = json.Unmarshal(data, &dynamicStrutForJson)
	if err != nil {
		logger.Error("Error", zap.Any("cancelProject", "REST API cannot unmarshal correctly."))

	}
	//make json looks beautiful.
	beautifulJson, _ := json.MarshalIndent(dynamicStrutForJson, "", "\t")

	//testing the output result into a json file.
	//err = ioutil.WriteFile("post.json", []byte(beautifulJson), 0644)
	return beautifulJson, err
	//status = helper.ParseCommonJSONResult(data)
	//return status, err
}
