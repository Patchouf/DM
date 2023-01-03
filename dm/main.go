package piscine

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
		sup := false
		textenmoins := false

		// Pour chaque caractère du code secret
		for _, letter := range secretCode {
			if sup {
				// Si nous devons ignorer un caractère, vérifier si c'est la fin de la séquence d'ignore
				if string(letter) == "-" {
					sup = false
				}
				if string(letter) == "-" {
					sup = true
				}
				// Passer au caractère suivant
				continue
			}

			if textenmoins {
				// Si nous devons ajouter un caractère au début de la chaîne de sortie,
				// vérifier si c'est la fin de la séquence d'ajout
				if string(letter) == "+" {
					textenmoins = false
				}
				if string(letter) == " " {
					textenmoins = true
				}
				// Ajouter le caractère au début de la chaîne de sortie
				secretMessage = string(letter) + secretMessage
				continue
			}

			// Si aucun des indicateurs n'est activé, ajouter la traduction correspondante au dictionnaire à la chaîne de sortie
			secretMessage += secretDictionary[string(letter)]
		}
		fmt.Println("Message secret :", secretMessage)
	}
}
