package main

import (
	"github.com/gcpug/zagane/zagane"
	"github.com/gostaticanalysis/ctxfield"
	"github.com/gostaticanalysis/dupimport"
	"github.com/gostaticanalysis/exclude"
	"github.com/gostaticanalysis/forcetypeassert"
	"github.com/gostaticanalysis/importgroup"
	"github.com/gostaticanalysis/loopdefer"
	"github.com/gostaticanalysis/nilerr"
	"github.com/gostaticanalysis/nofmt"
	"github.com/gostaticanalysis/notest"
	"github.com/gostaticanalysis/sqlrows/passes/sqlrows"
	"github.com/gostaticanalysis/typeswitch"
	"github.com/gostaticanalysis/unitconst"
	"github.com/gostaticanalysis/unused"
	"github.com/gostaticanalysis/wraperrfmt"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(exclude.All(append(
		zagane.Analyzers(), // for Cloud Spanner
		ctxfield.Analyzer,
		dupimport.Analyzer,
		exclude.By(forcetypeassert.Analyzer, exclude.TestFile),
		importgroup.Analyzer,
		loopdefer.Analyzer,
		nilerr.Analyzer,
		nofmt.Analyzer,
		notest.Analyzer,
		sqlrows.Analyzer,
		typeswitch.Analyzer,
		unitconst.Analyzer,
		unused.Analyzer,
		wraperrfmt.Analyzer,
	), exclude.GeneratedFile, exclude.FileWithPattern)...)
}
