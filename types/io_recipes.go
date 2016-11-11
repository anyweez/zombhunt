package types

type XmlRecipeList struct {
	Recipes []*XmlRecipe `xml:"recipe"`
}

type XmlRecipe struct {
	Name        string           `xml:"name,attr"`
	Count       uint32           `xml:"count,attr"`
	Ingredients []*XmlIngredient `xml:"ingredient"`
}

type XmlIngredient struct {
	Name  string `xml:"name,attr"`
	Count uint32 `xml:"count,attr"`
}
