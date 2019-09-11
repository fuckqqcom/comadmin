package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	text := `                    <p style="text-align: center;"><img class="rich_pages" data-copyright="0" data-ratio="0.4264264264264264" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_jpg/azXQmS1HA7ndCbgcUE9RfJTBU3XcC9vJ0eC4BwzD90OoGw4y3r9uA5sjT2A2POF0Gon6BPLncJRUfnkhoYEVAQ/640?wx_fmt=jpeg" data-type="jpeg" data-w="666" style=""/></p><p><br/></p><p><span style="color: rgb(0, 122, 170);"><strong><span style="font-size: 18px;">联邦快递所称将涉华为公司快件转至美国系“误操作”与事实不符</span></strong></span></p><p><br/></p><p>近期，国家有关部门依法对联邦快递（中国）有限公司未按名址投递快件一案实施调查发现，联邦快递关于将涉华为公司快件转至美国系“误操作”的说法与事实不符。另发现联邦快递涉嫌滞留逾百件涉华为公司进境快件。调查期间，还发现联邦快递其他违法违规线索。国家有关部门将秉持全面、客观、公正的原则，继续依法深入开展调查工作。<br/></p><p><br/></p><section class="" style="max-width: 100%;letter-spacing: 0.544px;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="1" data-tools="新媒体排版" data-id="12974" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section mpa-from-tpl="t" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section data-id="94324" mpa-from-tpl="t" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section data-width="100%" mpa-from-tpl="t" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section data-width="100%" mpa-from-tpl="t" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section data-style="line-height:24px;color:rgb(89, 89, 89); font-size:16px;" mpa-from-tpl="t" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="5" data-tools="新媒体排版" data-id="2239845" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="5" data-tools="新媒体排版" data-id="1075027" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="5" data-tools="新媒体排版" data-id="750326" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="5" data-tools="新媒体排版" data-id="453317" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="5" data-tools="新媒体排版" data-id="341631" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section class="" data-style-type="5" data-tools="新媒体排版" data-id="68900" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><span style="max-width: 100%;font-size: 14px;box-sizing: border-box !important;overflow-wrap: break-word !important;"><span style="max-width: 100%;color: rgb(136, 136, 136);box-sizing: border-box !important;overflow-wrap: break-word !important;">来源：新华社</span></span></section><section class="" data-style-type="5" data-tools="新媒体排版" data-id="68900" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><section data-role="outer" label="Powered by 135editor.com" style="max-width: 100%;box-sizing: border-box !important;overflow-wrap: break-word !important;"><hr style="max-width: 677px;box-sizing: border-box;letter-spacing: 0.54px;overflow-wrap: break-word !important;"/><p style="max-width: 100%;min-height: 1em;box-sizing: border-box !important;overflow-wrap: break-word !important;"><span style="max-width: 100%;font-size: 14px;color: rgb(136, 136, 136);box-sizing: border-box !important;overflow-wrap: break-word !important;">监制：李代祥</span></p><p style="max-width: 100%;min-height: 1em;box-sizing: border-box !important;overflow-wrap: break-word !important;"><span style="max-width: 100%;font-size: 14px;color: rgb(136, 136, 136);box-sizing: border-box !important;overflow-wrap: break-word !important;">编辑：李昂、陈子夏</span></p></section></section></section></section></section></section></section></section><p style="max-width: 100%;min-height: 1em;box-sizing: border-box !important;overflow-wrap: break-word !important;"><img class="__bg_gif " data-ratio="0.244" data-type="gif" data-before-oversubscription-url="https://mmbiz.qpic.cn/mmbiz_gif/azXQmS1HA7kekPnuE77wvuibxNz3qSuvpufX3QAxcBUSr9ibyx0F30WDHQMu5ux0nn21EFCiciaT3ibeJoEKic3aN9kQ/640?wx_fmt=gif" data-w="750" data-copyright="0" data-backw="677" data-backh="165" width="100%" data-src="https://mmbiz.qpic.cn/mmbiz_gif/azXQmS1HA7kekPnuE77wvuibxNz3qSuvpufX3QAxcBUSr9ibyx0F30WDHQMu5ux0nn21EFCiciaT3ibeJoEKic3aN9kQ/640?wx_fmt=gif" style="max-width: 677px;box-sizing: border-box;letter-spacing: 0.54px;width: 100%;overflow-wrap: break-word !important;visibility: visible !important;height: auto;"/></p><p style="max-width: 100%;min-height: 1em;letter-spacing: 0.544px;text-align: right;box-sizing: border-box !important;overflow-wrap: break-word !important;"><strong style="max-width: 100%;letter-spacing: 0.54px;font-size: 20px;color: rgb(0, 122, 170);box-sizing: border-box !important;overflow-wrap: break-word !important;">关注！<img class="__bg_gif" data-ratio="3" data-type="gif" data-w="300" data-copyright="0" width="30" data-src="https://mmbiz.qpic.cn/mmbiz_gif/azXQmS1HA7kw6CtFWdFYz6gIZyKzI8aXNx4zZ3ADUickzwAvGicPCyS2FvbRPmHntQibyDSdHuBibMhH6JiaIiabWnsQ/640?wx_fmt=gif" style="max-width: 677px;box-sizing: border-box;text-align: justify;overflow-wrap: break-word !important;visibility: visible !important;width: 30px !important;"/></strong></p></section></section></section></section></section></section></section></section></section>`

	fmt.Println(strings.TrimSpace(strings.ReplaceAll(text, `data-src="`, `src="`)))
	//src := "tools/utils"
	//dest := "sync.zip"
	//Zip(src, dest)
}

func read(c chan bool, i int) {
	fmt.Printf(" go func : %d \n", i)
	<-c
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
