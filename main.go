package main// ctrl +backspace delete word
import "fmt"

//shift +enter перевод строки
//main + Tag - live templates
func main() {
	score1:=10
	score2:=10
	score3:=10

	promoters :=0
	detractors :=0

	if score1 >= 9{
		promoters++
	}

	if score1 <= 6{
		detractors++
	}
	if score2 >= 9{
		promoters++
	}

	if score2 <= 6{
		detractors++
	}
	if score3 >= 9{
		promoters++
	}

	if score3 <= 6{
		detractors++
	}
	nps:=(promoters-detractors)*100/3
	fmt.Println(nps)
}



