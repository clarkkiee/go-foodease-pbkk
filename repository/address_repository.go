package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-foodease-be/dto"
	"go-foodease-be/models"
	"go-foodease-be/types"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

type (
	AddressRepository interface {
		ProduceCordFromText(ctx context.Context, tx *gorm.DB, street string) (*types.Coordinates, error)
		CreateAddress(ctx context.Context, tx *gorm.DB, address models.Address) (models.Address, error)
	}

	addressRepository struct {
		db *gorm.DB
	}
)

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) ProduceCordFromText(ctx context.Context, tx *gorm.DB, street string) (*types.Coordinates, error){
	url := fmt.Sprintf("https://geocode.search.hereapi.com/v1/geocode?q=%s&apiKey=%s", street, os.Getenv("GEOCODE_SECRET_API_KEY"))
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &types.Coordinates{}, nil
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %w", err)
	}
	defer resp.Body.Close()


	body, _ := io.ReadAll(resp.Body)

	var apiResponse dto.GeocodeAPIResponse
	json.Unmarshal(body, &apiResponse)
	
	var coords types.Coordinates
	if len(apiResponse.Items) > 0 {
		position := apiResponse.Items[0].Position
		coords.Latitude = position.Lat
		coords.Longitude = position.Lng
	} else {
		return &types.Coordinates{}, errors.New("failed to fetch coordinates data")
	}
	
	return &coords, nil
}

func (r *addressRepository) CreateAddress(ctx context.Context, tx *gorm.DB, address models.Address) (models.Address, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&address).Error; err != nil {
		return models.Address{}, err
	}

	return address, nil
}