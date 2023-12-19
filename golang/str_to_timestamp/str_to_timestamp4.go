package funcs

import (
	"context"
	"fmt"
	"git.garena.com/shopee/loan-service/credit_backend/unicorn/unicorn-extension/src/Logger"
	"reflect"
)

type RestorableIntTimeInfo interface {
	readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error)
	popInfo(ctx context.Context, info interface{}, time *Time4)
}

type TerminalSymbol interface {
	na(ctx context.Context)
}

func (_M *M) na(ctx context.Context) {}
func (_b *b) na(ctx context.Context) {}
func (_p *p) na(ctx context.Context) {}

type Time4 struct {
	year   int
	days   int
	isLeap bool
	month  int
	date   int
	hour   int
	minute int
	second int
	pm     bool
}

func (t *Time4) ToTimeStamp(ctx context.Context) int {

	result := 0

	yg := t.year - 1970
	result += yg*365 + yg/4 - yg/100 + yg/400
	t.isLeap = (t.year%4 == 0 && t.year%100 != 0) || t.year%400 == 0

	if t.days > 0 {
		result += t.days - 1
	} else {
		for i := 0; i < t.month-1; i++ {
			result += __month[i]
		}

		result += t.date - 1
		if t.isLeap && t.month > 2 {
			result += 1
		}
	}
	result *= 86400

	result += t.hour * 3600
	/**
	`%p` control pattern, is used to plus another 12 hours to the hour part
	*/
	if t.pm && t.hour < 13 {
		result += 12 * 3600
	}

	result += t.minute * 60
	result += t.second

	return result
}

func genericReturn(ctx context.Context, result interface{}) (interface{}, error) {
	if result == -1 {
		return -1, fmt.Errorf("error")
	} else {
		return result, nil
	}
}

func (_y *y) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _y.readFixedLengthDigitInt(ctx, 2, 100, t))
}

func (_y *y) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	if i > 69 {
		time.year = 1900 + i
	} else {
		time.year = 2000 + i
	}
}

func (_Y *Y) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _Y.readFixedLengthDigitInt(ctx, 4, 10000, t))
}

func (_Y *Y) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.year = i
}

func (_c *c) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _c.readVariableLengthDigitInt(ctx, 1, 2, 1, 12, false, t))
}

func (_c *c) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.month = i
}

func (_m *m) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	if i := _m.readFixedLengthDigitInt(ctx, 2, 13, t); i < 1 {
		return -1, fmt.Errorf("error")
	} else {
		return i, nil
	}
}

func (_m *m) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.month = i
}

var _MMapInt = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func (_M *M) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _M.readAlphabetsInt(ctx, _MMapInt, t))
}

func (_M *M) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.month = i
}

var _bMapInt = map[string]int{
	"Jan":   1,
	"Jan.":  1,
	"Feb":   2,
	"Feb.":  2,
	"Mar":   3,
	"Mar.":  3,
	"Apr":   4,
	"Apr.":  4,
	"May":   5,
	"May.":  5,
	"Jun":   6,
	"Jun.":  6,
	"June":  6,
	"Jul":   7,
	"Jul.":  7,
	"July":  7,
	"Aug":   8,
	"Aug.":  8,
	"Sep":   9,
	"Sep.":  9,
	"Sept.": 9,
	"Oct":   10,
	"Oct.":  10,
	"Nov":   11,
	"Nov.":  11,
	"Dec":   12,
	"Dec.":  12,
}

func (_b *b) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _b.readAlphabetsInt(ctx, _bMapInt, t))
}

func (_b *b) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.month = i
}

func (_e *e) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _e.readVariableLengthDigitInt(ctx, 1, 2, 1, 31, true, t))
}

func (_e *e) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.date = i
}

func (_d *d) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	if i := _d.readFixedLengthDigitInt(ctx, 2, 32, t); i < 1 {
		return -1, fmt.Errorf("error")
	} else {
		return i, nil
	}
}

