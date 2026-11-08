package cli

import "wruffic/api"

type Choice struct {
	Label string
	Kind  string
	ID    string
	Song  api.Song
}
