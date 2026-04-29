package database

import "API/internal/model"

func (db Database) InsertPost(post model.PostModel) error {
	sqlQueru := `
	INTSERT INTO tasks(title, text, tag, create_at) 
	VALUES($1, $2, $3, $4);
	`
	_, err := db.conn.Exec(db.ctx, sqlQueru, post.Title, post.Text, post.Tag, post.CreatedAt)
	return err
}

func (db Database) GetAllPost() (allPost map[string]model.PostModel, err error) {
	sqlQueru := `
	SELECT title, text, tag, created_at FROM posts
	LIMIT 100; 
	`
	rows, err := db.conn.Query(db.ctx, sqlQueru)
	if err != nil {
		return allPost, err
	}
	defer rows.Close()

	for rows.Next() {
		post := model.PostModel{}
		if err = rows.Scan(post.Title, post.Text, post.Tag, post.CreatedAt); err != nil {
			return allPost, err
		}
		allPost[post.Title] = post
	}

	return allPost, err
}
