package config

const Api_key = "f45565c453514dc9b89702aea30e34ed"
const Page_size = 2
const Amount_of_pages = 25
const Ordering_type = "-released" // aviliable ordering:
// name, released, added, created, updated, rating, metacritic
// possible reverse ordering via -value
const Db_conn_str string = "user=postgres password=1234 dbname=gg host = localhost port = 5432 sslmode=disable"
