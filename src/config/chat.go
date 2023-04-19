package config

import "strings"

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

// FixControlCharacters returns a copy of the Chat with new-line control characters corrected in the Text property.
func (c Chat) FixControlCharacters() Chat {
	c.Text = strings.ReplaceAll(c.Text, "\\n", "\n")

	extra := make([]Chat, 0)

	for _, e := range c.Extra {
		extra = append(extra, e.FixControlCharacters())
	}

	c.Extra = extra

	return c
}

// TODO rewrite this method to correctly convert Chat into the legacy system, allowing inheritance
func (c Chat) String() string {
	value := c.Text

	if c.Extra != nil {
		for _, e := range c.Extra {
			value += e.String()
		}
	}

	return value
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
