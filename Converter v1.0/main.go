package main

import "fmt"

func main() {
	var vhod string
	var usl int
	alph := make(map[string]int)
	alph["I"] = 1
	alph["II"] = 2
	alph["III"] = 3
	alph["IV"] = 4
	alph["V"] = 5
	alph["VI"] = 6
	alph["VII"] = 7
	alph["VIII"] = 8
	alph["IX"] = 9
	alph["X"] = 10
	alph["XX"] = 20
	alph["XXX"] = 30
	alph["XL"] = 40
	alph["L"] = 50
	alph["LX"] = 60
	alph["LXX"] = 70
	alph["LXXX"] = 80
	alph["XC"] = 90
	alph["C"] = 100
	alph["CC"] = 200
	alph["CCC"] = 300
	alph["CD"] = 400
	alph["D"] = 500
	alph["DC"] = 600
	alph["DCC"] = 700
	alph["DCCC"] = 800
	alph["CM"] = 900
	alph["M"] = 1000
	chislo := 0
	for i := true; i == true; {
		fmt.Print("Введите цифру от одного до трех: \n 1.Ввести римскую цифру \n 2.Получить результат \n 3.Выход")
		fmt.Scanf("%d\n", &usl)
		switch usl {
		case 1:
			{
				fmt.Print("Введите цифру")
				fmt.Scanf("%s\n", &vhod)
				chislo += alph[vhod]
			}
		case 2:
			{
				fmt.Println("Результат:", chislo)
			}
		case 3:
			{
				i = false
			}
		}
	}
}
