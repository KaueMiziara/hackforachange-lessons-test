package api

import (
	"encoding/json"
	"hackforachange-demo-backend/models"
	"os"
)

func LoadGrades() ([]models.Grade, error) {
	data, err := os.ReadFile("../data/database.json")
	if err != nil {
		return nil, err
	}

	var jsonData models.JsonData
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData.Grades, nil
}
