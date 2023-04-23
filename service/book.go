package service

// usecase

import (
	"challenge-8/model"
	"errors"
)

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	if len(in.Author) < 5 {
		return res, errors.New("invalid name author length")
	}
	// call repo
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateBook(in)
}

func (s *Service) GetAllBook() (res []model.Book, err error) {
	return s.repo.GetAllBook()
}

func (s *Service) GetBookByID(id int64) (res model.Book, err error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
	if len(in.Author) < 5 {
		return res, errors.New("invalid name author length")
	}
	return s.repo.UpdateBook(in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}

// func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
// 	// TODO: implement
// }
