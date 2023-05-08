package book

import "strconv"

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(id int, bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Book, error) {
	return s.repository.FindByID(id)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
	}
	return s.repository.Create(book)
}

func (s *service) Update(id int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(id)

	price, _ := bookRequest.Price.Int64()
	rating, _ := strconv.Atoi(bookRequest.Rating)

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = rating
	return s.repository.Update(book)
}
