// package controllers

// import (
// 	"monitoring-service/app/usecases"
// 	"monitoring-service/pkg/customerror"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type PemeriksaanDokterCombinedController struct {
// 	usecaseT1 usecases.PemeriksaanDokterTrimester1Usecase
// 	usecaseT3 usecases.PemeriksaanDokterTrimester3Usecase
// }

// func NewPemeriksaanDokterCombinedController(
// 	u1 usecases.PemeriksaanDokterTrimester1Usecase,
// 	u3 usecases.PemeriksaanDokterTrimester3Usecase,
// ) *PemeriksaanDokterCombinedController {
// 	return &PemeriksaanDokterCombinedController{
// 		usecaseT1: u1,
// 		usecaseT3: u3,
// 	}
// }

// // T1 Complete
// func (c *PemeriksaanDokterCombinedController) CreateT1(ctx echo.Context) error {
// 	var req usecases.PemeriksaanDokterTrimester1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
// 	}
// 	if err := c.usecaseT1.Create(&req); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, map[string]interface{}{
// 		"status_code": http.StatusCreated,
// 		"message":     "Data pemeriksaan dokter trimester 1 & lab jiwa berhasil disimpan",
// 	})
// }

// func (c *PemeriksaanDokterCombinedController) UpdateT1(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	var req usecases.PemeriksaanDokterTrimester1Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
// 	}
// 	if err := c.usecaseT1.Update(int32(id), &req); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"status_code": http.StatusOK,
// 		"message":     "Data berhasil diperbarui",
// 	})
// }

// func (c *PemeriksaanDokterCombinedController) GetT1ByID(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	data, err := c.usecaseT1.GetByID(int32(id))
// 	if err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, data)
// }

// func (c *PemeriksaanDokterCombinedController) GetT1ByKehamilan(ctx echo.Context) error {
// 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
// 	}
// 	list, err := c.usecaseT1.GetByKehamilanID(int32(kehamilanID))
// 	if err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}

// 	// Jika tidak ada data, return dengan struktur kosong
// 	if len(list) == 0 {
// 		return ctx.JSON(http.StatusOK, map[string]interface{}{
// 			"dokter":   nil,
// 			"lab_jiwa": nil,
// 		})
// 	}

// 	// Return hanya data pertama dengan struktur yang sesuai frontend
// 	firstData := list[0]
// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"dokter":   firstData.Dokter,
// 		"lab_jiwa": firstData.Lab,
// 	})
// }

// func (c *PemeriksaanDokterCombinedController) DeleteT1(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	if err := c.usecaseT1.Delete(int32(id)); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
// }

// // T3 Complete (opsional, analog)
// func (c *PemeriksaanDokterCombinedController) CreateT3(ctx echo.Context) error {
// 	var req usecases.PemeriksaanDokterTrimester3Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
// 	}
// 	if err := c.usecaseT3.Create(&req); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusCreated, map[string]interface{}{
// 		"status_code": http.StatusCreated,
// 		"message":     "Data pemeriksaan dokter trimester 3 & lanjutan berhasil disimpan",
// 	})
// }

// func (c *PemeriksaanDokterCombinedController) UpdateT3(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	var req usecases.PemeriksaanDokterTrimester3Request
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
// 	}
// 	if err := c.usecaseT3.Update(int32(id), &req); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"status_code": http.StatusOK,
// 		"message":     "Data berhasil diperbarui",
// 	})
// }

// func (c *PemeriksaanDokterCombinedController) GetT3ByID(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	data, err := c.usecaseT3.GetByID(int32(id))
// 	if err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, data)
// }

// func (c *PemeriksaanDokterCombinedController) GetT3ByKehamilan(ctx echo.Context) error {
// 	kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
// 	}
// 	list, err := c.usecaseT3.GetByKehamilanID(int32(kehamilanID))
// 	if err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}

// 	// Jika tidak ada data, return dengan struktur kosong
// 	if len(list) == 0 {
// 		return ctx.JSON(http.StatusOK, map[string]interface{}{
// 			"dokter":   nil,
// 			"lab_jiwa": nil,
// 		})
// 	}

// 	// Return hanya data pertama dengan struktur yang sesuai frontend
// 	firstData := list[0]
// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"dokter":   firstData.Dokter,
// 		"lab_jiwa": firstData.Lab,
// 	})
// }

// func (c *PemeriksaanDokterCombinedController) DeleteT3(ctx echo.Context) error {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
// 	}
// 	if err := c.usecaseT3.Delete(int32(id)); err != nil {
// 		statusCode := customerror.GetStatusCode(err)
// 		return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
// }


