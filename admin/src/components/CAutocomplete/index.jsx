import { Autocomplete, FormControl, InputLabel, MenuItem } from "@mui/material"

const CAutocomplete = ({value, onChange = () => {}, label, options, id, variant, required, ...props}) => {

  return (
    <FormControl fullWidth>
      <InputLabel required={required} size="small" id={'CSelect-' + id + '-label'}>{label}</InputLabel>
      <Autocomplete
        labelId={'CSelect-' + id + '-label'}
        value={value}
        label={label}
        onChange={onChange}
        variant={variant}
        size="small"
        {...props}
      >
         <MenuItem value={null}>---</MenuItem>
        {
          options?.map(option => (
            <MenuItem value={option.value}>{option.label}</MenuItem>
          ))
        }
      </Autocomplete>
    </FormControl>
  )
}

export default CAutocomplete
