package handler

// delivery, controller

import (
	"challenge-8/helper"
	"challenge-8/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = in.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetAllBook(c *gin.Context) {
	res, err := h.app.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	if len(res) == 0 {
		helper.Ok(c, res)
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.GetBookByID(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in.ID = idInt
	// call service
	res, err := h.app.UpdateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Successfully deleted book")
}

// func (h HttpServer) GetEmployeeByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// call service
// 	res, err := h.app.GetEmployeeByID(id)
// 	if err != nil {
// if err.Error() == helper.ErrNotFound {
// 	helper.NotFound(c, err.Error())
// 	return
// }
// 		helper.InternalServerError(c, err.Error())
// 		return
// 	}

// 	helper.Ok(c, res)
// }
