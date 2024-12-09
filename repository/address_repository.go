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
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	AddressRepository interface {
		ProduceCordFromText(ctx context.Context, tx *gorm.DB, street string) (*types.Coordinates, error)
		CreateAddress(ctx context.Context, tx *gorm.DB, address models.Address) (models.Address, error)
		GetAllAddressByCustomerId(ctx context.Context, tx *gorm.DB, customerId string) ([]dto.AddressResponse, error)
		GetAddressById(ctx context.Context, tx *gorm.DB, addressId string, customerId string) (dto.AddressResponse, error)
		UpdateAddressById(ctx context.Context, tx *gorm.DB, addressId string, address models.Address) (models.Address, error)
		DeleteAddressById(ctx context.Context, tx *gorm.DB, addressId string) error
		GetActiveAddress(ctx context.Context, tx *gorm.DB, entityId string) (dto.AddressResponse, error)
		SetActiveAddress(ctx context.Context, tx *gorm.DB, addressId string, customerId string) error
		GetUserActiveCoordinates(ctx context.Context, tx *gorm.DB, userId string) (dto.UserActiveCoordinatesResult, error) 
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
	
	fmt.Println(url)
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
	
	fmt.Println(apiResponse)
	var coords types.Coordinates
	
	if len(apiResponse.Items) > 0 {
		position := apiResponse.Items[0].Position
		coords.Latitude = position.Lat
		coords.Longitude = position.Lng
	} else {
		fmt.Println(coords)
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

func (r *addressRepository) GetAllAddressByCustomerId(ctx context.Context, tx *gorm.DB, customerId string) (addrResp []dto.AddressResponse, error error) {
	if tx == nil {
		tx = r.db
	}

	var addresses []models.Address

	err := tx.Raw(`SELECT id,
    street,
	ST_AsText(coordinates) AS coordinates,
    created_at,
    updated_at
	FROM addresses
	WHERE customer_id = ?`, customerId).Scan(&addresses).Error

	if err != nil {
		log.Fatalf("Error fetching addresses: %v", err)
	}

	for _, address := range addresses {
		addrResp = append(addrResp, dto.AddressResponse{
			ID: address.ID,
			Street: address.Street,
			Longitude: address.Coordinates.Longitude,
			Latitude: address.Coordinates.Latitude,
			CreatedAt: address.CreatedAt,
			UpdatedAt: address.UpdatedAt,
		})
	}
	return addrResp, nil
}

func (r *addressRepository) GetAddressById(ctx context.Context, tx *gorm.DB, addressId string, customerId string) (dto.AddressResponse, error){
	if tx == nil {
		tx = r.db
	}

	var address models.Address
	
	err := tx.Raw(`SELECT id,
    street,
	ST_AsText(coordinates) AS coordinates,
	customer_id,
    created_at,
    updated_at
	FROM addresses
	WHERE id = ?`, addressId).Scan(&address).Error

	if err != nil {
		return dto.AddressResponse{}, err
	}

	if address.CustomerID == uuid.Nil {
		return dto.AddressResponse{}, errors.New("address not found")
	}

	if strings.Compare(address.CustomerID.String(), customerId) != 0 {
		return dto.AddressResponse{}, errors.New("unauthorized to fetch another user address")
	}

	resp := dto.AddressResponse{
		ID: address.ID,
		Street: address.Street,
		Longitude: address.Coordinates.Longitude,
		Latitude: address.Coordinates.Latitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}

	return resp, nil
}

func (r *addressRepository) UpdateAddressById(ctx context.Context, tx *gorm.DB, addressId string, address models.Address) (models.Address, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&models.Address{}).Where("id = ?", addressId).Updates(map[string]interface{}{
		"street": address.Street,
		"coordinates": address.Coordinates,
	}).Error; err != nil {
		return models.Address{}, err
	}

	return address, nil
}

func (r *addressRepository) DeleteAddressById(ctx context.Context, tx *gorm.DB, addressId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&models.Address{}, "id = ?", addressId).Error; err != nil {
		return err
	}

	return nil
}

func (r *addressRepository) GetActiveAddress(ctx context.Context, tx *gorm.DB, entityId string) (dto.AddressResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var address models.Address
	err := tx.Raw(`SELECT id,
    street,
	ST_AsText(coordinates) AS coordinates,
	customer_id,
    created_at,
    updated_at
	FROM addresses
	WHERE id = (
		SELECT active_address_id FROM customers
		WHERE id = ?
	)`, entityId).Scan(&address).Error

	if err != nil {
		return dto.AddressResponse{}, err
	}

	response := dto.AddressResponse{
		ID: address.ID,
		Street: address.Street,
		Longitude: address.Coordinates.Longitude,
		Latitude: address.Coordinates.Latitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,		
	}

	return response, nil

}

func (r *addressRepository) SetActiveAddress(ctx context.Context, tx *gorm.DB, addressId string, customerId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&models.Customer{}).Where("id = ?", customerId).Update("active_address_id", addressId).Error; err != nil {
		return err
	}

	return nil
}

func (r *addressRepository) GetUserActiveCoordinates(ctx context.Context, tx *gorm.DB, userId string) (dto.UserActiveCoordinatesResult, error) {
	if tx == nil {
		tx = r.db
	}

	var res dto.UserActiveCoordinatesResult

	subQuery := tx.WithContext(ctx).Model(&models.Customer{}).Select("active_address_id").Where("id = ?", userId)
	if err := tx.WithContext(ctx).Model(&models.Address{}).Select("id", "coordinates").Where("id = (?)", subQuery).Take(&res).Error; err != nil {
		return dto.UserActiveCoordinatesResult{}, err
	}

	return res, nil

}