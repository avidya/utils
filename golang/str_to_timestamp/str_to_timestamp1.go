package funcs

import (
	"bytes"
	"context"
	"git.garena.com/shopee/loan-service/credit_backend/unicorn/unicorn-extension/src/Logger"
	"strconv"
	"strings"
	"time"
)

type TimeInfo interface {
	valid(ctx context.Context, info string) bool
	populate(ctx context.Context, info string, time *Time) bool
}

type BaseTimeInfo struct {
}

func (bti *BaseTimeInfo) dPopulate(ctx context.Context, info string, dict map[string]string, timeField *string) bool {
	if len(*timeField) > 0 {
		return false
	} else {
		v, _ := dict[info]
		*timeField = v
		return true
	}
}

func (bti *BaseTimeInfo) pPopulate(ctx context.Context, info string, timeField *string) bool {
	if len(*timeField) > 0 {
		return false
	} else {
		if len(info) < 2 {
			*timeField = "0" + info
		} else {
			*timeField = info
		}
		return true
	}
}

type Year struct {
	BaseTimeInfo
}

func (year *Year) populate(ctx context.Context, info string, time *Time) bool {
	if len(time.year) > 0 {
		return false
	} else {
		if len(info) == 2 {
			if info > "69" {
				time.year = "19" + info
			} else {
				time.year = "20" + info
			}
		} else {
			time.year = info
		}
		return true
	}
}

type Month struct {
	BaseTimeInfo
}

func (month *Month) populate(ctx context.Context, info string, time *Time) bool {
	return month.pPopulate(ctx, info, &time.month)
}

type Date struct {
	BaseTimeInfo
}

func (date *Date) populate(ctx context.Context, info string, time *Time) bool {
	return date.pPopulate(ctx, info, &time.date)
}

type Hour struct {
	BaseTimeInfo
}

func (hour *Hour) populate(ctx context.Context, info string, time *Time) bool {
	return hour.pPopulate(ctx, info, &time.hour)
}

type Minute struct {
	BaseTimeInfo
}

func (minute *Minute) populate(ctx context.Context, info string, time *Time) bool {
	return minute.pPopulate(ctx, info, &time.minute)
}

type Second struct {
	BaseTimeInfo
}

func (second *Second) populate(ctx context.Context, info string, time *Time) bool {
	return second.pPopulate(ctx, info, &time.second)
}

type y struct {
	Year
}

func (_y *y) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info > "01" && info < "99"
}

type Y struct {
	Year
}

func (_Y *Y) valid(ctx context.Context, info string) bool {
	return len(info) == 4 && info >= "0000" && info <= "9999"
}

type c struct {
	Month
}

func (_c *c) valid(ctx context.Context, info string) bool {
	return (len(info) == 1 && info >= "1" && info <= "9") || info == "10" || info == "11" || info == "12"
}

type m struct {
	Month
}

func (_m *m) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info >= "01" && info <= "12"
}

var _MMap = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

type M struct {
	BaseTimeInfo
}

func (_M *M) valid(ctx context.Context, info string) bool {
	_, ok := _MMap[info]
	return ok
}

func (_M *M) populate(ctx context.Context, info string, time *Time) bool {
	return _M.dPopulate(ctx, info, _MMap, &time.month)
}

var _bMap = map[string]string{
	"Jan":   "01",
	"Jan.":  "01",
	"Feb":   "02",
	"Feb.":  "02",
	"Mar":   "03",
	"Mar.":  "03",
	"Apr":   "04",
	"Apr.":  "04",
	"May":   "05",
	"May.":  "05",
	"Jun":   "06",
	"Jun.":  "06",
	"June":  "06",
	"Jul":   "07",
	"Jul.":  "07",
	"July":  "07",
	"Aug":   "08",
	"Aug.":  "08",
	"Sep":   "09",
	"Sep.":  "09",
	"Sept.": "09",
	"Oct":   "10",
	"Oct.":  "10",
	"Nov":   "11",
	"Nov.":  "11",
	"Dec":   "12",
	"Dec.":  "12",
}

type b struct {
	BaseTimeInfo
}

func (_b *b) valid(ctx context.Context, info string) bool {
	_, ok := _bMap[info]
	return ok
}

func (_b *b) populate(ctx context.Context, info string, time *Time) bool {
	return _b.dPopulate(ctx, info, _bMap, &time.month)
}

type e struct {
	Date
}

func (_e *e) valid(ctx context.Context, info string) bool {
	return (len(info) == 1 && info >= "1" && info <= "9") || (len(info) == 2 && info >= "10" && info <= "31")
}

type d struct {
	Date
}

func (_d *d) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info >= "01" && info <= "31"
}

type Ti struct {
	BaseTimeInfo
}

