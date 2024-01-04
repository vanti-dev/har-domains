# HAR file domain extractor

This tool analyses a HAR file and extracts each unique domain name from urls mentioned in the file.
A HAR file is used by many network tools to record the network traffic of a web page.

- To install: `go install github.com/vanti-dev/har-domains@latest`
- To build: `go build .`
- To run: `har-domains <path to har file>`

To get a HAR file from Chrome: https://support.google.com/admanager/answer/10358597

1. Open Chrome DevTools (F12)
2. Go to the Network tab
3. Reload the page
4. Click the "Export HAR" button (down arrow) or right click on the list of requests and select "Save all as HAR with content"
