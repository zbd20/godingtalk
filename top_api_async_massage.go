package godingtalk

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

const (
	AsyncSend    = "/asyncsend_v2"
	SendProgress = "/getsendprogress"
	SendResult   = "/getsendresult"
	Recall       = "/recall"
)

type topAPIAsyncMsgSendResponse struct {
	topAPIAsyncErrResponse
	TaskID int `json:"task_id"`
}

func (c *DingTalkClient) TopAPIAsyncSend(msgType string, userList []string, deptList []string, toAll bool, msgContent interface{}) (int, error) {
	if err := c.RefreshAccessToken(); err != nil {
		return 0, err
	}

	var resp topAPIAsyncMsgSendResponse
	if len(userList) > 100 {
		return 0, errors.New("can't more than 100 users at once")
	}
	if len(deptList) > 20 {
		return 0, errors.New("can't more than 20 departments at once")
	}

	toAllStr := "false"
	if toAll {
		toAllStr = "true"
	}

	var msg map[string]interface{}
	switch msgType {
	case "action_card":
		msg = map[string]interface{}{
			"msgtype":     msgType,
			"action_card": msgContent,
		}
	case "text":
		msg = map[string]interface{}{
			"msgtype": msgType,
			"text":    msgContent,
		}
	case "markdown":
		msg = map[string]interface{}{
			"msgtype":  msgType,
			"markdown": msgContent,
		}
	default:
		return 0, errors.New("invalid msgtype")
	}

	mbyte, err := json.Marshal(msg)
	if err != nil {
		return 0, err
	}

	params := url.Values{
		"access_token": {c.AccessToken},
		"agent_id":     {c.AgentID},
		"userid_list":  {strings.Join(userList, ",")},
		"to_all_user":  {toAllStr},
		"msg":          {string(mbyte)},
	}

	return resp.TaskID, c.topAPIAsyncRequest(params, AsyncSend, &resp)
}
