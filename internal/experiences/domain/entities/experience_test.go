package experiences_domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
)

func TestShouldBeAbleToCreateExperience(t *testing.T) {
	tech1, _ := technologies_domain.NewTechnology("TECH_TEST_1", "TECH_1.PNG", true)
	tech2, _ := technologies_domain.NewTechnology("TECH_TEST_2", "TECH_2.PNG", true)
	experience, _ := NewExperience(
		"POS_TEST",
		"COMPANY_TEST",
		"LOGO_TEST.SVG",
		"01/2023",
		"",
		[]string{"SUMMARY_TEST_1", "SUMMARY_TEST_2", "SUMMARY_TEST_3"},
		[]technologies_domain.Technology{*tech1, *tech2},
	)

	assert.NotEmpty(t, experience.ID)
	assert.Equal(t, len(experience.Summary), 3)
	assert.Equal(t, len(experience.Techs), 2)
}

func TestPositionShouldNotBeEmpty(t *testing.T) {
	experience := Experience{}

	assert.Error(t, experience.Validate(), "position is required")
}

func TestCompanyShouldNotBeEmpty(t *testing.T) {
	experience := Experience{
		Position: "POS_TEST",
	}

	assert.Error(t, experience.Validate(), "company is required")
}

func TestLogoShouldNotBeEmpty(t *testing.T) {
	experience := Experience{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
	}

	assert.Error(t, experience.Validate(), "logo is required")
}

func TestStartShouldNotBeEmpty(t *testing.T) {
	experience := Experience{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
		Logo:     "LOGO_TEST.SVG",
	}

	assert.Error(t, experience.Validate(), "period start is required")
}

func TestSummaryShouldNotBeEmpty(t *testing.T) {
	experience := Experience{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
		Logo:     "LOGO_TEST.SVG",
		Period:   Period{Start: "01/2023"},
	}

	assert.Error(t, experience.Validate(), "summary is required")
}

func TestTechsShouldNotBeEmpty(t *testing.T) {
	experience := Experience{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
		Logo:     "LOGO_TEST.SVG",
		Period:   Period{Start: "01/2023"},
		Summary:  []string{"SUM_1", "SUM_2"},
	}

	assert.Error(t, experience.Validate(), "techs are required")
}
