package main

import (
	"encoding/json"
	"sort"
)

/*
 * Json Diff tool hackerrank
 * Complete the 'getJSONDiff' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING json1
 *  2. STRING json2
 */

func getJSONDiff(json1 string, json2 string) []string {
	var toReturn []string

	// figure out which list is longer
	input1 := map[string]string{}
	json.Unmarshal([]byte(json1), &input1)

	input2 := map[string]string{}
	json.Unmarshal([]byte(json2), &input2)

	// make sure the longer map is the first input
	if len(input1) > len(input2) {
		toReturn = compare(input1, input2)
	} else {
		toReturn = compare(input2, input1)
	}

	sort.Strings(toReturn)
	return toReturn
}

func compare(map1 map[string]string, map2 map[string]string) []string {
	toReturn := []string{}

	for key, map1Val := range map1 {

		map2Val, map2Hit := map2[key]

		// If a key is present in only one json, it should not be considered
		if !map2Hit {
			continue
		}

		if map1Val != map2Val {
			toReturn = append(toReturn, key)
		}

	}

	return toReturn
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	json1 := readLine(reader)

	json2 := readLine(reader)

	result := getJSONDiff(json1, json2)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
