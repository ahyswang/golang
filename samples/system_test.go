package samples

import (
	"io/ioutil"
	"os"
	"testing"
)

//测试目录遍历
func ListDir(dirPath string, files *[]string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	PathSep := string(os.PathSeparator)
	for _, fi := range dir {
		fullPath := dirPath + PathSep + fi.Name()
		if fi.IsDir() {
			ListDir(fullPath, files)
		} else {
			*files = append(*files, fullPath)
		}
	}
	return nil
}

func TestDir(t *testing.T) {
	files := make([]string, 0)
	if ListDir("./", &files) != nil {
		t.Fatal("枚举文件失败")
	}
	t.Log(files)
}
