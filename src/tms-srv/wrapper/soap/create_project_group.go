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
	wsdl := "http://wsvr-stg-app-1.vmware.com:8080/ws/services/WorkflowWSWorkflowManager?wsdl"

	soap, err := gosoap.SoapClient(wsdl)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	/*
		Creates a project group containing projects associated with the provided locales for the provided files.
		This will result in creation of one project per locale, each having a set of tasks (one per each file).
		The files are copied into AIS location under the Client's folder's (see WSClient.getAisLocation()) project directory.

		The workflow, TM, TD, Workgroup, Cost Models, and other settings are inferred from the project type and client objects.

		Parameters:
			name - The name of the projet group
			description - The detaild descriptoin of the project group. This will be used as the descriptoin for each of the created projects.
			locales - The target locales for which to create the projects. Linkages will be set up as part of the project creation
			files - The files that have to be translated. These files are copied into the Client's folder's project directory prior to project creation.
			root - The root location of the files. When the above files are copied into the Client's folder's project diretory, the folder structure relative to this root is preserved. If this parameter is null, then the file structure under the Client's folder's project directory is flat. This parameter has to represent a parent directory of every file in the above paramter.
			client - The client for who this project is created. The client contains client-specific settings for TMs, TDs, and other parameters that will be used in the project creation.
			projectType - The type of a project that is created. Project type defines Workflow, Source Locale and other project specific parameters for project creation.
			customAisProperties - custom AIS properties to be applied during project group creation.
		Returns:
			The created project group
		Since:
			9.0.0

		WSProjectGroup createProjectGroup(String name,
									  String description,
									  WSLocale[] locales,
									  File[] files,
									  File root,
									  WSClient client,
									  WSProjectType projectType,
									  Map customAisProperties)
	*/

	//userWSLocalArray := [...] int{1,2,3}
	params := gosoap.Params{
		"token":       "1451573115",
		"name":        "henry_soap_test1",
		"description": "henry is testing",
		"workgroup":   "Test Group",
		"locales":     "",
		"assets":      "",
		"workflow":    "",
		"assetsMode":  "0",
		"costModel":   "",

		/*  "<>[UserWSLocaleArray]</locales>"     +
		"<>[AisWSNodeArray]</assets>"      +
		"<>[WorkflowWSWorkflow]</workflow> " +
		"<>0</assetsMode>"   +
		"<></costModel>"
		*/
	}

	_ = soap.Call("createProjectGroup", params)

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
