package plugin

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Redirect struct {
		Base `json:",squash"`
		To   string `json:"to"`
		Code string `json:"code"`
		When string `json:"when"`
	}

	HTTPSRedirect struct {
		Base                      `json:",squash"`
		middleware.RedirectConfig `json:",squash"`
	}

	HTTPSWWWRedirect struct {
		Base                      `json:",squash"`
		middleware.RedirectConfig `json:",squash"`
	}

	HTTPSNonWWWRedirect struct {
		Base                      `json:",squash"`
		middleware.RedirectConfig `json:",squash"`
	}

	WWWRedirect struct {
		Base                      `json:",squash"`
		middleware.RedirectConfig `json:",squash"`
	}

	NonWWWRedirect struct {
		Base                      `json:",squash"`
		middleware.RedirectConfig `json:",squash"`
	}
)

func (*Redirect) Init() (err error) {
	return
}

func (*Redirect) Priority() int {
	return -1
}

func (r *Redirect) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return r.Middleware(next)
}

func (r *HTTPSRedirect) Init() (err error) {
	r.Middleware = middleware.HTTPSRedirectWithConfig(r.RedirectConfig)
	return
}

func (*HTTPSRedirect) Priority() int {
	return -1
}

func (r *HTTPSRedirect) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return r.Middleware(next)
}

func (r *HTTPSWWWRedirect) Init() (err error) {
	r.Middleware = middleware.HTTPSWWWRedirectWithConfig(r.RedirectConfig)
	return
}

func (*HTTPSWWWRedirect) Priority() int {
	return -1
}

func (r *HTTPSWWWRedirect) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return r.Middleware(next)
}

func (r *HTTPSNonWWWRedirect) Init() (err error) {
	r.Middleware = middleware.HTTPSNonWWWRedirectWithConfig(r.RedirectConfig)
	return
}

func (*HTTPSNonWWWRedirect) Priority() int {
	return -1
}

func (r *HTTPSNonWWWRedirect) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return r.Middleware(next)
}

func (r *WWWRedirect) Init() (err error) {
	r.Middleware = middleware.WWWRedirectWithConfig(r.RedirectConfig)
	return
}

func (*WWWRedirect) Priority() int {
	return -1
}

func (r *WWWRedirect) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return r.Middleware(next)
}

func (r *NonWWWRedirect) Init() (err error) {
	r.Middleware = middleware.NonWWWRedirectWithConfig(r.RedirectConfig)
	return
}

func (*NonWWWRedirect) Priority() int {
	return -1
}

func (r *NonWWWRedirect) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return r.Middleware(next)
}
