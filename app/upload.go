package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"
	"upload/context"
	"upload/context/global"

	"os"

	"net/textproto"

	"github.com/jlaffaye/ftp"
)

func catch(c *context.Context, err error) {
	log.Println(err)
	c.JsonError(-1, "error, view log for detail")
}

func Upload(c *context.Context) {
	module := c.R.FormValue("module")
	userid := c.R.FormValue("openid")
	unique := false
	uniqueStr := c.R.FormValue("unique")
	if uniqueStr == "1" {
		unique = true
	}
	if userid == "" {
		c.JsonError(-1, "openid invalid")
		return
	}

	file, handler, err := c.R.FormFile("file")

	if err != nil {
		log.Println("get file error", err)
		c.JsonError(-1, "get file error")
		return
	}
	defer file.Close()

	// validate file type
	filetype := handler.Header.Get("Content-Type")
	inAllow := false
	for _, allowType := range global.Conf.AllowFileType {
		if filetype == allowType {
			inAllow = true
			break
		}
	}
	if !inAllow {
		c.JsonError(-2, fmt.Sprintf("forbidden file type %s", filetype))
		return
	}

	// limit file size

	data, err := ioutil.ReadAll(file)
	if err != nil {
		catch(c, err)
		return
	}

	filesize := len(data)
	if filesize > global.Conf.MaxFileSize {
		c.JsonError(-2, fmt.Sprintf("exceed max file size %d", filesize))
		return
	}

	extName := path.Ext(handler.Filename)

	var newFilename string
	if unique {
		newFilename = userid + extName
	} else {
		now := time.Now()
		currentDate := now.Format("20060102")
		newFilename = currentDate + "_" + userid + "_" + fmt.Sprintf("%d", now.UnixNano()/1e6) + extName
	}

	// 未配置FTP服务器, 直接传到UploadDir目录
	if global.Conf.FTPServer == 0 {
		if module != "" {
			newFilename = module + "/" + newFilename
		}

		newfile := path.Join(global.Conf.UploadDir, newFilename)
		dir := path.Dir(newfile)
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			} else {
				catch(c, err)
			}
		}

		err = ioutil.WriteFile(newfile, data, 0644)
		if err != nil {
			catch(c, err)
			return
		}

		uploadUrl := global.Conf.HTTPServerURL + newFilename
		c.JsonSuccess(uploadUrl)
		return
	}

	// 连接FTP, 上传文件
	con, err := ftp.Connect(global.Conf.FTPServerAddr)
	if err != nil {
		catch(c, err)
		return
	}

	err = con.Login(global.Conf.FTPUsername, global.Conf.FTPPassword)
	if err != nil {
		catch(c, err)
		return
	}
	defer con.Logout()

	err = con.ChangeDir(global.Conf.FTPUploadPath)
	if err != nil {
		catch(c, err)
		return
	}
	fullFilename := newFilename
	if module != "" {
		fullFilename = module + "/" + newFilename
		err = con.ChangeDir(module)
		if err != nil {
			if err.(*textproto.Error).Code == ftp.StatusFileUnavailable {
				err = con.MakeDir(module)
				if err != nil {
					catch(c, err)
					return
				}
				err = con.ChangeDir(module)
				if err != nil {
					catch(c, err)
					return
				}
			} else {
				catch(c, err)
				return
			}
		}

	}

	err = con.Stor(newFilename, bytes.NewBuffer(data))
	if err != nil {
		catch(c, err)
		return
	}

	uploadUrl := global.Conf.HTTPServerURL + fullFilename
	c.JsonSuccess(uploadUrl)
}
