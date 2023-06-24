package main

import (
	"fmt"
	"unicode/utf8"
)

type employee struct {
	salary  int
	country string
}

func main() {
	/*
			--------add/delete elements into map--------
			employeeSalary := make(map[string]int)
		    employeeSalary["steve"] = 12000

			employeeSalary := map[string]int {
		        "steve": 12000,
		        "jamie": 15000,
		    }
			delete(employeeSalary, "steve")

			-------accessing map--------
			salary_of_jamie := employeeSalary["jamie"]
			salary_of_not_existed_person_is_zero := employeeSalary{"Spongebob"}

			-------check existence------
			value, ok := map[key]
			if ok == true {
				fmt.Printf("key is present in map")
			}

			-------interate thru map------
			for key, value := range map {
		        fmt.Printf("map[%s] = %d\n", key, value)
		    }
	*/
	// --------map of structs------
	emp1 := employee{
		salary:  12000,
		country: "USA",
	}
	emp2 := employee{
		salary:  14000,
		country: "Canada",
	}
	emp3 := employee{
		salary:  13000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	}
	//----------concatnate string-------
	string1 := "Go"
	string2 := "is awesome"
	fmt.Printf(string1 + " " + string2 + "\n")
	result := fmt.Sprintf("%s %s\n", string1, string2) // format a string according to the input format specifier
	fmt.Printf(result)

	//---------rune---------
	// slice of rune: rune[]
	fmt.Printf("using rune for strings\n")
	name := "Señor"
	rune_represent_of_name := []rune(name)
	fmt.Printf("1111rune_represent_of_name is %x\n", rune_represent_of_name)
	// slice of rune to representt name variable:
	// some character can be represented by more than 1 byte of encoding, printf assume each code is 1 byte long
	// []rune(name) => converting name from string type to rune tuype
	// []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} <- unicode of character in hex
	charsAndBytePosition(name)
	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	// string function can take in (a slice of byte in hex  or decimal) or  (a slice of rune) and print it into the actual string
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(str))
	fmt.Printf("Number of bytes: %d\n", len(str)) // len and RuneCountInString function return different value
}

// -------------rune-------
// characters can be represented in several bytes in UTF-8 but rune can represent all characters with different bytes
func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}
