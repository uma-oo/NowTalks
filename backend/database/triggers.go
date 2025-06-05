package database

func (db *Database) Triggers() error {
	trigger_total_comments := `CREATE TRIGGER IF NOT EXISTS increment_total_comments 
			AFTER INSERT ON comments 
			FOR EACH ROW 
			BEGIN 
				UPDATE posts
				SET total_comments=total_comments+1 
				WHERE posts.postID = NEW.postID;
			END;`

	statement, err := db.Database.Prepare(trigger_total_comments)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}
