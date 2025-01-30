package controller

import (
	"errors"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/gomdori-foo/bearden/internal/core/common"
	"github.com/gomdori-foo/bearden/internal/core/provider"
	"github.com/gomdori-foo/bearden/internal/core/utils"
)

type ControllerFactory struct {
	constructor interface{}
	options ControllerFactoryOptions
}

type ControllerFactoryOptions struct {
	scope common.Scope
}

type ControllerFactoryWith struct {
	options ControllerFactoryOptions
}

type ControllerMetadata struct {
	prefix string
	routes []Route
}

func Controllers(constructors ...interface{}) []*ControllerFactory {
	controllerFactories := make([]*ControllerFactory, len(constructors))

	for i, constructor := range constructors {
		if controllerFactory, ok := constructor.(*ControllerFactory); ok {
			controllerFactories[i] = controllerFactory
		} else if utils.IsConstructor(constructor) {
			controllerFactories[i] = &ControllerFactory{
				constructor: constructor,
				options: NewDefaultOptions(),
			}
		} else {
			panic("invalid controller constructor")
		}
	}

	return controllerFactories
}

func Use(constructor interface{}, args ...*ControllerFactoryWith) *ControllerFactory {
	if !utils.IsConstructor(constructor) {
		panic("invalid controller constructor")
	}

	options := NewDefaultOptions()

	if len(args) > 0 {
		with := args[0]
		options.scope = with.options.scope
	}

	return &ControllerFactory{
		constructor: constructor,
		options: options,
	}
}

func With(options ControllerFactoryOptions) *ControllerFactoryWith {
	return &ControllerFactoryWith{
		options: options,
	}
}

func NewDefaultOptions() ControllerFactoryOptions {
	return ControllerFactoryOptions{
		scope: common.ScopeDefault,
	}
}

func (c *ControllerFactory) Create(providers []*provider.Provider) *Controller {
	constructorType := reflect.TypeOf(c.constructor)
	if constructorType.Kind() != reflect.Func {
		panic("controller constructor is not a function")
	}

	params := make([]interface{}, constructorType.NumIn())
	for i := 0; i < len(params); i++ {
		paramType := constructorType.In(i)
		provider := provider.FindProvider(providers, paramType)
		if provider == nil {
			return nil
		}

		instance := provider.Instance()
		instanceType := reflect.TypeOf(instance)
		if paramType.Kind() == reflect.Ptr && instanceType.Kind() != reflect.Ptr {
			val := reflect.New(instanceType)
			val.Elem().Set(reflect.ValueOf(instance))
			instance = val.Interface()
		} else if paramType.Kind() != reflect.Ptr && instanceType.Kind() == reflect.Ptr {
			instance = reflect.ValueOf(instance).Elem().Interface()
		}

		params[i] = instance
	}

	constructor := func() interface{} {
		constructorValue := reflect.ValueOf(c.constructor)
		paramValues := make([]reflect.Value, len(params))
		for i, param := range params {
			paramValues[i] = reflect.ValueOf(param)
		}
		result := constructorValue.Call(paramValues)[0].Interface()
		return result
	}

	metadata := c.getControllerMetadata()

	controller := NewController(metadata.prefix, metadata.routes, constructor)

	return controller
}

func (c *ControllerFactory) getControllerMetadata() ControllerMetadata {
	pc := reflect.ValueOf(c.constructor).Pointer()
	file, _ := runtime.FuncForPC(pc).FileLine(pc)
	fileContent, err := os.ReadFile(file)
	if err != nil {
		panic("failed to read file " + file)
	}

	content := string(fileContent)
	lines := strings.Split(content, "\n")

	prefix, err := getControllerPrefix(lines)
	if err != nil {
		panic("failed to get controller prefix")
	}

	routes, err := getControllerRoutes(lines)
	if err != nil {
		panic("failed to get controller routes")
	}

	return ControllerMetadata{
		prefix: prefix,
		routes: routes,
	}
}

