package consts

// 整型数据类型取值范围
// 在系统中一般做数据校验的时候,最小值取0,最大值不能超过对应数据的最大值
const (
	// MinUint8Value 无符号 8 位整型 (0 到 255)
	MinUint8Value = 0
	MaxUint8Value = 255

	// MinUint16Value 16 位整型 (0 到 65535)
	MinUint16Value = 0
	MaxUint16Value = 65535

	// MinUint32Value 32 位整型 (0 到 65535)
	MinUint32Value = 0
	MaxUint32Value = 4294967295

	// MinUint64Value 无符号 64 位整型 (0 到 18446744073709551615)
	MinUint64Value = 0
	MaxUint64Value = 4294967295

	// MinInt8Value 有符号 8 位整型 (-128 到 127)
	MinInt8Value = -128
	MaxInt8Value = 127

	// MinInt16Value 有符号 16 位整型 (-32768 到 32767)
	MinInt16Value = -32768
	MaxInt16Value = 32767

	// MinInt32Value 有符号 32 位整型 (-2147483648 到 2147483647)
	MinInt32Value = -2147483648
	MaxInt32Value = 2147483647

	// MinInt64Value 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	MinInt64Value = -9223372036854775808
	MaxInt64Value = 9223372036854775807
)
