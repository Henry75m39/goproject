package main

import (
	"encoding/xml"
	"fmt"
	"github.com/tiaguinho/gosoap"
	"log"
)

// GetIPLocationResponse will hold the Soap response
type GetCreateProjectGroupResponse struct {
	GetIPLocationResult string `xml:"GetIpLocationResult"`
}

// GetIPLocationResult will
type GetIPLocationResult struct {
	XMLName xml.Name `xml:"GeoIP"`
	Country string   `xml:"Country"`
	State   string   `xml:"State"`
}

var (
	r GetCreateProjectGroupResponse
)

func main() {
	//wsdl := "http://wsvr-stg-app-1.vmware.com:8080/ws/services/WorkflowWSWorkflowManager?wsdl"

	wsdl := "http://wsvr-stg-app-1.vmware.com:8080/ws/services/WorkflowWSProjectGroup?wsdl"

	soap, err := gosoap.SoapClient(wsdl)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	/*
		   createProjectGroup(
			   java.lang.String,
			   java.lang.String,
			   java.lang.String,
			   com.idiominc.webservices.data.UserWSWorkgroup,
			   com.idiominc.webservices.data.UserWSLocale[],
			   com.idiominc.webservices.data.AisWSNode[],
			   com.idiominc.webservices.data.WorkflowWSWorkflow,
			   int,
			   com.idiominc.webservices.data.CostmodelWSCostModel
		  )
	*/

	params := gosoap.Params{
		"sIp": "8.8.8.8",
	}

	_ = soap.Call("createProject", params)

	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

	// GetIpLocationResult will be a string. We need to parse it to XML

	err = soap.Unmarshal(&r)
	result := GetIPLocationResult{}
	err = xml.Unmarshal([]byte(r.GetIPLocationResult), &result)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	fmt.Println(result.Country)
	if result.Country != "US" {
		log.Fatalf("error: %+v", r)
	}

	log.Println("Country: ", result.Country)
	log.Println("State: ", result.State)
}
