# Bug reproduction

- Bug: checkpoint recovery and replay keep using a canceled caller context, so canceled work can append or publish after cancellation.
- Trigger: cancel the context before recovery, append, anchor publish, and replay-session calls.
- Error: the cancellation cases fail because canceled operations still report success or leave replay state behind.
