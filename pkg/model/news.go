package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

// News type
type News struct {
	ID        string
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var newStorage []*News

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews from request
func CreateNews(n *News) {
	n.ID = generateID()
	n.CreatedAt = time.Now()
	n.UpdatedAt = n.CreatedAt
	newStorage = append(newStorage, n)
}

// ListNews get all news
func ListNews() []*News {
	return newStorage
}

// GetNews from request id
func GetNews(id string) *News {
	for _, new := range newStorage {
		if new.ID == id {
			return new
		}
	}
	return nil
}

// DeleteNews from request id
func DeleteNews(id string) {
	for i, new := range newStorage {
		if new.ID == id {
			newStorage = append(newStorage[:i], newStorage[i+1:]...)
		}
	}
}
