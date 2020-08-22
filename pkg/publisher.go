package pkg

import (
	"context"
	"google.golang.org/api/option"
	"log"
)
import "google.golang.org/api/androidpublisher/v3"
import _ "google.golang.org/api/option"

type Service struct {
	*androidpublisher.Service
}

func createService(path string) Service {
	ctx := context.Background()
	service, err := androidpublisher.NewService(ctx, option.WithCredentialsFile(path))
	if err != nil {
		log.Fatalf("Unable to retrieve Android Publisher client: %v", err)
	}
	return Service{service}
}

func (s *Service) createEdit(packageName string) *androidpublisher.AppEdit {
	edit, err := s.Edits.Insert(packageName, nil).Do()
	if err != nil {
		log.Fatal(err)
	}
	return edit
}

func (s *Service) getTracks(packageName string, editId string) *androidpublisher.TracksListResponse {
	tracks, err := s.Edits.Tracks.List(packageName, editId).Fields("tracks").Do()
	if err != nil {
		log.Fatal(err)
	}
	return tracks
}

func Fetch(path string, packageName string) int64 {
	var versionCodes []int64
	var max int64
	service := createService(path)
	edit := service.createEdit(packageName)
	tracks := service.getTracks(packageName, edit.Id)
	for _, track := range tracks.Tracks {
		for _, release := range track.Releases {
			versionCodes = append(versionCodes, release.VersionCodes[0])
		}
	}
	for _, code := range versionCodes {
		if code > max {
			max = code
		}
	}
	return max
}
