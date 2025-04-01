#!/usr/bin/env bash

curl 'https://11.80.20.1/data?set=pwState:0' \
  -X 'POST' \
  -H 'Accept: */*' \
  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 0' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -b '-http-session-=6::http.session::7abfb4128b5b1c9bbde8cb4cf6bcffdb; tokenvalue=fb9b71d01b6fe7a5dd2677bbaa82e028; sysidledicon=ledIcon%20grayLed' \
  -H 'Origin: https://11.80.20.1' \
  -H 'Referer: https://11.80.20.1/sysSummaryData.html' \
  -H 'ST2: 29e7f64a84d8a7da12d8b7e8b1c48681' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0' \
  -H 'sec-ch-ua: "Chromium";v="134", "Not:A-Brand";v="24", "Microsoft Edge";v="134"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --insecure

