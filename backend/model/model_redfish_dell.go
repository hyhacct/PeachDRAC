package model

// 获取基础地址
type ModelRedfishDellBaseUrl struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Description  string `json:"Description"`
	Members      []struct {
		OdataID string `json:"@odata.id"`
	} `json:"Members"`
	Name string `json:"Name"`
}
