package main

import "fmt"
import s "strings"
//String Function
var p = fmt.Println
func main() {
	p("Contains", s.Contains("Test", "es")) //มี es อยู่ใน Test รึป่าว
	p("Count", s.Count("Test", "t")) //นับ
	p("HasPrefix", s.HasPrefix("Lord Varder", "Lord"))
	p("HasSuffix", s.HasSuffix("Lord Varder", "Varder"))
	p("Index", s.Index("Lord Varder", "V"))
	p("Join", s.Join([]string{"a","b","c"},"-"))
	p("Repeat", s.Repeat("a", 5)) //ทำซ้ำ
	p("Replace", s.Replace("foo", "o", "0", -1))// o เป็น 0 ตำแหน่ง -1
	p("Replace", s.Replace("fooo", "o", "0", 2)) //แทน O ด้วยเลข 0 ถ้าพบ 2 แทน 2 ตัวเลย
	p("Split", s.Split("a-b-c-d-e", "-"))
	p("ToLower", s.ToLower("Nattapon"))
	p("ToUpper", s.ToUpper("Nattapon"))
	p("Len", len("Nattapon"))











}

/*
	Contains 
	Count
	HasPrefix
	HasSuffix
	Index
	Join
	Repeat
	Replace
	Split
	ToLower
	ToUpper
*/