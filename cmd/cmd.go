package cmd

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type GlobalConfigType struct {
	Port               string        `yaml:"port"`
	DbUrl              string        `yaml:"dbUrl"`
	JwtSecret          string        `yaml:"jwtSecret"`
	JwtExp             time.Duration `yaml:"jwtExp"`
	TokenRefreshMinute time.Duration `yaml:"tokenRefreshMinute"`
	PageMin            int           `yaml:"pageMin"`
	PageMax            int           `yaml:"pageMax"`
	Nointer            bool          `yaml:"nointer"`
	Image              string        `yaml:"image"`
	EnableSqlLog       bool          `yaml:"enableSqlLog"`
}

var Config GlobalConfigType

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ConfigInit() {
	filename := "app.yml"
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", filename)
	if !FileExists(appConfigPath) {
		panic("config file not exits..")
	}

	yamlFile, err := ioutil.ReadFile(appConfigPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", Config)

}

type CmdMap map[string]*Cmd
type CmdHandler func([]string)

var cmdMap = make(CmdMap)

type Cmd struct {
	CmdStr  string
	Handler CmdHandler
	Usage   string
}

func (c Cmd) String() string {
	return fmt.Sprintf("command: %s", c.CmdStr)
}

func InstallCmd(str string, usage string, handler CmdHandler) *Cmd {
	r := &Cmd{CmdStr: str,
		Handler: handler, Usage: usage}
	cmdMap[str] = r
	return r
}

func Tip(s string) {
	log.Println(s)
	os.Exit(1)
}

func FetchArgs(lst []string) []string {
	for i, v := range lst {
		if len(v) > 1 && v[:1] == "-" {
			return lst[:i]
		}
	}
	return lst
}

func usage() {
	fmt.Println(`usage: [app] [cmd]`)
	for _, cmd := range cmdMap {
		fmt.Println(cmd)
	}
}

func (c *Cmd) exec(args []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(c.Usage)
			os.Exit(-1)
		}
	}()
	c.Handler(args)

}

func Exec() {
	args := os.Args
	if len(args) == 1 {
		usage()
	}
	args = args[1:]
	for i, arg := range args {
		if len(arg) > 1 && arg[:1] == "-" {
			if cmd, ok := cmdMap[arg[1:]]; ok {
				args = FetchArgs(args[i+1:])
				cmd.exec(args)
			} else {
				Tip("cmd not found...")
			}
		}
	}
}
