package i18n

type Language string

const (
	LanguageEn Language = "en"
	LanguageJa Language = "ja"
)

func GetCurrentLanguage(lang string) Language {
	if lang == string(LanguageJa) {
		return LanguageJa
	} else {
		return LanguageEn
	}
}
