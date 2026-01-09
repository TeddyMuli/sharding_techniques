package algorithms

var Competitors = []Sharder{
	NewModulo(),
	NewConsistent(),
	NewRange(10000),
	NewDirectory(),
	NewGeo(),
	NewRendezvous(),
}
