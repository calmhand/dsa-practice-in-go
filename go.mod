module remo/dsa

go 1.21.6

replace remo/bubble => ./sorts/Bubble-Sort

require (
	remo/bubble v0.0.0-00010101000000-000000000000
	remo/insertion v0.0.0-00010101000000-000000000000
	remo/selection v0.0.0-00010101000000-000000000000
)

replace remo/insertion => ./sorts/insertion/

replace remo/selection => ./sorts/selection
