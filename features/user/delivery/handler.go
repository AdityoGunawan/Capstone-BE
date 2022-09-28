package delivery

import (
	"capstone-project/features/user"
	"capstone-project/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &userDelivery{
		userUsecase: usecase,
	}
	e.POST("/login", handler.LoginUser)
}

func (handler *userDelivery) LoginUser(c echo.Context) error {
	data := user.UserCore{}
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Fail_Resp("user doesn't exist"))
	}

	token, err := handler.userUsecase.PostLogin(data)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "login success",
		"token":   token,
	})
}
