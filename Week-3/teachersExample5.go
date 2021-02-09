package week3
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string] int
}
func (db *singletonDatabase) GetPopulation(name string) int  {
	return db.capitals[name]
}
func readData(path string) (map[string]int, error) {
	file, err := os.Open( path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := map[string]int{}
	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}
var once sync.Once
var instance *singletonDatabase
func GetSingletonDatabase() *singletonDatabase  {
	once.Do(func() {
		caps,_ := readData("./capitals.txt")
		db := singletonDatabase{caps}
		instance = &db
	})
	return instance
}

func exa() {
	db := GetSingletonDatabase()
	population := db.GetPopulation("Osaka")
	fmt.Println("Population of Osaka is",population)
}