func getControllerPrefix(lines []string) (string, error) {
	for _, line := range lines {
		if strings.Contains(line, "@Controller") {
			// Extract the controller path from annotation
			// Example: @Controller() -> "/"
			// Example: @Controller("") -> "/"
			// Example: @Controller("/") -> "/"
			// Example: @Controller("foo") -> "/foo"
			// Example: @Controller("/foo") -> "/foo"
			// Example: @Controller("/foo/") -> "/foo"
			pathMatch := regexp.MustCompile(`@Controller\((.*?)\)`).FindStringSubmatch(line)
			if len(pathMatch) > 1 && pathMatch[1] != "" {
				// Extract path from quotes if present
				pathInQuotes := regexp.MustCompile(`"([^"]+)"`).FindStringSubmatch(pathMatch[1])
				if len(pathInQuotes) > 1 {
					if strings.HasPrefix(pathInQuotes[1], "/") {
						return pathInQuotes[1], nil
					} else if (strings.HasSuffix(pathInQuotes[1], "/")) {
						return pathInQuotes[1][:len(pathInQuotes[1])-1], nil
					} else {
						return "/" + pathInQuotes[1], nil
					}
				} else {
					return "/", nil
				}
			} else {
				return "/", nil
			}
		}
	}
	return "", errors.New("no controller prefix found")
}

func getControllerRoutes(lines []string) ([]Route, error) {
	routes := make([]Route, 0)

	var err error
	var method string = ""
	var path string = ""
	var handlerName string = ""

	for _, line := range lines {
		if strings.Contains(line, "@Get") {
			method = "Get"
			path, err = getControllerPath(`@Get\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "@Post") {
			method = "Post"
			path, err = getControllerPath(`@Post\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "@Put") {
			method = "Put"
			path, err = getControllerPath(`@Put\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "@Delete") {
			method = "Delete"
			path, err = getControllerPath(`@Delete\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "@Patch") {
			method = "Patch"
			path, err = getControllerPath(`@Patch\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "@Options") {
			method = "Options"
			path, err = getControllerPath(`@Options\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "@Head") {
			method = "Head"
			path, err = getControllerPath(`@Head\((.*?)\)`, line)
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(line, "func") {
			handlerName = extractMethodName(line)

			if method != "" {
				route := NewRoute(method, path, handlerName)
				routes = append(routes, route)
			}

			method = ""
			path = ""
			handlerName = ""
		}
	}

	return routes, nil
}

func getControllerPath(regex string, line string) (string, error) {
	pathMatch := regexp.MustCompile(regex).FindStringSubmatch(line)
	if len(pathMatch) > 1 && pathMatch[1] != "" {
		// Extract the method name from annotation
			// Example: @Get() -> ""
			// Example: @Get("") -> ""
			// Example: @Get("/") -> ""
			// Example: @Get("bar") -> "/bar"
			// Example: @Get("/bar") -> "/bar"
			// Example: @Get("/bar/") -> "/bar"
		pathInQuotes := regexp.MustCompile(`"([^"]+)"`).FindStringSubmatch(pathMatch[1])
		if len(pathInQuotes) > 1 {
			if strings.HasPrefix(pathInQuotes[1], "/") {
				return pathInQuotes[1], nil
			} else if (strings.HasSuffix(pathInQuotes[1], "/")) {
				return pathInQuotes[1][:len(pathInQuotes[1])-1], nil
			} else {
				return pathInQuotes[1], nil
			}
		} else {
			return "", nil
		}
	} else {
		return "", nil
	}
}

func extractMethodName(methodString string) string {
	// "func (c *FileController) Upload(ctx *gin.Context) {"

	parts := strings.Split(methodString, ") ")
	if len(parts) < 2 {
		return ""
	}
	
	methodName := strings.Split(parts[1], "(")[0]
	
	return methodName
}