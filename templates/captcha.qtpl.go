// This file is automatically generated by qtc from "captcha.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line captcha.qtpl:1
package templates

//line captcha.qtpl:1
import "github.com/bakape/captchouli/common"

//line captcha.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line captcha.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line captcha.qtpl:3
func StreamCaptcha(qw422016 *qt422016.Writer, colour, background, tag string, id [64]byte, images [9][16]byte) {
	//line captcha.qtpl:3
	qw422016.N().S(`<style>.captchouli-checkbox {display: none;}.captchouli-checkbox:checked ~ .captchouli-img {transform: scale(0.8);}.captchouli-img {margin: 2px;user-select: none;}.captchouli-width {width: 462px;}.captchouli-margin {margin: 4px 0;}</style><form method="post" class="captchouli-width" style="height:525px; background:`)
	//line captcha.qtpl:22
	qw422016.E().S(background)
	//line captcha.qtpl:22
	qw422016.N().S(`; color:`)
	//line captcha.qtpl:22
	qw422016.E().S(colour)
	//line captcha.qtpl:22
	qw422016.N().S(`; font-family:Sans-Serif;"><input type="text" name="`)
	//line captcha.qtpl:23
	qw422016.N().S(common.IDKey)
	//line captcha.qtpl:23
	qw422016.N().S(`" hidden value="`)
	//line captcha.qtpl:23
	streamencodeID(qw422016, id)
	//line captcha.qtpl:23
	qw422016.N().S(`"><input type="text" name="`)
	//line captcha.qtpl:24
	qw422016.N().S(common.ColourKey)
	//line captcha.qtpl:24
	qw422016.N().S(`" hidden value="`)
	//line captcha.qtpl:24
	qw422016.E().S(colour)
	//line captcha.qtpl:24
	qw422016.N().S(`"><input type="text" name="`)
	//line captcha.qtpl:25
	qw422016.N().S(common.BackgroundKey)
	//line captcha.qtpl:25
	qw422016.N().S(`" hidden value="`)
	//line captcha.qtpl:25
	qw422016.E().S(background)
	//line captcha.qtpl:25
	qw422016.N().S(`"><header class="captchouli-width captchouli-margin" style="text-align:center; font-size:130%; overflow:auto;">Select all images of <b>`)
	//line captcha.qtpl:27
	qw422016.E().S(tag)
	//line captcha.qtpl:27
	qw422016.N().S(`</b></header><div class="captchouli-width" style="height:462px;">`)
	//line captcha.qtpl:30
	buf := make([]byte, 4096)

	//line captcha.qtpl:31
	for i, img := range images {
		//line captcha.qtpl:31
		qw422016.N().S(`<label><input type="checkbox" name="captchouli-`)
		//line captcha.qtpl:33
		qw422016.N().D(i)
		//line captcha.qtpl:33
		qw422016.N().S(`" class="captchouli-checkbox"><img class="captchouli-img" draggable="false" src="`)
		//line captcha.qtpl:34
		streamthumbnail(qw422016, img, buf)
		//line captcha.qtpl:34
		qw422016.N().S(`"></label>`)
		//line captcha.qtpl:36
	}
	//line captcha.qtpl:36
	qw422016.N().S(`</div><input type="submit" class="captchouli-width captchouli-margin"></form>`)
//line captcha.qtpl:40
}

//line captcha.qtpl:40
func WriteCaptcha(qq422016 qtio422016.Writer, colour, background, tag string, id [64]byte, images [9][16]byte) {
	//line captcha.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line captcha.qtpl:40
	StreamCaptcha(qw422016, colour, background, tag, id, images)
	//line captcha.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line captcha.qtpl:40
}

//line captcha.qtpl:40
func Captcha(colour, background, tag string, id [64]byte, images [9][16]byte) string {
	//line captcha.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
	//line captcha.qtpl:40
	WriteCaptcha(qb422016, colour, background, tag, id, images)
	//line captcha.qtpl:40
	qs422016 := string(qb422016.B)
	//line captcha.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
	//line captcha.qtpl:40
	return qs422016
//line captcha.qtpl:40
}
