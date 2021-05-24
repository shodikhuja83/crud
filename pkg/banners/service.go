package banners

import (
	"context"
	"errors"
	"sync"
)

// Service ...
type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

// NewService ...
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

func (s *Service) Init() *Service {
	s.items = append(s.items,
		&Banner{
			ID:      1,
			Title:   "title1",
			Content: "content1",
			Button:  "button1",
			Link:    "link1",
		},
		&Banner{
			ID:      2,
			Title:   "title2",
			Content: "content2",
			Button:  "button2",
			Link:    "link2",
		},
		&Banner{
			ID:      3,
			Title:   "title3",
			Content: "content3",
			Button:  "button3",
			Link:    "link3",
		},
	)
	return s
}

// Banner ...
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

// All ...
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

// ByID ...
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}

// Save ...
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	id := len(s.items) + 1

	if item.ID == 0 {
		item.ID = int64(id)
		s.items = append(s.items, item)
	} else {
		for i, banner := range s.items {
			if banner.ID == item.ID {
				s.items[i] = item
			}
		}
	}
	return item, nil
}

// RemoveByID ...
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, banner := range s.items {
		if banner.ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}
