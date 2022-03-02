package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const url = `https://xskydata.jobs.feishu.cn/api/v1/search/job/posts`

func main() {
	// 先拉到总数
	countD, err := pull(Params{
		Limit:          1,
		Offset:         0,
		PortalType:     6,
		PortalEntrance: 1,
	})
	if err != nil {
		panic(err)
	}

	data, err := pull(Params{
		Limit:          countD.Data.Count,
		Offset:         0,
		PortalType:     6,
		PortalEntrance: 1,
	})
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(data.Data.JobPostList)
	if err != nil {
		panic(err)
	}
	writeToFile(bytes)
}

func writeToFile(bytes []byte) error {
	file, err := os.OpenFile("list.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(bytes)
	return err
}

func pull(p Params) (*DataTemp, error) {
	var (
		res         DataTemp
		err         error
		paramsBytes []byte
		req         *http.Request
		resp        *http.Response
		c           *http.Client
	)
	paramsBytes, err = json.Marshal(p)
	if err != nil {
		return nil, err
	}
	c = &http.Client{}
	req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(paramsBytes))
	if err != nil {
		return nil, err
	}
	// 这里我实在没找到x-csrf-token刷新的逻辑
	req.Header.Add("x-csrf-token", "EpKZ6i3MrB-wPxzeKPn3dCt2Y7Wo6tgpm5DD4Yi37ok=")
	req.Header.Add("cookie", "channel=saas-career; platform=pc; device-id=7070116063477155369; s_v_web_id=verify_l0852dxn_hCcvugya_qiTF_4W7r_8Hls_hIihEZyg7oe3; SLARDAR_WEB_ID=a31e4864-1896-4e1e-9169-88359f5090cc; atsx-csrf-token=EpKZ6i3MrB-wPxzeKPn3dCt2Y7Wo6tgpm5DD4Yi37ok=")
	req.Header.Add("website-path", "school")
	resp, err = c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &res)
	return &res, err
}