func (_d *d) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.date = i
}

func (_k *k) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _k.readVariableLengthDigitInt(ctx, 1, 2, 0, 23, false, t))
}

func (_k *k) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.hour = i
}

func (_H *H) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _H.readFixedLengthDigitInt(ctx, 2, 24, t))
}

func (_H *H) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.hour = i
}

func (_l *l) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _l.readVariableLengthDigitInt(ctx, 1, 2, 1, 12, false, t))
}

func (_l *l) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.hour = i
}

func (_h *h) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _h.readFixedLengthDigitInt(ctx, 2, 13, t))
}

func (_h *h) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.hour = i
}

func (_i *i) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _i.readFixedLengthDigitInt(ctx, 2, 60, t))
}

func (_i *i) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.minute = i
}

func (_S *S) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _S.readFixedLengthDigitInt(ctx, 2, 60, t))
}

func (_S *S) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.second = i
}

func (_T *Ti) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	b := _T.readFixedLength4(ctx, 8, t)
	if len(b) == 8 &&
		b[0] > 47 && b[0] < 51 && b[1] > 47 && b[1] < 58 && ((b[0]-48)*10+b[1]-48) < 24 &&
		b[3] > 47 && b[3] < 54 && b[4] > 47 && b[4] < 58 &&
		b[6] > 47 && b[6] < 54 && b[7] > 47 && b[7] < 58 &&
		b[2] == 58 /*`:`*/ && b[5] == 58 {
		return b, nil
	} else {
		return nil, fmt.Errorf("")
	}
}

func (_T *Ti) popInfo(ctx context.Context, info interface{}, time *Time4) {
	b, _ := info.([]byte)
	time.hour = int((b[0]-48)*10 + b[1] - 48)
	time.minute = int((b[3]-48)*10 + b[4] - 48)
	time.second = int((b[6]-48)*10 + b[7] - 48)
}

func (_p *p) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	s := string(_p.readFixedLength4(ctx, 2, t))
	if s == "PM" || s == "pm" || s == "AM" || s == "am" {
		return s, nil
	} else {
		return nil, fmt.Errorf("")
	}
}

func (_p *p) popInfo(ctx context.Context, info interface{}, time *Time4) {
	s, _ := info.(string)
	time.pm = s == "PM" || info == "pm"
}

type j struct {
	BaseTimeInfo
}

func (_j *j) readInfo(ctx context.Context, t *Tokenizer4) (interface{}, error) {
	return genericReturn(ctx, _j.readVariableLengthDigitInt(ctx, 1, 3, 1, 999, false, t))
}

func (_j *j) popInfo(ctx context.Context, info interface{}, time *Time4) {
	i, _ := info.(int)
	time.days = i
}

type InfoResult struct {
	index    int
	timeInfo RestorableIntTimeInfo
	info     interface{}
}

