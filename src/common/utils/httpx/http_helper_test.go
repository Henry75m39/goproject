package httpx

import (
	"GRM/src/common/utils/httpx"
	"testing"
)

//test httpx.DefaultGet method
func TestXXX(t *testing.T) {

	url := ""
	params := []map[string]string{
		{},
	}
	for _, param := range params {
		resp, err := httpx.DefaultGet(url, param)
		httpx.PrintResult(resp, err, t)
	}

}
