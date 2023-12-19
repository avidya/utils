package funcs

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStrToTimestamp_1(t *testing.T) {
	ctx := context.Background()
	assert.True(t, true)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023 11 08 16 42 40", "%Y %m %d %H %i %s") > 1000000000)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023   ,,1 8 16 42 40", "%Y ,,%c %e %H %i %s") == RF_StrToTimestamp1(ctx, "23 01 08 16 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp1(ctx, "6 2023 January 8 42 40 ", "%l %Y %M%%e %i %s") == RF_StrToTimestamp1(ctx, "23 01 08%%06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp1(ctx, "06 2023 Jan. 8 42 40 ", "%h %Y %b %e %i %s") == RF_StrToTimestamp1(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp1(ctx, "2023,pm 1 8 6 42 40", "%Y,%p %c %e %k %i %s") == RF_StrToTimestamp1(ctx, "23 01 08 18 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp1(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp1(ctx, "PM", "%p") > 0)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%i") == -1)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023 11 08 16 424 40", "%Y%m%d%H%i%s") == -1)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023 11 08 16,42 40", "%Y%m%d%H,%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023a11-08V16:42:40", "%Ya%m-%dV%H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023,11 08 16 42 40", "%Y,,%m %d %H %i %s") == -1)
	assert.True(t, RF_StrToTimestamp1(ctx, "2023,11 08 16 42 40", "%Z,,%m %d %H %i %s") == -1)
}

func TestStrToTimestamp_2(t *testing.T) {
	ctx := context.Background()
	assert.True(t, true)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023 11 08 16 42 40", "%Y %m %d %H %i %s") > 1000000000)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023   ,,1 8 16 42 40", "%Y ,,%c %e %H %i %s") == RF_StrToTimestamp2(ctx, "23 01 08 16 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp2(ctx, "6 2023 January 8 42 40", "%l %Y %M %e %i %s") == RF_StrToTimestamp2(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp2(ctx, "06 2023 Jan 8 42 40", "%h %Y %b %e %i %s") == RF_StrToTimestamp2(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp2(ctx, "2023,PM 1 8 6 42 40", "%Y,%p %c %e %k %i %s") == RF_StrToTimestamp2(ctx, "23 01 08 18 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp2(ctx, "20231108164240", "%Y%m%d%H%i%s") > 1000000000)
	//assert.True(t, StrToTimestamp2("PM", "%p") == 43200)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%i") == -1)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023 11 08 16 424 40", "%Y%m%d%H%i%s") == -1)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023110816,4240", "%Y%m%d%H,%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023a11-08V16:42:40", "%Ya%m-%dV%H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023,11 08 16 42 40", "%Y,,%m %d %H %i %s") == -1)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023,11 08 16 42 40", "%Z,,%m %d %H %i %s") == -1)
	assert.True(t, RF_StrToTimestamp2(ctx, "2023111", "%Y%c%e") == RF_StrToTimestamp2(ctx, "2023 11 01", "%Y %m %d"))
	assert.True(t, RF_StrToTimestamp2(ctx, "2023111", "%Y%e%c") == RF_StrToTimestamp2(ctx, "2023 01 11", "%Y %m %d"))
	fmt.Println(RF_StrToTimestamp2(ctx, "2023292", "%Y%e%c"))
	fmt.Println(RF_StrToTimestamp2(ctx, "2023131", "%Y%c%e"))
	fmt.Println((&Time4{
		year:  2023,
		month: 13,
		date:  1,
	}).ToTimeStamp(ctx))
	fmt.Println((&Time4{
		year:  2023,
		month: 1,
		date:  31,
	}).ToTimeStamp(ctx))

	/* the following assertion only belong to StrToTimestamp2 */
	assert.True(t, RF_StrToTimestamp2(ctx, "2023111 164240", "%Y%c%e %H%i%S") == RF_StrToTimestamp3(ctx, "2023-11-01 16:42:40", "%Y-%m-%d %H:%i:%s"))

}

