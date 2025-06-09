package config

import (
	"cmp"
	"os"

	"github.com/ascii-arcade/game-template/language"
	"github.com/charmbracelet/log"
)

const (
	MinimumHeight = 40
	MinimumWidth  = 120
)

var (
	Language *language.Language = setDefaultLanguage()

	Version  string = "dev"
	Debug    bool   = cmp.Or(os.Getenv("ASCII_ARCADE_DEBUG"), "false") == "true"
	Host     string = cmp.Or(os.Getenv("ASCII_ARCADE_HOST"), "localhost")
	SSHPort  string = cmp.Or(os.Getenv("ASCII_ARCADE_SSH_PORT"), "23234")
	HTTPPort string = cmp.Or(os.Getenv("ASCII_ARCADE_HTTP_PORT"), "8080")
)

func setDefaultLanguage() *language.Language {
	langCode := cmp.Or(os.Getenv("ASCII_ARCADE_LANG"), "EN")
	lang, exists := language.Languages[langCode]
	if !exists {
		log.Warn("Unknown language code %s, defaulting to English", langCode)
		return language.Languages["EN"]
	}
	return lang
}
