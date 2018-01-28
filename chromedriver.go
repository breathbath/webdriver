// Copyright 2013 Federico Sogaro. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webdriver

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type ChromeSwitches map[string]interface{}

type ChromeDriver struct {
	WebDriverCore
	//The port that ChromeDriver listens on. Default: 9515
	Port int
	//The URL path prefix to use for all incoming WebDriver REST requests. Default: ""
	BaseUrl string
	//The number of threads to use for handling HTTP requests. Default: 4
	Threads int
	//The path to use for the ChromeDriver server log. Default: ./chromedriver.log
	LogPath string
	// Log file to dump chromedriver stdout/stderr. If "" send to terminal. Default: ""
	LogFile string
	// Start method fails if Chromedriver doesn't start in less than StartTimeout. Default 20s.
	StartTimeout time.Duration
}

//create a new service using chromedriver.
//function returns an error if not supported switches are passed. Actual content
//of valid-named switches is not validate and is passed as it is.
//switch silent is removed (output is needed to check if chromedriver started correctly)
func NewChromeDriver(url string, port, threadsCount int, timeout time.Duration) *ChromeDriver {
	d := &ChromeDriver{}
	d.url = url
	d.Port = port
	d.Threads = threadsCount
	d.StartTimeout = timeout
	return d
}

func (d *ChromeDriver) Start() error {
	return nil
}

func (d *ChromeDriver) Stop() error {
	return nil
}

func (d *ChromeDriver) NewSession(desired, required Capabilities) (*Session, error) {
	//id, capabs, err := d.newSession(desired, required)
	//return &Session{id, capabs, d}, err
	session, err := d.newSession(desired, required)
	if err != nil {
		return nil, err
	}
	session.wd = d
	return session, nil
}

func (d *ChromeDriver) Sessions() ([]Session, error) {
	sessions, err := d.sessions()
	if err != nil {
		return nil, err
	}
	for i := range sessions {
		sessions[i].wd = d
	}
	return sessions, nil
}
