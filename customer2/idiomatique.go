package customer2

/* dans ce pckge nous reprenons l'exmple du package customer pour utiliser go idiomatique.
en Go idiomatique les champs de la struct sont publiques et un constructeur est utilisé
pour de raisons de conception et de reutilisabilité

Sans constructeur c'est possible, mais c'est limité
c := objects.Customer{
	UUID:  uuid.New().String(),
	Lastname:  lastname,
	Firstname: firstname,
	Email:     email,
}

❌ Aucun contrôle
❌ Répétition de logique partout
❌ Difficile à faire évoluer

Il est recommandé d'utiliser constructeur  même si les champs de la structure sont publiques
c := objects.NewCustomer(id, lastname, firstname, email)

*/
