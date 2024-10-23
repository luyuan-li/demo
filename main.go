package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var owner = "luyuan-li"

	var repo = "demo"

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?sha=main&path=address.json&since=2024-10-23T03:11:36Z", owner, repo)
	fmt.Println(url)

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加头部信息，如果需要认证，请取消注释并填入 TOKEN
	// req.Header.Set("Authorization", "token YOUR_PERSONAL_ACCESS_TOKEN")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 检查响应是否成功
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Error: %s\n", body)
		return
	}

	// 读取并解析 JSON 响应
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf(string(body))
}

// File 结构体对应 GitHub API 中的文件对象
type File struct {
	SHA       string `json:"sha"`
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
	Changes   int    `json:"changes"`
	Patch     string `json:"patch"`
}

// Commit 结构体对应 GitHub API 中的提交对象
type Commit struct {
	Files []File `json:"files"`
}

/*func main() {
	url := "https://api.github.com/repos/luyuan-li/demo/commits/a2ac8c025808c4db088be7f6b7b38558b8a96ffd"

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加头部信息，如果需要认证，请取消注释并填入 TOKEN
	// req.Header.Set("Authorization", "token YOUR_PERSONAL_ACCESS_TOKEN")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 检查响应是否成功
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Error: %s\n", body)
		return
	}

	// 读取并解析 JSON 响应
	body, _ := ioutil.ReadAll(resp.Body)
	var commit Commit
	if err := json.Unmarshal(body, &commit); err != nil {
		fmt.Println(err)
		return
	}

	// 输出文件的 diff 信息
	for _, file := range commit.Files {
		fmt.Printf("Filename: %s\n", file.Filename)
		fmt.Println("Patch:")
		patchLines := strings.Split(file.Patch, "\n")
		for _, line := range patchLines {
			fmt.Println(line) // 输出 diff 内容
		}
		fmt.Println("---")
	}
}*/
