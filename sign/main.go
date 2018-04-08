package main

import (
	"net/url"
	"sign/util"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	values := url.Values{}
	values.Set("action", "getuser")
	values.Set("account", "yanbin")
	values.Set("sid", "1448420002")
	values.Set("uuzu_op_id", "592")
	values.Set("verify", "13d943ab3d9f42d97a9bed4f21f39183")
	values.Set("game_money", "200")
	values.Set("order_id", "1519983153536541")
	values.Set("time", "1520232897")
	values.Set("u_money", "4.99")

	if !util.CheckAuth(values, "CejSh4i7kifA1NZn") {
		fmt.Println("111")
	}

	//http://10.16.40.249:4007/gm/uuzu/gta/charge?&account=yanbin&game_money=200&order_id=1519983153536541&sid=1448420002&time=1520232897&u_money=4.99&verify=13d943ab3d9f42d97a9bed4f21f39183&uuzu_op_id=592
//http://10.16.40.249:4007/gm/uuzu/gta/role?action=getuser&account=yanbin&time=1520232897&sid=1448420002&uuzu_op_id=592&verify=fd32df4007c224993e584a23b01d4152
}
