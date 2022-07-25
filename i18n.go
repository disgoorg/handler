package handler

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/log"
	"golang.org/x/text/message/catalog"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func newI18n(logger log.Logger, fallbackLanguage language.Tag) *I18n {
	return &I18n{
		Logger:  logger,
		Catalog: catalog.NewBuilder(catalog.Fallback(fallbackLanguage)),
	}
}

type I18n struct {
	Logger  log.Logger
	Catalog *catalog.Builder
}

func (i *I18n) LoadFromEmbedFS(fs embed.FS, dir string) error {
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		tag := strings.TrimSuffix(entry.Name(), ".json")
		lang, err := language.Parse(tag)
		if err != nil {
			i.Logger.Errorf("Failed to parse tag %s: %s", tag, err)
			continue
		}

		rawData, err := fs.ReadFile(fmt.Sprintf("%s/%s", dir, entry.Name()))
		if err != nil {
			i.Logger.Error("Failed to read language file: ", err)
			continue
		}

		var data map[string]interface{}
		if err = json.Unmarshal(rawData, &data); err != nil {
			i.Logger.Error("Failed to parse language file: ", err)
			continue
		}
		i.AddLanguage(lang, "", data)
	}
	return nil
}

func (i *I18n) AddLanguage(lang language.Tag, basePath string, data map[string]interface{}) {
	for key, value := range data {
		if value == nil {
			continue
		}
		newBasePath := strings.TrimPrefix(basePath+"."+key, ".")
		switch v := value.(type) {
		case string:
			if err := i.Catalog.SetString(lang, newBasePath, v); err != nil {
				i.Logger.Errorf("Failed to set string with path %s and value %s: %s", newBasePath, v, err)
			}
		case map[string]interface{}:
			i.AddLanguage(lang, newBasePath, v)
		}
	}
}

func (i *I18n) NewPrinter(locale discord.Locale) *message.Printer {
	lang, err := language.Parse(locale.Code())
	if err != nil {
		i.Logger.Errorf("error while parsing locale %s: %s", locale.Code(), err)
		lang = language.Und
	}
	return message.NewPrinter(lang, message.Catalog(i.Catalog))
}
