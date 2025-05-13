package limiter

type tokenDeque struct {
	items []*Token
}

func newTokenDeque() *tokenDeque {
	return &tokenDeque{}
}

func (d *tokenDeque) append(token *Token) {
	d.items = append(d.items, token)
}

func (d *tokenDeque) push(token *Token) {}

func (d *tokenDeque) pop() *Token {
	if len(d.items) == 0 {
		return nil
	}

	return d.items[0]
}

func (d *tokenDeque) len() int {
	return len(d.items)
}
