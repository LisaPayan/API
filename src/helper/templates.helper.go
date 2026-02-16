package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// Variable globale qui contiendra tous les templates chargés
var listeTemplate *template.Template

// Load charge tous les fichiers HTML depuis le dossier ../templates
func Load() {
	// Chargement des fichiers .html dans le dossier templates
	temp, tempErr := template.ParseGlob("../templates/*.html")
	if tempErr != nil {
		// En cas d'erreur, le programme s'arrête avec un message d'erreur
		log.Fatalf("Erreur template - %s", tempErr.Error())
		return
	}
	// Affectation des templates à la variable globale
	listeTemplate = temp
	fmt.Println("Template - chargement des templates terminé")
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer

	errRender := listeTemplate.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		http.Redirect(
			w,
			r,
			fmt.Sprintf(
				"/error?code=%d&message=%s",
				http.StatusInternalServerError,
				url.QueryEscape("Erreur lors du chargement de la page"),
			),
			http.StatusSeeOther,
		)
		return
	}
	_, _ = buffer.WriteTo(w)
}
