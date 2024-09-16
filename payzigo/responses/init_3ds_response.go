package responses

import "encoding/base64"

type Init3dsResponse struct {
	Status             string `json:"status"`
	Locale             string `json:"locale"`
	SystemTime         int64  `json:"systemTime"`
	ConversationID     string `json:"conversationId"`
	ThreeDSHTMLContent string `json:"threeDSHtmlContent"`
}

func (e *Init3dsResponse) GetHtmlContent() string {
	threedsDecoded, err := base64.StdEncoding.DecodeString(e.ThreeDSHTMLContent)
	if err != nil {
		return err.Error()
	}

	return string(threedsDecoded)
}
