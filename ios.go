package umeng

type IOSAps struct {
	Alert            string `json:"alert,omitempty"`             // 必填
	Badge            string `json:"badge,omitempty"`             // 选填
	Sound            string `json:"sound,omitempty"`             // 选填
	ContentAvailable string `json:"content-available,omitempty"` // 选填
	Category         string `json:"category,omitempty"`          // 选填 ios8才支持该字段
}

func IOSPayload(aps *IOSAps, extras ...map[string]interface{}) *map[string]interface{} {
	payload := make(map[string]interface{}, 0)
	payload["aps"] = aps
	if len(extras) > 0 {
		for _, config := range extras {
			for key, val := range config {
				payload[key] = val
			}
		}
	}
	return &payload
}
