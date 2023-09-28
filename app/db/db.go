package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	// CREATE VAR AND INTERFACE FOR THAT
	database *sql.DB
)

func NewDatabase() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "mysecretpassword"
	dbname := "postgres"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Could not open connection to the database")
		panic(err)
	}
	//defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Printf("Could not ping to the database")
		panic(err)
	}

	fmt.Println("Connected to the database!")

	setupDB(db)

	database = db
}

func setupDB(db *sql.DB) {
	// USER
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		cpf TEXT,
		password TEXT,
		numtrips TEXT,
		bio TEXT
	);`

	_, err := db.Exec(createUserTable)
	if err != nil {
		fmt.Printf("Could not create user table")
		panic(err)
	}

	// POST
	createPostTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title TEXT,
		content TEXT,
		postdate TIMESTAMP,
		userid INT
	);`

	_, err = db.Exec(createPostTable)
	if err != nil {
		fmt.Printf("Could not create post table")
		panic(err)
	}

	createPostUserFK := `ALTER TABLE posts DROP CONSTRAINT IF EXISTS posts_fk;
	ALTER TABLE posts
    ADD CONSTRAINT posts_fk FOREIGN KEY ( userid )
        REFERENCES users;`

	_, err = db.Exec(createPostUserFK)
	if err != nil {
		fmt.Printf("Could not create FK Post User")
		panic(err)
	}

	// LIKE
	createLikeTable := `
	CREATE TABLE IF NOT EXISTS likes (
		id SERIAL PRIMARY KEY,
		userid INT,
		postid INT
	);`

	_, err = db.Exec(createLikeTable)
	if err != nil {
		fmt.Printf("Could not create like table")
		panic(err)
	}

	createLikeUserFK := `ALTER TABLE likes DROP CONSTRAINT IF EXISTS likes_user_fk;
	ALTER TABLE likes
    ADD CONSTRAINT likes_user_fk FOREIGN KEY ( userid )
        REFERENCES users;`

	_, err = db.Exec(createLikeUserFK)
	if err != nil {
		fmt.Printf("Could not create FK Like User")
		panic(err)
	}

	createLikePostFK := `ALTER TABLE likes DROP CONSTRAINT IF EXISTS likes_post_fk;
	ALTER TABLE likes
    ADD CONSTRAINT likes_post_fk FOREIGN KEY ( postid )
        REFERENCES posts;`

	_, err = db.Exec(createLikePostFK)
	if err != nil {
		fmt.Printf("Could not create FK Like Post")
		panic(err)
	}

	// COMMENT
	createCommentTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		content TEXT,
		userid INT,
		postid INT
	);`

	_, err = db.Exec(createCommentTable)
	if err != nil {
		fmt.Printf("Could not create comment table")
		panic(err)
	}

	createCommentUserFK := `ALTER TABLE comments DROP CONSTRAINT IF EXISTS comments_user_fk;
	ALTER TABLE comments
    ADD CONSTRAINT comments_user_fk FOREIGN KEY ( userid )
        REFERENCES users;`

	_, err = db.Exec(createCommentUserFK)
	if err != nil {
		fmt.Printf("Could not create FK Comment User")
		panic(err)
	}

	createCommentPostFK := `ALTER TABLE comments DROP CONSTRAINT IF EXISTS comments_post_fk;
	ALTER TABLE comments
    ADD CONSTRAINT comments_post_fk FOREIGN KEY ( postid )
        REFERENCES posts;`

	_, err = db.Exec(createCommentPostFK)
	if err != nil {
		fmt.Printf("Could not create FK Comment Post")
		panic(err)
	}
}

func AuthenticateUser(username string) string {
	var storedPassword string

	err := database.QueryRow("SELECT password FROM users WHERE name = $1", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		} else {
			panic(err)
		}
	}

	return storedPassword
}
