package randx

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var (
	alpha = []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
	}
	alphaLen = len(alpha)

	digit = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}
	digitLen = len(digit)

	first = []string{
		"James",
		"John",
		"Robert",
		"Michael",
		"William",
		"David",
		"Richard",
		"Charles",
		"Joseph",
		"Thomas",
		"Christopher",
		"Daniel",
		"Paul",
		"Mark",
		"Donald",
		"George",
		"Kenneth",
		"Steven",
		"Edward",
		"Brian",
		"Ronald",
		"Anthony",
		"Kevin",
		"Jason",
		"Matthew",
		"Gary",
		"Timothy",
		"Jose",
		"Larry",
		"Jeffrey",
		"Frank",
		"Scott",
		"Eric",
		"Mary",
		"Patricia",
		"Linda",
		"Barbara",
		"Elizabeth",
		"Jennifer",
		"Maria",
		"Susan",
		"Margaret",
		"Dorothy",
		"Lisa",
		"Nancy",
		"Karen",
		"Betty",
		"Helen",
		"Sandra",
		"Donna",
		"Carol",
		"Ruth",
		"Sharon",
		"Michelle",
		"Laura",
		"Sarah",
		"Kimberly",
		"Deborah",
		"Jessica",
		"Shirley",
		"Cynthia",
		"Angela",
		"Melissa",
	}
	firstLen = len(first)

	last = []string{
		"Smith",
		"Johnson",
		"Williams",
		"Brown",
		"Jones",
		"Miller",
		"Davis",
		"Garcia",
		"Rodriguez",
		"Wilson",
		"Martinez",
		"Anderson",
		"Taylor",
		"Thomas",
		"Hernandez",
		"Moore",
		"Martin",
		"Jackson",
		"Thompson",
		"White",
		"Lopez",
		"Lee",
		"Gonzalez",
		"Harris",
		"Clark",
		"Lewis",
		"Robinson",
		"Walker",
		"Perez",
		"Hall",
		"Young",
		"Allen",
		"Sanchez",
		"Wright",
		"King",
		"Scott",
		"Green",
		"Baker",
		"Adams",
		"Nelson",
		"Hill",
		"Ramirez",
		"Campbell",
		"Mitchell",
		"Roberts",
		"Carter",
		"Phillips",
		"Evans",
		"Turner",
		"Torres",
		"Parker",
		"Collins",
		"Edwards",
		"Stewart",
		"Flores",
		"Morris",
		"Nguyen",
		"Murphy",
		"Rivera",
		"Cook",
		"Rogers",
		"Morgan",
		"Peterson",
		"Cooper",
		"Reed",
		"Bailey",
		"Bell",
		"Gomez",
		"Kelly",
		"Howard",
		"Ward",
		"Cox",
		"Diaz",
		"Richardson",
		"Wood",
		"Watson",
		"Brooks",
		"Bennett",
		"Gray",
		"James",
		"Reyes",
		"Cruz",
		"Hughes",
		"Price",
		"Myers",
		"Long",
		"Foster",
		"Sanders",
		"Ross",
		"Morales",
		"Powell",
		"Sullivan",
		"Russell",
	}
	lastLen = len(last)
)

func IntN(n int) int {
	return r.Intn(n)
}

func Int63() int64 {
	return r.Int63()
}

func AlphaN(n int) string {
	v := make([]string, n)
	for i := 0; i < n; i++ {
		v[i] = alpha[IntN(alphaLen)]
	}
	return strings.Join(v, "")
}

func DigitN(n int) string {
	v := make([]string, n)
	for i := 0; i < n; i++ {
		v[i] = digit[IntN(digitLen)]
	}
	return strings.Join(v, "")
}

func First() string {
	return first[r.Intn(firstLen)]
}

func Last() string {
	return last[r.Intn(lastLen)]
}

func FullName() string {
	return First() + " " + Last()
}

func Email(segments ...string) string {
	if len(segments) == 0 {
		segments = append(segments, First(), Last())
	}

	email := fmt.Sprintf("%v+%v@example.com", strings.Join(segments, "."), Int63())
	return strings.ToLower(email)
}

func Url(segments ...string) string {
	if len(segments) == 0 {
		segments = append(segments, AlphaN(12))
	}

	segments = append(segments, strconv.FormatInt(Int63(), 10), "example.com")
	return "http://" + strings.ToLower(strings.Join(segments, "."))
}
