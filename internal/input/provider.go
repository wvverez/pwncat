package input

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"pwncat/internal/config"
)

type Provider interface {
	Next() (*Value, error)
	Total() int64
}

type Value struct {
	Position int64
	Values   map[string][]byte
}

func NewProvider(cfg *config.Config) (Provider, error) {
	if cfg.Wordlist == "" {
		return NewStdin()
	}

	if strings.HasPrefix(cfg.Wordlist, "range:") {
		return NewRange(cfg.Wordlist)
	}

	return NewFile(cfg.Wordlist)
}

type FileProvider struct {
	file    *os.File
	scanner *bufio.Scanner
	pos     int64
	total   int64
}

func NewFile(path string) (*FileProvider, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	total := int64(0)
	for scanner.Scan() {
		total++
	}

	f.Seek(0, 0)
	return &FileProvider{
		file:    f,
		scanner: bufio.NewScanner(f),
		total:   total,
	}, nil
}

func (p *FileProvider) Next() (*Value, error) {
	if !p.scanner.Scan() {
		return nil, io.EOF
	}
	p.pos++
	return &Value{
		Position: p.pos,
		Values: map[string][]byte{
			"PWN": []byte(p.scanner.Text()),
		},
	}, nil
}

func (p *FileProvider) Total() int64 { return p.total }

type StdinProvider struct {
	scanner *bufio.Scanner
	pos     int64
}

func NewStdin() (*StdinProvider, error) {
	return &StdinProvider{
		scanner: bufio.NewScanner(os.Stdin),
	}, nil
}

func (p *StdinProvider) Next() (*Value, error) {
	if !p.scanner.Scan() {
		return nil, io.EOF
	}
	p.pos++
	return &Value{
		Position: p.pos,
		Values: map[string][]byte{
			"PWN": []byte(p.scanner.Text()),
		},
	}, nil
}

func (p *StdinProvider) Total() int64 { return 0 }

type RangeProvider struct {
	start   int
	end     int
	current int
	pos     int64
	total   int64
}

func NewRange(spec string) (*RangeProvider, error) {
	parts := strings.Split(spec, ":")
	if len(parts) != 3 {
		return nil, ErrInvalidRange
	}

	start, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}
	if start > end {
		return nil, ErrInvalidRange
	}

	return &RangeProvider{
		start:   start,
		end:     end,
		current: start,
		total:   int64(end - start + 1),
	}, nil
}

func (p *RangeProvider) Next() (*Value, error) {
	if p.current > p.end {
		return nil, io.EOF
	}
	val := strconv.Itoa(p.current)
	p.current++
	p.pos++
	return &Value{
		Position: p.pos,
		Values: map[string][]byte{
			"PWN": []byte(val),
		},
	}, nil
}

func (p *RangeProvider) Total() int64 { return p.total }

var ErrInvalidRange = &SourceError{"invalid range specification"}

type SourceError struct{ msg string }

func (e *SourceError) Error() string { return e.msg }
