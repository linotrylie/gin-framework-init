package passwordutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

// 测试 aes 正常加密解密
func TestAesEncrypt(t *testing.T) {
	orderInfo := `{"orderNo":"OS20220527152011111","platformOrderNo":"","orderName":"黄果树景区-齐天大圣水帘洞洞天","mchAppKey":"","supplierId":0,"resourceId":"111859916053942272","payType":5,"orderPrice":520.13,"checkType":1,"contactName":"林练来","contactPhone":"15522061206","idNumber":"522123199505014516","payAmount":520.13,"depositAmount":0,"remark":"","total":3,"peopleNumber":3,"timeStock":"","shuttleAddress":"","plateNumber":"","postAddress":"","verificationCode":"","changeOffSupply":0,"receiptSupply":0,"refundSupply":0,"orderMode":"","code":"","isDeposit":"","startDateAuto":"","isPay":"","isSms":"","isPrint":"","status":"","statusDes":"","payTime":1653667728,"startDate":1653667728,"endDate":1653667728,"touristorList":[{"orderNo":"OS20220527152011111","subOrderNo":"","certificateType":"0","idNumber":"522123199505014514","name":"林练来","phone":"15522061206","age":"","sex":"MALE","birthday":"","province":"","city":"","country":""},{"orderNo":"OS20220527152011111","subOrderNo":"","certificateType":"0","idNumber":"522123199505014515","name":"林练来1","phone":"15522061206","age":"","sex":"MALE","birthday":"","province":"","city":"","country":""},{"orderNo":"OS20220527152011111","subOrderNo":"","certificateType":"0","idNumber":"522123199505014516","name":"林练来2","phone":"15522061206","age":"","sex":"MALE","birthday":"","province":"","city":"","country":""}]}`
	var tests = []struct {
		key  []byte
		iv   []byte
		data []byte
	}{
		{[]byte("qqqqqqqqqqqqqqq1"), []byte("aaaaaaaaaaaaaaaa"), []byte("ihaveadreamshcsbdhcbsdhcbhsdbcsdc中文不会不会变速档好吧999999")},
		{[]byte("qqqqqqqqqqqqqqq1"), []byte("aaaaaaaaaaaaaaaa"), []byte("666666666")},
		{[]byte("qqqqqqqqqqqqqqq1"), []byte("aaaaaaaaaaaaaaaa"), []byte("eplnflhjhddmfoapljegainfboocoibjhpldhbcbifpohccdfabfkienhnnepadmnfghcmefdncimigeh")},
		{[]byte("1111111111111111"), []byte("2222222222222222"), []byte("lihjdlpndphhhfckkiaeoakjaigffhhnjhgjakpefpkkccphmofbbmmplojnooiephgopcpjjjkejhi")},
		{[]byte("aaaaaaaaaaaaaaaa"), []byte("bbbbbbbbbbbbbbbb"), []byte("lihjdlpndphhhfckkiaeoakjaigffhhnjhgjakpefpkkccphmofbbmmplojnooiephgopcpjjjkejhi")},
		{[]byte("ZyglMvRwH3oodB5BZyglMvRwH3oodB5a"), []byte("ZyglMvRwH3oodB5B"), []byte(orderInfo)},
	}

	for _, test := range tests {

		// 加密
		encrypt, err := AesCBCEncrypt(test.data, test.key, test.iv)
		assert.Nil(t, err)
		assert.NotEmpty(t, encrypt)

		// 解密
		decrypt, err := AesCBCDecrypt(encrypt, test.key, test.iv)
		assert.Nil(t, err)
		assert.NotEmpty(t, decrypt)
		assert.Equal(t, decrypt, test.data)
	}
}

// 测试 黄果树专用的 aes 加密解密
func TestHuanGuoShuAesEncrypt(t *testing.T) {
	var tests = []struct {
		key  []byte
		iv   []byte
		data []byte
	}{
		{[]byte("qqqqqqqqqqqqqqq1"), []byte("aaaaaaaaaaaaaaaa"), []byte("ihaveadreamshcsbdhcbsdhcbhsdbcsdc中文不会不会变速档好吧999999")},
		{[]byte("qqqqqqqqqqqqqqq1"), []byte("aaaaaaaaaaaaaaaa"), []byte("666666666")},
		{[]byte("qqqqqqqqqqqqqqq1"), []byte("aaaaaaaaaaaaaaaa"), []byte("eplnflhjhddmfoapljegainfboocokccphmofbbmmplojnooiephgopcpjjjkejhi")},
		{[]byte("1111111111111111"), []byte("2222222222222222"), []byte("eplnflhjhddmfoapljegainfboocokccphmofbbmmplojnooiephgopcpjjjkejhi")},
		{[]byte("1111111111111111"), []byte("2222222222222222"), []byte("Den kinesiska sidan svarade på Bidens kommentarer: schakaler kommer och det finns hagelgevär")},
		{[]byte("1111111111111111"), []byte("2222222222222222"), []byte("中国側はバイデン氏の発言に答えた：狼が来た猟銃がある")},
		{[]byte("1111111111111111"), []byte("2222222222222222"), []byte("[好的]中国側はバイデン氏の発言に答えた：狼が来た猟銃がある")},
		{[]byte("1111111111111111"), []byte("2222222222222222"), []byte("중국 측은 바이든의 발언에 대해 승냥이가 오면 사냥총이 있다고 대답했다")},
		{[]byte("1111111111111111qqqqqqqqqqqqqqq1"), []byte("2222222222222222"), []byte("我梦江南好,征辽亦偶然")},
	}
	for _, test := range tests {
		// 加密
		encrypt, err := AesCBCEncrypt(test.data, test.key, test.iv)
		assert.Nil(t, err)
		assert.NotEmpty(t, encrypt)
		encryptRes := huangGuoshuEncodeBytes(encrypt)
		fmt.Println("加密结果:", huangGuoshuEncodeBytes(encrypt))

		// 解密
		encryptEncode := huangGuoshuDecodeBytes(encryptRes)
		decrypt, err := AesCBCDecrypt([]byte(encryptEncode), test.key, test.iv)
		assert.Nil(t, err)
		assert.NotEmpty(t, decrypt)
		assert.Equal(t, decrypt, test.data)
		fmt.Println("解密结果:", string(decrypt))
	}
}

// 黄果树解密编码转换
func huangGuoshuDecodeBytes(str string) string {
	sliceInt8 := []int8{}
	for i := 0; i < len(str); i += 2 {
		c := str[i]
		a := int8((c - 'a') << 4)

		c1 := str[i+1]
		b := int8(c1 - 'a')

		sliceInt8 = append(sliceInt8, a+b)
	}

	bh := (*reflect.SliceHeader)(unsafe.Pointer(&sliceInt8))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// 黄果树加密编码转换
func huangGuoshuEncodeBytes(bytes []byte) string {
	sb := strings.Builder{}
	for i := 0; i < len(bytes); i++ {
		sb.WriteRune(rune((bytes[i]>>4)&0xF + 97))
		sb.WriteRune(rune((bytes[i] & 0xF) + 97))
	}
	return sb.String()
}
