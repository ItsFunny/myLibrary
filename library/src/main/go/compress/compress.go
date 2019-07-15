/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-08 14:36 
# @File : compress.go
# @Description : 
*/
package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
)

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}
	return buffer.Bytes(), nil
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	var out []byte
	if err != nil {
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}


type ZstdEncoder struct {
	File *os.File
	ZstdWriter *zstd.Writer
}



//level是压缩级别：1~19，越大越慢。默认为3

func NewZstdEncoder(filename string, level int) (*ZstdEncoder,error) {
	f,err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	w := zstd.NewWriterLevel(f,level)
	return &ZstdEncoder{File:f,ZstdWriter:w},nil
}

func (s *ZstdEncoder) Close() {

	s.ZstdWriter.Close()
	s.File.Close()
}

func ZstdCompressFile(dst,src string, compress_level int) error {
	//log.Println("Zstd ", src, "->", dst)
	s,err := NewZstdEncoder(dst,compress_level)
	defer s.Close()
	if err != nil {
		return err
	}
	src_file,err := os.Open(src)
	if err != nil {
		return err
	}
	defer src_file.Close()
	io.Copy(s.ZstdWriter, src_file)

	return nil
}

func ZstdDecompressFile(dst, src string) error {
	//log.Println("Zstd ", src, "->", dst)
	in_file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in_file.Close()
	out_file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out_file.Close()

	zst_reader := zstd.NewReader(in_file)
	io.Copy(out_file, zst_reader)

	return nil
}
