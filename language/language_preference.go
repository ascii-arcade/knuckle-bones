package language

type LanguagePreference struct {
	Lang *Language
}

func (lp *LanguagePreference) SetLanguage(abbr string) {
	lang, exists := Languages[abbr]
	if !exists {
		return
	}
	lp.Lang = lang
}
