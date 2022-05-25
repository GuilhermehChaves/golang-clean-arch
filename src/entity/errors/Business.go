package errorEntity

import "fmt"

type Business struct {
	message string
}

func NewBusinessError(message string) *Business {
	return &Business{
		message: message,
	}
}

func (b *Business) Error() string {
	return fmt.Sprintf("Business error: %s", b.message)
}
