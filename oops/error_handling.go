package oops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"github.com/pkg/errors"
)

const (
	pgxCode         = 1000
	jsonCode        = 2000
	internalCode    = 3000
	defaultCode     = 4000
	validationCode  = 5000
	grpcCode        = 6000
	timeParseError  = 7000
	httpRequestCode = 8000
)

// Error define um tipo erro para tratamento
type Error struct {
	Msg        string   `json:"msg"`
	Code       int      `json:"code"`
	Trace      []string `json:"-"`
	Err        error    `json:"-"`
	StatusCode int      `json:"-"`
}

// Error implementa a interface do tipo error
func (e *Error) Error() string {
	return e.Msg
}

// proccessError trata o erro para enviar uma mensagem para o usuário
func proccessError(rawError error) error {
	msg, code, responseStatus := "Erro desconhecido", 0, 400
	switch err := rawError.(type) {
	case *json.UnmarshalTypeError:
		msg, code = fmt.Sprintf("Tipo de valor %v não suportado no campo %v. Esperado tipo %v", err.Value, err.Field, err.Type.String()), jsonCode+1

	case validator.ValidationErrors:
		msg, code = "Campo "+err[0].Field()+" é obrigatorio", validationCode+1

	case *reflect.ValueError:
		msg, code = fmt.Sprintf("Não é possível acessar o valor do tipo %v", err.Kind.String()), internalCode+1

	case *strconv.NumError:
		msg, code = fmt.Sprintf("Não é possível converter valor %v", err.Num), internalCode+2

	case error:
		if err == io.EOF {
			msg, code = "Nenhum dado disponível para leitura", defaultCode+2
		}
	default:
		msg, code = err.Error(), defaultCode+1
	}

	return &Error{
		Msg:        msg,
		Err:        rawError,
		Code:       code,
		StatusCode: responseStatus,
	}
}

// Err constroi um instancia de erro a partir de um erro
func Err(err error) error {
	var e *Error
	if !errors.As(err, &e) || err == e {
		err = proccessError(err)
	}
	return errors.WithStack(err)
}

// Wrap encapsula o erro adicionando uma mensagem
func Wrap(err error, mensagem string) error {
	return errors.Wrap(Err(err), mensagem)
}

// DefinirErro adiciona uma mensagem e um status code para o erro
func DefinirErro(err error, c *gin.Context) {
	var e *Error

	if !errors.As(err, &e) {
		DefinirErro(Err(err), c)
		return
	}
	e.Msg = err.Error()

	c.JSON(e.StatusCode, e)
	c.Set("error", err)
	c.Abort()
}

// NovoErr cria uma nova instância de erro padrão
func NovoErr(mensagem string) error {
	return Err(&Error{
		Msg:        mensagem,
		Err:        errors.Errorf("Mensagem de erro interna: '%s'. Veja a stack para esse erro para ter informações adicionais.", mensagem),
		Code:       defaultCode,
		StatusCode: http.StatusBadRequest,
	})
}
