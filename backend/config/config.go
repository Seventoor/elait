package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// API Keys
	DeepgramAPIKey string
	OpenAIAPIKey   string

	SourceLanguage string // Sprache, die gesprochen wird -> ISO 639-1

	SpeechRate  string
	SpeechPitch string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Keine .env Datei gefunden, verwende System-Umgebungsvariablen")
	}

	return Config{
		DeepgramAPIKey: getEnv("DEEPGRAM_API_KEY", ""),
		OpenAIAPIKey:   getEnv("OPENAI_API_KEY", ""),

		SourceLanguage: "de",

		SpeechRate:  "1",
		SpeechPitch: "0",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
