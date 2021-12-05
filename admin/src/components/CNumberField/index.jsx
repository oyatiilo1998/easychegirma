import { TextField } from "@mui/material"

const CNumberField = ({ value, onChange = () => {}, ...props }) => {
  const changeHandler = (e) => {
    const val = e.target.value
    if (/^[\d.,:]*$/.test(val)) onChange(val)
  }

  const blurHandler = (e) => {
    const val = e.target.value
    onChange(Number(val.replace(",", ".")) || 0)
  }

  return (
    <TextField   
      {...props}
      value={value}
      onChange={changeHandler}
      onBlur={blurHandler}
    />
  )
}

export default CNumberField
