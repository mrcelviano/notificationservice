package commons

import (
	"errors"
	"flag"
	"os"
	"strings"
)

var environmentName string

const (
	// ConArgEnv Название ключа аргумента консоли, в котором будет указываться названия среды окружения
	ConArgEnv             = "env"
	GlobalEnvVariableName = "SOCIAL_TECH_ENV"
)

var (
	// ErrEmptyEnvName - ошибка, выдаваемая в том случае, если приложение запущено без указания среды
	ErrEmptyEnvName = errors.New("environment name is empty")
	// ErrInvalidEnvName - ошибка, выдаваемая в том случае, если приложение запущено с указанием неизвестной среды
	ErrInvalidEnvName = errors.New("invalid environment name")
)

const (
	// EnvNameLocal Название локального окружения
	EnvNameLocal = "local"
	// для тестов
	EnvNameTest = "test"
)

func checkEnvName(envName string) bool {
	return envName == EnvNameLocal || envName == EnvNameTest
}

//GetEnvFromArgs Получить среду из аргументов запуска
func GetEnvFromArgs() (string, bool) {
	var env string
	flag.StringVar(&env, ConArgEnv, "", "environment mode")
	flag.Parse()
	if len(env) > 0 {
		return env, true
	}
	return env, false
}

//GetEnvFromGlobal Получить среду из среды ОС
func GetEnvFromGlobal() (string, bool) {
	var env string
	env = os.Getenv(GlobalEnvVariableName)
	if len(env) > 0 {
		return env, true
	}
	return env, false
}

// GetEnvVar Получить имя среды (считывает либо из параметров запуска, либо из глобальных )
func GetEnvVar() string {
	if len(environmentName) == 0 {
		env, isProvide := GetEnvFromArgs()
		if isProvide == false {
			env, isProvide = GetEnvFromGlobal()
			if isProvide == false {
				panic(ErrEmptyEnvName)
			}

		}
		environmentName = strings.ToLower(env)
		if !checkEnvName(environmentName) {
			panic(ErrInvalidEnvName)
		}

	}
	return environmentName
}
