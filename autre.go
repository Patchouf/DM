package autre

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func autre() {
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
		var addedLetter string
		add := false

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
			if add {
				// Si nous devons ajouter une lettre vérifier si c'est la fin de la séquence d'ajout ou non
				if letter == '+' {
					add = false
					// Traiter la lettre ajoutée avant de passer au caractère suivant
					secretMessage += secretDictionary[string(addedLetter)] + secretDictionary[string(letter)]
					//addedLetter = ""
					continue
				} else {
					// Ajouter la lettre en cours de traitement à la variable addedLetter
					addedLetter += string(letter) + secretDictionary[string(addedLetter)] + secretDictionary[string(letter)]
					// Passer au caractère suivant
					continue
				}
			}

			// // Si le caractère en cours de traitement est un '+' indique que nous devons ajouter une lettre et passer au caractère suivant
			if letter == '+' {

				add = true
				continue
			}

			// Si le caractère saisie est un espace ca reinitialise la variable addedLetter
			if string(letter) == " " {
				addedLetter = ""
			}

			// Ajouter la lettre ajoutée et la traduction correspondante du dictionnaire à la chaîne de sortie
			secretMessage += secretDictionary[string(letter)]
			//addedLetter = ""
		}
		// Ajouter la lettre ajouté à la fin de la chaîne de sortie
		secretMessage += addedLetter + secretDictionary[string(addedLetter)]

		fmt.Println("Message secret :", secretMessage)
	}
}
