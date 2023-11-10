package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Promo struct {
	Nom     string
	Filiere string
	Niveau  int
	Nombre  int
}

type Etudiant struct {
	Nom    string
	Prenom string
	Sexe   int
	Age    int
}

type PageData struct {
	Promo         Promo
	ListeEtudiant []Etudiant
}

type UserInfo struct {
	Nom           string
	Prenom        string
	DateNaissance string
	Sexe          string
}

var (
	globalUserInfo UserInfo
)

func main() {

	temp, err := template.ParseGlob("./*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
	}

	// CHALLENGE 1 : /promo
	ListeEtudiant := []Etudiant{}
	ListeEtudiant = append(ListeEtudiant, Etudiant{"RODRIGUES", "Cyril", 1, 22})
	ListeEtudiant = append(ListeEtudiant, Etudiant{"MEDERREG", "Kheir-eddine", 0, 22})
	ListeEtudiant = append(ListeEtudiant, Etudiant{"PHILIPIERT", "Alan", 1, 22})

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		dataPage := PageData{
			Promo:         Promo{"Mentor'ac", "Informatique", 5, 3},
			ListeEtudiant: ListeEtudiant,
		}
		temp.ExecuteTemplate(w, "promo", dataPage)
	})

	// CHALLENGE 3 : /user/init

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "userInit", nil)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Erreur lors de la lecture des données du formulaire", http.StatusInternalServerError)
				return
			}

			userInfo := UserInfo{
				Nom:           r.FormValue("nom"),
				Prenom:        r.FormValue("prenom"),
				DateNaissance: r.FormValue("dateNaissance"),
				Sexe:          r.FormValue("sexe"),
			}

			globalUserInfo = userInfo

			http.Redirect(w, r, "/user/display", http.StatusSeeOther)
		} else {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "userDisplay", globalUserInfo)
	})

	// Gestion des fichiers dans assets
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// Serveur
	fmt.Println("Serveur lancé sur : http://localhost:8080")
	fmt.Println("Challenge1 ➡️ http://localhost:8080/promo")
	fmt.Println("Challenge3 ➡️ http://localhost:8080/user/init")
	http.ListenAndServe("localhost:8080", nil)
}
