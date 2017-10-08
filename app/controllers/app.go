package controllers

import (
	"braineedisp/app/models"
	"database/sql"
	"fmt"
	//Driver sql.
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

//App : Contrôleur revel
type App struct {
	*revel.Controller
}

//DB : variable d'accès à la BDD
var DB *sql.DB

//InitDB : ouvre l'accès à la BDD
func InitDB() {
	connstring := "golang@tcp(127.0.0.1:3306)/brain"

	var err error
	DB, err = sql.Open("mysql", connstring)
	if err != nil {
		fmt.Println("DB Error", err)
		revel.INFO.Println("DB Error", err)
	}

	fmt.Println("DB Connected")
	revel.INFO.Println("DB Connected")

	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Error", err)
		revel.INFO.Println("DB Error", err)
	}
}

//allons chercher dans la base de données
func getBrainee(b models.Brainee) models.Brainee {
	//envoi de la requête
	sql := "SELECT * FROM brainee WHERE id = ?"
	rows, err := DB.Query(sql, b.ID) //id est une clé primaire auto-incrémentée
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close() //idiomatique

	//on récupère les données
	for rows.Next() {
		err := rows.Scan(&b.ID, &b.Author, &b.Text, &b.Brand) //alors on assigne avec des pointeurs
		if err != nil {
			fmt.Println(err) //sinon c'est l'erreur
		}
		revel.INFO.Println(b.ID, b.Author, b.Text, b.Brand) //résultat dans la console
		return b
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err) //ou erreur dans la console
	}
	revel.INFO.Println(b.ID, b.Author, b.Brand, b.Text)

	return b
}

//écrivons dans la BDD
func sendBrainee(b models.Brainee) models.Brainee {
	sql := "INSERT INTO brainee (author, text, brand) VALUES (?, ?, ?)"
	_, err := DB.Exec(sql, b.Author, b.Text, b.Brand)
	if err != nil {
		panic(err)
	}
	//récupérons l'id pour afficher TOUTES les informations du brainee !
	sql = "SELECT id FROM brainee WHERE author = ? AND text = ? AND brand = ?"
	rows, err := DB.Query(sql, b.Author, b.Text, b.Brand)
	for rows.Next() {
		err := rows.Scan(&b.ID)
		if err != nil {
			revel.ERROR.Println(err)
		}
		return b
	}
	err = rows.Err()
	if err != nil {
		revel.ERROR.Println(err)
	}
	return b
}

//Index : il fallait en faire le rendu
func (c App) Index(id string, author string, brand string, text string) revel.Result {
	return c.Render()
}

/**
 * Brainees : Gérer les requêtes sur une deuxième page
 * création ou fetch du Brainee dans la BDD, affichage du texte dans le HTML
 */
func (c App) Brainees(id string, author string, brand string, text string) revel.Result {
	InitDB()

	var b models.Brainee
	var status string

	status = ""

	//dans le cas où on envoie un ID pour trouver un brainee
	if id != "" {
		b.ID = id

		b = getBrainee(b)

		status = "Voici le brainee que vous avez demandé !"
		id = b.ID
		author = b.Author
		brand = b.Brand
		text = b.Text

		if author == "" && brand == "" && text == "" {
			status = "Aucun brainee trouvé avec cet id : "
		}

		return c.Render(status, id, author, brand, text)
	}

	//si on veut plutôt en envoyer un
	if author != "" && brand != "" && text != "" {
		b.Author = author
		b.Brand = brand
		b.Text = text

		b = sendBrainee(b)

		status = "Votre brainee a bien été envoyé ! Vous pouvez le visualiser ici :"
		id = b.ID
		author = b.Author
		brand = b.Brand
		text = b.Text

		return c.Render(status, id, author, brand, text)
	}

	return c.Render(status, id, author, brand, text)

}
