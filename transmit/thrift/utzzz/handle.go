package utzzz

import "encoding/json"

type S_HeartBeatData struct {
	Account string `json:"account"`
}

type S_HeartBeat struct {
	Method        string          `json:"method"`
	Station       string          `json:"station"`
	IsAck         string          `json:"isAck"`
	HeartBeatData S_HeartBeatData `json:"heartBeatData"`
}

type S_HeartBeatAckData struct {
	ReturnCode string `json:"returnCode"`
	Msg        string `json:"msg"`
}

type S_HeartBeatACK struct {
	Method  string             `json:"method"`
	Station string             `json:"Station"`
	IsAck   string             `json:"isAck"`
	Data    S_HeartBeatAckData `json:"data"`
}

func HeartBeat(contnet string) *S_HeartBeatACK {
	heartBeat := &S_HeartBeat{}
	hb_ack := &S_HeartBeatACK{
		Method: "MS_SendHeartBeat_ACK",
		IsAck:  "1",
	}
	json.Unmarshal([]byte(contnet), &heartBeat)
	hb_ack_data := S_HeartBeatAckData{
		ReturnCode: "1",
		Msg:        "",
	}

	hb_ack.Data = hb_ack_data
	hb_ack.Station = heartBeat.Station
	return hb_ack
}
