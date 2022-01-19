package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Hex(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	// Find each instance of(hex) and turning the prior word into dec num
	for i := 0; i < len(n1); i++ {
		if n1[i] == "(hex)" {
			s, _ := strconv.ParseInt(n1[i-1], 16, 64)
			x := strconv.FormatInt(s, 10)
			n1[i-1] = x
			n1[i] = ""
		}
	}

	// Removing all empty strings
	finalSlice := []string{}
	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			finalSlice = append(finalSlice, n1[i])
		}
	}
	final := strings.Join(finalSlice, " ")
	result := []byte(final)
	return result
}

//
func Bin(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)
	// Finding each instance of (bin) and changing the prior word to a dec num
	for i := 0; i < len(n1); i++ {
		if n1[i] == "(bin)" {
			num := n1[i-1]
			intNum, _ := strconv.Atoi(num)
			emptySlice := []int{}

			for intNum != 0 {
				emptySlice = append(emptySlice, intNum%10)
				intNum /= 10
			}

			x := 0
			for i := 0; i < len(emptySlice); i++ {
				x = x + int(math.Pow(float64(emptySlice[i]*2), float64(i)))
			}
			newNum := strconv.Itoa(x)
			n1[i-1] = newNum
			n1[i] = ""
		}
	}
	finalSlice := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			finalSlice = append(finalSlice, n1[i])
		}
	}
	final := strings.Join(finalSlice, " ")
	return []byte(final)
}

func Up(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	for i := 0; i < len(n1); i++ {
		if n1[i] == "(up)" {
			n1[i] = ""
			n1[i-1] = strings.ToUpper(n1[i-1])
		}
	}
	finalSlice := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			finalSlice = append(finalSlice, n1[i])
		}
	}
	final := strings.Join(finalSlice, " ")
	return []byte(final)
}

func Low(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	for i := 0; i < len(n1); i++ {
		if n1[i] == "(low)" {
			n1[i] = ""
			n1[i-1] = strings.ToLower(n1[i-1])
		}
	}
	finalSlice := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			finalSlice = append(finalSlice, n1[i])
		}
	}
	final := strings.Join(finalSlice, " ")
	return []byte(final)
}

func Cap(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	for i := 0; i < len(n1); i++ {
		if n1[i] == "(cap)" {
			n1[i] = ""
			n1[i-1] = strings.Title(n1[i-1])
		}
	}
	finalSlice := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			finalSlice = append(finalSlice, n1[i])
		}
	}
	final := strings.Join(finalSlice, " ")
	return []byte(final)
}

func MultiUp(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	for i := 0; i < len(n1); i++ {
		if strings.Contains(n1[i], "(up,") {
			reg, _ := regexp.Compile("[^0-9]+")
			sNum := reg.ReplaceAllString(string(n1[i+1]), "")
			num, _ := strconv.Atoi(sNum)

			for j := num; j > 0; j-- {
				n1[i-j] = strings.ToUpper(n1[i-j])
			}
			n1[i] = ""
			n1[i+1] = ""

		}
	}
	emptyStr := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			emptyStr = append(emptyStr, n1[i])
		}
	}

	final := strings.Join(emptyStr, " ")

	return []byte(final)
}

func MultiLow(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	for i := 0; i < len(n1); i++ {
		if strings.Contains(n1[i], "(low,") {
			reg, _ := regexp.Compile("[^0-9]+")
			sNum := reg.ReplaceAllString(string(n1[i+1]), "")
			num, _ := strconv.Atoi(sNum)

			for j := num; j > 0; j-- {
				n1[i-j] = strings.ToLower(n1[i-j])
			}
			n1[i] = ""
			n1[i+1] = ""

		}
	}
	emptyStr := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			emptyStr = append(emptyStr, n1[i])
		}
	}
	final := strings.Join(emptyStr, " ")

	return []byte(final)
}

