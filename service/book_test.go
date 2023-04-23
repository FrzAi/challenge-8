package service

import (
	"challenge-8/model"
	"challenge-8/repository/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookByID(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{
				ID:        1,
				Name_book: "Golang",
				Author:    "fariz",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:        1,
			Name_book: "Golang",
			Author:    "fariz",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			BookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(BookRepo)
			}

			service := Service{
				repo: BookRepo,
			}

			res, err := service.GetBookByID(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

// func Test_BookService_GetAllBook(t *testing.T) {
// 	type testCase struct {
// 		name           string
// 		wantError      bool
// 		expectedResult model.Book
// 		expectedError  error
// 		onBookRepo     func(mock *mocks.MockBookRepo)
// 	}

// 	var testTable []testCase

// 	testTable = append(testTable, testCase{
// 		name:      "success",
// 		wantError: false,
// 		onBookRepo: func(mock *mocks.MockBookRepo) {
// 			mock.EXPECT().GetAllBook().Return(model.Book{
// 				ID:        1,
// 				Name_book: "Golang",
// 				Author:    "fariz",
// 			}, nil).Times(1)
// 		},
// 		expectedResult: model.Book{
// 			ID:        1,
// 			Name_book: "Golang",
// 			Author:    "fariz",
// 		},
// 	})

// 	// testTable = append(testTable, testCase{
// 	// 	name:          "record not found",
// 	// 	wantError:     true,
// 	// 	expectedError: errors.New("record not found"),
// 	// 	onBookRepo: func(mock *mocks.MockBookRepo) {
// 	// 		mock.EXPECT().GetAllBook().Return(model.Book{}, errors.New("record not found")).Times(1)
// 	// 	},
// 	// })

// 	// testTable = append(testTable, testCase{
// 	// 	name:          "unexpected error",
// 	// 	wantError:     true,
// 	// 	expectedError: errors.New("unexpected error"),
// 	// 	onBookRepo: func(mock *mocks.MockBookRepo) {
// 	// 		mock.EXPECT().GetAllBook().Return(model.Book{}, errors.New("unexpected error")).Times(1)
// 	// 	},
// 	// })

// 	for _, testCase := range testTable {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)

// 			BookRepo := mocks.NewMockBookRepo(mockCtrl)

// 			if testCase.onBookRepo != nil {
// 				testCase.onBookRepo(BookRepo)
// 			}

// 			service := Service{
// 				repo: BookRepo,
// 			}

// 			res, err := service.GetAllBook()

// 			if testCase.wantError {
// 				assert.EqualError(t, err, testCase.expectedError.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, testCase.expectedResult, res)
// 			}
// 		})
// 	}
// }

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			Name_book: "Golang",
			Author:    "fariz",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:        1,
				Name_book: "Golang",
				Author:    "fariz",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:        1,
			Name_book: "Golang",
			Author:    "fariz",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			Name_book: "Golang",
			Author:    "fariz",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid Name Author length",
		wantError: true,
		input: model.Book{
			Name_book: "Golang",
			Author:    "ai", // case negative
		},
		expectedError: errors.New("invalid name author length"),
	})

	// testTable = append(testTable, testCase{
	// 	name:      "invalid na",
	// 	wantError: true,
	// 	input: model.Book{
	// 		Name_book: "Golang",
	// 		Author:    "fariz", // case negative
	// 	},
	// 	expectedError: errors.New("invalid division"),
	// })

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			BookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(BookRepo)
			}

			service := Service{
				repo: BookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			Name_book: "Golang",
			Author:    "fariz",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{
				ID:        1,
				Name_book: "Golang",
				Author:    "fariz",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:        1,
			Name_book: "Golang",
			Author:    "fariz",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			Name_book: "Golang",
			Author:    "fariz",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid Name Author length",
		wantError: true,
		input: model.Book{
			Name_book: "Golang",
			Author:    "ai", // case negative
		},
		expectedError: errors.New("invalid name author length"),
	})

	// testTable = append(testTable, testCase{
	// 	name:      "invalid na",
	// 	wantError: true,
	// 	input: model.Book{
	// 		Name_book: "Golang",
	// 		Author:    "fariz", // case negative
	// 	},
	// 	expectedError: errors.New("invalid division"),
	// })

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			BookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(BookRepo)
			}

			service := Service{
				repo: BookRepo,
			}

			res, err := service.UpdateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

// func Test_BookService_DeleteBook(t *testing.T) {
// 	type testCase struct {
// 		name      string
// 		wantError bool
// 		input     model.Book
// 		// expectedResult model.Book
// 		expectedError error
// 		onBookRepo    func(mock *mocks.MockBookRepo)
// 	}

// 	var testTable []testCase

// 	testTable = append(testTable, testCase{
// 		name:      "success",
// 		wantError: false,
// 		input: model.Book{
// 			ID: 1,
// 		},
// 		onBookRepo: func(mock *mocks.MockBookRepo) {
// 			mock.EXPECT().DeleteBook(gomock.Any()).Return(nil).Times(1)
// 		},
// 	})

// 	testTable = append(testTable, testCase{
// 		name:          "record not found",
// 		wantError:     true,
// 		expectedError: errors.New("record not found"),
// 		onBookRepo: func(mock *mocks.MockBookRepo) {
// 			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("record not found")).Times(1)
// 		},
// 	})

// 	testTable = append(testTable, testCase{
// 		name:          "unexpected error",
// 		wantError:     true,
// 		expectedError: errors.New("unexpected error"),
// 		onBookRepo: func(mock *mocks.MockBookRepo) {
// 			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("unexpected error")).Times(1)
// 		},
// 	})

// 	for _, testCase := range testTable {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)

// 			BookRepo := mocks.NewMockBookRepo(mockCtrl)

// 			if testCase.onBookRepo != nil {
// 				testCase.onBookRepo(BookRepo)
// 			}

// 			service := Service{
// 				repo: BookRepo,
// 			}

// 			err := service.DeleteBook(1)

// 			if testCase.wantError {
// 				assert.EqualError(t, err, testCase.expectedError.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, testCase.expectedError, err)
// 			}
// 		})
// 	}
// }
