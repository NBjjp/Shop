package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

//时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

//获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

//获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

//获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//将string类型转换成int类型
func Int(str string) (int, error) {
	num, err1 := strconv.Atoi(str)
	return num, err1
}

//将int类型转换成string类型
func String(num int) string {
	str := strconv.Itoa(num)
	return str
}

//上传文件
func UploadImg(ctx *gin.Context, picName string) (string, error) {
	//1、获取上传的文件
	file, err1 := ctx.FormFile(picName)
	if err1 != nil {
		return "", err1
	}
	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}
	// 3、创建图片保存目录  static/upload/20210624
	day := GetDay()
	dir := "./static/upload/" + day

	err2 := os.MkdirAll(dir, 0666)
	if err2 != nil {
		fmt.Println(err2)
	}
	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnix(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	ctx.SaveUploadedFile(file, dst)
	return dst, nil
}
