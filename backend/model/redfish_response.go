package model

// 红帽响应
type RedfishResponse struct {
	Error struct {
		MessageExtendedInfo []struct {
			Message                     string   `json:"Message"`
			MessageArgs                 []string `json:"MessageArgs"`
			MessageArgsOdataCount       int      `json:"MessageArgs@odata.count"`
			MessageID                   string   `json:"MessageId"`
			RelatedProperties           []string `json:"RelatedProperties"`
			RelatedPropertiesOdataCount int      `json:"RelatedProperties@odata.count"`
			Resolution                  string   `json:"Resolution"`
			Severity                    string   `json:"Severity"`
		} `json:"@Message.ExtendedInfo"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
