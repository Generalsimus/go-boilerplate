package config

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/danielgtaylor/huma/v2"
	"github.com/joho/godotenv"
)

var Cfg *AppConfig

type AppConfig struct {
	ENV          string `env:"ENV" envDefault:"dev" json:"ENV" required:"true" enum:"dev,prod"`
	PORT         int    `env:"PORT" envDefault:"8080" json:"PORT" required:"true" minimum:"1024" maximum:"65535"`
	DATABASE_URL string `env:"DATABASE_URL,required" json:"DATABASE_URL" required:"true"`
}

func init() {
	_ = godotenv.Load()

	var tempCfg AppConfig // Load into a temporary variable first

	if err := env.Parse(&tempCfg); err != nil {
		log.Fatalf("🚨 Environment loading failed: %v", err)
	}

	// Convert to JSON Map for Huma Validation
	rawBytes, _ := json.Marshal(tempCfg)
	var jsonMap map[string]any
	json.Unmarshal(rawBytes, &jsonMap)

	validator := huma.NewModelValidator()
	errs := validator.Validate(reflect.TypeOf(AppConfig{}), jsonMap)

	if len(errs) > 0 {
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf("- %s", e.Error()))
		}
		log.Fatalf("🚨 Configuration Validation Failed:\n%s", strings.Join(errMsgs, "\n"))
	}

	// 3. ASSIGN TO GLOBAL
	// Only after validation passes do we set the global variable.
	Cfg = &tempCfg
	log.Println("✅ Global configuration initialized and validated!")
}
