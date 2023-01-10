package main

import (
	"fmt"
)

type Pig struct {
	pigHealth int
}

type Op struct {
	a int
	b int
}

const mod = 998244353

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	// кол-во итераций
	pigs := make([]*Pig, 0)

	ops := make([]Op, 0)

	for i := 0; i < 12; i++ {
		var a int
		var b int
		fmt.Scanf("%d %d\n", &a, &b)

		ops = append(ops, Op{a: a, b: b})

		switch ops[i].a {
		case 1:
			pig := &Pig{pigHealth: ops[i].b}
			pigs = append(pigs, pig)
		case 2:
			for _, pig := range pigs {
				pig.pigHealth -= ops[i].b
			}
		case 3:
			for j := 0; j < i; j++ {
				switch ops[j].a {
				case 1:
					pig := &Pig{pigHealth: ops[j].b}
					pigs = append(pigs, pig)
				case 2:
					for _, pig := range pigs {
						pig.pigHealth -= ops[j].b
					}
				}
			}

		}

	}
	var count int
	for _, v := range pigs {
		if v.pigHealth > 0 {
			count++
		}
	}
	fmt.Println(pigs)
	fmt.Println("Number of live pigs:", count%mod)
	fmt.Println("Выжили:", count)
}
