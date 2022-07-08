//
// badgerdbweb is a webserver base GUI for interacting with BadgerDB databases.
//
// For authorship see https://github.com/badarsebard/badgerdbweb
// MIT license is included in repository
//
package main

//go:generate go-bindata-assetfs -o web_static.go web/...

import (
	"flag"
	"fmt"
	badgerbrowserweb "github.com/badarsebard/badgerdbweb/web"
	"github.com/gin-gonic/gin"
	"os"
	"path"

	"github.com/dgraph-io/badger/v3"
	log "github.com/sirupsen/logrus"
)

const version = "v0.0.0"

var (
	showHelp   bool
	db         *badger.DB
	dbName     string
	port       string
	staticPath string
)

func usage(appName, version string) {
	fmt.Printf("Usage: %s [OPTIONS] [DB_NAME]", appName)
	fmt.Printf("\nOPTIONS:\n\n")
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			fmt.Printf("    -%s, -%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
		}
	})
	fmt.Printf("\n\nVersion %s\n", version)
}

func init() {
	// Read the static path from the environment if set.
	dbName = os.Getenv("BADGERDBWEB_DB_NAME")
	port = os.Getenv("BADGERDBWEB_PORT")
	// Use default values if environment not set.
	if port == "" {
		port = "8080"
	}
	// Setup for command line processing
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.StringVar(&dbName, "d", dbName, "Name of the database")
	flag.StringVar(&dbName, "db-name", dbName, "Name of the database")
	flag.StringVar(&port, "p", port, "Port for the web-ui")
	flag.StringVar(&port, "port", port, "Port for the web-ui")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	if showHelp == true {
		usage(appName, version)
		os.Exit(0)
	}

	// If non-flag options are included assume badger db is specified.
	if len(args) > 0 {
		dbName = args[0]
	}

	if dbName == "" {
		usage(appName, version)
		log.Printf("\nERROR: Missing badgerdb name\n")
		os.Exit(1)
	}

	fmt.Print(" ")
	log.Info("starting badgerdb-browser..")

	var err error
	db, err = badger.Open(badger.DefaultOptions(dbName))
	badgerbrowserweb.Db = db

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// OK, we should be ready to define/run web server safely.
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", badgerbrowserweb.Index)

	r.POST("/set", badgerbrowserweb.Set)
	r.POST("/get", badgerbrowserweb.Get)
	r.POST("/deleteKey", badgerbrowserweb.DeleteKey)
	r.POST("/prefixScan", badgerbrowserweb.PrefixScan)

	r.StaticFS("/web", assetFS())

	r.Run(":" + port)
}
