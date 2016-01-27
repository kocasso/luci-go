// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package prpc

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/luci/luci-go/common/clock/testclock"

	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDecoding(t *testing.T) {
	t.Parallel()

	Convey("requestFormat", t, func() {
		test := func(contentType string, expectedFormat format, expectedErr interface{}) {
			Convey("Content-Type: "+contentType, func() {
				actualFormat, err := requestFormat(contentType)
				So(err, ShouldErrLike, expectedErr)
				if err == nil {
					So(actualFormat, ShouldEqual, expectedFormat)
				}
			})
		}

		test("", formatBinary, nil)
		test(mtPRPC, formatBinary, nil)
		test(mtPRPCBinary, formatBinary, nil)
		test(mtPRPCJSNOPB, formatJSONPB, nil)
		test(mtPRPCText, formatText, nil)
		test(
			mtPRPC+"; encoding=blah",
			0,
			`encoding parameter: invalid value "blah". Valid values: "json", "binary", "text"`)
		test(mtPRPC+"; boo=true", 0, `unexpected parameter "boo"`)

		test(mtJSON, formatJSONPB, nil)
		test(mtJSON+"; whatever=true", formatJSONPB, nil)

		test("x", 0, `"x" is not supported`)
		test("x,y", 0, "mime: expected slash after first token")
	})

	Convey("readMessage", t, func() {
		var msg HelloRequest
		read := func(contentType string, body []byte) *protocolError {
			req := &http.Request{
				Body:   ioutil.NopCloser(bytes.NewBuffer(body)),
				Header: http.Header{},
			}
			req.Header.Set("Content-Type", contentType)
			return readMessage(req, &msg)
		}

		testLucy := func(contentType string, body []byte) {
			err := read(contentType, body)
			So(err, ShouldBeNil)
			So(msg.Name, ShouldEqual, "Lucy")
		}

		Convey("binary", func() {
			testMsg := &HelloRequest{Name: "Lucy"}
			body, err := proto.Marshal(testMsg)
			So(err, ShouldBeNil)

			Convey(mtPRPC, func() {
				testLucy(mtPRPC, body)
			})
			Convey(mtPRPCBinary, func() {
				testLucy(mtPRPCBinary, body)
			})
			Convey("malformed body", func() {
				err := read(mtPRPCBinary, []byte{0})
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
			Convey("empty body", func() {
				err := read(mtPRPCBinary, nil)
				So(err, ShouldBeNil)
			})
		})

		Convey("json", func() {
			body := []byte(`{"name": "Lucy"}`)
			Convey(mtJSON, func() {
				testLucy(mtJSON, body)
			})
			Convey(mtPRPCJSNOPB, func() {
				testLucy(mtPRPCJSNOPB, body)
			})
			Convey("malformed body", func() {
				err := read(mtPRPCJSNOPB, []byte{0})
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
			Convey("empty body", func() {
				err := read(mtPRPCJSNOPB, nil)
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("text", func() {
			Convey(mtPRPCText, func() {
				body := []byte(`name: "Lucy"`)
				testLucy(mtPRPCText, body)
			})
			Convey("malformed body", func() {
				err := read(mtPRPCText, []byte{0})
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
			Convey("empty body", func() {
				err := read(mtPRPCText, nil)
				So(err, ShouldBeNil)
			})
		})

		Convey("unsupported media type", func() {
			err := read("blah", nil)
			So(err, ShouldNotBeNil)
			So(err.status, ShouldEqual, http.StatusUnsupportedMediaType)
		})
	})

	Convey("parseHeader", t, func() {
		c := context.Background()

		header := func(name, value string) http.Header {
			return http.Header{
				name: []string{value},
			}
		}

		Convey(headerTimeout, func() {
			Convey("Works", func() {
				now := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
				c, _ = testclock.UseTime(c, now)

				c, err := parseHeader(c, header(headerTimeout, "1M"))
				So(err, ShouldBeNil)

				deadline, ok := c.Deadline()
				So(ok, ShouldBeTrue)
				So(deadline, ShouldHappenWithin, time.Second, now.Add(time.Minute))
			})

			Convey("Fails", func() {
				c2, err := parseHeader(c, header(headerTimeout, "blah"))
				So(c2, ShouldEqual, c)
				So(err, ShouldErrLike, headerTimeout+` header: unit is not recognized: "blah"`)
			})
		})

		Convey("Content-Type", func() {
			c, err := parseHeader(c, header("Content-Type", "blah"))
			So(err, ShouldBeNil)
			_, ok := metadata.FromContext(c)
			So(ok, ShouldBeFalse)
		})

		Convey("Accept", func() {
			c, err := parseHeader(c, header("Accept", "blah"))
			So(err, ShouldBeNil)
			_, ok := metadata.FromContext(c)
			So(ok, ShouldBeFalse)
		})

		Convey("Unrecognized headers", func() {
			test := func(c context.Context, header http.Header, expectedMetadata metadata.MD) {
				c, err := parseHeader(c, header)
				So(err, ShouldBeNil)
				md, ok := metadata.FromContext(c)
				So(ok, ShouldBeTrue)
				So(md, ShouldResembleV, expectedMetadata)
			}

			headers := http.Header{
				"X": []string{"1"},
				"Y": []string{"1", "2"},
			}

			Convey("without metadata in context", func() {
				test(c, headers, metadata.MD(headers))
			})

			Convey("with metadata in context", func() {
				c = metadata.NewContext(c, metadata.MD{
					"X": []string{"0"},
					"Z": []string{"1"},
				})
				test(c, headers, metadata.MD{
					"X": []string{"0", "1"},
					"Y": []string{"1", "2"},
					"Z": []string{"1"},
				})
			})

			Convey("binary", func() {
				Convey("Works", func() {
					const name = "Lucy"
					b64 := base64.StdEncoding.EncodeToString([]byte(name))
					test(c, header("Name-Bin", b64), metadata.MD{
						"Name": []string{name},
					})
				})
				Convey("Fails", func() {
					c2, err := parseHeader(c, header("Name-Bin", "zzz"))
					So(c2, ShouldEqual, c)
					So(err, ShouldErrLike, "Name-Bin header: illegal base64 data at input byte 0")
				})
			})
		})
	})
}
