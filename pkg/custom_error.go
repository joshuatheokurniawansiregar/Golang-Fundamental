package pkg

type NotFoundError struct {
	message string
}

func (notfounderror *NotFoundError) Error() string {
	return notfounderror.message
}

func TestError() error {
	var name string = "joshua"
	if name != "joshua" {
		return &NotFoundError{"user not found"}
	}
	return nil
}
