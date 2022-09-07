package gotenableutils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// StructToJSON marshals struct data to JSON
func StructToJSON(structData interface{}) *bytes.Buffer {
	body, err := json.Marshal(structData)
	ErrFatal(err)
	return bytes.NewBuffer(body)
}

// JSONToStruct unmarshals JSON to struct
func JSONToStruct(jsonData []byte, model interface{}) {
	ErrFatal(json.Unmarshal(jsonData, model))
}

// JSONToIface unmarshals JSON data to an interface
func JSONToIface(jsonData []byte) interface{} {
	var ifaceData interface{}
	ErrFatal(json.Unmarshal(jsonData, &ifaceData))
	return ifaceData
}

// PPrintJSON pretty prints JSON data
func PPrintJSON(jsonData []byte) {
	ifaceData := JSONToIface(jsonData)
	prettyData, err := json.MarshalIndent(ifaceData, ``, `  `)
	ErrFatal(err)
	fmt.Println(string(prettyData))
}
