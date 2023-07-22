package schemes

type SchemeContact struct {
	ID       string `json:"id"        validate:"uuid"`
	Realname string `json:"real_name" validate:"required,lowercase"`
	Whatsapp uint64 `json:"whatsapp"  validate:"required,numeric"`
	Email    string `json:"email"     validate:"required"`
}
