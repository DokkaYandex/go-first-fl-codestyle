package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

func defence(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(5, 10))
	} else if charClass == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(-2, 2))
	} else if charClass == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", charName, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"

}

func special(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	} else if charClass == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	} else if charClass == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	}
	return "неизвестный класс персонажа"
}

func StartTraining(ChaName, CharClass string) string {
	if CharClass == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", ChaName)
	}

	if CharClass == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", ChaName)
	}

	if CharClass == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", ChaName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(ChaName, CharClass))

		} else if cmd == "defence" {
			fmt.Println(defence(ChaName, CharClass))

		} else if cmd == "special" {
			fmt.Println(special(ChaName, CharClass))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func СhoiseСlass() string {
	var ApproveChoice string
	var CharClass string

	for ApproveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &CharClass)

		if CharClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")

		} else if CharClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")

		} else if CharClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &ApproveChoice)
		ApproveChoice = strings.ToLower(ApproveChoice)
	}
	return CharClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var CharName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &CharName)

	fmt.Printf("Здравствуй, %s\n", CharName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	CharClass := СhoiseСlass()

	fmt.Println(StartTraining(CharName, CharClass))
}
