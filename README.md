## Welcome to the golang web scraper 

![](https://img.shields.io/badge/golang-1.17-52a7f7) ![](https://img.shields.io/badge/-selenium-ff69b4) ![](https://img.shields.io/badge/-postgresql-3294f0) ![](https://img.shields.io/badge/-docker-32c7f0) ![](https://img.shields.io/badge/-htmlquery-4f75ff)


>  Work in progress... *(i mean it)*

---

### How to launch: 
(the only working part for now is ag_api_getter)
 1. Download dependencies from go mod
 2. Launch local postgresql with given configs from confif.go
 3. Create game table via uncommenting special function in main.go (with first launch)
 4. Run main.go

***
**What is used for what:**
 - **Go** - the basis of the parser, sends requests, searches for data *(+ specifically htmlquery*)
 - **Selenium** (via golang) - handles javascript on found pages if necessary

***

**Other projects using this parser:**
 - **[Gamers Gazette](https://github.com/An9rewRyan/Gamers-Gazette)** - web platform based on this parser
 - **[Golang game news_parser](https://github.com/An9rewRyan/golang_game_news_parser)** - game parser which patterns were used here (also part of GG)

---

Contact me if you have some questions or suggestions via:
 - Telegram: **[@Michael_J_Goldberg](https://t.me/Michael_J_Goldberg)**
 - Vk - **[vk.com/mj_the_reviewer](https://vk.com/mj_the_reviewer)**
 - Discord - **[YUUJIRO HANMA#6379](https://discordapp.com/users/389483338865311745/)**

***

**What resources does the parser use:**
 - **[dtf.ru](https://dtf.ru/)**
 - **[igromania.ru](https://www.igromania.ru/)**
 - **[kanobu.ru](https://kanobu.ru/videogames/)**
 - **[playground.ru](https://www.playground.ru/)**
 - **[stopgame.ru](https://stopgame.ru/)**
 - **[vgtimes.ru](https://vgtimes.ru/)**

 
