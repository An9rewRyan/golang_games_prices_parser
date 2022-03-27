package utils

import (
	"database/sql"
	"fmt"
	"prices_parser/config"
	"strings"

	"github.com/lib/pq"
)

func Connect_database() *sql.DB {
	db, err := sql.Open("postgres", config.Db_conn_str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database connected succefully")
	}
	return db
}

func Create_article_table() {
	Db := Connect_database()
	defer Db.Close()
	create_game_table_query := `create table games(
								game_id serial primary key,
								title varchar(100) not null,
								slug varchar(100) not null,
								release_date timestamp null,
								background_image varchar(250) null,
								ag_rating float4 null,
								playtime integer null,
								metacritic integer null, 
								min_requirenments text null,
								rec_requirenments text null,
								platforms text[],
								genres text[],
								tags text[],
								screenshots text[]
								)`
	_, err := Db.Exec(create_game_table_query)
	if err != nil {
		fmt.Println(err)
	}
}

func Write_game_to_db(title string, slug string, release_date string, background_image string,
	ag_rating float32, playtime int16, metacritic int8, min_requirenments string, rec_requirenments string,
	platforms string, genres string, tags string, screenshots string) {
	Db := Connect_database()
	defer Db.Close()

	platforms_m := strings.Split(platforms[1:len(platforms)-1], ",")
	genres_m := strings.Split(genres[1:len(genres)-1], ",")
	tags_m := strings.Split(tags[1:len(tags)-1], ",")
	screenshots_m := strings.Split(screenshots[1:len(screenshots)-1], ",")

	write_game := fmt.Sprintf(`insert into games 
	(	title, slug, release_date, background_image,
		ag_rating, playtime, metacritic, min_requirenments,
		rec_requirenments, platforms, genres, tags, screenshots)
		values ('%s', '%s', '%s', '%s', %f, %d, %d, '%s', '%s', $1, $2, $3, $4)`,
		title, slug, release_date, background_image, ag_rating, playtime, metacritic, min_requirenments, rec_requirenments)
	fmt.Println(write_game)
	_, err := Db.Exec(write_game, pq.Array(platforms_m), pq.Array(genres_m), pq.Array(tags_m), pq.Array(screenshots_m))
	if err != nil {
		fmt.Println("pq error:", err)
	}
}
