package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type VideoInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    Data   `json:"data"`
}
type Rights struct {
	Bp            int `json:"bp"`
	Elec          int `json:"elec"`
	Download      int `json:"download"`
	Movie         int `json:"movie"`
	Pay           int `json:"pay"`
	Hd5           int `json:"hd5"`
	NoReprint     int `json:"no_reprint"`
	Autoplay      int `json:"autoplay"`
	UgcPay        int `json:"ugc_pay"`
	IsCooperation int `json:"is_cooperation"`
	UgcPayPreview int `json:"ugc_pay_preview"`
	NoBackground  int `json:"no_background"`
	CleanMode     int `json:"clean_mode"`
	IsSteinGate   int `json:"is_stein_gate"`
}
type Owner struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}
type Stat struct {
	Aid        int    `json:"aid"`
	View       int    `json:"view"`
	Danmaku    int    `json:"danmaku"`
	Reply      int    `json:"reply"`
	Favorite   int    `json:"favorite"`
	Coin       int    `json:"coin"`
	Share      int    `json:"share"`
	NowRank    int    `json:"now_rank"`
	HisRank    int    `json:"his_rank"`
	Like       int    `json:"like"`
	Dislike    int    `json:"dislike"`
	Evaluation string `json:"evaluation"`
	ArgueMsg   string `json:"argue_msg"`
}
type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}
type Pages struct {
	Cid       int       `json:"cid"`
	Page      int       `json:"page"`
	From      string    `json:"from"`
	Part      string    `json:"part"`
	Duration  int       `json:"duration"`
	Vid       string    `json:"vid"`
	Weblink   string    `json:"weblink"`
	Dimension Dimension `json:"dimension"`
}
type Subtitle struct {
	AllowSubmit bool          `json:"allow_submit"`
	List        []interface{} `json:"list"`
}
type UserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"`
}
type Data struct {
	Bvid      string    `json:"bvid"`
	Aid       int       `json:"aid"`
	Videos    int       `json:"videos"`
	Tid       int       `json:"tid"`
	Tname     string    `json:"tname"`
	Copyright int       `json:"copyright"`
	Pic       string    `json:"pic"`
	Title     string    `json:"title"`
	Pubdate   int       `json:"pubdate"`
	Ctime     int       `json:"ctime"`
	Desc      string    `json:"desc"`
	State     int       `json:"state"`
	Duration  int       `json:"duration"`
	Rights    Rights    `json:"rights"`
	Owner     Owner     `json:"owner"`
	Stat      Stat      `json:"stat"`
	Dynamic   string    `json:"dynamic"`
	Cid       int       `json:"cid"`
	Dimension Dimension `json:"dimension"`
	NoCache   bool      `json:"no_cache"`
	Pages     []Pages   `json:"pages"`
	Subtitle  Subtitle  `json:"subtitle"`
	UserGarb  UserGarb  `json:"user_garb"`
}

type PlayURLInfo struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	TTL     int     `json:"ttl"`
	Data    URLData `json:"data"`
}
type Durl struct {
	Order     int         `json:"order"`
	Length    int         `json:"length"`
	Size      int         `json:"size"`
	Ahead     string      `json:"ahead"`
	Vhead     string      `json:"vhead"`
	URL       string      `json:"url"`
	BackupURL interface{} `json:"backup_url"`
}
type SupportFormats struct {
	Quality        int    `json:"quality"`
	Format         string `json:"format"`
	NewDescription string `json:"new_description"`
	DisplayDesc    string `json:"display_desc"`
	Superscript    string `json:"superscript"`
}
type URLData struct {
	From              string           `json:"from"`
	Result            string           `json:"result"`
	Message           string           `json:"message"`
	Quality           int              `json:"quality"`
	Format            string           `json:"format"`
	Timelength        int              `json:"timelength"`
	AcceptFormat      string           `json:"accept_format"`
	AcceptDescription []string         `json:"accept_description"`
	AcceptQuality     []int            `json:"accept_quality"`
	VideoCodecid      int              `json:"video_codecid"`
	SeekParam         string           `json:"seek_param"`
	SeekType          string           `json:"seek_type"`
	Durl              []Durl           `json:"durl"`
	SupportFormats    []SupportFormats `json:"support_formats"`
}

func getVideoInfo(bid string) VideoInfo {
	url := "http://api.bilibili.com/x/web-interface/view?bvid=" + bid
	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", url, nil)
	//req, _ := http.NewRequest("POST", url, strings.NewReader("bvid%3dBV117411r7R1"))
	//req.Header.Set("User-Agent", "curl/7.64.1")
	//req.Header.Set("Host", "api.bilibili.com")
	//resp, err := client.Do(req)
	resp, err := http.Get(url)
	//fmt.Println(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", content)
	//	ct := string(content)
	var data VideoInfo
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("%v", data)
	return data
}

func getDownloadURLs(bid, cid string) []string {

	url := fmt.Sprintf("http://api.bilibili.com/x/player/playurl?cid=%s&qn=0&type=&otype=json&fourk=1&bvid=%s&fnver=0&fnval=1", cid, bid)
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", content)
	var data PlayURLInfo
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	playURLs := []string{}
	for _, item := range data.Data.Durl {
		playURLs = append(playURLs, item.URL)
	}
	return playURLs
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeFile(URLs []string, filename string) {
	filename = "./" + filename + ".mp4"
	// if checkFileIsExist(filename) { //如果文件存在

	// 	f, err = f, err
	// 	fmt.Println("文件存在")
	// } else {
	// 	f, err = f, err
	// 	fmt.Println("文件不存在")
	// }

	// defer f.Close()
	for _, url := range URLs {
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile(filename, body, 0600)

	}

}

func main() {
	bid := "BV125411H7GT"
	videoinfo := getVideoInfo(bid)
	//fmt.Printf("%v", videoinfo)
	//videos := fmt.Println(videoinfo.Data.Videos)
	pages := videoinfo.Data.Pages
	//fmt.Println(pages[0])
	//var cids []int
	for _, page := range pages {
		var fileName string
		if videoinfo.Data.Videos == 1 {
			fileName = videoinfo.Data.Title
		} else {
			fileName = page.Part
		}
		cid := strconv.Itoa(page.Cid)
		playURLs := getDownloadURLs(bid, cid)

		writeFile(playURLs, fileName)

	}
	//fmt.Println(cids)
	return
}
