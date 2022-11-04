package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	things := []string{
		"twelve Drummers Drumming",
		"eleven Pipers Piping",
		"ten Lords-a-Leaping",
		"nine Ladies Dancing",
		"eight Maids-a-Milking",
		"seven Swans-a-Swimming",
		"six Geese-a-Laying",
		"five Gold Rings",
		"four Calling Birds",
		"three French Hens",
		"two Turtle Doves",
		"a Partridge in a Pear Tree",
	}
	days := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}

	thingsBrought := things[11]

	if i != 1 {
		thingsBrought = strings.Join([]string{strings.Join(things[12-i:11], ", "), thingsBrought}, ", and ")
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[i-1], thingsBrought)
}

func Song() string {
	o := make([]string, 12)
	for i := 1; i <= 12; i++ {
		o[i-1] = Verse(i)
	}
	return strings.Join(o, "\n")
}
