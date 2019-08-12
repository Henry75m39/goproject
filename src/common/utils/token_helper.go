package utils

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/db"
	"GRM/src/common/utils/helper"
	"GRM/src/common/utils/log"
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type UserInfo struct {
	SessionID      *string `json:"sessionId"`
	InstanceID     int     `json:"instanceId"`
	ExpirationTime string  `json:"expirationTime"`
	UserDetails    struct {
		UserType               string `json:"userType"`
		Username               string `json:"username"`
		FullName               string `json:"fullName"`
		Fingerprint            string `json:"fingerprint"`
		RegionalSettingsLocale string `json:"regionalSettingsLocale"`
		Language               struct {
			ID           int    `json:"id"`
			LanguageCode string `json:"languageCode"`
			CountryCode  string `json:"countryCode"`
			IsoCode      string `json:"isoCode"`
			Locale       string `json:"locale"`
		} `json:"language"`
	} `json:"userDetails"`
	LastUpdateTime  int64  `json:"lastUpdateTime"`
	DaysToPwdExpire int    `json:"daysToPwdExpire"`
	LoginOutcome    string `json:"loginOutcome"`
}

func GetToken() (token *string) {
	var c configs.WSConfig
	var contents []byte
	var err error
	var expired bool
	var timeInSeconds float64

	timeInSeconds = 1800 * 1000

	logger := log.Instance()

	isKeyExist := db.ValidateKeyInDb("sessionId")

	if isKeyExist {
		timeStamp, _ := db.Instance.Get([]byte("timeStamp"), nil)
		//calculate time differ, and save new timestamp.
		t, _ := time.Parse("2006-01-02 15:04:05", string(timeStamp))
		now := time.Now()
		subSecond := now.Sub(t)

		if subSecond.Seconds() <= timeInSeconds {
			//not expired
			cachedToken, err := db.Instance.Get([]byte("sessionId"), nil)
			strToken := string(cachedToken)
			if err == nil {
				logger.Info("Info", zap.Any("Key:", "sessionId"))
				logger.Info("Info", zap.Any("Value:", string(cachedToken)))
			}
			//db.Instance.Close()
			expired = false
			return &strToken
		} else {
			//session expired need to get a new session ID from WorldServer.
			expired = true
		}
	}

	if expired || !isKeyExist {
		contents = helper.GetJsonContents()
		//unmarshal the json to struct object
		err = json.Unmarshal([]byte(contents), &c)
		if err != nil {
			logger.Info("ERROR:", zap.Any("Parsing JSON file:", err))
			return
		}
		//construct login Restful API path
		apiPath := c.WSLoginAPI
		userName := c.WSInfo.Username
		wsPwd := c.WSInfo.Password

		jsonData := map[string]string{"username": userName, "password": wsPwd}
		jsonValue, _ := json.Marshal(jsonData)

		request, _ := http.NewRequest("POST", apiPath, bytes.NewBuffer(jsonValue))
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		rsp, err := client.Do(request)

		data, err := ioutil.ReadAll(rsp.Body)
		token = getSessionID(data)

		//store a timestamp
		t := time.Now().Format("2006-01-02 15:04:05")
		err = db.Instance.Put([]byte("timeStamp"), []byte(t), nil)

		err = db.Instance.Put([]byte("sessionId"), []byte(*token), nil)

		if err != nil {
			logger.Error("ERROR", zap.Any("Put new sessionId failed caused by", err))
		}
	}

	err = db.Instance.Close()
	if err != nil {
		logger.Error("ERROR", zap.Any("Related to LevelDB", err))

	}
	return token
}

func getSessionID(jsonResult []byte) (sessionID *string) {
	var r UserInfo
	_ = json.Unmarshal(jsonResult, &r)
	sessionID = r.SessionID
	//for _, wsItem := range r.SessionID {

	//for _, cwItem := range wsItem. {
	//	recPlainResult = recPlainResult + cwItem.W
	//}
	//}
	return sessionID
}
