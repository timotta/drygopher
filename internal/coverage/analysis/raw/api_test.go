package raw_test

import (
	"errors"
	"testing"

	"github.com/eltorocorp/drygopher/internal/pckg"

	"github.com/eltorocorp/drygopher/internal/coverage/analysis/raw"
	"github.com/eltorocorp/drygopher/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetRawCoverageAnalysisForPackage_Normally_ReturnsAnalysis(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)
	commandAPI := new(mocks.CommandAPI)

	commandAPI.On("Output").Return([]byte("testresult"), nil)
	execAPI.On("Command", mock.Anything, mock.Anything, mock.Anything).Return(commandAPI)

	profileData := "headerthatshouldgetskipped\ndata"
	osioAPI.On("ReadFile", mock.Anything).Return([]byte(profileData), nil)

	rawAPI := raw.New(osioAPI, execAPI)
	result, err := rawAPI.GetRawCoverageAnalysisForPackage("somepkg")
	assert.NoError(t, err)
	assert.Equal(t, []string{"data"}, result)
}

func Test_GetRawCoverageAnalysisForPackage_TestCmdError_ReturnsError(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)
	commandAPI := new(mocks.CommandAPI)

	commandAPI.On("Output").Return(nil, errors.New("test error"))
	execAPI.On("Command", mock.Anything, mock.Anything, mock.Anything).Return(commandAPI)

	rawAPI := raw.New(osioAPI, execAPI)
	_, err := rawAPI.GetRawCoverageAnalysisForPackage("somepkg")

	assert.EqualError(t, err, "test error")
}

func Test_GetRawCoverageAnalysisForPackage_NoTestsForPackage_ReturnsEmptySlice(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)
	commandAPI := new(mocks.CommandAPI)

	// go test will output a line that starts with a question mark for packages
	// that have no test files.
	commandAPI.On("Output").Return([]byte("?"), nil)
	execAPI.On("Command", mock.Anything, mock.Anything, mock.Anything).Return(commandAPI)

	rawAPI := raw.New(osioAPI, execAPI)
	result, err := rawAPI.GetRawCoverageAnalysisForPackage("somepkg")
	assert.NoError(t, err)
	assert.Equal(t, []string{}, result)
}

func Test_GetRawCoverageAnalysisForPackage_ErrorReadingTmpFile_ReturnsError(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)
	commandAPI := new(mocks.CommandAPI)

	commandAPI.On("Output").Return([]byte("testresult"), nil)
	execAPI.On("Command", mock.Anything, mock.Anything, mock.Anything).Return(commandAPI)
	osioAPI.On("ReadFile", mock.Anything).Return(nil, errors.New("test error"))

	rawAPI := raw.New(osioAPI, execAPI)
	_, err := rawAPI.GetRawCoverageAnalysisForPackage("somepkg")

	assert.EqualError(t, err, "test error")
}

func Test_AggregateRawPackageAnalysisData_Normally_AggregatesWithoutError(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)

	rawAPI := raw.New(osioAPI, execAPI)
	rawData := []string{
		"mode: count",
		"somepackage:0.0,0.0 1 0",
	}
	stats, err := rawAPI.AggregateRawPackageAnalysisData("somepackage", rawData)
	expectedStats := pckg.Stats{
		Covered:         0,
		Estimated:       false,
		Package:         "somepackage",
		RawCoverageData: rawData,
		Statements:      1,
		Uncovered:       1,
	}
	assert.Equal(t, expectedStats, *stats)
	assert.NoError(t, err)
}

func Test_AggregateRawPackageAnalysisData_ErrorParsing_ReturnsError(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)
	rawAPI := raw.New(osioAPI, execAPI)

	testCases := []struct {
		name    string
		rawData []string
	}{
		{name: "badCallCount", rawData: []string{"mode: count", "somepackage:0.0,0.0 0 X"}},
		{name: "badStmtCount", rawData: []string{"mode: count", "somepackage:0.0,0.0 X 0"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stats, err := rawAPI.AggregateRawPackageAnalysisData("somepackage", tc.rawData)
			assert.Nil(t, stats)
			assert.EqualError(t, err, "strconv.ParseFloat: parsing \"X\": invalid syntax")
		})
	}
}

// See notes in the AggregateRawPackageAnalysisData method regarding this behavior.
func Test_AggregateRawPackageAnalysisData_CallsPresent_TotalCoveredEqualsCallCount(t *testing.T) {
	osioAPI := new(mocks.OSIOAPI)
	execAPI := new(mocks.ExecAPI)

	rawAPI := raw.New(osioAPI, execAPI)
	rawData := []string{
		"mode: count",
		"somepackage:0.0,0.0 5 1",
	}
	stats, err := rawAPI.AggregateRawPackageAnalysisData("somepackage", rawData)
	expectedStats := pckg.Stats{
		Covered:         5,
		Estimated:       false,
		Package:         "somepackage",
		RawCoverageData: rawData,
		Statements:      5,
		Uncovered:       0,
	}
	assert.Equal(t, expectedStats, *stats)
	assert.NoError(t, err)
}
