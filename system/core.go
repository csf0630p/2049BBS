package system

import (
	"log"
	"runtime"

	"net/url"
	"strings"

	"github.com/ego008/youdb"
	"github.com/gorilla/securecookie"
	"github.com/terminus2049/2049bbs/util"
	"github.com/weint/config"
)

type MainConf struct {
	HttpPort       int
	HttpsOn        bool
	Domain         string // 若启用https 则该domain 为注册的域名，eg: domain.com、www.domain.com
	HttpsPort      int
	PubDir         string
	ViewDir        string
	Youdb          string
	CookieSecure   bool
	CookieHttpOnly bool
	OldSiteDomain  string
	TLSCrtFile     string
	TLSKeyFile     string
}

type SiteConf struct {
	GoVersion         string
	MD5Sums           string
	Name              string
	Desc              string
	AdminEmail        string
	MainDomain        string // 上传图片后添加网址前缀, eg: http://domian.com 、http://234.21.35.89:8082
	MainNodeIds       string
	MustLoginNodeIds  string
	NotHomeNodeIds    string
	TimeZone          int
	HomeShowNum       int
	PageShowNum       int
	TagShowNum        int
	CategoryShowNum   int
	TitleMaxLen       int
	ContentMaxLen     int
	PostInterval      int
	CommentListNum    int
	CommentInterval   int
	Authorized        bool
	RegReview         bool
	CloseReg          bool
	AutoDataBackup    bool
	AutoGetTag        bool
	GetTagApi         string
	UploadSuffix      string
	UploadImgOnly     bool
	UploadImgResize   bool
	UploadMaxSize     int
	UploadMaxSizeByte int64
}

type AppConf struct {
	Main *MainConf
	Site *SiteConf
}

type Application struct {
	Cf *AppConf
	Db *youdb.DB
	Sc *securecookie.SecureCookie
}

func LoadConfig(filename string) *config.Engine {
	c := &config.Engine{}
	c.Load(filename)
	return c
}

func (app *Application) Init(c *config.Engine, currentFilePath string) {

	mcf := &MainConf{}
	c.GetStruct("Main", mcf)

	// check domain
	if strings.HasPrefix(mcf.Domain, "http") {
		dm, err := url.Parse(mcf.Domain)
		if err != nil {
			log.Fatal("domain fmt err", err)
		}
		mcf.Domain = dm.Host
	} else {
		mcf.Domain = strings.Trim(mcf.Domain, "/")
	}

	scf := &SiteConf{}
	c.GetStruct("Site", scf)
	scf.GoVersion = runtime.Version()
	fMd5, _ := util.HashFileMD5(currentFilePath)
	scf.MD5Sums = fMd5
	scf.MainDomain = strings.Trim(scf.MainDomain, "/")
	log.Println("MainDomain:", scf.MainDomain)
	if scf.TimeZone < -12 || scf.TimeZone > 12 {
		scf.TimeZone = 0
	}
	if scf.UploadMaxSize < 1 {
		scf.UploadMaxSize = 1
	}
	scf.UploadMaxSizeByte = int64(scf.UploadMaxSize) << 20

	app.Cf = &AppConf{mcf, scf}
	db, err := youdb.Open(mcf.Youdb)
	if err != nil {
		log.Fatalf("Connect Error: %v", err)
	}
	app.Db = db

	// set main node
	db.Hset("keyValue", []byte("main_category"), []byte(scf.MainNodeIds))

	app.Sc = securecookie.New(securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32))
	//app.Sc.SetSerializer(securecookie.JSONEncoder{})

	log.Println("youdb Connect to", mcf.Youdb)
}

func (app *Application) Close() {
	app.Db.Close()
	log.Println("db cloded")
}
