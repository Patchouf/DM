package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Dictionnaire qui associe chaque caractère de l'alphabet à sa traduction dans le code secret
	secretDictionary := map[string]string{
		"a": "a",
		"b": "bé",
		"c": "cé",
		"d": "dé",
		"e": "eu",
		"f": "aife",
		"g": "gé",
		"h": "ache",
		"i": "i",
		"j": "ji",
		"k": "ka",
		"l": "aile",
		"m": "aime",
		"n": "aine",
		"o": "o",
		"p": "pé",
		"q": "cu",
		"r": "aire",
		"s": "aisse",
		"t": "té",
		"u": "ou",
		"v": "vé",
		"w": "doublevé",
		"x": "ixe",
		"y": "igrec",
		"z": "zède",
	}

	// lire les lettres entres
	lectureText := bufio.NewReader(os.Stdin)

	// Boucle infinie qui demandera à l'utilisateur d'entrer un code secret jusqu'à ce qu'il entre "q" pour quitter
	for {
		fmt.Print("Entrez le code secret (ou tapez 'q' pour quitter) : ")
		secretCode, _ := lectureText.ReadString('\n')
		secretCode = strings.TrimSpace(secretCode)
		if secretCode == "q" {
			break
		}

		// Initialisation de la chaîne de sortie à vide
		secretMessage := ""

		// Indicateurs pour savoir si nous devons ignorer un caractère ou ajouter un caractère au début de la chaîne de sortie
		ignore := false

		adda := false

		// Pour chaque caractère du code secret
		for _, letter := range secretCode {
			if ignore {
				// Vérifier si c'est la fin de la séquence d'ignore
				if letter == '-' {
					ignore = false
				} else {
					// Passer au caractère suivant
					continue
				}
			}

			// Si le caractère est un "-", activer l'ignore
			if letter == '-' {
				ignore = true
				continue
			}

			if adda {
				//var add string
				// Si nous devons ajouter un caractère au début de la chaîne de sortie,
				// vérifier si c'est la fin de la séquence d'ajout
				if string(letter) == "+" {
					//add =string(letter)
					adda = false
				}
				if string(letter) == " " {
					adda = true

				}

				// Ajouter le caractère au début de la chaîne de sortie
				secretMessage = secretMessage + string(letter) //+ add
				continue
			}

			// Si aucun des indicateurs n'est activé, ajouter la traduction correspondante au dictionnaire à la chaîne de sortie
			secretMessage += secretDictionary[string(letter)]
		}
		fmt.Println("Message secret :", secretMessage)
	}
}