func TestStrToTimestamp_3(t *testing.T) {
	ctx := context.Background()
	//assert.True(t, true)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023 11 08 16 42 40", "%Y %m %d %H %i %s") > 1000000000)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023   ,,1 8 16 42 40", "%Y ,,%c %e %H %i %s") == RF_StrToTimestamp3(ctx, "23 01 08 16 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp3(ctx, "6 2023 January 8 42 40 ", "%l %Y %M %e %i %s") == RF_StrToTimestamp3(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp3(ctx, "06 2023 Jan. 8 42 40 ", "%h %Y %b %e %i %s") == RF_StrToTimestamp3(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp3(ctx, "2023,pm 1 8 6 42 40", "%Y,%p %c %e %k %i %s") == RF_StrToTimestamp3(ctx, "23 01 08 18 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp3(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp3(ctx, "PM", "%p") > 0)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%i") == -1)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023 11 08 16 424 40", "%Y%m%d%H%i%s") == -1)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023 11 08 16,42 40", "%Y%m%d%H,%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023a11-08V16:42:40", "%Ya%m-%dV%H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023,11 08 16 42 40", "%Y,,%m %d %H %i %s") == -1)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023,11 08 16 42 40", "%Z,,%m %d %H %i %s") == -1)
	/* the following assertion only belong to StrToTimestamp3 */
	assert.True(t, RF_StrToTimestamp3(ctx, "20231108164240", "%Y%m%d%H%i%S") == RF_StrToTimestamp3(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s"))
	assert.True(t, RF_StrToTimestamp3(ctx, "2023%1108164240", "%Y%%%m%d%H%i%S") == RF_StrToTimestamp3(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s"))
	assert.True(t, RF_StrToTimestamp3(ctx, "20231108164240", "%Y%%%m%d %H%i%S") == -1)
	assert.True(t, RF_StrToTimestamp3(ctx, "2023111 164240", "%Y%c%e %H%i%S") == RF_StrToTimestamp3(ctx, "2023-11-01 16:42:40", "%Y-%c-%d %H:%i:%s"))
	assert.True(t, RF_StrToTimestamp3(ctx, "2023111 164240", "%Y%e%c %H%i%S") == RF_StrToTimestamp3(ctx, "2023-01-11 16:42:40", "%Y-%m-%d %H:%i:%s"))
	//assert.True(t, StrToTimestamp3("2023131 164240", "%Y%c%e %H%i%S") == StrToTimestamp3("2023-01-31 16:42:40", "%Y-%m-%d %H:%i:%s"))

}

func TestStrToTimestamp_4(t *testing.T) {
	ctx := context.Background()
	assert.True(t, true)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023 11 08 16 42 40", "%Y %m %d %H %i %s") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023   ,,1 8 16 42 40", "%Y ,,%c %e %H %i %s") == RF_StrToTimestamp4(ctx, "23 01 08 16 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp4(ctx, "6 2023 January 8 42 40 ", "%l %Y %M %e %i %s") == RF_StrToTimestamp4(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "06 2023 Jan. 8 42 40 ", "%h %Y %b %e %i %s") == RF_StrToTimestamp4(ctx, "23 01 08 06:42:40", "%y %m %d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023,pm 1 8 6 42 40", "%Y,%p %c %e %k %i %s") == RF_StrToTimestamp4(ctx, "23 01 08 18 42 40", "%y %m %d %H %i %s"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "PM", "%p") > 0)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%i") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023 11 08 16,42 40", "%Y%m%d%H,%i%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023a11-08V16:42:40", "%Ya%m-%dV%H:%i:%s") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023,11 08 16 42 40", "%Y,,%m %d %H %i %s") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023,11 08 16 42 40", "%Z,,%m %d %H %i %s") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "20231108164240", "%Y%m%d%H%i%S") == RF_StrToTimestamp4(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023%1108164240", "%Y%%%m%d%H%i%S") == RF_StrToTimestamp4(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20231108164240", "%Y%%%m%d %H%i%S") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023111 164240", "%Y%c%e %H%i%S") == RF_StrToTimestamp4(ctx, "2023-11-01 16:42:40", "%Y-%c-%d %H:%i:%s"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023111 164240", "%Y%e%c %H%i%S") == RF_StrToTimestamp4(ctx, "2023-01-11 16:42:40", "%Y-%m-%d %H:%i:%s"))

	/* the following assertion only belong to StrToTimestamp4 */
	assert.True(t, RF_StrToTimestamp4(ctx, "2023 11 08 16 40 423", "%Y%m%d%H%s%i") == RF_StrToTimestamp4(ctx, "2023 11 08 16 42 40", "%Y%m%d%H%i%s"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023131 164240", "%Y%c%e %H%i%S") == RF_StrToTimestamp4(ctx, "2023-01-31 16:42:40", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20231101 4240", "%Y%c%e%H %i%S") == RF_StrToTimestamp4(ctx, "2023-01-01 01:42:40", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20231421140", "%Y%c%i%e%k%S") == RF_StrToTimestamp4(ctx, "2023-01-01 01:42:40", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20231121113", "%Y%c%i%e%k%S") == RF_StrToTimestamp4(ctx, "2023-01-01 01:12:13", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "202311211113", "%Y%c%i%e%k%S") == RF_StrToTimestamp4(ctx, "2023-11-01 01:21:13", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023112111113", "%Y%c%i%e%k%S") == RF_StrToTimestamp4(ctx, "2023-11-11 01:21:13", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20231121111113", "%Y%c%i%e%k%S") == RF_StrToTimestamp4(ctx, "2023-11-11 11:21:13", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20231111", "%Y%k%c%e") == RF_StrToTimestamp4(ctx, "2023-01-01 11:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023141", "%Y%c%e%k") == RF_StrToTimestamp4(ctx, "2023-01-04 01:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023141", "%Y%c%e%k") == RF_StrToTimestamp4(ctx, "2023-01-04 01:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "202432", "%Y%c%e") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "7032", "%y%c%e") > 2000000 && RF_StrToTimestamp4(ctx, "7032", "%y%c%e") < 9000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023Jan", "%Y%b") == RF_StrToTimestamp4(ctx, "2023-01-01 00:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20232111", "%Y%c%e%k") == RF_StrToTimestamp4(ctx, "2023-02-11 01:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "20232311", "%Y%c%e%k") == RF_StrToTimestamp4(ctx, "2023-02-03 11:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "202322901", "%Y%c%e%k%i") == RF_StrToTimestamp4(ctx, "2023-02-02 09:01:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023229", "%Y%c%e%k") == RF_StrToTimestamp4(ctx, "2023-02-02 09:00:00", "%Y-%m-%d %T"))

	assert.True(t, RF_StrToTimestamp4(ctx, "12023", "%e%Y") == RF_StrToTimestamp4(ctx, "2023-01-01 00:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "12023", "%e%Y") == RF_StrToTimestamp4(ctx, "2023-01-01 00:00:00", "%Y-%m-%d %T"))

	assert.True(t, RF_StrToTimestamp4(ctx, "2023366", "%Y%j") == RF_StrToTimestamp4(ctx, "2024-01-01 00:00:00", "%Y-%m-%d %T"))
	assert.True(t, RF_StrToTimestamp4(ctx, "2023133", "%Y%j%c%e") == RF_StrToTimestamp4(ctx, "2023-01-01 00:00:00", "%Y-%m-%d %T"))

	assert.True(t, RF_StrToTimestamp4(ctx, "2023-11-13 13:00:54", "%Y-%m-%d %T") > 1000000000)
	assert.True(t, RF_StrToTimestamp4(ctx, "11/13/2023", "%m/%d/%Y") > 1000000000)

	/*error case*/
	assert.True(t, RF_StrToTimestamp4(ctx, "2023-01-00 00:00:00", "%Y-%m-%d %T") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023", "%Y%e") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023", "%Y%M") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "12023", "%e%c%T") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023-01-01 0v:00:00", "%Y-%m-%d %T") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023-00-01 00:00:00", "%Y-%m-%d %T") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023-01-01 00:00:00pM", "%Y-%m-%d %T%p") == -1)
	assert.True(t, RF_StrToTimestamp4(ctx, "2023Januu", "%Y%b") == -1)

}

func TestFast(t *testing.T) {
	ctx := context.Background()
	time := Time{"1970", "01", "01", "08", "00", "00", false}
	assert.True(t, time.ToTimeStamp(ctx) == 28800)
	time = Time{"1971", "01", "01", "08", "00", "00", false}
	assert.True(t, time.ToTimeStamp(ctx) == 31564800)
	time = Time{"1994", "01", "01", "08", "00", "00", false}
	assert.True(t, time.ToTimeStamp(ctx) == 757411200)
	time = Time{"2023", "11", "18", "01", "20", "31", false}
	assert.True(t, time.ToTimeStamp(ctx) == 1700270431)
}

func TestGolangTimeFormat(t *testing.T) {
	//ttime, _ := time.Parse("2006-1-2 3:4:5", "2023-1-1 1:1:1")
	//fmt.Println(ttime.Unix())
	//assert.True(t, ttime.Unix() == RF_StrToTimestamp1(context.Background(), "2023-1-1 1:01:01", "%Y-%c-%e %k:%i:%s"))
	//
	//ttime2, _ := time.Parse("200612 3:4:5", "2023111 1:1:1")
	//assert.True(t, ttime2.Unix() == RF_StrToTimestamp1(context.Background(), "2023-11-1 1:01:01", "%Y-%c-%e %k:%i:%s"))
	//
	//ttime3, _ := time.Parse("200621 3:4:5", "2023111 1:1:1")
	//assert.True(t, ttime3.Unix() == RF_StrToTimestamp1(context.Background(), "2023-1-11 1:01:01", "%Y-%c-%e %k:%i:%s"))
	//
	//ttime4, _ := time.Parse("200612 3:4:5", "202311 1:1:1")
	//fmt.Println(ttime4.Unix())
	//
	//ttime5, _ := time.Parse("200612 3:4:5", "20231231 1:1:1")
	//fmt.Println(ttime5.Unix())
	//
	//ttime6, _ := time.Parse("200612 3:4:5", "20231111 1:1:1")
	//fmt.Println(ttime6.Unix())
}

func TestPerformance(t *testing.T) {
	ctx := context.Background()
	beginTime := time.Now()
	for i := 0; i < 1000000; i++ {
		RF_StrToTimestamp1(context.Background(), "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s")
	}
	fmt.Printf("StrToTimestamp cost: %d\n", time.Now().Sub(beginTime))

	beginTime = time.Now()
	for i := 0; i < 1000000; i++ {
		ttime, _ := time.Parse("2006-01-02 03:04:05", "2023-11-08 16:42:40")
		ttime.Unix()
	}
	fmt.Printf("time.Parse cost: %d\n", time.Now().Sub(beginTime))

	for i := 0; i < 1000000; i++ {
		RF_StrToTimestamp3(ctx, "2023-11-08 16:42:40", "%Y-%m-%d %H:%i:%s")
	}
	fmt.Printf("StrToTimestamp3 cost: %d\n", time.Now().Sub(beginTime))
}

//func TestPerformance2(t *testing.T) {
//	beginTime := time.Now()
//	for i := 0; i < 100000; i++ {
//		TestStrToTimestamp_1(t)
//	}
//	fmt.Printf("StrToTimestamp cost: %d\n", time.Now().Sub(beginTime))
//
//	beginTime = time.Now()
//	for i := 0; i < 100000; i++ {
//		TestStrToTimestamp_2(t)
//	}
//	fmt.Printf("StrToTimestamp2 cost: %d\n", time.Now().Sub(beginTime))
//
//	beginTime = time.Now()
//	for i := 0; i < 100000; i++ {
//		TestStrToTimestamp_3(t)
//	}
//	fmt.Printf("StrToTimestamp3 cost: %d\n", time.Now().Sub(beginTime))
//}
