package helper

import (
	"GRM/src/common/configs"
	"GRM/src/common/utils/db"
	"GRM/src/common/utils/log"
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type UserInfo struct {
	SessionID       *string     `json:"sessionId"`
	ExpirationTime  string      `json:"expirationTime"`
	UserDetails     UserDetails `json:"userDetails"`
	LoginOutcome    string      `json:"loginOutcome"`
	LastUpdateTime  int64       `json:"lastUpdateTime"`
	DaysToPwdExpire int         `json:"daysToPwdExpire"`
}
type Language struct {
	ID           int    `json:"id"`
	LanguageCode string `json:"languageCode"`
	CountryCode  string `json:"countryCode"`
	Locale       string `json:"locale"`
}
type UserDetails struct {
	Username               string   `json:"username"`
	FullName               string   `json:"fullName"`
	Fingerprint            string   `json:"fingerprint"`
	Language               Language `json:"language"`
	RegionalSettingsLocale string   `json:"regionalSettingsLocale"`
}

func GetToken() (token *string) {
	var c configs.WSConfig
	var contents []byte
	var err error
	var expired bool
	var sessionTimeOutDuration int64

	//initial setting for the WorldServer session timeout in 20 minutes.
	//calculation in seconds.
	sessionTimeOutDuration = 20 * 60

	logger := log.Instance()

	isKeyExist := db.ValidateKeyInDb("sessionId")

	if isKeyExist {
		timeStamp, _ := db.Instance.Get([]byte("timeStamp"), nil)
		//calculate time differ, and save new timestamp.
		nowTimeInStr := time.Now().Format("2006-01-02 15:04:05")

		differSecond := getHourDiffer(string(timeStamp), nowTimeInStr)

		if differSecond <= sessionTimeOutDuration {
			//not expired
			cachedToken, err := db.Instance.Get([]byte("sessionId"), nil)
			//update cached token timestamp
			err = db.Instance.Put([]byte("timeStamp"), []byte(nowTimeInStr), nil)
			strToken := string(cachedToken)
			if err != nil {
				logger.Error("ERROR:", zap.Any("LevelDB:", "DB occurs error when Get"))
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
		contents = GetJsonContents()
		//unmarshal the json to struct object
		err = json.Unmarshal([]byte(contents), &c)
		if err != nil {
			logger.Error("ERROR:", zap.Any("Parsing JSON file:", err))
			panic("Error")
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
		if err != nil {
			logger.Error("Error:", zap.Any("Request handling error occurred:", err))
		}

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

	//err = db.Instance.Close()
	//if err != nil {
	//	logger.Error("ERROR", zap.Any("Related to LevelDB", err))

	//}
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

func getHourDiffer(startTime, endTime string) int64 {
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		//internal testing for time calculation
		hour := float64(diff) / 3600.00
		fmt.Println(hour)
		second := hour * 3600.00
		fmt.Println(second)
		return diff
	} else {
		return 0
	}
}
