package main

import (
	"log"
	"time"
)

type Service interface {
	getLinkForShortID(shortID string) (*Link, error)
	incrementGlobalCounter() uint
	incrementLinkCounter(link *Link)
	newLink(URL string) *Link
	close()
}

type ServiceImpl struct {
	linkRepo    LinkRepository
	counterRepo CounterRepository
}

func (s *ServiceImpl) newLink(URL string) *Link {
	var timenow = time.Now().UTC()
	var counter = s.incrementGlobalCounter()
	var shortID = baseconv.Encode(counter)

	link := Link{
		ID:         URL + "-" + shortID,
		URL:        URL,
		Clicks:     0,
		ShortID:    shortID,
		ShortIDInt: counter,
		CreatedAt:  &timenow,
		UpdatedAt:  &timenow,
	}

	err := s.linkRepo.InsertLink(link)
	if err != nil {
		log.Fatal("Error inserting new link ", err)
	}

	return &link
}

func (s *ServiceImpl) getLinkForShortID(shortID string) (*Link, error) {
	if len(shortID) > 10 {
		return nil, errInvalidLink
	}

	shortIDInt := baseconv.Decode(shortID)
	link, err := s.linkRepo.FindLinkByShortIdInt(shortIDInt)

	return link, err
}

func (s *ServiceImpl) incrementLinkCounter(link *Link) {
	var timenow = time.Now().UTC()
	link.Clicks++
	link.UpdatedAt = &timenow
	err := s.linkRepo.UpdateLink(*link)
	if err != nil {
		log.Fatal("Error incrementing counter for link", link, err)
	}
}

func (s *ServiceImpl) incrementGlobalCounter() uint {
	var timenow = time.Now().UTC()

	counter, err := s.counterRepo.FindCounterById("1")
	if counter == nil {
		counter = &Counter{
			ID:        "1",
			Count:     0,
			CreatedAt: &timenow,
			StatType:  "counter",
		}
	}

	counter.Count++
	counter.UpdatedAt = &timenow

	err = s.counterRepo.UpsertCounter(*counter)
	if err != nil {
		log.Fatal("Failed updating global counter ", err)
	}

	return counter.Count
}

func (s *ServiceImpl) close() {
	s.linkRepo.close()
	s.counterRepo.close()
}