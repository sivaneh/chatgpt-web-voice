package voice

import "strings"

const (
	defaultVoice        = "cove"
	defaultVoiceMode    = "wingman"
	defaultLanguageCode = "auto"
)

// Choice is one public identifier and display label.
type Choice struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// VoiceChoice describes one upstream realtime voice.
type VoiceChoice struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// LanguageChoice provides stable labels without tying downstream clients to
// the built-in page's current locale.
type LanguageChoice struct {
	ID     string `json:"id"`
	NameZH string `json:"name_zh"`
	NameEN string `json:"name_en"`
}

// SessionDefaults are the values applied when a downstream field is omitted.
type SessionDefaults struct {
	Voice        string `json:"voice"`
	VoiceMode    string `json:"voice_mode"`
	LanguageCode string `json:"language_code"`
}

// DataChannelConfig is the negotiated channel required by ChatGPT voice.
type DataChannelConfig struct {
	Label      string `json:"label"`
	Negotiated bool   `json:"negotiated"`
	ID         int    `json:"id"`
}

// ICEServerConfig is compatible with the browser RTCIceServer shape.
type ICEServerConfig struct {
	URLs []string `json:"urls"`
}

// WebRTCConfig contains connection bootstrap values safe for downstream use.
type WebRTCConfig struct {
	DataChannel  DataChannelConfig `json:"data_channel"`
	ICEServers   []ICEServerConfig `json:"ice_servers"`
	ReceiveAudio bool              `json:"receive_audio"`
	ReceiveVideo bool              `json:"receive_video"`
}

// PublicConfig is the complete non-secret downstream capability document.
type PublicConfig struct {
	Version    string           `json:"version"`
	Defaults   SessionDefaults  `json:"defaults"`
	Voices     []VoiceChoice    `json:"voices"`
	VoiceModes []Choice         `json:"voice_modes"`
	Languages  []LanguageChoice `json:"languages"`
	WebRTC     WebRTCConfig     `json:"webrtc"`
}

var voiceChoices = []VoiceChoice{
	{ID: "breeze", Name: "Breeze", Description: "Animated"},
	{ID: "cove", Name: "Cove", Description: "Composed"},
	{ID: "ember", Name: "Ember", Description: "Confident"},
	{ID: "fathom", Name: "Arbor", Description: "Easygoing"},
	{ID: "glimmer", Name: "Sol", Description: "Savvy"},
	{ID: "juniper", Name: "Juniper", Description: "Open"},
	{ID: "maple", Name: "Maple", Description: "Cheerful"},
	{ID: "orbit", Name: "Spruce", Description: "Calm"},
	{ID: "vale", Name: "Vale", Description: "Bright"},
}

var voiceModeChoices = []Choice{
	{ID: defaultVoiceMode, Name: "Standard"},
}

