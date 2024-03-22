package main

import "fmt"

func main() {

	fmt.Println("Create custom type - struct")

	type person struct {
		name          string
		second_name   string
		fav_ice_cream []string
	}

	fifonz := person{
		name:          "Fifonż",
		second_name:   "Nowakowski",
		fav_ice_cream: []string{"chocolade", "banana"},
	}

	alfons := person{
		name:          "Alfons",
		second_name:   "Nowak",
		fav_ice_cream: []string{"wanilla", "banana"},
	}

	fmt.Printf("%v\n", fifonz)
	fmt.Printf("%v\n", alfons)

	for _, k := range alfons.fav_ice_cream {
		fmt.Printf("%v\n", k)
	}

	for _, k := range fifonz.fav_ice_cream {
		fmt.Printf("%v\n", k)
	}

	fmt.Println("=================")

	fmt.Println("Compose struct into map")
	person_map := make(map[string]person)

	person_map[alfons.second_name] = alfons
	person_map[fifonz.second_name] = fifonz

	//or
	// person := map[string]person{
	// 	fifonz.second_name: fifonz,
	// 	alfons.second_name: alfons,
	// }

	for k, v := range person_map {
		fmt.Printf("%v %v\n", k, v)
	}

	fmt.Println("=================")

	fmt.Println("Compose struct into other struct")

	type Engine struct {
		electric bool
	}

	type Car struct {
		engine Engine
		make   string
		model  string
		doors  int
		color  string
	}

	fiat_coupe := Car{
		engine: Engine{
			electric: false,
		},
		make:  "Fiat",
		model: "Coupe",
		doors: 2,
		color: "Red",
	}
	dodge_chall := Car{
		engine: Engine{
			electric: false,
		},
		make:  "Dodge",
		model: "Challenger",
		doors: 2,
		color: "Grey",
	}

	fmt.Println(fiat_coupe)
	fmt.Println(dodge_chall)

	fmt.Println(fiat_coupe.make, fiat_coupe.model)
	fmt.Println(dodge_chall.make, dodge_chall.model)

	fmt.Println("=================")

	fmt.Println("Anonymus structs")

	me := struct{
		first string
		friends map[string]int
		fav_drinks []string
	}{
		first: "Matthew",
		friends: map[string]int{
			"Kubuś": 22,
			"Buła": 22,
			"Gaba": 22,
		},
		fav_drinks: []string{"water", "pepsi"},
	}

	fmt.Print(me)

	for _, v := range me.friends{
		fmt.Print(me.first, " - firends - ", v, "\n")
	}
	for _, v := range me.fav_drinks{
		fmt.Print(me.first, " - fav drinks - ", v, "\n")
	}
}
