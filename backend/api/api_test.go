package api

import (
	"hackforachange-demo-backend/models"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoadGrades(t *testing.T) {
	grades, err := LoadGrades()

	if err != nil {
		t.Error("Error loading grades: ", err)
	}

	expectedGrade1 := models.Grade{
		ID:   1,
		Name: "1st Grade",
		Subjects: []models.Subject{
			{ID: 1, Name: "Mathematics"},
			{ID: 2, Name: "Science"},
			{ID: 3, Name: "Portuguese"},
		},
	}

	diff := cmp.Diff(grades[0], expectedGrade1)

	if diff != "" {
		t.Errorf("Mismatch in loaded grade 1: %s", diff)
	}
}
