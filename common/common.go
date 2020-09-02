package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	ERR_CODE_CRT_USER = "FSS80001"
	ERR_CODE_CRT_ORG  = "FSS80002"
	ERR_CODE_CRT_FLOW = "FSS80003"
	ERR_CODE_GET_FLOW = "FSS80004"
	ERR_CODE_UP_FILE  = "FSS80005"
	ERR_CODE_DWN_FLOW = "FSS80006"
	ERR_CODE_QRY_FLOW = "FSS80007"

	ERR_CODE_SIGN_FLOW = "FSS80009"

	STATUS_DISABLED = 0
	STATUS_ENABLED  = 1

	OSS_SWITCH_ON  = "1"
	OSS_SWITCH_OFF = "0"

	STATUS_INIT = 0
	STATUS_PDF  = 1
	STATUS_USER = 2
	STATUS_ORG  = 2
	STATUS_CRT  = 3
	STATUS_SIGN = 4
	STATUS_CA   = 5
	STATUS_OSS  = 7
	STATUS_END  = 9

	STATUS_FAIL = ""

	DEFAULT_DOMAIN    = "http://pics2.crfchina.com/"
	DEFAULT_PATH      = "html/"
	DEFAULT_XLS_PATH  = "xlsx/"
	DEFAULT_FONT_SIZE = 14

	SYSTEM_CAGATE = "ca-gate"
	SYSTEM_FTS2   = "FTS2"

	GROUP_OSS = "oss"

	FIELD_LOGIN_PASS  = "user_pwd"
	FIELD_ERRORS      = "pwderr_cnt"
	FIELD_UPDATE_TIME = "update_date"
	FIELD_PROC_STATUS = "status"

	EMPTY_STRING = ""
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
		ERR_CODE_TOOBUSY: "短信发送太频繁",
		ERR_CODE_PDFERR:  "创建PDF文件出错",
	}
	//系统配置信息
	ConfMap        = make(map[string]string, 0)
	TOPWAY_URL     string
	API_ID         string
	API_KEY        string
	KEY_OSS_BUCKET = "cagate"
)

type ErrorResp struct {
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

const (
	NOW_TIME_FORMAT = "2006-01-02 15:04:05"

	DATA_REPORT      = 1
	AGREEMENT_REPORT = 2

	CHAPTER_LEVEL1 = 1
	CHAPTER_LEVEL2 = 2

	TMPL_SUFFIX = "P"

	KEY_EQB_HOST   = "eqb_host"
	KEY_EQB_APPID  = "eqb_appid"
	KEY_EQB_SECRET = "eqb_secret"
	KEY_EQB_TOKEN  = "eqb_token"

	KEY_EQB_INTERVAL = "eqb_interval"
	KEY_PDFPATH      = "local_pdfpath"

	KEY_KAFKA_HOST        = "kafka_host"
	KEY_DATA_REPORT_TOPIC = "kafka_dataReport_topic"
	KEY_AGRT_REPORT_TOPIC = "kafka_agrtReport_topic"
)

/*
	返回结构
*/
type UserPara struct {
	UserName string   `json:"fullname"`
	IdNo     string   `json:"idCardNum"`
	Position Position `json:"position"`
}

type Position struct {
	Page      int `json:"page"`
	PositionX int `json:"positionX"`
	PositionY int `json:"positionY"`
}

type ResponseMsgDetail struct {
	Code     string `json:"code"`
	Mesage   string `json:"message"`
	FileName string `json:"filename"`
	Group    string `json:"group"`
	FileId   string `json:"fileId"`
}

type Notify_Request_Msg struct {
	RcsFileId              string `json:"rcsFileId"`
	LoanNo                 string `json:"loanNo"`
	Template               string `json:"template"`
	FinalProtocolFileGroup string `json:"finalProtocolFileGroup"`
	FinalProtocolFileName  string `json:"finalProtocolFileName"`
	CrfUid                 string `json:"crfUid"`
}

type Notify_Response_Msg struct {
	Result    string `json:"result"`
	RcsFileId string `json:"rcsFileId"`
	LoanNo    string `json:"loanNo"`
}

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

func HttpPost(url string, packBuf []byte) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(packBuf))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req.WithContext(context.TODO()))
	if err != nil {
		log.Printf("client.Do%v", err)
		return EMPTY_STRING, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return EMPTY_STRING, err
	}
	var notifyResp Notify_Response_Msg
	if err = json.Unmarshal(respBytes, &notifyResp); err != nil {
		return EMPTY_STRING, err
	}
	if notifyResp.Result == "SUCCESS" {
		return notifyResp.LoanNo, nil
	} else {
		return notifyResp.LoanNo, fmt.Errorf("回调通知失败")
	}
}
