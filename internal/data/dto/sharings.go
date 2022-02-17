package dto

// CreateSharing_Params Object
type CreateSharing_Params struct {
	Content string `json:"content,omitempty" example:"aaaabbbbcccc"`
}

// Sharing Object
type Sharing struct {
	Hash    string `json:"hash,omitempty" example:"11c85195ae99540ac07f80e2905e6e39aaefc4ac94cd380f366e79ba83560566"`
	Content string `json:"content,omitempty" example:"aaaabbbbcccc"`
}
