# Bug reproduction

- Bug: audit export cleanup is skipped on an early encode error, and deferred cleanup errors overwrite or discard the original operation error.
- Trigger: encode a batch containing the bad event while the resource close function also fails, then exercise span close and metrics flush.
- Error: `resource leaked`, `err=cleanup`, `got=span close failed`, and `flush=<nil>`.
