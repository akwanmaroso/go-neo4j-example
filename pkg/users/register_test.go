package users_test

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/akwanmaroso/go-neo4j-example/pkg/users"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeUserRepository struct {
}

func (f FakeUserRepository) RegisterUser(user users.User) error {
	return nil
}

var _ = Describe("Users", func() {
	It("should register", func() {
		handler := users.UserHandler{
			Path:           "/users",
			UserRepository: FakeUserRepository{},
		}
		testResponseWritter := httptest.NewRecorder()
		requestBody := strings.NewReader("{\"user\":{\"email\":\"user@example.com\", \"password\":\"secret\", \"username\":\"user\"}}")
		handler.Register(testResponseWritter, httptest.NewRequest("POST", "/users", requestBody))

		Expect(testResponseWritter.Code).To(Equal(201))
		responseBody, _ := ioutil.ReadAll(testResponseWritter.Body)
		Expect(string(responseBody)).To(Equal("{\"user\":{\"username\":\"user\",\"email\":\"user@example.com\"}}"))
	})
})
