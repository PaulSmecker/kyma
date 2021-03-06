// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import apperrors "github.com/kyma-project/kyma/components/metadata-service/internal/apperrors"

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Get provides a mock function with given fields: bucketName, objectName
func (_m *Repository) Get(bucketName string, objectName string) ([]byte, apperrors.AppError) {
	ret := _m.Called(bucketName, objectName)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(bucketName, objectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(bucketName, objectName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Put provides a mock function with given fields: bucketName, objectName, resource
func (_m *Repository) Put(bucketName string, objectName string, resource []byte) apperrors.AppError {
	ret := _m.Called(bucketName, objectName, resource)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string, []byte) apperrors.AppError); ok {
		r0 = rf(bucketName, objectName, resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: bucketName, objectName
func (_m *Repository) Remove(bucketName string, objectName string) apperrors.AppError {
	ret := _m.Called(bucketName, objectName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(bucketName, objectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
