package godocs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/seletskiy/hierr"
)

type ErrorBind struct {
	Arg      string
	Err      error
	NotValid bool
	NotFound bool
}

func (bind ErrorBind) Error() string {
	if bind.NotValid {
		return hierr.Errorf(
			bind.Err, "%s has invalid value", bind.Arg,
		).Error()
	}

	return fmt.Sprintf("%s is not specified", bind.Arg)
}

func Bind(
	doc, version string,
	receiver interface{}, settings ...interface{},
) error {
	args, err := Parse(doc, version, settings...)
	fmt.Printf("XXXXXX bind.go:32: args: %#v\n", args)
	if err != nil {
		return err
	}

	return bind(receiver, args)
}

func bind(
	value interface{},
	args map[string]interface{},
) error {
	resource := reflect.Indirect(reflect.ValueOf(value))
	if resource.Kind() == reflect.Map {
		value = args
		return nil
	}

	if resource.Kind() != reflect.Struct {
		panic("resource bind must be a struct")
	}

	resourceStruct := resource.Type()
	for index := 0; index < resourceStruct.NumField(); index++ {
		var (
			resourceField = resource.Field(index)
			structField   = resourceStruct.Field(index)
		)

		key := structField.Tag.Get("arg")
		if key != "" {
			arg, ok := args[key]
			if !ok {
				return ErrorBind{
					Arg:      key,
					NotFound: true,
				}
			}

			fieldValue, err := convert(
				reflect.ValueOf(arg),
				structField.Type,
			)
			if err != nil {
				return ErrorBind{
					Arg:      key,
					NotValid: true,
					Err:      err,
				}
			}

			err = set(resourceField, fieldValue)
			if err != nil {
				return ErrorBind{
					Arg:      key,
					NotValid: true,
					Err:      err,
				}
			}
		}

		for resourceField.Kind() == reflect.Ptr {
			resourceField = resourceField.Elem()
		}

		if resourceField.Kind() == reflect.Struct {
			err := bind(resourceField.Addr().Interface(), args)
			if err != nil {
				return err
			}

			continue
		}

		if resourceField.Kind() == reflect.Slice {
			for i := 0; i < resourceField.Len(); i++ {
				field := reflect.Indirect(resourceField.Index(i))
				if field.Kind() == reflect.Struct {
					err := bind(field.Addr().Interface(), args)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func convert(
	source reflect.Value,
	_type reflect.Type,
) (value reflect.Value, err error) {
	defer func() {
		tears := recover()
		if tears != nil {
			err = errors.New(fmt.Sprint(tears))
		}
	}()

	return source.Convert(_type), nil
}

func set(source reflect.Value, value reflect.Value) (err error) {
	defer func() {
		tears := recover()
		if tears != nil {
			err = errors.New(fmt.Sprint(tears))
		}
	}()

	source.Set(value)
	return nil
}
