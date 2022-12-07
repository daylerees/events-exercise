package events

import "events-exercise/internal/accounts"

// StreamReader is an interface that allows a client to read events from a stream.
type StreamReader interface {

	// Read reads an event at a time from the stream.
	Read(any) error
}

// Stream provides access to a stream of events through the use of a StreamReader.
type Stream struct{}

// Read reads events from the stream and passes them to the given StreamReader.
func (s *Stream) Read(sr StreamReader) error {

	for _, event := range s.eventSequence() {
		err := sr.Read(event)
		if err != nil {
			return err
		}
	}

	return nil
}

// eventSequence returns a sequence of hard-coded events.
func (s *Stream) eventSequence() []any {
	return []any{
		accounts.UserAccountCreatedEvent{
			UserId: "0af3a961-5146-46b5-93f8-95c0ab687007",
		},
		accounts.UserAccountCreatedEvent{
			UserId: "d60f3e10-b707-4c76-b165-da38b95aa4b9",
		},
		accounts.UserAccountCreatedEvent{
			UserId: "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserAccountUpdatedEvent{
			UserId:   "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			FullName: "Rosetta Brandon",
			Email:    "rbrandon@test.com",
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserAccountUpdatedEvent{
			UserId:   "0af3a961-5146-46b5-93f8-95c0ab687007",
			FullName: "Anthony Swiss",
			Email:    "anthony.swiss@test.com",
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_GREEN,
		},
		accounts.UserLostBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "0af3a961-5146-46b5-93f8-95c0ab687007",
			BadgeColour: accounts.BadgeColour_GREEN,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserLostBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "d60f3e10-b707-4c76-b165-da38b95aa4b9",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserAccountUpdatedEvent{
			UserId:   "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			FullName: "Neves Firmino",
			Email:    "neves88@test.com",
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_BLUE,
		},
		accounts.UserAccountUpdatedEvent{
			UserId: "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			Email:  "neves.firmino@test.com",
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_RED,
		},
		accounts.UserGainedBadgeEvent{
			UserId:      "6c5031e7-ff1c-4986-ac27-05a2737cd2f4",
			BadgeColour: accounts.BadgeColour_GREEN,
		},
		accounts.UserAccountUpdatedEvent{
			UserId:   "0af3a961-5146-46b5-93f8-95c0ab687007",
			FullName: "Anthony Swiss-Jones",
		},
	}
}
