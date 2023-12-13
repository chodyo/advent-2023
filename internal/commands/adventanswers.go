//nolint:gochecknoglobals,gochecknoinits // I like this better than having a wrapper function
package commands

var adventAnswers map[string]int

func init() {
	adventAnswers = map[string]int{
		`day 01 part 1 sample 1`: 142,
		`day 01 part 1 sample 2`: 209,
		`day 01 part 1 full`:     56042,
		`day 01 part 2 sample 1`: 142,
		`day 01 part 2 sample 2`: 281,
		`day 01 part 2 full`:     55358,

		`day 02 part 1 sample`: 8,
		`day 02 part 1 full`:   2771,
		`day 02 part 2 sample`: 2286,
		`day 02 part 2 full`:   70924,

		`day 03 part 1 sample`: 4361,
		`day 03 part 1 full`:   553825,
		`day 03 part 2 sample`: 467835,
		`day 03 part 2 full`:   93994191,

		`day 04 part 1 sample`: 13,
		`day 04 part 1 full`:   18519,
		`day 04 part 2 sample`: 30,
		`day 04 part 2 full`:   11787590,
	}
}
