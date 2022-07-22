package strutil

import (
	"fmt"
	"testing"
)

func TestGetJsonKey(t *testing.T) {
	json := `{
    "orderNo":"PFT8888888885",
    "orderName":"2022-07-05-2022-07-05[林练来]",
    "resourceId":"9638966",
    "payType":0,
    "orderPrice":0.2,
    "checkType":1,
    "contactName":"林练来",
    "contactPhone":"13484068986",
    "idNumber":"522123199504014516",
    "payAmount":1,
    "depositAmount":0,
    "remark":"",
    "total":2,
    "peopleNumber":2,
    "timeStock":"",
    "verificationCode":"",
    "changeOffSupply":1,
    "receiptSupply":0,
    "refundSupply":1,
    "orderMode":"",
    "code":"test123,tets5555,test1234567890",
    "qrcode":"http://open.12301dev.com/code/test123.png?type=manual_code,http://open.12301dev.com/code/tets5555.png?type=manual_code,http://open.12301dev.com/code/test1234567890.png?type=manual_code",
    "refundChanrge":0,
    "smsCode":"",
    "isDeposit":0,
    "startDateAuto":0,
    "isPay":1,
    "isSms":0,
    "isPrint":0,
    "status":14,
    "payTime":1656992355,
    "validationTime":0,
    "cancelTime":0,
    "startDate":"2022-07-05",
    "endDate":"2022-07-05",
    "leaveTime":0,
    "createTime":1656992355,
    "updateTime":0,
    "touristors":[
        {
            "order_No":"PFT8888888885",
            "subOrderNo":"415173610584408074",
            "certificateType":"0",
            "idNumber":"520181198709021370",
            "name":"林练来2",
            "phone":"15522061206",
            "qrcodeUrl":"",
            "createTime":1656992355
        },
        {
            "orderNo":"PFT8888888885",
            "subOrderNo":"415173610584473610",
            "certificateType":"0",
            "idNumber":"522422199210117210",
            "name":"daidai",
            "phone":"15522061206",
            "qrcodeUrl":"",
            "createTime":1656992355
        }
    ]
}`
	res := GetJsonKeys([]byte(json))
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestGetRandom(t *testing.T) {
	for i := 0; i < 8; i++ {
		str := GetRandom(16, LetterNumbers)
		fmt.Println(str)
	}
}

func TestIdCard(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{"111111111111111111", false},
		{"41080119930228457X", true},
		{"510801197609022309", true},
		{"150401198107053872", true},
		{"130133197204039024", true},
		{"632722197112040806", true},
		{"130683199011300601", true},
		{"350111199409241690", true},
		{"522323198705037737", true},
		{"510182197109294463", true},
		{"653221197910077436", true},
		{"533526197206260908", true},
		{"230305198909078721", true},
		{"232304198204030301", true},
		{"411425198812189711", true},
		{"350521197404071798", true},
		{"542128198709025957", true},
		{"350321198401316749", true},
		{"440804197710034663", true},
		{"372900197507012999", true},
		{"41080119930228457X", true},
		{"510801197609022309", true},
		{"150401198107053872", true},
		{"130133197204039024", true},
		{"430102197606046442", true},
		{"632722197112040806", true},
		{"130683199011300601", true},
		{"350111199409241690", true},
		{"522323198705037737", true},
		{"510182197109294463", true},
		{"533526197206260908", true},
		{"230305198909078721", true},
		{"232304198204030301", true},
		{"411425198812189711", true},
		{"350521197404071798", true},
		{"522422199210117210", true},
		{"510902199308240717", true},
		{"520202199508177832", true},
		{"522125199303010079", true},
		{"522726199012130918", true},
		{"420106197904123614", true},
		{"520121199408201042", true},
		{"522126199101171516", true},
		{"522129199306255034", true},
		{"522501199402062043", true},
		{"370303199404213126", true},
		{"522422198703081012", true},
		{"52212719891103653X", true},
		{"522422199210117210", true},
		{"520181198709021370", true},
		{"520123198612291215", true},
		{"520103199402195216", true},
		{"522723199402190212", true},
		{"522123199504014516", true},
		{"520201198806090419", true},
		{"230225198212204513", true},
		{"230602198612227519", true},
		{"2", false},
		{"2222", false},
		{"522422199310227654", false}, // 存在问题
		{"51080119760902230999", false},
		{"51080160902230999", false},
	}

	for _, test := range tests {
		res := CheckIdCard(test.input)
		if res != test.output {
			t.Fatalf("inCard = %s,expect = %v,actual = %v", test.input, test.output, res)
		}
	}
}
