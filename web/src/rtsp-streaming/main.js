const RTSP2web = require('rtsp2web')

let port = 9999

new RTSP2web({
  port,
  audio: false,
})
