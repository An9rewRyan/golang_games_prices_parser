package config

import (
	"os"

	"github.com/tebeka/selenium"
)

const Selenium_path = "selenium/selenium-server.jar"
const Gecko_driver_path = "selenium/geckodriver"
const Chrome_driver_path = "selenium/chromedriver"
const Selenium_port = 8080

var Selenium_opts = []selenium.ServiceOption{
	selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in.
	selenium.GeckoDriver(Gecko_driver_path), // Specify the path to GeckoDriver in order to use Firefox.
	selenium.Output(os.Stderr),
}
