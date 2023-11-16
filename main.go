package main

import (
	"database/sql"
	"errors"
	"fmt"
	"image/color"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	dbpath, err := os.UserHomeDir()
    if err != nil {
        log.Fatal("Could not open user home directory", err )
    }
    fmt.Println( dbpath )
	dbpath += "/.firmware-checker/"
	dbname := "firmware-checker.db"
	dbparams := "file:" + dbpath + dbname + "?cache=shared&mode=rwc"
	db, err := sql.Open("sqlite3", dbparams)
	if err != nil {
		log.Fatal("cant access database:", err)
	}
	//Check if we can access the database.
	err = db.Ping()
	//If we cant then maybe it does not yet exist so lets make it.
	if err != nil {
		if _, err := os.Stat(dbpath); err == nil {
			// Folder exists only create the db file
			file, err := os.OpenFile(dbpath + dbname, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
			if err != nil {
				log.Fatal("Cant create database", err)
			}
			file.Close()
		  } else if errors.Is(err, os.ErrNotExist) {
			//Create the folder then the db file
			err := os.Mkdir(dbpath, 0700)
			if err != nil {
				log.Fatal("Unable to make directory", err)
			}
			file, err := os.OpenFile(dbpath + dbname, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
			if err != nil {
				log.Fatal("Cant create database", err)
			}
			file.Close()
		  } else {
			log.Fatal("Unable to create database", err)
		  }
	}
	//@todo create tables if not exist
	prepStmt := `CREATE TABLE IF NOT EXISTS devices (device TEXT PRIMARY KEY, url TEXT NOT NULL, lastcheck INTEGER)`
	statement, err := db.Prepare(prepStmt)
	if err != nil {
		log.Fatal("could not create device table", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal("could not create device table", err)
	}
	prepStmt = `CREATE TABLE IF NOT EXISTS settings (id INTEGER PRIMARY KEY AUTOINCREMENT, cron TEXT, version TEXT)`
	statement, err = db.Prepare(prepStmt)
	if err != nil {
		log.Fatal("could not create settings table", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal("could not create settings table", err)
	}
}

func main() {
	mainApp := app.New()
	//@todo set version dynamically
	applicationWindow := mainApp.NewWindow("Firmware Checker v1.0.0-beta")
	applicationWindow.Resize(fyne.NewSize(1024, 768))

	tabs := container.NewAppTabs(
		container.NewTabItem("Devices", deviceList()),
		container.NewTabItem("URL Settings", urlList()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	applicationWindow.SetContent(tabs)
	applicationWindow.ShowAndRun()
}

func deviceList() (fyne.CanvasObject) {
	//todo get the current list of devices from the database
	//@todo bind variable to the widget
	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	//grid := container.New(layout.NewGridLayout(1), text1, text2, text3)
	grid := container.New(layout.NewVBoxLayout(), text1, text2, text3)
	return grid
	//return widget.NewLabel("Device List")
}

func urlList() (fyne.CanvasObject) {
	//todo get the current list of urls from the database
	//@todo bind variable to the widget
	return widget.NewLabel("URL List")
}