package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const loadNerve = 200

type propertyFeatures struct {
	load         int
	price        int
	propertyname string
}
type propertyRespons struct {
	loadRespons         int
	priceRespons        int
	averageWeight       float64
	propertynameRespons string
}

func transportNerve(transport []propertyRespons) []propertyRespons {
	amountCarried := 0
	var goodsReceived []propertyRespons
	for _, v := range transport {
		if amountCarried != loadNerve {
			amountCarried += v.loadRespons
			goodsReceived = append(goodsReceived, v)
		}
	}
	return goodsReceived
}

func averageCalculation(calcultion []propertyRespons) []propertyRespons {

	goodsMap := make(map[propertyRespons]float64)
	for _, p := range calcultion {
		goodsMap[propertyRespons{loadRespons: p.loadRespons,
			priceRespons:        p.priceRespons,
			propertynameRespons: p.propertynameRespons}] = p.averageWeight
	}

	var dataList []propertyRespons
	for i, v := range goodsMap {
		d := propertyRespons{loadRespons: i.loadRespons,
			priceRespons:        i.priceRespons,
			propertynameRespons: i.propertynameRespons,
			averageWeight:       v}
		dataList = append(dataList, d)
	}

	sort.SliceStable(dataList, func(i, j int) bool {
		return dataList[i].averageWeight > dataList[j].averageWeight
	})

	return dataList

}

func inverseProportion(property []propertyFeatures) []propertyRespons {

	var propertyWeight []propertyRespons
	for _, v := range property {
		var averageValue float64
		averageValue = float64(float64(v.price) / float64(v.load))
		propertyWeight = append(propertyWeight, propertyRespons{
			loadRespons:         v.load,
			priceRespons:        v.price,
			averageWeight:       averageValue,
			propertynameRespons: v.propertyname})
	}
	return propertyWeight
}

func ReadCsvFile(fileName string) ([]propertyFeatures, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var personList []propertyFeatures
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		record := strings.Split(line, ",")
		loads, _ := strconv.Atoi(record[0])
		prices, _ := strconv.Atoi(record[1])
		personList = append(personList, propertyFeatures{load: loads, price: prices, propertyname: record[2]})
	}
	return personList, nil
}
func WriteFileCsv(employee []propertyRespons, resultFileName string) error {
	file, err := os.OpenFile(resultFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	for _, h := range employee {
		s := fmt.Sprintf("%s,%v\n", h.propertynameRespons,h.averageWeight)
		_, err = dataWriter.WriteString(s)
		if err != nil {
			return err
		}
	}

	dataWriter.Flush()
	return nil
}

func main() {
	file, err := ReadCsvFile("goods.csv")
	if err != nil {
		log.Fatal(err)
	}
	inverse := inverseProportion(file)
	averega := averageCalculation(inverse)
	trans := transportNerve(averega)
	WriteFileCsv(trans, "goodsReceived")
}
