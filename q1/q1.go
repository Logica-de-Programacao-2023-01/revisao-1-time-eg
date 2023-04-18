package q1

import "errors"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, errors.New("Valor da compra inválido")
	}

	var totalCompras float64

	for _, valorCompra := range purchaseHistory {
		totalCompras += valorCompra
	}

	if len(purchaseHistory) == 0 {
		return currentPurchase * 0.1, nil
	}

	if totalCompras > 1000 {
		return currentPurchase * 0.1, nil
	}

	if totalCompras <= 1000 {
		return currentPurchase * 0.05, nil
	}

	if totalCompras <= 500 {
		return currentPurchase * 0.02, nil
	}

	mediaCompras := totalCompras / float64(len(purchaseHistory))

	if mediaCompras > 1000 {
		return currentPurchase * 0.2, nil
	}
	return 0, nil
}
