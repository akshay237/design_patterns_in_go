package main

type englishToSpanishAdapter struct {
	speaker EnglishSpeaker
}

func translate(msg string) string {
	return "¡Hola!"
}

func (e *englishToSpanishAdapter) greetInSpanish() string {
	spanishMessage := translate(e.speaker.greetInEnglish())
	return spanishMessage
}
