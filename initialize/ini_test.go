package initialize

import (
	"testing"
)

func TestValidateCode(t *testing.T) {
	var tests = []struct {
		path string
	}{
		{"/Users/daidai/code/golang/projects/company-projects/blx-equity-platform/admin/core/dao/authdao/node_dao.go"},
		{"/Users/daidai/code/golang/projects/company-projects/blx-equity-platform/admin/core/dao/authdao/role_dao.go"},
		{"/Users/daidai/code/golang/projects/company-projects/blx-equity-platform/admin/core/dao/authdao/user_dao.go"},
		{"/Users/daidai/code/golang/projects/company-projects/legou/main.go"},
	}

	for _, test := range tests {
		err := GormCodeValidate(test.path)
		if err != nil {
			t.Fatalf("gormCodeValidate error:%v", err.Error())
		}
	}
}
