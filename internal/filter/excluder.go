package filter

import (
	"pwncat/internal/config"
	"pwncat/internal/http"
)

type Excluder struct {
	conditions []Condition
}

func NewExcluder(cfg *config.Config) (Condition, error) {
	e := &Excluder{}

	if cfg.ExcludeStatus != "" {
		e.conditions = append(e.conditions, NewStatusMatcher(cfg.ExcludeStatus))
	}
	if cfg.ExcludeSize != "" {
		e.conditions = append(e.conditions, NewSizeMatcher(cfg.ExcludeSize))
	}
	if cfg.ExcludeRegex != "" {
		re, err := NewRegexMatcher(cfg.ExcludeRegex)
		if err != nil {
			return nil, err
		}
		e.conditions = append(e.conditions, re)
	}

	return e, nil
}

func (e *Excluder) Match(resp *http.Response) bool {
	for _, cond := range e.conditions {
		if cond.Match(resp) {
			return true
		}
	}
	return false
}
