package numtostr_test

import (
	"testing"

	"github.com/gostaticanalysis/numtostr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, numtostr.Analyzer, "a")
}

func TestPackageAlias(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, numtostr.Analyzer, "b")
}
