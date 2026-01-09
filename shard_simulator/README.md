# Shard Simulator
This is meant to simulate and benchmark various shard mapping techniques.
It is split into 2 phases:
 - Phase 1: This is purely mathematical and programmatic and uses code and concurrency to simulate shards and test metrics such as: skew and movement.
 - Phase 2: This uses real databases for benchmarking to measure network I/O
 
 The algorithms benchmarked are:
 	- Modulo Hashing
  - Consistent Hashing
  - Rendezvous
  - Range based Hashing
  - Directory based Hashing
  - Geo based Hashing
  