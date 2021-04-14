package main

import (
	"encoding/json"
	//"fmt"
)

type Ingredient struct {
	IdIngredient	string
	StrIngredient	string
	StrDescription	string
	StrType			string	
	StrAlcohol		string
	StrABV			string
}

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
	Ingridients     [15]string
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
	Meashures       [15]string
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
	listOfIngridients   string
}

func squeezeCocktail(cocktail *Cocktail){
	cocktail.Ingridients[0] = cocktail.StrIngredient1
	cocktail.Ingridients[1] = cocktail.StrIngredient2
	cocktail.Ingridients[2] = cocktail.StrIngredient3
	cocktail.Ingridients[3] = cocktail.StrIngredient4
	cocktail.Ingridients[4] = cocktail.StrIngredient5
	cocktail.Ingridients[5] = cocktail.StrIngredient6
	cocktail.Ingridients[6] = cocktail.StrIngredient7
	cocktail.Ingridients[7] = cocktail.StrIngredient8
	cocktail.Ingridients[8] = cocktail.StrIngredient9
	cocktail.Ingridients[9] = cocktail.StrIngredient10
	cocktail.Ingridients[10] = cocktail.StrIngredient11
	cocktail.Ingridients[11] = cocktail.StrIngredient12
	cocktail.Ingridients[12] = cocktail.StrIngredient13
	cocktail.Ingridients[13] = cocktail.StrIngredient14
	cocktail.Ingridients[14] = cocktail.StrIngredient15
	cocktail.Meashures[0] = cocktail.StrMeasure1
	cocktail.Meashures[1] = cocktail.StrMeasure2
	cocktail.Meashures[2] = cocktail.StrMeasure3
	cocktail.Meashures[3] = cocktail.StrMeasure4
	cocktail.Meashures[4] = cocktail.StrMeasure5
	cocktail.Meashures[5] = cocktail.StrMeasure6
	cocktail.Meashures[6] = cocktail.StrMeasure7
	cocktail.Meashures[7] = cocktail.StrMeasure8
	cocktail.Meashures[8] = cocktail.StrMeasure9
	cocktail.Meashures[9] = cocktail.StrMeasure10
	cocktail.Meashures[10] = cocktail.StrMeasure11
	cocktail.Meashures[11] = cocktail.StrMeasure12
	cocktail.Meashures[12] = cocktail.StrMeasure13
	cocktail.Meashures[13] = cocktail.StrMeasure14
	cocktail.Meashures[14] = cocktail.StrMeasure15
}

type Cocktails struct {
	Drinks [1]Cocktail //at least one empty element. because err returned
}

type Ingreds struct{
	Ingredients [1]Ingredient
}

func getRandomCocktail() (Cocktail, error) {
	var cocktails Cocktails
	resp, err := getRequest("https://www.thecocktaildb.com/api/json/v1/1/random.php")

	if err != nil {
		return cocktails.Drinks[0], err
	}

	err = json.Unmarshal([]byte(resp), &cocktails)
	if err != nil {
		return cocktails.Drinks[0], err
	}
	
	squeezeCocktail(&cocktails.Drinks[0])
	return cocktails.Drinks[0], err
}

func searchByIngredient(ingredient string) (Cocktails, error){
	var cocktails Cocktails
	resp, err := getRequest("https://www.thecocktaildb.com/api/json/v1/1/filter.php?i=" + ingredient)

	if err != nil {
		return cocktails, err
	}

	err = json.Unmarshal([]byte(resp), &cocktails)
	if err != nil {
		return cocktails, err
	}

	return cocktails, err
}

func lookUpIngredientById(id string) (Ingredient, error){
	var ingredients Ingreds
	
	resp, err := getRequest("https://www.thecocktaildb.com/api/json/v1/1/lookup.php?iid=" + id)

	if err != nil {
		return ingredients.Ingredients[0], err
	}

	err = json.Unmarshal([]byte(resp), &ingredients)
	if err != nil {
		return ingredients.Ingredients[0], err
	}

	return ingredients.Ingredients[0], err
}

func lookUpFullCocktailDetailById(id string) (Cocktail, error){
	var cocktails Cocktails
	resp, err := getRequest("https://www.thecocktaildb.com/api/json/v1/1/lookup.php?i=" + id)

	if err != nil {
		return cocktails.Drinks[0], err
	}

	err = json.Unmarshal([]byte(resp), &cocktails)
	if err != nil {
		return cocktails.Drinks[0], err
	}

	return cocktails.Drinks[0], err
}


func searchIngredientByName(name string) (Ingredient, error){
	var ingredients Ingreds
	
	resp, err := getRequest("https://www.thecocktaildb.com/api/json/v1/1/search.php?i=" + name)

	if err != nil {
		return ingredients.Ingredients[0], err
	}

	err = json.Unmarshal([]byte(resp), &ingredients)
	if err != nil {
		return ingredients.Ingredients[0], err
	}

	return ingredients.Ingredients[0], err
}

func searchCocktailByName(name string) (Cocktail, error){
	var cocktails Cocktails

	resp, err := getRequest("https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + name)

	if err != nil {
		return cocktails.Drinks[0], err
	}

	err = json.Unmarshal([]byte(resp), &cocktails)
	if err != nil {
		return cocktails.Drinks[0], err
	}

	return cocktails.Drinks[0], err
}