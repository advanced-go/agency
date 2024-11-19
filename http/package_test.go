package http

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"net/http"
)

func ExampleExchange() {
	r, _ := http.NewRequest("", "http://localhost:8083/github/advanced-go/agency:resiliency?action=start", nil)
	resp, status := Exchange(r)
	if status.OK() {
		//buf, _ := io.ReadAll(resp.Body, nil)
		fmt.Printf("test: Exchange(r) -> [status:%v] [status-code:%v] [%v]\n", status, resp.StatusCode, resp.Header.Get(core.XAuthority))
	}

	//Output:
	//test: Exchange(r) -> [status:OK] [status-code:200] [github/advanced-go/log]

}