var registration4 = map[string]RestorableIntTimeInfo{
	"y": &y{},
	"Y": &Y{},
	"c": &c{},
	"m": &m{},
	"M": &M{},
	"e": &e{},
	"d": &d{},
	"T": &Ti{},
	"p": &p{},
	"j": &j{},
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
 * @Description: A date string parsing tool. ver.4\n
 * this function take the first input as a date string, and parse the date time info according to the format which
 * defined in the second parameter. If parsed successfully, a unix timestamp (in second) will be returned. if any error
 * occurs, -1 will be returned.\n
 * *NOTE*: the timestamp will be calibrated based on the local time zone.\n
 * \n
 * continuous digit parsing behavior:\n
 * when there's no explicit seperator between date time info element, this function will try to extract every part in
 * greedy mode. for instance, when parsing "20231111" with "%Y%c%e%k", it will be interpreted as 2023/11/01 01:00:00.
 * if switch the position of %k and %c, then we will get 2023/01/01 11:00:00 instead.\n
 * month/date correlation:\n
 * suppose the format string has substring like "%Y%c%e…", when encounter the input like "2023229", `%c` will get `2`,
 * `%e` will get `2` also, since `2023` is not a leap year, and the residual `9` will be left to the following control
 * character to match.\n
 * \n
 * this function will try to parse the complete date time info based on the format string, if any date part missing, the
 * corresponding default value will be taken from the `1970 01 01 00:00:00`. it's to say, after "11 08" get parsed by
 * "%m %d", the hour, minute, second part will all be stuffed by 0, and the year will be set 1970 by default.\n
 * \n
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
 * | %j | Day of year (001 .. 999)                 |\n
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
 *   `%b`: all the valid input includes: Jan, Jan., Feb, Feb., Mar, Mar., Apr, Apr., May, May., Jun, Jun., June, Jul, Jul.
 *       July, Aug, Aug., Sep, Sep., Sept. Oct, Oct., Nov, Nov., Dec, Dec.\n
 *   `%c`: when in continuous digit situation, this control character will work in greedy mode. e.g. when find `11` in
 *       character stream in date string. this character will take it as November instead of January with a residual
 *       `1` left to the following control character to match.\n
 *   `%e`: will also work in greedy mode, alike to %c\n
 *   `%j`: will also work in greedy mode, alike to %c. NOTE: \n
 *       1.) this control character will override the infos matched by month and date.\n
 *       2.) the XXX matched by this control character can actually great than 365. if this overflow situation occurs,
 *       the minuend will be counted to the next year. e.g. after `2023366` being parsed by `%Y%j`, it will be processed
 *       as the equivalent `2024/01/01`.\n
 *   `%k`: will also work in greedy mode, alike to %c\n
 *   `%l`: will also work in greedy mode, alike to %c\n
 *   `%p`: can be used standalone, and also can be composed with any hour control character, if %p successfully match
 *       `PM` or `pm` in date string, and meanwhile, the matched hour is less than 12, then additional 12 hours will be
 *       added. otherwise, this control character will just be neglected.\n
 *   `%y`: if the value XX is less than 70, then it will be processed as 19XX. otherwise, 20XX.\n
 *\n
 * @Example 1: RF_StrToTimestamp("2023 11 08 16 42 40", "%Y %m %d %H %i %s") return 1699461760 (GMT)
 * @Example 2: RF_StrToTimestamp("6 2023 JanuaryX8 42 40 ", "%l %Y %MX%e %i %s") is equal to\n
 *                RF_StrToTimestamp("23 01 08%06:42:40", "%y %m%%%d %T")
 * @Example 3: RF_StrToTimestamp("PM", "%p") return 43200(GMT)
 * @Example 4: RF_StrToTimestamp("2023131 164240", "%Y%c%e %H%i%S") is equal to\n
 *                RF_StrToTimestamp("2023-01-31 16:42:40", "%Y-%m-%d %T"))
 * @Example 5: RF_StrToTimestamp("20231101 4240", "%Y%c%e%H %i%S") is equal to\n
 *                RF_StrToTimestamp("2023-01-01 01:42:40", "%Y-%m-%d %T")
 * @Example 6: RF_StrToTimestamp("20231121113", "%Y%c%i%e%k%S") is equal to\n
 *                RF_StrToTimestamp("2023-01-01 01:12:13", "%Y-%m-%d %T")
 * @Example 7: RF_StrToTimestamp("202311211113", "%Y%c%i%e%k%S") is equal to\n
 *                RF_StrToTimestamp("2023-11-01 01:21:13", "%Y-%m-%d %T")
 * @Example 8: RF_StrToTimestamp("202322901", "%Y%c%e%k%i") is equal to\n
 *                RF_StrToTimestamp("2023-02-02 09:01:00", "%Y-%m-%d %T")
 * @Example 9: RF_StrToTimestamp("2023133", "%Y%j%c%e") is equal to\n
 *                RF_StrToTimestamp("2023-01-01 00:00:00", "%Y-%m-%d %T"))
 * @Param str: date string to be parsed
 * @Param format: pattern string, contains any number of control characters
 * @Return: timestamp in second if parsed successfully, or -1 if any error occurs
 */
