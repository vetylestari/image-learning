package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/Renos-id/go-starter-template/lib/response"
	"github.com/ggicci/httpin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	v *validator.Validate
)

// HTTP middleware setting a value on the request context
func RequestValidation[K interface{}](data K, v *validator.Validate, trans ut.Translator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			input, b := r.Context().Value(httpin.Input), new(bytes.Buffer)
			json.NewEncoder(b).Encode(input)
			decoder := json.NewDecoder(b)
			err := decoder.Decode(&data)
			if err != nil {
				resp := response.WriteError(500, "Failed Decoed HTTP Request", true, err)
				resp.ToJSON(w, r)
				return
			}

			//Check If Header Exists
			metaValue := reflect.ValueOf(&data).Elem()
			rns_user_id_header := metaValue.FieldByName("RnsUserId")
			rns_user_name_header := metaValue.FieldByName("RnsUserName")
			if rns_user_id_header != (reflect.Value{}) {
				user_id, _ := strconv.ParseInt(r.Header.Get("x-rns-user-id"), 10, 0)
				rns_user_id_header.SetInt(user_id)
			}
			if rns_user_name_header != (reflect.Value{}) {
				rns_user_name_header.SetString(r.Header.Get("x-rns-user-name"))
			}

			if err != nil {
				resp := response.WriteError(500, "Failed decode Request in Request Validation", true, err)
				resp.ToJSON(w, r)
				return
			}
			err = v.Struct(data)
			if err != nil {
				errs := translateError(err, trans)
				resp := response.WriteError(422, "Validation Failed", false, errs)
				resp.ToJSON(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), "body", data)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func translateError(err error, trans ut.Translator) (errs response.ValidationErrors) {
	var errors response.ValidationErrors
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Sprint(e.Translate(trans))
		errs = append(errors, response.ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: strings.Replace(translatedErr, "_", " ", 1),
		})
	}
	return errs
}
