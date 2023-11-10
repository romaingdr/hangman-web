package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func main() {

	temp, err := template.ParseGlob("./*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
	}

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

	ListeEtudiant := []Etudiant{}
	ListeEtudiant = append(ListeEtudiant, Etudiant{"RODRIGUES", "Cyril", 1, 22})
	ListeEtudiant = append(ListeEtudiant, Etudiant{"MEDERREG", "Kheir-eddine", 0, 22})
	ListeEtudiant = append(ListeEtudiant, Etudiant{"PHILIPIERT", "Alan", 1, 22})

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		dataPage := PageData{
			Promo:         Promo{"Mentor'ac", "Informatique", 5, 3},
			ListeEtudiant: ListeEtudiant,
		}
		temp.ExecuteTemplate(w, "index", dataPage)
	})
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
