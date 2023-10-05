package api

import (
	"encoding/json"
	"hackforachange-demo-backend/models"
	"os"
	"path/filepath"
)

func LoadGrades() ([]models.Grade, error) {
	path, err := filepath.Abs("data/database.json")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var jsonData models.JsonData
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData.Grades, nil
}

func GetSubjectByGradeId(gradeID int) ([]models.Subject, error) {
	grades, err := LoadGrades()
	if err != nil {
		return nil, err
	}

	for _, grade := range grades {
		if grade.ID == gradeID {
			return grade.Subjects, nil
		}
	}

	return nil, nil
}
