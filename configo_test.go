package configo_test

import (
	"os"
	"testing"

	"github.com/creasty/configo"
)

type Sample1 struct {
	ValFromDotenv string
	ValFromEnv    string
	ValFromYaml   string

	Nested1 struct {
		Override1 string
		Override2 string
		Override3 string
	}

	Nested2 struct {
		Override1 string
		Override2 string
	}
}

type Sample2 struct {
	Required string `valid:"required"`
}

func init() {
	os.Clearenv()
}

func load(in interface{}) error {
	return configo.Load(in, configo.Option{
		Dir: "./data/config",
	})
}

func TestLoad(t *testing.T) {
	s := &Sample1{}

	os.Setenv("APP_ENV", "production")
	os.Setenv("APP_VALFROMENV", "env")
	os.Setenv("APP_NESTED1_OVERRIDE3", "env")

	if err := load(s); err != nil {
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

	if s.Nested1.Override1 != "production" {
		t.Error("should be overrided by production.yml")
	}

	if s.Nested1.Override2 != "production.local" {
		t.Error("should be overrided by production.local.yml")
	}

	if s.Nested1.Override3 != "env" {
		t.Error("should be overrided by env")
	}

	if s.Nested2.Override1 != "production" {
		t.Error("should be overrided by nested/production.yml")
	}

	if s.Nested2.Override2 != "production.local" {
		t.Error("should be overrided by nested/production.local.yml")
	}
}

func TestValidation(t *testing.T) {
	s := &Sample2{}

	if err := load(s); err == nil {
		t.Error("should return error if a struct is not valid")
	}

	os.Setenv("APP_REQUIRED", "foo")

	if err := load(s); err != nil {
		t.Error("should validate a struct")
	}
}
