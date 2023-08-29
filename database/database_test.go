package database

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"user_management_service/models"
	"user_management_service/testhelpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  *Database
	ctx         context.Context
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	db, err := sql.Open("postgres", pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewDatabase(db)
	suite.repository = repository
}
func (suite *CustomerRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *CustomerRepoTestSuite) TestCreateCustomer() {
	t := suite.T()

	err := suite.repository.CreateUser(&models.User{
		Name:  "Test",
		Email: "test@test",
	})
	assert.NoError(t, err)
}

func (suite *CustomerRepoTestSuite) TestGetCustomerByEmail() {
	t := suite.T()

	customer, err := suite.repository.GetUserByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Test", customer.Name)
	assert.Equal(t, "test@test", customer.Email)
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
