package main

import "fmt"

/*
{
    "l": 8,
    "w": 4,
    "e": [6, 7],
}

{
    "c": 0,
    "w": 6,
    "e": [2, 7, 8],
}

[0, 2, 7, 8, 6, 2, 7, 8]
[0, 2, 6, 7]


("a...a", "a...a") 100 букв

n = len(res), m=len(in)
n * (2m) +  n = n(1+2m) = n*m
*/

func Fuzzysearch(res string, in string) bool {
	if (res == "" && in == "") || res == "" {
		return true
	}

	if in == "" {
		return false
	}

	inArr := []rune(in)

	outMap := make(map[rune][]int, len(inArr))

	for ri, r := range inArr {
		for _, s := range res {
			if s == r {
				_, ok := outMap[s]
				if !ok {
					outMap[s] = make([]int, 0)
				}

				outMap[s] = append(outMap[s], ri)
			}
		}
	}

	// outMap - map with key which is rune from left side and value - array of rune positions from 2 arg
	arrSorted := make([]int, 0)
	for _, r := range res {
		outr, ok := outMap[r]
		if ok {
			var curMin int = -1

			if len(arrSorted) > 0 {
				curMin = arrSorted[len(arrSorted)-1]
			}

			mino := getMin(outr, curMin)
			if mino == curMin {
				return false
			}

			arrSorted = append(arrSorted, mino)
		}
	}

	return checkSortedAscending(arrSorted)
}

func checkSortedAscending(slice []int) bool {
	if len(slice) < 1 {
		return false
	}

	curValue := slice[0]

	for _, v := range slice {
		if v < curValue {
			return false
		}
	}

	return true
}

func getMin(slice []int, curMin int) int {
	for _, s := range slice {
		if s > curMin {
			return s
		}
	}

	return curMin
}

func main() {
	fmt.Println(Fuzzysearch("car", "cartwheel"))       // true
	fmt.Println(Fuzzysearch("cwhl", "cartwheel"))      // true
	fmt.Println(Fuzzysearch("cwheel", "cartwheel"))    // true
	fmt.Println(Fuzzysearch("cartwheel", "cartwheel")) // true
	fmt.Println(Fuzzysearch("cwheeel", "cartwheel"))   // false
	fmt.Println(Fuzzysearch("lw", "cartwheel"))        // false
	fmt.Println(Fuzzysearch("wl", "cartwheel"))        // true
	fmt.Println(Fuzzysearch("cewe", "caertwheel"))     // true
	fmt.Println(Fuzzysearch("ceew", "caertwheel"))     // false
}
