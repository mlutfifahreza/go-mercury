package gallery_service

import "go-mercury/internal/data/gallery_db"

func (s Service) RegisterUser(link gallery_db.Link) (int, error) {
	affectedCount, err := s.db.UpdateLink(link)
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}
