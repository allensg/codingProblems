package main

/*
 * Server Traffic Monitor Hackerrank
 * Complete the 'getMaxTrafficTime' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY start
 *  2. INTEGER_ARRAY end
 */

func getMaxTrafficTime(start []int32, end []int32) int32 {
	// Write your code here

	endTime := end[len(end)-1]
	maxTraffic, maxTrafficIndex := 0, 0
	for currentSecond := 0; currentSecond < int(endTime); currentSecond++ {
		// maybe subtraction? like for each second, check each element in the array
		currentSecondActiveCount := 0
		// get start and end for each second
		for index, startValue := range start {
			endValue := end[index]

			// if endValue - currentSecond > its start value it counts
			//diff := endValue - int32(currentSecond)
			if currentSecond >= int(startValue) && int(endValue) >= currentSecond {
				currentSecondActiveCount = currentSecondActiveCount + 1
			}
		}
		//fmt.Printf("currentSecond: %d, activeCount: %d, maxTraffic: %d\n", currentSecond, currentSecondActiveCount, maxTraffic)
		// then its a max problem where you keep track of the most recent max and its index.
		if currentSecondActiveCount > maxTraffic {
			//  fmt.Println("inside active reassign")
			maxTraffic = currentSecondActiveCount
			maxTrafficIndex = currentSecond
		}

	}

	return int32(maxTrafficIndex)

}

/*
possible better solution by preloading values into a map but it passes less tests and i don't have time to work out the details.
  startMap := map[int]int{}
    endMap := map[int]int{}
    for index, startValue := range start {
        endValue := end[index]
        startMapValue, startHit := startMap[int(startValue)]
        endMapValue, endHit := endMap[int(endValue)]

        if startHit {
            startMap[int(startValue)] = startMapValue + 1
        } else {
            startMap[int(startValue)] = 1
        }


        if endHit {
            endMap[int(endValue)] = endMapValue + 1
        } else {
            endMap[int(endValue)] = 1
        }
    }


    // make a loop between 0 and the max end time
    endTime := end[len(end)-1]
    maxTraffic, maxTrafficIndex := 0,0
    rollingSum := 0
    // set up a rolling sum by checking the maps
    for currentSecond := 0; currentSecond < int(endTime); currentSecond++ {
        // get start and end for each second

        startMapValue, startHit := startMap[currentSecond]
        endMapValue, endHit := endMap[currentSecond]

        if startHit {
            rollingSum = rollingSum + startMapValue
        }

        if endHit {
            rollingSum = rollingSum - endMapValue
        }


        //fmt.Printf("currentSecond: %d, activeCount: %d, maxTraffic: %d\n", currentSecond, currentSecondActiveCount, maxTraffic)
        // then its a max problem where you keep track of the most recent max and its index.
        if rollingSum > maxTraffic {
          //  fmt.Println("inside active reassign")
            maxTraffic = rollingSum
            maxTrafficIndex = currentSecond
        }

    }

    return int32(maxTrafficIndex)
*/

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	startCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var start []int32

	for i := 0; i < int(startCount); i++ {
		startItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		startItem := int32(startItemTemp)
		start = append(start, startItem)
	}

	endCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var end []int32

	for i := 0; i < int(endCount); i++ {
		endItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		endItem := int32(endItemTemp)
		end = append(end, endItem)
	}

	result := getMaxTrafficTime(start, end)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

*/
