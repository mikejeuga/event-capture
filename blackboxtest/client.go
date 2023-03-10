package blackboxtest

import (
	"context"
	"github.com/mikejeuga/event-capture/specifications"
	"log"
	"net/http"
	"time"
)

type TestClient struct {
	baseURL    string
	httpDriver *http.Client
	aLogger    log.Logger
}

func NewTestClient() *TestClient {
	return &TestClient{
		baseURL: "http://localhost:8080/",
		httpDriver: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   5 * time.Second,
		},
		aLogger: log.Logger{},
	}
}

func (t TestClient) FindEvent(ctx context.Context, sessionID string) (specifications.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (t TestClient) CaptureCopyAndPaste(ctx context.Context, cp specifications.CopyAndPaste) error {
	//TODO implement me
	panic("implement me")
}

func (t TestClient) CaptureResize(ctx context.Context, rz specifications.Resize) error {
	//TODO implement me
	panic("implement me")
}

func (t TestClient) CaptureTimeOnPage(ctx context.Context, tStp specifications.TimeTaken) error {
	//TODO implement me
	panic("implement me")
}
