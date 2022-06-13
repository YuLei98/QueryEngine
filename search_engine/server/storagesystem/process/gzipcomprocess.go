package process

import (
	"bytes"
	"compress/flate"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Gzipcomprocess struct {
}
func (this *Gzipcomprocess) ExecTime(fn func()) float64 {
	start := time.Now()
	fn()
	tc := float64(time.Since(start).Nanoseconds())
	return tc / 1e6
}

// Write 写入二进制数据到磁盘文件
func (this *Gzipcomprocess) Write(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	//日志文件开
	//log.Println("Write:", filename)

	// compressData := this.Compression(buffer.Bytes())

	// err = ioutil.WriteFile(filename, compressData, 0600)
	// if err != nil {
	// 	panic(err)
	// }



	//compressData := this.Compression(buffer.Bytes())

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func (this *Gzipcomprocess) Encoder(data interface{}) []byte {
	if data == nil {
		return nil
	}
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		log.Println("出现panic");
		//panic(err)
	}
	return buffer.Bytes()
}

func (this *Gzipcomprocess) Decoder(data []byte, v interface{}) {
	if data == nil {
		return
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(v)
	if err != nil {
		log.Println("出现panic",err);
		 
		//panic(err)
	}
	
}

// Compression 压缩数据
func (this *Gzipcomprocess) Compression(data []byte) []byte {
	buf := new(bytes.Buffer)
	write, err := flate.NewWriter(buf, flate.DefaultCompression)
	defer write.Close()

	if err != nil {
		panic(err)
	}

	write.Write(data)
	write.Flush()
	//log.Println("原大小：", len(data), "压缩后大小：", buf.Len(), "压缩率：", fmt.Sprintf("%.2f", float32(buf.Len())/float32(len(data))), "%")
	return buf.Bytes()
}

//this.Decompression 解压缩数据
func (this *Gzipcomprocess) Decompression(data []byte) []byte {
	return this.DecompressionBuffer(data).Bytes()
}

func (this *Gzipcomprocess) DecompressionBuffer(data []byte) *bytes.Buffer {
	buf := new(bytes.Buffer)
	read := flate.NewReader(bytes.NewReader(data))
	defer read.Close()

	buf.ReadFrom(read)
	return buf
}

// Read 从磁盘文件加载二进制数据
func (this *Gzipcomprocess) Read(filename string) (data []byte, err error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			//忽略
			return
		}
		panic(err)
	}
	
	// //解压
	// decoData := this.Decompression(raw)
	
	// buffer := bytes.NewBuffer(decoData)
	// dec := gob.NewDecoder(buffer)
	// err = dec.Decode(&data)
	// if err != nil {
	// 	panic(err)
	// }

	//不解压
		buffer := bytes.NewBuffer(raw)
			dec := gob.NewDecoder(buffer)
			err = dec.Decode(&data)
			if err != nil {
				panic(err)
			}
	return
}