package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/odamilola36/golang_api/dto"
	"github.com/odamilola36/golang_api/entity"
	"github.com/odamilola36/golang_api/helper"
	"github.com/odamilola36/golang_api/service"
)


type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	//this is where we put our services
	authService service.AuthService
	jwtService service.Jwtservice
}


func NewAuthController(jwtService service.Jwtservice, authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context)  {
	var loginDTO dto.LoginDTO
	//bound
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredentials(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check your credentials", "Invalid credentials", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context)  {
	var registerDto dto.RegisterDTO
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		response := helper.BuildErrorResponse("Failed to register", errDto.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Failed to process registration", "Duplicate email address", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else { 
		createUser := c.authService.CreateUser(registerDto)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		response := helper.BuildResponse(true, "OK!", createUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
