package formating

import (
	"encoding/json"
	"strconv"

	"github.com/eiko-team/eiko/misc/structures"

	"go.mongodb.org/mongo-driver/bson"
)

func getFirstValue(data bson.M, keys []string) string {
	for _, str := range keys {
		if data[str] != nil {
			return data[str].(string)
		}
	}
	return ""
}

func getName(data bson.M) string {
	return getFirstValue(data, []string{"product_name", "product_name_fr", "product_name_en", "product_name_ha"})
}

func getCompagnyName(data bson.M) string {
	return getFirstValue(data, []string{"brands"})
}

func getAdditive(data bson.M) []string {
	return []string{""}
}

func getAllergen(data bson.M) []string {
	return []string{""}
}

func getEnergie(data bson.M) float64 {
	str := getFirstValue(data, []string{"nutriment.energy", "nutriment.energy_value", "nutriment.energy_serving"})
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getFat(data bson.M) float64 {
	str := getFirstValue(data, []string{"nutriment_level.fat", "nutriment.fat", "nutriment.fat_value", "nutriment.fat_serving"})
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getFiber(data bson.M) float64 {
	str := getFirstValue(data, []string{"nutriscore_points.fiber", "nutriscore_points.fiber_value"})
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getGlucides(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getLipides(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getProteins(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getSodium(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getSaturatedFat(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getSugarGlucides(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getEnergy(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getAlcool(data bson.M) float64 {
	str := ""
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getBack(data bson.M) string {
	return ""
}

func getComposition(data bson.M) string {
	return ""
}

func getFront(data bson.M) string {
	return ""
}

func getManufacturing(data bson.M) string {
	return ""
}

func getCode(data bson.M) []string {
	return []string{""}
}

func getCategories(data bson.M) []string {
	return []string{""}
}

func getTags(data bson.M) []string {
	return []string{""}
}

func getPackaging(data bson.M) []string {
	return []string{""}
}

func getIngredient(data bson.M) []string {
	return []string{""}
}

func getVitamins(data bson.M) []string {
	return []string{""}
}

func getNutriScore(data bson.M) string {
	return ""
}

func getGrammes(data bson.M) int {
	str := ""
	res, _ := strconv.Atoi(str)
	return res
}

func getMLitre(data bson.M) int {
	str := ""
	res, _ := strconv.Atoi(str)
	return res
}

func getLabel(data bson.M) []string {
	return []string{""}
}

func BsonToString(data bson.M) (string, error) {
	struc := structures.Consumable{
		Additive:      getAdditive(data),
		Alcool:        getAlcool(data),
		Allergen:      getAllergen(data),
		Back:          getBack(data),
		Categories:    getCategories(data),
		Code:          getCode(data),
		Company:       getCompagnyName(data),
		Energie:       getEnergie(data),
		Energy:        getEnergy(data),
		Fat:           getFat(data),
		Fiber:         getFiber(data),
		Front:         getFront(data),
		Glucides:      getGlucides(data),
		Grammes:       getGrammes(data),
		Ingredient:    getIngredient(data),
		Label:         getLabel(data),
		Lipides:       getLipides(data),
		Manufacturing: getManufacturing(data),
		MLitre:        getMLitre(data),
		Name:          getName(data),
		NutriScore:    getNutriScore(data),
		Packaging:     getPackaging(data),
		Proteins:      getProteins(data),
		SaturatedFat:  getSaturatedFat(data),
		Sodium:        getSodium(data),
		SugarGlucides: getSugarGlucides(data),
		Tags:          getTags(data),
		Vitamins:      getVitamins(data),
	}
	r, err := json.Marshal(struc)
	res := string(r)
	return res, err
}
