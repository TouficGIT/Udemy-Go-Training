package Player

// Rappel :
// Majuscule = publique
// Minuscule = privé

//ex: "password" est privé

//Avatar : structure of a player avater
type Avatar struct {
	Url string
}

//Player : structure of a player
type Player struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}

// Une bonne pratique pour l'initialisation d'une struc :
// => Créer un fonc qui renvoit une nvelle instance de notre struct

//New function : Retrieve a new player struct
func New(name string) Player {
	return Player{
		Name:     name,
		password: "defaultPassword",
	}
}