func RF_StrToTimestamp4(ctx context.Context, str string, format string) int64 {
	time := Time4{
		year:  1970,
		month: 1,
		date:  1,
	}
	t := &Tokenizer4{
		b:      []byte(str),
		l:      len(str),
		rpBuff: []*RestorePoint{},
	}
	var controlMode bool
	fs := []byte(format)
	result := []InfoResult{}

	for ; t.i < len(fs); t.i++ {

		if t.lastestRP != nil {

			t.i = t.lastestRP.i  //reset outer index
			t.c = t.lastestRP.c1 //reset inner index

			// reset the Restore Point buffer, remove the RPs behind the specified one.
			for di, rp := range t.rpBuff {
				if rp.i == t.lastestRP.i && rp.c1 == t.lastestRP.c1 && rp.c2 == t.lastestRP.c2 {
					t.rpBuff = t.rpBuff[:di]
					break
				}
			}

			// reset the result also.
			for si, info := range result {
				if info.index >= t.i {
					result = result[:si]
					break
				}
			}
		}

		b := fs[t.i]
		if b == 32 /*space*/ {
			continue
		}
		if controlMode {
			if b == 37 /*`%`*/ {
				if ok, _b := t.match(ctx, b); ok { // in this case means we meet another `%`, so, switch the `controlMode` off
					controlMode = false
					continue
				} else {
					Logger.Errorf(ctx, "unmatched character: %d, `%%` is expected", _b)
					return -1
				}
			} else if timeInfo, ok := registration4[string(b)]; !ok {
				Logger.Errorf(ctx, "unknown control character: %s", string(b))
				return -1
			} else {
				if info, err := timeInfo.readInfo(ctx, t); err != nil {
					if _, ok := timeInfo.(TerminalSymbol); len(t.rpBuff) == 0 || ok {
						Logger.Errorf(ctx, "invalid pattern... ")
						return -1
					} else {
						t.lastestRP = t.rpBuff[len(t.rpBuff)-1]
						t.i-- // it's not an elegant, but a brain-less way to avoid escaping from the loop
						continue
					}
				} else {
					result = append(result, InfoResult{t.i, timeInfo, info})
					t.lastestRP = nil
				}
			}
		} else if b != 37 {
			if ok, _b := t.match(ctx, b); !ok {
				Logger.Errorf(ctx, "unmatched character: %d, `%d` is expected", _b, b)
				return -1
			}
		}
		controlMode = b == 37

		if t.i == len(fs)-1 {
			dedup := map[reflect.Type]interface{}{}
			for _, r := range result {
				r.timeInfo.popInfo(ctx, r.info, &time)
				if _, ok := dedup[reflect.TypeOf(r.timeInfo)]; ok {
					return -1
				} else {
					dedup[reflect.TypeOf(r.timeInfo)] = struct {
					}{}
				}
			}
			if !monthDateCheck(ctx, &time) {
				for _, rp := range t.rpBuff {
					if rp.isDate {
						t.lastestRP = rp
					}
				}
				controlMode = true
				t.i-- //again
			}
		}
	}

	return int64(correctTime(ctx, time.ToTimeStamp(ctx)))
}

func monthDateCheck(ctx context.Context, time *Time4) bool {
	switch time.month {
	case 4, 6, 9, 11:
		return time.date < 31
	case 2:
		return time.date < 29 || time.isLeap && time.date == 29
	default:
		return time.date < 32
	}
}

type RestorePoint struct {
	i      int  /*index to track the outer loop*/
	c1     int  /*index to begin*/
	c2     int  /*boundary index*/
	isDate bool /*date has to be treated in a very unusual way*/
}

