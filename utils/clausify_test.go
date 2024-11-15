package utils

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// 测试当前文件夹下所有文件：go test -v
// 指定测试文件命令：go test -v cal_test.go cal.go

func assert(t *testing.T, x, y any) {
	//判断两个值是否深度相等,两个变量在结构和内容上完全一致才算是深度相等
	if !reflect.DeepEqual(x, y) {
		t.Fatalf("AssertionError: %v != %v", x, y)
	}
}

func getURLQuery(uri string) map[string][]string {
	u, _ := url.Parse(uri)
	return u.Query()
}

func TestInvalidOperator(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?username__xy=josh")
	fmt.Println("q=", q)
	c, err := Clausify(q)
	fmt.Println("c=", c)
	assert(t, err.Error(), "Invalid operator")
}

func TestEqual(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?username=josh") //q= map[username:[josh]]
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, c.Conditions, "username = ?")
	assert(t, len(c.Variables), 1)
	q = getURLQuery("https://httpbin.org/?username=josh&age=30") //q= map[age:[30] username:[josh]]
	c, _ = Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, strings.Contains(c.Conditions, "username = ?"), true)
	assert(t, strings.Contains(c.Conditions, "age = ?"), true)
	assert(t, len(c.Variables), 2)
}

func TestNotEqual(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?username__neq=josh")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, c.Conditions, "username != ?")
	assert(t, len(c.Variables), 1)
	q = getURLQuery("https://httpbin.org/?username__neq=josh&age__neq=30")
	c, _ = Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, strings.Contains(c.Conditions, "username != ?"), true)
	assert(t, strings.Contains(c.Conditions, "age != ?"), true)
	assert(t, len(c.Variables), 2)
}

func TestGreaterThan(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__gt=15&name=book")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price > ?"))
	assert(t, true, strings.Contains(c.Conditions, "name = ?"))
	assert(t, len(c.Variables), 2)
}

func TestGreaterThanEqual(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__gte=15&name=book")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price >= ?"))
	assert(t, true, strings.Contains(c.Conditions, "name = ?"))
	assert(t, len(c.Variables), 2)
}

func TestLessThan(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__lt=15&name=book")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price < ?"))
	assert(t, true, strings.Contains(c.Conditions, "name = ?"))
	assert(t, len(c.Variables), 2)
}

func TestLessThanEqual(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__lte=15&name=book")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price <= ?"))
	assert(t, true, strings.Contains(c.Conditions, "name = ?"))
	assert(t, len(c.Variables), 2)
}

func TestLike(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__lte=15&name__like=book")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price <= ?"))
	assert(t, true, strings.Contains(c.Conditions, "name LIKE ?"))
	assert(t, len(c.Variables), 2)
}

func TestILike(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__lte=15&name__ilike=book&category=fruits")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price <= ?"))
	assert(t, true, strings.Contains(c.Conditions, "name ILIKE ?"))
	assert(t, len(c.Variables), 3)
}

func TestNotLike(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?price__lte=15&name__nlike=book")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "price <= ?"))
	assert(t, true, strings.Contains(c.Conditions, "name NOT LIKE ?"))
	assert(t, len(c.Variables), 2)
}

func TestIn(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?id__in=2,4,6")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "id IN ?"))
	assert(t, len(c.Variables), 1)
}

func TestNotIn(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?id__nin=2,4,6")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "id NOT IN ?"))
	assert(t, len(c.Variables), 1)
}

func TestBetween(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?category=fruits&price__between=10,20")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, " BETWEEN ? AND ?"))
	assert(t, len(c.Variables), 3)
}

func TestNotBetween(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?category=fruits&price__nbetween=10,20")
	c, _ := Clausify(q)
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, " NOT BETWEEN ? AND ?"))
	assert(t, len(c.Variables), 3)
}

type MyClausifier struct {
	Separator string
}

func (m MyClausifier) Clausify(k string, vv []string) (Condition, error) {
	op := strings.Split(k, m.Separator)
	if op[1] == "<>" {
		return Condition{
			Expression: concat(op[0], " <> ?"),
			Variables:  []interface{}{vv[0]},
		}, nil
	}
	return Condition{}, nil
}

func TestCustomOperator(t *testing.T) {
	q := getURLQuery("https://httpbin.org/?id-<>=1")
	c, _ := With(q, MyClausifier{Separator: "-"})
	log.Printf("%+v, %+v\n", c.Conditions, c.Variables)
	assert(t, true, strings.Contains(c.Conditions, "id <> ?"))
	assert(t, len(c.Variables), 1)
}

func TestCreateCondition(t *testing.T) {
	clause := Clause{
		Conditions: "",
		Variables:  []interface{}{},
	}

	clause.CreateCondition("name", LIKE, []string{"book"})
	clause.CreateCondition("type", IN, []string{"1,2,3,4,5"})
	clause.CreateCondition("study_date", BETWEEN, []string{"2023-11-01 00:00:00", "2024-10-31 23:59:59"})
	fmt.Println(clause)

	strSQL := clause.BuildSQLStatement()
	fmt.Println(strSQL)
}
