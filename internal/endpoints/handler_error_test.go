package endpoints

import (
	internalerrors "campaign-manager/internal/internalErrors"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_Returning_Object(t *testing.T) {
	assert := assert.New(t)
	campaignId := "3k12jl3k12j3"
	mockedEndpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return map[string]string{
			"id": campaignId,
		}, 200, nil
	}

	handleFunc := HandlerError(mockedEndpoint)
	req, _ := http.NewRequest("GET", "", nil)
	res := httptest.NewRecorder()

	handleFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Contains(res.Body.String(), campaignId)
}

func Test_HandlerError_With_InternalServerError(t *testing.T) {
	assert := assert.New(t)
	mockedEndpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternalError
	}

	handleFunc := HandlerError(mockedEndpoint)
	req, _ := http.NewRequest("POST", "", nil)
	res := httptest.NewRecorder()

	handleFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternalError.Error())
}

func Test_HandlerError_With_ValidationError(t *testing.T) {
	assert := assert.New(t)
	errorMessage := "validation error"
	mockedEndpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New(errorMessage)
	}

	handleFunc := HandlerError(mockedEndpoint)
	req, _ := http.NewRequest("POST", "", nil)
	res := httptest.NewRecorder()

	handleFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), errorMessage)
}

func Test_HandlerError_With_ResourceNotFound(t *testing.T) {
	assert := assert.New(t)
	mockedEndpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrResourceNotFound
	}

	handleFunc := HandlerError(mockedEndpoint)
	req, _ := http.NewRequest("POST", "", nil)
	res := httptest.NewRecorder()

	handleFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusNotFound, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrResourceNotFound.Error())
}
