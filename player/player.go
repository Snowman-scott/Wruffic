package player

import (
	"fmt"
	// "time"

	"github.com/gen2brain/go-mpv"

	"wruffic/api"
)


type Player struct {
	client *api.Client
	mpv    *mpv.Mpv
}

func New(client *api.Client) (*Player, error) {
	m := mpv.New()

	if err := m.SetOptionString("video", "no"); err != nil {
			return nil, fmt.Errorf("mpv set video: %w", err)
		}
	if err := m.Initialize(); err != nil {
		return nil, fmt.Errorf("mpv init: %w", err)
	}

	return &Player{client: client, mpv: m}, nil
}

func (p *Player) PlayAlbum(Songs []api.Song) error{
	for _, s := range Songs {
		fmt.Printf("Track: %d\nTitle: %s\nDuration: %d seconds", s.Track, s.Title, s.Duration)
		url, err := p.client.StreamURL(s.ID)
		if err != nil {
			return fmt.Errorf("wuff stweam faw %s failed with: %v", s.Title, err)
		}

		err = p.mpv.Command([]string{"loadfile", url, "append-play"})
		if err != nil {
			return fmt.Errorf("qewing %s Faiwed with : %v", s.Title, err)
		}
	}
	return nil
}

func (p *Player) Close() {
	p.mpv.TerminateDestroy()
}
