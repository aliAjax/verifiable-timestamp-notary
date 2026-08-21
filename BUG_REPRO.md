# Bug reproduction

- Bug: cancellation is ignored across the quorum collection window and protocol decode path, allowing late commitments or counts after the window is canceled.
- Trigger: cancel the caller context before adding a commitment, collecting, checking readiness, decoding, and closing the window.
- Error: the cancellation cases fail because late work is accepted or the canceled window still reports a count/ready state.
