package product

import (
	domain "alirah/app/domain/product"
)

func SingleResource(product *domain.Product) *domain.Product {
	return &domain.Product{
		Id:          product.Id,
		Title:       product.Title,
		Description: product.Description,
		Image:       product.Image,
		Price:       product.Price,
	}
}

func Collection(products *[]domain.Product) []*domain.Product {
	resources := make([]*domain.Product, 0)
	for _, value := range *products {
		resource := SingleResource(&value)
		resources = append(resources, resource)
	}
	return resources
}
