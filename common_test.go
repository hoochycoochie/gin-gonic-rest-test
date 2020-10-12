package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article
var tmpUserList []user

// This function is used for setup before executing the test functions

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func GetRouter(withTemplate bool) *gin.Engine {
	r := gin.Default()

	if withTemplate {
		r.LoadHTMLGlob("templates/*")
	}

	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpUserList = userList
	tmpArticleList = articleList
}

// This function is used to restore the main lists from the temporary one
func restoreLists() {
	userList = tmpUserList
	articleList = tmpArticleList
}
