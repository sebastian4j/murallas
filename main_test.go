package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMurallas(t *testing.T) {
	desde := []int{2, 0, 0, 3}
	hasta := []int{4, 5, 6, 5}
	res1 := buscar("*|**|*|", &desde, &hasta)
	esperado := []int{0, 2, 3, 0}
	assert.ElementsMatch(t, *res1, esperado)
	//
	desde = []int{0, 3, 1, 6}
	hasta = []int{10, 9, 7, 10}
	res1 = buscar("**|*|***|**", &desde, &hasta)
	esperado = []int{4, 3, 1, 0}
	assert.ElementsMatch(t, *res1, esperado)
	//
	desde = []int{0, 2, 5, 5, 0}
	hasta = []int{10, 6, 7, 10, 7}
	res1 = buscar("|****|*|**|", &desde, &hasta)
	esperado = []int{7, 0, 1, 3, 5}
	assert.ElementsMatch(t, *res1, esperado)
}
