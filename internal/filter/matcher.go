package filter

import (
	"regexp"
	"strconv"
	"strings"

	"pwncat/internal/config"
	"pwncat/internal/http"
)

type Condition interface {
	Match(resp *http.Response) bool
}

type Matcher struct {
	conditions []Condition
}

func NewMatcher(cfg *config.Config) (Condition, error) {
	m := &Matcher{}

	if cfg.MatchStatus != "" {
		m.conditions = append(m.conditions, NewStatusMatcher(cfg.MatchStatus))
	}
	if cfg.MatchSize != "" {
		m.conditions = append(m.conditions, NewSizeMatcher(cfg.MatchSize))
	}
	if cfg.MatchRegex != "" {
		re, err := NewRegexMatcher(cfg.MatchRegex)
		if err != nil {
			return nil, err
		}
		m.conditions = append(m.conditions, re)
	}

	if len(m.conditions) == 0 {
		m.conditions = append(m.conditions, NewDefaultMatcher())
	}

	return m, nil
}

func (m *Matcher) Match(resp *http.Response) bool {
	for _, cond := range m.conditions {
		if !cond.Match(resp) {
			return false
		}
	}
	return true
}

type StatusMatcher struct{ codes []int }

func NewStatusMatcher(spec string) *StatusMatcher {
	parts := strings.Split(spec, ",")
	codes := make([]int, 0, len(parts))
	for _, p := range parts {
		if c, err := strconv.Atoi(p); err == nil {
			codes = append(codes, c)
		}
	}
	return &StatusMatcher{codes: codes}
}

func (m *StatusMatcher) Match(resp *http.Response) bool {
	for _, code := range m.codes {
		if resp.StatusCode == code {
			return true
		}
	}
	return false
}

type SizeMatcher struct{ Min, Max int }

func NewSizeMatcher(spec string) *SizeMatcher {
	parts := strings.Split(spec, "-")
	if len(parts) != 2 {
		return &SizeMatcher{Min: 0, Max: int(^uint(0) >> 1)}
	}
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])
	return &SizeMatcher{Min: min, Max: max}
}

func (m *SizeMatcher) Match(resp *http.Response) bool {
	return resp.Size >= m.Min && resp.Size <= m.Max
}

type RegexMatcher struct{ re *regexp.Regexp }

func NewRegexMatcher(pattern string) (*RegexMatcher, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexMatcher{re: re}, nil
}

func (m *RegexMatcher) Match(resp *http.Response) bool {
	return m.re.MatchString(resp.Body)
}

type DefaultMatcher struct{}

func NewDefaultMatcher() *DefaultMatcher { return &DefaultMatcher{} }

func (m *DefaultMatcher) Match(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}
