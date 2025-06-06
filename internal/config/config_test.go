package config

import (
	"os"
	"testing"
)

func TestNewConfig_WithEnvPort(t *testing.T) {
	t.Run("Test with env port", func(t *testing.T) {
		os.Setenv("PORT", "7456")
		defer os.Unsetenv("PORT")

		cfg := NewConfig()

		if cfg.Port != "7456" {
			t.Errorf("Expected port '7456', got '%s'", cfg.Port)
		}
	})

	t.Run("Test defualt port", func(t *testing.T) {
		os.Unsetenv("PORT")

		cfg := NewConfig()

		if cfg.Port != "8080" {
			t.Errorf("Expected default port '8080', got '%s'", cfg.Port)
		}
	})
}
