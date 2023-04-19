package q2

func AverageLettersPerWord(text string) (float64, error) {
	//função para contar a média de letras por palavra em um texto.  
	//receber um texto e retornar a média de letras por palavra. A função deve retornar um erro caso o texto não possua nenhuma palavra > "texto vazio".
	letras := 0
	espacos := 0
	
	if text == nil or " "{
		return 0, fmt.Errorf("texto vazio")
	}
	
	for i := range text{
		if i != " "{
			letras += 1
		}
		else{
			espacos += 1
		}
	}
	
	media := letras / (espacos + 1)
	fmt.Print("Média de letras por palavra:", media \n"Erro: nil")
	return media, nil
}
