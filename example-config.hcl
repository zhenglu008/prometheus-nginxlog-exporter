port = 4040

namespace "nginx" {
  source_files = [
    "test.log",
    "foo.log"
  ]
  format = "$remote_addr - $remote_user [$time_local] \"$request\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\" \"$http_x_forwarded_for\""
}