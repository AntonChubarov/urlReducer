package httpcontrollers

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"server/domain"
	"strings"
	"testing"
)

func TestPostLinkController(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := []struct {
		description  string
		service      *domain.MockIService
		wantHTTPCode int
		wantError    error
	}{
		{
			description:  "should return statusOK",
			wantHTTPCode: http.StatusOK,
			service: func(mock *domain.MockIService) *domain.MockIService {
				mock.EXPECT().SaveLink(gomock.Any()).Return("", nil).Times(1)
				return mock
			}(domain.NewMockIService(ctrl)),
			wantError:   nil,
		},
	}

	for _, c := range cases {
		cc := c
		t.Run(c.description, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			req.Header.Add(echo.HeaderContentType, echo.MIMETextPlain)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handle := NewPostController(cc.service)
			// act
			c.SetPath("/")
			// assert
			if err := handle.Post(c); err == nil {
				assert.Equal(t, cc.wantHTTPCode, rec.Code)
			} else {
				assert.Error(t, cc.wantError, err)
			}
		})
	}
}