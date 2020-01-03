package main

import refinement "github.com/mazuninky/refinement"

func main()  {
	numberType := refinement.MustNewRegexType(`[0-9]+`)
	numberPack := numberType.Pack("45")
	number, err := numberPack.Unpack()
	if err != nil {
		panic(err)
	}
	println(number)
}