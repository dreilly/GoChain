package chain

import (
	"encoding/json"
	"errors"
	"fmt"
	"gochain/config"
	"gochain/utility"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Chain : the shape of what is returned
// should have a name corresponding to the file it's stored in
// and the actual chain consisting of links
type Chain struct {
	Name       string
	ChainLinks []link
	MetaData   chainMetaData
}

type link struct {
	Date   time.Time
	Symbol rune
}

type chainMetaData struct {
	Description  string
	CreationDate time.Time
}

// PrintChain : prints this chain
func (chain *Chain) PrintChain(name string, detail bool) error {
	c := config.Configuration{}
	c.GetConfiguration()
	chainPath := filepath.Join(c.ChainDirectory, "Chains", name+".chain")
	file, err := os.Open(chainPath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(chain)
	if decodeErr != nil {
		return decodeErr
	}

	if detail {
		printVerbose(chain.ChainLinks)
	} else {
		printSimple(chain.ChainLinks)
	}

	return nil
}

func printSimple(links []link) {
	for _, c := range links {
		fmt.Printf("[%c]", c.Symbol)
	}
	fmt.Println()
}

func printVerbose(links []link) {
	for _, c := range links {
		fmt.Printf("\tDate: \t%s", c.Date.String())
		fmt.Println()
		linked := c.Symbol == 'X'
		fmt.Printf("\tLinked: %t", linked)
		fmt.Println()
		fmt.Println()
	}
	fmt.Println()
}

// GetChain : returns a chain with a given name
func (chain *Chain) GetChain(name string) (myErr error) {
	c := config.Configuration{}
	c.GetConfiguration()
	chainPath := filepath.Join(c.ChainDirectory, "Chains", name+".chain")
	file, err := os.Open(chainPath)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(chain)
	if decodeErr != nil {
		return decodeErr
	}

	return
}

// CreateChain : create a chain with a given name
// will create a chain directory if none exists
// directory will default to user's home directory
// or if the ~/gochain.json exists will use the
// defaultDirectory value from there
func CreateChain(name string) error {
	chain := Chain{}
	chain.Name = name
	createChainDir()
	_, err := createChain(name)

	return err
}

func createChainDir() error {
	config := config.Configuration{}
	config.GetConfiguration()

	dirPath := filepath.Join(config.ChainDirectory, "Chains")
	if exist, err := objExists(dirPath); !exist {
		if err == nil {
			os.MkdirAll(dirPath, os.ModePerm)
		} else {
			return err
		}
	} else {
		return errors.New("directory creation failed")
	}

	return nil
}

func objExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func createChain(name string) (bool, error) {
	config := config.Configuration{}
	config.GetConfiguration()
	chainPath := filepath.Join(config.ChainDirectory, "Chains", name+".chain")
	if exist, existErr := objExists(chainPath); !exist {
		if existErr == nil {
			newChainFile, err := os.Create(chainPath)
			if err != nil {
				return false, err
			}
			chainSkeleton, err := getNewFileLayout(name)
			newChainFile.Write(chainSkeleton)
			newChainFile.Sync()
			newChainFile.Close()
		} else {
			return false, existErr
		}
	} else {
		return false, errors.New("a chain with name already exists")
	}

	return true, nil
}

func getNewFileLayout(name string) ([]byte, error) {
	l := make([]link, 1, 1)
	t := time.Now()
	l[0].Date = t
	l[0].Symbol = ' '

	chain := Chain{
		name,
		l,
		chainMetaData{
			"This is my new chain!",
			t,
		},
	}
	jm, err := json.Marshal(chain)
	if err != nil {
		return nil, err
	}

	return jm, nil
}

// GetAllChains : returns list of all chains in chain dir
func GetAllChains() (list []string, readError error) {
	config := config.Configuration{}
	config.GetConfiguration()
	chainPath := filepath.Join(config.ChainDirectory, "Chains")
	files, err := ioutil.ReadDir(chainPath)
	if err != nil {
		return nil, err
	}
	fileNames := make([]string, len(files), len(files))
	for i, f := range files {
		basename := f.Name()
		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		fileNames[i] = name
	}

	return fileNames, nil
}

// CreateLink : creates link on current chain
func (chain *Chain) CreateLink(name string) error {
	chain.GetChain(name)
	t := time.Now()
	lastLinkIndex := len(chain.ChainLinks) - 1

	xTime := chain.ChainLinks[lastLinkIndex].Date
	yTime := utility.StripTimeValues(xTime)
	ddiff := getDaysSince(yTime)

	emptyLinksToFill := 0
	if ddiff > 1 {
		emptyLinksToFill = ddiff - 1
	}
	if emptyLinksToFill > 0 {
		for i := 0; i < emptyLinksToFill; i++ {
			chain.ChainLinks = append(chain.ChainLinks, link{t, ' '})
		}
	}
	if ddiff != 0 {
		chain.ChainLinks = append(chain.ChainLinks, link{t, 'X'})
	} else {
		if chain.ChainLinks[lastLinkIndex].Symbol == 'X' {
			fmt.Println("Link for today already exists")
		} else {
			chain.ChainLinks[lastLinkIndex].Symbol = 'X'
		}
	}
	chain.writeChainToFile(name)

	return nil
}

func (chain *Chain) writeChainToFile(name string) error {
	conf := config.Configuration{}
	conf.GetConfiguration()

	chainPath := filepath.Join(conf.ChainDirectory, "Chains", name+".chain")
	if exist, err := objExists(chainPath); exist {
		if err != nil {
			return err
		}
		jm, jErr := json.Marshal(chain)
		if jErr != nil {
			return jErr
		}
		wErr := ioutil.WriteFile(chainPath, jm, 0666)
		if wErr != nil {
			return wErr
		}
	} else {
		return errors.New("Can not find chain")
	}

	return nil
}

func getDaysSince(from time.Time) int {
	_, _, d, _, _, _ := utility.Diff(from, time.Now())
	return d
}
