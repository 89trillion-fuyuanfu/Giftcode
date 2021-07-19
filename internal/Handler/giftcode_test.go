package Test

import (
	"Giftcode/internal/service"
	"testing"
)

// 单元测试
func Test_giftcode(t *testing.T) {
	// 赋予测试参数

	cases := []struct {
		a      string
		expect string
	}{

		{"9Xf9GAx5", "礼品内容：三只小猫咪 创建人姓名：扶元富 创建时间：2021"},
	}
	for _, c := range cases {
		if r := service.Getstring(c.a); r != c.expect {
			t.Errorf("test failed")
		}
	}

}
