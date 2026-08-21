package quorum

import "context"
type CollectionWindow struct { collector *Collector; closed bool }
func NewCollectionWindow(ctx context.Context, c *Collector) *CollectionWindow { return &CollectionWindow{collector: c} }
func (w *CollectionWindow) Add(x Commitment) bool { return w.collector.Add(x) }
func (w *CollectionWindow) Close() { w.closed = true }
