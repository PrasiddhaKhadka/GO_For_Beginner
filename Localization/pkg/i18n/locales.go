package i18n

import (
	"embed"
	"encoding/json"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type TranslationKey string

//go:embed en.json ja.json
var localeFs embed.FS

var localizers = make(map[Language]*goi18n.Localizer)

func init() {
	bundle := goi18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
}
