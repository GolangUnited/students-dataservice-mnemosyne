package lessons

const (
	GetLessonByIdQuery    = `SELECT * FROM lessons WHERE id = $1`
	AddLessonQuery        = `INSERT INTO lessons (presentation, video_link, lesson_date, homework, lecturer_id, group_id, language) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	UpdateLessonByIdQuery = `UPDATE lessons 
								SET presentation = $1,
	    							video_link = $2,
	    							lesson_date = $3,
	   				 				homework =$4,
	    							lecturer_id =$5,
	    							group_id =$6, 
	    							language =$7,
								    updated_at = $8
								    WHERE id = $9`
	DeactivateLessonByIdQuery = `UPDATE lessons SET deleted = true, updated_at = $1 WHERE id = $2`
	ActivateLessonByIdQuery   = `UPDATE lessons SET deleted = false, updated_at = $1 WHERE id = $2`
)
