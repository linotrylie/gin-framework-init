package passwordutil

import (
	"equity/utils/fileutil"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestRsaGenKey 通过直接生成秘钥对字符串方式测试生成秘钥对,加密,解密
func TestRsaGenKey(t *testing.T) {
	rsaKey, err := RsaGenKeyToBytes(2048)
	assert.Nil(t, err)

	// 私钥
	privateKey := rsaKey.PrivateKey
	// 公钥
	publicKey := rsaKey.PublicKey

	// 公钥私钥不能为空
	assert.NotEmpty(t, privateKey)
	assert.NotEmpty(t, publicKey)

	// 私钥公钥不相等
	assert.NotEqual(t, privateKey, publicKey)

	// 待加密字符串
	wantEncryptData := "i have a dream!"

	// 加密
	EncryptData, err := RSAEncrypt([]byte(wantEncryptData), publicKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, EncryptData)

	// 解密
	decrypt, err := RSADecrypt(EncryptData, privateKey)
	assert.Nil(t, err)
	assert.Equal(t, wantEncryptData, string(decrypt))
}

// TestRsaGenKeyByFile 通过存文件方式测试生成秘钥对,加密,解密
func TestRsaGenKeyByFile(t *testing.T) {
	err := RsaGenKey(2048)
	assert.Nil(t, err)

	getwd, _ := os.Getwd()
	publicFilePath := getwd + fileutil.FileSeparator() + "public.pem"
	privateFilePath := getwd + fileutil.FileSeparator() + "private.pem"
	assert.FileExistsf(t, publicFilePath, fmt.Sprintf("文件不存在,file=%s", publicFilePath))
	assert.FileExistsf(t, privateFilePath, fmt.Sprintf("文件不存在,file=%s", privateFilePath))

	// 通过文件加密
	wantEncryptData := `thisisatestring`
	EncryptData, err := RSAEncryptByFile([]byte(wantEncryptData), publicFilePath)
	assert.Nil(t, err)
	assert.NotEmpty(t, EncryptData)

	// 解密
	DecryptData, err := RSADecryptByFile(EncryptData, privateFilePath)
	assert.Nil(t, err)

	// 判断解密信息和明文是否相等
	assert.Equal(t, string(DecryptData), wantEncryptData)

	// 删除公钥文件
	err = os.Remove(publicFilePath)
	assert.Nil(t, err)

	// 删除私钥文件
	err = os.Remove(privateFilePath)
	assert.Nil(t, err)
}

func TestMd5Encode(t *testing.T) {
	md5 := MD5Encode("25bbdfdb9c674da8e2ab36d37a2c226e&_pid=2583764&format=json&id_number=522123199505014516&is_pay=0&item_id=111859916053942272&leave_date=2022-06-09&method=item_orders&mobile=15522061206&name=林练来&orders_id=OS20220527152011111&players[0][id_ntype]=0&players[0][id_number]=522123199505014514&players[0][mobile]=15522061206&players[0][name]=林练来&price=0.000000&size=3&start_date=2022-06-08&start_date_auto=0&25bbdfdb9c674da8e2ab36d37a2c226e")
	fmt.Println(md5)
}
