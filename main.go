package main

import (
	"github.com/gcpug/zagane/zagane"
	"github.com/gostaticanalysis/ctxfield"
	"github.com/gostaticanalysis/debugcode/debugcode"
	"github.com/gostaticanalysis/dupimport"
	"github.com/gostaticanalysis/elseless"
	"github.com/gostaticanalysis/exclude"
	"github.com/gostaticanalysis/forcetypeassert"
	"github.com/gostaticanalysis/importgroup"
	"github.com/gostaticanalysis/innertypealias"
	"github.com/gostaticanalysis/loopdefer"
	"github.com/gostaticanalysis/nilerr"
	"github.com/gostaticanalysis/noctor"
	"github.com/gostaticanalysis/nofmt"
	"github.com/gostaticanalysis/notest"
	"github.com/gostaticanalysis/sqlrows/passes/sqlrows"
	"github.com/gostaticanalysis/testhelper"
	"github.com/gostaticanalysis/typednil"
	"github.com/gostaticanalysis/typeswitch"
	"github.com/gostaticanalysis/unitconst"
	"github.com/gostaticanalysis/unused"
	"github.com/gostaticanalysis/wraperrfmt"
	"github.com/gostaticanalysis/zapvet/zapvet"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(exclude.All(append(append(append(
		zagane.Analyzers(), // for Cloud Spanner
		ctxfield.Analyzer,
		dupimport.Analyzer,
		elseless.Analyzer,
		exclude.By(forcetypeassert.Analyzer, exclude.TestFile),
		importgroup.Analyzer,
		innertypealias.Analyzer,
		loopdefer.Analyzer,
		nilerr.Analyzer,
		noctor.Analyzer,
		nofmt.Analyzer,
		notest.Analyzer,
		sqlrows.Analyzer,
		testhelper.Analyzer,
		typednil.Analyzer,
		typeswitch.Analyzer,
		unitconst.Analyzer,
		unused.Analyzer,
		wraperrfmt.Analyzer),
		zapvet.Analyzers()...), // for zap
		debugcode.Analyzers()...,
	), exclude.GeneratedFile, exclude.FileWithPattern)...)
}
