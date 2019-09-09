package utils

import (
	"archive/zip"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Pdf struct {
	Path    string //保存路径
	Title   string //pdf标题
	Content string //pdf正文
}

func (p Pdf) Mkdir() bool {
	if _, err := os.Stat(p.Path); err != nil {
		return true
	}
	err := os.Mkdir(p.Path, os.ModePerm)
	if err != nil {
		log.Printf("file path is exist...")
		return false
	}
	return true
}

func (p Pdf) isExist() bool {
	info, err := os.Stat(p.Path)
	if os.IsExist(err) {
		return true
	}
	//文件夹
	if info.IsDir() {
		return true
	}

	//文件
	if info.Size() == 0 {
		return false
	}

	return false
}

func (p Pdf) HtmlToPdf() {
	if runtime.GOOS == "windows" {
		wkhtmltopdf.SetPath("tools/utils/pdf.exe")
	}
	pdf, err := wkhtmltopdf.NewPDFGenerator()

	pdf.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(p.Content)))
	err = pdf.Create()
	if err != nil {
		log.Fatal(err)
	}
	p.Mkdir()
	err = pdf.WriteFile(fmt.Sprintf("%s/%s.%s", p.Path, p.Title, ".pdf"))
	if err != nil {
		log.Fatal(err)
	}

}

func Zip(srcFile string, destZip string) error {
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
