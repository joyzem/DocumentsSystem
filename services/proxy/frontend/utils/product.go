package utils

import (
	"fmt"

	productDto "github.com/joyzem/documents/services/product/dto"
	"github.com/levigross/grequests"
)

func GetProducts() (*productDto.GetProductsResponse, error) {
	productsUrl := fmt.Sprintf("%s/products", GetProductsAddress())
	productsResp, err := grequests.Get(productsUrl, nil)
	var products productDto.GetProductsResponse
	productsResp.JSON(&products)
	return &products, err
}

func GetProductById(id int) (*productDto.ProductByIdResponse, error) {
	productUrl := fmt.Sprintf("%s/products/%d", GetProductsAddress(), id)
	productResp, err := grequests.Get(productUrl, &grequests.RequestOptions{
		JSON: productDto.ProductByIdRequest{
			Id: id,
		}})
	var product productDto.ProductByIdResponse
	productResp.JSON(&product)
	return &product, err
}
