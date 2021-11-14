package types

type Request struct {
	ID       string
	Location string
	IP       string
	Data     map[string]string
}

func (r Request) IsValid() bool {
	if len(r.IP) == 0 {
		return false
	}
	if len(r.ID) == 0 {
		return false
	}
	if len(r.Location) == 0 {
		return false
	}

	return true
}
