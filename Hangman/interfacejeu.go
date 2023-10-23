package hangman

import (
	"fmt"
	"strings"
)

// Fonction d'Interface de jeu
func InterfaceJeu() {
	LettresTrouvees = []string{}
	//Supprime le fichier mot.txt pour éviter les conflits
	SupprimerMot()
	// Ecrit le mot aléatoire dans le fichier mot.txt (en le créant)
	Ecriremot()
	Word := LireFichierMot()
	fmt.Println("Le mot à deviner : ")
	fmt.Println("")
	//On affiche le mot en tiret + espace
	tentatives := 9
	var MotTrouvé bool
	var choixlettres string
	// tant que les tentatives sont au dessus de 0 et le mot n'est pas trouvé boucle continue.
	for tentatives > 0 && MotTrouvé == false {
		AffichageTirets(Word)
		AffichagePendu(tentatives)
		fmt.Println("")
		fmt.Print("Entrez une lettre ou le mot entier : ")
		fmt.Scan(&choixlettres)
		// CAS SI le joueur indique une lettre
		if len(choixlettres) == 1 {
			if strings.Contains(Word, choixlettres) {
				fmt.Println("Vous avez trouvé une lettre")
				LettresTrouver(choixlettres)
			} else {
				fmt.Println("Cette lettre n'est pas dans le mot")
				// décremente les tentatives car le joueur se trompe
				tentatives--
			}
		}
		//CAS SI le joueur indique un mot
		if len(choixlettres) > 1 {
			if Word == choixlettres {
				fmt.Println("Vous avez trouvé le mot")
				//MOT TROUVE = TRUE donc la boucle se termine
				MotTrouvé = true
			} else {
				fmt.Println("Le mot n'est pas bon")
				// LE MOT n'est pas bon donc les tentavies se réduisent de 2
				tentatives -= 2
			}
		}

	}
	//FIN DE LA BOUCLE avec mot = TRUE donc gagné
	if MotTrouvé == true {
		fmt.Println("Vous avez gagné")
		fmt.Println("")
		fmt.Println("")
		InterfaceJeu()
	} else {
		// FIN DE LA BOUCLE avec tentatives = 0 donc perdu
		fmt.Println(` 
 +-------+  
     |   |  
     O   |  
    /|\  |  
    / \  |  
	 |  
==========`)
		fmt.Println("")
		fmt.Println("Vous n'avez plus de tentatives")
		fmt.Println("")
		fmt.Printf("Le mot était %s", Word)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		InterfaceJeu()
	}

}
