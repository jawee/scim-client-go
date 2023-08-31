package internal

import (
    "reflect"
	"fmt"
	"strconv"
)

func MainFunc() {
    fmt.Printf("Hello internal\n")
    asdf()
}

type FooBar struct {
	Foo string
	Bar int
}

func asdf() {
	structValue := FooBar{Foo: "foo", Bar: 10}
	fields := reflect.TypeOf(structValue)
	values := reflect.ValueOf(structValue)

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")

		switch field.Type.Kind() {
		case reflect.String:
			v := value.String()
			fmt.Print(v, "\n")
		case reflect.Int:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		case reflect.Int32:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		case reflect.Int64:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		default:
			panic("Not support type of struct")
		}
	}
}

