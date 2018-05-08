package confutil

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const delimiter = "_"

type Config struct {
	confMap map[string]string
	node string
}

// 根据文件路径初始文config对象
// @param path
func InitConfig(path string) (conf * Config) {
	c := new(Config)
	c.confMap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.node = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.node) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		first := strings.TrimSpace(s[:index])
		if len(first) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.node + delimiter + first
		c.confMap[key] = strings.TrimSpace(second)
	}
	return c
}

func (c Config) Read(node, key string) string {
	key = node + delimiter + key
	v, found := c.confMap[key]
	if !found {
		return ""
	}
	return v
}
