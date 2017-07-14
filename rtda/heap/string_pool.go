package heap

import "unicode/utf16"

// 字符串池，key表示Go字符串（UTF8），value表示Java字符串（UTF16）
var internedStrings = map[string]*Object{}

// 根据Go字符串返回响应的Java字符串实例
func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	// 将Go字符串形式转换成Java字符串形式
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}
	// 创建Java字符串，赋值
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	// 添加至字符串池
	internedStrings[goStr] = jStr
	return jStr
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
