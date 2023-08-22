package utils

import (
	"fmt"
	"time"
	"os"
	"errors"
	"io/ioutil"
	"regexp"

	article "myblogs/user/gin/pbfile/article"
)

var (
	html_part1 = "<!DOCTYPE html><html lang=\"zh-CN\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">    <title>技术专区</title><link rel=\"stylesheet\" href=\"styles.css\"><style>html, body {display: flex;justify-content: center;align-items: center;height: 100%;margin: 0;padding: 0;} .article {text-align: center;}.back-link {position: absolute;top: 10px;left: 10px;}</style></head><body><a class=\"back-link\" href=\"/"
	html_part6 = "\">返回上级</a><div class=\"article\"><h1 class=\"article-title\">"
	html_part2 = "</h1><div class=\"article-meta\"><span class=\"article-author\"><br>"
	html_part3 = "<br><br></span><span class=\"article-date\">"
	html_part4 = "</span></div><div class=\"article-content\">"
	html_part5 = "</div></div></body></html>"
)

func GetTime() string {
	t := time.Now()
	y := fmt.Sprint(t.Year())
	m := fmt.Sprint(int(t.Month()))
	d := fmt.Sprint(t.Day())
	h := fmt.Sprint(t.Hour())
	mi := fmt.Sprint(t.Minute())
	s := fmt.Sprint(t.Second())
	date := y+"-"+m+"-"+d+" "+h+":"+mi+":"+s
	return date
}

func SaveHTMLFile(art *article.Article, content string, id int64) error {
	filePath := "./articles/" + art.Column + "/" + fmt.Sprint(id) + ".html"
	file, err := os.OpenFile(filePath, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return errors.New("open file:" + filePath + " error!")
	}

	data := html_part1 + art.Column + html_part6 + art.Title + html_part2 + "作者:" + art.Author + html_part3 + "发表时间:" + art.Date + html_part4 + content + html_part5
	_, err = file.Write([]byte(data))
	if err != nil {
		return errors.New("write file:" + filePath + " error!")
	}

	return nil
}

func GetArticleContent(column string, id int64) (string, string) {
	filePath := "./articles/" + column + "/" + fmt.Sprint(id) + ".html"
    // 读取 HTML 文件内容
    htmlContent, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("无法读取文件:", err)
        return "", ""
    }

    // 将文件内容转换为字符串
    htmlString := string(htmlContent)

	re := regexp.MustCompile(`<div class="article-content">([\s\S]*?)</div>`)
	match := re.FindStringSubmatch(htmlString)

	articleContent := match[1]

	re = regexp.MustCompile(`<h1 class="article-title">([\s\S]*?)</h1>`)
	match = re.FindStringSubmatch(htmlString)

	titleContent := match[1]
	fmt.Println("title:", titleContent)

    // 打印提取的内容
    return titleContent, articleContent
}