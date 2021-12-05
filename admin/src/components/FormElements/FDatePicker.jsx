import { DatePicker } from "@mui/lab"
import { TextField } from "@mui/material"
import _ from "lodash"

const FDatePicker = ({formik, name, label, width, inputProps, ...props}) => {

  return (
    <DatePicker
      inputFormat="dd.MM.yyyy"
      mask="__.__.____"
      toolbarFormat="dd.MM.yyyy"
      value={_.get(formik.values, name)}
      name={name}
      onChange={value => formik.setFieldValue(name, value)}
      {...props}
      renderInput={(params) => (
        <TextField 
          {...params}
          style={{ width }}
          size="small"
          error={_.get(formik.touched, name) && Boolean(_.get(formik.errors, name))}
          helperText={(_.get(formik.touched, name) && _.get(formik.errors, name)) ?? " "}
          {...inputProps}
          label={label}
        />
      )}
    />
  )
}

export default FDatePicker
