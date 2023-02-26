package product

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(index int) (*Product, error) {

	if index > len(allProducts) {
		return nil, errors.New("Out of range")

	}

	return &allProducts[index], nil
}

func (s *Service) Delete(index int) error {

	// a := []string{"A", "B", "C", "D", "E"}
	// i := 2
	// Удалить элемент по индексу i из a.
	// // 1. Копировать последний элемент в индекс i.
	// a[i] = a[len(a)-1]

	// // 2. Удалить последний элемент (записать нулевое значение).
	// a[len(a)-1] = ""

	// // 3. Усечь срез.
	// a = a[:len(a)-1]

	// if  len(allProducts) > index:
	// 	return 1

	if index > len(allProducts) {
		return errors.New("Out of range")

	}

	allProducts[index] = allProducts[len(allProducts)-1]
	allProducts[len(allProducts)-1] = Product{""}
	allProducts = allProducts[:len(allProducts)-1]

	return nil

}

func (s *Service) New(name string) error {
	allProducts = append(allProducts, Product{Title: name})
	return nil
}

func (s *Service) Edit(index int, name string) error {

	if index > len(allProducts) {
		return errors.New("Out of range")

	}

	allProducts[index] = Product{Title: name}
	return nil
}
