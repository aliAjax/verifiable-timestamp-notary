package audit

import (
    "errors"
    "testing"
    "example.com/verifiable-timestamp-notary/internal/observability"
)

func TestAuditBatchClosesResourceA7(t *testing.T) { r := &BatchResource{}; _, _ = EncodeBatch([]Event{{ID: "bad"}}, r); if !ResourceClosed(r) { t.Fatal("resource leaked") } }
func TestAuditCleanupPreservesEncodeErrorA7(t *testing.T) { _, err := EncodeWithCleanup([]Event{{ID: "bad"}}, func() error { return errors.New("cleanup") }); if err == nil || err.Error() != "encode failed" { t.Fatalf("err=%v", err) } }
func TestSpanClosePreservesOriginalErrorA7(t *testing.T) { original := errors.New("encode"); s := observability.NewSpan("t", "export"); if got := s.CloseWithError(original); got == nil || got.Error() != original.Error() { t.Fatalf("got=%v", got) } }
func TestMetricsFlushReturnsCleanupErrorA7(t *testing.T) { original := errors.New("flush"); reg := observability.NewRegistry(); if err := reg.Flush(func() error { return original }); !errors.Is(err, original) { t.Fatalf("flush=%v", err) } }
