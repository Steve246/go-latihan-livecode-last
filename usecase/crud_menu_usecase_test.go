package usecase

import (
	"errors"
	"go_livecode_persiapan/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMenu = model.Menu{
	Id: "C001",
	Menu_Name: "Ayam Bakar",
}

type repoMock struct {
	mock.Mock
}

type MenuUseCaseTestSuite struct {
	suite.Suite
	repoMock *repoMock
}

func (r *repoMock) Update(menuFood *model.Menu, id string) error {
	args := r.Called(menuFood, id)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil

}

func (r *repoMock) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil

}

func (r *repoMock) Create(newMenu *model.Menu) error {
	args := r.Called(newMenu)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func(suite *MenuUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
}

func (suite *MenuUseCaseTestSuite) TestCustomerCreate_Success(){
	//Expected
	 newDumyMenu := model.Menu{}

	suite.repoMock.On("Create", newDumyMenu).Return(nil) //balikin nil, tapi diinput data customer baru

	//ACTUAL UJI CODE USECASE

	customerUsecaseTest := NewCrudMenuUsecase(suite.repoMock)

	err := customerUsecaseTest.CreateMenu(&newDumyMenu)
	//return nil 

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), nil, err) //bandingin nil data dummy dan error dari usecase asli

}

func (suite *MenuUseCaseTestSuite) TestCustomerCreate_Failed(){
	//Expected
	newDumyMenu := model.Menu{}

	suite.repoMock.On("Create", newDumyMenu).Return(errors.New("failed")) 

	//ACTUAL UJI CODE USECASE

	customerUsecaseTest := NewCrudMenuUsecase(suite.repoMock)

	err := customerUsecaseTest.CreateMenu(&newDumyMenu)
	

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)

}




func TestMenuRepositoryTestSuite(t *testing.T){
	suite.Run(t, new(MenuUseCaseTestSuite))
}

