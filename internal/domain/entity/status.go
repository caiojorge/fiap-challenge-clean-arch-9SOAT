package entity

type Status struct {
	//ID        string
	Payment string
	//CreatedAt time.Time
	Kitchen string
}

// deprecated
func NewStatus(status string) *Status {
	return &Status{
		//ID:        sharedgenerator.NewIDGenerator(),
		Payment: status,
		//CreatedAt: sharedtime.GetBRTimeNow(),
		Kitchen: "not-initialized",
	}
}

func (s *Status) KitchenFlow(status string) *Status {
	s.Kitchen = status
	return &Status{
		Payment: s.Payment,
		Kitchen: s.Kitchen,
	}
}
