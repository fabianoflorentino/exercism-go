package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {

	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	/*
		The expression is composed of two parts separated by the logical operator "||" (or). Let's analyze each part individually:

		First part: (!archerIsAwake && !knightIsAwake && prisonerIsAwake)

		In this part, we have three conditions joined by the logical operator "&&" (and).
		The first condition !archerIsAwake checks if the archer is not awake.
		The second condition !knightIsAwake checks if the knight is not awake.
		The third condition prisonerIsAwake checks if the prisoner is awake.
		All three conditions need to be true for this part of the expression to return true.

		Second part: (petDogIsPresent && !archerIsAwake)

		In this part, we have two conditions joined by the logical operator "&&" (and).
		The first condition petDogIsPresent checks if the pet dog is present.
		The second condition !archerIsAwake checks if the archer is not awake.
		Both conditions need to be true for this part of the expression to return true.

		The complete expression returns true if at least one of the two parts is true.
		This means that if the prisoner is awake and neither the archer nor the knight are awake,
		or if the pet dog is present and the archer is not awake, the expression will return true. Otherwise, it will return false.
	*/
	return (!archerIsAwake && !knightIsAwake && prisonerIsAwake) || (petDogIsPresent && !archerIsAwake)
}
