package formating

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/eiko-team/eiko/misc/structures"
)

func split(s string, sep rune) []string {
	var res []string
	var buff string
	var pos int
	var parenthese int
	for _, c := range s {
		if c == ' ' && pos == 0 {
			continue
		}
		if c == '(' {
			parenthese++
		}
		if c == ')' {
			parenthese--
		}
		if c == sep && parenthese == 0 {
			res = append(res, buff)
			buff = ""
			pos = 0
		} else {
			buff += string(c)
			pos++
		}
	}
	return append(res, buff)
}

func betterList(list string) []string {
	list = strings.Replace(list, ".", "", -1)
	tmp := split(list, ',')
	return tmp
}

func additiveList(list string) []string {
    // FORMAT: '([ elt -> lang:elt ] )+'
    // RETURN [elt, lang:elt, ...]
    list = strings.Replace(list, " ", "", -1)
    list = strings.Replace(list, "[", "", -1)
    list = strings.Replace(list, "]", ",", -1)
    l := split(list, ',')
    if len(l) == 0 {
        return nil
    }
    res := []string{}
    for _, val := range l {
        elts := strings.Split(val, "->")
        if len(elts) != 2 {
            // res = append(res, val)
            continue
        } else {
            res = append(res, elts...)
        }
    }
    return res
}

func namesToPos(names []string, colName string) int {
	for pos, val := range names {
		if val == colName {
			return pos
		}
	}
	return -1
}

func getElt(names []string, data []string, colName string) string {
	pos := namesToPos(names, colName)
	if pos == -1 {
		return ""
	}
	return data[pos]
}

func getElts(names []string, data []string, colNames []string) []string {
	res := make([]string, len(colNames))
	for i, val := range colNames {
		if val == "" {
			continue
		}
		res[i] = `{"` + val + `":"` + getElt(names, data, val) + `"}`
	}
	return res
}

func getAdditive(names []string, data []string) []string {
    res := additiveList(getElt(names, data, "additives"))
    if tag := getElt(names, data, "additives_tags"); len(tag) != 0 {
        tags := betterList(tag)
        res = append(res, tags...)
    }
    return res
}

func getAlcool(names []string, data []string) float64 {
	str := getElt(names, data, "alcohol_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getAllergen(names []string, data []string) []string {
    // TODO Import allergents from MongoDb database
    // BODY it might be usefull to import data from both database and csv file
	return betterList(getElt(names, data, ""))
}

func getBack(names []string, data []string) string {
	return getElt(names, data, "image_small_url")
}

func getCategories(names []string, data []string) []string {
	return betterList(getElt(names, data, "categories_fr"))
}

func getCode(names []string, data []string) []string {
	return []string{getElt(names, data, "code")}
}

func getCompany(names []string, data []string) string {
	return getElt(names, data, "brands")
}

func getEnergy(names []string, data []string) float64 {
	str := getElt(names, data, "energy_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getFat(names []string, data []string) float64 {
	str := getElt(names, data, "fat_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getFiber(names []string, data []string) float64 {
	str := getElt(names, data, "fiber_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getFront(names []string, data []string) string {
	return getElt(names, data, "image_url")
}

func getGlucides(names []string, data []string) float64 {
	str := getElt(names, data, "glucose_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getGrammes(names []string, data []string) int {
	str := getElt(names, data, "serving_size")
	if str == "" {
		return 0
	}
	res, _ := strconv.Atoi(str)
	return res
}

func getIngredient(names []string, data []string) []string {
	return betterList(getElt(names, data, "ingredients_text"))
}

func getLabel(names []string, data []string) []string {
	return getElts(names, data, []string{"labels", "labels_tags"})
}

func getManufacturing(names []string, data []string) string {
	return getElt(names, data, "manufacturing_places")
}

func getMLitre(names []string, data []string) int {
	// TODO: find proper volume
	str := getElt(names, data, "serving_size")
	if str == "" {
		return 0
	}
	res, _ := strconv.Atoi(str)
	return res
}

func getName(names []string, data []string) string {
	return getElt(names, data, "product_name")
}

func getNutriScore(names []string, data []string) string {
	return getElt(names, data, "nutrition_grade_fr")
}

func getPackaging(names []string, data []string) []string {
	return getElts(names, data, []string{"packaging", "packaging_tags"})
}

func getProteins(names []string, data []string) float64 {
	str := getElt(names, data, "proteins_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getSaturatedFat(names []string, data []string) float64 {
	str := getElt(names, data, "saturated-fat_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getSodium(names []string, data []string) float64 {
	str := getElt(names, data, "sodium_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getSugarGlucides(names []string, data []string) float64 {
	str := getElt(names, data, "sugars_100g")
	if str == "" {
		return 0
	}
	res, _ := strconv.ParseFloat(str, 32)
	return res
}

func getTags(names []string, data []string) []string {
	return getElts(names, data, []string{"categories", "categories_tags"})
}

func getVitamins(names []string, data []string) []string {
	return getElts(names, data, []string{"vitamin-a_100g", "vitamin-d_100g", "vitamin-e_100g", "vitamin-k_100g", "vitamin-c_100g", "vitamin-b1_100g", "vitamin-b2_100g", "vitamin-pp_100g", "vitamin-b6_100g", "vitamin-b9_100g", "vitamin-b12_100g"})
}

func ProductToString(names []string, data []string) (string, error) {
	struc := structures.Consumable{
		Additive:      getAdditive(names, data),
		Alcool:        getAlcool(names, data),
		Allergen:      getAllergen(names, data),
		Back:          getBack(names, data),
		Categories:    getCategories(names, data),
		Code:          getCode(names, data),
		Company:       getCompany(names, data),
		Energy:        getEnergy(names, data),
		Fat:           getFat(names, data),
		Fiber:         getFiber(names, data),
		Front:         getFront(names, data),
		Glucides:      getGlucides(names, data),
		Grammes:       getGrammes(names, data),
		Ingredient:    getIngredient(names, data),
		Label:         getLabel(names, data),
		Manufacturing: getManufacturing(names, data),
		MLitre:        getMLitre(names, data),
		Name:          getName(names, data),
		NutriScore:    getNutriScore(names, data),
		Packaging:     getPackaging(names, data),
		Proteins:      getProteins(names, data),
		SaturatedFat:  getSaturatedFat(names, data),
		Sodium:        getSodium(names, data),
		SugarGlucides: getSugarGlucides(names, data),
		Tags:          getTags(names, data),
		Vitamins:      getVitamins(names, data),
	}
	r, err := json.Marshal(struc)
	res := string(r)
	return res, err
}
