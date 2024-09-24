package mocker

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/mocker/utils"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type MServer struct {
	host    string
	port    uint64
	cors    bool
	data    []byte
	r       *gin.Engine
	workDir string
}

func NewMServer(workDir string) (*MServer, error) {
	file, err := os.OpenFile(filepath.Join(workDir, "mocker.json"), os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// check router
	if !gjson.GetBytes(data, "router").Exists() {
		return nil, errors.New("router not found")
	}

	return &MServer{
		workDir: workDir,
		r:       gin.Default(),
		data:    data,
		host:    utils.JsonValueWithDefault[string](data, "host", "0.0.0.0"),
		port:    utils.JsonValueWithDefault[uint64](data, "port", 8080),
		cors:    utils.JsonValueWithDefault[bool](data, "host", false),
	}, nil
}

func (s MServer) fileHandler(router string, value gjson.Result) error {
	res := value.Get("dir")
	if !res.Exists() {
		return fmt.Errorf("router[%s],dir not found", router)
	}
	dir := strings.Replace(res.String(), "${WorkDir}", s.workDir, 1)
	if !utils.DirExists(dir) {
		return fmt.Errorf("router[%s],dir[%s] not exist", router, dir)
	}
	s.r.StaticFS(router, http.Dir(dir))
	return nil
}

func (s MServer) apiServer(value gjson.Result) error {
	return nil
}

func (s MServer) apiGroupServer(value gjson.Result) error {
	return nil
}

func (s MServer) crudHandler(value gjson.Result) error {
	return nil
}

func (s MServer) proxyServer(value gjson.Result) error {
	return nil
}

func (s MServer) Setup() (err error) {
	gjson.GetBytes(s.data, "router").ForEach(func(key, value gjson.Result) bool {
		var rType = value.Get("type").String()
		switch RouterType(rType) {
		case RouterTypeFS:
			err = s.fileHandler(key.String(), value)
		case RouterTypeCrud:
			err = s.crudHandler(value)
		case RouterTypeAPI:
			err = s.apiServer(value)
		case RouterTypeProxy:
			err = s.proxyServer(value)
		case RouterTypeGroup:
			err = s.apiGroupServer(value)
		default:
			err = fmt.Errorf("unknown router type:%s", rType)
		}
		return err == nil
	})
	return s.r.Run(fmt.Sprintf("%s:%d", s.host, s.port))
}
