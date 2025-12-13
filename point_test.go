package flow

import (
	xlsx "github.com/openzee/xlsx-loader"
)

func LoadExcel() []*xlsx.DiscretePointMetadata {
	rst, _ := xlsx.ParseExcel("test/test_online.xlsx")
	return rst
}
