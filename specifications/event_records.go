package specifications

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"testing"
)

type CopyAndPaste struct {
	EventType  string `json:"eventType"`
	WebsiteUrl string `json:"websiteUrl"`
	SessionID  string `json:"sessionId"`
	Pasted     bool   `json:"pasted"`
	FormID     string `json:"formID"`
}

type Dimension struct {
	Width  string
	Height string
}

type Event struct {
	WebsiteUrl         string `json:"websiteUrl"`
	SessionId          string `json:"sessionId"`
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int             // Seconds
}

type Resize struct {
}

type TimeTaken struct {
}

type EventRecorder interface {
	FindEvent(ctx context.Context, sessionID string) (Event, error)
	CaptureCopyAndPaste(ctx context.Context, cp CopyAndPaste) error
	CaptureResize(ctx context.Context, rz Resize) error
	CaptureTimeOnPage(ctx context.Context, tStp TimeTaken) error
}

type EventSpecifications struct {
	Recorder EventRecorder
}

func NewEventSpecifications(recorder EventRecorder) *EventSpecifications {
	return &EventSpecifications{Recorder: recorder}
}

func (s EventSpecifications) CopyAndPasteRecording(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	t.Run("A CopyAndPaste event is recorded", func(t *testing.T) {
		tc := testcase.NewT(t, nil)

		copyAndPaste := CopyAndPaste{
			EventType:  "copyAndPaste",
			WebsiteUrl: "http://mikejeuga.com",
			SessionID:  uuid.New().String(),
			Pasted:     true,
			FormID:     "homePage",
		}

		err := s.Recorder.CaptureCopyAndPaste(ctx, copyAndPaste)
		tc.Must.NoError(err, "the recording of copy and paste event should not error out")

		event, err := s.Recorder.FindEvent(ctx, copyAndPaste.SessionID)
		tc.Must.NoError(err, "retrieving an event should not error out")

		tc.Must.True(event.CopyAndPaste["pasted"])
	})
}
