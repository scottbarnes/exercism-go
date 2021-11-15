package allergies

// type allergy struct {
// 	name  string
// 	value uint
// }

var allergens = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// AllergicTo takes a total allergy number and verifies if a person is allergic
// to a specific allergen.
func AllergicTo(allergies uint, allergen string) bool {
	for k := range allergens {
		// Powers of two always have unique combinations in binary, because each power is of the form:
		// 0001 = 1
		// 0010 = 2
		// 0100 = 4
		// 1000 = 8, etc.
		// Therefore if total&allergen return the allergen, the allergen is part of the unique
		// combination that comprises the total. e.g., if the total is 9:
		// 1001 = 9
		// ----
		// 0001 = 1
		// 1000 = 8
		if allergens[k&int(allergies)] == allergen {
			return true
		}
	}

	return false
}

// Allergies takes a total allergy number and returns a string of all matched
// allergens.
func Allergies(allergiesTotal uint) []string {
	var result []string
	for _, v := range allergens {
		if AllergicTo(allergiesTotal, v) {
			result = append(result, v)
		}
	}
	return result
}
