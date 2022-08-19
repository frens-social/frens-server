package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database Database `mapstructure:"database"`
	Router   Router   `mapstructure:"router"`
}

type Database struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
}

type Router struct {
	Port            string `mapstructure:"port" validate:"required"`
	Prefork         bool   `mapstructure:"prefork"`
	BodyLimit       int    `mapstructure:"body_limit"`
	ReadBufferSize  int    `mapstructure:"read_buffer_size"`
	WriteBufferSize int    `mapstructure:"write_buffer_size"`
	ReadTimeout     int    `mapstructure:"read_timeout"`
	WriteTimeout    int    `mapstructure:"write_timeout"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
}

var C Configuration
var FiberConfig fiber.Config

func (c *Configuration) Validate() error {
	v := validator.New()

	var errs []*ErrorResponse
	err := v.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, &ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}
	}

	for _, err := range errs {
		logrus.Errorf("Invalid configuration: %s %s %v", err.FailedField, err.Tag, err.Value)
	}

	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("%v", errs))
	} else {
		return nil
	}
}

func Load() {
	vp := viper.New()
	vp.SetConfigName("conf") // name of config file (without extension)
	vp.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath(".")    // optionally look for config in the working directory

	err := vp.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(err)
	}

	err = vp.Unmarshal(&C)
	if err != nil {
		panic(err)
	}

	// Validate configuration
	if err := C.Validate(); err != nil {
		logrus.Fatal("Invalid configuration.")
	}

	// create router config
	FiberConfig = fiber.Config{
		Prefork:         C.Router.Prefork,
		BodyLimit:       C.Router.BodyLimit,
		ReadBufferSize:  C.Router.ReadBufferSize,
		WriteBufferSize: C.Router.WriteBufferSize,
		ReadTimeout:     time.Second * time.Duration(C.Router.ReadTimeout),
		WriteTimeout:    time.Second * time.Duration(C.Router.WriteTimeout),
	}

	//todo, add fsnotify to watch for changes
	logrus.Info("Configuration successfully loaded.")
}
