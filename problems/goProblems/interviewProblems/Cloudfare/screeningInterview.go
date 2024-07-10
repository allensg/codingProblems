package main

// happy path
// [512, 800, 888], 5 => [5121234567, 8001234567, 8881234567, 5127654321, 8007654321]

// one number for 800
// [512, 800, 888], 5 => [5121234567, 8001234567, 8881234567, 5125430191, 8887650098]

// only 1 for 800, 1 for 888
// [512, 800, 888], 5 => [8002345678, 8881234567]

// [*512*, 800, 888]; []
// [512, *800*, 888]; []
// [512, 800, *888*]; [800]

// [800, 888]; 512
// [800, 888, 512]

/*
 * Complete the 'GetPhoneNumbers' function below.
 *
 */
func GetPhoneNumbers(desiredCodes []string, desiredCount int) []string {

	toReturn := []string{}

	for len(toReturn) != desiredCount && len(desiredCodes) > 0 {
		current := desiredCodes[0]
		desiredCodes := desiredCodes[1:]

		number, found := GetAvailablePhoneNumber(current)

		if found {
			toReturn = append(toReturn, number)
			desiredCodes = append(desiredCodes, current)
		}

	}

	return toReturn
}

// 512 => 5121234567, true
// 512 => 5127654321, true
// ..
// 512 => '', false

// assume that you have at your disposal a function
func GetAvailablePhoneNumber(areaCode string) (string, bool) {
	// suppose this function returns
	// ("<an available phone number>"", true) if a phone number starting by areaCode is available
	// ("", false) otherwise
}
