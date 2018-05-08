package api

import (
	"crypto/md5"
	"database/sql"
	"transurl/dbutil"
	"fmt"
	"sync"
	"transurl/util"
	"hash/crc32"
)

var DB *sql.DB = dbutil.GetConn()

var urlPre = util.GetInstanceConf().Read("Url", "PREFIX_URL_ADDRESS")


// 定义数据库表tb_url信息
type TbUrl struct {
	Id int64 `db:"id"`
	ShortUrl string `db:"short_url"`
	OriginUrl string `db:"origin_url"`
	code string `db:"url_code"`
}

// 长URL转换为短URL
// @param longUrl
// @return shortUrl, err
func LongToShort(longUrl string) (shortUrl string, err error) {

	// 获取长连接的hashcode值
	hashcode := crc32.ChecksumIEEE([]byte (longUrl))
	code := transTo62(int64(hashcode))
	// 长URL过长，将其使用md5加密缩短
	urlMd5 := md5.Sum([]byte (longUrl))
	md5code := fmt.Sprintf("%x", urlMd5)
	code += md5code
	// 对下面的代码进行同步操作，避免大量并发导致同一个长连接对应多个短连接
	var mu = &sync.Mutex{}
	mu.Lock() // 加锁
	res := DB.QueryRow("select id, origin_url, short_url, url_code from tb_url where url_code=?", code)
	url := &TbUrl{}
	err = res.Scan(&url.Id, &url.OriginUrl, &url.ShortUrl, &url.code)
	//log.Println("short url ==", url.ShortUrl)
	if err == sql.ErrNoRows {
		err = nil
		// 此时说明数据库中没有数据，需要新添加数据
		shortUrl, errRet := generateShortUrl(longUrl, code)
		if errRet == nil {
			err = errRet
			return shortUrl, err
		}
	}
	mu.Unlock() // 释放锁
	if err != nil {
		fmt.Println("DB.query fail,", err.Error())
		return
	}

	return url.ShortUrl, err
}

// 将短的URL转化为长的URL
// @param shortUrl
// @ return longUrl, err
func ShortToLongUrl(shortUrl string) (originUrl string, err error) {
	row := DB.QueryRow("select origin_url from tb_url where short_url=?", shortUrl)
	err = row.Scan(&originUrl)
	if err != nil {
		fmt.Println("ShortToLongUrl fail, fail info: ", err.Error())
	}
	return originUrl, err
}

// 生成短URL
// 根据ID，将ID转换为62进制的字符串返回
func generateShortUrl(longUrl string, code string) (shortUrl string, err error){
	result, err := DB.Exec("insert into tb_url(origin_url,url_code) values(?,?)", longUrl, code)
	if err != nil {
		fmt.Println("generate short url fail ", err.Error())
		panic(err)
	}
	id, _ := result.LastInsertId()

	// 将ID转换为62进制数
	shortUrl = transTo62(id)
	shortUrl = urlPre + shortUrl

	_, e := DB.Exec("update tb_url set short_url=? where id=?", shortUrl, id)
	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
	return shortUrl, e
}

// 将十进制转换为62进制   0-9a-zA-Z 六十二进制
func transTo62(id int64) string {
	// 1 -- > 1  10-- > a  61-- > Z
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for {
		var result byte
		number := id % 62
		result = charset[number]
		var tmp []byte
		tmp = append(tmp,result)
		shortUrl = append(tmp, shortUrl...)
		id = id / 62
		if id == 0{
			break
		}
	}
	//fmt.Println(string(shortUrl))
	return string(shortUrl)
}
