module remo/dsa

go 1.21.6

replace remo/bubble => ./sorts/bubble

replace remo/insertion => ./sorts/insertion/

replace remo/selection => ./sorts/selection

replace remo/quick => ./sorts/quick

replace remo/merge => ./sorts/merge

require (
	remo/binary v0.0.0-00010101000000-000000000000
	remo/bubble v0.0.0-00010101000000-000000000000
	remo/heap v0.0.0-00010101000000-000000000000
	remo/insertion v0.0.0-00010101000000-000000000000
	remo/merge v0.0.0-00010101000000-000000000000
	remo/quick v0.0.0-00010101000000-000000000000
	remo/selection v0.0.0-00010101000000-000000000000
)

replace remo/heap => ./sorts/heap

replace remo/binary => ./search/binary/
