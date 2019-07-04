package networkutil

import (
	"fmt"
	"testing"
)
func TestTable(t *testing.T) {
	table := Table(GetLocalIp(4))
	fmt.Println(table)
}
func TestGetLocalIpv4(t *testing.T) {

}
