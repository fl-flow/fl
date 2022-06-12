package parameterparser

import (
  "fmt"
  "testing"
  "encoding/json"
  "fl/common/dag_error"
  "fl/common/parser/parameter_parser"
)


func testParse(rawParameter string) (parameterparser.Parameter, error){
  var parameter parameterparser.Parameter
  ok := json.Unmarshal([]byte(rawParameter), &parameter)
  if ok != nil {
    return parameter, &dagError.DagError{Code: 12000}
  }
  parameter1, ok1 := parameterparser.Parse(parameter)
  if ok1 != nil {
    return parameter, ok1
  }
  return parameter1, nil
}


func TestParse(t *testing.T) {
  _, parameterparserok1 := testParse(`
    {
      "common": "CCCC",
      "tasks": ""
    }
  `)
  fmt.Println(parameterparserok1)

  parameterparser, parameterparserok := testParse(`
    {
      "common": "CCCC",
      "tasks": {"a": "z", "f": "d"}
    }
  `)
  fmt.Println(parameterparserok)
  fmt.Println(parameterparser)


  // // parser unit test
  // tasks, parameter, e := parser.Parse(`
  //   {
  //     "dag": {
  //       "a": {
  //         "input": [],
  //         "output": ["d"],
  //         "cmd": "cmd"
  //       }
  //     },
  //     "parameter": {
  //       "common": "CCCC",
  //       "tasks": {"a": "z"}
  //     }
  //   }
  // `)
  // fmt.Println(e)
  // fmt.Println(tasks)
  // fmt.Println(parameter)
}
