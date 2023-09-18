package response

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Renos-id/go-starter-template/lib"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func SetLogging(logg *logrus.Logger) {
	log = logg
}

type CommonResponse struct {
	Status      bool   `json:"status"`
	Message     string `json:"message"`
	Note        string `json:"-"`
	Data        any    `json:"data,omitempty"`
	Error       any    `json:"error,omitempty"`
	Code        int    `json:"-"`
	SendToSlack bool   `json:"-"`
}

// respondwithJSON write json response format
func (cr CommonResponse) ToJSON(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(cr)
	if IsApiFailed(cr.Code) {
		if cr.SendToSlack {
			sendToSlack(response, r, cr.Code, cr.Note)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(cr.Code)
	w.Write(response)
	return
}

func sendToSlack(crb []byte, r *http.Request, code int, note string) {
	mapData := make(logrus.Fields)
	json.Unmarshal(crb, &mapData)
	log.WithField("App Name", os.Getenv("APP_NAME")).
		WithField("Method", r.Method).
		WithField("Request URL", r.URL).
		WithField("Code", code).
		WithFields(mapData).
		WithField("HTTPBodyRequest", lib.StructToJSON(r.Context().Value("body"))).
		Error(note)
}

func WriteSuccess(message string, data any) CommonResponse {
	return CommonResponse{
		Status:      true,
		Message:     message,
		Data:        data,
		Code:        200,
		SendToSlack: false,
	}
}

func WriteError(code int, note string, sendToSlack bool, err any) CommonResponse {
	// var errors ValidationErrors
	if code == 0 {
		code = 500
	}
	var body = CommonResponse{
		Status:      false,
		Message:     http.StatusText(code),
		Note:        note,
		Code:        code,
		SendToSlack: sendToSlack,
	}

	switch e := err.(type) {
	case ValidationErrors:
		body.Error = err
	case error:
		switch e.(error) {
		case sql.ErrNoRows:
			body.Message = "Data does not exists!"
		case io.EOF:
			body.Error = "Failed to read HTTP Request"
		default:
			body.Error = e.Error()
		}
	case map[string]string:
		body.Data = err
	default:
		body.Error = err
	}

	return body
}
