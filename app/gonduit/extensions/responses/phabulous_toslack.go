package responses

import "github.com/SpinlockLabs/phabulous/app/gonduit/extensions/entities"

// PhabulousToSlackResponse is the response to calling phabulous.toslack.
type PhabulousToSlackResponse map[string]*entities.SlackUser
