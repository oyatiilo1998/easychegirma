import { FormControl, InputLabel, MenuItem, Select } from "@mui/material"
import { id } from "date-fns/locale"
import _ from 'lodash'

const FSelect = ({formik, name, label, width, options, ...props}) => {
  return (
    <FormControl style={{width, marginBottom: '20px'}} >
      <InputLabel id={id}>{label}</InputLabel>
      <Select
        labelId={id}
        value={_.get(formik.values, name)}
        label={label}
        size="small"
        fullWidth
        onChange={(e) => formik.setFieldValue(name, e.target.value)}
      >
        {
          options.map(option => (
            <MenuItem value={option.value} >{option.label}</MenuItem>
          ))
        } 
      </Select>
    </FormControl>
  )
}

export default FSelect
