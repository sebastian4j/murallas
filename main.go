package main

import (
	"log"
	"sort"
)

func buscar(input string, desde *[]int, hasta *[]int) *[]int {
	log.Println(input)
	data := make(map[int]int)
	resultado := []int{}
	chars := []rune(input)
	im := false
	cuenta := 0
	for a, b := range chars {
		t := string(b)
		if t == "|" {
			if !im { // inicio de muralla
				im = true
				data[a] = cuenta
			} else { // termino de muralla
				data[a] = cuenta
				cuenta = 0
			}
		} else if t == "*" && im {
			cuenta++
		}
	}
	log.Println("data")
	keys := []int{}
	for a, b := range data {
		log.Println(a, b)
		keys = append(keys, a)
	}
	log.Println("keys")
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, b := range keys {
		log.Println(b)
	}

	for i, j := range *desde {
		inxdesde := j
		indxhasta := (*hasta)[i]
		log.Println()
		log.Println("original", inxdesde, indxhasta)
		// obtener el inicio
		for _, v := range keys {
			if v == inxdesde {
				// el indice es correcto para buscar datos
				break
			} else if v > inxdesde {
				// el primer mayor será el valido
				inxdesde = v
				break
			}
		}
		log.Println("inicio encontrado", inxdesde)
		log.Println("hasta: ", indxhasta)
		// buscar las coincidencias
		contados := 0
		primero := true
		for _, v := range keys {
			if v >= inxdesde && v <= indxhasta {
				log.Println("indice lectura: ", v, data[v])
				if !primero {
					contados += data[v]
				} else {
					primero = false
				}
			}
		}
		log.Println(contados)
		resultado = append(resultado, contados)
	}
	return &resultado
}
