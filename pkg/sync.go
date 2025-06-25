package efw

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RulesResponse struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	// TODO
	// implement the fields of the Rule struct based on the JSON structure
	Port      int    `json:"port"`
	Protocol  string `json:"protocol"`  // e.g., tcp or udp
	Action    string `json:"action"`    // e.g., allow or deny
	Direction string `json:"direction"` // e.g., inbound or outbound
}

// Add this constant at the top of the file
const rulesURL = "https://app.staging.enclave.aws.sidechannel.com/cdn/storage/public_files/ogZM5tWyhkBJ87Xpo/original/ogZM5tWyhkBJ87Xpo.json"

// Sync retrieves the firewall rules from the remote endpoint
// and applies them to nftables, creating a new table called "efw".
// The default state is blocking, so all rules are defined for allowing traffic either inbound or outbound
func (e *EFW) Sync() error {
	// TODO
	// implement EFW.Sync()
	// this function should load the firewall rules from this endpoint: https://app.staging.enclave.aws.sidechannel.com/cdn/storage/public_files/ogZM5tWyhkBJ87Xpo/original/ogZM5tWyhkBJ87Xpo.json
	fmt.Println("TODO implement EFW.Sync()")

	// use the RulesResponse struct to unmarshal the JSON data

	// next apply the rules to nftables using the appropriate library
	// you can either use the "nft" package or a purego solution would be to use the "github.com/google/nftables" package
	// either are acceptable in this challenge
	// you will apply your rules to a new table called "efw"

	fmt.Println("Fetching firewall rules from remote endpoint...")

	// Step 1: Make HTTP GET request
	resp, err := http.Get(rulesURL)
	if err != nil {
		return fmt.Errorf("failed to fetch rules: %w", err)
	}
	defer resp.Body.Close()

	// Step 2: Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Step 3: Parse JSON into RulesResponse struct
	var rulesResp RulesResponse
	if err := json.Unmarshal(body, &rulesResp); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Step 4: Print parsed rules (for now, just debug output)
	fmt.Printf("Fetched %d rule(s):\n", len(rulesResp.Rules))
	for _, rule := range rulesResp.Rules {
		fmt.Printf("Port: %d, Protocol: %s, Action: %s, Direction: %s\n", rule.Port, rule.Protocol, rule.Action, rule.Direction)
	}

	// (We’ll add rule application logic next)
	return nil
}

// package efw

// import (
// 	"encoding/binary"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"github.com/google/nftables"
// 	"github.com/google/nftables/expr"
// 	"golang.org/x/sys/unix"
// )

// const rulesURL = "https://app.staging.enclave.aws.sidechannel.com/cdn/storage/public_files/ogZM5tWyhkBJ87Xpo/original/ogZM5tWyhkBJ87Xpo.json"

// type RulesResponse struct {
// 	Rules []Rule `json:"rules"`
// }

// type Rule struct {
// 	Port      int    `json:"port"`
// 	Protocol  string `json:"protocol"`  // "tcp" or "udp"
// 	Action    string `json:"action"`    // only "allow" for now
// 	Direction string `json:"direction"` // "inbound" or "outbound"
// }

// func (e *EFW) Sync() error {
// 	fmt.Println("Fetching firewall rules from remote endpoint...")

// 	resp, err := http.Get(rulesURL)
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch rules: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	var rulesResp RulesResponse
// 	if err := json.Unmarshal(body, &rulesResp); err != nil {
// 		return fmt.Errorf("failed to parse JSON: %w", err)
// 	}

// 	fmt.Printf("Fetched %d rule(s):\n", len(rulesResp.Rules))
// 	for _, rule := range rulesResp.Rules {
// 		fmt.Printf("  Port: %d, Protocol: %s, Action: %s, Direction: %s\n",
// 			rule.Port, rule.Protocol, rule.Action, rule.Direction)
// 	}

// 	c := &nftables.Conn{}

// 	// Reset existing table (if any)
// 	c.DelTable(&nftables.Table{
// 		Name:   "efw",
// 		Family: nftables.TableFamilyINet,
// 	})
// 	_ = c.Flush()

// 	// Create new table
// 	table := &nftables.Table{
// 		Name:   "efw",
// 		Family: nftables.TableFamilyINet,
// 	}
// 	c.AddTable(table)

// 	// Create input and output chains
// 	inputChain := &nftables.Chain{
// 		Name:     "input",
// 		Table:    table,
// 		Type:     nftables.ChainTypeFilter,
// 		Hooknum:  nftables.ChainHookInput,
// 		Priority: nftables.ChainPriorityFilter,
// 		Policy:   nftables.ChainPolicyDrop,
// 	}
// 	outputChain := &nftables.Chain{
// 		Name:     "output",
// 		Table:    table,
// 		Type:     nftables.ChainTypeFilter,
// 		Hooknum:  nftables.ChainHookOutput,
// 		Priority: nftables.ChainPriorityFilter,
// 		Policy:   nftables.ChainPolicyDrop,
// 	}
// 	c.AddChain(inputChain)
// 	c.AddChain(outputChain)

// 	// Add allow rules
// 	for _, rule := range rulesResp.Rules {
// 		if rule.Action != "allow" {
// 			continue // only allow rules are supported in this version
// 		}

// 		var chain *nftables.Chain
// 		switch rule.Direction {
// 		case "inbound":
// 			chain = inputChain
// 		case "outbound":
// 			chain = outputChain
// 		default:
// 			continue
// 		}

// 		exprs := []expr.Any{
// 			// Match IP protocol (tcp/udp)
// 			&expr.Payload{
// 				DestRegister: 1,
// 				Base:         expr.PayloadBaseNetworkHeader,
// 				Offset:       9, // protocol field
// 				Len:          1,
// 			},
// 			&expr.Cmp{
// 				Register: 1,
// 				Op:       expr.CmpOpEq,
// 				Data:     []byte{protocolToByte(rule.Protocol)},
// 			},
// 			// Match destination port
// 			&expr.Payload{
// 				DestRegister: 1,
// 				Base:         expr.PayloadBaseTransportHeader,
// 				Offset:       2, // dest port offset
// 				Len:          2,
// 			},
// 			&expr.Cmp{
// 				Register: 1,
// 				Op:       expr.CmpOpEq,
// 				Data:     portToBytes(rule.Port),
// 			},
// 			// Accept traffic
// 			&expr.Verdict{
// 				Kind: expr.VerdictAccept,
// 			},
// 		}

// 		c.AddRule(&nftables.Rule{
// 			Table: table,
// 			Chain: chain,
// 			Exprs: exprs,
// 		})
// 	}

// 	if err := c.Flush(); err != nil {
// 		return fmt.Errorf("failed to apply rules: %w", err)
// 	}

// 	fmt.Println("✅ Firewall rules applied successfully using Go-native nftables.")
// 	return nil
// }

// func protocolToByte(proto string) byte {
// 	switch proto {
// 	case "tcp":
// 		return unix.IPPROTO_TCP
// 	case "udp":
// 		return unix.IPPROTO_UDP
// 	default:
// 		return 0
// 	}
// }

// func portToBytes(port int) []byte {
// 	b := make([]byte, 2)
// 	binary.BigEndian.PutUint16(b, uint16(port))
// 	return b
// }
