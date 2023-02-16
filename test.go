package main

import(
	"fmt"
	"flag"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var(
		array []int
		maxLen []int 
		resultToo []int
		IDs []string
	)

	inputA := flag.Int("i", 255, "Network address")
	inputB := flag.Int("b", 255, "Network address")
	boolFlag := flag.Bool("+", false, "Summa")
	flag.Parse()
	A := binary(*inputA, array)
	B := binary(*inputB, array)
	maxLen = append(maxLen, len(A), len(B))
	sort.Ints(maxLen)
	if *boolFlag {
		summa(A, B, resultToo, maxLen, IDs)
	} else {
		check := false
		if len(A) < maxLen[len(maxLen)-1] {
			A = viravnivanie(A, maxLen)
			B = reverse(B)

		} else {
			B = viravnivanie(B, maxLen)
			A = reverse(A)
		}
		if *inputA < *inputB {
			A, B = B, A
		} else {
			A, B = A, B
		}
		A = reverse(A)
		B = reverse(B)

		for i := 0; i < maxLen[len(maxLen)-1]; i++ {
			result := (A[i] - B[i])
			if result == 0 && check == false && A[i] == 0 {
				resultToo = append(resultToo, 0)
				check = false
			} else if result == 1 && check == false && A[i] == 1 {
				resultToo = append(resultToo, 1)
				check = false
			}  else if result == 0 && check == false && A[i] == 1 {
				resultToo = append(resultToo, 0)
				check = false
			} else if result == -1 && check == false && A[i] == 0 {
				resultToo = append(resultToo, 1)
				check = true
			} else if result == 0 && check == true && A[i] == 0 {
				resultToo = append(resultToo, 1)
				check = true
			} else if result == 1 && check == true && A[i] == 1 {
				resultToo = append(resultToo, 0 )
				check = false
			} else if result == 0 && check == true && A[i] == 1 {
				resultToo = append(resultToo, 1 )
				check = true
			} else if result == -1 && check == true && A[i] == 0 {
				resultToo = append(resultToo, 0 )
				check = true
			} 
		}
	
		if *inputA < *inputB {
			resStr := strConv(reverse(resultToo), IDs)
			fmt.Println("-", resStr)
		} else {
			resStr := strConv(reverse(resultToo), IDs)
			fmt.Println(resStr)
		}
	}	
}

func summa(valueA []int, valueB []int, valueRes []int, valueLen [] int, valueStr []string) {
	check := false
	if len(valueA) < valueLen[len(valueLen)-1] {
		valueA = viravnivanie(valueA, valueLen)
		valueB = reverse(valueB)
	} else {
		valueB = viravnivanie(valueB, valueLen)
		valueA = reverse(valueA)
	}

	valueA = reverse(valueA)
	valueB = reverse(valueB)

	for i:= 0; i < valueLen[len(valueLen)-1]; i++ {
		result := valueA[i] + valueB[i]
		if result == 0 && check == false {
			valueRes = append(valueRes, 0)
		} else if result == 1 && check == false  { 
			valueRes = append(valueRes, 1)
		} else if result == 2 && check == false  {
			valueRes = append(valueRes, 0)
			check = true
		} else if result == 0 && check == true { 
			valueRes = append(valueRes, 1)
			check = false
		} else if result == 1 && check == true { 
			valueRes = append(valueRes, 0)
			check = true
		} else if result == 2 && check == true { 
			valueRes = append(valueRes, 1)
			check = true
		}

	}
	if check == true {
		valueRes = append(valueRes, 1)
		check = false
	}

	resStr := strConv(reverse(valueRes), valueStr)
	fmt.Println(resStr)

}

func strConv(value []int, value1 []string) (output string) { 
	for _, i := range value {
		value1 = append(value1, strconv.Itoa(i))
	}
	/*fmt.Println(strings.Join(value1, ""))*/
	return strings.Join(value1, "")
}

func viravnivanie(value []int, maxLen []int) (output []int) {
	value = reverse(value)
	per := maxLen[len(maxLen)-1] - len(value)
	for i := 0; i < per; i++ {
		value = append([]int{0}, value...)
	}
	return value
}

func binary(value int, array []int) (outputBinaryArray []int) {
	for  value >= 1 {
        array = append(array, value % 2)
        value = value / 2
    }
    return array /*reverse(array)*/ // ----ЗАБЫЛ ЗАЧЕМ НАДО ВЕРНУТЬСЯ ВОЗМОЖНО СОКРАТИТЬ
}

func reverse(array []int) (outputArray []int) {
	for i:= 1; i <= len(array); i++ {
    	outputArray = append(outputArray, array[len(array) - i])
    }
    return outputArray
}