func MultiCap(name []byte) []byte {
	n := string(name)
	n1 := strings.Fields(n)

	for i := 0; i < len(n1); i++ {
		if strings.Contains(n1[i], "(cap,") {
			reg, _ := regexp.Compile("[^0-9]+")
			sNum := reg.ReplaceAllString(string(n1[i+1]), "")
			num, _ := strconv.Atoi(sNum)

			for j := num; j > 0; j-- {
				n1[i-j] = strings.Title(n1[i-j])
			}
			n1[i] = ""
			n1[i+1] = ""

		}
	}
	emptyStr := []string{}

	for i := 0; i < len(n1); i++ {
		if n1[i] != "" {
			emptyStr = append(emptyStr, n1[i])
		}
	}

	final := strings.Join(emptyStr, " ")

	return []byte(final)
}

func Punc(name []byte) []byte {
	str := string(name)
	r := []rune(str)
	var s string

	for i := len(r) - 1; i >= 0; i-- {
		if i == 0 || i == 1 {
			s += string(r[i])
		} else if (r[i] == '.' || r[i] == ',' || r[i] == '?' || r[i] == '!' || r[i] == ':' || r[i] == ';') && (r[i-1] == '.' || r[i-1] == ',' || r[i-1] == '?' || r[i-1] == '!' || r[i-1] == ':' || r[i-1] == ';') {
			// If there are two puncs in a row
			s += string(r[i])
		} else if (r[i] == '.' || r[i] == ',' || r[i] == '?' || r[i] == '!' || r[i] == ':' || r[i] == ';') && r[i-1] == ' ' {
			// If there is a punc with a space before
			s += string(r[i])
			i = i - 1
		} else if r[i] == ' ' && (r[i+1] == ' ' || r[i-1] == ' ') {
			// Dont add double spaces
			continue
		} else if !(r[i] == '.' || r[i] == ',' || r[i] == '?' || r[i] == '!' || r[i] == ':' || r[i] == ';' || r[i] == ' ') && (r[i-1] == '.' || r[i-1] == ',' || r[i-1] == '?' || r[i-1] == '!' || r[i-1] == ':' || r[i-1] == ';') {
			// Make sure there is a space after punc
			s += string(r[i])
			s += " "
		} else {
			s += string(r[i])
		}
	}

	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return []byte(string(rev))
}

func Marks(name []byte) []byte {
	str := string(name)
	r := []rune(str)
	open := false
	newR := []rune{}

	for i := 0; i < len(r); i++ {
		if i == 73 {
			newR = append(newR, r[i])
			continue
		}

		if !open && r[i] == 39 {
			open = true
			newR = append(newR, r[i])
			if r[i+1] == ' ' {
				i = i + 1
			}
		} else if open && r[i] == 39 {
			open = false
			newR = append(newR, r[i])
		} else if open && r[i] == ' ' && r[i+1] == 39 {
			continue
		} else {
			newR = append(newR, r[i])
		}
	}

	return []byte(string(newR))
}

func Vowel(name []byte) []byte {
	str := string(name)
	r := strings.Fields(str)

	for i := 0; i < len(r); i++ {
		if r[i] == "a" {
			for j, letter := range r[i+1] {
				if j == 0 && (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h' || letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H') {
					r[i] = "an"
				}
			}
		} else if r[i] == "A" {
			for k, letter := range r[i+1] {
				if k == 0 && (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h' || letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H') {
					r[i] = "An"
				}
			}
		}
	}
	r1 := strings.Join(r, " ")
	return []byte(r1)
}

func main() {
	// Take file name from arguement
	name, _ := ioutil.ReadFile(os.Args[1])

	// Change all words with (hex) to the decimal version. (cast to string and use parseint)
	name = Hex(name)
	// Change all instances of (bin) to decimal
	name = Bin(name)
	// Change all instances of (up) to uppercase
	name = Up(name)

	// Change all instance of (low) to Lowercase
	name = Low(name)
	// Change all instances of (cap) to Title
	name = Cap(name)
	//
	name = MultiUp(name)
	name = MultiLow(name)
	name = MultiCap(name)

	name = Vowel(name)

	name = Punc(name)

	name = Marks(name)

	// Return result file
	err := ioutil.WriteFile(os.Args[2], name, 0o644)
	if err != nil {
		fmt.Print(err)
	}
}
