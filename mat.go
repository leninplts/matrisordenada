/*
*
* Estudiantes:
*	Lenin Jose Mamani Huaylla
*	Jhon Javier Cutipa Vargas
*
 */
package main

import "fmt"
import "sort"

func sortn(nums []int) {
	sort.Ints(nums[:])
	fmt.Println(nums)
}

func main() {

	var twoD [4][4]int
	fmt.Println("ingrese 16 valores: ")
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD); j++ {
			fmt.Scan(&twoD[i][j])
		}
	}
	fmt.Println("matriz no ordenada: ")
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD); j++ {
			fmt.Print(twoD[i][j], " ")
		}
		fmt.Printf("\n")
	}

	var cont int
	var nums [16]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD); j++ {
			nums[cont] = twoD[i][j]
			cont++
		}
	}

	//fmt.Println(twoD)
	fmt.Println("numeros de la matriz no ordenada: ", nums)

	//la division de matrices en cuatro
	var part = len(twoD) / 2
	var p = part * part

	//primera parte
	parte1 := make([]int, p)

	cont = 0

	for i := 0; i < len(twoD)/2; i++ {
		for j := 0; j < len(twoD)/2; j++ {
			parte1[cont] = twoD[i][j]
			cont++
		}
	}

	//segunda parte
	parte2 := make([]int, p)

	cont = 0

	for i := len(twoD) / 2; i < len(twoD); i++ {
		for j := 0; j < len(twoD)/2; j++ {
			parte2[cont] = twoD[i][j]
			cont++
		}
	}

	//tercera parte
	parte3 := make([]int, p)

	cont = 0

	for i := 0; i < len(twoD)/2; i++ {
		for j := len(twoD) / 2; j < len(twoD); j++ {
			parte3[cont] = twoD[i][j]
			cont++
		}
	}

	//cuarta parte
	parte4 := make([]int, p)

	cont = 0

	for i := len(twoD) / 2; i < len(twoD); i++ {
		for j := len(twoD) / 2; j < len(twoD); j++ {
			parte4[cont] = twoD[i][j]
			cont++
		}
	}

	//fmt.Println(partes)

	go sortn(parte1[:])
	go sortn(parte2[:])
	go sortn(parte3[:])
	go sortn(parte4[:])

	var ordnums12 [8]int
	copy(ordnums12[:], parte1[:]) // [:] creates a slice that references the entire array
	copy(ordnums12[len(parte1):], parte2[:])

	var ordnums34 [8]int
	copy(ordnums34[:], parte3[:]) // [:] creates a slice that references the entire array
	copy(ordnums34[len(parte3):], parte4[:])

	var ordnums [16]int
	copy(ordnums[:], ordnums12[:]) // [:] creates a slice that references the entire array
	copy(ordnums[len(ordnums12):], ordnums34[:])
	sortn(ordnums[:])

	fmt.Println("parte 1: ", parte1)
	fmt.Println("parte 2: ", parte2)
	fmt.Println("parte 3: ", parte3)
	fmt.Println("parte 4: ", parte4)
	fmt.Println("todos los numeros ordenando: ", ordnums)

	cont = 0
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD); j++ {
			twoD[i][j] = ordnums[cont]
			cont++
		}
	}

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD); j++ {
			fmt.Print(twoD[i][j], " ")
		}
		fmt.Printf("\n")
	}

}
