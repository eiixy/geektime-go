package code

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type ReadFromFile struct {
	path     string
	interval time.Duration
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error %s", err))
	}
	// 移动字符指针到文件末尾
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(r.interval)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("read file error %s", err))
		}
		rc <- line[:len(line)-1]
	}
}

type WriteToDatabase struct {
	dsn string
}

func (w *WriteToDatabase) Write(wc chan string) {
	for s := range wc {
		fmt.Println(s)
	}
}

type LogProcess struct {
	rc     chan []byte
	wc     chan string
	reader Reader
	writer Writer
}

func (l *LogProcess) Process() {
	for bytes := range l.rc {
		l.wc <- string(bytes)
	}
}

func Handler() {
	// 读取器
	reader := &ReadFromFile{
		path:     "D:\\environments\\phpstudy_pro\\Extensions\\Nginx1.15.11\\logs\\access.log",
		interval: 100 * time.Millisecond,
	}

	// 写入器
	writer := &WriteToDatabase{
		dsn: "",
	}

	log := &LogProcess{
		rc:     make(chan []byte),
		wc:     make(chan string),
		reader: reader,
		writer: writer,
	}

	go log.reader.Read(log.rc)
	go log.Process()
	go log.writer.Write(log.wc)
}
