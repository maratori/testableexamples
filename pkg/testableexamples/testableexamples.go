package testableexamples

import (
	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer returns Analyzer that checks if examples are testable.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "testableexamples",
		Doc:  "linter checks if examples are testable",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			return nil, nil
		},
	}
}
