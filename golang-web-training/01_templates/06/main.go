package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type beer struct {
	Name    string
	Brewery string
	Type    string
	Alc     int
}

type cider struct {
	Name    string
	Brewery string
	Fruit   string
	Alc     int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	a := beer{
		Name:    "CBS",
		Brewery: "Founders",
		Type:    "stout",
		Alc:     11,
	}

	b := beer{
		Name:    "Pirkka Lager",
		Brewery: "Pirkka",
		Type:    "lager",
		Alc:     4,
	}

	c := beer{
		Name:    "Joku Ipa",
		Brewery: "Stadin Panimo",
		Type:    "IPA",
		Alc:     7,
	}

	d := cider{
		Name:    "Situkka",
		Brewery: "Jooh",
		Fruit:   "apple",
		Alc:     4,
	}

	e := cider{
		Name:    "Päärynä siideri",
		Brewery: "Maatila",
		Fruit:   "pear",
		Alc:     5,
	}

	beers := []beer{a, b, c}
	ciders := []cider{d, e}

	data := struct {
		Drinks []beer
		Extras []cider
	}{
		beers,
		ciders,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
