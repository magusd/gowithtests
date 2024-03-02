package hello

import "fmt"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	template := GetTemplate(lang)
	return fmt.Sprintf(template, name)
}

const (
	english = "en"
	spanish = "es"
	french  = "fr"

	EnglishTemplate = "Hello, %s!"
	SpanishTemplate = "Hola, %s!"
	FrenchTemplate  = "Bonjour, %s!"
)

func GetTemplate(lang string) string {

	template := EnglishTemplate

	switch lang {
	case english:
		template = EnglishTemplate
	case spanish:
		template = SpanishTemplate
	case french:
		template = FrenchTemplate
	}
	return template
}
