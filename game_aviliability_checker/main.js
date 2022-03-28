const puppeteer = require("puppeteer");
const fs = require("fs");
const config = require("/config")

//this huge thing has necessary args for increasing performance 


async function get_js_rendered_page (link) {
    console.log(link)
    const browser = await puppeteer.launch({
      userDataDir: './cache',
      args: minimal_args
    });
    page = await browser.newPage();
    await page.setDefaultNavigationTimeout(0); 
    await page.goto(link, {
      waitUntil: 'load',
    });

    let bodyHTML = await page.evaluate(() => document.documentElement.outerHTML);
    
    await browser.close();
    return bodyHTML;
}