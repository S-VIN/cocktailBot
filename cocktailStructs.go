package main

import (
	"encoding/json"
)

type Cocktail struct {
	IdDrink             string
	StrDrink            string
	StrDrinkAlternate   string
	StrTags             string
	StrVideo            string
	StrCategory         string
	StrIBA              string
	StrAlcoholic        string
	StrGlass            string
	StrInstructions     string
	StrDrinkThumb       string
	StrIngredient1      string
	StrIngredient2      string
	StrIngredient3      string
	StrIngredient4      string
	StrIngredient5      string
	StrIngredient6      string
	StrIngredient7      string
	StrIngredient8      string
	StrIngredient9      string
	StrIngredient10     string
	StrIngredient11     string
	StrIngredient12     string
	StrIngredient13     string
	StrIngredient14     string
	StrIngredient15     string
	StrMeasure1         string
	StrMeasure2         string
	StrMeasure3         string
	StrMeasure4         string
	StrMeasure5         string
	StrMeasure6         string
	StrMeasure7         string
	StrMeasure8         string
	StrMeasure9         string
	StrMeasure10        string
	StrMeasure11        string
	StrMeasure12        string
	StrMeasure13        string
	StrMeasure14        string
	StrMeasure15        string
	StrImageSource      string
	StrImageAttribution string
}
type Cocktails struct {
	Drinks []Cocktail
}


func getRandomCocktail() (Cocktail, error){
var cocktail Cocktail
	resp, err := getRequest("www.thecocktaildb.com/api/json/v1/1/random.php")
	
	if (err != nil){
		return cocktail, err
	}
	
	var temp Cocktails
	err = json.Unmarshal([]byte(resp), &temp)
	if (err != nil){
		return cocktail, err
	}	

	return temp.Drinks[0], err
}
