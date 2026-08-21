# Bug reproduction

- Bug: key rotation leaves the state machine and registry pointing at the old/pending key, so promotion succeeds without making the new version current or usable.
- Trigger: transition through a rotation window, record the new key, promote it, and check current/usable state.
- Error: the target cases fail because the new key is not active/current and subsequent signing still uses the old version.
