package test

import (
	"fmt"
	"github.com/leverwwz/go-substrate-rpc-client/types"
	"github.com/leverwwz/go-substrate/client"
	"github.com/leverwwz/go-substrate/expand"
	"github.com/leverwwz/go-substrate/expand/base"
	"github.com/leverwwz/go-substrate/expand/bifrost"
	"reflect"
	"strings"
	"testing"
)

/*
https://polkadot.js.org/apps/#/explorer, wss://api.crust.network/
*/
func Test_Chain(t *testing.T) {
	c, err := client.New("wss://bifrost-rpc.liebi.com/ws")
	if err != nil {
		t.Fatal(err)
	}

	//b := polkadot.PolkadotEventRecords{}
	b := bifrost.BifrostEventRecords{}
	existMap := getEventTypesFieldName(b)
	fmt.Println(c.ChainName)
	fmt.Println(c.Meta.Version)
	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if IsExist(typeName, existMap) {
					continue
				}
				fmt.Printf("%s		[]Event%s%s\n", typeName, mod.Name, event.Name)
				if len(event.Args) == 0 {
					fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
				} else {
					as := ""
					for _, arg := range event.Args {
						s := fmt.Sprintf("	%s    types.\n", arg)
						as = as + s
					}
					fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n%v	Topics []types.Hash\n}\n", mod.Name, event.Name, as)
				}

				//fmt.Println(event.Args)
				fmt.Println("------------------------------------------------")
			}
		}
	}
}

func Test_Chain2(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {
				for _, arg := range event.Args {
					a := string(arg)
					if a[len(a)-1:] == ">" {
						typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
						fmt.Printf("%s		[]Event%s%s\n", typeName, mod.Name, event.Name)
						fmt.Println(arg)
						fmt.Println("------------------------------------------------")
					}
				}


			}
		}
	}
}

func Test_SubGameChain(t *testing.T) {
	c, err := client.New("wss://mainnet.subgame.org")
	if err != nil {
		t.Fatal(err)
	}

	//b := polkadot.PolkadotEventRecords{}
	b := bifrost.BifrostEventRecords{}
	existMap := getEventTypesFieldName(b)
	fmt.Println(c.ChainName)
	fmt.Println(c.Meta.Version)
	for _, mod := range c.Meta.AsMetadataV12.Modules {
		if mod.HasEvents {
			for _, event := range mod.Events {
				typeName := fmt.Sprintf("%s_%s", mod.Name, event.Name)
				if IsExist(typeName, existMap) {
					continue
				}
				fmt.Printf("%s		[]Event%s%s\n", typeName, mod.Name, event.Name)
				if len(event.Args) == 0 {
					fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
				} else {
					as := ""
					for _, arg := range event.Args {
						s := fmt.Sprintf("	%s    types.\n", arg)
						as = as + s
					}
					fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n%v	Topics []types.Hash\n}\n", mod.Name, event.Name, as)
				}

				//fmt.Println(event.Args)
				fmt.Println("------------------------------------------------")
			}
		}
	}
}

func Test_New(t *testing.T) {
	//b := polkadot.PolkadotEventRecords{}
	////tp:=reflect.TypeOf(b)
	////fmt.Println(tp.NumField())
	////fmt.Println(tp.Field(1).Name)
	//getEventTypesFieldName(b)

	s := "sp_std::marker::PhantomData<(AccountId, Event)>"
	namespaceSlice := strings.Split(s, "::")
	fmt.Println(len(namespaceSlice))
	fmt.Println(namespaceSlice)
}

func getEventTypesFieldName(ier expand.IEventRecords) []string {
	var existMap []string
	//first types.EventRecords
	te := types.EventRecords{}
	tep := reflect.TypeOf(te)
	for i := 0; i < tep.NumField(); i++ {
		existMap = append(existMap, tep.Field(i).Name)
	}

	// second
	be := base.BaseEventRecords{}
	bep := reflect.TypeOf(be)
	for i := 0; i < bep.NumField(); i++ {
		if bep.Field(i).Name == "EventRecords" {
			continue
		}
		existMap = append(existMap, bep.Field(i).Name)
	}
	// third parse IEventRecords
	ierp := reflect.TypeOf(ier)
	for i := 0; i < ierp.NumField(); i++ {
		if ierp.Field(i).Name == "BaseEventRecords" {
			continue
		}

		existMap = append(existMap, ierp.Field(i).Name)
	}
	return existMap
}

