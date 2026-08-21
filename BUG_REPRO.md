# Bug reproduction

- Bug: missing-claim errors lose their sentinel while crossing the API boundary, so the transport reports a server error and marks a not-found response retryable.
- Trigger: classify a wrapped missing claim through the notary, API, and transport helpers.
- Error: the target cases fail because the status is 500 instead of not-found/404 and the error chain is not classifiable.
