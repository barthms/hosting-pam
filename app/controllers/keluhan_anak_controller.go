package controllers

import (
	"errors"
	"monitoring-service/app/models"
	"monitoring-service/app/usecases"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type KeluhanAnakController struct {
	useCase usecases.KeluhanAnakUseCase
}

func NewKeluhanAnakController(useCase usecases.KeluhanAnakUseCase) *KeluhanAnakController {
	return &KeluhanAnakController{useCase}
}

func isKeluhanAnakForbidden(err error) bool {
	return errors.Is(err, usecases.ErrKeluhanAnakForbidden)
}

func (c *KeluhanAnakController) Create(ctx echo.Context) error {
	var data models.KeluhanAnak
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	if err := c.useCase.Create(&data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, models.Response{StatusCode: http.StatusCreated, Data: data})
}

func (c *KeluhanAnakController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "id tidak valid"})
	}
	var data models.KeluhanAnak
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	if err := c.useCase.Update(uint(id), &data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KeluhanAnakController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "id tidak valid"})
	}
	if err := c.useCase.Delete(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Message: "Deleted"})
}

func (c *KeluhanAnakController) GetByAnakID(ctx echo.Context) error {
	anakID, err := strconv.Atoi(ctx.Param("anak_id"))
	if err != nil || anakID <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "anak_id tidak valid"})
	}
	data, err := c.useCase.GetByAnakID(uint(anakID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KeluhanAnakController) GetByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "id tidak valid"})
	}
	data, err := c.useCase.GetByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KeluhanAnakController) GetByAnakIDForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	anakIDRaw := strings.TrimSpace(ctx.QueryParam("anak_id"))
	if anakIDRaw == "" {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "anak_id wajib diisi"})
	}

	anakID, err := strconv.Atoi(anakIDRaw)
	if err != nil || anakID <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "anak_id tidak valid"})
	}

	data, err := c.useCase.GetByAnakIDForIbu(uint(anakID), uint(claims.UserID))
	if err != nil {
		if isKeluhanAnakForbidden(err) {
			return ctx.JSON(http.StatusForbidden, models.Response{StatusCode: http.StatusForbidden, Message: "Anda tidak memiliki akses ke data anak ini"})
		}
		return ctx.JSON(http.StatusInternalServerError, models.Response{StatusCode: http.StatusInternalServerError, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}

func (c *KeluhanAnakController) GetByIDForIbu(ctx echo.Context) error {
	claims, ok := getAuthClaims(ctx)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, models.Response{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"})
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{StatusCode: http.StatusBadRequest, Message: "id tidak valid"})
	}

	data, err := c.useCase.GetByIDForIbu(uint(id), uint(claims.UserID))
	if err != nil {
		if isKeluhanAnakForbidden(err) {
			return ctx.JSON(http.StatusForbidden, models.Response{StatusCode: http.StatusForbidden, Message: "Anda tidak memiliki akses ke data anak ini"})
		}
		return ctx.JSON(http.StatusNotFound, models.Response{StatusCode: http.StatusNotFound, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, models.Response{StatusCode: http.StatusOK, Data: data})
}
