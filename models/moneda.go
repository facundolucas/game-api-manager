package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Moneda struct {
	gorm.Model
	Value        int
	Description  string
	MonedaTypeID uint
	MonedaType   MonedaType
	ImageURL     string
	Models       []MonedaModel
}

type MonedaType struct {
	gorm.Model
	Code        string
	Description string
}

type MonedaModel struct {
	gorm.Model
	Path        string
	Description string
	MonedaID    uint
}

func (m *Moneda) SaveMoneda() (*Moneda, error) {

	var err error

	err = DB.Omit("MonedaType").Create(&m).Error

	if err != nil {
		return &Moneda{}, err
	}

	return m, nil
}

func GetMonedasOrderedByValue(code string) ([]Moneda, error) {

	var monedas []Moneda

	result := DB.
		Preload("Models").
		Joins("JOIN moneda_types ON monedas.type_id = moneda_types.id").
		Order("Value").
		Where("moneda_types.code = ?", code).
		Find(&monedas)

	if result.Error != nil {
		return nil, result.Error
	}

	return monedas, nil
}

func GetMonedaByID(id string) (Moneda, error) {

	var m Moneda

	if err := DB.First(&m, id).Error; err != nil {
		return m, errors.New("moneda not found")
	}

	return m, nil

}

func (m *Moneda) UpdateMoneda() (*Moneda, error) {

	var err error

	err = DB.Save(&m).Error

	if err != nil {
		return &Moneda{}, err
	}

	return m, nil
}