var languageChoices = []LanguageChoice{
	{ID: "auto", NameZH: "Auto-detect", NameEN: "Auto-detect"},
	{ID: "en", NameZH: "English", NameEN: "English"},
	{ID: "es", NameZH: "Spanish", NameEN: "Spanish"},
	{ID: "pt", NameZH: "Portuguese", NameEN: "Portuguese"},
	{ID: "fr", NameZH: "French", NameEN: "French"},
	{ID: "de", NameZH: "German", NameEN: "German"},
	{ID: "ja", NameZH: "Japanese", NameEN: "Japanese"},
	{ID: "id", NameZH: "Indonesian", NameEN: "Indonesian"},
	{ID: "ru", NameZH: "Russian", NameEN: "Russian"},
	{ID: "it", NameZH: "Italian", NameEN: "Italian"},
	{ID: "tr", NameZH: "Turkish", NameEN: "Turkish"},
	{ID: "ar", NameZH: "Arabic", NameEN: "Arabic"},
	{ID: "hi", NameZH: "Hindi", NameEN: "Hindi"},
	{ID: "ko", NameZH: "Korean", NameEN: "Korean"},
	{ID: "nl", NameZH: "Dutch", NameEN: "Dutch"},
	{ID: "pl", NameZH: "Polish", NameEN: "Polish"},
	{ID: "vi", NameZH: "Vietnamese", NameEN: "Vietnamese"},
	{ID: "uk", NameZH: "Ukrainian", NameEN: "Ukrainian"},
	{ID: "sv", NameZH: "Swedish", NameEN: "Swedish"},
	{ID: "da", NameZH: "Danish", NameEN: "Danish"},
	{ID: "nb", NameZH: "Norwegian Bokmål", NameEN: "Norwegian Bokmål"},
	{ID: "no", NameZH: "Norwegian", NameEN: "Norwegian"},
	{ID: "th", NameZH: "Thai", NameEN: "Thai"},
	{ID: "ro", NameZH: "Romanian", NameEN: "Romanian"},
	{ID: "ms", NameZH: "Malay", NameEN: "Malay"},
	{ID: "bn", NameZH: "Bangla", NameEN: "Bangla"},
	{ID: "mr", NameZH: "Marathi", NameEN: "Marathi"},
	{ID: "ta", NameZH: "Tamil", NameEN: "Tamil"},
	{ID: "te", NameZH: "Telugu", NameEN: "Telugu"},
	{ID: "gu", NameZH: "Gujarati", NameEN: "Gujarati"},
	{ID: "ur", NameZH: "Urdu", NameEN: "Urdu"},
	{ID: "ml", NameZH: "Malayalam", NameEN: "Malayalam"},
	{ID: "kn", NameZH: "Kannada", NameEN: "Kannada"},
	{ID: "sw", NameZH: "Swahili", NameEN: "Swahili"},
	{ID: "zh", NameZH: "Mandarin Chinese", NameEN: "Mandarin Chinese"},
	{ID: "af", NameZH: "Afrikaans", NameEN: "Afrikaans"},
	{ID: "hy", NameZH: "Armenian", NameEN: "Armenian"},
	{ID: "az", NameZH: "Azerbaijani", NameEN: "Azerbaijani"},
	{ID: "be", NameZH: "Belarusian", NameEN: "Belarusian"},
	{ID: "bs", NameZH: "Bosnian", NameEN: "Bosnian"},
	{ID: "bg", NameZH: "Bulgarian", NameEN: "Bulgarian"},
	{ID: "ca", NameZH: "Catalan", NameEN: "Catalan"},
	{ID: "hr", NameZH: "Croatian", NameEN: "Croatian"},
	{ID: "cs", NameZH: "Czech", NameEN: "Czech"},
	{ID: "et", NameZH: "Estonian", NameEN: "Estonian"},
	{ID: "fi", NameZH: "Finnish", NameEN: "Finnish"},
	{ID: "gl", NameZH: "Galician", NameEN: "Galician"},
	{ID: "ka", NameZH: "Georgian", NameEN: "Georgian"},
	{ID: "el", NameZH: "Greek", NameEN: "Greek"},
	{ID: "he", NameZH: "Hebrew", NameEN: "Hebrew"},
	{ID: "hu", NameZH: "Hungarian", NameEN: "Hungarian"},
	{ID: "is", NameZH: "Icelandic", NameEN: "Icelandic"},
	{ID: "kk", NameZH: "Kazakh", NameEN: "Kazakh"},
	{ID: "lv", NameZH: "Latvian", NameEN: "Latvian"},
	{ID: "lt", NameZH: "Lithuanian", NameEN: "Lithuanian"},
	{ID: "mk", NameZH: "Macedonian", NameEN: "Macedonian"},
	{ID: "mi", NameZH: "Māori", NameEN: "Māori"},
	{ID: "ne", NameZH: "Nepali", NameEN: "Nepali"},
	{ID: "fa", NameZH: "Persian", NameEN: "Persian"},
	{ID: "sr", NameZH: "Serbian", NameEN: "Serbian"},
	{ID: "sk", NameZH: "Slovak", NameEN: "Slovak"},
	{ID: "sl", NameZH: "Slovenian", NameEN: "Slovenian"},
	{ID: "tl", NameZH: "Filipino", NameEN: "Filipino"},
	{ID: "cy", NameZH: "Welsh", NameEN: "Welsh"},
	{ID: "amh", NameZH: "Amharic", NameEN: "Amharic"},
	{ID: "mya", NameZH: "Burmese", NameEN: "Burmese"},
	{ID: "yue", NameZH: "Cantonese (Traditional Chinese)", NameEN: "Cantonese (Traditional Chinese)"},
	{ID: "fil", NameZH: "Filipino", NameEN: "Filipino"},
	{ID: "gle", NameZH: "Irish", NameEN: "Irish"},
	{ID: "mon", NameZH: "Mongolian", NameEN: "Mongolian"},
	{ID: "som", NameZH: "Somali", NameEN: "Somali"},
	{ID: "zh-cn", NameZH: "Simplified Chinese", NameEN: "Simplified Chinese"},
	{ID: "zh-tw", NameZH: "Traditional Chinese (Taiwan)", NameEN: "Traditional Chinese (Taiwan)"},
	{ID: "zh-hk", NameZH: "Traditional Chinese (Hong Kong)", NameEN: "Traditional Chinese (Hong Kong)"},
}

