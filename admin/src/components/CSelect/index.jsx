import { FormControl, InputLabel, MenuItem, Select } from "@mui/material"

const CSelect = ({value, onChange = () => {}, label, options, id, variant, required, error, helperText ,...props}) => {

  return (
    <FormControl fullWidth  >
      <InputLabel required={required} size="small" id={'CSelect-' + id + '-label'}>{label}</InputLabel>
      <Select
        labelId={'CSelect-' + id + '-label'}
        value={value}
        label={label}
        onChange={onChange}
        variant={variant}
        error={error}
        size="small"
        {...props}
      >
         <MenuItem value={null}>---</MenuItem>
        {
          options?.map((option, index) => (
            <MenuItem key={index} value={option.value}>{option.label}</MenuItem>
          ))
        }
      </Select>
    </FormControl>
  )
}

export default CSelect
