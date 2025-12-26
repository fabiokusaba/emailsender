package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiokusaba/emailsender/internal/internalerrors"
	"github.com/stretchr/testify/assert"
)

func TestHandlerError_when_endpoint_returns_internal_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 500, internalerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func TestHandlerError_when_endpoint_returns_domain_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 400, errors.New("domain error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func TestHandlerError_when_endpoint_returns_obj_and_status(t *testing.T) {
	assert := assert.New(t)
	type BodyForTest struct {
		Id int
	}

	objExpected := BodyForTest{Id: 1}
	var objReturned BodyForTest

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 200, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpected, objReturned)
}
