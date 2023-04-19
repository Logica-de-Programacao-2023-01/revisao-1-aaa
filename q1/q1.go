package q1

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	
	discount := 0
	soma := 0
	valor := currentPurchase
	
	//conseguir o total do histórico de compras
	for i := range purchaseHistory{
		soma += purchaseHistory[i]
	}

	//comparar em ordem crescente para o desconto maior ser
	
	if soma <= 0.0{
		
		return errors.New("Valor de compra insuficiente.")
	}
	if soma <= 500.0{
		discount = 0.02 * valor
	}
	if soma <= 1000.0{
		discount = 0.5 * valor
	}
	if soma == nil{
		discount = 0.1 * valor
	}
	if soma > 1000.0{
		discount = 0.1 * valor
	}
	if (soma / len) > 1000.0{
		discount = 0.2 * valor
	}

	return desconto, nil
	fmt.Print("Valor do desconto:")
	fmt.Print("Erro:", err)
	return 0, nil
}
