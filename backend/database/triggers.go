package database

// func (db *Database) Triggers() error {
// 	trigger_total_comments := `CREATE TRIGGER IF NOT EXISTS increment_total_comments 
// 			AFTER INSERT ON comments 
// 			FOR EACH ROW 
// 			BEGIN 
// 				UPDATE posts
// 				SET total_comments=total_comments+1 
// 				WHERE posts.postID = NEW.postID;
// 			END;`
// 	trigger_total_likes_comments_insert := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_likes_comments_insert
// 	   AFTER INSERT ON comment_reactions
// 		FOR EACH ROW
// 		BEGIN

// 		UPDATE comments
// 		SET total_likes = total_likes + 1
// 		WHERE comments.id = NEW.comment_id
// 		AND NEW.reaction_id = 1 ;
// 		END;`

// 	trigger_total_likes_comments_update := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_likes_comments_update
// 		AFTER UPDATE ON comment_reactions
// 		FOR EACH ROW
// 		BEGIN

// 		UPDATE comments
// 		SET total_likes = total_likes + 1
// 		WHERE comments.id = NEW.comment_id
// 		AND OLD.reaction_id=0
// 		AND NEW.reaction_id = 1 ;


// 		UPDATE comments
// 		SET 
//         total_likes = total_likes + 1,
//         total_dislikes = CASE 
//             WHEN total_dislikes > 0 THEN total_dislikes - 1
//             ELSE 0
//         END
// 		WHERE comments.id = NEW.comment_id AND total_dislikes -1 >= 0
// 		AND OLD.reaction_id=-1
// 		AND NEW.reaction_id = 1 ;

// 		UPDATE comments
// 		SET total_likes = CASE 
//         WHEN total_likes > 0 THEN total_likes - 1
//         ELSE 0
//         END
// 		WHERE comments.id = NEW.comment_id
// 		AND OLD.reaction_id=1
// 		AND NEW.reaction_id = 0;
// 		END;`

// 	trigger_total_dislikes_comments_insert := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_dislikes_comments_insert
// 	   AFTER INSERT ON comment_reactions
// 		FOR EACH ROW
// 		BEGIN

// 		UPDATE comments
// 		SET total_dislikes = total_dislikes + 1
// 		WHERE comments.id = NEW.comment_id
// 		AND NEW.reaction_id = -1 ;
// 		END;`
// 	trigger_total_dislikes_comments_update := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_dislikes_comments_update
// 		AFTER UPDATE ON comment_reactions
// 		FOR EACH ROW
// 		BEGIN

// 		UPDATE comments
// 		SET total_dislikes = total_dislikes + 1
// 		WHERE comments.id = NEW.comment_id
// 		AND OLD.reaction_id=0 
// 		AND NEW.reaction_id = -1 ;

// 		UPDATE comments
// 		SET 
//         total_dislikes = total_dislikes + 1,
//         total_likes = CASE 
//             WHEN total_likes > 0 THEN total_likes - 1
//             ELSE 0
//         END
// 		WHERE comments.id = NEW.comment_id
// 		AND OLD.reaction_id=1
// 		AND NEW.reaction_id = -1 ;

// 		UPDATE comments
// 		SET total_dislikes = CASE 
//         WHEN total_dislikes > 0 THEN total_dislikes - 1
//         ELSE 0
//         END
// 		WHERE comments.id = NEW.comment_id
// 		AND OLD.reaction_id=-1
// 		AND NEW.reaction_id = 0;
// 		END;`
// 	trigger_total_likes_posts_insert := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_likes_posts_insert
// 		AFTER INSERT ON post_reaction
// 		 FOR EACH ROW
// 		 BEGIN
 
// 		 UPDATE posts
// 		 SET total_likes = total_likes + 1
// 		 WHERE posts.id = NEW.post_id 
// 		 AND NEW.reaction_id = 1 ;
// 		 END;`
// 	trigger_total_likes_posts_update := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_likes_posts_update
// 		 AFTER UPDATE ON post_reaction
// 		 FOR EACH ROW
// 		 BEGIN
 
// 		 UPDATE posts
// 		 SET total_likes = total_likes + 1
// 		 WHERE posts.id = NEW.post_id
// 		 AND OLD.reaction_id=0
// 		 AND NEW.reaction_id = 1 ;
 
// 		 UPDATE posts
// 		 SET 
// 		 total_likes = total_likes + 1,
// 		 total_dislikes = CASE 
// 			 WHEN total_dislikes > 0 THEN total_dislikes - 1
// 			 ELSE 0
// 		 END
// 		 WHERE posts.id = NEW.post_id AND total_dislikes -1 >= 0
// 		 AND OLD.reaction_id=-1
// 		 AND NEW.reaction_id = 1 ;
 
// 		 UPDATE posts
// 		 SET total_likes = CASE 
// 		 WHEN total_likes > 0 THEN total_likes - 1
// 		 ELSE 0
// 		 END
// 		 WHERE posts.id = NEW.post_id
// 		 AND OLD.reaction_id=1
// 		 AND NEW.reaction_id = 0;
// 		 END;`
// 	trigger_total_dislikes_posts_insert := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_dislikes_posts_insert
// 		 AFTER INSERT ON post_reaction
// 		  FOR EACH ROW
// 		  BEGIN
  
// 		  UPDATE posts
// 		  SET total_dislikes = total_dislikes + 1
// 		  WHERE posts.id = NEW.post_id
// 		  AND NEW.reaction_id = -1 ;
// 		  END;`
// 	trigger_total_dislikes_posts_update := `CREATE TRIGGER IF NOT EXISTS increment_or_decrement_total_dislikes_posts_update
// 		  AFTER UPDATE ON post_reaction
// 		  FOR EACH ROW
// 		  BEGIN
  
// 		  UPDATE posts
// 		  SET total_dislikes = total_dislikes + 1
// 		  WHERE posts.id = NEW.post_id
// 		  AND OLD.reaction_id=0 
// 		  AND NEW.reaction_id = -1 ;
  
// 		  UPDATE posts
// 		  SET 
// 		  total_dislikes = total_dislikes + 1,
// 		  total_likes = CASE 
// 			  WHEN total_likes > 0 THEN total_likes - 1
// 			  ELSE 0
// 		  END
// 		  WHERE posts.id = NEW.post_id
// 		  AND OLD.reaction_id=1
// 		  AND NEW.reaction_id = -1 ;
  
// 		  UPDATE posts
// 		  SET total_dislikes = CASE 
// 		  WHEN total_dislikes > 0 THEN total_dislikes - 1
// 		  ELSE 0
// 		  END
// 		  WHERE posts.id = NEW.post_id
// 		  AND OLD.reaction_id=-1
// 		  AND NEW.reaction_id = 0;
// 		  END;`
// 	triggers := []string{
// 		trigger_total_comments, trigger_total_likes_comments_insert,
// 		trigger_total_likes_comments_update, trigger_total_dislikes_comments_insert,
// 		trigger_total_dislikes_comments_update, trigger_total_likes_posts_insert, trigger_total_likes_posts_update, trigger_total_dislikes_posts_insert, trigger_total_dislikes_posts_update,
// 	}
// 	for _, query := range triggers {
// 		statement, err := db.Database.Prepare(query)
// 		if err != nil {
// 			return err
// 		}
// 		defer statement.Close()
// 		statement.Exec()
// 	}
// 	return nil
// }
