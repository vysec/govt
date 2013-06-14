/*
A few test cases for the `govt` package.

We cannot have many test cases, because the public API is limited to four requests per minute.
So here, we demonstrate that the scheme works, and leave it at that

Written by Willi Ballenthin while at Mandiant.
June, 2013.
*/
package govt

import "testing"

func TestPublicApi(t *testing.T) {
	var apikey string = ""
	if apikey == "" {
		t.Error("Unfortunately, you must edit the test case and provide your API key")
    return
	}

	govt := Client{Apikey: apikey}
	govt.UseDefaultUrl()

	var testMd5 string = "eeb024f2c81f0d55936fb825d21a91d6"
	report, err := govt.GetReport(testMd5)
	if err != nil {
		t.Error("Error requesting report: ", err.Error())
    return
	}
	if report.Md5 != testMd5 {
		t.Error("Requested MD5 does not match result: ", testMd5, " vs. ", report.Md5)
    return
	}
}
