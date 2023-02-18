package store

type SuggestionKey string
type SuggestionText string
type SuggestionRef int
type SuggestionRefs map[SuggestionText]SuggestionRef
type Suggestions map[SuggestionKey]SuggestionRefs

func (s Suggestions) Get(key, text string) *SuggestionRef {
	suggestion := s[SuggestionKey(key)]
	if nil == suggestion {
		return nil
	}
	ref, ok := suggestion[SuggestionText(text)]
	if !ok {
		return nil
	}
	return &ref
}

func (s Suggestions) AddSuggestion(key string, text string) {
	refs := s[SuggestionKey(key)]
	if nil == refs {
		refs = make(SuggestionRefs)
		s[SuggestionKey(key)] = refs
	}

	_, ok := refs[SuggestionText(text)]
	if !ok {
		refs[SuggestionText(text)] = 0
	}

	refs[SuggestionText(text)]++
}


