package ws_files

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/helper"
	"GRM/src/common/utils/log"
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func Upload(file *multipart.FileHeader) (status []string, err error) {

	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	_ = w.Close()

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
	apiPath := c.WSFiles.Files

	if apiPath == "" {
		logger.Error("Error", zap.Any("createProjectForTask", "REST API cannot read from JSON config file "))
		return
	}

	// Now that you have a form, you can submit it to your handler.

	req, err := http.NewRequest("POST", apiPath, &b)

	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("token", *token)
	if err != nil {
		panic(err)
	}
	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//Reads the body of response and closes the body reader when done reading it.
	data, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	status = helper.ParseCommonJSONResult(data)
	return status, err
	// Check the response
	//if res.StatusCode != http.StatusOK {
	//fmt.Println(http.StatusOK, res.Status)
	//	err = fmt.Errorf("bad status: %s", res.Status)
	//}

}
