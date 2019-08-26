package helper

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

//Not available to use still need to enhance and modify in future
func UploadSingleFile(token *string, filename string, targetUrl string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")

	}

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")

	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)

	//
	//req, _ := http.NewRequest("POST", targetUrl, bodyBuf )

	//req.Header.Set("Content-Type", contentType)
	//req.Header.Add("token", *token)

	//client := &http.Client{}
	//resp, err := client.Do(req)
	//

	return resp, err

	/*
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(resp.Status)
		fmt.Println(string(respBody))
		return nil

	*/
}

//Almost workable to use still need to enhance and modify in future
// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

/*
func main() {
	path, _ := os.Getwd()
	path += "/test.pdf"
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	request, err := newfileUploadRequest(apiPath, extraParams, "file", "/tmp/doc.pdf")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)

		fmt.Println(body)
	}
}
*/

// sample usage
//func main() {
//target_url := "http://wsvr-stg-app-1.vmware.com:8080/ws-api/v1/files"
//filename := "/Volumes/Henry/WSLibProject/src/tms-srv/wrapper/upload.txt"
//_, _ = postFile(filename, target_url)
//}
