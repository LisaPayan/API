# PROJET GROUPIE TRACKER

## DESCRIPTION 
Ce projet est un site web qui exploite une API ("https://restcountries.com"). Il est developpé en Go, HTML et CSS.

---

### MON THEME 
J'ai choisi une api sur tous les pays du monde. On peut découvrir certaines informations sur les pays du monde entier et ainsi les parcourir et les découvrir grâce à ce site.

---

## MES FONCTIONNALITES
1. **Recherche**
    - On peut  effectuer une recherche avec un nom (le nom officiel ou le nom utilisé)ou avec un continent(ce dernier écrit en anglais).

2. **Filtrage**
    - On peut filtrer les pays en fonction des continents, de la langue parlée, du nombre d'habitants minimum.
    - On avons la possibilité de cumuler les filtres.

3. **Pagination**
    - Cette fonction permet d'afficher les pays 15 par 15 avec la possibilité d'aller à la page suivante ou revenir à la page précédente.

4. **Favoris**
    - Avec cette dernière, nous avons la possibilté d'ajouter un pays en favoris, d'accéder aux pays mis en favoris et de les supprimer.
    - Cette fonction utilise un fichier JSON pour gérer la liste des favoris.

---    

## Architecture du site

- **Page d’accueil** : Présentation du thème et du site. ("/")  
- **Page collection** : Liste des pays avec la fonction recherche et favoris. ("/countries")
- **Page ressource (détails)** : Informations détaillées sur chaque pays. ("/countries/details") 
- **Page catégorie** : Affichage des ressources par catégorie grâce au système de filtre. ("/filter)  
- **Page favoris** : Gestion de la liste de favoris. ("/favoris") 
- **Page recherche** : Résultats de recherche. ("/countries/search") 
- **Page pagination** : Affiche tous les pays du monde 15 par 15. ("/countries/pagination") 
- **Page erreur** : Route de redirection en cas d'erreur. ("/error)

- **README** : En dehors, c'est le "guide pratique" pour comprendre et utiliser ce projet.(README.md)
- **A-propos** : En dehors, c'est un pdf qui permet d’avoir un retour sur le projet. (A_propos.pdf)

---

## API utilisée

- API REST utilisée : [REST Countries](https://restcountries.com/)  
- Endpoints exploités :  
  1. `GET /v3.1/all?fields=name,flags,region` – Récupération de tous les pays.  
  2. `GET /v3.1/name/{name}` – Récupération d’un pays par son nom.  
  3. `GET /v3.1/all?fields=name,flags,region,languages,population` - Récupération pays pour filtrage ou recherche selon les besoins.  

**Format JSON** pour toutes les requêtes, sans quota d’utilisation.  

---

## Installation et lancement

1. **Cloner le dépôt**
```bash
git clone <https://github.com/LisaPayan/API.git>
cd ./src
go run main.go