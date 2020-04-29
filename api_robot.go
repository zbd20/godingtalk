package godingtalk

import (
	"net/url"
)

type RobotAtList struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

//SendRobotTextMessage can send a text message to a group chat
func (c *DingTalkClient) SendRobotTextMessage(accessToken string, msg string) (data MessageResponse, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}

//SendRobotMarkdownMessage can send a text message to a group chat
func (c *DingTalkClient) SendRobotMarkdownMessage(accessToken string, title string, msg string) (data MessageResponse, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  msg,
		},
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}

// SendRobotTextAtMessage can send a text message and at user to a group chat
func (c *DingTalkClient) SendRobotTextAtMessage(accessToken string, msg string, at *RobotAtList) error {
	var data OAPIResponse
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
		"at": at,
	}
	err := c.httpRPC("robot/send", params, request, &data)
	return err
}

//SendRobotMarkdownMessage can send a text message to a group chat
func (c *DingTalkClient) SendRobotMarkdownAtMessage(accessToken string, title string, msg string, at *RobotAtList) (data MessageResponse, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  msg,
		},
		"at": at,
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}

func (c *DingTalkClient) SendRobotActionCardMessage(accessToken string, title string, msg string, btns []map[string]interface{}, btnOrientation string) (data MessageResponse, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"title":          title,
			"text":           msg,
			"btnOrientation": btnOrientation,
			"btns":           btns,
		},
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}

func (c *DingTalkClient) SendRobotActionCardAtMessage(accessToken string, title string, msg string, btns []map[string]interface{}, btnOrientation string, at *RobotAtList) (data MessageResponse, err error) {
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"title":          title,
			"text":           msg,
			"btnOrientation": btnOrientation,
			"btns":           btns,
		},
		"at": at,
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}
