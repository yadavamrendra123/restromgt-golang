package services

import (
	"errors"
	"gorm.io/gorm"
	"restro-mgt/database"
	"restro-mgt/models"
	"time"
)

func CreateRestaurant(restaurant *models.Restaurant) error {
	result := database.DB.Create(restaurant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	result := database.DB.Find(&restaurants)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range restaurants {
		if restaurants[i].CreatedAt.IsZero() {
			restaurants[i].CreatedAt = time.Time{}
		}
		if restaurants[i].UpdatedAt.IsZero() {
			restaurants[i].UpdatedAt = time.Time{}
		}
	}

	return restaurants, nil
}

func GetRestaurantByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	result := database.DB.First(&restaurant, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("restaurant not found")
		}
		return nil, result.Error
	}
	return &restaurant, nil
}

func UpdateRestaurant(restaurant *models.Restaurant) error {
	result := database.DB.Save(restaurant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRestaurant(id uint) error {
	result := database.DB.Delete(&models.Restaurant{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
