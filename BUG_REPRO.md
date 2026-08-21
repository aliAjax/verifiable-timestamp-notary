# Bug reproduction

- Bug: concurrent snapshot reads and writes race, and a previously returned history slice can change after a later write.
- Trigger: run the concurrent snapshot isolation case while replacing snapshots and reading IDs/digests.
- Error: `WARNING: DATA RACE`; the failing run also reports that the history snapshot changed.