package controllers

import (
    "monitoring-service/app/models" // Ditambahkan
    "monitoring-service/app/usecases"
    "monitoring-service/pkg/customerror"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

type PemeriksaanDokterCombinedController struct {
    usecaseT1 usecases.PemeriksaanDokterTrimester1Usecase
    usecaseT3 usecases.PemeriksaanDokterTrimester3Usecase
}

func NewPemeriksaanDokterCombinedController(
    u1 usecases.PemeriksaanDokterTrimester1Usecase,
    u3 usecases.PemeriksaanDokterTrimester3Usecase,
) *PemeriksaanDokterCombinedController {
    return &PemeriksaanDokterCombinedController{
        usecaseT1: u1,
        usecaseT3: u3,
    }
}

// T1 Complete
func (c *PemeriksaanDokterCombinedController) CreateT1(ctx echo.Context) error {
    var req models.PemeriksaanDokterTrimester1 // DIUBAH
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
    }
    if err := c.usecaseT1.Create(&req); err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusCreated, map[string]interface{}{
        "status_code": http.StatusCreated,
        "message":     "Data pemeriksaan dokter trimester 1 & lab jiwa berhasil disimpan",
    })
}

func (c *PemeriksaanDokterCombinedController) UpdateT1(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    var req models.PemeriksaanDokterTrimester1 // DIUBAH
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
    }
    if err := c.usecaseT1.Update(int32(id), &req); err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "status_code": http.StatusOK,
        "message":     "Data berhasil diperbarui",
    })
}

func (c *PemeriksaanDokterCombinedController) GetT1ByID(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    data, err := c.usecaseT1.GetByID(int32(id))
    if err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, data)
}

func (c *PemeriksaanDokterCombinedController) GetT1ByKehamilan(ctx echo.Context) error {
    kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
    }
    list, err := c.usecaseT1.GetByKehamilanID(int32(kehamilanID))
    if err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }

    if len(list) == 0 {
        return ctx.JSON(http.StatusOK, map[string]interface{}{
            "dokter":   nil,
            "lab_jiwa": nil,
        })
    }

    firstData := list[0]
    // DIUBAH: Karena model tidak punya field Dokter & Lab, data utuh dikembalikan ke key "dokter"
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "dokter":   firstData, 
        "lab_jiwa": nil,        
    })
}

func (c *PemeriksaanDokterCombinedController) DeleteT1(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    if err := c.usecaseT1.Delete(int32(id)); err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
}

// T3 Complete (opsional, analog)
func (c *PemeriksaanDokterCombinedController) CreateT3(ctx echo.Context) error {
    var req models.PemeriksaanDokterTrimester3 // DIUBAH
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
    }
    if err := c.usecaseT3.Create(&req); err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusCreated, map[string]interface{}{
        "status_code": http.StatusCreated,
        "message":     "Data pemeriksaan dokter trimester 3 & lanjutan berhasil disimpan",
    })
}

func (c *PemeriksaanDokterCombinedController) UpdateT3(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    var req models.PemeriksaanDokterTrimester3 // DIUBAH
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "format request tidak valid: " + err.Error()})
    }
    if err := c.usecaseT3.Update(int32(id), &req); err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "status_code": http.StatusOK,
        "message":     "Data berhasil diperbarui",
    })
}

func (c *PemeriksaanDokterCombinedController) GetT3ByID(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    data, err := c.usecaseT3.GetByID(int32(id))
    if err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, data)
}

func (c *PemeriksaanDokterCombinedController) GetT3ByKehamilan(ctx echo.Context) error {
    kehamilanID, err := strconv.ParseInt(ctx.QueryParam("kehamilan_id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "kehamilan_id required"})
    }
    list, err := c.usecaseT3.GetByKehamilanID(int32(kehamilanID))
    if err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }

    if len(list) == 0 {
        return ctx.JSON(http.StatusOK, map[string]interface{}{
            "dokter":   nil,
            "lab_jiwa": nil,
        })
    }

    firstData := list[0]
    // DIUBAH: Sama seperti T1
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "dokter":   firstData,
        "lab_jiwa": nil,
    })
}

func (c *PemeriksaanDokterCombinedController) DeleteT3(ctx echo.Context) error {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
    }
    if err := c.usecaseT3.Delete(int32(id)); err != nil {
        statusCode := customerror.GetStatusCode(err)
        return ctx.JSON(statusCode, map[string]string{"message": err.Error()})
    }
    return ctx.JSON(http.StatusOK, map[string]string{"message": "Data berhasil dihapus"})
}