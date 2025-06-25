package efw

import (
	"fmt"

	"github.com/google/nftables"
	"github.com/google/nftables/expr"
)

// Status retrieves the status of the EFW
// and outputs it to the console
func (e *EFW) Status() {
	// TODO
	// implement EFW.Status()
	// this function should retrieve the status of the EFW
	// and the current state of the firewall
	// you will want to represent the output of the efw table
	fmt.Println("TODO implement EFW.Status()")
	c := &nftables.Conn{}

	// List all tables
	tables, err := c.ListTables()
	if err != nil {
		fmt.Printf("Error retrieving tables: %v\n", err)
		return
	}

	for _, table := range tables {
		if table.Name != "efw" {
			continue
		}

		fmt.Printf("Table: %s (Family: %s)\n", table.Name, table.Family.String())

		// List chains in this table
		chains, err := c.ListChains()
		if err != nil {
			fmt.Printf("Error retrieving chains: %v\n", err)
			return
		}

		for _, chain := range chains {
			if chain.Table.Name != "efw" {
				continue
			}
			fmt.Printf("  Chain: %s (Hook: %v, Policy: %v)\n", chain.Name, chain.Hooknum, chain.Policy)

			// List rules in this chain
			rules, err := c.GetRules(table, chain)
			if err != nil {
				fmt.Printf("    Error retrieving rules: %v\n", err)
				continue
			}

			for i, rule := range rules {
				fmt.Printf("    Rule #%d:\n", i+1)
				for _, e := range rule.Exprs {
					switch exp := e.(type) {
					case *expr.Payload:
						fmt.Printf("      Match: Payload (base=%v, offset=%d, len=%d)\n", exp.Base, exp.Offset, exp.Len)
					case *expr.Cmp:
						fmt.Printf("      Match: Compare register=%d, op=%v, data=%v\n", exp.Register, exp.Op, exp.Data)
					case *expr.Verdict:
						fmt.Printf("      Verdict: %s\n", verdictName(exp.Kind))
					default:
						fmt.Printf("      Expr: %T\n", exp)
					}
				}
			}
		}
	}
}
