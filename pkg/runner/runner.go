package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Run(database string, bin string, dir string, backup int, verbose int) error {

	var log = func(text string) {
		// fmt.Sprintf("verbose:%d", verbose)
		// log.
		if verbose > 0 {
			// log.DefaultLogger.Println(text)
			log.Println(text)
		}
	}
	var logArr = func(text string, arr []string) {
		log(text)
		for _, v := range arr {
			log(fmt.Sprintf("  %s", v))
		}
	}
	var err error
	log("任务开始")
	cmd := fmt.Sprintf("su - postgres -c \"%s -c %s -p %d \"| gzip > %s/%s__%s.gz",
		bin,      // bin路劲
		database, // 数据库
		5432,     // 端口号
		dir,      // 备份目录
		database, // 数据库,
		time.Now().Format("20060102150405"),
	)
	log(fmt.Sprintf("cmd:%s", cmd))
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	if err != nil {
		return err
	}
	log(fmt.Sprintf("结果:%s", output))
	glob_pattern := fmt.Sprintf("%s/%s__*.gz", dir, database)
	log(fmt.Sprintf("glob_pattern:%s", glob_pattern))
	matches, err := filepath.Glob(glob_pattern)
	if err != nil {
		return err
	}
	// log.
	logArr("文件目录", matches)
	var ffileKey = func(filePath string) (int64, error) {
		nameArr := strings.Split(filePath, "/")
		name := nameArr[len(nameArr)-1]
		pieces := strings.Split(name, ".")
		withoutExt := pieces[0]
		PrefixPieces := strings.Split(withoutExt, "__")
		if len(PrefixPieces) < 2 {
			return int64(0), nil
		}
		return strconv.ParseInt(PrefixPieces[1], 10, 64)
	}
	sort.Slice(matches, func(i, j int) bool {
		key1, _ := ffileKey(matches[i])
		key2, _ := ffileKey(matches[j])
		return key1 > key2
	})
	logArr("排序文件目录", matches)
	len_matches := len(matches)
	log(fmt.Sprintf("文件数目:%d, 最大备份数目:%d", len_matches, backup))
	if len_matches > backup {
		remove_file_list := matches[backup:]
		logArr("要移除的文件列表", remove_file_list)
		for _, v := range remove_file_list {

			err = os.Remove(v)
			log(fmt.Sprintf("移除路劲:%s", v))
			if err != nil {
				return err
			}
		}
	}
	return nil

}
