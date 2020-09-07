package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ERR_CODE_SUCCESS = "0000000"
	ERR_CODE_DBERROR = "FSS1001"
	ERR_CODE_TOKENER = "FSS1003"
	ERR_CODE_PARTOEN = "FSS1005"
	ERR_CODE_JSONERR = "FSS2001"
	ERR_CODE_URLERR  = "FSS2005"
	ERR_CODE_NOTFIND = "FSS3000"
	ERR_CODE_EXPIRED = "FSS6000"
	ERR_CODE_TYPEERR = "FSS4000"
	ERR_CODE_STATUS  = "FSS5000"
	ERR_CODE_FAILED  = "FSS9000"
	ERR_CODE_TOOBUSY = "FSS6010"
	ERR_CODE_PDFERR  = "FSS4040"
	ERR_CODE_PARAMS  = "FSS2044"
)

var (
	ERROR_MAP map[string]string = map[string]string{
		ERR_CODE_SUCCESS: "执行成功",
		ERR_CODE_DBERROR: "DB执行错误",
		ERR_CODE_JSONERR: "JSON格式错误",
		ERR_CODE_EXPIRED: "时效已经到期",
		ERR_CODE_TYPEERR: "类型转换错误",
		ERR_CODE_STATUS:  "状态不正确",
		ERR_CODE_PARAMS:  "输入参数不合法",
		ERR_CODE_TOKENER: "获取TOKEN失败",
		ERR_CODE_PARTOEN: "解析TOKEN错误",
		ERR_CODE_NOTFIND: "查询没发现提示",
	}
	//系统配置信息
	ConfMap = make(map[string]string, 0)
)

type ErrorResp struct {
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

const (
	NOW_TIME_FORMAT = "2006-01-02 15:04:05"
)

func PrintHead(a ...interface{}) {
	log.Println("========》", a)
}

func PrintTail(a ...interface{}) {
	log.Println("《========", a)
}

func Write_Response(response interface{}, w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8087")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type,Action, Module,Authorization")
	fmt.Fprintf(w, string(json))
}
