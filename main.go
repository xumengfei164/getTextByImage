package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"os"
	"path/filepath"
)

func main() {
	var outputFolderPath string
	var filePath string
	// 图片地址
	fmt.Print("请输入图片地址：（例如：/Users/Ven/images/image1.png或D:/Ven/images/image1.png）\n:")
	_, err := fmt.Scan(&filePath)
	if err != nil {
		fmt.Println("发生错误:", err)
		return
	}
	// 语言选择
	var lang int
	fmt.Print("请选择待识别图片中的语言种类（默认中文简体）：\n1->中文简体\n2->英语\n:")
	_, err = fmt.Scan(&lang)
	if err != nil {
		fmt.Println("发生错误:", err)
		return
	}
	var langTyp string
	switch lang {
	case 1:
		langTyp = "chi_sim"
	case 2:
		langTyp = "eng"
	default:
		langTyp = "chi_sim"

	}
	// 生成文字
	fmt.Print("请输入文字生成地址：（输入回车在图片所在目录下生成内容文件）\n:")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("发生错误:", err)
		return
	}
	if len(userInput) == 1 && userInput[0] == '\n' {
		outputFolderPath = filepath.Dir(filePath) + "/output.txt"
	} else {
		outputFolderPath = string(userInput)
	}
	//fmt.Println(filePath, langTyp)
	//return
	text, err := getTextByImage(filePath, langTyp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = writeFile(outputFolderPath, text)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
	fmt.Println("调用成功.")
}

func getTextByImage(path, lang string) (string, error) {
	if lang == "" {
		lang = "eng"
	}
	client := gosseract.NewClient()
	defer client.Close()

	// 设置图像文件路径
	err := client.SetImage(path)
	if err != nil {
		err = errors.New(fmt.Sprintf("图片路径设置错误:%v", err))
		return "", err
	}
	// 设置语言（可选）
	err = client.SetLanguage(lang)
	if err != nil {
		err = errors.New(fmt.Sprintf("语言设置错误:%v", err))
		return "", err
	}
	// 执行 OCR
	text, err := client.Text()
	if err != nil {
		err = errors.New(fmt.Sprintf("生成文本错误:%v", err))
		return "", err
	}

	// 返回识别内容
	return text, nil
}

func writeFile(filePath, content string) error {
	// 创建或打开文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将内容写入文件
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	//fmt.Printf("文件 %s 写入成功\n", filePath)
	return nil
}
