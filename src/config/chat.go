package config

// Chat is a Minecraft formatted Chat object, see: https://wiki.vg/Chat.
type Chat struct {
	Text          string      `yaml:"text" json:"text"`
	Color         string      `yaml:"color" json:"color,omitempty"`
	Bold          bool        `yaml:"bold" json:"bold,omitempty"`
	Italic        bool        `yaml:"italic" json:"italic,omitempty"`
	Underlined    bool        `yaml:"underlined" json:"underlined,omitempty"`
	Strikethrough bool        `yaml:"strikethrough" json:"strikethrough,omitempty"`
	Obfuscated    bool        `yaml:"obfuscated" json:"obfuscated,omitempty"`
	Font          string      `yaml:"font" json:"font,omitempty"`
	ClickEvent    *ClickEvent `yaml:"click_event" json:"clickEvent,omitempty"`
	HoverEvent    *HoverEvent `yaml:"hover_event" json:"hoverEvent,omitempty"`
	Extra         []Chat      `yaml:"extra" json:"extra,omitempty"`
}

// ClickEvent is a Minecraft `clickEvent` serialized in the Chat object.
type ClickEvent struct {
	Action string `yaml:"action" json:"action"`
	Value  string `yaml:"value" json:"value"`
}

// HoverEvent is a Minecraft `hoverEvent` serialized in the Chat object.
type HoverEvent struct {
	Action   string      `yaml:"action" json:"action"`
	Contents interface{} `yaml:"contents" json:"contents"`
}
