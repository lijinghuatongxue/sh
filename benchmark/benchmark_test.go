package main

import (
	Meutils "github.com/lijinghuatongxue/utils"
	"testing"
)

func BenchmarkBase64(b *testing.B) {
	//Meutils.Base64Encode("111213123sdasdas")
	Meutils.WfMain("./writeTest.txt", "sdsadasdasdasdasdasdas\n", false, false)
}
