package memento20

type InputText struct {
	content string
}

func (t *InputText) Append(content string) {
	t.content += content
}

func (t *InputText) GetText() string {
	return t.content
}

func (t *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: t.content}
}

func (t *InputText) Restore(s *Snapshot) {
	t.content = s.Content()
}

type Snapshot struct {
	content string
}

func (s *Snapshot) Content() string {
	return s.content
}
