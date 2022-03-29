package main

import (
	"fmt"
	"name_parser/config"

	"github.com/tebeka/selenium"
)

func main() {
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(config.Selenium_path, config.Selenium_port, config.Selenium_opts...)
	if err != nil {
		fmt.Println(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", config.Selenium_port))
	if err != nil {
		fmt.Println(err)
	}
	defer wd.Quit()
	if err := wd.Get("https://kanobu.ru/"); err != nil {
		panic(err)
	}
	fmt.Println(wd.PageSource())

}
