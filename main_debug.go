package main

import(
    "fmt"
    "strconv"
    "math/rand"
    "time"
)
//Число компьютера
type Computer struct{
    Nums []int
}
//Метод генерации случайного числа компьютера
func (c *Computer) gen(){
    rand.Seed(time.Now().UnixNano())
    for{
        var bol bool
        tmp := rand.Intn(10)
        if len(c.Nums) == 0 {
            c.Nums = append(c.Nums, tmp)
            continue
        }else if len(c.Nums) == 4{
            break
        }

        for _, v := range c.Nums {
            if v == tmp{
                bol = true
                break
            }
        }
        if bol{
            continue
        }
        c.Nums = append(c.Nums, tmp)
    }
}
//Метода проверки на коров и быков
func (c *Computer) check(userNumber []int) bool{
    var cows, bulls int
    for _, v := range userNumber{
        for _, v2 := range c.Nums{
            if v == v2{
                cows++
            }
        }
    }

    for k, v := range userNumber{
        if c.Nums[k] == v{
            bulls++
        }
    }

    fmt.Println("Коров:", cows, "Быков:", bulls)
    if bulls == 4 {
        return true
    }else{
        return false
    }
}

func (c *Computer) Read(userNumber []int) bool{
    for k, v := range userNumber{
        for k2, v2 := range userNumber{
            if v == v2 && k != k2{
                fmt.Println("Вы ввели некорректное число, циферки не должны совпадать")
                return false
            }
        }
    }
    return true
}

func main(){
    var comp Computer
    comp.gen()
    var input string
    var userNumber []int
    fmt.Println("Загаданное число:", comp.Nums)
    fmt.Println("Компьютер загадал число, отгадайте его")
    for pop := 0; pop < 100; pop++{
        //Запускаем бесконечный цикл, пока не будет введено корректное число
        //Как только вводиться корректная строка, которая конвертируется
        //число, и при этом длина данного числа равна 4
        //То мы берем эту строку и разбиваем на массив символов, и отдельно
        //вытаскиваем цифры и записываем их в массив чисел
        userNumber = make([]int, 0, 0)
        fmt.Print("Введите число: ")
        fmt.Scanln(&input)
        _, err := strconv.Atoi(input)
        if input != "" && err == nil && len([]byte(input)) == 4{
            //преобразуем введенные символы в слайс байтов
            //идем по этому слайсу
            //берем элемент преобразуем его в строку, а потом в число
            //и записываем в слайс ПользовательскогоЧисла
            number := []byte(input)
            for i := 0; i < len(number); i++{
                tmp := string(number[i])
                toIntTmp, _ := strconv.Atoi(tmp)
                userNumber = append(userNumber, toIntTmp)
            }
            fmt.Println()
            fmt.Println("----DEBUG-START----")
            fmt.Println("Пользователь ввел:", userNumber)
            fmt.Println("Загаданное число:", comp.Nums)
            fmt.Println("----DEBUG-END----")
            fmt.Println()
            if comp.Read(userNumber) && comp.check(userNumber) {
                fmt.Println("Ты выиграл")
                break
            }else{
                continue
            }
        }else{
            fmt.Println("Некорректный ввод")
        }
        if pop == 99{
            fmt.Println("Вы проиграли")
        }

    }
}
