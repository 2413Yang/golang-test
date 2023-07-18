package main

import "fmt"
import "strings"

func urlProcess(url string) string {
	if result := strings.HasPrefix(url, "http://"); !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}
func pathProcess(path string) string {
	if !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}
func main() {
	var (
		url  string
		path string
	)
	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println("url:", url)
	fmt.Println("path:", path)
}
