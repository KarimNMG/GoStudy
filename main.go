package main

import (
	"fmt"
	//"github.com/KarimNMG/GoStudy/controllers"
)

func startWebServer(port int, number int) (int, error) {
	fmt.Println("web server has started...", port)
	fmt.Println("Success", number)
	return port, nil
}
func main() {

	///---------------Panic functions
	
	
	
	//--------------LOOPs-------------------

	// mp1 := map[string]int {"http":70, "https":4010}
	// for k, v := range(mp1) {
	// 	println(k, v)
	// }

	// slice := []int{5, 6, 7}

	// for i, v := range slice {
	// 	println(i, v)
	// 	// very good for mapp
	// }

	// for i := 0; i < len(slice); i++ {
	// 	println(slice[i])
	// }

	// for i := 0; i < 5; i++ {
	// 	println("i", i)
	// }

	// var j int
	// for j < 5 {
	// 	println("j", j)
	// 	j++
	// 	if j == 3 {
	// 		break
	// 	}
	// }
	///-------------------controllers , constructors,
	// controllers.RegisterControllers()
	// http.ListenAndServe(":3000", nil)
	//-------------------functions----------------------
	// port := 2000
	// _, err := startWebServer(port, 55)
	// fmt.Println(err)
	//---------------------------------------
	// structs from other file
	// u := models.User{
	// 	Id:        1,
	// 	FirstName: "Karim",
	// 	LastName:  "Abdelatif",
	// }

	// fmt.Printf(u.FirstName)
	//fmt.Println("Hellow world")
	//--------------------------------------------------
	// var pi float32 = 3.14
	// fmt.Println(pi)

	// firstName := "Karim"
	// fmt.Println(firstName)

	// c := complex(3, 4)
	// fmt.Println(c)

	// r, im := real(c), imag(c)
	// fmt.Println(r, im)

	//-----------------------------------------------
	// pointers

	// var firstName *string = new(string)
	// *firstName = "Karim"

	// fmt.Println(*firstName)

	// lastName := "Abdelatif"

	// ptr := &lastName

	// fmt.Println(ptr, *ptr)

	// const

	// const pi = 3.14
	// fmt.Println(pi)
	//-------------------------------------------
	// collections
	// var arr [3]int
	// arr[0] = 5
	// arr[1] = 5
	// arr[2] = 8
	// fmt.Println(arr)

	// arrays and slice
	// arr := []int{5, 8, 1}
	// fmt.Println(arr)

	// arr = append(arr, 44, 55, 77)

	// fmt.Println(arr)

	// s1 := arr[1:]
	// s2 := arr[:2]
	// s3 := arr[1:2]
	// fmt.Println(s1, s2, s3)

	// emp := []int{}

	// emp = append(emp, 44, 55, 77)

	// fmt.Println(emp)

	// s4 := emp[1:]
	// s5 := emp[:2]
	// s6 := emp[1:2]

	// fmt.Println(s4, s5, s6)

	// fmt.Println(emp)

	// map
	// map1 := map[string]int{"k1": 5}

	// fmt.Println(map1)

	// map1["k1"] = 44
	// map1["k2"] = 488

	// fmt.Println(map1["k2"])
	// fmt.Println(map1)

	// delete(map1, "k1")
	// fmt.Println(map1)
	// fmt.Println(map1["k1"])

	//-----------------------------------------------
	// STRUCT

	// type user struct {
	// 	Id        int
	// 	FirstName string
	// 	LastName  string
	// }
	// var u user
	// u.Id = 1
	// u.FirstName = "Karim"
	// u.LastName = "Abdelatif"
	// fmt.Println(u)

	// u2 := user{
	// 	Id:        1,
	// 	FirstName: "Ahmed",
	// 	LastName:  "Ali",
	// }
	// fmt.Println(u2)
}
