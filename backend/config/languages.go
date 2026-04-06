package config

type Language struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

var Languages = []Language{
	{Name: "Afrikaans (South Africa)", Code: "af-ZA"},
	{Name: "Arabic", Code: "ar-XA"},
	{Name: "Basque (Spain)", Code: "eu-ES"},
	{Name: "Bengali (India)", Code: "bn-IN"},
	{Name: "Bulgarian (Bulgaria)", Code: "bg-BG"},
	{Name: "Catalan (Spain)", Code: "ca-ES"},
	{Name: "Chinese (Hong Kong)", Code: "yue-HK"},
	{Name: "Croatian (Croatia)", Code: "hr-HR"},
	{Name: "Czech (Czech Republic)", Code: "cs-CZ"},
	{Name: "Danish (Denmark)", Code: "da-DK"},
	{Name: "Dutch (Belgium)", Code: "nl-BE"},
	{Name: "Dutch (Netherlands)", Code: "nl-NL"},
	{Name: "English (Australia)", Code: "en-AU"},
	{Name: "English (India)", Code: "en-IN"},
	{Name: "English (United Kingdom)", Code: "en-GB"},
	{Name: "English (USA)", Code: "en-US"},
	{Name: "Estonian (Estonia)", Code: "et-EE"},
	{Name: "Filipino (Philippines)", Code: "fil-PH"},
	{Name: "Finnish (Finland)", Code: "fi-FI"},
	{Name: "French (Canada)", Code: "fr-CA"},
	{Name: "French (France)", Code: "fr-FR"},
	{Name: "Galician (Spain)", Code: "gl-ES"},
	{Name: "German (Germany)", Code: "de-DE"},
	{Name: "Greek (Greece)", Code: "el-GR"},
	{Name: "Gujarati (India)", Code: "gu-IN"},
	{Name: "Hebrew (Israel)", Code: "he-IL"},
	{Name: "Hindi (India)", Code: "hi-IN"},
	{Name: "Hungarian (Hungary)", Code: "hu-HU"},
	{Name: "Icelandic (Iceland)", Code: "is-IS"},
	{Name: "Indonesian (Indonesia)", Code: "id-ID"},
	{Name: "Italian (Italy)", Code: "it-IT"},
	{Name: "Japanese (Japan)", Code: "ja-JP"},
	{Name: "Kannada (India)", Code: "kn-IN"},
	{Name: "Korean (South Korea)", Code: "ko-KR"},
	{Name: "Latvian (Latvia)", Code: "lv-LV"},
	{Name: "Lithuanian (Lithuania)", Code: "lt-LT"},
	{Name: "Malay (Malaysia)", Code: "ms-MY"},
	{Name: "Malayalam (India)", Code: "ml-IN"},
	{Name: "Mandarin (China)", Code: "cmn-CN"},
	{Name: "Mandarin (Taiwan)", Code: "cmn-TW"},
	{Name: "Marathi (India)", Code: "mr-IN"},
	{Name: "Norwegian (Norway)", Code: "nb-NO"},
	{Name: "Polish (Poland)", Code: "pl-PL"},
	{Name: "Portuguese (Brazil)", Code: "pt-BR"},
	{Name: "Portuguese (Portugal)", Code: "pt-PT"},
	{Name: "Punjabi (India)", Code: "pa-IN"},
	{Name: "Romanian (Romania)", Code: "ro-RO"},
	{Name: "Russian (Russia)", Code: "ru-RU"},
	{Name: "Serbian (Cyrillic)", Code: "sr-RS"},
	{Name: "Slovak (Slovakia)", Code: "sk-SK"},
	{Name: "Slovenian (Slovenia)", Code: "sl-SI"},
	{Name: "Spanish (Spain)", Code: "es-ES"},
	{Name: "Spanish (USA)", Code: "es-US"},
	{Name: "Swedish (Sweden)", Code: "sv-SE"},
	{Name: "Tamil (India)", Code: "ta-IN"},
	{Name: "Telugu (India)", Code: "te-IN"},
	{Name: "Thai (Thailand)", Code: "th-TH"},
	{Name: "Turkish (Turkey)", Code: "tr-TR"},
	{Name: "Ukrainian (Ukraine)", Code: "uk-UA"},
	{Name: "Urdu (India)", Code: "ur-IN"},
	{Name: "Vietnamese (Vietnam)", Code: "vi-VN"},
}

func IsValidLangCode(code string) bool {
	for _, lang := range Languages {
		if lang.Code == code {
			return true
		}
	}
	return false
}
