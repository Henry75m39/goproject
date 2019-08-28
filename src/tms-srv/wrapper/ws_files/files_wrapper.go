package ws_files

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/helper"
	"GRM/src/common/utils/log"
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func Upload(file *multipart.File, handler *multipart.FileHeader) (data []byte, err error) {
	var c configs.WSConfig
	var contents []byte

	logger := log.Instance()
	token := helper.GetToken()

	contents = helper.GetJsonContents()
	//unmarshal the json to struct object
	_ = json.Unmarshal([]byte(contents), &c)
	//construct login Restful API path
	apiPath := c.WSFiles.Files

	if apiPath == "" {
		logger.Error("Error", zap.Any("Upload File", "REST API cannot read from JSON config file."))
		return
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(handler.Filename))
	if err != nil {
		logger.Error("Error", zap.Any("Upload File", "Create form file occurs error."))

	}
	_, err = io.Copy(part, *file)
	_ = writer.Close()

	req, err := http.NewRequest("POST", apiPath, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("token", *token)

	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//Reads the body of response and closes the body reader when done reading it.
	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	dynamicStrutForJson := make(map[string]interface{})
	err = json.Unmarshal(data, &dynamicStrutForJson)
	if err != nil {
		logger.Error("Error", zap.Any("FileUpload:", "REST API cannot unmarshal correctly."))

	}
	//make json looks beautiful.
	beautifulJson, _ := json.MarshalIndent(dynamicStrutForJson, "", "\t")

	return beautifulJson, err

	//status = helper.ParseCommonJSONResult(data)
	//return status, err
	// Check the response
	//if res.StatusCode != http.StatusOK {
	//fmt.Println(http.StatusOK, res.Status)
	//	err = fmt.Errorf("bad status: %s", res.Status)
	//}

}
