package technologies_domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameShouldNotBeBlank(t *testing.T) {
	tech := Technology{}

	assert.Error(t, tech.Validate(), "name is required")
}

func TestSrcShouldNotBeBlank(t *testing.T) {
	tech := Technology{
		Name: "test",
	}

	assert.Error(t, tech.Validate(), "src is required")
}

func TestShouldBeAbleToCreateTechnology(t *testing.T) {
	tech, _ := NewTechnology("test name", "src.png", true)

	assert.NotEmpty(t, tech.ID)
	assert.Equal(t, tech.Name, "test name")
	assert.Equal(t, tech.Src, "src.png")
	assert.Equal(t, tech.Display, true)
}
