package funcs

import (
	"context"
	"git.garena.com/shopee/loan-service/credit_backend/unicorn/unicorn-extension/src/Logger"
)

type ReadableTimeInfo interface {
	TimeInfo
	read(ctx context.Context, t *Tokenizer) []byte
}

func (_y *y) read(ctx context.Context, t *Tokenizer) []byte {
	return _y.readFixedLengthDigit(ctx, 2, t)
}

func (_Y *Y) read(ctx context.Context, t *Tokenizer) []byte {
	return _Y.readFixedLengthDigit(ctx, 4, t)
}

func (_c *c) read(ctx context.Context, t *Tokenizer) []byte {
	return _c.readVariableLengthDigit(ctx, 1, 2, t)
}

func (_m *m) read(ctx context.Context, t *Tokenizer) []byte {
	return _m.readFixedLength(ctx, 2, t)
}

func (_M *M) read(ctx context.Context, t *Tokenizer) []byte {
	return _M.readAlphabets(ctx, t)
}

func (_b *b) read(ctx context.Context, t *Tokenizer) []byte {
	return _b.readAlphabets(ctx, t)
}

func (_e *e) read(ctx context.Context, t *Tokenizer) []byte {
	return _e.readVariableLengthDigit(ctx, 1, 2, t)
}

func (_d *d) read(ctx context.Context, t *Tokenizer) []byte {
	return _d.readFixedLength(ctx, 2, t)
}

func (_T *Ti) read(ctx context.Context, t *Tokenizer) []byte {
	return _T.readFixedLength(ctx, 8, t)
}

func (_k *k) read(ctx context.Context, t *Tokenizer) []byte {
	return _k.readVariableLengthDigit(ctx, 1, 2, t)
}

func (_H *H) read(ctx context.Context, t *Tokenizer) []byte {
	return _H.readFixedLengthDigit(ctx, 2, t)
}

func (_l *l) read(ctx context.Context, t *Tokenizer) []byte {
	return _l.readVariableLengthDigit(ctx, 1, 2, t)
}

func (_h *h) read(ctx context.Context, t *Tokenizer) []byte {
	return _h.readFixedLengthDigit(ctx, 2, t)
}

func (_p *p) read(ctx context.Context, t *Tokenizer) []byte {
	return _p.readFixedLength(ctx, 2, t)
}

func (_i *i) read(ctx context.Context, t *Tokenizer) []byte {
	return _i.readFixedLengthDigit(ctx, 2, t)
}

func (_S *S) read(ctx context.Context, t *Tokenizer) []byte {
	return _S.readFixedLengthDigit(ctx, 2, t)
}

/**
 * @Author: feng.gao@shopee.com
 * @Description: A date string parsing tool. ver.3\n
 * this function take the first input as a date string, and parse the date time info according to the format which
 * defined in the second parameter. If parsed successfully, a unix timestamp (in second) will be returned. if any error
 * occurs, -1 will be returned.\n
 * *NOTE*: the timestamp will be calibrated based on the local time zone.\n
 * \n
 * continuous digit parsing behavior & known flaws:\n
 * when there's no explicit seperator between date time info element, this function will try to extract every part in
 * greedy mode. for instance, when parsing "2023111" with "%Y%c%e", it will be interpreted as 2023/11/01. if change the
 * order of %c and %e, then we will get 2023/01/11 instead.\n
 * a known flaw: when parsing "2023131" with "%Y%c%e", it should have been processed as 2023/01/31 without any ambiguity.
 * but this version will return -1\n
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
 *   %c: when in continuous digit situation, this control character will work in greedy mode. e.g. when find `11` in
 *       character stream in date string. this character will take it as November instead of January with a residual
 *       `1` left to the following control character to match.
 *   %e: will also work in greedy mode, alike to %c
 *   %k: will also work in greedy mode, alike to %c
 *   %l: will also work in greedy mode, alike to %c
 *   %y: if the value XX is small then 70, then it will be processed as 19XX. otherwise, 20XX.\n
 *\n
 * @Example 1: RF_StrToTimestamp("2023 11 08 16 42 40", "%Y %m %d %H %i %s") return 1699461760 (GMT)
 * @Example 2: RF_StrToTimestamp("6 2023 JanuaryX8 42 40 ", "%l %Y %MX%e %i %s") is equal to
 *                RF_StrToTimestamp("23 01 08%06:42:40", "%y %m%%%d %T")
 * @Example 3: RF_StrToTimestamp("PM", "%p") return 43200(GMT)
 * @Example 4: RF_StrToTimestamp("2023111 164240", "%Y%c%e %H%i%S") is equal to
 *                RF_StrToTimestamp("2023-11-01 16:42:40", "%Y-%m-%d %H:%i:%s")
 * @Param str: date string to be parsed
 * @Param format: pattern string, contains any number of control characters
 * @Return: timestamp in second if parsed successfully, or -1 if any error occurs
 */
