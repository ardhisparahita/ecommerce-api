package mapper

import (
	"github.com/ardhisparahita/ecommerce-api/internal/domain"
	"github.com/ardhisparahita/ecommerce-api/internal/dto/response"
)

func ToCartResponse(cart *domain.Cart) *response.CartResponse {
	subtotal := cart.Product.Price * float64(cart.Quantity)

	return &response.CartResponse{
		ID:       cart.Product.ID,
		Quantity: cart.Quantity,
		Subtotal: subtotal,
		Product: response.CartProductResponse{
			ID:       cart.Product.ID,
			Name:     cart.Product.Name,
			Price:    cart.Product.Price,
			ImageURL: cart.Product.ImageURL,
		},
	}
}

func ToCartResponses(carts []domain.Cart) ([]response.CartResponse, float64) {
	var (
		responses  []response.CartResponse
		grandTotal float64
	)

	for _, cart := range carts {
		subtotal := cart.Product.Price * float64(cart.Quantity)

		responses = append(responses, response.CartResponse{
			ID:       cart.Product.ID,
			Quantity: cart.Quantity,
			Subtotal: subtotal,
			Product: response.CartProductResponse{
				ID:       cart.Product.ID,
				Name:     cart.Product.Name,
				Price:    cart.Product.Price,
				ImageURL: cart.Product.ImageURL,
			},
		})
	}

	return responses, grandTotal
}
