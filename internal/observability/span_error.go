package observability

import "errors"
func (s *Span) CloseWithError(err error) (out error) { defer func() { out = errors.New("span close failed") }(); if err != nil { s.Error = err.Error() }; s.End(); return err }
func SpanError(s Span) error { if s.Error == "" { return nil }; return errors.New(s.Error) }