func RF_StrToTimestamp3(ctx context.Context, str string, format string) int64 {
	time := Time{}
	t := &Tokenizer{
		b: []byte(str),
		c: 0,
		l: len(str),
	}
	var controlMode = false
	fs := []byte(format)
	for i := 0; i < len(fs); i++ {
		b := fs[i]
		if b == 32 {
			continue
		}
		if controlMode {
			if b == 37 {
				if t.match(ctx, b) {
					controlMode = false
					continue
				} else {
					Logger.Errorf(ctx, "unmatched character: %s", string(b))
					return -1
				}
			} else if timeInfo, ok := registration[string(b)].(ReadableTimeInfo); !ok {
				Logger.Errorf(ctx, "unknown control character: %s", string(b))
				return -1
			} else {
				info := string(timeInfo.read(ctx, t))
				if !timeInfo.valid(ctx, info) {
					Logger.Errorf(ctx, "error pattern: %s", info)
					return -1
				} else if !timeInfo.populate(ctx, info, &time) {
					Logger.Errorf(ctx, "duplicate definition: %s", info)
					return -1
				}
			}
		} else if b != 37 && !t.match(ctx, b) {
			Logger.Errorf(ctx, "unmatched character: %s", string(b))
			return -1
		}
		controlMode = b == 37
	}
	return int64(time.ToTimeStamp(ctx))
}

func (t *Tokenizer) skipSpace(ctx context.Context) {
	for ; t.c < t.l && t.b[t.c] == 32; t.c++ {
	}
}

var emptyBytes []byte

func (bti *BaseTimeInfo) readFixedLength(ctx context.Context, l int, t *Tokenizer) []byte {
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

func (bti *BaseTimeInfo) readFixedLengthDigit(ctx context.Context, l int, t *Tokenizer) []byte {
	if t.c == t.l {
		return emptyBytes
	}
	t.skipSpace(ctx)
	o := t.c
	for ; t.c-o < l && t.c < t.l && t.b[t.c] >= 48 /*`0`*/ && t.b[t.c] <= 57; /*`9`*/ t.c++ {
	}
	if t.c-o == l {
		r := t.b[o:t.c]
		t.skipSpace(ctx)
		return r
	} else {
		return emptyBytes
	}
}

func (bti *BaseTimeInfo) readVariableLengthDigit(ctx context.Context, m, n int, t *Tokenizer) []byte {
	if t.c == t.l {
		return emptyBytes
	}
	t.skipSpace(ctx)
	o := t.c
	for ; t.b[t.c] >= 48 /*`0`*/ && t.b[t.c] <= 57 && t.c-o < n; /*`9`*/ t.c++ {
	}
	if t.c-o >= m-1 {
		r := t.b[o:t.c]
		t.skipSpace(ctx)
		return r
	} else {
		return emptyBytes
	}
}

func (bti *BaseTimeInfo) readAlphabets(ctx context.Context, t *Tokenizer) []byte {
	if t.c == t.l {
		return emptyBytes
	}
	t.skipSpace(ctx)
	o := t.c
	for ; t.c < t.l && (t.b[t.c] >= 65 /*`A`*/ && t.b[t.c] <= 90 /*`Z`*/) ||
		(t.b[t.c] >= 97 /*`a`*/ && t.b[t.c] <= 123 /*`z`*/) || t.b[t.c] == 46; /*`.`*/ t.c++ {
	}
	r := t.b[o:t.c]
	t.skipSpace(ctx)
	return r
}
