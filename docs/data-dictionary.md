# Data Dictionary

`Claim.digest` is a lowercase hexadecimal digest; the service never stores the original file. `Batch.root` is the canonical Merkle root. `Checkpoint.previous_hash` links checkpoint records. `Proof.path` is an ordered sibling path. `SignerKey` carries validity and revocation windows. `AuditEvent.hash` commits the event and previous hash.
