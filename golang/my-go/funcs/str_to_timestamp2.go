package funcs

import (
	"context"
	"strings"
	"time"
)

var translation = map[string]string{
	"%y": "06",
	"%Y": "2006",
	"%c": "1",
	"%m": "01",
	"%M": "January",
	"%e": "2",
	"%d": "02",
	"%T": "15:04:05",
	"%p": "PM",
	"%k": "3",
	"%H": "15",
	"%l": "3",
	"%h": "03",
	"%i": "04",
	"%S": "05",
	"%s": "05",
	"%b": "Jan",
}

/**
 * @Author: feng.gao@shopee.com
 * @Description: A date string parsing tool. ver.2\n
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
func RF_StrToTimestamp2(ctx context.Context, str string, format string) int64 {
	layout := convert(ctx, format)
	if t, e := time.Parse(layout, str); e != nil {
		return -1
	} else {
		return t.Unix()
	}
}

func convert(ctx context.Context, format string) string {
	for k, v := range translation {
		format = strings.Replace(format, k, v, -1)
	}
	return format
}
