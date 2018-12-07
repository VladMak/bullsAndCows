package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Число компьютера
type Computer struct {
	Nums []int
}

//Метод генерации случайного числа компьютера
func (c *Computer) gen() {
	rand.Seed(time.Now().UnixNano())
	for {
		var bol bool
		tmp := rand.Intn(10)
		if len(c.Nums) == 0 {
			c.Nums = append(c.Nums, tmp)
			continue
		} else if len(c.Nums) == 4 {
			break
		}

		for _, v := range c.Nums {
			if v == tmp {
				bol = true
				break
			}
		}
		if bol {
			continue
		}
		c.Nums = append(c.Nums, tmp)
	}
}

//Метода проверки на коров и быков
func (c *Computer) check(userNumber []int) bool {
	var cows, bulls int
	for _, v := range userNumber {
		for _, v2 := range c.Nums {
			if v == v2 {
				cows++
			}
		}
	}

	for k, v := range userNumber {
		if c.Nums[k] == v {
			bulls++
		}
	}

	fmt.Println("Коров:", cows, "Быков:", bulls)
	if bulls == 4 {
		return true
	} else {
		return false
	}
}

func (c *Computer) Read(userNumber []int) bool {
	for k, v := range userNumber {
		for k2, v2 := range userNumber {
			if v == v2 && k != k2 {
				fmt.Println("Вы ввели некорректное число, циферки не должны совпадать")
				return false
			}
		}
	}
	return true
}

func manyConvertOperation(number []byte, userNumber []int) []int {
	for i := 0; i < len(number); i++ {
		tmp := string(number[i])
		toIntTmp, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Println("Вы ввели недопустимые символы")
			continue
		} else {
			userNumber = append(userNumber, toIntTmp)
		}
	}
	return userNumber
}

func debugInfo(userNumber, nums []int) {
	fmt.Println()
	fmt.Println("----DEBUG-START----")
	fmt.Println("Пользователь ввел:", userNumber)
	fmt.Println("Загаданное число:", nums)
	fmt.Println("----DEBUG-END----")
	fmt.Println()

}

func doUserWin(comp Computer, userNumber []int) bool {
	if comp.Read(userNumber) && comp.check(userNumber) {
		fmt.Println("Ты выиграл")
		return true
	} else {
		return false
	}
}

func main() {
	var comp Computer
	comp.gen()
	var input string
	var userNumber []int
	fmt.Println("Загаданное число:", comp.Nums)
	fmt.Println("Компьютер загадал число, отгадайте его")
	for pop := 0; pop < 100; pop++ {
		//Запускаем бесконечный цикл, пока не будет введено корректное число
		//Как только вводиться корректная строка, которая конвертируется
		//число, и при этом длина данного числа равна 4
		//То мы берем эту строку и разбиваем на массив символов, и отдельно
		//вытаскиваем цифры и записываем их в массив чисел
		userNumber = make([]int, 0, 0)
		input = ""
		fmt.Print("Введите число: ")
		fmt.Scanln(&input)
		_, err := strconv.Atoi(input)
		//Проверка на ошибочные входные данные
		if input != "" && err == nil && len([]byte(input)) == 4 {
			//преобразуем введенные символы в слайс байтов
			//идем по этому слайсу
			//берем элемент преобразуем его в строку, а потом в число
			//и записываем в слайс ПользовательскогоЧисла
			number := []byte(input)
			//Конвертируем байты в числа
			userNumber = manyConvertOperation(number, userNumber)
			debugInfo(userNumber, comp.Nums)
			if doUserWin(comp, userNumber) {
				break
			} else {
				continue
			}
			//если ошибка
		} else {
			if input == "" {
				fmt.Println("Вы ввели пустое сообщение, необходимо ввести четыре цифры")
				continue
			} else if len([]byte(input)) != 4 {
				fmt.Println("Вы ввели больше символов, чем необходимо, требуется четыре")
				continue
			}
			fmt.Println("Вы ввели недопустимые символы, должны быть введены только цифры")
		}
		if pop == 99 {
			fmt.Println("Вы проиграли")
		}

	}
}