func (_T *Ti) valid(ctx context.Context, info string) bool {
	ts := strings.Split(info, ":")
	return len(ts) == 3 && ts[0] >= "00" && ts[0] < "24" && ts[1] >= "00" && ts[1] < "60" && ts[2] >= "00" && ts[2] < "60"
}

func (_T *Ti) populate(ctx context.Context, info string, time *Time) bool {
	ts := strings.Split(info, ":")
	if len(time.hour) > 0 {
		return false
	} else {
		time.hour = ts[0]
	}
	if len(time.minute) > 0 {
		return false
	} else {
		time.minute = ts[1]
	}
	if len(time.second) > 0 {
		return false
	} else {
		time.second = ts[2]
	}
	return true
}

type k struct {
	Hour
}

func (_k *k) valid(ctx context.Context, info string) bool {
	return (len(info) == 1 && info >= "0" && info <= "9") || (len(info) == 2 && info >= "10" && info < "24")
}

type H struct {
	Hour
}

func (_H *H) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info >= "00" && info < "24"
}

type l struct {
	Hour
}

func (_l *l) valid(ctx context.Context, info string) bool {
	return len(info) == 1 && info >= "0" && info <= "9" || info == "10" || info == "11" || info == "12"
}

type h struct {
	Hour
}

func (_h *h) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info >= "00" && info <= "12"
}

type p struct {
	BaseTimeInfo
}

func (_p *p) valid(ctx context.Context, info string) bool {
	return info == "AM" || info == "am" || info == "PM" || info == "pm"
}

func (_p *p) populate(ctx context.Context, info string, time *Time) bool {
	time.pm = info == "PM" || info == "pm"
	return true
}

type i struct {
	Minute
}

func (_i *i) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info >= "00" && info < "60"
}

type S struct {
	Second
}

func (_S *S) valid(ctx context.Context, info string) bool {
	return len(info) == 2 && info >= "00" && info < "60"
}

type Time struct {
	year   string
	month  string
	date   string
	hour   string
	minute string
	second string
	pm     bool
}

var __month = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func (t *Time) ToTimeStamp(ctx context.Context) int {

	y, _ := strconv.Atoi(t.year)
	mon, _ := strconv.Atoi(t.month)
	d, _ := strconv.Atoi(t.date)
	h, _ := strconv.Atoi(t.hour)
	minute, _ := strconv.Atoi(t.minute)
	s, _ := strconv.Atoi(t.second)

	if y == 0 {
		y = 1970
	}
	if mon == 0 {
		mon = 1
	}
	if d == 0 {
		d = 1
	}

	result := 0

	yg := y - 1970
	result += yg*365 + yg/4 - yg/100 + yg/400

	for i := 0; i < mon-1; i++ {
		result += __month[i]
	}

	result += d - 1
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 && mon > 2 {
		result += 1
	}

	result *= 86400

	result += h * 3600
	if t.pm {
		result += 12 * 3600
	}

	result += minute * 60
	result += s

	return result
}

var registration = map[string]TimeInfo{
	"y": &y{},
	"Y": &Y{},
	"c": &c{},
	"m": &m{},
	"M": &M{},
	"e": &e{},
	"d": &d{},
	"T": &Ti{},
	"p": &p{},
	"k": &k{},
	"H": &H{},
	"l": &l{},
	"h": &h{},
	"i": &i{},
	"S": &S{},
	"s": &S{},
	"b": &b{},
}

