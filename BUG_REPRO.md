# Bug reproduction

- Bug: Merkle validation wraps leaf and duplicate errors without preserving the error chain, and the classifier falls back to an unknown/retryable class.
- Trigger: validate an empty digest and a duplicate leaf through normalized, proof, and incremental paths.
- Error: the sentinel classification cases fail because `errors.Is` cannot identify the expected leaf/duplicate error.
