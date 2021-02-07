package main

import (
	"os"
	"sort"
	"strings"
	"text/template"
)

//-----------------------------------------------------
//  Provided data

var catalogTemplate = `
======== Film Catalog ========
      
English
	Love Actually, Fish Called Wanda or Harry Potter

French
	Amelie, MonOncle, Breathless or Amour

Czech
	Ucho, Pelisky or Ostre Sledovane Vlaky

==============================
`

var filmCategories = map[string][]string{
	"english": {
		"Fish Called Wanda", "bread", "Harry Potter", "Love Actually",
	},

	"french": {
		"Prophet", "Amelie", "Cache", "La Grande Illusion",
		"Les Miserable", "The Artist", "Mon Oncle", "Van Gogh",
		"Pierrot Le Fou", "Breathless", "Buffet Froid", "Amour",
	},

	"czech": {
		"Ucho", "Pelisky", "Ostre Sledovane Vlaky",
		"Svetaci", "Postriziny", "Slavnost snezenek",
	},
}

var cinemaCategories = []string{
	"Pelisky", "Love Actually", "Breathless",
	"Amour", "Ucho", "Harry Potter", "Fish Called Wanda",
	"Ostre Sledovane Vlaky", "Amelie", "Mon Oncle",
}

var cinemaFilms = map[string]float32{
	"Harry Potter":  5,
	"Pelisky":       2.5,
	"Ucho":          3.25,
	"Breathless":    3,
	"Amelie":        2,
	"Amour":         1.75,
	"Love Actually": 2.5,
}

//-----------------------------------------------------

type categories struct {
	English string
	French  string
	Czech   string
}

func main() {
	err := printOutProgram(filmCategories)
	if err != nil {
		panic(err)
	}

	cinemaCats := createCinemaFilmCategories(filmCategories, cinemaCategories)
	err = printOutProgram(cinemaCats)
	if err != nil {
		panic(err)
	}

	ordered := orderFilmsByLength(cinemaCats, cinemaFilms)
	err = printOutProgram(ordered)
	if err != nil {
		panic(err)
	}
}

func printOutProgram(fc map[string][]string) error {
	categories := categories{
		English: collate(fc["english"]),
		French:  collate(fc["french"]),
		Czech:   collate(fc["czech"]),
	}

	t, err := template.New("catalog").
		Parse(`
======== Film Catalog ========
      
English
	{{ .English }}

French
	{{ .French }}

Czech
	{{ .Czech }}

==============================
`)
	if err != nil {
		return err
	}

	return t.Execute(os.Stdout, categories)
}

func collate(items []string) string {
	l := len(items)
	toPrint := strings.Join([]string{strings.Join(items[:(l-1)], ", "), items[l-1]}, " or ")

	return toPrint
}

func createCinemaFilmCategories(filmCats map[string][]string, cinemaCategories []string) map[string][]string {
	cinemaCats := make(map[string][]string)

	for _, chd := range cinemaCategories {
		for cat, items := range filmCats {
			for _, it := range items {
				if chd == it {
					cinemaCats[cat] = append(cinemaCats[cat], it)
				}
			}
		}
	}
	return cinemaCats
}

func orderFilmsByLength(fc map[string][]string, fpt map[string]float32) map[string][]string {
	ordered := make(map[string][]string)
	for c, its := range fc {
		sort.Slice(its, func(i, j int) bool {
			l := its[i]
			r := its[j]

			// films whose length is unknown are shown at the end
			if fpt[l] == 0.0 {
				return fpt[l] > fpt[r]
			}
			return fpt[l] < fpt[r]
		})
		ordered[c] = its
	}
	return ordered
}
