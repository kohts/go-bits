package main

import (
	"fmt"
	"log"
	"pkohts/config"
	"reflect"
	"strings"
)

func printStruct(v interface{}, ind int) {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		fmt.Printf(strings.Repeat(" ", ind))
		if value.Kind() == reflect.Struct {
		} else {
			fmt.Printf("[%02d] %v %v == %v\n", i, field.Name, field.Tag, value.Interface())
		}

	}
}

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("can't read configuration: %v", err)
	}

	for i, a := range config.Assistants {
		fmt.Printf("%+v: %+v\n", i, a)
	}

	//printStruct(config, 0)
}
