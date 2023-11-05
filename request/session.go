package request

type LcSession struct {
	Session   string `json:"session"`
	CsrfToken string `json:"csrfToken"`
}
