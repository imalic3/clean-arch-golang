// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	types "clean-arch-golang/types"

	mock "github.com/stretchr/testify/mock"
)

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

// Articles provides a mock function with given fields:
func (_m *ArticleRepository) Articles() ([]types.Article, error) {
	ret := _m.Called()

	var r0 []types.Article
	if rf, ok := ret.Get(0).(func() []types.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: title
func (_m *ArticleRepository) Get(title string) (types.Article, error) {
	ret := _m.Called(title)

	var r0 types.Article
	if rf, ok := ret.Get(0).(func(string) types.Article); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(types.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NumberOfArticles provides a mock function with given fields:
func (_m *ArticleRepository) NumberOfArticles() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Store provides a mock function with given fields: article
func (_m *ArticleRepository) Store(article types.Article) error {
	ret := _m.Called(article)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Article) error); ok {
		r0 = rf(article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewArticleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewArticleRepository creates a new instance of ArticleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArticleRepository(t mockConstructorTestingTNewArticleRepository) *ArticleRepository {
	mock := &ArticleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
