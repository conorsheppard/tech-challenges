package build_menu

import (
	"fmt"
	"testing"
)

func Test_BuildMenu(t *testing.T) {
	tests := []struct {
		expectedOutput string
		courseItems    map[string][]string
		chefDishes     []string
		restaurantName string
	}{
		{
			expectedOutput: ``,
			courseItems: map[string][]string{
				"antipasti": {},
				"mains":     {},
				"desserts":  {},
			},
			chefDishes:     []string{},
			restaurantName: "",
		},
		{
			expectedOutput: `~~~~ Nacho Average Mexican Diner ~~~~
             ** Menu **
Antipasti
tacos, chips or plantain
Mains
nachos, sushi, pasta or rice
Desserts
pears, churros or cheese`,
			courseItems: map[string][]string{
				"antipasti": {
					"tacos", "chips", "plantain", "quesadilla",
				},

				"mains": {
					"steak", "nachos", "chicken", "pizza",
					"sandwich", "taco", "sushi", "burger",
					"pork", "pasta", "wrap", "rice",
				},

				"desserts": {
					"pears", "churros", "cheese",
					"icecream", "sorbet", "cupcake",
				},
			},
			chefDishes: []string{
				"churros", "tacos", "pasta",
				"rice", "pears", "plantain", "chips",
				"cheese", "nachos", "sushi",
			},
			restaurantName: "Nacho Average Mexican Diner",
		},
	}

	var testsPassed, testsFailed int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)
		result := buildMenu(test.courseItems, test.chefDishes, test.restaurantName)

		if result != test.expectedOutput {
			fmt.Printf("FAIL: test %d\nExpected:\n%s\n\nGot:\n%s", i, test.expectedOutput, result) // replace with t.Errorf()
			testsFailed++
		} else {
			fmt.Printf("PASS: test %d\n", i)
			testsPassed++
		}
	}
	fmt.Printf("Total tests ran: %d\n", len(tests))
	fmt.Printf("Total tests passed: %d\n", testsPassed)
	fmt.Printf("Total tests failed: %d\n", testsFailed)
}
