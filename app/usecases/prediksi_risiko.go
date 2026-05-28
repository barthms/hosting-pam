package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"net/http"
)

type PrediksiRisikoUsecase interface {
    Predict(req *models.PrediksiRisikoRequest) (*models.PrediksiRisikoResponse, error)
}

type prediksiRisikoUsecase struct {
    mlServiceURL string
}

func NewPrediksiRisikoUsecase(mlServiceURL string) PrediksiRisikoUsecase {
    return &prediksiRisikoUsecase{mlServiceURL: mlServiceURL}
}

func (u *prediksiRisikoUsecase) Predict(req *models.PrediksiRisikoRequest) (*models.PrediksiRisikoResponse, error) {
    jsonReq, err := json.Marshal(req)
	 fmt.Println("📤 [ML Request] Payload ke Python:", string(jsonReq)) // <-- log ini
	 
    // if err != nil {
    //     return nil, err
    // }

    resp, err := http.Post(u.mlServiceURL+"/predict", "application/json", bytes.NewBuffer(jsonReq))
    if err != nil {
        return nil, errors.New("ML service tidak tersedia")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("ML service mengembalikan error")
    }

    var result models.PrediksiRisikoResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return &result, nil
}