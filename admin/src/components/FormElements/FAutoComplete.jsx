import { Autocomplete, TextField } from "@mui/material"
import _ from "lodash"
import { useState } from "react"

const options = [
  {
    label: '123',
    value: 111
  }
]

const FAutoComplete = ({formik, label, name}) => {
  const [value, setValue] = useState(111)
  
  return (
    <Autocomplete
      disablePortal
      id="combo-box-demo"
      options={options}
      sx={{ width: 300 }}
      isOptionEqualToValue={(loption, lvalue) => loption.value !== lvalue}
      value={value}
      onChange={(e, val) => console.log("Value ===> ", val)}
      renderInput={(params) => <TextField {...params} label="Movie" />}
    />
  )
}

export default FAutoComplete
