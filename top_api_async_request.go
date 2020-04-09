package godingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const topAPIAsyncRootURL = "https://oapi.dingtalk.com/topapi/message/corpconversation"

type TopAPIAsyncResponse interface {
	checkError() error
}

type TopAPIAsyncRequest struct {
	AgentId    int                    `json:"agent_id"`
	UserIdList []string               `json:"userid_list"`
	DeptIdList []string               `json:"dept_id_list"`
	ToAllUser  bool                   `json:"to_all_user"`
	Msg        map[string]interface{} `json:"msg"`
}

type TopCMCAResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	TaskId  int    `json:"task_id"`
}

type topAPIAsyncErrResponse struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestID string `json:"request_id"`
}

func (result *topAPIAsyncErrResponse) checkError() (err error) {
	if result.ErrCode != 0 || len(result.SubCode) != 0 {
		err = fmt.Errorf("%#v", result)
	}

	return err
}

func (c *DingTalkClient) topAPIAsyncRequest(params url.Values, path string, respData TopAPIAsyncResponse) error {
	resp, err := http.PostForm(topAPIAsyncRootURL+path, params)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Server error: " + resp.Status)
	}
	defer resp.Body.Close()

	respBuf, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err := json.Unmarshal(respBuf, &respData)
		if err != nil {
			return err
		}
		return respData.checkError()
	}

	return err
}
