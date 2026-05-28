// package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"monitoring-service/app/models"
// 	"monitoring-service/app/usecases"

// 	"github.com/labstack/echo/v4"
// )

// type GrafikPeningkatanBBController struct {
// 	usecase usecases.GrafikPeningkatanBBUsecase
// }

// func NewGrafikPeningkatanBBController(u usecases.GrafikPeningkatanBBUsecase) *GrafikPeningkatanBBController {
// 	return &GrafikPeningkatanBBController{usecase: u}
// }

// type createGrafikBBRequest struct {
// 	KehamilanID     int32   `json:"kehamilan_id"`
// 	BeratBadan      float64 `json:"berat_badan"`
// 	MingguKehamilan int     `json:"minggu_kehamilan"`
// }

// //
// // CREATE
// //
// func (c *GrafikPeningkatanBBController) Create(ctx echo.Context) error {
// 	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
// 	if !ok || claims == nil {
// 		return ctx.JSON(http.StatusUnauthorized, models.Response{
// 			StatusCode: http.StatusUnauthorized,
// 			Message:    "Unauthorized",
// 		})
// 	}

// 	var req createGrafikBBRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 		})
// 	}

// 	// ✅ VALIDASI DASAR
// 	if req.KehamilanID == 0 {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "kehamilan_id wajib diisi",
// 		})
// 	}

// 	if req.BeratBadan <= 0 {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "berat badan harus > 0",
// 		})
// 	}

// 	if req.MingguKehamilan < 0 || req.MingguKehamilan > 42 {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "minggu kehamilan harus 0 - 42",
// 		})
// 	}

// 	g := &models.GrafikPeningkatanBB{
// 		KehamilanID:     req.KehamilanID,
// 		BeratBadan:      &req.BeratBadan,
// 		MingguKehamilan: &req.MingguKehamilan,
// 	}

// 	if err := c.usecase.Create(g); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusCreated, models.Response{
// 		StatusCode: http.StatusCreated,
// 		Data:       g,
// 	})
// }

// //
// // GET BY ID
// //
// func (c *GrafikPeningkatanBBController) GetByID(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "invalid id",
// 		})
// 	}

// 	data, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{
// 			StatusCode: http.StatusNotFound,
// 			Message:    err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       data,
// 	})
// }

// //
// // GET BY KEHAMILAN ID (UNTUK GRAFIK)
// //
// func (c *GrafikPeningkatanBBController) GetByKehamilanID(ctx echo.Context) error {
// 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "kehamilan_id wajib",
// 		})
// 	}

// 	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       list,
// 	})
// }

// //
// // UPDATE
// //
// func (c *GrafikPeningkatanBBController) Update(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "invalid id",
// 		})
// 	}

// 	var req createGrafikBBRequest
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 		})
// 	}

// 	existing, err := c.usecase.GetByID(int32(id))
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, models.Response{
// 			StatusCode: http.StatusNotFound,
// 			Message:    "data tidak ditemukan",
// 		})
// 	}

// 	if req.BeratBadan > 0 {
// 		existing.BeratBadan = &req.BeratBadan
// 	}

// 	if req.MingguKehamilan >= 0 && req.MingguKehamilan <= 42 {
// 		existing.MingguKehamilan = &req.MingguKehamilan
// 	}

// 	if err := c.usecase.Update(existing); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Data:       existing,
// 	})
// }

// //
// // DELETE
// //
// func (c *GrafikPeningkatanBBController) Delete(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, models.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "invalid id",
// 		})
// 	}

// 	if err := c.usecase.Delete(int32(id)); err != nil {
// 		return ctx.JSON(http.StatusInternalServerError, models.Response{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 		})
// 	}

// 	return ctx.JSON(http.StatusOK, models.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "deleted",
// 	})
// }

package controllers

import (
	"net/http"
	"strconv"

	"monitoring-service/app/models"
	"monitoring-service/app/usecases"

	"github.com/labstack/echo/v4"
)

type GrafikPeningkatanBBController struct {
	usecase usecases.GrafikPeningkatanBBUsecase
}

func NewGrafikPeningkatanBBController(u usecases.GrafikPeningkatanBBUsecase) *GrafikPeningkatanBBController {
	return &GrafikPeningkatanBBController{usecase: u}
}

type createGrafikBBRequest struct {
	KehamilanID     int32   `json:"kehamilan_id"`
	BeratBadan      float64 `json:"berat_badan"`
	MingguKehamilan int     `json:"minggu_kehamilan"`
}

// CREATE
func (c *GrafikPeningkatanBBController) Create(ctx echo.Context) error {
	var req createGrafikBBRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// VALIDASI
	if req.KehamilanID == 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id wajib diisi",
		})
	}

	if req.BeratBadan <= 0 {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "berat badan harus > 0",
		})
	}

	if req.MingguKehamilan < 0 || req.MingguKehamilan > 42 {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "minggu kehamilan harus 0 - 42",
		})
	}

	g := &models.GrafikPeningkatanBB{
		KehamilanID:     req.KehamilanID,
		BeratBadan:      &req.BeratBadan,
		MingguKehamilan: &req.MingguKehamilan,
	}

	err := c.usecase.Create(g)
	if err != nil {

		// 🔥 HANDLE ERROR BISNIS
		if err.Error() == "data minggu ini sudah ada" {
			return ctx.JSON(http.StatusConflict, models.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
			})
		}

		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Data:       g,
	})
}

// GET BY ID
func (c *GrafikPeningkatanBBController) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	data, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

// GET BY KEHAMILAN ID (UNTUK GRAFIK)
func (c *GrafikPeningkatanBBController) GetByKehamilanID(ctx echo.Context) error {
	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "kehamilan_id wajib",
		})
	}

	list, err := c.usecase.GetByKehamilanID(int32(kehamilanID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       list,
	})
}

// UPDATE
func (c *GrafikPeningkatanBBController) Update(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	var req createGrafikBBRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	existing, err := c.usecase.GetByID(int32(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "data tidak ditemukan",
		})
	}

	// VALIDASI & UPDATE FIELD
	if req.BeratBadan > 0 {
		existing.BeratBadan = &req.BeratBadan
	}

	if req.MingguKehamilan >= 0 && req.MingguKehamilan <= 42 {
		existing.MingguKehamilan = &req.MingguKehamilan
	}

	err = c.usecase.Update(existing)
	if err != nil {

		if err.Error() == "data minggu ini sudah ada" {
			return ctx.JSON(http.StatusConflict, models.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
			})
		}

		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       existing,
	})
}

// DELETE
func (c *GrafikPeningkatanBBController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
		})
	}

	if err := c.usecase.Delete(int32(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "data berhasil dihapus",
	})
}

// GET GRAFIK BB UNTUK IBU (MODUL IBU) - /v2
func (c *GrafikPeningkatanBBController) GetGrafikBBForOrangtua(ctx echo.Context) error {
	claims, ok := ctx.Get("auth_claims").(*models.AuthClaims)
	if !ok || claims == nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "token tidak valid",
		})
	}

	data, err := c.usecase.GetGrafikBBForOrangtua(claims.UserID)
	if err != nil {
		if err.Error() == "kehamilan aktif tidak ditemukan" {
			return ctx.JSON(http.StatusNotFound, models.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
