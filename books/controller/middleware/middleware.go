package middleware

import (
	"bytes"
	"io"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
	"github.com/valyala/fasttemplate"
)

type GoMiddleware struct {
}

func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

type (
	Skipper      func(c echo.Context) bool
	LoggerConfig struct {
		Skipper  Skipper
		Format   string `yaml:"format"`
		Output   io.Writer
		template *fasttemplate.Template
		colorer  *color.Color
		pool     *sync.Pool
	}
)

func DefaultSkipper(echo.Context) bool {
	return false
}

var (
	DefaultLoggerConfig = LoggerConfig{
		Skipper: DefaultSkipper,
		Format:  `"method":"${method}","uri":"${uri}"  `,
	}
)

func (m *GoMiddleware) Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultLoggerConfig.Skipper
	}
	if config.Format == "" {
		config.Format = DefaultLoggerConfig.Format
	}
	if config.Output == nil {
		config.Output = DefaultLoggerConfig.Output
	}

	config.template = fasttemplate.New(config.Format, "${", "}")
	config.colorer = color.New()
	config.colorer.SetOutput(config.Output)
	config.pool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 256))
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			if err = next(c); err != nil {
				c.Error(err)
			}
			buf := config.pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer config.pool.Put(buf)

			if _, err = config.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "uri":
					return buf.WriteString(req.RequestURI)
				case "method":
					return buf.WriteString(req.Method)
				}
				return 0, nil
			}); err != nil {
				return
			}

			if config.Output == nil {
				_, err = c.Logger().Output().Write(buf.Bytes())
				return
			}
			_, err = config.Output.Write(buf.Bytes())
			return
		}
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
