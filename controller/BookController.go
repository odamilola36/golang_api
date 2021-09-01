package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/odamilola36/golang_api/dto"
	"github.com/odamilola36/golang_api/entity"
	"github.com/odamilola36/golang_api/helper"
	"github.com/odamilola36/golang_api/service"
	"net/http"
	"strconv"
)

type BookController interface {
	All(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.Jwtservice
}

func NewBookController(bookService service.BookService, jwtService service.Jwtservice) BookController {
	return &bookController{
		bookService: bookService,
		jwtService:  jwtService,
	}
}

func (c *bookController) getUserIdByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}

func (b bookController) All(ctx *gin.Context) {
	var books []entity.Book
	books = b.bookService.All()
	var res = helper.BuildResponse(true, "OK!", books)
	ctx.JSON(http.StatusOK, res)
}

func (b bookController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book = b.bookService.FindById(id)
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("no book with the given id", "no data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	} else {
		res := helper.BuildResponse(true, "OK!", book)
		ctx.JSON(http.StatusOK, res)
		return
	}
}

func (b bookController) Insert(ctx *gin.Context) {
	var bookCreate dto.BookCreateDTO
	var errDto = ctx.ShouldBind(bookCreate)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		authToken := ctx.GetHeader("Authorization")
		userId := b.getUserIdByToken(authToken)
		convUserId, err := strconv.ParseUint(userId, 10, 64)
		if err == nil {
			bookCreate.UserID = convUserId
		}
		result := b.bookService.Insert(bookCreate)
		response := helper.BuildResponse(true, "OK!", result)
		ctx.JSON(http.StatusOK, response)
	}
}

func (b bookController) Update(ctx *gin.Context) {
	var bookUpdate dto.BookUpdateDTO
	errDto := ctx.ShouldBind(bookUpdate)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to proccess request", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authToken := ctx.GetHeader("Authorization")
	token, err := b.jwtService.ValidateToken(authToken)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])
	if b.bookService.IsAllowedToEdit(userId, bookUpdate.ID) {
		id, errId := strconv.ParseUint(userId, 10, 64)
		if errId == nil {
			bookUpdate.UserID = id
		}
		result := b.bookService.Update(bookUpdate)
		response := helper.BuildResponse(true, "OK!", result)
		ctx.JSON(http.StatusOK, response)
		return
	} else {
		response := helper.BuildErrorResponse("Invalid action", "user cannot update book", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
}

func (b bookController) Delete(ctx *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get id", "No param id found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	book.ID = id
	authToken := ctx.GetHeader("Authorization")
	token, err := b.jwtService.ValidateToken(authToken)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])
	if b.bookService.IsAllowedToEdit(userId, uint64(id)) {
		b.bookService.Delete(book)
		res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("Invalid action", "user cannot update book", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
}
