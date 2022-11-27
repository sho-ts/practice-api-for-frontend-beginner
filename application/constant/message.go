package constant

import (
	"strconv"
)

var MESSAGE_NOT_FOUND = "コンテンツが見つかりませんでした"

var MESSAGE_HAS_QUERY_VALUE_PROBLEM = "クエリパラメータに問題があります"

var MESSAGE_UNKNOWN_ERROR = "エラーが発生しました"

func GetRequiredValidateErrorMessage(key string) string {
	return key + "は必須です"
}

func GetMaximumCharVaridateErrorMessage(key string, max int) string {
	return key + "は" + strconv.Itoa(max) + "文字以内である必要があります"
}
