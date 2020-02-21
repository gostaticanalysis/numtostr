package main

import (
	"github.com/gostaticanalysis/numtostr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(numtostr.Analyzer) }