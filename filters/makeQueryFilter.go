package filters

import (
	"encoding/json"
	"fmt"
	"golang-examples/manipulation"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func makeFilters(data interface{}) map[string]string {
	var dataMap map[string]interface{}
	dataByte, _ := json.Marshal(data)
	json.Unmarshal(dataByte, &dataMap)

	dataFilters := make(map[string]string)
	for k, _ := range dataMap {
		fake := manipulation.SnakeCaseToCamelCase(k)

		dataFilters[fake] = k
	}

	return dataFilters
}

func MakeQueryFilters(filters map[string]interface{}, data interface{}) (queryMod []qm.QueryMod) {
	allowedFilters := makeFilters(data)

	for filterKey, filterVal := range filters {
		if realFilterName, ok := allowedFilters[filterKey]; ok {
			filterValList := manipulation.InterfaceToInterfaceSlice(filterVal)
			queryMod = append(queryMod, qm.AndIn(fmt.Sprintf("%s IN ?", realFilterName), filterValList...))
		}
	}
	return
}
