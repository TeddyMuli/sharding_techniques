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
  
  The results can be found in the visualization Directory
  
## Requirements
  - Docker
  - Golang 1.23 or later
  - Python
  - [uv](https://docs.astral.sh/uv/guides/install-python/)
  
## How to Run
1. Clone the repo
```bash
git clone https://github.com/TeddyMuli/sharding_techniques
cd sharding_techniques
```
2. Create a virtual environment
```bash
uv venv virt
source virt/bin/activate
```
3. Install go dependcies
```bash
go mod tidy
```
4. Install python dependcies
```bash
uv pip install -r requirements.txt
```
5. Run phase 1
```bash
go run -race cmd/phase_1/main.go
```
6. Run phase 2
```bash
go run -race cmd/phase_2/main.go
```
7. Create visualizations
```bash
cd visualization
./phase_1/visualize.py
./phase_2/visualize.py
```
8. View the generate images
