package api

import (
	"encoding/json"
	"hackforachange-demo-backend/models"
	"os"
	"path/filepath"
)

func LoadGrades() ([]models.Grade, error) {
	data, err := os.ReadFile(filepath.Join("data", "database.json"))
	if err != nil {
		return nil, err
	}

	var grades []models.Grade
	if err := json.Unmarshal(data, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}
