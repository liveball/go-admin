package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
)

//JSON def json struct
type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//Render writes data with custom ContentType for implement gin Render interface
func (js JSON) Render(w http.ResponseWriter) error {
	var jb []byte
	if jb, err = json.Marshal(js); err != nil {
		return errors.WithStack(err)
	}
	if _, err = w.Write(jb); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// WriteContentType write json ContentType for implement gin Render interface
func (js JSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
}

func myRender(c *gin.Context, code int, r render.Render) {
	r.WriteContentType(c.Writer)
	if code > 0 {
		c.Status(code)
	}

	if err := r.Render(c.Writer); err != nil {
		c.Error(err)
	}
}

// JSONOutput def json output
func JSONOutput(c *gin.Context, data interface{}, err error) {
	code := http.StatusOK
	msg := "ok"
	if err != nil {
		code = 500
		msg = err.Error()
	}

	myRender(c, http.StatusOK, JSON{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
