# Bug reproduction

- Bug: a failed worker job is dropped from the result stream, the collector loses the error, dispatch can return an empty result set, and the expected count is short by one.
- Trigger: process a mixed batch containing a job whose work function returns an error, then inspect the result channel and count.
- Error: `missing error result`, `results=[] err=<nil>`, and `expected count mismatch`.
