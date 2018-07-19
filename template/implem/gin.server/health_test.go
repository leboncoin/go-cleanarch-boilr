package server

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/h2non/baloo.v3"

	"{{GitServer}}/{{Organization}}/{{Name}}/implem/mock.uc"
)

func TestHealth_happycase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ucHandler := uc.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().Health().Return(nil).Times(1)


	gE := gin.Default()
	NewRouter(ucHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	baloo.New(ts.URL).
		Get("/health").
		Expect(t).
		Status(http.StatusOK).
		Done()
}

func TestHealth_fail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ucHandler := uc.NewMockHandler(mockCtrl)
	ucHandler.EXPECT().Health().Return(errors.New("")).Times(1)


	gE := gin.Default()
	NewRouter(ucHandler).SetRoutes(gE)

	ts := httptest.NewServer(gE)
	defer ts.Close()

	baloo.New(ts.URL).
		Get("/health").
		Expect(t).
		Status(http.StatusInternalServerError).
		Done()
}