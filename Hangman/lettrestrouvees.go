package hangman

var LettresTrouvees []string

// Si la lettre est trouvée, l'envoi dans le slice LettresTrouvees
func LettresTrouver(choixlettres string) []string {
	LettresTrouvees = append(LettresTrouvees, choixlettres)
	return LettresTrouvees
}
