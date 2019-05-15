package main

import ("fmt")

//Map
func main() {
	/*x:=make(map[string] string)
	x["TH"]="Thailand"
	x["JP"]="Japan"
	x["EN"]="England"
	x["H"]="Hydrogen"
	x["Li"]="Lithium"*/
	x:=map[string] string{
		"TH":"Thailand",
		"JP":"Japan",
	}

	//fmt.Println(x["EN"])
	//fmt.Println(x["Li"])
	fmt.Println(x)


}