var allowedRealtimeVoices = choiceSetVoices(voiceChoices)
var allowedVoiceModes = choiceSet(voiceModeChoices)
var allowedLanguageCodes = languageSet(languageChoices)

var realtimeVoiceAliases = map[string]string{
	"arbor":  "fathom",
	"sol":    "glimmer",
	"spruce": "orbit",
}

// Config returns a copy of the non-secret capability document.
func Config() PublicConfig {
	return PublicConfig{
		Version: "v1",
		Defaults: SessionDefaults{
			Voice:        defaultVoice,
			VoiceMode:    defaultVoiceMode,
			LanguageCode: defaultLanguageCode,
		},
		Voices:     append([]VoiceChoice(nil), voiceChoices...),
		VoiceModes: append([]Choice(nil), voiceModeChoices...),
		Languages:  append([]LanguageChoice(nil), languageChoices...),
		WebRTC: WebRTCConfig{
			DataChannel: DataChannelConfig{Label: "oai-events", Negotiated: true, ID: 0},
			ICEServers: []ICEServerConfig{
				{URLs: []string{"stun:stun.l.google.com:19302"}},
				{URLs: []string{"stun:stun4.l.google.com:19302"}},
			},
			ReceiveAudio: true,
			ReceiveVideo: false,
		},
	}
}

func normalizeSessionOptions(voice, voiceMode, languageCode string) (SessionDefaults, error) {
	voice = strings.ToLower(strings.TrimSpace(voice))
	if voice == "" {
		voice = defaultVoice
	}
	if alias, ok := realtimeVoiceAliases[voice]; ok {
		voice = alias
	}
	if _, ok := allowedRealtimeVoices[voice]; !ok {
		return SessionDefaults{}, &ServiceError{Message: "unsupported voice", StatusCode: 400}
	}

	voiceMode = strings.ToLower(strings.TrimSpace(voiceMode))
	if voiceMode == "" {
		voiceMode = defaultVoiceMode
	}
	if _, ok := allowedVoiceModes[voiceMode]; !ok {
		return SessionDefaults{}, &ServiceError{Message: "unsupported voice_mode", StatusCode: 400}
	}

	languageCode = strings.ToLower(strings.TrimSpace(languageCode))
	if languageCode == "" {
		languageCode = defaultLanguageCode
	}
	if _, ok := allowedLanguageCodes[languageCode]; !ok {
		return SessionDefaults{}, &ServiceError{Message: "unsupported language_code", StatusCode: 400}
	}

	return SessionDefaults{Voice: voice, VoiceMode: voiceMode, LanguageCode: languageCode}, nil
}

func choiceSet(items []Choice) map[string]struct{} {
	result := make(map[string]struct{}, len(items))
	for _, item := range items {
		result[item.ID] = struct{}{}
	}
	return result
}

func choiceSetVoices(items []VoiceChoice) map[string]struct{} {
	result := make(map[string]struct{}, len(items))
	for _, item := range items {
		result[item.ID] = struct{}{}
	}
	return result
}

func languageSet(items []LanguageChoice) map[string]struct{} {
	result := make(map[string]struct{}, len(items))
	for _, item := range items {
		result[item.ID] = struct{}{}
	}
	return result
}
