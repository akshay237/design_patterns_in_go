package main

import "fmt"

func main() {
	espanol := SpanishSpeaker{}
	englishSpeaker := EnglishSpeaker{}

	fmt.Println("Without translation")
	fmt.Println("Espanol says: ", espanol.greetInSpanish())
	fmt.Println("English women says: ", englishSpeaker.greetInEnglish())

	adaptedEnglishSpeaker := englishToSpanishAdapter{
		speaker: englishSpeaker,
	}

	fmt.Println("With translation")
	fmt.Println("Espanol says: ", espanol.greetInSpanish())
	fmt.Println("English women says: ", englishSpeaker.greetInEnglish())
	fmt.Println("English woman says in spanish: ", adaptedEnglishSpeaker.greetInSpanish())

}
