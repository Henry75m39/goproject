package ws_tasks

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/helper"
	"GRM/src/common/utils/httpx"
	"GRM/src/tms-srv/entity"
	"encoding/json"
	"errors"
)

func GetAllTasks(tasks entity.Tasks) ([]byte, error) {
	var c configs.WSConfig
	var contents []byte

	params := make(map[string]string)

	if tasks.Ids != "" {
		params["ids"] = tasks.Ids
	}
	if tasks.Fields != "" {
		params["fields"] = tasks.Fields
	}
	if tasks.ExcludedFields != "" {
		params["excludedField"] = tasks.ExcludedFields
	}
	if tasks.OutdatedTasks != "" {
		params["outdataedTasks"] = tasks.OutdatedTasks
	}
	if tasks.ProjectId != "" {
		params["projectId"] = tasks.ProjectId
	}
	if tasks.View != "" {
		params["view"] = tasks.View
	}

	headers := make(map[string]string)

	token := helper.GetToken()
	headers["token"] = *token

	contents = helper.GetJsonContents()
	//unmarshal the json to struct object
	err := json.Unmarshal([]byte(contents), &c)
	if err != nil {
		return nil, errors.New("new request is fail ")
	}
	//construct login Restful API path
	apiPath := c.WSTasks.Tasks

	if apiPath == "" {
		return nil, errors.New("new request is fail ")
	}

	data, err := httpx.Get(apiPath, params, headers)

	return data, err

}
