package lasagna

// PreparationTime calculates the total preparation time for a lasagna
// based on the number of layers and the average time per layer.
// If avgTimePerLayer is 0, a default value of 2 minutes per layer is used.
//
// Parameters:
//   - layers: A slice of strings representing the layers of the lasagna.
//   - avgTimePerLayer: An integer representing the average preparation time per layer in minutes.
//
// Returns:
//   - An integer representing the total preparation time in minutes.
func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}

	return len(layers) * avgTimePerLayer
}

// Quantities calculates the total quantity of noodles and sauce needed
// for a given list of lasagna layers.
//
// Each "noodles" layer requires 50 grams of noodles, and each "sauce"
// layer requires 0.2 liters of sauce. Other layer types are ignored.
//
// Parameters:
//
//	layers []string - A slice of strings representing the layers of the lasagna.
//
// Returns:
//
//	int - The total amount of noodles required in grams.
//	float64 - The total amount of sauce required in liters.
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// AddSecretIngredient takes two slices of strings: `friendList` and `myList`.
// It replaces the last element of `myList` with the last element of `friendList`,
// effectively adding a "secret ingredient" from the friend's list to your list.
// If `friendList` is empty, the function does nothing.
func AddSecretIngredient(friendList []string, myList []string) {
	if len(friendList) == 0 {
		return
	}

	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe scales the ingredient quantities for a lasagna recipe based on the desired number of portions.
//
// Parameters:
//   - quantities: A slice of float64 representing the original quantities of each ingredient.
//   - portions: An integer representing the desired number of portions. If portions is 0, the function returns nil.
//
// Returns:
//
//	A slice of float64 containing the scaled quantities of each ingredient. The scaling factor is calculated
//	as the desired number of portions divided by 2.0.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	if portions == 0 {
		return nil
	}

	scaledQuantities := make([]float64, len(quantities))
	scalingFactor := float64(portions) / 2.0

	for i, quantity := range quantities {
		scaledQuantities[i] = quantity * scalingFactor
	}

	return scaledQuantities
}
