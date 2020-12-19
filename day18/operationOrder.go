package day18

import (
	"fmt"
	"github.com/a-bleier/aoc2020/fileio"
	"regexp"
	"strconv"
	"unicode"
)

var p1 int //precedence addition
var p2 int //precedence multiplication

func Run() {
	lines := fileio.GetLinesFromFile("day18/input.txt")
	sum := 0
	p1, p2 = 0, 0
	for _, l := range lines {
		postFixTokens := ShuntingYard(tokenizeInfixExpression(remvoveBlanks(l)))
		sum += evaluatePostfix(postFixTokens)
	}
	fmt.Println("A: ", sum)

	sum = 0
	p1, p2 = 1, 0
	for _, l := range lines {
		postFixTokens := ShuntingYard(tokenizeInfixExpression(remvoveBlanks(l)))
		sum += evaluatePostfix(postFixTokens)
	}
	fmt.Println("B: ", sum)

}

func evaluatePostfix(postfix []string) int {
	var stack = make([]int, 0)
	for _, t := range postfix {
		if isNumber(t) {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		} else if t == "+" {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		} else if t == "*" {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		}
	}
	return stack[0]
}

func ShuntingYard(tokensInfix []string) []string {
	var tokensPostfix []string
	var stack []string

	for _, token := range tokensInfix {

		topTokenIndex := len(stack) - 1
		if isNumber(token) {
			tokensPostfix = append(tokensPostfix, token)
		} else if token == "," {
			for stack[topTokenIndex] != "(" {
				tokensPostfix = append(tokensPostfix, stack[topTokenIndex])
				stack = stack[:topTokenIndex]
				topTokenIndex--
			}
		} else if isOperator(token) {
			for len(stack) > 0 && isOperator(stack[topTokenIndex]) && precedence(token) <= precedence(stack[topTokenIndex]) {
				tokensPostfix = append(tokensPostfix, stack[topTokenIndex])
				stack = stack[:topTokenIndex]
				topTokenIndex--
			}
			stack = append(stack, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for stack[topTokenIndex] != "(" {
				tokensPostfix = append(tokensPostfix, stack[topTokenIndex])
				stack = stack[:topTokenIndex]
				topTokenIndex--
			}
			//Removes "("
			stack = stack[:topTokenIndex]
			topTokenIndex--
		}

	}

	//empty stack

	for len(stack) != 0 {
		tokensPostfix = append(tokensPostfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return tokensPostfix
}
func isNumber(expr string) bool {
	for _, ch := range expr {
		if !(unicode.IsDigit(ch)) && ch != '.' {
			return false
		}
	}
	return true
}

func isOperator(token string) bool {
	//TODO Write operator lookup
	switch token {
	case
		"+",
		"-",
		"*",
		"/",
		"^":
		return true

	}
	return false
}
func tokenizeInfixExpression(expr string) []string {
	var tokens []string
	var pattern string
	pattern = "(\\b\\w*[\\.]?\\w+\\b|[\\(\\)\\+\\*\\-\\/\\,\\^])"
	regexpInfix, _ := regexp.Compile(pattern)
	tokens = regexpInfix.FindAllString(expr, -1)
	return tokens
}

func remvoveBlanks(rawExpr string) string {
	var newString string
	for i := 0; i < len(rawExpr); i++ {
		if rawExpr[i] != ' ' {
			newString = newString + string(rawExpr[i])
		}
	}
	return newString
}
func precedence(ch string) int {
	if ch == "+" || ch == "-" {
		return p1
	} else if ch == "*" || ch == "/" {
		return p2
	}
	return 0
}
