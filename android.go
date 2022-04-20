package umeng

type AndroidBody struct {
	Title       string      `json:"title,omitempty"`
	Text        string      `json:"text,omitempty"`
	Icon        string      `json:"icon,omitempty"`
	LargeIcon   string      `json:"largeIcon,omitempty"`
	Img         string      `json:"img,omitempty"`
	Sound       string      `json:"sound,omitempty"`
	BuilderId   string      `json:"builder_id,omitempty"`
	PlayVibrate bool        `json:"play_vibrate,omitempty"`
	PlayLights  bool        `json:"play_lights,omitempty"`
	PlaySound   bool        `json:"play_sound,omitempty"`
	AfterOpen   string      `json:"after_open,omitempty"`
	Url         string      `json:"url,omitempty"`
	Activity    string      `json:"activity,omitempty"`
	Custom      interface{} `json:"custom,omitempty"`
}

type AndroidPayload struct {
	DisplayType string                 `json:"display_type,omitempty"`
	Body        *AndroidBody           `json:"body,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
}
