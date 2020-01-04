package main// ctrl +backspace delete word
import "fmt"

//shift +enter перевод строки
//main + Tag - live templates
func main() {

	score1:=10
	score2:=7
	score3:=10




	promoters :=0
	detractors :=0

	//refactoring - улучшение структуры кода без модификации поведения
	//ctrl alt v - позволять менять константы
	//shift fn f6

	promoutersLowerBound := 9
	detractorsUpperBound := 6

	if score1 >= promoutersLowerBound {
		promoters++
	}


	if score1 <= detractorsUpperBound {
		detractors++
	}
	if score2 >= promoutersLowerBound {
		promoters++
	}

	if score2 <= detractorsUpperBound {
		detractors++
	}
	if score3 >= promoutersLowerBound {
		promoters++
	}

	if score3 <= detractorsUpperBound {
		detractors++
	}
	nps:=(promoters-detractors)*100/3
	fmt.Println(nps)
}



