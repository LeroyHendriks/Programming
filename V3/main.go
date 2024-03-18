package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var kentekens []Kenteken

type Kenteken struct {
	Name string `json:"kenteken"`
}

func lees_json() {
	// Lees JSON-bestand in
	jsonBytes, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Fout bij het lezen van het JSON-bestand:", err)
		return
	}

	// Unmarshal JSON-gegevens in Go-gegevensstructuur
	if err := json.Unmarshal(jsonBytes, &kentekens); err != nil {
		fmt.Println("Fout bij het omzetten van JSON naar Go-gegevensstructuur:", err)
		return
	}
}

func kenteken_toevoegen() {

	var inputPlate string

	fmt.Println("Voer uw kenteken in:")
	fmt.Scan(&inputPlate)

	lees_json()

	found := false
	for _, kenteken := range kentekens {
		if kenteken.Name == inputPlate {
			fmt.Printf("Het kenteken %s staat al in het JSON-bestand.\n", inputPlate)
			found = true
			break
		}
	}

	// Voeg het nieuwe kenteken toe aan de slice
	if !found {
		kentekens = append(kentekens, Kenteken{Name: inputPlate})

		// Marshall de geüpdatete gegevens terug naar JSON
		updatedJSON, err := json.Marshal(kentekens)
		if err != nil {
			fmt.Println("Fout bij het marshalling van JSON:", err)
			return
		}

		// Schrijf de gemarshallede gegevens terug naar het JSON-bestand
		if err := os.WriteFile("data.json", updatedJSON, 0644); err != nil {
			fmt.Println("Fout bij het schrijven naar het JSON-bestand:", err)
			return
		}
		fmt.Println("Kenteken toegevoegd en opgeslagen in het JSON-bestand.")
	}
}
func main() {
	// variabelen voor kenteken toevoegen
	var inputJaNee string

	// Kenteken toevoegen aan json
	fmt.Println("Wilt u een kenteken toevoegen? Ja/Nee")
	fmt.Scan(&inputJaNee)

	if inputJaNee == "ja" {
		kenteken_toevoegen()

	}

	lees_json()

	// Vraag gebruiker om het kenteken op te zoeken
	fmt.Println("Voer een kenteken in om te zoeken:")
	var searchPlate string
	fmt.Scan(&searchPlate)

	// Doorloop de kentekens en zoek naar het ingevoerde kenteken
	found := false
	for _, kenteken := range kentekens {
		if kenteken.Name == searchPlate {
			fmt.Printf("Kenteken %s gevonden, ", kenteken.Name)
			found = true
			break
		}
	}

	currentTime := time.Now()
	hour := currentTime.Hour()

	morgen := "Goedemorgen"
	middag := "Goedemiddag"
	nacht := "Goedeavond"

	if found {
		if hour >= 7 && hour < 12 {
			fmt.Println(morgen)
		} else if hour >= 12 && hour < 18 {
			fmt.Println(middag)
		} else if hour >= 18 && hour < 23 {
			fmt.Println(nacht)
		} else {
			fmt.Println("De parkeerplaats is 's nachts gesloten. Excuus voor het ongemak.")
		}
	} else {
		fmt.Println("Je hebt geen toegang tot de parkeerplaats. Excuses voor het ongemak.")
	}
}
