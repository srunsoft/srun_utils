package utils_test

import (
	utils "e.coding.net/srun/go_utils/srun_utils"
	"testing"
)

func TestReadExcel(t *testing.T) {
	excel, err := utils.ReadExcel("../tmp/eps_add_template.xlsx", "批量添加SIM卡")
	if err != nil {
		t.Error(err)
	}
	t.Log(excel)
	if len(excel) != 126376 {
		t.Error("len(excel) = ", len(excel))
	}
}
