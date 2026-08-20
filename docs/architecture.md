# Architecture Notes

The service separates claim lifecycle from Merkle construction, signer quorum, checkpoint chaining, proof verification, key lifecycle, evidence indexing, and audit compliance. Ports are represented by Go interfaces so PostgreSQL, KMS, HSM, object storage, and external anchor implementations can be added without changing the domain.

Threat model: digest validation prevents malformed inputs; body limits prevent request flooding; nonce/idempotency prevents replay; clock windows isolate unhealthy signers; quorum prevents a single signer from issuing a proof; audit hash links detect tampering. Production mTLS/OIDC middleware is an adapter boundary.

Recovery procedure: stop batch workers, load the last checkpoint cursor, verify the checkpoint chain, replay only claims in pending/aggregating states, reject duplicate signer commitments, and publish a new checkpoint after quorum recovery.
