package hatena

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCount(t *testing.T) {

	Convey("Given target url", t, func() {
		cli := NewHatena(nil)
		u := "https://www.google.co.jp"

		Convey("When call", func() {
			count, err := cli.BookMarks.Count(u)

			Convey("Then return count", func() {
				t.Logf("err %v", err)
				t.Logf("result %v", count)

				So(err, ShouldBeNil)
				So(count, ShouldBeGreaterThan, 0)
			})
		})
	})
}