/**
 * @Author: feng.gao@shopee.com
 * @Description: A date string parsing tool. ver.1\n
 * this function take the first input as a date string, and parse the date time info according to the format which
 * defined in the second parameter. If parsed successfully, a unix timestamp (in second) will be returned. if any error
 * occurs, -1 will be returned.\n
 * *NOTE*: the timestamp will be calibrated based on the local time zone.\n
 *
 * all the date time info MUST BE expressed in splittable way. i.e. "2023 11 08" can be parsed by "%Y %m %d"(note the
 * space between every time info element.) and "2023-11-08 16:42:40" can be parsed by "%Y-%m-%d %H:%i:%s". the space
 * count doesn't have to matched precisely, which means "2023   11 08" can also be parsed by "%Y %m %d", besides that,
 * all the other character MUST BE matched exactly.\n
 *
 * this function will try to parse the complete date time info based on the format string, if any date part missing, the
 * corresponding default value will be taken from the `1970 01 01 00:00:00`. it's to say, after "11 08" get parsed by
 * "%m %d", the hour, minute, second part will all be stuffed by 0, and the year will be set 1970 by default.\n
 *
 * all the supported format control characters are listed below:\n
 * +----+------------------------------------------+\n
 * | %b | Abbreviated month name (Jan .. Dec)      |\n
 * +----+------------------------------------------+\n
 * | %c | Month, numeric (1 .. 12)                 |\n
 * +----+------------------------------------------+\n
 * | %d | Day of the month, numeric (01 .. 31)     |\n
 * +----+------------------------------------------+\n
 * | %e | Day of the month, numeric (1 .. 31)      |\n
 * +----+------------------------------------------+\n
 * | %H | Hour (00 .. 23)                          |\n
 * +----+------------------------------------------+\n
 * | %h | Hour (01 .. 12)                          |\n
 * +----+------------------------------------------+\n
 * | %i | Minutes, numeric (00 .. 59)              |\n
 * +----+------------------------------------------+\n
 * | %k | Hour (0 .. 23)                           |\n
 * +----+------------------------------------------+\n
 * | %l | Hour (1 .. 12)                           |\n
 * +----+------------------------------------------+\n
 * | %M | Month name (January .. December)         |\n
 * +----+------------------------------------------+\n
 * | %m | Month, numeric (01 .. 12)                |\n
 * +----+------------------------------------------+\n
 * | %p | AM/am or PM/pm                           |\n
 * +----+------------------------------------------+\n
 * | %S | Seconds (00 .. 59)                       |\n
 * +----+------------------------------------------+\n
 * | %s | Seconds (00 .. 59)                       |\n
 * +----+------------------------------------------+\n
 * | %T | Time, 24-hour (hh:mm:ss)                 |\n
 * +----+------------------------------------------+\n
 * | %Y | Year, numeric, four digits               |\n
 * +----+------------------------------------------+\n
 * | %y | Year, numeric (two digits)               |\n
 * +----+------------------------------------------+\n
 * | %% | A literal % character                    |\n
 * +----+------------------------------------------+\n
 * more detailed description:\n
 *   %b: all the valid input includes: Jan, Jan., Feb, Feb., Mar, Mar., Apr, Apr., May, May., Jun, Jun., June, Jul, Jul.
 *       July, Aug, Aug., Sep, Sep., Sept. Oct, Oct., Nov, Nov., Dec, Dec.\n
 *   %y: if the value XX is small then 70, then it will be processed as 19XX. otherwise, 20XX instead.\n\n
 * @Example 1: RF_StrToTimestamp("2023 11 08 16 42 40", "%Y %m %d %H %i %s") return 1699461760 (GMT)
 * @Example 2: RF_StrToTimestamp("6 2023 JanuaryX8 42 40 ", "%l %Y %MX%e %i %s") is equal to
 *                RF_StrToTimestamp("23 01 08%06:42:40", "%y %m%%%d %T")
 * @Example 3: RF_StrToTimestamp("PM", "%p") return 43200(GMT)
 * @Param str: date string to be parsed
 * @Param format: pattern string, contains any number of control characters
 * @Return: timestamp in second if parsed successfully, or -1 if any error occurs
 */
func RF_StrToTimestamp1(ctx context.Context, str string, format string) int64 {
	time := Time{}
	tokenizer := &Tokenizer{
		b: []byte(str),
		c: 0,
		l: len(str),
	}
	var controlMode = false
	fs := []byte(format)
	for i := 0; i < len(fs); i++ {
		b := fs[i]
		if controlMode {
			if timeInfo, ok := registration[string(b)]; !ok {
				Logger.Errorf(ctx, "unknown control character: %s", string(b))
				return -1
			} else {
				var info string
				if i == len(fs)-1 /*means touch the end*/ {
					info = tokenizer.nextToken(ctx, 32 /*wout a predefined EOF char... what a shame on u, golang*/)
				} else {
					if fs[i+1] == 37 {
						info = tokenizer.nextToken(ctx, 32)
					} else {
						//in this case, we should move the cursor manually
						i++
						info = tokenizer.nextToken(ctx, fs[i])
					}
				}
				if !timeInfo.valid(ctx, info) {
					Logger.Errorf(ctx, "error pattern: %s", info)
					return -1
				} else if !timeInfo.populate(ctx, info, &time) {
					Logger.Errorf(ctx, "duplicate definition: %s", info)
					return -1
				}
			}
		} else if b != 37 && !tokenizer.match(ctx, b) {
			Logger.Errorf(ctx, "unmatched character: %s", string(b))
			return -1
		}
		controlMode = b == 37
	}
	return int64(correctTime(ctx, time.ToTimeStamp(ctx)))
}

func correctTime(ctx context.Context, ts int) int {
	_, offset := time.Now().Zone()
	return ts + offset
}

type Tokenizer struct {
	b []byte
	c int
	l int
}

func (t *Tokenizer) match(ctx context.Context, s byte) bool {
	if t.c < len(t.b) && t.b[t.c] == s {
		t.c++
		return true
	} else {
		return false
	}
}

func (t *Tokenizer) nextToken(ctx context.Context, s byte) string {
	var buffer bytes.Buffer
	for ; t.c < len(t.b) && t.b[t.c] != s; t.c++ {
		buffer.WriteByte(t.b[t.c])
	}
	for ; t.c < len(t.b) && (t.b[t.c] == s || t.b[t.c] == 32); t.c++ {
	}
	return buffer.String()
}
