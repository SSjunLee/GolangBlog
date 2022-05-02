// @Author ljn 2022/5/1 11:23:00
package models

import "time"

type Archive struct {
	Posts []Post    `json:"posts"`
	Time  time.Time `json:"time"`
}

func PostArchive() ([]Archive, error) {
	posts := make([]Post, 0, 8)
	achieves := make([]Archive, 0, 8)
	err := Db.Select("id", "title", "path", "created").
		Where("kind = ? AND status = ?", KindArticle, 2).
		Order("created desc").
		Limit(8).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}
	for _, p := range posts {
		if idx := indexOfArchive(achieves, p.Created); idx != -1 {
			achieves[idx].Posts = append(achieves[idx].Posts, p)
		} else {
			achieves = append(achieves, Archive{
				Posts: []Post{p},
				Time:  p.Created,
			})
		}
	}
	return achieves, nil
}

func indexOfArchive(achieves []Archive, t time.Time) int {
	for i, a := range achieves {
		if t.Year() == a.Time.Year() && t.Month() == a.Time.Month() {
			return i
		}
	}
	return -1
}
