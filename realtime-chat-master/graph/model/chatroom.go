/*
@Time : 2020/11/4 16:11
@Author : lai
@Description :
@File : chatroom
*/
package model

type Chatroom struct {
	Name      string    `json:"name"`
	Messages  []Message `json:"messages"`
	Observers map[string]struct {
		Username string
		Message  chan *Message
	}
}
