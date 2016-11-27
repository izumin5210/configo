package configo_test

import (
	"os"
	"testing"

	"github.com/creasty/configo"
)

type Sample struct {
	ValFromDotenv string
	ValFromEnv    string
	ValFromYaml   string

	Nested struct {
		Override1 string
		Override2 string
	}
}

func TestLoad(t *testing.T) {
	s := &Sample{}

	os.Clearenv()
	os.Setenv("APP_ENV", "production")
	os.Setenv("APP_VALFROMENV", "env")

	if err := configo.Load(s, configo.Option{Dir: "./data/config"}); err != nil {
		t.Error(err)
		return
	}

	if s.ValFromEnv != "env" {
		t.Error("should read from env")
	}

	if s.ValFromDotenv != "dotenv" {
		t.Error("should read from dotenv")
	}

	if s.ValFromYaml != "yaml" {
		t.Error("should read from default.yml")
	}

	if s.Nested.Override1 != "production" {
		t.Error("should read from production.yml")
	}

	if s.Nested.Override2 != "production.local" {
		t.Error("should read from production.local.yml")
	}
}
