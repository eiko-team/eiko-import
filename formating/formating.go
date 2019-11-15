package formating

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/eiko-team/eiko/misc/structures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getFirstValue(data bson.M, keys []string) string {
	for _, str := range keys {
		if str == "" {
			continue
		}
		if data[str] != nil {
			return data[str].(string)
		}
	}
	return ""
}

func getAllValues(data bson.M, keys []string) []string {
	var res []string
	for _, str := range keys {
		if str == "" {
			continue
		}
		if data[str] != nil {
			if reflect.ValueOf(data[str]).Kind() == reflect.Slice {
				for _, d := range data[str].(primitive.A) {
					res = append(res, string(d.(string)))
				}
			}
		}
		if len(res) > 0 {
			// Got a match
			return res
		}
	}
	return res
}

func getNestedValues(data bson.M, key string, keys []string,
	selector string) string {
	var res string
	if key == "" || data[key] == nil {
		return ""
	}

	d := data[key].(bson.M)
	for _, str := range keys {
		if str == "" {
			continue
		}
		if d[str] != nil {
			for k, val := range d[str].(bson.M) {
				fmt.Println(k, val)
			}
			fmt.Println("\n", d[str].(bson.M),
				"\n",
				d[str].(bson.M)["sizes"],
				d[str].(bson.M)[selector],
				selector, "\n\n\n\n\n\n\n")
			panic(d[str].(bson.M)[selector].(string))
			return string(d[str].(bson.M)[selector].(string))
		}
	}
	return res
}

func getName(data bson.M) string {
	return getFirstValue(data, []string{"product_name", "product_name_fr", "product_name_en", "product_name_ha"})
}

func getCompagnyName(data bson.M) string {
	return getFirstValue(data, []string{"brands"})
}

func getAdditive(data bson.M) []string {
	return getAllValues(data, []string{"additives_n", "additives_tags", "additives_old_n", "additives_old_tags", "additives_original_tags", "additives_prev_original_tags", "additives_debug_tags"})
}

func getAllergen(data bson.M) []string {
	return getAllValues(data, []string{"allergens", "allergens_from_ingredients", "allergens_from_user", "allergens_tags", "allergens_hierarchy"})
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
	// return getFirstValue(data, []string{"images.ingredients_fr"})
	// str := getNestedValues(data, "images", []string{"ingredients_fr"}, "full")
	// if str != "" {
	// 	panic(str)
	// }
	str := getFirstValue(data, []string{"image_url", "image_small_url"})
	if str != "" {
		panic(str)
	}
	return str
}

func getComposition(data bson.M) string {
	return getFirstValue(data, []string{""})
}

func getFront(data bson.M) string {
	return getFirstValue(data, []string{""})
}

func getManufacturing(data bson.M) string {
	return getFirstValue(data, []string{""})
}

func getCode(data bson.M) []string {
	return getAllValues(data, []string{""})
}

func getCategories(data bson.M) []string {
	return getAllValues(data, []string{""})
}

func getTags(data bson.M) []string {
	return getAllValues(data, []string{""})
}

func getPackaging(data bson.M) []string {
	return getAllValues(data, []string{""})
}

func getIngredient(data bson.M) []string {
	return getAllValues(data, []string{""})
}

func getVitamins(data bson.M) []string {
	return getAllValues(data, []string{""})
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
	return getAllValues(data, []string{""})
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
