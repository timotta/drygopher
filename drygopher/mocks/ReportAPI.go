// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import pckg "github.com/eltorocorp/drygopher/drygopher/coverage/pckg"

// ReportAPI is an autogenerated mock type for the ReportAPI type
type ReportAPI struct {
	mock.Mock
}

// BuildCoverageReport provides a mock function with given fields: allPackages, exclusionPatterns
func (_m *ReportAPI) BuildCoverageReport(allPackages pckg.Group, exclusionPatterns []string) (string, error) {
	ret := _m.Called(allPackages, exclusionPatterns)

	var r0 string
	if rf, ok := ret.Get(0).(func(pckg.Group, []string) string); ok {
		r0 = rf(allPackages, exclusionPatterns)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pckg.Group, []string) error); ok {
		r1 = rf(allPackages, exclusionPatterns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}