type Tokenizer4 struct {
	b         []byte
	c         int /*index to track the inner loop*/
	i         int /*index to track the outer loop*/
	l         int
	lastestRP *RestorePoint
	rpBuff    []*RestorePoint
}

func (t *Tokenizer4) skipSpace(ctx context.Context) {
	for ; t.c < t.l && t.b[t.c] == 32; t.c++ {
	}
}

func (t *Tokenizer4) match(ctx context.Context, s byte) (bool, byte) {
	b := byte(0)
	if t.c < len(t.b) {
		b := t.b[t.c]
		if b == s {
			t.c++
			return true, 0
		} else {
			return false, b
		}
	} else {
		return false, b
	}
}

func (bti *BaseTimeInfo) readFixedLength4(ctx context.Context, l int, t *Tokenizer4) []byte {
	t.skipSpace(ctx)
	if t.c+l > t.l {
		return emptyBytes
	} else {
		o := t.c
		t.c += l
		r := t.b[o:t.c]
		t.skipSpace(ctx)
		return r
	}
}

func (bti *BaseTimeInfo) readFixedLengthDigitInt(ctx context.Context, l, a int, t *Tokenizer4) int {
	if t.c == t.l {
		return -1
	}
	t.skipSpace(ctx)
	o := t.c
	r := -1
	for ; t.c-o < l && t.c < t.l; /*`9`*/ t.c++ {
		if t.b[t.c] >= 48 /*`0`*/ && t.b[t.c] <= 57 {
			if r == -1 {
				r = int(t.b[t.c]) - 48
			} else {
				r = r*10 + int(t.b[t.c]) - 48
			}
		} else {
			return -1
		}
	}

	if r < a && t.c-o == l {
		t.skipSpace(ctx)
		return r
	} else {
		t.skipSpace(ctx)
		return -1
	}
}

func (bti *BaseTimeInfo) readVariableLengthDigitInt(ctx context.Context, m, n, a, b int, isDate bool, t *Tokenizer4) int {
	if t.c == t.l { //hit the end, return -1 directly
		return -1
	}
	t.skipSpace(ctx)
	o := t.c
	r := -1
	for ; t.c < t.l && t.b[t.c] >= 48 /*`0`*/ && t.b[t.c] <= 57 /*`9`*/ && t.c-o < n; t.c++ {
		if t.lastestRP != nil && t.c == t.lastestRP.c2 { // to avoid retrying
			return r
		}
		if r == -1 {
			r = int(t.b[t.c]) - 48
		} else {
			r = r*10 + int(t.b[t.c]) - 48
		}
		if r >= a && r <= b {
			if r == 0 { // it's a very special case, can't be applied with the previous way..
				t.c++
				return 0
			}
			if t.c-o >= m && t.c-o < n {
				t.rpBuff = append(t.rpBuff, &RestorePoint{t.i, o, t.c, isDate})
			}
		} else if r > b {
			return (r - int(t.b[t.c]) + 48) / 10
		} else {
			return -1
		}
	}
	t.skipSpace(ctx)

	return r
}

func (bti *BaseTimeInfo) readAlphabetsInt(ctx context.Context, d map[string]int, t *Tokenizer4) int {
	if t.c == t.l {
		return -1
	}
	t.skipSpace(ctx)
	o := t.c
	for ; t.c < t.l && ((t.b[t.c] >= 65 /*`A`*/ && t.b[t.c] <= 90 /*`Z`*/) ||
		(t.b[t.c] >= 97 /*`a`*/ && t.b[t.c] <= 123 /*`z`*/) || t.b[t.c] == 46); /*`.`*/ t.c++ {
	}
	r := t.b[o:t.c]
	t.skipSpace(ctx)
	if v, ok := d[string(r)]; ok {
		return v
	} else {
		return -1
	}
}
