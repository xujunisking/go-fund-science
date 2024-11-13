package utils

import (
	"errors"
	"strconv"
	"strings"
)

//用于解析SQL查询语句脚本

// ErrInvalidOperator describes an invalid operator error
var ErrInvalidOperator = errors.New("Invalid operator")

// Concat concatenate strings(连接字符串)
func concat(ss ...string) string {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteString(s)
	}
	return sb.String()
}

// 逻辑运算符
type LogicalOperator string

var loList = map[LogicalOperator]string{
	Equal: "=", NEqual: "!=",
	GreaterThan: ">", GreaterThanOrEqual: ">=",
	LessThan: "<", LessThanOrEqual: "<=",
	IN: "IN", NOTIN: "NOT IN",
	LIKE: "LIKE", ILIKE: "ILIKE", NLIKE: "NOT LIKE",
	BETWEEN: "BETWEEN", NBETWEEN: "NOT BETWEEN",
}

const (
	Equal              LogicalOperator = "eq"
	NEqual                             = "neq"
	GreaterThan                        = "gt"
	GreaterThanOrEqual                 = "gte"
	LessThan                           = "lt"
	LessThanOrEqual                    = "lte"
	IN                                 = "in"
	NOTIN                              = "nin"
	LIKE                               = "like"
	ILIKE                              = "ilike"
	NLIKE                              = "nlike"
	BETWEEN                            = "between"
	NBETWEEN                           = "nbetween"
)

type Clausifier interface {
	Clausify(k string, vv []string) (Condition, error)
}

// Condition describes a SQL Clause condition(SQL子句条件结构体)
type Condition struct {
	Expression string
	Variables  []interface{}
}

// Clause describe a SQL Where clause(SQL Where子句条件结构体)
type Clause struct {
	Conditions string
	Variables  []interface{}
}

func NewClause() Clause {
	return Clause{
		Conditions: "",
		Variables:  []interface{}{},
	}
}

// AddCondition adds a where clause condition to the current where clause(附加查询条件用AND连接)
func (c *Clause) AddAndCondition(cond Condition) {
	if c.Conditions == "" {
		c.Conditions = cond.Expression
	} else {
		c.Conditions = concat(c.Conditions, " AND ", cond.Expression)
	}
	c.Variables = append(c.Variables, cond.Variables...)
}

// 附加查询条件用OR连接
// 参数hasParentheses表示是否有圆括号
func (c *Clause) AddOrCondition(cond1 Condition, cond2 Condition, hasParentheses bool) {

	var expression string
	if hasParentheses {
		expression = "(" + concat(cond1.Expression, " OR ", cond2.Expression) + ")"
	} else {
		expression = concat(cond1.Expression, " OR ", cond2.Expression)
	}

	if c.Conditions == "" {
		c.Conditions = expression
	} else {
		c.Conditions = concat(c.Conditions, " AND ", expression)
	}

	c.Variables = append(c.Variables, cond1.Variables...)
	c.Variables = append(c.Variables, cond2.Variables...)
}

func (c *Clause) CreateCondition(key string, symbol LogicalOperator, value []string) {

	var Expression string
	var Variables []interface{}
	switch symbol {
	case Equal, NEqual, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
		Expression = concat(key, " ", loList[symbol], " ? ")
		Variables = []interface{}{value[0]}
	case IN, NOTIN:
		Expression = concat(key, " ", loList[symbol], " ? ")
		Variables = []interface{}{concat("(", value[0], ")")}
	case LIKE, ILIKE, NLIKE:
		Expression = concat(key, " ", loList[symbol], " ? ")
		Variables = []interface{}{concat("'%", value[0], "%'")}
	case BETWEEN, NBETWEEN:
		Expression = concat(key, " ", loList[symbol], " ? AND ? ")
		Variables = []interface{}{concat("'", value[0], "'"), concat("'", value[1], "'")}
	}

	if c.Conditions == "" {
		c.Conditions = Expression
	} else {
		c.Conditions = concat(c.Conditions, "AND ", Expression)
	}

	c.Variables = append(c.Variables, Variables...)
}

func (c *Clause) BuildSQLStatement() string {
	if c.Conditions == "" {
		return ""
	}

	strSQL := c.Conditions
	for _, v := range c.Variables {
		strSQL = strings.Replace(strSQL, "?", v.(string), 1)
	}
	return strSQL
}

// SQL转意字符映射
var operators = map[string]string{
	"eq": "=", "neq": "!=",
	"gt": ">", "gte": ">=",
	"lt": "<", "lte": "<=",
	"in": "IN", "nin": "NOT IN",
	"like": "LIKE", "ilike": "ILIKE", "nlike": "NOT LIKE",
	"between": "BETWEEN", "nbetween": "NOT BETWEEN",
}

// QSClausifier is the default clausifier
type QSClausifier struct {
	Separator   string
	Placeholder string
	Operators   map[string]string
}

// GetOperator returns the operator key
func (c QSClausifier) GetOperator(k string) (string, string) {
	op := strings.Split(k, c.Separator)
	if len(op) == 2 {
		return op[0], op[1]
	}
	return k, "eq"
}

// BuildCondition return condition variables with the right type
func (c QSClausifier) BuildCondition(k string, o string, v string) Condition {
	cond := Condition{}
	var nv []interface{}
	for _, e := range strings.Split(v, ",") {
		if val, err := strconv.Atoi(e); err == nil {
			nv = append(nv, val)
			continue
		}
		if strings.Contains(o, "LIKE") {
			e = concat("%", e, "%")
		}
		nv = append(nv, e)
	}
	// edge cases
	if strings.Contains(o, "IN") {
		cond.Expression = concat(k, " ", o, " ", c.Placeholder)
		cond.Variables = append(cond.Variables, nv)
	} else {
		if strings.Contains(o, "BETWEEN") {
			cond.Expression = concat(k, " ", o, " ", c.Placeholder, " AND ", c.Placeholder)
		} else {
			cond.Expression = concat(k, " ", o, " ", c.Placeholder)
		}
		cond.Variables = append(cond.Variables, nv...)
	}
	return cond
}

// Clausify is the
func (c QSClausifier) Clausify(k string, vv []string) (Condition, error) {
	cond := Condition{}
	k, op := c.GetOperator(k)
	//映射取值(成功v为值ok为true;失败v为空字符串ok为false)
	// v, _ := c.Operators["op"]
	// fmt.Println(v == "")
	if _, ok := c.Operators[op]; !ok {
		return cond, ErrInvalidOperator
	}
	return c.BuildCondition(k, c.Operators[op], vv[0]), nil
}

// With tuns url.Query into where clause condtion by passing a custom operator
func With(q map[string][]string, cf Clausifier) (Clause, error) {
	c := Clause{}
	for k, v := range q {
		cond, err := cf.Clausify(k, v)
		if err != nil {
			return c, err
		}
		c.AddAndCondition(cond)
	}
	return c, nil
}

// Clausify takes an url.Query and turns it into a Where clause conditions
func Clausify(q map[string][]string) (Clause, error) {
	return With(q, QSClausifier{
		Separator: "__", Operators: operators, Placeholder: "?"})
}
