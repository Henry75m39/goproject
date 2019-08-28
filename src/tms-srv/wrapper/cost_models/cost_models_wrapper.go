package cost_models

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/helper"
	"GRM/src/common/utils/log"
	"GRM/src/tms-srv/entity"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func CostModels(models entity.CostModels) ([]byte, error) {
	var c configs.WSConfig
	var contents []byte

	params := make(map[string]string)

	if models.SortBy != "" {
		params["sortBy"] = models.SortBy
	}

	if models.SortDirection != "" {
		params["sortDirection"] = models.SortDirection
	}

	if models.ExcludedFields != "" {
		params["excludedFiedls"] = models.ExcludedFields
	}

	if models.Fields != "" {
		params["fields"] = models.Fields
	}

	if models.LocaleId != "" {
		params["localeId"] = models.LocaleId
	}

	if models.ScopingId != "" {
		params["scopingId"] = models.ScopingId
	}

	logger := log.Instance()
	token := helper.GetToken()

	contents = helper.GetJsonContents()
	//unmarshal the json to struct object
	err := json.Unmarshal([]byte(contents), &c)
	if err != nil {
		logger.Error("Error", zap.Any("CostModels", "REST API cannot unmarshal from JSON file. "))
	}
	//construct login Restful API path
	apiPath := c.WSCostModels.CostModels

	if apiPath == "" {
		logger.Error("Error", zap.Any("CostModels", "REST API cannot read from JSON config file. "))
	}

	req, _ := http.NewRequest("GET", apiPath, nil)

	req.Header.Add("token", *token)
	//Query parses RawQuery and returns the corresponding values.
	q := req.URL.Query()

	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	rsp, err := client.Do(req)
	data, _ := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()

	dynamicStrutForJson := make(map[string]interface{})
	err = json.Unmarshal(data, &dynamicStrutForJson)
	if err != nil {
		logger.Error("Error", zap.Any("CostModels", "REST API cannot unmarshal correctly."))

	}
	//make json looks beautiful.
	beautifulJson, _ := json.MarshalIndent(dynamicStrutForJson, "", "\t")

	//testing the output result into a json file.
	//err = ioutil.WriteFile("post.json", []byte(beautifulJson), 0644)
	return beautifulJson, err

	//status = helper.ParseCommonJSONResult(data)
	//return data, err
}
