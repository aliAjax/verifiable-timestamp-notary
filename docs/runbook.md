# Operations Runbook

Check `/healthz`, `/readyz`, `/metrics`, and `/debug/audit` before changing traffic. A quorum error is expected when fewer than the configured number of healthy signers are inside the clock-skew window. Restore signer health or route to a configured disaster-recovery signer set; do not lower quorum without an audited policy change.

For time rollback alerts, fence the affected signer, preserve its commitments, rotate its key, and verify all proofs issued during the disputed window. Proof revocation is explicit and leaves the original evidence immutable.
