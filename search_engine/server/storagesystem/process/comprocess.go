package process

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Compressor struct {
}

//从文件中二进制数据
func (this *Compressor) ReadFile_Object(filename string) (data []byte, err error) {
	pathurl := strings.TrimSpace(filename)

	file, err := os.Open(pathurl)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("解码失败", err.Error())
		return
	} else {
		fmt.Println("解码成功")
		//数据解压缩
		data = this.DoZlibUnCompress(data)
	}
	return
}

//写入文件二进制数据
func (this *Compressor) WriteFil_eObject(filename string, data []byte) (err error) {

	pathurl := strings.TrimSpace(filename)

	file, err := os.Create(pathurl)
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()
	//数据压缩
	data = this.DoZlibCompress(data)

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
	return
}

//进行zlib压缩
func (this *Compressor) DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	if len(in.Bytes()) > len(src) {
		fmt.Println("出发狗压缩")
	}
	log.Println("原大小：", len(src), "压缩后大小：", len(in.Bytes()), "压缩率：", fmt.Sprintf("%.2f", float32(len(in.Bytes()))/float32(len(src))), "%")
	return in.Bytes()
}

//进行zlib解压缩
func (this *Compressor) DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

//按行读取文件
func (this *Compressor) ReadFileOnline(filepath string) (strs []string, err error) {
	fmt.Println("go语言读取文件")
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	flag := 0
	for flag < 100 {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		} else {
			strs = append(strs, line)
		}
		flag++
	}
	return
}
