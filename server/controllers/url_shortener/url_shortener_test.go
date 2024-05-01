package urlshortener

import (
	"errors"
	"testing"
	"time"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/model"
	"pacstall.dev/webserver/types/repository"
	"pacstall.dev/webserver/utils/expect"
)

type mockShortenedLinkRepository struct {
	mockFindOneByLinkIdCalled int
	mockIncrementVisitsCalled int

	mockFindOneByLinkId func(linkId string) (repository.ShortenedLink, error)
	mockIncrementVisits func(id uint) error
}

func (r *mockShortenedLinkRepository) FindOneByLinkId(linkId string) (repository.ShortenedLink, error) {
	r.mockFindOneByLinkIdCalled += 1
	return r.mockFindOneByLinkId(linkId)
}

func (r *mockShortenedLinkRepository) IncrementVisits(id uint) error {
	r.mockIncrementVisitsCalled += 1
	return r.mockIncrementVisits(id)
}

type mockMatomoTrackerService struct {
	mockTrackShortLinkCalled int
	mockTrackShortLink       func(user, userAgent, urlRef, link string)
}

func (s *mockMatomoTrackerService) TrackShortLink(user, userAgent, urlRef, link string) {
	s.mockTrackShortLinkCalled += 1
	s.mockTrackShortLink(user, userAgent, urlRef, link)
}

const GOROUTINE_DELAY = 30 * time.Millisecond

func Test_UrlShortenerController_findShortenedLinkAndTrack_NotFound(t *testing.T) {
	shortenedLinkRepository := &mockShortenedLinkRepository{
		mockFindOneByLinkId: func(linkId string) (repository.ShortenedLink, error) {
			return nil, errors.New("not found")
		},
		mockIncrementVisits: func(id uint) error {
			return errors.New("not found")
		},
	}
	matomoTrackerService := &mockMatomoTrackerService{
		mockTrackShortLink: func(user, userAgent, urlRef, link string) {},
	}
	urlShortenerController := New(
		config.MatomoConfiguration{
			Enabled: true,
		},
		shortenedLinkRepository,
		matomoTrackerService,
	)

	linkId := "does-not-exist"
	doNotTrack := false

	_, found := urlShortenerController.findShortenedLinkAndTrack(linkId, doNotTrack, "", "", "")

	// Stuff happening in the background
	time.Sleep(GOROUTINE_DELAY)

	expect.False(t, "short url found", found)
	expect.Equals(t, "ShortenedLinkRepository.FindOneByLinkId calls", 1, shortenedLinkRepository.mockFindOneByLinkIdCalled)
	expect.Equals(t, "ShortenedLinkRepository.IncrementVisits calls", 0, shortenedLinkRepository.mockIncrementVisitsCalled)
	expect.Equals(t, "MatomoTrackerService.TrackShortLink calls", 0, matomoTrackerService.mockTrackShortLinkCalled)
}

func Test_UrlShortenerController_findShortenedLinkAndTrack_Found_DoNotTrack(t *testing.T) {
	shortenedLinkRepository := &mockShortenedLinkRepository{
		mockFindOneByLinkId: func(linkId string) (repository.ShortenedLink, error) {
			return &model.ShortenedLink{
				LinkId: linkId,
			}, nil
		},
		mockIncrementVisits: func(id uint) error {
			return nil
		},
	}
	matomoTrackerService := &mockMatomoTrackerService{
		mockTrackShortLink: func(user, userAgent, urlRef, link string) {},
	}
	urlShortenerController := New(
		config.MatomoConfiguration{
			Enabled: true,
		},
		shortenedLinkRepository,
		matomoTrackerService,
	)

	linkId := "linkId"
	doNotTrack := true

	_, found := urlShortenerController.findShortenedLinkAndTrack(linkId, doNotTrack, "", "", "")

	// Stuff happening in the background
	time.Sleep(GOROUTINE_DELAY)

	expect.True(t, "short url not found", found)
	expect.Equals(t, "ShortenedLinkRepository.FindOneByLinkId calls", 1, shortenedLinkRepository.mockFindOneByLinkIdCalled)
	expect.Equals(t, "ShortenedLinkRepository.IncrementVisits calls", 0, shortenedLinkRepository.mockIncrementVisitsCalled)
	expect.Equals(t, "MatomoTrackerService.TrackShortLink calls", 0, matomoTrackerService.mockTrackShortLinkCalled)
}

func Test_UrlShortenerController_findShortenedLinkAndTrack_Found_DoTrack(t *testing.T) {
	shortenedLinkRepository := &mockShortenedLinkRepository{
		mockFindOneByLinkId: func(linkId string) (repository.ShortenedLink, error) {
			return &model.ShortenedLink{
				LinkId: linkId,
			}, nil
		},
		mockIncrementVisits: func(id uint) error {
			return nil
		},
	}
	matomoTrackerService := &mockMatomoTrackerService{
		mockTrackShortLink: func(user, userAgent, urlRef, link string) {},
	}
	urlShortenerController := New(
		config.MatomoConfiguration{
			Enabled: true,
		},
		shortenedLinkRepository,
		matomoTrackerService,
	)

	linkId := "linkId"
	doNotTrack := false

	_, found := urlShortenerController.findShortenedLinkAndTrack(linkId, doNotTrack, "", "", "")

	// Stuff happening in the background
	time.Sleep(GOROUTINE_DELAY)

	expect.True(t, "short url not found", found)
	expect.Equals(t, "ShortenedLinkRepository.FindOneByLinkId calls", 1, shortenedLinkRepository.mockFindOneByLinkIdCalled)
	expect.Equals(t, "ShortenedLinkRepository.IncrementVisits calls", 1, shortenedLinkRepository.mockIncrementVisitsCalled)
	expect.Equals(t, "MatomoTrackerService.TrackShortLink calls", 1, matomoTrackerService.mockTrackShortLinkCalled)
}

func Test_UrlShortenerController_findShortenedLinkAndTrack_Found_MatomoDisabled(t *testing.T) {
	shortenedLinkRepository := &mockShortenedLinkRepository{
		mockFindOneByLinkId: func(linkId string) (repository.ShortenedLink, error) {
			return &model.ShortenedLink{
				LinkId: linkId,
			}, nil
		},
		mockIncrementVisits: func(id uint) error {
			return nil
		},
	}
	matomoTrackerService := &mockMatomoTrackerService{
		mockTrackShortLink: func(user, userAgent, urlRef, link string) {},
	}
	urlShortenerController := New(
		config.MatomoConfiguration{
			Enabled: false,
		},
		shortenedLinkRepository,
		matomoTrackerService,
	)

	linkId := "linkId"
	doNotTrack := false

	_, found := urlShortenerController.findShortenedLinkAndTrack(linkId, doNotTrack, "", "", "")

	// Stuff happening in the background
	time.Sleep(GOROUTINE_DELAY)

	expect.True(t, "short url not found", found)
	expect.Equals(t, "ShortenedLinkRepository.FindOneByLinkId calls", 1, shortenedLinkRepository.mockFindOneByLinkIdCalled)
	expect.Equals(t, "ShortenedLinkRepository.IncrementVisits calls", 1, shortenedLinkRepository.mockIncrementVisitsCalled)
	expect.Equals(t, "MatomoTrackerService.TrackShortLink calls", 0, matomoTrackerService.mockTrackShortLinkCalled)
}