func IsExist(typeName string, existTypes []string) bool {
	for _, v := range existTypes {
		if typeName == v {
			return true
		}
	}
	return false
}

func Test_GetAllType(t *testing.T) {
	// kList, err := getAllTypes("wss://kusama-rpc.polkadot.io", 13)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	pList, err := getAllTypes("wss://rpc.polkadot.io", 12)
	if err != nil {
		t.Fatal(err)
	}
	sList, err := getAllTypes("wss://mainnet.subgame.org", 12)
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(pList)
	//fmt.Println(kList)
	//fmt.Println(sList)
	var sameList []string
	for _, p := range pList {
		haveSame := false
		for _, k := range sList {
			if p == k {
				haveSame = true
				break
			}
		}
		if haveSame == true {
			sameList = append(sameList, p)
		}
	}
	for _, s := range sameList {
		fmt.Println(s)
	}
}

func getAllTypes(url string, version int) ([]string, error) {
	c, err := client.New(url)
	if err != nil {
		return nil, err
	}
	var tmpLists []string
	if version == 13 {
		for _, mod := range c.Meta.AsMetadataV13.Modules {

			if mod.HasEvents {
				for _, event := range mod.Events {

					if len(event.Args) == 0 {
						//fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
					} else {
						//as:=""
						for _, arg := range event.Args {
							norPrint := false
							for _, as := range tmpLists {
								if as == string(arg) {
									norPrint = true
								}
							}
							if norPrint {
								continue
							}
							tmpLists = append(tmpLists, string(arg))
						}

					}
				}
			}
		}
	} else {
		for _, mod := range c.Meta.AsMetadataV12.Modules {

			if mod.HasEvents {
				for _, event := range mod.Events {

					if len(event.Args) == 0 {
						//fmt.Printf("type Event%s%s struct { \n	Phase    types.Phase\n	\n	Topics []types.Hash\n}\n", mod.Name, event.Name)
					} else {
						//as:=""
						for _, arg := range event.Args {
							norPrint := false
							for _, as := range tmpLists {
								if as == string(arg) {
									norPrint = true
								}
							}
							if norPrint {
								continue
							}
							tmpLists = append(tmpLists, string(arg))
						}

					}
				}
			}
		}
	}

	return tmpLists, nil
}

type ChainTypes struct {
	ChainName string                   `json:"chain_name"`
	Modules   []map[string]interface{} `json:"modules"`
}

func ConvertType(name string) string {

	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "T::", "")
	name = strings.ReplaceAll(name, "VecDeque<", "Vec<")
	name = strings.ReplaceAll(name, "<T>", "")
	name = strings.ReplaceAll(name, "<T as Trait>::", "")
	name = strings.ReplaceAll(name, "<T, I>", "")
	name = strings.ReplaceAll(name, "\n", " ")
	name = strings.ReplaceAll(name, `&'static[u8]`, "Bytes")

	switch name {
	case "()", "<InherentOfflineReport as InherentOfflineReport>::Inherent":
		name = "Null"
	case "Vec<u8>":
		name = "Bytes"
	case "<Lookup as StaticLookup>::Source":
		name = "Address"
	case "<Balance as HasCompact>::Type":
		name = "Compact<Balance>"
	case "<BlockNumber as HasCompact>::Type":
		name = "Compact<BlockNumber>"
	case "<Moment as HasCompact>::Type":
		name = "Compact<Moment>"
	case "<T as Trait<I>>::Proposal":
		name = "Proposal"
	case "wasm::PrefabWasmModule":
		name = "PrefabWasmModule"
	}
	if strings.Contains(name, "::") {
		subNames := strings.Split(name, "::")
		name = subNames[len(subNames)-1]
	}
	return name
}
