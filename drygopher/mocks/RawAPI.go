// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import pckg "github.com/eltorocorp/drygopher/drygopher/coverage/pckg"

// RawAPI is an autogenerated mock type for the RawAPI type
type RawAPI struct {
	mock.Mock
}

// AggregateRawPackageAnalysisData provides a mock function with given fields: pkg, rawPkgCoverageData
func (_m *RawAPI) AggregateRawPackageAnalysisData(pkg string, rawPkgCoverageData []string) (*pckg.Stats, error) {
	ret := _m.Called(pkg, rawPkgCoverageData)

	var r0 *pckg.Stats
	if rf, ok := ret.Get(0).(func(string, []string) *pckg.Stats); ok {
		r0 = rf(pkg, rawPkgCoverageData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pckg.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(pkg, rawPkgCoverageData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRawCoverageAnalysisForPackage provides a mock function with given fields: pkg
func (_m *RawAPI) GetRawCoverageAnalysisForPackage(pkg string) ([]string, error) {
	ret := _m.Called(pkg)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(pkg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pkg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}