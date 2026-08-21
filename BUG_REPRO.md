# Bug reproduction

- Bug: evidence constructors, filters, paging, retention, and copy helpers return slices that alias caller-owned storage.
- Trigger: mutate the source, destination, or returned slice after each operation and compare the original records.
- Error: failures report source/list mutation such as `list aliases source`, `filtered result aliases list`, or `copy aliased source`.
