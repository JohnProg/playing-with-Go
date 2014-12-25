package data

type Attendee struct {
	Person //type embedding for composition
	interests []string
}