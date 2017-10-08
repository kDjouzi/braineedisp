package models

//Brainee : structure des brainees
type Brainee struct {
	//Première lettre majuscule car ce sont des données publiques
	ID     string //ID plutôt que Id : idiomatique
	Author string
	Brand  string
	Text   string
}
