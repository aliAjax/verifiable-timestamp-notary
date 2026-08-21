package evidence

import (
    "testing"
    "time"
)

func recordsForTest() []Record { now := time.Now().Unix(); return []Record{{ID: "a", State: "open", CreatedUnix: now}, {ID: "b", State: "closed", CreatedUnix: now}, {ID: "c", State: "open", CreatedUnix: now}} }
func TestEvidenceListCopiesSourceIsolation(t *testing.T) { src := recordsForTest(); list := NewEvidenceList(src); src[0].ID = "changed"; if list.All()[0].ID != "a" { t.Fatalf("list aliases source: %v", list.All()) } }
func TestEvidenceFilterCopiesSourceIsolation(t *testing.T) { src := recordsForTest(); list := NewEvidenceList(src); got := list.Filter("open"); if len(got) != 2 || src[1].ID != "b" { t.Fatalf("got=%v src=%v", got, src) }; got[0].ID = "changed"; if list.All()[0].ID != "a" { t.Fatalf("filtered result aliases list: %v", list.All()) } }
func TestEvidencePageDoesNotOverwriteInput(t *testing.T) { src := recordsForTest(); dst := src[:2]; got := PageInto(dst, src, 1, 1); if got[0].ID != "b" || src[0].ID != "a" { t.Fatalf("got=%v src=%v", got, src) } }
func TestEvidenceRetentionDoesNotAliasInput(t *testing.T) { src := recordsForTest(); got := BatchKeep(src, time.Now(), time.Minute); RestoreRecord(got, 0, Record{ID: "changed"}); if src[0].ID != "a" { t.Fatalf("got=%v src=%v", got, src) } }
func TestEvidenceCopyIsIndependent(t *testing.T) { src := recordsForTest(); got := CopyRecords(src); RestoreRecord(got, 0, Record{ID: "changed"}); if src[0].ID != "a" { t.Fatal("copy aliased source") } }
